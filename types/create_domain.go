package types

type CreateDomainRequest struct {
	ServiceID     string
	EnvironmentID string
	Domain        string
	IsGenerated   bool
}

func NewCreateDomainRequest(serviceID string, environmentID string, domain string) *CreateDomainRequest {
	return &CreateDomainRequest{
		ServiceID:     serviceID,
		EnvironmentID: environmentID,
		Domain:        domain,
		IsGenerated:   true,
	}
}

var _ ToGraphQLRequest = (*CreateDomainRequest)(nil)

func (r *CreateDomainRequest) ToRequest() (map[string]interface{}, error) {
	return map[string]interface{}{
		"query": `mutation CreateDomain($serviceID: ObjectID!, $environmentID: ObjectID!, $domain: String!, $isGenerated: Boolean!) {
			createDomain(serviceID: $serviceID, environmentID: $environmentID, domain: $domain, isGenerated: $isGenerated) {
				_id
			}
		}`,
		"variables": map[string]interface{}{
			"serviceID":     r.ServiceID,
			"environmentID": r.EnvironmentID,
			"domain":        r.Domain,
			"isGenerated":   r.IsGenerated,
		},
	}, nil
}

type CreateDomainResponse struct {
	DomainID string `json:"_id"`
}
