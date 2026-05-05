# ExportsSearchServiceSearchResponse

ExportsSearchServiceSearchResponse is the response for searching system log exports.


## Fields

| Field                                                                                  | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `List`                                                                                 | [][shared.Exporter](../../../pkg/models/shared/exporter.md)                            | :heavy_minus_sign:                                                                     | The list of system log exports matching the search criteria.                           |
| `NextPageToken`                                                                        | `*string`                                                                              | :heavy_minus_sign:                                                                     | The token to retrieve the next page of results, or empty if there are no more results. |