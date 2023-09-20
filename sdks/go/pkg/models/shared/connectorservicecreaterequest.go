// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ConnectorServiceCreateRequest - The ConnectorServiceCreateRequest message.
type ConnectorServiceCreateRequest struct {
	// The catalogId field.
	CatalogID *string `json:"catalogId,omitempty"`
	// Contains an arbitrary serialized message along with a @type that describes the type of the serialized message.
	Config map[string]interface{} `json:"config,omitempty"`
	// The description field.
	Description *string `json:"description,omitempty"`
	// The ConnectorExpandMask is used to expand related objects on a connector.
	ExpandMask *ConnectorExpandMask `json:"expandMask,omitempty"`
	// The userIds field.
	UserIds []string `json:"userIds,omitempty"`
}
