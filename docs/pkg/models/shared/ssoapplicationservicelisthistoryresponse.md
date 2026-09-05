# SSOApplicationServiceListHistoryResponse

SSOApplicationServiceListHistoryResponse returns SSO application history
 entries.


## Fields

| Field                                                                                           | Type                                                                                            | Required                                                                                        | Description                                                                                     |
| ----------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------- |
| `List`                                                                                          | [][shared.SSOApplicationHistoryEntry](../../../pkg/models/shared/ssoapplicationhistoryentry.md) | :heavy_minus_sign:                                                                              | The page of history entries, newest first.                                                      |
| `NextPageToken`                                                                                 | `*string`                                                                                       | :heavy_minus_sign:                                                                              | Pagination token for the next page, or empty if there are no more results.                      |