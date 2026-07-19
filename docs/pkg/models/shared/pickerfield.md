# PickerField

The PickerField message.

This message contains a oneof named type. Only a single field of the following list may be set at a time:
  - appUserPicker
  - resourcePicker
  - c1UserPicker



## Fields

| Field                                                                        | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `AppUserPicker`                                                              | [*shared.AppUserFilter](../../../pkg/models/shared/appuserfilter.md)         | :heavy_minus_sign:                                                           | N/A                                                                          |
| `C1UserPicker`                                                               | [*shared.C1UserFilter](../../../pkg/models/shared/c1userfilter.md)           | :heavy_minus_sign:                                                           | N/A                                                                          |
| `ResourcePicker`                                                             | [*shared.AppResourceFilter](../../../pkg/models/shared/appresourcefilter.md) | :heavy_minus_sign:                                                           | N/A                                                                          |