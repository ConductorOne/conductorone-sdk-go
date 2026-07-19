# CohortUserWithCoverage

CohortUserWithCoverage pairs a user with the count of selected entitlements they hold.


## Fields

| Field                                                           | Type                                                            | Required                                                        | Description                                                     |
| --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- |
| `CoveredCount`                                                  | `*int`                                                          | :heavy_minus_sign:                                              | Number of selected_entitlements that this user currently holds. |
| `User`                                                          | [*shared.User](../../../pkg/models/shared/user.md)              | :heavy_minus_sign:                                              | N/A                                                             |