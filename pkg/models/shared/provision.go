// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// Provision - The provision step references a provision policy for this step.
type Provision struct {
	// ProvisionPolicy is a oneOf that indicates how a provision step should be processed.
	//
	// This message contains a oneof named typ. Only a single field of the following list may be set at a time:
	//   - connector
	//   - manual
	//   - delegated
	//
	ProvisionPolicy *ProvisionPolicy `json:"provisionPolicy,omitempty"`
	// A field indicating whether this step is assigned.
	Assigned *bool `json:"assigned,omitempty"`
}

func (o *Provision) GetProvisionPolicy() *ProvisionPolicy {
	if o == nil {
		return nil
	}
	return o.ProvisionPolicy
}

func (o *Provision) GetAssigned() *bool {
	if o == nil {
		return nil
	}
	return o.Assigned
}
