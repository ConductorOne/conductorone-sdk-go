// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"time"
)

// PolicyPolicyType -  Indicates the type of this policy. Can also be used to get the value from policySteps.
type PolicyPolicyType string

const (
	PolicyPolicyTypePolicyTypeUnspecified   PolicyPolicyType = "POLICY_TYPE_UNSPECIFIED"
	PolicyPolicyTypePolicyTypeGrant         PolicyPolicyType = "POLICY_TYPE_GRANT"
	PolicyPolicyTypePolicyTypeRevoke        PolicyPolicyType = "POLICY_TYPE_REVOKE"
	PolicyPolicyTypePolicyTypeCertify       PolicyPolicyType = "POLICY_TYPE_CERTIFY"
	PolicyPolicyTypePolicyTypeAccessRequest PolicyPolicyType = "POLICY_TYPE_ACCESS_REQUEST"
	PolicyPolicyTypePolicyTypeProvision     PolicyPolicyType = "POLICY_TYPE_PROVISION"
)

func (e PolicyPolicyType) ToPointer() *PolicyPolicyType {
	return &e
}

func (e *PolicyPolicyType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "POLICY_TYPE_UNSPECIFIED":
		fallthrough
	case "POLICY_TYPE_GRANT":
		fallthrough
	case "POLICY_TYPE_REVOKE":
		fallthrough
	case "POLICY_TYPE_CERTIFY":
		fallthrough
	case "POLICY_TYPE_ACCESS_REQUEST":
		fallthrough
	case "POLICY_TYPE_PROVISION":
		*e = PolicyPolicyType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PolicyPolicyType: %v", v)
	}
}

// Policy -  A policy describes the behavior of the ConductorOne system when processing a task. You can describe the type, approvers, fallback behavior, and escalation processes.
type Policy struct {
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
	//  The description of the Policy.
	//
	Description *string `json:"description,omitempty"`
	//  The display name of the Policy.
	//
	DisplayName *string `json:"displayName,omitempty"`
	//  The ID of the Policy.
	//
	ID *string `json:"id,omitempty"`
	//  A map of string(policy type) to steps in a policy. This structure is leftover from a previous design, and should only ever have one key->value set.
	//
	PolicySteps map[string]PolicySteps `json:"policySteps,omitempty"`
	//  Indicates the type of this policy. Can also be used to get the value from policySteps.
	//
	PolicyType *PolicyPolicyType `json:"policyType,omitempty"`
	//  An array of actions (ordered) to take place after a policy completes processing.
	//
	PostActions []PolicyPostActions `json:"postActions,omitempty"`
	//  A policy configuration option that allows for reassinging tasks to delgated users. This level of delegation referrs to the individual delegates users set on their account.
	//
	ReassignTasksToDelegates *bool `json:"reassignTasksToDelegates,omitempty"`
	//  Whether this policy is a builtin system policy. Builtin system policies cannot be edited.
	//
	SystemBuiltin *bool      `json:"systemBuiltin,omitempty"`
	UpdatedAt     *time.Time `json:"updatedAt,omitempty"`
}

func (o *Policy) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *Policy) GetDeletedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.DeletedAt
}

func (o *Policy) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *Policy) GetDisplayName() *string {
	if o == nil {
		return nil
	}
	return o.DisplayName
}

func (o *Policy) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Policy) GetPolicySteps() map[string]PolicySteps {
	if o == nil {
		return nil
	}
	return o.PolicySteps
}

func (o *Policy) GetPolicyType() *PolicyPolicyType {
	if o == nil {
		return nil
	}
	return o.PolicyType
}

func (o *Policy) GetPostActions() []PolicyPostActions {
	if o == nil {
		return nil
	}
	return o.PostActions
}

func (o *Policy) GetReassignTasksToDelegates() *bool {
	if o == nil {
		return nil
	}
	return o.ReassignTasksToDelegates
}

func (o *Policy) GetSystemBuiltin() *bool {
	if o == nil {
		return nil
	}
	return o.SystemBuiltin
}

func (o *Policy) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

// PolicyInput -  A policy describes the behavior of the ConductorOne system when processing a task. You can describe the type, approvers, fallback behavior, and escalation processes.
type PolicyInput struct {
	//  The description of the Policy.
	//
	Description *string `json:"description,omitempty"`
	//  The display name of the Policy.
	//
	DisplayName *string `json:"displayName,omitempty"`
	//  A map of string(policy type) to steps in a policy. This structure is leftover from a previous design, and should only ever have one key->value set.
	//
	PolicySteps map[string]PolicyStepsInput `json:"policySteps,omitempty"`
	//  Indicates the type of this policy. Can also be used to get the value from policySteps.
	//
	PolicyType *PolicyPolicyType `json:"policyType,omitempty"`
	//  An array of actions (ordered) to take place after a policy completes processing.
	//
	PostActions []PolicyPostActions `json:"postActions,omitempty"`
	//  A policy configuration option that allows for reassinging tasks to delgated users. This level of delegation referrs to the individual delegates users set on their account.
	//
	ReassignTasksToDelegates *bool `json:"reassignTasksToDelegates,omitempty"`
}

func (o *PolicyInput) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *PolicyInput) GetDisplayName() *string {
	if o == nil {
		return nil
	}
	return o.DisplayName
}

func (o *PolicyInput) GetPolicySteps() map[string]PolicyStepsInput {
	if o == nil {
		return nil
	}
	return o.PolicySteps
}

func (o *PolicyInput) GetPolicyType() *PolicyPolicyType {
	if o == nil {
		return nil
	}
	return o.PolicyType
}

func (o *PolicyInput) GetPostActions() []PolicyPostActions {
	if o == nil {
		return nil
	}
	return o.PostActions
}

func (o *PolicyInput) GetReassignTasksToDelegates() *bool {
	if o == nil {
		return nil
	}
	return o.ReassignTasksToDelegates
}
