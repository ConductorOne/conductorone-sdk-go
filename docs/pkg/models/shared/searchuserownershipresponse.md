# SearchUserOwnershipResponse

The SearchUserOwnershipResponse message contains a paginated list of ownership entries.


## Fields

| Field                                                                                | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `List`                                                                               | [][shared.UserOwnershipEntry](../../../pkg/models/shared/userownershipentry.md)      | :heavy_minus_sign:                                                                   | The list of ownership entries for the requested user.                                |
| `NextPageToken`                                                                      | `*string`                                                                            | :heavy_minus_sign:                                                                   | Pagination token for the next page of results. Empty when there are no more results. |