# GetAttributionRollupsResponse

GetAttributionRollupsResponse contains one page of ranked attribution groups.


## Fields

| Field                                                                       | Type                                                                        | Required                                                                    | Description                                                                 |
| --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- |
| `Groups`                                                                    | [][shared.AttributionGroup](../../../pkg/models/shared/attributiongroup.md) | :heavy_minus_sign:                                                          | Groups ordered by settled spend descending, then key ascending.             |
| `NextPageToken`                                                             | `*string`                                                                   | :heavy_minus_sign:                                                          | Token for the next page. Empty when no more groups remain.                  |