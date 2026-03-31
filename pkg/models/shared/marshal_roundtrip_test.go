package shared

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestMarshalRoundTrip_IntegerStringTags verifies that types with integer:"string"
// tags correctly round-trip through JSON marshal/unmarshal.
// This is a regression test for https://linear.app/conductorone/issue/IGA-719
func TestMarshalRoundTrip_IntegerStringTags(t *testing.T) {
	tests := []struct {
		name     string
		json     string
		target   interface{ MarshalJSON() ([]byte, error) }
		validate func(t *testing.T, marshaled string)
	}{
		{
			name:   "FacetValue with count as string",
			json:   `{"count":"42","displayName":"test","value":"val1"}`,
			target: &FacetValue{},
			validate: func(t *testing.T, marshaled string) {
				var m map[string]json.RawMessage
				require.NoError(t, json.Unmarshal([]byte(marshaled), &m))
				assert.Equal(t, `"42"`, string(m["count"]), "count should be serialized as a string")
			},
		},
		{
			name:   "FacetRange with integer string fields",
			json:   `{"count":"10","from":"1","to":"100"}`,
			target: &FacetRange{},
			validate: func(t *testing.T, marshaled string) {
				var m map[string]json.RawMessage
				require.NoError(t, json.Unmarshal([]byte(marshaled), &m))
				assert.Equal(t, `"10"`, string(m["count"]))
				assert.Equal(t, `"1"`, string(m["from"]))
				assert.Equal(t, `"100"`, string(m["to"]))
			},
		},
		{
			name:   "Facets with count",
			json:   `{"count":"5","displayName":"category","fieldName":"type"}`,
			target: &Facets{},
			validate: func(t *testing.T, marshaled string) {
				var m map[string]json.RawMessage
				require.NoError(t, json.Unmarshal([]byte(marshaled), &m))
				assert.Equal(t, `"5"`, string(m["count"]))
			},
		},
		{
			name:   "AutomationExecutionRef with id",
			json:   `{"id":"12345"}`,
			target: &AutomationExecutionRef{},
			validate: func(t *testing.T, marshaled string) {
				var m map[string]json.RawMessage
				require.NoError(t, json.Unmarshal([]byte(marshaled), &m))
				assert.Equal(t, `"12345"`, string(m["id"]))
			},
		},
		{
			name:   "Int64Field with defaultValue",
			json:   `{"defaultValue":"999"}`,
			target: &Int64Field{},
			validate: func(t *testing.T, marshaled string) {
				var m map[string]json.RawMessage
				require.NoError(t, json.Unmarshal([]byte(marshaled), &m))
				assert.Equal(t, `"999"`, string(m["defaultValue"]))
			},
		},
		{
			name:   "Escalation with expiration",
			json:   `{"expiration":"3600","escalationComment":"urgent"}`,
			target: &Escalation{},
			validate: func(t *testing.T, marshaled string) {
				var m map[string]json.RawMessage
				require.NoError(t, json.Unmarshal([]byte(marshaled), &m))
				assert.Equal(t, `"3600"`, string(m["expiration"]))
			},
		},
		{
			name:   "Task with numericId",
			json:   `{"id":"task-1","numericId":"456","state":"TASK_STATE_OPEN"}`,
			target: &Task{},
			validate: func(t *testing.T, marshaled string) {
				var m map[string]json.RawMessage
				require.NoError(t, json.Unmarshal([]byte(marshaled), &m))
				assert.Equal(t, `"456"`, string(m["numericId"]))
			},
		},
		{
			name:   "NumberField with integer string fields",
			json:   `{"maxValue":"100","minValue":"0","step":"1"}`,
			target: &NumberField{},
			validate: func(t *testing.T, marshaled string) {
				var m map[string]json.RawMessage
				require.NoError(t, json.Unmarshal([]byte(marshaled), &m))
				assert.Equal(t, `"100"`, string(m["maxValue"]))
				assert.Equal(t, `"0"`, string(m["minValue"]))
				assert.Equal(t, `"1"`, string(m["step"]))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Unmarshal from JSON (simulating API response)
			err := json.Unmarshal([]byte(tt.json), tt.target)
			require.NoError(t, err, "unmarshal should succeed")

			// Marshal back to JSON (the operation reported in IGA-719)
			data, err := json.Marshal(tt.target)
			require.NoError(t, err, "marshal should succeed")

			// Validate the output
			tt.validate(t, string(data))

			// Verify full round-trip: unmarshal the marshaled data again
			target2 := tt.target
			err = json.Unmarshal(data, target2)
			require.NoError(t, err, "second unmarshal should succeed")

			data2, err := json.Marshal(target2)
			require.NoError(t, err, "second marshal should succeed")
			assert.JSONEq(t, string(data), string(data2), "round-trip should produce identical JSON")
		})
	}
}

// TestMarshalRoundTrip_UserServiceListResponse tests the exact scenario
// from the issue: List Users -> Marshal back to JSON.
func TestMarshalRoundTrip_UserServiceListResponse(t *testing.T) {
	input := `{
		"list": [
			{
				"user": {
					"id": "user-123",
					"email": "test@example.com",
					"displayName": "Test User",
					"status": "ENABLED",
					"directoryStatus": "ENABLED",
					"department": "Engineering"
				}
			}
		],
		"nextPageToken": "token123"
	}`

	var resp UserServiceListResponse
	err := json.Unmarshal([]byte(input), &resp)
	require.NoError(t, err)

	data, err := json.Marshal(resp)
	require.NoError(t, err)

	// Verify key fields are preserved
	var check map[string]json.RawMessage
	require.NoError(t, json.Unmarshal(data, &check))
	assert.Contains(t, string(check["list"]), `"user-123"`)
	assert.Contains(t, string(check["list"]), `"test@example.com"`)
	assert.Equal(t, `"token123"`, string(check["nextPageToken"]))

	// Verify second round-trip
	var resp2 UserServiceListResponse
	require.NoError(t, json.Unmarshal(data, &resp2))
	data2, err := json.Marshal(resp2)
	require.NoError(t, err)
	assert.JSONEq(t, string(data), string(data2))
}
