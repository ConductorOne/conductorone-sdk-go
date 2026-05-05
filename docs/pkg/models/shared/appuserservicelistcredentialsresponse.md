# AppUserServiceListCredentialsResponse

The response message for listing credentials of an app user.


## Fields

| Field                                                                         | Type                                                                          | Required                                                                      | Description                                                                   |
| ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
| `List`                                                                        | [][shared.AppUserCredential](../../../pkg/models/shared/appusercredential.md) | :heavy_minus_sign:                                                            | The list of credential results.                                               |
| `NextPageToken`                                                               | `*string`                                                                     | :heavy_minus_sign:                                                            | The token for fetching the next page of results.                              |