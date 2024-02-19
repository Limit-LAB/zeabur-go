package request

import (
	"context"

	"github.com/limit-lab/zeabur-go/types"
)

func (c *Client) CreateService(ctx context.Context, req *types.CreateServiceRequest) (*types.CreateServiceResponse, error) {
	type wrapper struct {
		CreateService *types.CreateServiceResponse `json:"createService"`
	}

	resp, err := doRequest[*types.CreateServiceRequest, wrapper](ctx, c, req)
	if err != nil {
		return nil, err
	}
	return resp.CreateService, nil
}
