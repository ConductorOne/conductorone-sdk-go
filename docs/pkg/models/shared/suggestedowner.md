# SuggestedOwner

SuggestedOwner is a candidate owner for an unowned NHI.


## Fields

| Field                                                                               | Type                                                                                | Required                                                                            | Description                                                                         |
| ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- |
| `ConfidencePercent`                                                                 | `*int64`                                                                            | :heavy_minus_sign:                                                                  | Max of reasons[].percent at detection time.                                         |
| `Reasons`                                                                           | [][shared.SuggestedOwnerReason](../../../pkg/models/shared/suggestedownerreason.md) | :heavy_minus_sign:                                                                  | The reasons field.                                                                  |
| `UserID`                                                                            | `*string`                                                                           | :heavy_minus_sign:                                                                  | C1 User ID; readers resolve the name.                                               |