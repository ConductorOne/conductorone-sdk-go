# XAAClientAudienceMappingServiceListHistoryResponse

XAAClientAudienceMappingServiceListHistoryResponse returns client audience
 mapping history entries.


## Fields

| Field                                                                                                               | Type                                                                                                                | Required                                                                                                            | Description                                                                                                         |
| ------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------- |
| `List`                                                                                                              | [][shared.XAAClientAudienceMappingHistoryEntry](../../../pkg/models/shared/xaaclientaudiencemappinghistoryentry.md) | :heavy_minus_sign:                                                                                                  | The page of history entries, newest first.                                                                          |
| `NextPageToken`                                                                                                     | `*string`                                                                                                           | :heavy_minus_sign:                                                                                                  | Pagination token for the next page, or empty if there are no more results.                                          |