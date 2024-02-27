package types

type GetEnvironmentsRequest struct {
	ProjectID string
}

func NewGetEnvironmentsRequest(projectID string) *GetEnvironmentsRequest {
	return &GetEnvironmentsRequest{
		ProjectID: projectID,
	}
}

func (r *GetEnvironmentsRequest) ToRequest() (map[string]interface{}, error) {
	return map[string]interface{}{
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

type GetEnvironmentsResponse struct {
	Environments []Environment `json:"environments"`
}

type Environment struct {
	ID string `json:"_id"`
}
