# PickerField

The PickerField message.

This message contains a oneof named type. Only a single field of the following list may be set at a time:
  - appUserPicker
  - resourcePicker



## Fields

| Field                                                                        | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `AppResourceFilter`                                                          | [*shared.AppResourceFilter](../../../pkg/models/shared/appresourcefilter.md) | :heavy_minus_sign:                                                           | The AppResourceFilter message.                                               |
| `AppUserFilter`                                                              | [*shared.AppUserFilter](../../../pkg/models/shared/appuserfilter.md)         | :heavy_minus_sign:                                                           | The AppUserFilter message.                                                   |