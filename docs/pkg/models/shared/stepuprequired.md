# StepUpRequired

StepUpRequired demands a stronger re-authentication before access is granted.


## Fields

| Field                                                                             | Type                                                                              | Required                                                                          | Description                                                                       |
| --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- |
| `Level`                                                                           | [*shared.StepUpRequiredLevel](../../../pkg/models/shared/stepuprequiredlevel.md)  | :heavy_minus_sign:                                                                | The assurance level the step-up must reach.                                       |
| `MaxAgeSeconds`                                                                   | `*int`                                                                            | :heavy_minus_sign:                                                                | How fresh the step-up must be, in seconds.                                        |
| `Types`                                                                           | [][shared.StepUpRequiredTypes](../../../pkg/models/shared/stepuprequiredtypes.md) | :heavy_minus_sign:                                                                | The credential types that may satisfy the step-up.                                |