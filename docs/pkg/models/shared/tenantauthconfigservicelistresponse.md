# TenantAuthConfigServiceListResponse

The TenantAuthConfigServiceListResponse message.


## Fields

| Field                                                                                | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `List`                                                                               | [][shared.TenantAuthConfig](../../../pkg/models/shared/tenantauthconfig.md)          | :heavy_minus_sign:                                                                   | The list of authentication provider configurations.                                  |
| `NextPageToken`                                                                      | `*string`                                                                            | :heavy_minus_sign:                                                                   | A token to retrieve the next page of results, or empty if there are no more results. |