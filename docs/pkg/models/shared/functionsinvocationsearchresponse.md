# FunctionsInvocationSearchResponse

FunctionsInvocationSearchResponse is the response for searching function invocations.


## Fields

| Field                                                                           | Type                                                                            | Required                                                                        | Description                                                                     |
| ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- |
| `List`                                                                          | [][shared.FunctionInvocation](../../../pkg/models/shared/functioninvocation.md) | :heavy_minus_sign:                                                              | The list of function invocations, ordered by created_at descending.             |
| `NextPageToken`                                                                 | **string*                                                                       | :heavy_minus_sign:                                                              | The pagination token for fetching the next page.                                |