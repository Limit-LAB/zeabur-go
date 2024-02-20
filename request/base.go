package request

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/limit-lab/zeabur-go/types"
)

func doRequest[RequestT types.ToGraphQLRequest, ResponseWrapper any](
	ctx context.Context, c *Client, reqObj RequestT,
) (*ResponseWrapper, error) {
	body, err := reqObj.ToRequest()
	if err != nil {
		return nil, err
	}

	var bodyReader io.Reader
	bs, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewBuffer(bs)

	url := fmt.Sprintf("%s/graphql", c.apiPath.String())
	req, err := http.NewRequestWithContext(ctx, "POST", url, bodyReader)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	req.Header.Add("Accept-Encoding", "gzip, deflate")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed with status = %d", resp.StatusCode)
	}

	var reader io.Reader

	if resp.Header.Get("Content-Encoding") == "gzip" {
		grd, err := gzip.NewReader(resp.Body)
		if err != nil {
			return nil, err
		}
		reader = grd
	} else {
		reader = resp.Body
	}

	type data struct {
		Data *ResponseWrapper `json:"data"`
	}
	var rs data
	if err := json.NewDecoder(reader).Decode(&rs); err != nil {
		return nil, err
	}
	return rs.Data, nil
}
