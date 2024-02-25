package types

type GetProjectRequest struct {
	ProjectID string
}

func NewGetProjectRequest(projectID string) *GetProjectRequest {
	return &GetProjectRequest{
		ProjectID: projectID,
	}
}

func (r *GetProjectRequest) ToRequest() (map[string]interface{}, error) {
	return map[string]interface{}{
		"query": `query GetProject($projectID: ObjectID!) {
			project(_id: $projectID) {
				_id
				owner {
					_id
					name
					email
					avatarURL
				}
				collaborators {
					_id
					name
					email
					avatarURL
				}
				name
				description
				iconURL
			}
		}`,
		"variables": map[string]interface{}{
			"projectID": r.ProjectID,
		},
	}, nil
}

type GetProjectResponse struct {
	ProjectID     string `json:"_id"`
	Owner         User   `json:"owner"`
	Collaborators []User `json:"collaborators"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	IconURL       string `json:"iconURL"`
}

type User struct {
	ID        string `json:"_id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	AvatarURL string `json:"avatarURL"`
}
