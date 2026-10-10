# EdgeFindingSettingsEntry

EdgeFindingSettingsEntry is a requested change to one category. Separate
 from EdgeFindingCategorySetting for the same reason FindingSettingsEntry is:
 an update needs enum validation and presence on `enabled`.


## Fields

| Field                                                                                                      | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `Category`                                                                                                 | [*shared.EdgeFindingSettingsEntryCategory](../../../pkg/models/shared/edgefindingsettingsentrycategory.md) | :heavy_minus_sign:                                                                                         | The category field.                                                                                        |
| `Enabled`                                                                                                  | `*bool`                                                                                                    | :heavy_minus_sign:                                                                                         | Target state. Required: an omitted field must not read as false.                                           |