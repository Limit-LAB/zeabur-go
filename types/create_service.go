package types

type ServiceTemplate string

const (
	TemplateGit ServiceTemplate = "GIT"
)

type CreateServiceRequest struct {
	ProjectID       string
	ServiceName     string
	ServiceTemplate ServiceTemplate
}

func NewCreateServiceRequest(projectID string, serviceName string) *CreateServiceRequest {
	return &CreateServiceRequest{
		ProjectID:       projectID,
		ServiceName:     serviceName,
		ServiceTemplate: TemplateGit,
	}
}

func (r *CreateServiceRequest) ToRequest() (map[string]interface{}, error) {
	return map[string]interface{}{
		"query": `mutation CreateService($projectID: ObjectID!, $template: ServiceTemplate!, $serviceName: String!) {
			createService(projectID: $projectID, template: $template, name: $serviceName) {
				_id
			}
		}`,
		"variables": map[string]interface{}{
			"projectID":   r.ProjectID,
			"template":    r.ServiceTemplate,
			"serviceName": r.ServiceName,
		},
	}, nil
}

type CreateServiceResponse struct {
	ServiceID string `json:"_id"`
}
