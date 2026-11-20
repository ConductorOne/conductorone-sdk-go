# GetSyncErrorDigestResponse

GetSyncErrorDigestResponse contains a fresh bounded summary of recent errors, using structured errors and heuristic activity matching.


## Fields

| Field                                                                               | Type                                                                                | Required                                                                            | Description                                                                         |
| ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- |
| `Entries`                                                                           | [][shared.SyncErrorDigestEntry](../../../pkg/models/shared/syncerrordigestentry.md) | :heavy_minus_sign:                                                                  | Most frequent distinct redacted messages, with deterministic tie ordering.          |