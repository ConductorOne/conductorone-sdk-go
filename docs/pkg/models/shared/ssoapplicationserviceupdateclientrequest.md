# SSOApplicationServiceUpdateClientRequest

SSOApplicationServiceUpdateClientRequest replaces display name, redirect
 URIs, private-key JWKS, or tightens legacy PKCE to required.


## Fields

| Field                                                                                                  | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `Client`                                                                                               | [*shared.SSOApplicationOIDCClientConfig](../../../pkg/models/shared/ssoapplicationoidcclientconfig.md) | :heavy_check_mark:                                                                                     | N/A                                                                                                    |
| `ClientID`                                                                                             | `string`                                                                                               | :heavy_check_mark:                                                                                     | Generated client ID to update.                                                                         |