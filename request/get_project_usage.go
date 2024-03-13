package request

import (
	"context"

	"github.com/limit-lab/zeabur-go/types"
)

func (c *Client) GetProjectUsage(ctx context.Context, req *types.GetProjectUsageRequest) (*types.GetProjectUsageResponse, error) {
	type wrapper struct {
		ProjectUsage *types.GetProjectUsageResponse `json:"projectUsage"`
	}

	resp, err := doRequest[*types.GetProjectUsageRequest, wrapper](ctx, c, req)
	if err != nil {
		return nil, err
	}
	return resp.ProjectUsage, nil
}
