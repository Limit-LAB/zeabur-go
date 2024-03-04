package types

type DeleteServiceRequest struct {
	ServiceID string
}

func NewDeleteServiceRequest(serviceID string) *DeleteServiceRequest {
	return &DeleteServiceRequest{
		ServiceID: serviceID,
	}
}

func (r *DeleteServiceRequest) ToRequest() (map[string]interface{}, error) {
	return map[string]interface{}{
		"query": `mutation DeleteService($serviceID: ObjectID!) {
			deleteService(_id: $serviceID)
		}`,
		"variables": map[string]interface{}{
			"serviceID": r.ServiceID,
		},
	}, nil
}

type DeleteServiceResponse bool
