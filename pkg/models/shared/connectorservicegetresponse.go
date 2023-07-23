// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ConnectorServiceGetResponse - The ConnectorServiceGetResponse message.
type ConnectorServiceGetResponse struct {
	// The ConnectorView message.
	ConnectorView *ConnectorView `json:"connectorView,omitempty"`
	// The expanded field.
	Expanded []map[string]interface{} `json:"expanded,omitempty"`
}

func (o *ConnectorServiceGetResponse) GetConnectorView() *ConnectorView {
	if o == nil {
		return nil
	}
	return o.ConnectorView
}

func (o *ConnectorServiceGetResponse) GetExpanded() []map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.Expanded
}
