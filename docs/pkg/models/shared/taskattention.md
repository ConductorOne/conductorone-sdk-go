# TaskAttention

Why an open task cannot progress without intervention.


## Fields

| Field                                                                          | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `AttentionAt`                                                                  | [*time.Time](https://pkg.go.dev/time#Time)                                     | :heavy_minus_sign:                                                             | N/A                                                                            |
| `Cause`                                                                        | [*shared.TaskAttentionCause](../../../pkg/models/shared/taskattentioncause.md) | :heavy_minus_sign:                                                             | Why the task cannot progress.                                                  |