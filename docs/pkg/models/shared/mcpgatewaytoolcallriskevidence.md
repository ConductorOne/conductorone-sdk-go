# McpGatewayToolCallRiskEvidence

McpGatewayToolCallRiskEvidence carries the rule that fired and the
 subject/summary the UI renders. No detector writes this yet -- see
 FINDING_TYPE_MCP_GATEWAY_TOOL_CALL_RISK.


## Fields

| Field                                                                       | Type                                                                        | Required                                                                    | Description                                                                 |
| --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- |
| `Evidence`                                                                  | [][shared.KeyValue](../../../pkg/models/shared/keyvalue.md)                 | :heavy_minus_sign:                                                          | The evidence field.                                                         |
| `Kind`                                                                      | `*string`                                                                   | :heavy_minus_sign:                                                          | Human-readable rule category, e.g. "Prompt injection in tool output".       |
| `RuleID`                                                                    | `*string`                                                                   | :heavy_minus_sign:                                                          | Detection rule id, e.g. "MCP-201".                                          |
| `Subject`                                                                   | `*string`                                                                   | :heavy_minus_sign:                                                          | Human-readable subject, e.g. "agent-support-triage -> okta/reset_password". |
| `Summary`                                                                   | `*string`                                                                   | :heavy_minus_sign:                                                          | The summary field.                                                          |