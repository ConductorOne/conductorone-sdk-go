# CampaignHealthSnapshot

Campaign health snapshot. Read-only; updated by backend maintenance processors.


## Fields

| Field                                                            | Type                                                             | Required                                                         | Description                                                      |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `CheckedAt`                                                      | [*time.Time](https://pkg.go.dev/time#Time)                       | :heavy_minus_sign:                                               | N/A                                                              |
| `PhantomLockedCount`                                             | `*int`                                                           | :heavy_minus_sign:                                               | Number of pending actions locked by terminal (dead) submissions. |