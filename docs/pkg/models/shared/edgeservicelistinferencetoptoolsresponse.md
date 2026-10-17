# EdgeServiceListInferenceTopToolsResponse

The EdgeServiceListInferenceTopToolsResponse message.


## Fields

| Field                                                                                                                    | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `Tools`                                                                                                                  | [][shared.TBTopTool](../../../pkg/models/shared/tbtoptool.md)                                                            | :heavy_minus_sign:                                                                                                       | Ordered by call_count descending, limited to page_size. No pagination<br/> cursor -- this is a fixed top-N, not a full list. |