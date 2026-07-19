# ListAIGovernanceSettingsHistoryResponse

ListAIGovernanceSettingsHistoryResponse contains a page of AI governance
 settings change-history entries, newest first.


## Fields

| Field                                                                                                       | Type                                                                                                        | Required                                                                                                    | Description                                                                                                 |
| ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- |
| `List`                                                                                                      | [][shared.AIGovernanceSettingsHistoryEntry](../../../pkg/models/shared/aigovernancesettingshistoryentry.md) | :heavy_minus_sign:                                                                                          | The page of history entries, newest first.                                                                  |
| `NextPageToken`                                                                                             | `*string`                                                                                                   | :heavy_minus_sign:                                                                                          | Pagination token for the next page, or empty if there are no more results.                                  |