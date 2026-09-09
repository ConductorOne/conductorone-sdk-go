# SSOApplicationServiceCreateClientResponse

SSOApplicationServiceCreateClientResponse contains the generated client and
 its one-time secret, when applicable.


## Fields

| Field                                                                                      | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `Client`                                                                                   | [*shared.SSOApplicationOIDCClient](../../../pkg/models/shared/ssoapplicationoidcclient.md) | :heavy_minus_sign:                                                                         | N/A                                                                                        |
| `ClientSecret`                                                                             | `*string`                                                                                  | :heavy_minus_sign:                                                                         | Returned once for client_secret_basic/client_secret_post; empty for<br/> none/private_key_jwt. |