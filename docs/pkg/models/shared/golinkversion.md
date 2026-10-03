# GoLinkVersion

GoLinkVersion represents a point-in-time snapshot of a GoLink.


## Fields

| Field                                                  | Type                                                   | Required                                               | Description                                            |
| ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ |
| `ChangedByUserID`                                      | `*string`                                              | :heavy_minus_sign:                                     | The changedByUserId field.                             |
| `CreatedAt`                                            | [*time.Time](https://pkg.go.dev/time#Time)             | :heavy_minus_sign:                                     | N/A                                                    |
| `GolinkID`                                             | `*string`                                              | :heavy_minus_sign:                                     | The golinkId field.                                    |
| `ID`                                                   | `*string`                                              | :heavy_minus_sign:                                     | The id field.                                          |
| `Snapshot`                                             | [*shared.GoLink](../../../pkg/models/shared/golink.md) | :heavy_minus_sign:                                     | N/A                                                    |
| `Version`                                              | `*int`                                                 | :heavy_minus_sign:                                     | The version field.                                     |