// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// The CIDRRestriction message.
type CIDRRestriction struct {
	// The enabled field.
	Enabled *bool `json:"enabled,omitempty"`
	// The sourceCidr field.
	SourceCidr []string `json:"sourceCidr,omitempty"`
}

func (o *CIDRRestriction) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *CIDRRestriction) GetSourceCidr() []string {
	if o == nil {
		return nil
	}
	return o.SourceCidr
}
