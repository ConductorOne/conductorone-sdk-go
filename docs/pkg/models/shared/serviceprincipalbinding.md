# ServicePrincipalBinding

ServicePrincipalBinding is one row in the binding store, naming a
 subject's link to a single service principal.


## Fields

| Field                                      | Type                                       | Required                                   | Description                                |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| `CreatedAt`                                | [*time.Time](https://pkg.go.dev/time#Time) | :heavy_minus_sign:                         | N/A                                        |
| `ServicePrincipalID`                       | `*string`                                  | :heavy_minus_sign:                         | The servicePrincipalId field.              |
| `UpdatedAt`                                | [*time.Time](https://pkg.go.dev/time#Time) | :heavy_minus_sign:                         | N/A                                        |