# Contacts

Contacts represents the contact configuration for an organization.


## Fields

| Field                                                         | Type                                                          | Required                                                      | Description                                                   |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| `BillingEmails`                                               | []`string`                                                    | :heavy_minus_sign:                                            | Email addresses of billing contacts for this organization.    |
| `CreatedAt`                                                   | [*time.Time](https://pkg.go.dev/time#Time)                    | :heavy_minus_sign:                                            | N/A                                                           |
| `OperationsEmails`                                            | []`string`                                                    | :heavy_minus_sign:                                            | Email addresses of operations contacts for this organization. |
| `SecurityEmails`                                              | []`string`                                                    | :heavy_minus_sign:                                            | Email addresses of security contacts for this organization.   |
| `UpdatedAt`                                                   | [*time.Time](https://pkg.go.dev/time#Time)                    | :heavy_minus_sign:                                            | N/A                                                           |