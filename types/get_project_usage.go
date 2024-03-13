package types

import "fmt"

type GroupByEntity string

const (
	GroupByEntityCategory = GroupByEntity("CATEGORY")
)

type EntityType string

const (
	EntityTypeMemory = EntityType("Memory")
	EntityTypeCPU    = EntityType("CPU")
	EntityTypeEgress = EntityType("Egress")
)

type GetProjectUsageRequest struct {
	ProjectID          string
	UsageGroupByEntity GroupByEntity
}

func NewGetProjectUsageRequest(projectID string) *GetProjectUsageRequest {
	return &GetProjectUsageRequest{
		ProjectID:          projectID,
		UsageGroupByEntity: GroupByEntityCategory,
	}
}

func (r *GetProjectUsageRequest) ToRequest() (map[string]interface{}, error) {
	return map[string]interface{}{
		"query": fmt.Sprintf(`query GetProjectUsage($projectID: ObjectID!) {
			projectUsage(projectID: $projectID, usageGroupByEntity: %s) {
				usages {
					entity
					usage
				}
				periodStart
				periodEnd
			}
		}`, r.UsageGroupByEntity),
		"variables": map[string]interface{}{
			"projectID": r.ProjectID,
		},
	}, nil
}

type GetProjectUsageResponse struct {
	Usages      []Usage `json:"usages"`
	PeriodStart string  `json:"periodStart"`
	PeriodEnd   string  `json:"periodEnd"`
}

type Usage struct {
	Entity EntityType `json:"entity"`
	Usage  float64    `json:"usage"`
}
