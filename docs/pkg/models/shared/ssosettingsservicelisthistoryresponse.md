# SSOSettingsServiceListHistoryResponse

SSOSettingsServiceListHistoryResponse returns SSO settings history entries.


## Fields

| Field                                                                                     | Type                                                                                      | Required                                                                                  | Description                                                                               |
| ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- |
| `List`                                                                                    | [][shared.SSOSettingsHistoryEntry](../../../pkg/models/shared/ssosettingshistoryentry.md) | :heavy_minus_sign:                                                                        | The page of history entries, newest first.                                                |
| `NextPageToken`                                                                           | `*string`                                                                                 | :heavy_minus_sign:                                                                        | Pagination token for the next page, or empty if there are no more results.                |