package request

import (
	"context"

	"github.com/limit-lab/zeabur-go/types"
)

func (c *Client) GetDeployments(ctx context.Context, req *types.GetDeploymentsRequest) (*types.GetDeploymentsResponse, error) {
	type wrapper struct {
		Deployments types.GetDeploymentsResponse `json:"deployments"`
	}

	resp, err := doRequest[*types.GetDeploymentsRequest, wrapper](ctx, c, req)
	if err != nil {
		return nil, err
	}
	return &resp.Deployments, nil
}
