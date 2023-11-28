# SearchUsersResponse

The SearchUsersResponse message.


## Fields

| Field                                                                                                                                                                                                                                                                                                                                          | Type                                                                                                                                                                                                                                                                                                                                           | Required                                                                                                                                                                                                                                                                                                                                       | Description                                                                                                                                                                                                                                                                                                                                    |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `Expanded`                                                                                                                                                                                                                                                                                                                                     | [][shared.SearchUsersResponseExpanded](../../../pkg/models/shared/searchusersresponseexpanded.md)                                                                                                                                                                                                                                              | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                             | List of related objects                                                                                                                                                                                                                                                                                                                        |
| `List`                                                                                                                                                                                                                                                                                                                                         | [][shared.UserView](../../../pkg/models/shared/userview.md)                                                                                                                                                                                                                                                                                    | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                             | The list of results containing up to X results, where X is the page size defined in the request                                                                                                                                                                                                                                                |
| `NextPageToken`                                                                                                                                                                                                                                                                                                                                | **string*                                                                                                                                                                                                                                                                                                                                      | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                             | The nextPageToken is shown for the next page if the number of results is larger than the max page size. The server returns one page of results and the nextPageToken until all results are retreived. To retrieve the next page, use the same request and append a pageToken field with the value of nextPageToken shown on the previous page. |