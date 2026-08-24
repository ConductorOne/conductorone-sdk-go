package shared

import (
	"encoding/json"
	"testing"
)

// TestMarshalRoundTrip_IntegerStringTags verifies that types with
// integer:"string" / number:"string" struct tags (where the API serializes
// int64 values as JSON strings like "12345") round-trip correctly through
// json.Marshal / json.Unmarshal. Each case unmarshals a JSON payload where the
// integer:"string" field is a JSON string, marshals it back, and asserts the
// field is still serialized as a JSON string.
func TestMarshalRoundTrip_IntegerStringTags(t *testing.T) {
	tests := []struct {
		name     string
		json     string
		target   interface{}
		validate func(t *testing.T, out map[string]json.RawMessage)
	}{
		{
			name:   "FacetValue",
			json:   `{"count":"42","displayName":"test","value":"val1"}`,
			target: &FacetValue{},
			validate: func(t *testing.T, out map[string]json.RawMessage) {
				if got := string(out["count"]); got != `"42"` {
					t.Errorf("count = %s, want \"42\"", got)
				}
			},
		},
		{
			name:   "FacetRange",
			json:   `{"count":"10","from":"1","to":"100"}`,
			target: &FacetRange{},
			validate: func(t *testing.T, out map[string]json.RawMessage) {
				if got := string(out["count"]); got != `"10"` {
					t.Errorf("count = %s, want \"10\"", got)
				}
				if got := string(out["from"]); got != `"1"` {
					t.Errorf("from = %s, want \"1\"", got)
				}
				if got := string(out["to"]); got != `"100"` {
					t.Errorf("to = %s, want \"100\"", got)
				}
			},
		},
		{
			name:   "Facets",
			json:   `{"count":"5","displayName":"category","fieldName":"type"}`,
			target: &Facets{},
			validate: func(t *testing.T, out map[string]json.RawMessage) {
				if got := string(out["count"]); got != `"5"` {
					t.Errorf("count = %s, want \"5\"", got)
				}
			},
		},
		{
			name:   "AutomationExecutionRef",
			json:   `{"id":"12345"}`,
			target: &AutomationExecutionRef{},
			validate: func(t *testing.T, out map[string]json.RawMessage) {
				if got := string(out["id"]); got != `"12345"` {
					t.Errorf("id = %s, want \"12345\"", got)
				}
			},
		},
		{
			name:   "Int64Field",
			json:   `{"defaultValue":"999"}`,
			target: &Int64Field{},
			validate: func(t *testing.T, out map[string]json.RawMessage) {
				if got := string(out["defaultValue"]); got != `"999"` {
					t.Errorf("defaultValue = %s, want \"999\"", got)
				}
			},
		},
		{
			name:   "NumberField",
			json:   `{"maxValue":"100","minValue":"0","step":"1"}`,
			target: &NumberField{},
			validate: func(t *testing.T, out map[string]json.RawMessage) {
				if got := string(out["maxValue"]); got != `"100"` {
					t.Errorf("maxValue = %s, want \"100\"", got)
				}
				if got := string(out["minValue"]); got != `"0"` {
					t.Errorf("minValue = %s, want \"0\"", got)
				}
				if got := string(out["step"]); got != `"1"` {
					t.Errorf("step = %s, want \"1\"", got)
				}
			},
		},
		{
			name:   "AppResourceInput",
			json:   `{"grantCount":"7"}`,
			target: &AppResourceInput{},
			validate: func(t *testing.T, out map[string]json.RawMessage) {
				if got := string(out["grantCount"]); got != `"7"` {
					t.Errorf("grantCount = %s, want \"7\"", got)
				}
			},
		},
		{
			name:   "ExecuteAutomationResponse",
			json:   `{"executionId":"123"}`,
			target: &ExecuteAutomationResponse{},
			validate: func(t *testing.T, out map[string]json.RawMessage) {
				if got := string(out["executionId"]); got != `"123"` {
					t.Errorf("executionId = %s, want \"123\"", got)
				}
			},
		},
		{
			name:   "FileField",
			json:   `{"maxFileSize":"1048576"}`,
			target: &FileField{},
			validate: func(t *testing.T, out map[string]json.RawMessage) {
				if got := string(out["maxFileSize"]); got != `"1048576"` {
					t.Errorf("maxFileSize = %s, want \"1048576\"", got)
				}
			},
		},
		{
			name:   "Int64Rules",
			json:   `{"const":"5","gt":"1","gte":"2","in":["3","4"],"lt":"10","lte":"9","notIn":["6","7"]}`,
			target: &Int64Rules{},
			validate: func(t *testing.T, out map[string]json.RawMessage) {
				if got := string(out["const"]); got != `"5"` {
					t.Errorf("const = %s, want \"5\"", got)
				}
				if got := string(out["gt"]); got != `"1"` {
					t.Errorf("gt = %s, want \"1\"", got)
				}
				if got := string(out["gte"]); got != `"2"` {
					t.Errorf("gte = %s, want \"2\"", got)
				}
				if got := string(out["in"]); got != `["3","4"]` {
					t.Errorf("in = %s, want [\"3\",\"4\"]", got)
				}
				if got := string(out["lt"]); got != `"10"` {
					t.Errorf("lt = %s, want \"10\"", got)
				}
				if got := string(out["lte"]); got != `"9"` {
					t.Errorf("lte = %s, want \"9\"", got)
				}
				if got := string(out["notIn"]); got != `["6","7"]` {
					t.Errorf("notIn = %s, want [\"6\",\"7\"]", got)
				}
			},
		},
		{
			name:   "PayloadWorkflowStep",
			json:   `{"workflowExecutionId":"456"}`,
			target: &PayloadWorkflowStep{},
			validate: func(t *testing.T, out map[string]json.RawMessage) {
				if got := string(out["workflowExecutionId"]); got != `"456"` {
					t.Errorf("workflowExecutionId = %s, want \"456\"", got)
				}
			},
		},
		{
			name:   "RequestCatalogView",
			json:   `{"memberCount":"12"}`,
			target: &RequestCatalogView{},
			validate: func(t *testing.T, out map[string]json.RawMessage) {
				if got := string(out["memberCount"]); got != `"12"` {
					t.Errorf("memberCount = %s, want \"12\"", got)
				}
			},
		},
		{
			name:   "SearchAutomationExecutionsRequest",
			json:   `{"executionId":"789"}`,
			target: &SearchAutomationExecutionsRequest{},
			validate: func(t *testing.T, out map[string]json.RawMessage) {
				if got := string(out["executionId"]); got != `"789"` {
					t.Errorf("executionId = %s, want \"789\"", got)
				}
			},
		},
		{
			name:   "SFixed64Rules",
			json:   `{"const":"5","gt":"1","gte":"2","in":["3","4"],"lt":"10","lte":"9","notIn":["6","7"]}`,
			target: &SFixed64Rules{},
			validate: func(t *testing.T, out map[string]json.RawMessage) {
				if got := string(out["const"]); got != `"5"` {
					t.Errorf("const = %s, want \"5\"", got)
				}
				if got := string(out["in"]); got != `["3","4"]` {
					t.Errorf("in = %s, want [\"3\",\"4\"]", got)
				}
			},
		},
		{
			name:   "SInt64Rules",
			json:   `{"const":"5","gt":"1","gte":"2","in":["3","4"],"lt":"10","lte":"9","notIn":["6","7"]}`,
			target: &SInt64Rules{},
			validate: func(t *testing.T, out map[string]json.RawMessage) {
				if got := string(out["const"]); got != `"5"` {
					t.Errorf("const = %s, want \"5\"", got)
				}
				if got := string(out["in"]); got != `["3","4"]` {
					t.Errorf("in = %s, want [\"3\",\"4\"]", got)
				}
			},
		},
		{
			name:   "TaskAuditErrorResult",
			json:   `{"errorCount":"3"}`,
			target: &TaskAuditErrorResult{},
			validate: func(t *testing.T, out map[string]json.RawMessage) {
				if got := string(out["errorCount"]); got != `"3"` {
					t.Errorf("errorCount = %s, want \"3\"", got)
				}
			},
		},
		{
			name:   "WebhookSourceWorkflowStep",
			json:   `{"workflowExecutionId":"321"}`,
			target: &WebhookSourceWorkflowStep{},
			validate: func(t *testing.T, out map[string]json.RawMessage) {
				if got := string(out["workflowExecutionId"]); got != `"321"` {
					t.Errorf("workflowExecutionId = %s, want \"321\"", got)
				}
			},
		},
		// Pre-existing cases: Escalation and Task already implement
		// MarshalJSON (from PR #97 and earlier work). They are NOT among the
		// 17 affected types and must not be modified.
		{
			name:   "Escalation",
			json:   `{"expiration":"1234567890"}`,
			target: &Escalation{},
			validate: func(t *testing.T, out map[string]json.RawMessage) {
				if got := string(out["expiration"]); got != `"1234567890"` {
					t.Errorf("expiration = %s, want \"1234567890\"", got)
				}
			},
		},
		{
			name:   "Task",
			json:   `{"numericId":"42"}`,
			target: &Task{},
			validate: func(t *testing.T, out map[string]json.RawMessage) {
				if got := string(out["numericId"]); got != `"42"` {
					t.Errorf("numericId = %s, want \"42\"", got)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := json.Unmarshal([]byte(tt.json), tt.target); err != nil {
				t.Fatalf("unmarshal: %v", err)
			}
			out, err := json.Marshal(tt.target)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var m map[string]json.RawMessage
			if err := json.Unmarshal(out, &m); err != nil {
				t.Fatalf("unmarshal map: %v", err)
			}
			tt.validate(t, m)
		})
	}
}

// TestMarshalRoundTrip_UserServiceListResponse verifies a full
// UserServiceListResponse (containing a UserView with a User) round-trips
// through json.Marshal / json.Unmarshal.
func TestMarshalRoundTrip_UserServiceListResponse(t *testing.T) {
	userID := "user-123"
	email := "user@example.com"
	displayName := "Test User"
	nextPageToken := "token-1"

	in := UserServiceListResponse{
		List: []UserView{
			{
				UserID: &userID,
				User: &User{
					ID:          &userID,
					Email:       &email,
					DisplayName: &displayName,
				},
			},
		},
		NextPageToken: &nextPageToken,
	}

	out, err := json.Marshal(&in)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}

	var got UserServiceListResponse
	if err := json.Unmarshal(out, &got); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}

	if len(got.List) != 1 {
		t.Fatalf("list length = %d, want 1", len(got.List))
	}
	if got.List[0].UserID == nil || *got.List[0].UserID != userID {
		t.Errorf("list[0].userId = %v, want %q", got.List[0].UserID, userID)
	}
	if got.List[0].User == nil || got.List[0].User.Email == nil || *got.List[0].User.Email != email {
		t.Errorf("list[0].user.email = %v, want %q", got.List[0].User, email)
	}
	if got.NextPageToken == nil || *got.NextPageToken != nextPageToken {
		t.Errorf("nextPageToken = %v, want %q", got.NextPageToken, nextPageToken)
	}
}
