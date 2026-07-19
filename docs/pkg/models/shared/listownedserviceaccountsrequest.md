# ListOwnedServiceAccountsRequest

The ListOwnedServiceAccountsRequest message.


## Fields

| Field                                                                        | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ExpandMask`                                                                 | [*shared.AppUserExpandMask](../../../pkg/models/shared/appuserexpandmask.md) | :heavy_minus_sign:                                                           | N/A                                                                          |
| `PageSize`                                                                   | `*int`                                                                       | :heavy_minus_sign:                                                           | The maximum number of results to return per page.                            |
| `PageToken`                                                                  | `*string`                                                                    | :heavy_minus_sign:                                                           | The token for fetching the next page of results.                             |