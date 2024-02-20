package request

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/limit-lab/zeabur-go/types"
)

func (c *Client) Deploy(ctx context.Context, req *types.DeployRequest) error {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	err := writer.WriteField("environment", req.EnvironmentID)
	if err != nil {
		return err
	}
	fileWriter, err := writer.CreateFormFile("code", "code.zip")
	if err != nil {
		return err
	}
	_, err = io.Copy(fileWriter, req.CodeZip)
	if err != nil {
		return err
	}
	err = writer.Close()
	if err != nil {
		return err
	}

	url := fmt.Sprintf("%s/projects/%s/services/%s/deploy",
		c.apiPath.String(), req.ProjectID, req.ServiceID)
	httpReq, err := http.NewRequestWithContext(ctx, "POST", url, body)
	if err != nil {
		return err
	}
	httpReq.Header.Add("Content-Type", writer.FormDataContentType())
	httpReq.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	httpReq.Header.Add("Accept-Encoding", "gzip, deflate")

	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}
