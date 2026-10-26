# XAAScopeDiscoveryFinding

XAAScopeDiscoveryFinding is one per-URI or per-value observation from a
 scope discovery run.


## Fields

| Field                                                                                              | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `Code`                                                                                             | [*shared.XAAScopeDiscoveryFindingCode](../../../pkg/models/shared/xaascopediscoveryfindingcode.md) | :heavy_minus_sign:                                                                                 | What was observed.                                                                                 |
| `Message`                                                                                          | `*string`                                                                                          | :heavy_minus_sign:                                                                                 | Human-readable detail.                                                                             |
| `ResourceURI`                                                                                      | `*string`                                                                                          | :heavy_minus_sign:                                                                                 | The resource URI the finding applies to. Empty for the authorization<br/> server fallback.         |
| `ScopeValue`                                                                                       | `*string`                                                                                          | :heavy_minus_sign:                                                                                 | The scope value the finding applies to, for per-value findings.                                    |