package types

type Region string

const (
	RegionSfo1 Region = "sfo1"
	RegionHkg1 Region = "hkg1"
)

type CreateProjectRequest struct {
	Region Region
}

func NewCreateProjectRequest(region Region) *CreateProjectRequest {
	return &CreateProjectRequest{
		Region: region,
	}
}

func (r *CreateProjectRequest) ToRequest() (*map[string]interface{}, error) {
	return &map[string]interface{}{
		"query": `mutation CreateProject($region: String!) {
			createProject(region: $region) {
				_id
			}
		}`,
		"variables": map[string]interface{}{
			"region": r.Region,
		},
	}, nil
}

type CreateProjectResponse struct {
	ProjectID string `json:"_id"`
}
