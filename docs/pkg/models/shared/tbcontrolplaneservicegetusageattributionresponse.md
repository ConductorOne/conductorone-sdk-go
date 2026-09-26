# TBControlPlaneServiceGetUsageAttributionResponse

The TBControlPlaneServiceGetUsageAttributionResponse message.


## Fields

| Field                                                                                 | Type                                                                                  | Required                                                                              | Description                                                                           |
| ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- |
| `Rows`                                                                                | [][shared.TBUsageAttributionRow](../../../pkg/models/shared/tbusageattributionrow.md) | :heavy_minus_sign:                                                                    | Ordered by (input_tokens + output_tokens) descending, limited to<br/> page_size.      |