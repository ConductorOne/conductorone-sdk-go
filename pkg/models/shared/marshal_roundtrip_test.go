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
		name      string
		json      string
		newTarget func() interface{ MarshalJSON() ([]byte, error) }
		validate  func(t *testing.T, marshaled string)
	}{
		{
			name:      "FacetValue with count as string",
			json:      `{"count":"42","displayName":"test","value":"val1"}`,
			newTarget: func() interface{ MarshalJSON() ([]byte, error) } { return &FacetValue{} },
			validate: func(t *testing.T, marshaled string) {
				var m map[string]json.RawMessage
				require.NoError(t, json.Unmarshal([]byte(marshaled), &m))
				assert.Equal(t, `"42"`, string(m["count"]), "count should be serialized as a string")
			},
		},
		{
			name:      "FacetRange with integer string fields",
			json:      `{"count":"10","from":"1","to":"100"}`,
			newTarget: func() interface{ MarshalJSON() ([]byte, error) } { return &FacetRange{} },
			validate: func(t *testing.T, marshaled string) {
				var m map[string]json.RawMessage
				require.NoError(t, json.Unmarshal([]byte(marshaled), &m))
				assert.Equal(t, `"10"`, string(m["count"]))
				assert.Equal(t, `"1"`, string(m["from"]))
				assert.Equal(t, `"100"`, string(m["to"]))
			},
		},
		{
			name:      "Facets with count",
			json:      `{"count":"5"}`,
			newTarget: func() interface{ MarshalJSON() ([]byte, error) } { return &Facets{} },
			validate: func(t *testing.T, marshaled string) {
				var m map[string]json.RawMessage
				require.NoError(t, json.Unmarshal([]byte(marshaled), &m))
				assert.Equal(t, `"5"`, string(m["count"]))
			},
		},
		{
			name:      "AutomationExecutionRef with id",
			json:      `{"id":"12345"}`,
			newTarget: func() interface{ MarshalJSON() ([]byte, error) } { return &AutomationExecutionRef{} },
			validate: func(t *testing.T, marshaled string) {
				var m map[string]json.RawMessage
				require.NoError(t, json.Unmarshal([]byte(marshaled), &m))
				assert.Equal(t, `"12345"`, string(m["id"]))
			},
		},
		{
			name:      "Int64Field with defaultValue",
			json:      `{"defaultValue":"999"}`,
			newTarget: func() interface{ MarshalJSON() ([]byte, error) } { return &Int64Field{} },
			validate: func(t *testing.T, marshaled string) {
				var m map[string]json.RawMessage
				require.NoError(t, json.Unmarshal([]byte(marshaled), &m))
				assert.Equal(t, `"999"`, string(m["defaultValue"]))
			},
		},
		{
			name:      "Escalation with expiration",
			json:      `{"expiration":"3600","escalationComment":"urgent"}`,
			newTarget: func() interface{ MarshalJSON() ([]byte, error) } { return &Escalation{} },
			validate: func(t *testing.T, marshaled string) {
				var m map[string]json.RawMessage
				require.NoError(t, json.Unmarshal([]byte(marshaled), &m))
				assert.Equal(t, `"3600"`, string(m["expiration"]))
			},
		},
		{
			name:      "Task with numericId",
			json:      `{"id":"task-1","numericId":"456","state":"TASK_STATE_OPEN"}`,
			newTarget: func() interface{ MarshalJSON() ([]byte, error) } { return &Task{} },
			validate: func(t *testing.T, marshaled string) {
				var m map[string]json.RawMessage
				require.NoError(t, json.Unmarshal([]byte(marshaled), &m))
				assert.Equal(t, `"456"`, string(m["numericId"]))
			},
		},
		{
			name:      "NumberField with integer string fields",
			json:      `{"maxValue":"100","minValue":"0","step":"1"}`,
			newTarget: func() interface{ MarshalJSON() ([]byte, error) } { return &NumberField{} },
			validate: func(t *testing.T, marshaled string) {
				var m map[string]json.RawMessage
				require.NoError(t, json.Unmarshal([]byte(marshaled), &m))
				assert.Equal(t, `"100"`, string(m["maxValue"]))
				assert.Equal(t, `"0"`, string(m["minValue"]))
				assert.Equal(t, `"1"`, string(m["step"]))
			},
		},
		{
			name:      "AppResourceInput with grantCount",
			json:      `{"grantCount":"7"}`,
			newTarget: func() interface{ MarshalJSON() ([]byte, error) } { return &AppResourceInput{} },
			validate: func(t *testing.T, marshaled string) {
				var m map[string]json.RawMessage
				require.NoError(t, json.Unmarshal([]byte(marshaled), &m))
				assert.Equal(t, `"7"`, string(m["grantCount"]))
			},
		},
		{
			name:      "ExecuteAutomationResponse with executionId",
			json:      `{"executionId":"123"}`,
			newTarget: func() interface{ MarshalJSON() ([]byte, error) } { return &ExecuteAutomationResponse{} },
			validate: func(t *testing.T, marshaled string) {
				var m map[string]json.RawMessage
				require.NoError(t, json.Unmarshal([]byte(marshaled), &m))
				assert.Equal(t, `"123"`, string(m["executionId"]))
			},
		},
		{
			name:      "FileField with maxFileSize",
			json:      `{"maxFileSize":"1048576"}`,
			newTarget: func() interface{ MarshalJSON() ([]byte, error) } { return &FileField{} },
			validate: func(t *testing.T, marshaled string) {
				var m map[string]json.RawMessage
				require.NoError(t, json.Unmarshal([]byte(marshaled), &m))
				assert.Equal(t, `"1048576"`, string(m["maxFileSize"]))
			},
		},
		{
			name:      "Int64Rules with integer string fields",
			json:      `{"const":"5","gt":"1","gte":"2","in":["3","4"],"lt":"10","lte":"9","notIn":["6","7"]}`,
			newTarget: func() interface{ MarshalJSON() ([]byte, error) } { return &Int64Rules{} },
			validate: func(t *testing.T, marshaled string) {
				var m map[string]json.RawMessage
				require.NoError(t, json.Unmarshal([]byte(marshaled), &m))
				assert.Equal(t, `"5"`, string(m["const"]))
				assert.Equal(t, `"1"`, string(m["gt"]))
				assert.Equal(t, `"2"`, string(m["gte"]))
				assert.Equal(t, `["3","4"]`, string(m["in"]))
				assert.Equal(t, `"10"`, string(m["lt"]))
				assert.Equal(t, `"9"`, string(m["lte"]))
				assert.Equal(t, `["6","7"]`, string(m["notIn"]))
			},
		},
		{
			name:      "PayloadWorkflowStep with workflowExecutionId",
			json:      `{"workflowExecutionId":"456"}`,
			newTarget: func() interface{ MarshalJSON() ([]byte, error) } { return &PayloadWorkflowStep{} },
			validate: func(t *testing.T, marshaled string) {
				var m map[string]json.RawMessage
				require.NoError(t, json.Unmarshal([]byte(marshaled), &m))
				assert.Equal(t, `"456"`, string(m["workflowExecutionId"]))
			},
		},
		{
			name:      "RequestCatalogView with memberCount",
			json:      `{"memberCount":"12"}`,
			newTarget: func() interface{ MarshalJSON() ([]byte, error) } { return &RequestCatalogView{} },
			validate: func(t *testing.T, marshaled string) {
				var m map[string]json.RawMessage
				require.NoError(t, json.Unmarshal([]byte(marshaled), &m))
				assert.Equal(t, `"12"`, string(m["memberCount"]))
			},
		},
		{
			name:      "SearchAutomationExecutionsRequest with executionId",
			json:      `{"executionId":"789"}`,
			newTarget: func() interface{ MarshalJSON() ([]byte, error) } { return &SearchAutomationExecutionsRequest{} },
			validate: func(t *testing.T, marshaled string) {
				var m map[string]json.RawMessage
				require.NoError(t, json.Unmarshal([]byte(marshaled), &m))
				assert.Equal(t, `"789"`, string(m["executionId"]))
			},
		},
		{
			name:      "SFixed64Rules with integer string fields",
			json:      `{"const":"5","gt":"1","gte":"2","in":["3","4"],"lt":"10","lte":"9","notIn":["6","7"]}`,
			newTarget: func() interface{ MarshalJSON() ([]byte, error) } { return &SFixed64Rules{} },
			validate: func(t *testing.T, marshaled string) {
				var m map[string]json.RawMessage
				require.NoError(t, json.Unmarshal([]byte(marshaled), &m))
				assert.Equal(t, `"5"`, string(m["const"]))
				assert.Equal(t, `["3","4"]`, string(m["in"]))
			},
		},
		{
			name:      "SInt64Rules with integer string fields",
			json:      `{"const":"5","gt":"1","gte":"2","in":["3","4"],"lt":"10","lte":"9","notIn":["6","7"]}`,
			newTarget: func() interface{ MarshalJSON() ([]byte, error) } { return &SInt64Rules{} },
			validate: func(t *testing.T, marshaled string) {
				var m map[string]json.RawMessage
				require.NoError(t, json.Unmarshal([]byte(marshaled), &m))
				assert.Equal(t, `"5"`, string(m["const"]))
				assert.Equal(t, `["3","4"]`, string(m["in"]))
			},
		},
		{
			name:      "TaskAuditErrorResult with errorCount",
			json:      `{"errorCount":"3"}`,
			newTarget: func() interface{ MarshalJSON() ([]byte, error) } { return &TaskAuditErrorResult{} },
			validate: func(t *testing.T, marshaled string) {
				var m map[string]json.RawMessage
				require.NoError(t, json.Unmarshal([]byte(marshaled), &m))
				assert.Equal(t, `"3"`, string(m["errorCount"]))
			},
		},
		{
			name:      "WebhookSourceWorkflowStep with workflowExecutionId",
			json:      `{"workflowExecutionId":"321"}`,
			newTarget: func() interface{ MarshalJSON() ([]byte, error) } { return &WebhookSourceWorkflowStep{} },
			validate: func(t *testing.T, marshaled string) {
				var m map[string]json.RawMessage
				require.NoError(t, json.Unmarshal([]byte(marshaled), &m))
				assert.Equal(t, `"321"`, string(m["workflowExecutionId"]))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Unmarshal from JSON (simulating API response) into a fresh instance
			target := tt.newTarget()
			err := json.Unmarshal([]byte(tt.json), target)
			require.NoError(t, err, "unmarshal should succeed")

			// Marshal back to JSON (the operation reported in IGA-719)
			data, err := json.Marshal(target)
			require.NoError(t, err, "marshal should succeed")

			// Validate the output
			tt.validate(t, string(data))

			// Verify full round-trip: unmarshal the marshaled data into a FRESH
			// instance (not the same pointer), then marshal again. This catches
			// stale-field or aliasing bugs that a same-pointer round-trip would miss.
			target2 := tt.newTarget()
			err = json.Unmarshal(data, target2)
			require.NoError(t, err, "second unmarshal should succeed")

			data2, err := json.Marshal(target2)
			require.NoError(t, err, "second marshal should succeed")
			assert.JSONEq(t, string(data), string(data2), "round-trip should produce identical JSON")
		})
	}
}

// TestMarshalRoundTrip_UserServiceListResponse is a smoke test that a
// UserServiceListResponse (a list of UserView objects) round-trips through
// JSON marshal/unmarshal. Note: the User/UserView types carry no
// integer:"string" fields, so this does NOT exercise the integer:"string"
// path — the integer:"string" coverage lives in
// TestMarshalRoundTrip_IntegerStringTags.
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
