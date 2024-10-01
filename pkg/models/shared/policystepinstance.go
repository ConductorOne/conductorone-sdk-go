// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// PolicyStepInstanceState - The state of the step, which is either active or done.
type PolicyStepInstanceState string

const (
	PolicyStepInstanceStatePolicyStepStateUnspecified PolicyStepInstanceState = "POLICY_STEP_STATE_UNSPECIFIED"
	PolicyStepInstanceStatePolicyStepStateActive      PolicyStepInstanceState = "POLICY_STEP_STATE_ACTIVE"
	PolicyStepInstanceStatePolicyStepStateDone        PolicyStepInstanceState = "POLICY_STEP_STATE_DONE"
)

func (e PolicyStepInstanceState) ToPointer() *PolicyStepInstanceState {
	return &e
}

// PolicyStepInstance - The policy step instance includes a reference to an instance of a policy step that tracks state and has a unique ID.
//
// This message contains a oneof named instance. Only a single field of the following list may be set at a time:
//   - approval
//   - provision
//   - accept
//   - reject
//   - wait
type PolicyStepInstance struct {
	// This policy step indicates that a ticket should have an approved outcome. This is a terminal approval state and is used to explicitly define the end of approval steps.
	//  The instance is just a marker for it being copied into an active policy.
	AcceptInstance *AcceptInstance `json:"accept,omitempty"`
	// The approval instance object describes the way a policy step should be approved as well as its outcomes and state.
	//
	// This message contains a oneof named outcome. Only a single field of the following list may be set at a time:
	//   - approved
	//   - denied
	//   - reassigned
	//   - restarted
	//   - reassignedByError
	//
	ApprovalInstance *ApprovalInstance `json:"approval,omitempty"`
	// A provision instance describes the specific configuration of an executing provision policy step including actions taken and notification id.
	//
	// This message contains a oneof named outcome. Only a single field of the following list may be set at a time:
	//   - completed
	//   - cancelled
	//   - errored
	//   - reassignedByError
	//
	ProvisionInstance *ProvisionInstance `json:"provision,omitempty"`
	// This policy step indicates that a ticket should have a denied outcome. This is a terminal approval state and is used to explicitly define the end of approval steps.
	//  The instance is just a marker for it being copied into an active policy.
	RejectInstance *RejectInstance `json:"reject,omitempty"`
	// Used by the policy engine to describe an instantiated wait step.
	//
	// This message contains a oneof named until. Only a single field of the following list may be set at a time:
	//   - condition
	//
	//
	// This message contains a oneof named outcome. Only a single field of the following list may be set at a time:
	//   - succeeded
	//   - timedOut
	//
	WaitInstance *WaitInstance `json:"wait,omitempty"`
	// The ID of the PolicyStepInstance. This is required by many action submission endpoints to indicate what step you're approving.
	ID *string `json:"id,omitempty"`
	// The policy generation id refers to the version of the policy that this step was created from.
	PolicyGenerationID *string `json:"policyGenerationId,omitempty"`
	// The state of the step, which is either active or done.
	State *PolicyStepInstanceState `json:"state,omitempty"`
}

func (o *PolicyStepInstance) GetAcceptInstance() *AcceptInstance {
	if o == nil {
		return nil
	}
	return o.AcceptInstance
}

func (o *PolicyStepInstance) GetApprovalInstance() *ApprovalInstance {
	if o == nil {
		return nil
	}
	return o.ApprovalInstance
}

func (o *PolicyStepInstance) GetProvisionInstance() *ProvisionInstance {
	if o == nil {
		return nil
	}
	return o.ProvisionInstance
}

func (o *PolicyStepInstance) GetRejectInstance() *RejectInstance {
	if o == nil {
		return nil
	}
	return o.RejectInstance
}

func (o *PolicyStepInstance) GetWaitInstance() *WaitInstance {
	if o == nil {
		return nil
	}
	return o.WaitInstance
}

func (o *PolicyStepInstance) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *PolicyStepInstance) GetPolicyGenerationID() *string {
	if o == nil {
		return nil
	}
	return o.PolicyGenerationID
}

func (o *PolicyStepInstance) GetState() *PolicyStepInstanceState {
	if o == nil {
		return nil
	}
	return o.State
}
