// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// The RequestCatalogExpandMask includes the paths in the catalog view to expand in the return value of this call.
type RequestCatalogExpandMask struct {
	// An array of paths to be expanded in the response. May be any combination of "*", "created_by_user_id", "app_ids", and "access_entitlements".
	Paths []string `json:"paths,omitempty"`
}

func (o *RequestCatalogExpandMask) GetPaths() []string {
	if o == nil {
		return nil
	}
	return o.Paths
}
