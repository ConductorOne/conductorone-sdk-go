// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// The ConnectorServiceCreateDelegatedRequest message contains the fields required to create a connector.
type ConnectorServiceCreateDelegatedRequest struct {
	// The AppManagedStateBindingRef message.
	AppManagedStateBindingRef *AppManagedStateBindingRef `json:"appManagedStateBindingRef,omitempty"`
	// The ConnectorExpandMask is used to expand related objects on a connector.
	ConnectorExpandMask *ConnectorExpandMask `json:"expandMask,omitempty"`
	// The catalogId describes which catalog entry this connector is an instance of. For example, every Okta connector will have the same catalogId indicating it is an Okta connector.
	CatalogID *string `json:"catalogId,omitempty"`
	// The description of the connector.
	Description *string `json:"description,omitempty"`
	// The displayName of the connector.
	DisplayName *string `json:"displayName,omitempty"`
	// The userIds field is used to define the integration owners of the connector.
	UserIds []string `json:"userIds,omitempty"`
}

func (o *ConnectorServiceCreateDelegatedRequest) GetAppManagedStateBindingRef() *AppManagedStateBindingRef {
	if o == nil {
		return nil
	}
	return o.AppManagedStateBindingRef
}

func (o *ConnectorServiceCreateDelegatedRequest) GetConnectorExpandMask() *ConnectorExpandMask {
	if o == nil {
		return nil
	}
	return o.ConnectorExpandMask
}

func (o *ConnectorServiceCreateDelegatedRequest) GetCatalogID() *string {
	if o == nil {
		return nil
	}
	return o.CatalogID
}

func (o *ConnectorServiceCreateDelegatedRequest) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *ConnectorServiceCreateDelegatedRequest) GetDisplayName() *string {
	if o == nil {
		return nil
	}
	return o.DisplayName
}

func (o *ConnectorServiceCreateDelegatedRequest) GetUserIds() []string {
	if o == nil {
		return nil
	}
	return o.UserIds
}
