# Escalation

The Escalation message.

This message contains a oneof named escalation_policy. Only a single field of the following list may be set at a time:
  - replacePolicy
  - reassignToApprovers
  - cancelTicket
  - skipStep



## Fields

| Field                                                                            | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `CancelTicket`                                                                   | [*shared.CancelTicket](../../../pkg/models/shared/cancelticket.md)               | :heavy_minus_sign:                                                               | N/A                                                                              |
| `EscalationComment`                                                              | `*string`                                                                        | :heavy_minus_sign:                                                               | The escalationComment field.                                                     |
| `Expiration`                                                                     | `*int64`                                                                         | :heavy_minus_sign:                                                               | The expiration field.                                                            |
| `ReassignToApprovers`                                                            | [*shared.ReassignToApprovers](../../../pkg/models/shared/reassigntoapprovers.md) | :heavy_minus_sign:                                                               | N/A                                                                              |
| `ReplacePolicy`                                                                  | [*shared.ReplacePolicy](../../../pkg/models/shared/replacepolicy.md)             | :heavy_minus_sign:                                                               | N/A                                                                              |
| `SkipStep`                                                                       | [*shared.SkipStep](../../../pkg/models/shared/skipstep.md)                       | :heavy_minus_sign:                                                               | N/A                                                                              |