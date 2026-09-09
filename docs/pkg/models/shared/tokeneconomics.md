# TokenEconomics

TokenEconomics summarizes token usage and cost for settled calls.


## Fields

| Field                                                             | Type                                                              | Required                                                          | Description                                                       |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `BillableTokens`                                                  | `*int64`                                                          | :heavy_minus_sign:                                                | Total billable input, output, cache-read, and cache-write tokens. |
| `BlendedCostPerMillionTokensNano`                                 | `*int64`                                                          | :heavy_minus_sign:                                                | Blended settled cost per million billable tokens.                 |
| `CacheReadTokens`                                                 | `*int64`                                                          | :heavy_minus_sign:                                                | Input tokens served from cache.                                   |
| `CacheReadValueNano`                                              | `*int64`                                                          | :heavy_minus_sign:                                                | Savings from cache reads compared with uncached input prices.     |
| `CacheWriteCostNano`                                              | `*int64`                                                          | :heavy_minus_sign:                                                | Cost of writing tokens to cache.                                  |
| `CacheWriteTokens`                                                | `*int64`                                                          | :heavy_minus_sign:                                                | Tokens written to cache.                                          |
| `InputTokens`                                                     | `*int64`                                                          | :heavy_minus_sign:                                                | Input tokens processed.                                           |
| `OutputTokens`                                                    | `*int64`                                                          | :heavy_minus_sign:                                                | Output tokens generated.                                          |
| `ReasoningTokens`                                                 | `*int64`                                                          | :heavy_minus_sign:                                                | Reasoning tokens generated.                                       |
| `SettledCalls`                                                    | `*int64`                                                          | :heavy_minus_sign:                                                | Number of calls with settled spend.                               |