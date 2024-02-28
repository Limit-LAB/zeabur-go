package types

type UpdateRootDirectoryRequest struct {
	ServiceID     string
	RootDirectory string
}

func NewUpdateRootDirectoryRequest(serviceID string, rootDirectory string) *UpdateRootDirectoryRequest {
	return &UpdateRootDirectoryRequest{
		ServiceID:     serviceID,
		RootDirectory: rootDirectory,
	}
}

func (r *UpdateRootDirectoryRequest) ToRequest() (map[string]interface{}, error) {
	return map[string]interface{}{
		"query": `mutation UpdateRootDirectory($rootDirectory: String!, $serviceID: ObjectID!) {
			updateRootDirectory(rootDirectory: $rootDirectory, serviceID: $serviceID)
		}`,
		"variables": map[string]interface{}{
			"rootDirectory": r.RootDirectory,
			"serviceID":     r.ServiceID,
		},
	}, nil
}

type UpdateRootDirectoryResponse bool
