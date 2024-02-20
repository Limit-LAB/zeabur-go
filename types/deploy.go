package types

import "io"

type DeployRequest struct {
	ProjectID     string
	ServiceID     string
	EnvironmentID string
	CodeZip       io.Reader
}

func NewDeployRequest(projectID string, serviceID string, environmentID string, codeZip io.Reader) *DeployRequest {
	return &DeployRequest{
		ProjectID:     projectID,
		ServiceID:     serviceID,
		EnvironmentID: environmentID,
		CodeZip:       codeZip,
	}
}
