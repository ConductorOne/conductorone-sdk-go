# TaskAuditRequestDefaultsApplied

TaskAuditRequestDefaultsApplied records which tier of the request-settings
 precedence chain supplied the defaults for a grant request. The rule ID, not
 its name, is stored; consumers resolve the current display name via
 (app_id, routing_rule_id).


## Fields

| Field                                                                                                                | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `AppID`                                                                                                              | `*string`                                                                                                            | :heavy_minus_sign:                                                                                                   | The appId field.                                                                                                     |
| `RoutingRuleID`                                                                                                      | `*string`                                                                                                            | :heavy_minus_sign:                                                                                                   | The routingRuleId field.                                                                                             |
| `Source`                                                                                                             | [*shared.TaskAuditRequestDefaultsAppliedSource](../../../pkg/models/shared/taskauditrequestdefaultsappliedsource.md) | :heavy_minus_sign:                                                                                                   | The source field.                                                                                                    |