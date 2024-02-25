package request

import (
	"context"

	"github.com/limit-lab/zeabur-go/types"
)

func (c *Client) GetProject(ctx context.Context, req *types.GetProjectRequest) (*types.GetProjectResponse, error) {
	type wrapper struct {
		Project *types.GetProjectResponse `json:"project"`
	}

	resp, err := doRequest[*types.GetProjectRequest, wrapper](ctx, c, req)
	if err != nil {
		return nil, err
	}
	return resp.Project, nil
}
