# ListSuggestionsResponse

The ListSuggestionsResponse message.


## Fields

| Field                                                                                                   | Type                                                                                                    | Required                                                                                                | Description                                                                                             |
| ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- |
| `List`                                                                                                  | [][shared.RoleMiningManagementSuggestion](../../../pkg/models/shared/roleminingmanagementsuggestion.md) | :heavy_minus_sign:                                                                                      | The list of role mining suggestions.                                                                    |
| `NextPageToken`                                                                                         | `*string`                                                                                               | :heavy_minus_sign:                                                                                      | Token to retrieve the next page of results, empty if no more results.                                   |