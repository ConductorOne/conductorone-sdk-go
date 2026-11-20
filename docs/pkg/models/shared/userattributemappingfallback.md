# UserAttributeMappingFallback

The UserAttributeMappingFallback message.

This message contains a oneof named fallback. Only a single field of the following list may be set at a time:
  - app
  - appUserProfile
  - celExpression



## Fields

| Field                                                                                                                | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `App`                                                                                                                | [*shared.UserAttributeTypeConfigApp](../../../pkg/models/shared/userattributetypeconfigapp.md)                       | :heavy_minus_sign:                                                                                                   | N/A                                                                                                                  |
| `AppUserProfile`                                                                                                     | [*shared.UserAttributeTypeConfigAppUserProfile](../../../pkg/models/shared/userattributetypeconfigappuserprofile.md) | :heavy_minus_sign:                                                                                                   | N/A                                                                                                                  |
| `CelExpression`                                                                                                      | [*shared.UserAttributeTypeConfigCelExpression](../../../pkg/models/shared/userattributetypeconfigcelexpression.md)   | :heavy_minus_sign:                                                                                                   | N/A                                                                                                                  |