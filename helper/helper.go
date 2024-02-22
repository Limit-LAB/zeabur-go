package helper

import (
	"context"
	"fmt"
	"io"

	"github.com/limit-lab/zeabur-go/request"
	"github.com/limit-lab/zeabur-go/types"
)

type Helper struct {
	c *request.Client
}

func NewHelper(client *request.Client) *Helper {
	return &Helper{
		c: client,
	}
}

func (h *Helper) DeployProject(ctx context.Context,
	codeZip io.Reader, serviceName, domainName string) (string, error) {
	projectResp, err := h.c.CreateProject(ctx,
		types.NewCreateProjectRequest(types.RegionSfo1))
	if err != nil {
		return "", fmt.Errorf("failed to create project: %w", err)
	}

	envResp, err := h.c.GetEnvironment(ctx,
		types.NewGetEnvironmentRequest(projectResp.ProjectID))
	if err != nil {
		return "", fmt.Errorf("failed to get environment: %w", err)
	}
	if len(envResp.Environments) == 0 {
		return "", fmt.Errorf("no environment found")
	}

	serviceResp, err := h.c.CreateService(ctx,
		types.NewCreateServiceRequest(projectResp.ProjectID, serviceName))
	if err != nil {
		return "", fmt.Errorf("failed to create service: %w", err)
	}

	deployReq := types.NewDeployRequest(
		projectResp.ProjectID,
		serviceResp.ServiceID,
		envResp.Environments[0].ID,
		codeZip,
	)
	_, err = h.c.Deploy(ctx, deployReq)
	if err != nil {
		return "", fmt.Errorf("failed to deploy project: %w", err)
	}

	domainResp, err := h.c.CreateDomain(ctx, types.NewCreateDomainRequest(
		serviceResp.ServiceID, envResp.Environments[0].ID, domainName))
	if err != nil {
		return "", fmt.Errorf("failed to create domain: %w", err)
	}

	return domainResp.Domain, nil
}
