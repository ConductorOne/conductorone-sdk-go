# FindingRiskFactor

The FindingRiskFactor message.


## Fields

| Field                                                                                        | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `Description`                                                                                | `*string`                                                                                    | :heavy_minus_sign:                                                                           | The description field.                                                                       |
| `Name`                                                                                       | `*string`                                                                                    | :heavy_minus_sign:                                                                           | The name field.                                                                              |
| `Severity`                                                                                   | [*shared.FindingRiskFactorSeverity](../../../pkg/models/shared/findingriskfactorseverity.md) | :heavy_minus_sign:                                                                           | The severity field.                                                                          |
| `Weight`                                                                                     | `*int64`                                                                                     | :heavy_minus_sign:                                                                           | Weight of this factor in the score calculation [0, 100].                                     |