# FunctionsInvocationSearchRequest

FunctionsInvocationSearchRequest is the request for searching function invocations.
 Results are returned in descending order by created_at (newest first).


## Fields

| Field                                            | Type                                             | Required                                         | Description                                      |
| ------------------------------------------------ | ------------------------------------------------ | ------------------------------------------------ | ------------------------------------------------ |
| `PageSize`                                       | **int*                                           | :heavy_minus_sign:                               | The number of results to return per page.        |
| `PageToken`                                      | **string*                                        | :heavy_minus_sign:                               | The pagination token for fetching the next page. |