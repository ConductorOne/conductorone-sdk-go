# C1TableComponent

C1TableComponent renders a tabular view: typed columns + rows, capped and
 paginated client-side; the full data set lives behind the artifact link.


## Fields

| Field                                                                   | Type                                                                    | Required                                                                | Description                                                             |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `ArtifactURL`                                                           | [*shared.DynamicString](../../../pkg/models/shared/dynamicstring.md)    | :heavy_minus_sign:                                                      | N/A                                                                     |
| `Columns`                                                               | []`string`                                                              | :heavy_minus_sign:                                                      | The columns field.                                                      |
| `PageSize`                                                              | `*int`                                                                  | :heavy_minus_sign:                                                      | Rows per page for client-side pagination; 0 shows all rows on one page. |
| `Rows`                                                                  | [][shared.C1TableRow](../../../pkg/models/shared/c1tablerow.md)         | :heavy_minus_sign:                                                      | The rows field.                                                         |
| `Sources`                                                               | [][shared.C1ChartSource](../../../pkg/models/shared/c1chartsource.md)   | :heavy_minus_sign:                                                      | Provenance: the queries the producing function ran.                     |
| `Title`                                                                 | [*shared.DynamicString](../../../pkg/models/shared/dynamicstring.md)    | :heavy_minus_sign:                                                      | N/A                                                                     |
| `TotalRows`                                                             | `*int64`                                                                | :heavy_minus_sign:                                                      | Full count when rows are truncated.                                     |