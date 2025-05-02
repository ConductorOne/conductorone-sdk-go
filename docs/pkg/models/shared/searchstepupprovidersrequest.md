# SearchStepUpProvidersRequest

Request message for searching step-up providers


## Fields

| Field                                                                         | Type                                                                          | Required                                                                      | Description                                                                   |
| ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
| `PageSize`                                                                    | **int*                                                                        | :heavy_minus_sign:                                                            | Maximum number of results to return                                           |
| `PageToken`                                                                   | **string*                                                                     | :heavy_minus_sign:                                                            | Token for pagination                                                          |
| `ProviderType`                                                                | [*shared.ProviderType](../../../pkg/models/shared/providertype.md)            | :heavy_minus_sign:                                                            | The providerType field.                                                       |
| `Query`                                                                       | **string*                                                                     | :heavy_minus_sign:                                                            | Filter by name (partial match)                                                |
| `Refs`                                                                        | [][shared.StepUpProviderRef](../../../pkg/models/shared/stepupproviderref.md) | :heavy_minus_sign:                                                            | The refs field.                                                               |