# EscalationInstance

The EscalationInstance message.

This message contains a oneof named escalation_policy. Only a single field of the following list may be set at a time:
  - replacePolicy
  - reassignToApprovers
  - cancelTicket
  - skipStep



## Fields

| Field                                                                            | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `CancelTicket`                                                                   | [*shared.CancelTicket](../../../pkg/models/shared/cancelticket.md)               | :heavy_minus_sign:                                                               | The CancelTicket message.                                                        |
| `ReassignToApprovers`                                                            | [*shared.ReassignToApprovers](../../../pkg/models/shared/reassigntoapprovers.md) | :heavy_minus_sign:                                                               | The ReassignToApprovers message.                                                 |
| `ReplacePolicy`                                                                  | [*shared.ReplacePolicy](../../../pkg/models/shared/replacepolicy.md)             | :heavy_minus_sign:                                                               | The ReplacePolicy message.                                                       |
| `SkipStep`                                                                       | [*shared.SkipStep](../../../pkg/models/shared/skipstep.md)                       | :heavy_minus_sign:                                                               | The SkipStep message.                                                            |
| `AlreadyEscalated`                                                               | **bool*                                                                          | :heavy_minus_sign:                                                               | The alreadyEscalated field.                                                      |
| `EscalationComment`                                                              | **string*                                                                        | :heavy_minus_sign:                                                               | The escalationComment field.                                                     |
| `ExpiresAt`                                                                      | [*time.Time](https://pkg.go.dev/time#Time)                                       | :heavy_minus_sign:                                                               | N/A                                                                              |