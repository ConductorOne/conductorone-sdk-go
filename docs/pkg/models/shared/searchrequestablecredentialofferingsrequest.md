# SearchRequestableCredentialOfferingsRequest

The SearchRequestableCredentialOfferingsRequest message.


## Fields

| Field                                           | Type                                            | Required                                        | Description                                     |
| ----------------------------------------------- | ----------------------------------------------- | ----------------------------------------------- | ----------------------------------------------- |
| `AppID`                                         | `*string`                                       | :heavy_minus_sign:                              | Limits the search to one app.                   |
| `PageSize`                                      | `*int`                                          | :heavy_minus_sign:                              | The pageSize field.                             |
| `PageToken`                                     | `*string`                                       | :heavy_minus_sign:                              | The pageToken field.                            |
| `Query`                                         | `*string`                                       | :heavy_minus_sign:                              | The query field.                                |
| `TargetUserID`                                  | `*string`                                       | :heavy_minus_sign:                              | The user to search for. Defaults to the caller. |