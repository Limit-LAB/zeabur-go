package request

import (
	"context"

	"github.com/limit-lab/zeabur-go/types"
)

func (c *Client) DeleteService(ctx context.Context, req *types.DeleteServiceRequest) (*types.DeleteServiceResponse, error) {
	type wrapper struct {
		DeleteService *types.DeleteServiceResponse `json:"deleteService"`
	}

	resp, err := doRequest[*types.DeleteServiceRequest, wrapper](ctx, c, req)
	if err != nil {
		return nil, err
	}
	return resp.DeleteService, nil
}
