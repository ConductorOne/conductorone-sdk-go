# EdgeIntegrationServiceListResponse

EdgeIntegrationServiceListResponse is one page of integration metadata.


## Fields

| Field                                                                     | Type                                                                      | Required                                                                  | Description                                                               |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| `List`                                                                    | [][shared.EdgeIntegration](../../../pkg/models/shared/edgeintegration.md) | :heavy_minus_sign:                                                        | Integrations in the requested application.                                |
| `NextPageToken`                                                           | `*string`                                                                 | :heavy_minus_sign:                                                        | Token for the next page; empty when complete.                             |