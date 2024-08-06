// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// A FacetValue message contains count and value of the facet entry.
type FacetValue struct {
	// The count of the values in this facet.
	Count *int64 `integer:"string" json:"count,omitempty"`
	// The name of this facet.
	DisplayName *string `json:"displayName,omitempty"`
	// The icon for this facet.
	IconURL *string `json:"iconUrl,omitempty"`
	// The value of this facet.
	Value *string `json:"value,omitempty"`
}

func (o *FacetValue) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *FacetValue) GetDisplayName() *string {
	if o == nil {
		return nil
	}
	return o.DisplayName
}

func (o *FacetValue) GetIconURL() *string {
	if o == nil {
		return nil
	}
	return o.IconURL
}

func (o *FacetValue) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}
