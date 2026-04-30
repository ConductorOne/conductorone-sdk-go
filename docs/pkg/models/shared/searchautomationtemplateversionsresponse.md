# SearchAutomationTemplateVersionsResponse

The SearchAutomationTemplateVersionsResponse message.


## Fields

| Field                                                                                         | Type                                                                                          | Required                                                                                      | Description                                                                                   |
| --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- |
| `List`                                                                                        | [][shared.AutomationTemplateVersion](../../../pkg/models/shared/automationtemplateversion.md) | :heavy_minus_sign:                                                                            | The page of template versions matching the search criteria.                                   |
| `NextPageToken`                                                                               | `*string`                                                                                     | :heavy_minus_sign:                                                                            | Token to retrieve the next page of results, empty when no more results exist.                 |