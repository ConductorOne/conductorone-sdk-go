# APIKeySearchServiceSearchResponse

The APIKeySearchServiceSearchResponse message.


## Fields

| Field                                                                                | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `List`                                                                               | [][shared.APIKey](../../../pkg/models/shared/apikey.md)                              | :heavy_minus_sign:                                                                   | The list of API keys matching the search criteria.                                   |
| `NextPageToken`                                                                      | `*string`                                                                            | :heavy_minus_sign:                                                                   | A token to retrieve the next page of results, or empty if there are no more results. |