# FindingRiskScore

The FindingRiskScore message.


## Fields

| Field                                                                         | Type                                                                          | Required                                                                      | Description                                                                   |
| ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
| `OriginalScore`                                                               | `*int64`                                                                      | :heavy_minus_sign:                                                            | The originalScore field.                                                      |
| `OverrideByUserID`                                                            | `*string`                                                                     | :heavy_minus_sign:                                                            | The overrideByUserId field.                                                   |
| `OverrideScore`                                                               | `*int64`                                                                      | :heavy_minus_sign:                                                            | The overrideScore field.                                                      |
| `RiskFactors`                                                                 | [][shared.FindingRiskFactor](../../../pkg/models/shared/findingriskfactor.md) | :heavy_minus_sign:                                                            | The riskFactors field.                                                        |
| `Score`                                                                       | `*int64`                                                                      | :heavy_minus_sign:                                                            | The score field.                                                              |
| `SystemScore`                                                                 | `*int64`                                                                      | :heavy_minus_sign:                                                            | The systemScore field.                                                        |