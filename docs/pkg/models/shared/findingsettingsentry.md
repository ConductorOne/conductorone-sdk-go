# FindingSettingsEntry

FindingSettingsEntry is a requested change to one type, which is why it is a
 separate message from FindingTypeSetting rather than the same one reused: an
 update needs enum validation and presence on `enabled` so an omitted field is
 an error, while a response always carries a value and must not make callers
 handle an absent one.


## Fields

| Field                                                                                                                          | Type                                                                                                                           | Required                                                                                                                       | Description                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ |
| `Enabled`                                                                                                                      | `*bool`                                                                                                                        | :heavy_minus_sign:                                                                                                             | Target state. Required: explicit presence keeps an omitted field from<br/> reading as false and silently switching a detector off. |
| `FindingType`                                                                                                                  | [*shared.FindingSettingsEntryFindingType](../../../pkg/models/shared/findingsettingsentryfindingtype.md)                       | :heavy_minus_sign:                                                                                                             | The finding type to configure. Must be a detector-backed type.                                                                 |