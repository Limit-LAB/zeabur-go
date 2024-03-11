package request

import (
	"context"

	"github.com/limit-lab/zeabur-go/types"
)

func (c *Client) DeleteProject(ctx context.Context, req *types.DeleteProjectRequest) (*types.DeleteProjectResponse, error) {
	type wrapper struct {
		DeleteService *types.DeleteProjectResponse `json:"deleteProject"`
	}

	resp, err := doRequest[*types.DeleteProjectRequest, wrapper](ctx, c, req)
	if err != nil {
		return nil, err
	}
	return resp.DeleteService, nil
}
