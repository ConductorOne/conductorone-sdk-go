# XAASettingsServiceListHistoryResponse

XAASettingsServiceListHistoryResponse returns cross-app-access settings
 history entries.


## Fields

| Field                                                                                     | Type                                                                                      | Required                                                                                  | Description                                                                               |
| ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- |
| `List`                                                                                    | [][shared.XAASettingsHistoryEntry](../../../pkg/models/shared/xaasettingshistoryentry.md) | :heavy_minus_sign:                                                                        | The page of history entries, newest first.                                                |
| `NextPageToken`                                                                           | `*string`                                                                                 | :heavy_minus_sign:                                                                        | Pagination token for the next page, or empty if there are no more results.                |