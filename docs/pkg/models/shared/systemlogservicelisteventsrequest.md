# SystemLogServiceListEventsRequest

The SystemLogServiceListEventsRequest message.


## Fields

| Field                                                                | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `PageSize`                                                           | **int*                                                               | :heavy_minus_sign:                                                   | The pageSize field.                                                  |
| `PageToken`                                                          | **string*                                                            | :heavy_minus_sign:                                                   | The pageToken field.                                                 |
| `Since`                                                              | [*time.Time](https://pkg.go.dev/time#Time)                           | :heavy_minus_sign:                                                   | N/A                                                                  |
| `SinceEventUID`                                                      | **string*                                                            | :heavy_minus_sign:                                                   | The sinceEventUid field.                                             |
| `SortDirection`                                                      | [*shared.SortDirection](../../../pkg/models/shared/sortdirection.md) | :heavy_minus_sign:                                                   | The sortDirection field.                                             |
| `Until`                                                              | [*time.Time](https://pkg.go.dev/time#Time)                           | :heavy_minus_sign:                                                   | N/A                                                                  |