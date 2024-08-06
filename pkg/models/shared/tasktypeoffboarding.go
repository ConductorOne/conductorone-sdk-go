// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/conductorone/conductorone-sdk-go/pkg/utils"
	"time"
)

// TaskTypeOffboardingOutcome - The outcome field.
type TaskTypeOffboardingOutcome string

const (
	TaskTypeOffboardingOutcomeOffboardingOutcomeUnspecified TaskTypeOffboardingOutcome = "OFFBOARDING_OUTCOME_UNSPECIFIED"
	TaskTypeOffboardingOutcomeOffboardingOutcomeInProgress  TaskTypeOffboardingOutcome = "OFFBOARDING_OUTCOME_IN_PROGRESS"
	TaskTypeOffboardingOutcomeOffboardingOutcomeDone        TaskTypeOffboardingOutcome = "OFFBOARDING_OUTCOME_DONE"
	TaskTypeOffboardingOutcomeOffboardingOutcomeError       TaskTypeOffboardingOutcome = "OFFBOARDING_OUTCOME_ERROR"
	TaskTypeOffboardingOutcomeOffboardingOutcomeCancelled   TaskTypeOffboardingOutcome = "OFFBOARDING_OUTCOME_CANCELLED"
)

func (e TaskTypeOffboardingOutcome) ToPointer() *TaskTypeOffboardingOutcome {
	return &e
}
func (e *TaskTypeOffboardingOutcome) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "OFFBOARDING_OUTCOME_UNSPECIFIED":
		fallthrough
	case "OFFBOARDING_OUTCOME_IN_PROGRESS":
		fallthrough
	case "OFFBOARDING_OUTCOME_DONE":
		fallthrough
	case "OFFBOARDING_OUTCOME_ERROR":
		fallthrough
	case "OFFBOARDING_OUTCOME_CANCELLED":
		*e = TaskTypeOffboardingOutcome(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TaskTypeOffboardingOutcome: %v", v)
	}
}

// The TaskTypeOffboarding message.
type TaskTypeOffboarding struct {
	// The outcome field.
	Outcome     *TaskTypeOffboardingOutcome `json:"outcome,omitempty"`
	OutcomeTime *time.Time                  `json:"outcomeTime,omitempty"`
	// The subjectUserId field.
	SubjectUserID *string `json:"subjectUserId,omitempty"`
}

func (t TaskTypeOffboarding) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TaskTypeOffboarding) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TaskTypeOffboarding) GetOutcome() *TaskTypeOffboardingOutcome {
	if o == nil {
		return nil
	}
	return o.Outcome
}

func (o *TaskTypeOffboarding) GetOutcomeTime() *time.Time {
	if o == nil {
		return nil
	}
	return o.OutcomeTime
}

func (o *TaskTypeOffboarding) GetSubjectUserID() *string {
	if o == nil {
		return nil
	}
	return o.SubjectUserID
}
