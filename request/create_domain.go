package request

import (
	"context"

	"github.com/limit-lab/zeabur-go/types"
)

func (c *Client) CreateDomain(ctx context.Context, req *types.CreateDomainRequest) (*types.CreateDomainResponse, error) {
	type wrapper struct {
		AddDomain *types.CreateDomainResponse `json:"addDomain"`
	}

	resp, err := doRequest[*types.CreateDomainRequest, wrapper](ctx, c, req)
	if err != nil {
		return nil, err
	}
	return resp.AddDomain, nil
}
