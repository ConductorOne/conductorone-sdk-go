package tools

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

func strToPtr(str string) *string {
	return &str
}

func TestGetPaths(t *testing.T) {
	tests := []struct {
		input shared.AppEntitlementView
		want  []Path
	}{
		{
			input: shared.AppEntitlementView{
				AppPath:             strToPtr("as"),
				AppResourcePath:     strToPtr("asd"),
				AppResourceTypePath: strToPtr("asdads"),
			},
			want: []Path{
				{
					Key:   "AppPath",
					Value: strToPtr("as"),
				},
				{
					Key:   "AppResourcePath",
					Value: strToPtr("asd"),
				},
				{
					Key:   "AppResourceTypePath",
					Value: strToPtr("asdads"),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run("TestGetPaths", func(t *testing.T) {
			got := GetPaths(&tt.input)
			for i, gotPath := range got {
				if gotPath.Key != tt.want[i].Key {
					t.Errorf("GetPaths() Key Error = %v, want %v", gotPath.Key, tt.want[i].Key)
				}
				if *gotPath.Value != *tt.want[i].Value {
					t.Errorf("GetPaths() Value Error = %v, want %v", *gotPath.Value, tt.want[i].Value)
				}
			}
		})
	}
}

func TestGetMarshalledObject(t *testing.T) {
	// Mock data
	mockAppID := "123"
	now := time.Now()
	mockAppResourceType := shared.AppResourceType{
		AppID:       &mockAppID,
		CreatedAt:   &now,
		DisplayName: strToPtr("Example App Resource Type"),
		ID:          strToPtr("456"),
		UpdatedAt:   &now,
	}

	// Marshal to JSON
	jsonBytes, err := json.Marshal(mockAppResourceType)
	if err != nil {
		t.Error("Error marshaling AppResourceType to JSON:", err)
		return
	}

	// Unmarshal into map[string]interface{}
	var result map[string]interface{}
	err = json.Unmarshal(jsonBytes, &result)
	if err != nil {
		t.Error("Error unmarshaling JSON into map[string]interface{}:", err)
		return
	}

	mockResponse := shared.RequestCatalogSearchServiceSearchEntitlementsResponseExpanded{
		AtType:               strToPtr(atTypeAppResourceType),
		AdditionalProperties: result,
	}

	tests := []struct {
		name    string
		input   shared.RequestCatalogSearchServiceSearchEntitlementsResponseExpanded
		wantErr bool
	}{
		{
			name:    "Valid AppResourceType",
			input:   mockResponse,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GetMarshalledObject(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMarshalledObject() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}