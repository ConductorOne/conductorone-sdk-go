// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ConnectorExpandMask - The ConnectorExpandMask message.
type ConnectorExpandMask struct {
	// The paths field.
	Paths []string `json:"paths,omitempty"`
}

func (o *ConnectorExpandMask) GetPaths() []string {
	if o == nil {
		return nil
	}
	return o.Paths
}
