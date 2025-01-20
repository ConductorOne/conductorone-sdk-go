// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// The WaitCondition message.
type WaitCondition struct {
	// The condition that has to be true for this wait condition to continue.
	Condition *string `json:"condition,omitempty"`
}

func (o *WaitCondition) GetCondition() *string {
	if o == nil {
		return nil
	}
	return o.Condition
}