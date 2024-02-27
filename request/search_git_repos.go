package request

import (
	"context"

	"github.com/limit-lab/zeabur-go/types"
)

func (c *Client) SearchGitRepos(ctx context.Context, req *types.SearchGitReposRequest) (*types.SearchGitReposResponse, error) {
	type wrapper struct {
		SearchGitRepositories types.SearchGitReposResponse `json:"searchGitRepositories"`
	}

	resp, err := doRequest[*types.SearchGitReposRequest, wrapper](ctx, c, req)
	if err != nil {
		return nil, err
	}
	return &resp.SearchGitRepositories, nil
}
