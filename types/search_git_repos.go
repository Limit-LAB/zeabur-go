package types

type Provider string

const (
	ProviderGithub Provider = "GitHub"
)

type SearchGitReposRequest struct {
	Provider Provider
	Limit    uint32
	Keyword  string
}

func NewSearchGitReposRequest(limit uint32, keyword string) *SearchGitReposRequest {
	return &SearchGitReposRequest{
		Provider: ProviderGithub,
		Limit:    limit,
		Keyword:  keyword,
	}
}

func (r *SearchGitReposRequest) ToRequest() (map[string]interface{}, error) {
	return map[string]interface{}{
		"query": `query SearchGitRepos($provider: GitProvider!, $limit: Int!, $keyword: String) {
			searchGitRepositories(provider: $provider, Limit: $limit, keyword: $keyword) {
				id
				name
				provider
				url
				owner
			}
		}`,
		"variables": map[string]interface{}{
			"provider": r.Provider,
			"limit":    r.Limit,
			"keyword":  r.Keyword,
		},
	}, nil
}

type SearchGitReposResponse []Repo

type Repo struct {
	ID       uint64 `json:"id"`
	Name     string `json:"name"`
	Provider string `json:"provider"`
	URL      string `json:"url"`
	Owner    string `json:"owner"`
}
