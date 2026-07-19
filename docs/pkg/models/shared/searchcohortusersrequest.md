# SearchCohortUsersRequest

The SearchCohortUsersRequest message.


## Fields

| Field                                                                   | Type                                                                    | Required                                                                | Description                                                             |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `PageSize`                                                              | `*int`                                                                  | :heavy_minus_sign:                                                      | Maximum number of users to return per page.                             |
| `PageToken`                                                             | `*string`                                                               | :heavy_minus_sign:                                                      | Pagination token from a previous response.                              |
| `ProfileFilters`                                                        | [][shared.ProfileFilter](../../../pkg/models/shared/profilefilter.md)   | :heavy_minus_sign:                                                      | Additional profile filters to narrow the cohort user search.            |
| `SelectedEntitlements`                                                  | [][shared.EntitlementRef](../../../pkg/models/shared/entitlementref.md) | :heavy_minus_sign:                                                      | Optional list of entitlements to compute per-user coverage for.         |