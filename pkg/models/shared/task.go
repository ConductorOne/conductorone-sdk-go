// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/conductorone/conductorone-sdk-go/pkg/utils"
	"time"
)

type Actions string

const (
	ActionsTaskActionTypeUnspecified                              Actions = "TASK_ACTION_TYPE_UNSPECIFIED"
	ActionsTaskActionTypeClose                                    Actions = "TASK_ACTION_TYPE_CLOSE"
	ActionsTaskActionTypeApprove                                  Actions = "TASK_ACTION_TYPE_APPROVE"
	ActionsTaskActionTypeDeny                                     Actions = "TASK_ACTION_TYPE_DENY"
	ActionsTaskActionTypeComment                                  Actions = "TASK_ACTION_TYPE_COMMENT"
	ActionsTaskActionTypeDelete                                   Actions = "TASK_ACTION_TYPE_DELETE"
	ActionsTaskActionTypeReassign                                 Actions = "TASK_ACTION_TYPE_REASSIGN"
	ActionsTaskActionTypeRestart                                  Actions = "TASK_ACTION_TYPE_RESTART"
	ActionsTaskActionTypeSendReminder                             Actions = "TASK_ACTION_TYPE_SEND_REMINDER"
	ActionsTaskActionTypeProvisionComplete                        Actions = "TASK_ACTION_TYPE_PROVISION_COMPLETE"
	ActionsTaskActionTypeProvisionCancelled                       Actions = "TASK_ACTION_TYPE_PROVISION_CANCELLED"
	ActionsTaskActionTypeProvisionErrored                         Actions = "TASK_ACTION_TYPE_PROVISION_ERRORED"
	ActionsTaskActionTypeProvisionAppUserTargetCreated            Actions = "TASK_ACTION_TYPE_PROVISION_APP_USER_TARGET_CREATED"
	ActionsTaskActionTypeRollbackSkipped                          Actions = "TASK_ACTION_TYPE_ROLLBACK_SKIPPED"
	ActionsTaskActionTypeHardReset                                Actions = "TASK_ACTION_TYPE_HARD_RESET"
	ActionsTaskActionTypeEscalateToEmergencyAccess                Actions = "TASK_ACTION_TYPE_ESCALATE_TO_EMERGENCY_ACCESS"
	ActionsTaskActionTypeChangePolicy                             Actions = "TASK_ACTION_TYPE_CHANGE_POLICY"
	ActionsTaskActionTypeRecalculateDenialFromBasePolicyDecisions Actions = "TASK_ACTION_TYPE_RECALCULATE_DENIAL_FROM_BASE_POLICY_DECISIONS"
	ActionsTaskActionTypeSetInsightsAndRecommendation             Actions = "TASK_ACTION_TYPE_SET_INSIGHTS_AND_RECOMMENDATION"
	ActionsTaskActionTypeSetAnalysisID                            Actions = "TASK_ACTION_TYPE_SET_ANALYSIS_ID"
)

func (e Actions) ToPointer() *Actions {
	return &e
}

func (e *Actions) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "TASK_ACTION_TYPE_UNSPECIFIED":
		fallthrough
	case "TASK_ACTION_TYPE_CLOSE":
		fallthrough
	case "TASK_ACTION_TYPE_APPROVE":
		fallthrough
	case "TASK_ACTION_TYPE_DENY":
		fallthrough
	case "TASK_ACTION_TYPE_COMMENT":
		fallthrough
	case "TASK_ACTION_TYPE_DELETE":
		fallthrough
	case "TASK_ACTION_TYPE_REASSIGN":
		fallthrough
	case "TASK_ACTION_TYPE_RESTART":
		fallthrough
	case "TASK_ACTION_TYPE_SEND_REMINDER":
		fallthrough
	case "TASK_ACTION_TYPE_PROVISION_COMPLETE":
		fallthrough
	case "TASK_ACTION_TYPE_PROVISION_CANCELLED":
		fallthrough
	case "TASK_ACTION_TYPE_PROVISION_ERRORED":
		fallthrough
	case "TASK_ACTION_TYPE_PROVISION_APP_USER_TARGET_CREATED":
		fallthrough
	case "TASK_ACTION_TYPE_ROLLBACK_SKIPPED":
		fallthrough
	case "TASK_ACTION_TYPE_HARD_RESET":
		fallthrough
	case "TASK_ACTION_TYPE_ESCALATE_TO_EMERGENCY_ACCESS":
		fallthrough
	case "TASK_ACTION_TYPE_CHANGE_POLICY":
		fallthrough
	case "TASK_ACTION_TYPE_RECALCULATE_DENIAL_FROM_BASE_POLICY_DECISIONS":
		fallthrough
	case "TASK_ACTION_TYPE_SET_INSIGHTS_AND_RECOMMENDATION":
		fallthrough
	case "TASK_ACTION_TYPE_SET_ANALYSIS_ID":
		*e = Actions(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Actions: %v", v)
	}
}

