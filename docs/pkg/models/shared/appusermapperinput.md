# AppUserMapperInput

AppUserMapper configures custom account mapping for uplift.


## Fields

| Field                                                                                             | Type                                                                                              | Required                                                                                          | Description                                                                                       |
| ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- |
| `MappingCases`                                                                                    | [][shared.AppUserMapperMatchCaseInput](../../../pkg/models/shared/appusermappermatchcaseinput.md) | :heavy_minus_sign:                                                                                | Ordered list of match cases. Each case defines a pair of CEL key extractors.                      |