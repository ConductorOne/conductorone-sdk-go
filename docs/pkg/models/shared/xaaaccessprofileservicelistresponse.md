# XAAAccessProfileServiceListResponse

XAAAccessProfileServiceListResponse returns a page of access profiles.


## Fields

| Field                                                                       | Type                                                                        | Required                                                                    | Description                                                                 |
| --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- |
| `AccessProfiles`                                                            | [][shared.XAAAccessProfile](../../../pkg/models/shared/xaaaccessprofile.md) | :heavy_minus_sign:                                                          | The page of access profiles.                                                |
| `NextPageToken`                                                             | `*string`                                                                   | :heavy_minus_sign:                                                          | Token for the next page, or empty if there are no more results.             |