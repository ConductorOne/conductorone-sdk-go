# OrgDomain

OrgDomain represents a verified email domain associated with the tenant.


## Fields

| Field                                           | Type                                            | Required                                        | Description                                     |
| ----------------------------------------------- | ----------------------------------------------- | ----------------------------------------------- | ----------------------------------------------- |
| `CreatedAt`                                     | [*time.Time](https://pkg.go.dev/time#Time)      | :heavy_minus_sign:                              | N/A                                             |
| `DeletedAt`                                     | [*time.Time](https://pkg.go.dev/time#Time)      | :heavy_minus_sign:                              | N/A                                             |
| `Domain`                                        | `*string`                                       | :heavy_minus_sign:                              | The verified domain name (e.g., "example.com"). |
| `ID`                                            | `*string`                                       | :heavy_minus_sign:                              | The unique identifier of the domain record.     |
| `UpdatedAt`                                     | [*time.Time](https://pkg.go.dev/time#Time)      | :heavy_minus_sign:                              | N/A                                             |