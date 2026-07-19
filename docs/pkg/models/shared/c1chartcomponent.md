# C1ChartComponent

C1ChartComponent renders a chart from typed data. The data oneof is keyed by
 shape — each shape carries its own style enum, so an invalid combination
 (e.g. a pie chart with a time axis) is unrepresentable.

This message contains a oneof named data. Only a single field of the following list may be set at a time:
  - timeSeries
  - categorical



## Fields

| Field                                                                                  | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ArtifactURL`                                                                          | [*shared.DynamicString](../../../pkg/models/shared/dynamicstring.md)                   | :heavy_minus_sign:                                                                     | N/A                                                                                    |
| `Categorical`                                                                          | [*shared.C1ChartCategoricalData](../../../pkg/models/shared/c1chartcategoricaldata.md) | :heavy_minus_sign:                                                                     | N/A                                                                                    |
| `Sources`                                                                              | [][shared.C1ChartSource](../../../pkg/models/shared/c1chartsource.md)                  | :heavy_minus_sign:                                                                     | Provenance: the queries the producing function ran.                                    |
| `TimeSeries`                                                                           | [*shared.C1ChartTimeSeriesData](../../../pkg/models/shared/c1charttimeseriesdata.md)   | :heavy_minus_sign:                                                                     | N/A                                                                                    |
| `Title`                                                                                | [*shared.DynamicString](../../../pkg/models/shared/dynamicstring.md)                   | :heavy_minus_sign:                                                                     | N/A                                                                                    |