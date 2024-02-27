package request

import (
	"context"

	"github.com/limit-lab/zeabur-go/types"
)

func (c *Client) GetEnvironments(ctx context.Context, req *types.GetEnvironmentsRequest) (*types.GetEnvironmentsResponse, error) {
	resp, err := doRequest[*types.GetEnvironmentsRequest, types.GetEnvironmentsResponse](ctx, c, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
