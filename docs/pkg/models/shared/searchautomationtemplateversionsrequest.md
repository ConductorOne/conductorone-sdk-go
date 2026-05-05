# SearchAutomationTemplateVersionsRequest

The SearchAutomationTemplateVersionsRequest message.


## Fields

| Field                                                                      | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `AutomationTemplateID`                                                     | `*string`                                                                  | :heavy_minus_sign:                                                         | The automation template whose version history to search.                   |
| `PageSize`                                                                 | `*int`                                                                     | :heavy_minus_sign:                                                         | Maximum number of results to return per page.                              |
| `PageToken`                                                                | `*string`                                                                  | :heavy_minus_sign:                                                         | Pagination token from a previous SearchAutomationTemplateVersionsResponse. |