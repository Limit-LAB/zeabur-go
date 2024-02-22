package request

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/limit-lab/zeabur-go/types"
)

func (c *Client) Deploy(ctx context.Context, req *types.DeployRequest) (*types.DeployResponse, error) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	err := writer.WriteField("environment", req.EnvironmentID)
	if err != nil {
		return nil, err
	}
	fileWriter, err := writer.CreateFormFile("code", "code.zip")
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(fileWriter, req.CodeZip)
	if err != nil {
		return nil, err
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/projects/%s/services/%s/deploy",
		c.apiPath.String(), req.ProjectID, req.ServiceID)
	httpReq, err := http.NewRequestWithContext(ctx, "POST", url, body)
	if err != nil {
		return nil, err
	}
	httpReq.Header.Add("Content-Type", writer.FormDataContentType())
	httpReq.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	httpReq.Header.Add("Accept-Encoding", "gzip, deflate")

	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

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

	var rs types.DeployResponse
	if err := json.NewDecoder(reader).Decode(&rs); err != nil {
		return nil, err
	}

	return &rs, nil
}
