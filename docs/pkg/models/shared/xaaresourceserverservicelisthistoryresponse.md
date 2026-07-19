# XAAResourceServerServiceListHistoryResponse

XAAResourceServerServiceListHistoryResponse returns resource server history
 entries.


## Fields

| Field                                                                                                 | Type                                                                                                  | Required                                                                                              | Description                                                                                           |
| ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- |
| `List`                                                                                                | [][shared.XAAResourceServerHistoryEntry](../../../pkg/models/shared/xaaresourceserverhistoryentry.md) | :heavy_minus_sign:                                                                                    | The page of history entries, newest first.                                                            |
| `NextPageToken`                                                                                       | `*string`                                                                                             | :heavy_minus_sign:                                                                                    | Pagination token for the next page, or empty if there are no more results.                            |