# SearchMyVendedCredentialsRequest

SearchMyVendedCredentialsRequest has the same filters as the general
 inventory search, but its distinct type carries the self-only permission.
 The server always scopes results to the signed-in requester's issuance
 tickets; callers cannot provide an identity-user filter.


## Fields

| Field                                                                            | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `AppID`                                                                          | `*string`                                                                        | :heavy_minus_sign:                                                               | The appId field.                                                                 |
| `ExpandMask`                                                                     | [*shared.AppSecretExpandMask](../../../pkg/models/shared/appsecretexpandmask.md) | :heavy_minus_sign:                                                               | N/A                                                                              |
| `PageSize`                                                                       | `*int`                                                                           | :heavy_minus_sign:                                                               | The pageSize field.                                                              |
| `PageToken`                                                                      | `*string`                                                                        | :heavy_minus_sign:                                                               | The pageToken field.                                                             |
| `Query`                                                                          | `*string`                                                                        | :heavy_minus_sign:                                                               | The query field.                                                                 |
| `Refs`                                                                           | [][shared.AppSecretRef](../../../pkg/models/shared/appsecretref.md)              | :heavy_minus_sign:                                                               | The refs field.                                                                  |
| `ResourceTypeIds`                                                                | []`string`                                                                       | :heavy_minus_sign:                                                               | The resourceTypeIds field.                                                       |