# CredentialInventoryPolicyServiceSearchResponse

The CredentialInventoryPolicyServiceSearchResponse message.


## Fields

| Field                                                                                         | Type                                                                                          | Required                                                                                      | Description                                                                                   |
| --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- |
| `List`                                                                                        | [][shared.CredentialInventoryPolicy](../../../pkg/models/shared/credentialinventorypolicy.md) | :heavy_minus_sign:                                                                            | The page of matching policies.                                                                |
| `NextPageToken`                                                                               | `*string`                                                                                     | :heavy_minus_sign:                                                                            | A token to fetch the next page, or empty if there are no more results.                        |