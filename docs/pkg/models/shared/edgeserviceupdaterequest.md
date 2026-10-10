# EdgeServiceUpdateRequest

EdgeServiceUpdateRequest updates an existing Edge via update_mask. Editable
 paths: mode, enabled_kinds, egress_rules, credential_injections,
 inference_routes, allow_inference_bearer_tokens. The legacy inference mask
 may clear or replace one graph.


## Fields

| Field                                              | Type                                               | Required                                           | Description                                        |
| -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- |
| `Edge`                                             | [*shared.Edge](../../../pkg/models/shared/edge.md) | :heavy_minus_sign:                                 | N/A                                                |
| `UpdateMask`                                       | `*string`                                          | :heavy_minus_sign:                                 | N/A                                                |