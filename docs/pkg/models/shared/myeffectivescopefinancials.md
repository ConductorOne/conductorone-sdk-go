# MyEffectiveScopeFinancials

MyEffectiveScopeFinancials contains caller-owned budget amounts. It is
 present only for SUBJECT and SUBJECT_APP scopes that name the caller.


## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `AccountInitialized`                                               | `*bool`                                                            | :heavy_minus_sign:                                                 | Whether this budget account has recorded activity.                 |
| `HasRemainingNano`                                                 | `*bool`                                                            | :heavy_minus_sign:                                                 | Whether remaining_nano is available for this scope.                |
| `LifetimeConsumedNano`                                             | `*int64`                                                           | :heavy_minus_sign:                                                 | Spend settled over the lifetime of the budget account.             |
| `LimitNano`                                                        | `*int64`                                                           | :heavy_minus_sign:                                                 | Effective spend limit. Zero when tracking is true.                 |
| `RemainingNano`                                                    | `*int64`                                                           | :heavy_minus_sign:                                                 | Spend headroom after settled and reserved spend.                   |
| `ReservedNano`                                                     | `*int64`                                                           | :heavy_minus_sign:                                                 | Spend reserved by in-progress calls.                               |
| `Tracking`                                                         | `*bool`                                                            | :heavy_minus_sign:                                                 | Whether this scope records spend without enforcing a finite limit. |
| `WindowCommittedNano`                                              | `*int64`                                                           | :heavy_minus_sign:                                                 | Settled plus reserved spend in the current budget period.          |
| `WindowConsumedNano`                                               | `*int64`                                                           | :heavy_minus_sign:                                                 | Spend settled in the current budget period.                        |