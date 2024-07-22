// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/utils"
)

// RequestCatalogManagementServiceGetResponseExpanded - Contains an arbitrary serialized message along with a @type that describes the type of the serialized message.
type RequestCatalogManagementServiceGetResponseExpanded struct {
	// The type of the serialized message.
	AtType               *string        `json:"@type,omitempty"`
	AdditionalProperties map[string]any `additionalProperties:"true" json:"-"`
}

func (r RequestCatalogManagementServiceGetResponseExpanded) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RequestCatalogManagementServiceGetResponseExpanded) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *RequestCatalogManagementServiceGetResponseExpanded) GetAtType() *string {
	if o == nil {
		return nil
	}
	return o.AtType
}

func (o *RequestCatalogManagementServiceGetResponseExpanded) GetAdditionalProperties() map[string]any {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

// RequestCatalogManagementServiceGetResponse - The request catalog management service get response returns a request catalog view with the expanded items in the expanded array indicated by the expand mask in the request.
type RequestCatalogManagementServiceGetResponse struct {
	// The request catalog view contains the serialized request catalog and paths to objects referenced by the request catalog.
	RequestCatalogView *RequestCatalogView `json:"requestCatalogView,omitempty"`
	// List of serialized related objects.
	Expanded []RequestCatalogManagementServiceGetResponseExpanded `json:"expanded,omitempty"`
}

func (o *RequestCatalogManagementServiceGetResponse) GetRequestCatalogView() *RequestCatalogView {
	if o == nil {
		return nil
	}
	return o.RequestCatalogView
}

func (o *RequestCatalogManagementServiceGetResponse) GetExpanded() []RequestCatalogManagementServiceGetResponseExpanded {
	if o == nil {
		return nil
	}
	return o.Expanded
}
