package request

import (
	"context"

	"github.com/limit-lab/zeabur-go/types"
)

func (c *Client) Invite(ctx context.Context, req *types.InviteRequest) (*types.InviteID, error) {
	type wrapper struct {
		InviteID *types.InviteID `json:"inviteUserToProject"`
	}

	resp, err := doRequest[*types.InviteRequest, wrapper](ctx, c, req)
	if err != nil {
		return nil, err
	}
	return resp.InviteID, nil
}