// Annotations - Contains an arbitrary serialized message along with a @type that describes the type of the serialized message.
type Annotations struct {
	// The type of the serialized message.
	AtType               *string                `json:"@type,omitempty"`
	AdditionalProperties map[string]interface{} `additionalProperties:"true" json:"-"`
}

func (a Annotations) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *Annotations) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Annotations) GetAtType() *string {
	if o == nil {
		return nil
	}
	return o.AtType
}

func (o *Annotations) GetAdditionalProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

// Processing - The processing state of a task as defined by the `processing_enum`
type Processing string

const (
	ProcessingTaskProcessingTypeUnspecified Processing = "TASK_PROCESSING_TYPE_UNSPECIFIED"
	ProcessingTaskProcessingTypeProcessing  Processing = "TASK_PROCESSING_TYPE_PROCESSING"
	ProcessingTaskProcessingTypeWaiting     Processing = "TASK_PROCESSING_TYPE_WAITING"
	ProcessingTaskProcessingTypeDone        Processing = "TASK_PROCESSING_TYPE_DONE"
)

func (e Processing) ToPointer() *Processing {
	return &e
}

func (e *Processing) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "TASK_PROCESSING_TYPE_UNSPECIFIED":
		fallthrough
	case "TASK_PROCESSING_TYPE_PROCESSING":
		fallthrough
	case "TASK_PROCESSING_TYPE_WAITING":
		fallthrough
	case "TASK_PROCESSING_TYPE_DONE":
		*e = Processing(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Processing: %v", v)
	}
}

// Recommendation - The recommendation field.
type Recommendation string

const (
	RecommendationInsightRecommendationUnspecified Recommendation = "INSIGHT_RECOMMENDATION_UNSPECIFIED"
	RecommendationInsightRecommendationApprove     Recommendation = "INSIGHT_RECOMMENDATION_APPROVE"
	RecommendationInsightRecommendationDeny        Recommendation = "INSIGHT_RECOMMENDATION_DENY"
	RecommendationInsightRecommendationReview      Recommendation = "INSIGHT_RECOMMENDATION_REVIEW"
)

func (e Recommendation) ToPointer() *Recommendation {
	return &e
}

func (e *Recommendation) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "INSIGHT_RECOMMENDATION_UNSPECIFIED":
		fallthrough
	case "INSIGHT_RECOMMENDATION_APPROVE":
		fallthrough
	case "INSIGHT_RECOMMENDATION_DENY":
		fallthrough
	case "INSIGHT_RECOMMENDATION_REVIEW":
		*e = Recommendation(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Recommendation: %v", v)
	}
}

// TaskState - The current state of the task as defined by the `state_enum`
type TaskState string

const (
	TaskStateTaskStateUnspecified TaskState = "TASK_STATE_UNSPECIFIED"
	TaskStateTaskStateOpen        TaskState = "TASK_STATE_OPEN"
	TaskStateTaskStateClosed      TaskState = "TASK_STATE_CLOSED"
)

func (e TaskState) ToPointer() *TaskState {
	return &e
}

func (e *TaskState) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "TASK_STATE_UNSPECIFIED":
		fallthrough
	case "TASK_STATE_OPEN":
		fallthrough
	case "TASK_STATE_CLOSED":
		*e = TaskState(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TaskState: %v", v)
	}
}

