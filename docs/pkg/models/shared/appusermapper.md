# AppUserMapper

AppUserMapper configures custom account mapping for uplift.


## Fields

| Field                                                                                   | Type                                                                                    | Required                                                                                | Description                                                                             |
| --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- |
| `AppID`                                                                                 | `*string`                                                                               | :heavy_minus_sign:                                                                      | The app this mapper belongs to.                                                         |
| `MappingCases`                                                                          | [][shared.AppUserMapperMatchCase](../../../pkg/models/shared/appusermappermatchcase.md) | :heavy_minus_sign:                                                                      | Ordered list of match cases. Each case defines a pair of CEL key extractors.            |