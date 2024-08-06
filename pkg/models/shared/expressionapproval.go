// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// The ExpressionApproval message.
type ExpressionApproval struct {
	// Configuration to allow self approval of if the user is specified and also the target of the ticket.
	AllowSelfApproval *bool `json:"allowSelfApproval,omitempty"`
	// The assignedUserIds field.
	AssignedUserIds []string `json:"assignedUserIds,omitempty"`
	// Array of dynamic expressions to determine the approvers.  The first expression to return a non-empty list of users will be used.
	Expressions []string `json:"expressions,omitempty"`
	// Configuration to allow a fallback if the expression does not return a valid list of users.
	Fallback *bool `json:"fallback,omitempty"`
	// Configuration to specific which users to fallback to if and the expression does not return a valid list of users.
	FallbackUserIds []string `json:"fallbackUserIds,omitempty"`
}

func (o *ExpressionApproval) GetAllowSelfApproval() *bool {
	if o == nil {
		return nil
	}
	return o.AllowSelfApproval
}

func (o *ExpressionApproval) GetAssignedUserIds() []string {
	if o == nil {
		return nil
	}
	return o.AssignedUserIds
}

func (o *ExpressionApproval) GetExpressions() []string {
	if o == nil {
		return nil
	}
	return o.Expressions
}

func (o *ExpressionApproval) GetFallback() *bool {
	if o == nil {
		return nil
	}
	return o.Fallback
}

func (o *ExpressionApproval) GetFallbackUserIds() []string {
	if o == nil {
		return nil
	}
	return o.FallbackUserIds
}