// Task - A fully-fleged task object. Includes its policy, references to external apps, its type, its processing history, and more.
type Task struct {
	// A policy instance is an object that contains a reference to the policy it was created from, the currently executing step, the next steps, and the history of previously completed steps.
	PolicyInstance *PolicyInstance `json:"policy,omitempty"`
	// Task Type provides configuration for the type of task: certify, grant, or revoke
	//
	// This message contains a oneof named task_type. Only a single field of the following list may be set at a time:
	//   - grant
	//   - revoke
	//   - certify
	//
	TaskType *TaskType `json:"type,omitempty"`
	// The actions that can be performed on the task by the current user.
	Actions []Actions `json:"actions,omitempty"`
	// The ID of the analysis object associated with this task created by an analysis workflow if the analysis feature is enabled for your tenant.
	AnalysisID *string `json:"analysisId,omitempty"`
	// An array of `google.protobuf.Any` annotations with various base64-encoded data.
	Annotations []Annotations `json:"annotations,omitempty"`
	// The count of comments.
	CommentCount *int       `json:"commentCount,omitempty"`
	CreatedAt    *time.Time `json:"createdAt,omitempty"`
	// The ID of the user that is the creator of this task. This may not always match the userId field.
	CreatedByUserID *string    `json:"createdByUserId,omitempty"`
	DeletedAt       *time.Time `json:"deletedAt,omitempty"`
	// The description of the task. This is also known as justification.
	Description *string `json:"description,omitempty"`
	// The display name of the task.
	DisplayName *string `json:"displayName,omitempty"`
	// A field indicating whether this task was created using an emergency access flow, or escalated to emergency access. On task creation, it will also use the app entitlement's emergency policy when possible.
	EmergencyAccess *bool `json:"emergencyAccess,omitempty"`
	// An array of external references to the task. Historically that has been items like Jira task IDs. This is currently unused, but may come back in the future for integrations.
	ExternalRefs []ExternalRef `json:"externalRefs,omitempty"`
	// The ID of the task.
	ID *string `json:"id,omitempty"`
	// The insightIds field.
	InsightIds []string `json:"insightIds,omitempty"`
	// A human-usable numeric ID of a task which can be included in place of the fully qualified task id in path parmeters (but not search queries).
	NumericID *string `json:"numericId,omitempty"`
	// The policy generation id refers to the current policy's generation ID. This is changed when the policy is changed on a task.
	PolicyGenerationID *string `json:"policyGenerationId,omitempty"`
	// The processing state of a task as defined by the `processing_enum`
	Processing *Processing `json:"processing,omitempty"`
	// The recommendation field.
	Recommendation *Recommendation `json:"recommendation,omitempty"`
	// The current state of the task as defined by the `state_enum`
	State *TaskState `json:"state,omitempty"`
	// An array of IDs belonging to Identity Users that are allowed to review this step in a task.
	StepApproverIds []string   `json:"stepApproverIds,omitempty"`
	UpdatedAt       *time.Time `json:"updatedAt,omitempty"`
	// The ID of the user that is the target of this task. This may be empty if we're targeting a specific app user that has no known identity user.
	UserID *string `json:"userId,omitempty"`
}

func (t Task) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *Task) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Task) GetPolicyInstance() *PolicyInstance {
	if o == nil {
		return nil
	}
	return o.PolicyInstance
}

func (o *Task) GetTaskType() *TaskType {
	if o == nil {
		return nil
	}
	return o.TaskType
}

func (o *Task) GetActions() []Actions {
	if o == nil {
		return nil
	}
	return o.Actions
}

func (o *Task) GetAnalysisID() *string {
	if o == nil {
		return nil
	}
	return o.AnalysisID
}

func (o *Task) GetAnnotations() []Annotations {
	if o == nil {
		return nil
	}
	return o.Annotations
}

func (o *Task) GetCommentCount() *int {
	if o == nil {
		return nil
	}
	return o.CommentCount
}

func (o *Task) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *Task) GetCreatedByUserID() *string {
	if o == nil {
		return nil
	}
	return o.CreatedByUserID
}

func (o *Task) GetDeletedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.DeletedAt
}

func (o *Task) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *Task) GetDisplayName() *string {
	if o == nil {
		return nil
	}
	return o.DisplayName
}

func (o *Task) GetEmergencyAccess() *bool {
	if o == nil {
		return nil
	}
	return o.EmergencyAccess
}

func (o *Task) GetExternalRefs() []ExternalRef {
	if o == nil {
		return nil
	}
	return o.ExternalRefs
}

func (o *Task) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Task) GetInsightIds() []string {
	if o == nil {
		return nil
	}
	return o.InsightIds
}

func (o *Task) GetNumericID() *string {
	if o == nil {
		return nil
	}
	return o.NumericID
}

func (o *Task) GetPolicyGenerationID() *string {
	if o == nil {
		return nil
	}
	return o.PolicyGenerationID
}

func (o *Task) GetProcessing() *Processing {
	if o == nil {
		return nil
	}
	return o.Processing
}

func (o *Task) GetRecommendation() *Recommendation {
	if o == nil {
		return nil
	}
	return o.Recommendation
}

func (o *Task) GetState() *TaskState {
	if o == nil {
		return nil
	}
	return o.State
}

func (o *Task) GetStepApproverIds() []string {
	if o == nil {
		return nil
	}
	return o.StepApproverIds
}

func (o *Task) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *Task) GetUserID() *string {
	if o == nil {
		return nil
	}
	return o.UserID
}
