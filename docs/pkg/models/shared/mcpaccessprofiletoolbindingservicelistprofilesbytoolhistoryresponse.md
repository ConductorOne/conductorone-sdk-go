# MCPAccessProfileToolBindingServiceListProfilesByToolHistoryResponse

Contains a page of change-history entries for the toolsets one tool has belonged to,
 sorted newest first.


## Fields

| Field                                                                                                                     | Type                                                                                                                      | Required                                                                                                                  | Description                                                                                                               |
| ------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------- |
| `List`                                                                                                                    | [][shared.MCPAccessProfileToolBindingHistoryEntry](../../../pkg/models/shared/mcpaccessprofiletoolbindinghistoryentry.md) | :heavy_minus_sign:                                                                                                        | The page of history entries, newest first.                                                                                |
| `NextPageToken`                                                                                                           | `*string`                                                                                                                 | :heavy_minus_sign:                                                                                                        | Pagination token for the next page, or empty if there are no more results.                                                |