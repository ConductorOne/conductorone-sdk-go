package tools

import (
	"testing"

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
				Path{
					Key:   "AppPath",
					Value: strToPtr("as"),
				},
				Path{
					Key:   "AppResourcePath",
					Value: strToPtr("asd"),
				},
				Path{
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
