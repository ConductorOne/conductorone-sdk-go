// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// DirectoryServiceListResponse - The DirectoryServiceListResponse message.
type DirectoryServiceListResponse struct {
	// The expanded field.
	Expanded []map[string]interface{} `json:"expanded,omitempty"`
	// The list field.
	List []DirectoryView `json:"list,omitempty"`
	// The nextPageToken field.
	NextPageToken *string `json:"nextPageToken,omitempty"`
}