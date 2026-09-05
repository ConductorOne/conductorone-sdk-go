# CredentialExpiringEvidence

The CredentialExpiringEvidence message.


## Fields

| Field                                                   | Type                                                    | Required                                                | Description                                             |
| ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- |
| `Expired`                                               | `*bool`                                                 | :heavy_minus_sign:                                      | Whether the expiry was already past when last observed. |
| `ExpiresAt`                                             | [*time.Time](https://pkg.go.dev/time#Time)              | :heavy_minus_sign:                                      | N/A                                                     |