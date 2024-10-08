// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// ResourceType - The resourceType field.
type ResourceType string

const (
	ResourceTypeRole    ResourceType = "ROLE"
	ResourceTypeGroup   ResourceType = "GROUP"
	ResourceTypeLicense ResourceType = "LICENSE"
	ResourceTypeProject ResourceType = "PROJECT"
	ResourceTypeCatalog ResourceType = "CATALOG"
	ResourceTypeCustom  ResourceType = "CUSTOM"
)

func (e ResourceType) ToPointer() *ResourceType {
	return &e
}

// The CreateManuallyManagedResourceTypeRequest message.
type CreateManuallyManagedResourceTypeRequest struct {
	// The name field.
	Name *string `json:"name,omitempty"`
	// The resourceType field.
	ResourceType *ResourceType `json:"resourceType,omitempty"`
}

func (o *CreateManuallyManagedResourceTypeRequest) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *CreateManuallyManagedResourceTypeRequest) GetResourceType() *ResourceType {
	if o == nil {
		return nil
	}
	return o.ResourceType
}
