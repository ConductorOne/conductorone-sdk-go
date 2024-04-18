// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/conductorone/conductorone-sdk-go/pkg/utils"
	"time"
)

// TaskTypeGrantOutcome - The outcome of the grant.
type TaskTypeGrantOutcome string

const (
	TaskTypeGrantOutcomeGrantOutcomeUnspecified TaskTypeGrantOutcome = "GRANT_OUTCOME_UNSPECIFIED"
	TaskTypeGrantOutcomeGrantOutcomeGranted     TaskTypeGrantOutcome = "GRANT_OUTCOME_GRANTED"
	TaskTypeGrantOutcomeGrantOutcomeDenied      TaskTypeGrantOutcome = "GRANT_OUTCOME_DENIED"
	TaskTypeGrantOutcomeGrantOutcomeError       TaskTypeGrantOutcome = "GRANT_OUTCOME_ERROR"
	TaskTypeGrantOutcomeGrantOutcomeCancelled   TaskTypeGrantOutcome = "GRANT_OUTCOME_CANCELLED"
)

func (e TaskTypeGrantOutcome) ToPointer() *TaskTypeGrantOutcome {
	return &e
}

func (e *TaskTypeGrantOutcome) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "GRANT_OUTCOME_UNSPECIFIED":
		fallthrough
	case "GRANT_OUTCOME_GRANTED":
		fallthrough
	case "GRANT_OUTCOME_DENIED":
		fallthrough
	case "GRANT_OUTCOME_ERROR":
		fallthrough
	case "GRANT_OUTCOME_CANCELLED":
		*e = TaskTypeGrantOutcome(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TaskTypeGrantOutcome: %v", v)
	}
}

// The TaskTypeGrant message indicates that a task is a grant task and all related details.
type TaskTypeGrant struct {
	// The ID of the app entitlement.
	AppEntitlementID *string `json:"appEntitlementId,omitempty"`
	// The ID of the app.
	AppID *string `json:"appId,omitempty"`
	// The ID of the app user.
	AppUserID     *string `json:"appUserId,omitempty"`
	GrantDuration *string `json:"grantDuration,omitempty"`
	// The ID of the user.
	IdentityUserID *string `json:"identityUserId,omitempty"`
	// The outcome of the grant.
	Outcome     *TaskTypeGrantOutcome `json:"outcome,omitempty"`
	OutcomeTime *time.Time            `json:"outcomeTime,omitempty"`
	// The TaskGrantSource message tracks which external URL was the source of the specificed grant ticket.
	TaskGrantSource *TaskGrantSource `json:"source,omitempty"`
}

func (t TaskTypeGrant) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TaskTypeGrant) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TaskTypeGrant) GetAppEntitlementID() *string {
	if o == nil {
		return nil
	}
	return o.AppEntitlementID
}

func (o *TaskTypeGrant) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

func (o *TaskTypeGrant) GetAppUserID() *string {
	if o == nil {
		return nil
	}
	return o.AppUserID
}

func (o *TaskTypeGrant) GetGrantDuration() *string {
	if o == nil {
		return nil
	}
	return o.GrantDuration
}

func (o *TaskTypeGrant) GetIdentityUserID() *string {
	if o == nil {
		return nil
	}
	return o.IdentityUserID
}

func (o *TaskTypeGrant) GetOutcome() *TaskTypeGrantOutcome {
	if o == nil {
		return nil
	}
	return o.Outcome
}

func (o *TaskTypeGrant) GetOutcomeTime() *time.Time {
	if o == nil {
		return nil
	}
	return o.OutcomeTime
}

func (o *TaskTypeGrant) GetTaskGrantSource() *TaskGrantSource {
	if o == nil {
		return nil
	}
	return o.TaskGrantSource
}
