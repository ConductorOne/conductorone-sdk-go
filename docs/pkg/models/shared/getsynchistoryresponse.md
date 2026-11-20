# GetSyncHistoryResponse

GetSyncHistoryResponse contains newest-first runs and the trend within that window.


## Fields

| Field                                                                        | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `Runs`                                                                       | [][shared.SyncRunSummary](../../../pkg/models/shared/syncrunsummary.md)      | :heavy_minus_sign:                                                           | Recent completed and in-progress runs, deduplicated by lifecycle identifier. |
| `Streak`                                                                     | [*shared.SyncFailureStreak](../../../pkg/models/shared/syncfailurestreak.md) | :heavy_minus_sign:                                                           | N/A                                                                          |