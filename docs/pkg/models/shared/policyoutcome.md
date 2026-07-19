# PolicyOutcome

PolicyOutcome is the effect of a matched rule. Exactly one kind is set.

This message contains a oneof named kind. Only a single field of the following list may be set at a time:
  - allow
  - deny
  - stepUpRequired
  - challengeRequired
  - enrollmentRequired



## Fields

| Field                                                                          | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `Allow`                                                                        | [*shared.Allow](../../../pkg/models/shared/allow.md)                           | :heavy_minus_sign:                                                             | N/A                                                                            |
| `ChallengeRequired`                                                            | [*shared.ChallengeRequired](../../../pkg/models/shared/challengerequired.md)   | :heavy_minus_sign:                                                             | N/A                                                                            |
| `Deny`                                                                         | [*shared.Deny](../../../pkg/models/shared/deny.md)                             | :heavy_minus_sign:                                                             | N/A                                                                            |
| `EnrollmentRequired`                                                           | [*shared.EnrollmentRequired](../../../pkg/models/shared/enrollmentrequired.md) | :heavy_minus_sign:                                                             | N/A                                                                            |
| `StepUpRequired`                                                               | [*shared.StepUpRequired](../../../pkg/models/shared/stepuprequired.md)         | :heavy_minus_sign:                                                             | N/A                                                                            |