// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// TaskExpandMask -  The task expand mask is an array of strings that specifes the related objects the requester wishes to have returned when making a request where the expand mask is part of the input. Use '*' to view all possible responses.
type TaskExpandMask struct {
	//  A list of IDs associated with a task
	//
	Paths []string `json:"paths,omitempty"`
}

func (o *TaskExpandMask) GetPaths() []string {
	if o == nil {
		return nil
	}
	return o.Paths
}
