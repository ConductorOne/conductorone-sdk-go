# SessionPolicyPolicyOutcome

PolicyOutcome is the effect of a matched rule. Exactly one kind is set. For
 session continuous-evaluation, the meaningful kinds are Allow (continue),
 Deny (terminate), and StepUpRequired.

This message contains a oneof named kind. Only a single field of the following list may be set at a time:
  - allow
  - deny
  - stepUpRequired
  - challengeRequired
  - enrollmentRequired



## Fields

| Field                                                                                                    | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `Allow`                                                                                                  | [*shared.SessionPolicyAllow](../../../pkg/models/shared/sessionpolicyallow.md)                           | :heavy_minus_sign:                                                                                       | N/A                                                                                                      |
| `ChallengeRequired`                                                                                      | [*shared.SessionPolicyChallengeRequired](../../../pkg/models/shared/sessionpolicychallengerequired.md)   | :heavy_minus_sign:                                                                                       | N/A                                                                                                      |
| `Deny`                                                                                                   | [*shared.SessionPolicyDeny](../../../pkg/models/shared/sessionpolicydeny.md)                             | :heavy_minus_sign:                                                                                       | N/A                                                                                                      |
| `EnrollmentRequired`                                                                                     | [*shared.SessionPolicyEnrollmentRequired](../../../pkg/models/shared/sessionpolicyenrollmentrequired.md) | :heavy_minus_sign:                                                                                       | N/A                                                                                                      |
| `StepUpRequired`                                                                                         | [*shared.SessionPolicyStepUpRequired](../../../pkg/models/shared/sessionpolicystepuprequired.md)         | :heavy_minus_sign:                                                                                       | N/A                                                                                                      |