package types

type InviteRequest struct {
	ProjectID    string
	InviteeEmail string
}

func NewInviteRequest(projectID string, inviteeEmail string) *InviteRequest {
	return &InviteRequest{
		ProjectID:    projectID,
		InviteeEmail: inviteeEmail,
	}
}

func (r *InviteRequest) ToRequest() (map[string]interface{}, error) {
	return map[string]interface{}{
		"query": `mutation Invite($inviteeEmail: String!, $projectID: ObjectID!) {
			inviteUserToProject(inviteeEmail: $inviteeEmail, projectID: $projectID)
		}`,
		"variables": map[string]interface{}{
			"inviteeEmail": r.InviteeEmail,
			"projectID":    r.ProjectID,
		},
	}, nil
}

type InviteID string
