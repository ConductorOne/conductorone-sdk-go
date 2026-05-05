# PersonalClientSearchServiceSearchResponse

The PersonalClientSearchServiceSearchResponse message.


## Fields

| Field                                                                                | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `List`                                                                               | [][shared.PersonalClient](../../../pkg/models/shared/personalclient.md)              | :heavy_minus_sign:                                                                   | The list of personal client credentials matching the search criteria.                |
| `NextPageToken`                                                                      | `*string`                                                                            | :heavy_minus_sign:                                                                   | A token to retrieve the next page of results, or empty if there are no more results. |