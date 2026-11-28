# AppUserSourceReason

Why a record is on its account.

This message contains a oneof named reason. Only a single field of the following list may be set at a time:
  - created
  - ruleMatch
  - merged
  - legacy
  - resourceIdMatch



## Fields

| Field                                                                                                          | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `Created`                                                                                                      | [*shared.AppUserSourceReasonCreated](../../../pkg/models/shared/appusersourcereasoncreated.md)                 | :heavy_minus_sign:                                                                                             | N/A                                                                                                            |
| `Legacy`                                                                                                       | [*shared.AppUserSourceReasonLegacy](../../../pkg/models/shared/appusersourcereasonlegacy.md)                   | :heavy_minus_sign:                                                                                             | N/A                                                                                                            |
| `Merged`                                                                                                       | [*shared.AppUserSourceReasonMerged](../../../pkg/models/shared/appusersourcereasonmerged.md)                   | :heavy_minus_sign:                                                                                             | N/A                                                                                                            |
| `PlacedAt`                                                                                                     | [*time.Time](https://pkg.go.dev/time#Time)                                                                     | :heavy_minus_sign:                                                                                             | N/A                                                                                                            |
| `ResourceIDMatch`                                                                                              | [*shared.AppUserSourceReasonResourceIDMatch](../../../pkg/models/shared/appusersourcereasonresourceidmatch.md) | :heavy_minus_sign:                                                                                             | N/A                                                                                                            |
| `RuleMatch`                                                                                                    | [*shared.AppUserSourceReasonRuleMatch](../../../pkg/models/shared/appusersourcereasonrulematch.md)             | :heavy_minus_sign:                                                                                             | N/A                                                                                                            |