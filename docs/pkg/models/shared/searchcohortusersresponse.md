# SearchCohortUsersResponse

The SearchCohortUsersResponse message.


## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `List`                                                                | [][shared.User](../../../pkg/models/shared/user.md)                   | :heavy_minus_sign:                                                    | The list of users matching the cohort and optional filters.           |
| `NextPageToken`                                                       | `*string`                                                             | :heavy_minus_sign:                                                    | Token to retrieve the next page of results, empty if no more results. |