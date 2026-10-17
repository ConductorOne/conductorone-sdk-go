# EdgeInferenceModelFamily

EdgeInferenceModelFamily groups registered models, e.g. "Claude Opus".


## Fields

| Field                                                                                           | Type                                                                                            | Required                                                                                        | Description                                                                                     |
| ----------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------- |
| `Family`                                                                                        | `*string`                                                                                       | :heavy_minus_sign:                                                                              | Family name; models with no family are grouped under their own display name.                    |
| `Models`                                                                                        | [][shared.EdgeInferenceRegistryModel](../../../pkg/models/shared/edgeinferenceregistrymodel.md) | :heavy_minus_sign:                                                                              | The models field.                                                                               |