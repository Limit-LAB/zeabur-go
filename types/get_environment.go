package types

type GetEnvironmentRequest struct {
	ProjectID string
}

func NewGetEnvironmentRequest(projectID string) *GetEnvironmentRequest {
	return &GetEnvironmentRequest{
		ProjectID: projectID,
	}
}

func (r *GetEnvironmentRequest) ToRequest() (*map[string]interface{}, error) {
	return &map[string]interface{}{
		"query": `query GetEnvironment($projectID: ObjectID!) {
			environments(projectID: $projectID) {
				_id
			}
		}`,
		"variables": map[string]interface{}{
			"projectID": r.ProjectID,
		},
	}, nil
}

type GetEnvironmentResponse struct {
	Environments []Environment `json:"environments"`
}

type Environment struct {
	ID string `json:"_id"`
}
