# PersonalDeviceServiceSearchResponse

The PersonalDeviceServiceSearchResponse message.


## Fields

| Field                                                                        | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `List`                                                                       | [][shared.PersonalDevice](../../../pkg/models/shared/personaldevice.md)      | :heavy_minus_sign:                                                           | The devices the calling user has registered, matching the search criteria.   |
| `NextPageToken`                                                              | `*string`                                                                    | :heavy_minus_sign:                                                           | A token to retrieve the next page of results, or empty if there are no more. |