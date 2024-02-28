package request

import (
	"context"

	"github.com/limit-lab/zeabur-go/types"
)

func (c *Client) UpdateRootDirectory(ctx context.Context, req *types.UpdateRootDirectoryRequest) (*types.UpdateRootDirectoryResponse, error) {
	type wrapper struct {
		UpdateRootDirectory types.UpdateRootDirectoryResponse `json:"updateRootDirectory"`
	}

	resp, err := doRequest[*types.UpdateRootDirectoryRequest, wrapper](ctx, c, req)
	if err != nil {
		return nil, err
	}
	return &resp.UpdateRootDirectory, nil
}
