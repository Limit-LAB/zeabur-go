package types

type DeploymentStatus string

const (
	StatusBuilding  DeploymentStatus = "BUILDING"
	StatusCanceled  DeploymentStatus = "CANCELED"
	StatusCrashed   DeploymentStatus = "CRASHED"
	StatusDeploying DeploymentStatus = "DEPLOYING"
	StatusFailed    DeploymentStatus = "FAILED"
	StatusPending   DeploymentStatus = "PENDING"
	StatusRemoved   DeploymentStatus = "REMOVED"
	StatusRunning   DeploymentStatus = "RUNNING"
	StatusUnknown   DeploymentStatus = "UNKNOWN"
)

type GetDeploymentsRequest struct {
	EnvironmentID string
	ServiceID     string
	Cursor        string
	Filter        DeploymentStatus
	Limit         uint32
}

func NewGetDeploymentsRequest(environmentID string, serviceID string) *GetDeploymentsRequest {
	return &GetDeploymentsRequest{
		EnvironmentID: environmentID,
		ServiceID:     serviceID,
		Limit:         5,
	}
}

func (r *GetDeploymentsRequest) ToRequest() (map[string]interface{}, error) {
	jsonReq := map[string]interface{}{
		"query": `query GetDeployments($environmentID: ObjectID!, $serviceID: ObjectID!, $cursor: String, $limit: Int, $filter: DeploymentStatus) {
			deployments(environmentID: $environmentID, serviceID: $serviceID, cursor: $cursor, perPage: $limit, filter: $filter) {
				edges {
					cursor
					node {
						_id
						status
						commitMessage
						createdAt
						finishedAt
						planType
						planMeta
					}
				}
			}
		}`,
		"variables": map[string]interface{}{
			"environmentID": r.EnvironmentID,
			"serviceID":     r.ServiceID,
			"limit":         r.Limit,
		},
	}

	if r.Cursor != "" {
		jsonReq["variables"].(map[string]interface{})["cursor"] = r.Cursor
	}
	if r.Filter != "" {
		jsonReq["variables"].(map[string]interface{})["filter"] = r.Filter
	}

	return jsonReq, nil
}

type GetDeploymentsResponse struct {
	DeploymentEdges []DeploymentEdge `json:"edges"`
}

type DeploymentEdge struct {
	Cursor         string         `json:"cursor"`
	DeploymentNode DeploymentNode `json:"node"`
}

type DeploymentNode struct {
	DeploymentID  string                 `json:"_id"`
	Status        DeploymentStatus       `json:"status"`
	CommitMessage string                 `json:"commitMessage"`
	CreatedAt     string                 `json:"createdAt"`
	FinishedAt    string                 `json:"finishedAt"`
	PlanType      string                 `json:"planType"`
	PlanMeta      map[string]interface{} `json:"planMeta"`
}
