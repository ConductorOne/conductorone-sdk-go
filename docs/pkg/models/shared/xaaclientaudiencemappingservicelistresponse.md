# XAAClientAudienceMappingServiceListResponse

XAAClientAudienceMappingServiceListResponse returns a page of mappings.


## Fields

| Field                                                                                       | Type                                                                                        | Required                                                                                    | Description                                                                                 |
| ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- |
| `ClientAudienceMappings`                                                                    | [][shared.XAAClientAudienceMapping](../../../pkg/models/shared/xaaclientaudiencemapping.md) | :heavy_minus_sign:                                                                          | The page of mappings.                                                                       |
| `NextPageToken`                                                                             | `*string`                                                                                   | :heavy_minus_sign:                                                                          | Token for the next page, or empty if there are no more results.                             |