# SendGridProviderConfig

SendGridProviderConfig configures sending via a customer's SendGrid account.


## Fields

| Field                                                                                                                                  | Type                                                                                                                                   | Required                                                                                                                               | Description                                                                                                                            |
| -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| `APIKey`                                                                                                                               | `*string`                                                                                                                              | :heavy_minus_sign:                                                                                                                     | Customer's SendGrid API key. Write-only: accepted on create/update, never returned in Get.<br/> Empty on update means "keep existing key". |