# AccessReviewColumnConfig

Configuration for which columns are visible in the reviewer task list.


## Fields

| Field                                                                                                                    | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `Columns`                                                                                                                | [][shared.Columns](../../../pkg/models/shared/columns.md)                                                                | :heavy_minus_sign:                                                                                                       | Ordered list of columns visible to reviewers.<br/> If empty, the default column set for the campaign's default_view is used. |