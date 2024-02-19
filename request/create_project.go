package request

import (
	"context"

	"github.com/limit-lab/zeabur-go/types"
)

func (c *Client) CreateProject(ctx context.Context, req *types.CreateProjectRequest) (*types.CreateProjectResponse, error) {
	type wrapper struct {
		CreateProject *types.CreateProjectResponse `json:"createProject"`
	}

	resp, err := doRequest[*types.CreateProjectRequest, wrapper](ctx, c, req)
	if err != nil {
		return nil, err
	}
	return resp.CreateProject, nil
}
