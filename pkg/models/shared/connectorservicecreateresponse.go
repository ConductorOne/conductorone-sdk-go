// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ConnectorServiceCreateResponse -  The ConnectorServiceCreateResponse is the response returned from creating a connector.
type ConnectorServiceCreateResponse struct {
	//  The ConnectorView object provides a connector response object, as well as JSONPATHs to related objects provided by expanders.
	//
	ConnectorView *ConnectorView `json:"connectorView,omitempty"`
	//  The array of expanded items indicated by the request.
	//
	Expanded []map[string]interface{} `json:"expanded,omitempty"`
}

func (o *ConnectorServiceCreateResponse) GetConnectorView() *ConnectorView {
	if o == nil {
		return nil
	}
	return o.ConnectorView
}

func (o *ConnectorServiceCreateResponse) GetExpanded() []map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.Expanded
}
