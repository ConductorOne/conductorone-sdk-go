# XAAScopeServiceListResponse

XAAScopeServiceListResponse returns a page of scopes.


## Fields

| Field                                                           | Type                                                            | Required                                                        | Description                                                     |
| --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- |
| `NextPageToken`                                                 | `*string`                                                       | :heavy_minus_sign:                                              | Token for the next page, or empty if there are no more results. |
| `Scopes`                                                        | [][shared.XAAScope](../../../pkg/models/shared/xaascope.md)     | :heavy_minus_sign:                                              | The page of scopes.                                             |