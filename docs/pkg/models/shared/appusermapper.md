# AppUserMapper

AppUserMapper configures custom account mapping for uplift.


## Fields

| Field                                                                                   | Type                                                                                    | Required                                                                                | Description                                                                             |
| --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- |
| `MappingCases`                                                                          | [][shared.AppUserMapperMatchCase](../../../pkg/models/shared/appusermappermatchcase.md) | :heavy_minus_sign:                                                                      | Ordered list of match cases. Each case defines a pair of CEL key extractors.            |