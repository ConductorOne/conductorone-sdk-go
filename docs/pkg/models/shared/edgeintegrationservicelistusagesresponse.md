# EdgeIntegrationServiceListUsagesResponse

EdgeIntegrationServiceListUsagesResponse contains matching Edges in a candidate page.


## Fields

| Field                                                                               | Type                                                                                | Required                                                                            | Description                                                                         |
| ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- |
| `List`                                                                              | [][shared.EdgeIntegrationUsage](../../../pkg/models/shared/edgeintegrationusage.md) | :heavy_minus_sign:                                                                  | Saved references found in this page. May be empty even when another page exists.    |
| `NextPageToken`                                                                     | `*string`                                                                           | :heavy_minus_sign:                                                                  | Continue scanning with this token, even if list is empty.                           |