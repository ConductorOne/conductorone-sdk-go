package utils

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestUnmarshalMapWithNestedStringInt64 reproduces the issue where a map[string]Struct
// containing nested structs with integer:"string" tags fails to unmarshal because the
// map handling falls through to standard json.Unmarshal which doesn't understand the tag.
// This is the root cause of the Escalation.Expiration unmarshal failure (IGA-731).
func TestUnmarshalMapWithNestedStringInt64(t *testing.T) {
	// Inner struct with integer:"string" tag (like Escalation)
	type Escalation struct {
		Expiration *int64 `integer:"string" json:"expiration,omitempty"`
	}

	// Intermediate struct (like Approval) - no custom UnmarshalJSON
	type Approval struct {
		Escalation *Escalation `json:"escalation,omitempty"`
	}

	// Step struct (like PolicyStep)
	type PolicyStep struct {
		Approval *Approval `json:"approval,omitempty"`
	}

	// Steps container (like PolicySteps) - used as map value
	type PolicySteps struct {
		Steps []PolicyStep `json:"steps,omitempty"`
	}

	// Top-level struct (like Policy)
	type Policy struct {
		PolicySteps map[string]PolicySteps `json:"policySteps,omitempty"`
	}

	input := `{
		"policySteps": {
			"grant": {
				"steps": [{
					"approval": {
						"escalation": {
							"expiration": "3600"
						}
					}
				}]
			}
		}
	}`

	var policy Policy
	err := UnmarshalJSON([]byte(input), &policy, "", false, nil)
	require.NoError(t, err, "should unmarshal string-encoded int64 in map values")

	require.NotNil(t, policy.PolicySteps)
	steps, ok := policy.PolicySteps["grant"]
	require.True(t, ok)
	require.Len(t, steps.Steps, 1)
	require.NotNil(t, steps.Steps[0].Approval)
	require.NotNil(t, steps.Steps[0].Approval.Escalation)
	require.NotNil(t, steps.Steps[0].Approval.Escalation.Expiration)
	assert.Equal(t, int64(3600), *steps.Steps[0].Approval.Escalation.Expiration)
}

// TestUnmarshalMapWithNestedStringInt64_NumericValue ensures the fix also handles
// the case where the API returns a numeric value instead of a string.
func TestUnmarshalMapWithNestedStringInt64_NumericValue(t *testing.T) {
	type Inner struct {
		Value *int64 `integer:"string" json:"value,omitempty"`
	}
	type Outer struct {
		Items map[string]Inner `json:"items,omitempty"`
	}

	// When the API returns a string-encoded integer
	input := `{"items": {"a": {"value": "42"}}}`
	var out Outer
	err := UnmarshalJSON([]byte(input), &out, "", false, nil)
	require.NoError(t, err)
	require.NotNil(t, out.Items["a"].Value)
	assert.Equal(t, int64(42), *out.Items["a"].Value)
}

// TestMarshalMapWithNestedStringInt64 ensures marshaling also works correctly.
func TestMarshalMapWithNestedStringInt64(t *testing.T) {
	type Inner struct {
		Value *int64 `integer:"string" json:"value,omitempty"`
	}
	type Outer struct {
		Items map[string]Inner `json:"items,omitempty"`
	}

	val := int64(42)
	out := Outer{
		Items: map[string]Inner{
			"a": {Value: &val},
		},
	}

	data, err := MarshalJSON(out, "", false)
	require.NoError(t, err)

	var raw map[string]json.RawMessage
	require.NoError(t, json.Unmarshal(data, &raw))

	var items map[string]json.RawMessage
	require.NoError(t, json.Unmarshal(raw["items"], &items))

	// The value should be marshaled as a string due to integer:"string" tag
	assert.Equal(t, `{"value":"42"}`, string(items["a"]))
}
