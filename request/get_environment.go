package request

import (
	"context"

	"github.com/limit-lab/zeabur-go/types"
)

func (c *Client) GetEnvironment(ctx context.Context, req *types.GetEnvironmentRequest) (*types.GetEnvironmentResponse, error) {
	resp, err := doRequest[*types.GetEnvironmentRequest, types.GetEnvironmentResponse](ctx, c, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
