# FindingTypeSetting

FindingTypeSetting is one finding type's detection switch as it currently
 stands. Named for a single type on purpose: the stored model
 c1.models.finding.v1.FindingSettings is the tenant-wide object holding every
 type, and one name for both granularities reads as the same thing twice.
 Display copy for the type is client-owned; this carries state only.


## Fields

| Field                                                                                                                                              | Type                                                                                                                                               | Required                                                                                                                                           | Description                                                                                                                                        |
| -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- |
| `Enabled`                                                                                                                                          | `*bool`                                                                                                                                            | :heavy_minus_sign:                                                                                                                                 | Whether the system detects this finding type. Types never configured read<br/> back their shipped default, which is per type rather than uniformly on. |
| `FindingType`                                                                                                                                      | [*shared.FindingTypeSettingFindingType](../../../pkg/models/shared/findingtypesettingfindingtype.md)                                               | :heavy_minus_sign:                                                                                                                                 | The findingType field.                                                                                                                             |