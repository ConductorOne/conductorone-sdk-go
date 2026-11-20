# SyncCountDelta

SyncCountDelta compares one aggregate count; missing map entries count as zero.


## Fields

| Field                                                               | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `After`                                                             | `*int64`                                                            | :heavy_minus_sign:                                                  | Count in the after run.                                             |
| `Before`                                                            | `*int64`                                                            | :heavy_minus_sign:                                                  | Count in the before run.                                            |
| `Delta`                                                             | `*int64`                                                            | :heavy_minus_sign:                                                  | After count minus before count.                                     |
| `LargeDecrease`                                                     | `*bool`                                                             | :heavy_minus_sign:                                                  | True when a positive before count decreases by at least 30 percent. |