# AccessReviewExcludedResourceType

A resource type whose grants are excluded from a user-first access review campaign's scope.


## Fields

| Field                                                           | Type                                                            | Required                                                        | Description                                                     |
| --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- |
| `AccessReviewID`                                                | `*string`                                                       | :heavy_minus_sign:                                              | The ID of the access review campaign this exclusion belongs to. |
| `AppID`                                                         | `*string`                                                       | :heavy_minus_sign:                                              | The ID of the application that owns the resource type.          |
| `CreatedAt`                                                     | [*time.Time](https://pkg.go.dev/time#Time)                      | :heavy_minus_sign:                                              | N/A                                                             |
| `DeletedAt`                                                     | [*time.Time](https://pkg.go.dev/time#Time)                      | :heavy_minus_sign:                                              | N/A                                                             |
| `ResourceTypeID`                                                | `*string`                                                       | :heavy_minus_sign:                                              | The ID of the resource type excluded from scope.                |
| `TenantID`                                                      | `*string`                                                       | :heavy_minus_sign:                                              | The tenant that owns this exclusion.                            |
| `UpdatedAt`                                                     | [*time.Time](https://pkg.go.dev/time#Time)                      | :heavy_minus_sign:                                              | N/A                                                             |