# TaskTypeFinding

The TaskTypeFinding message.


## Fields

| Field                                                                                  | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `FindingID`                                                                            | `*string`                                                                              | :heavy_minus_sign:                                                                     | Reference to the source finding.                                                       |
| `FindingType`                                                                          | `*string`                                                                              | :heavy_minus_sign:                                                                     | The finding type discriminator.                                                        |
| `Outcome`                                                                              | [*shared.TaskTypeFindingOutcome](../../../pkg/models/shared/tasktypefindingoutcome.md) | :heavy_minus_sign:                                                                     | The outcome field.                                                                     |
| `OutcomeTime`                                                                          | [*time.Time](https://pkg.go.dev/time#Time)                                             | :heavy_minus_sign:                                                                     | N/A                                                                                    |