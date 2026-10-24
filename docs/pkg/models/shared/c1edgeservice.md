# C1EdgeService

C1EdgeService names the hosted Edge or AuthZEN service a trust under the
 C1_EDGE provider lets the fleet act for.


## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `AppID`                                                            | `*string`                                                          | :heavy_minus_sign:                                                 | App owning the hosted service.                                     |
| `ResourceID`                                                       | `*string`                                                          | :heavy_minus_sign:                                                 | ID of the hosted Edge or AuthZEN server.                           |
| `ResourceKind`                                                     | [*shared.ResourceKind](../../../pkg/models/shared/resourcekind.md) | :heavy_minus_sign:                                                 | Kind of the hosted service.                                        |