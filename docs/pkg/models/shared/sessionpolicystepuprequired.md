# SessionPolicyStepUpRequired

StepUpRequired demands a stronger re-authentication before the session may
 continue.


## Fields

| Field                                                                                                       | Type                                                                                                        | Required                                                                                                    | Description                                                                                                 |
| ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- |
| `Level`                                                                                                     | [*shared.Level](../../../pkg/models/shared/level.md)                                                        | :heavy_minus_sign:                                                                                          | The level field.                                                                                            |
| `MaxAgeSeconds`                                                                                             | `*int`                                                                                                      | :heavy_minus_sign:                                                                                          | How fresh the step-up must be, in seconds.                                                                  |
| `Types`                                                                                                     | [][shared.SessionPolicyStepUpRequiredTypes](../../../pkg/models/shared/sessionpolicystepuprequiredtypes.md) | :heavy_minus_sign:                                                                                          | The types field.                                                                                            |