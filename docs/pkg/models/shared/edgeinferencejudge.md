# EdgeInferenceJudge

The EdgeInferenceJudge message.


## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `BaseThreshold`                                                    | `*float64`                                                         | :heavy_minus_sign:                                                 | The baseThreshold field.                                           |
| `MaxOutputTokens`                                                  | `*int64`                                                           | :heavy_minus_sign:                                                 | Maximum tokens available to classify a stage result (1 to 16,384). |
| `Prompt`                                                           | `*string`                                                          | :heavy_minus_sign:                                                 | The prompt field.                                                  |
| `RecentTurnWindow`                                                 | `*int64`                                                           | :heavy_minus_sign:                                                 | The recentTurnWindow field.                                        |
| `ThresholdStep`                                                    | `*float64`                                                         | :heavy_minus_sign:                                                 | The thresholdStep field.                                           |