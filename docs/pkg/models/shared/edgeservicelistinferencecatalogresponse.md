# EdgeServiceListInferenceCatalogResponse

EdgeServiceListInferenceCatalogResponse contains every catalog entry.


## Fields

| Field                                                                                            | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `Entries`                                                                                        | [][shared.EdgeInferenceCatalogEntry](../../../pkg/models/shared/edgeinferencecatalogentry.md)    | :heavy_minus_sign:                                                                               | The entries field.                                                                               |
| `Families`                                                                                       | [][shared.EdgeInferenceModelFamily](../../../pkg/models/shared/edgeinferencemodelfamily.md)      | :heavy_minus_sign:                                                                               | C1's registered models grouped by family, with the ways an Edge can serve each.                  |
| `Tiers`                                                                                          | [][shared.EdgeInferenceTierOption](../../../pkg/models/shared/edgeinferencetieroption.md)        | :heavy_minus_sign:                                                                               | The model each C1 tier currently uses for the caller's tenant. A tier with no model is left out. |