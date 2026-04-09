# ServiceAccountMisclassificationType

ServiceAccountMisclassificationType: account classified as human but
 detected as service account, or vice versa.
 Target: AppUserTarget (the misclassified account).


## Fields

| Field                                                                            | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `CurrentAccountType`                                                             | [*shared.CurrentAccountType](../../../pkg/models/shared/currentaccounttype.md)   | :heavy_minus_sign:                                                               | What the account is currently classified as in ConductorOne.                     |
| `DetectedAccountType`                                                            | [*shared.DetectedAccountType](../../../pkg/models/shared/detectedaccounttype.md) | :heavy_minus_sign:                                                               | What the detector thinks the account actually is.                                |