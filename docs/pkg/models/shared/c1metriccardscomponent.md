# C1MetricCardsComponent

C1MetricCardsComponent renders a row of aggregate stat cards.


## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `Cards`                                                               | [][shared.C1MetricCard](../../../pkg/models/shared/c1metriccard.md)   | :heavy_minus_sign:                                                    | The cards field.                                                      |
| `Sources`                                                             | [][shared.C1ChartSource](../../../pkg/models/shared/c1chartsource.md) | :heavy_minus_sign:                                                    | Provenance: the queries the producing function ran.                   |
| `Title`                                                               | [*shared.DynamicString](../../../pkg/models/shared/dynamicstring.md)  | :heavy_minus_sign:                                                    | N/A                                                                   |