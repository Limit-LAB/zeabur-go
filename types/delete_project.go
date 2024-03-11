package types

type DeleteProjectRequest struct {
	ProjectID string
}

func NewDeleteProjectRequest(projectID string) *DeleteProjectRequest {
	return &DeleteProjectRequest{
		ProjectID: projectID,
	}
}

func (r *DeleteProjectRequest) ToRequest() (map[string]interface{}, error) {
	return map[string]interface{}{
		"query": `mutation DeleteProject($projectID: ObjectID!) {
			deleteProject(_id: $projectID)
		}`,
		"variables": map[string]interface{}{
			"projectID": r.ProjectID,
		},
	}, nil
}

type DeleteProjectResponse bool
