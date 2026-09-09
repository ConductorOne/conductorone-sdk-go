# ShadowMcpEvidence

ShadowMcpEvidence carries counts only -- occurrence detail (which users,
 devices, harnesses) is resolved dynamically, not stored here.


## Fields

| Field                                                                       | Type                                                                        | Required                                                                    | Description                                                                 |
| --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- |
| `DeviceCount`                                                               | `*int64`                                                                    | :heavy_minus_sign:                                                          | Distinct devices observed running the product outside the governed gateway. |
| `UserCount`                                                                 | `*int64`                                                                    | :heavy_minus_sign:                                                          | Distinct users observed running the product outside the governed gateway.   |