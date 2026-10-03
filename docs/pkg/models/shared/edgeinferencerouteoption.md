# EdgeInferenceRouteOption

EdgeInferenceRouteOption describes an already configured route to callers
 with application management access. It never includes credential material.


## Fields

| Field                                                       | Type                                                        | Required                                                    | Description                                                 |
| ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- |
| `DisplayName`                                               | `*string`                                                   | :heavy_minus_sign:                                          | Human-readable route name.                                  |
| `RouteID`                                                   | `*string`                                                   | :heavy_minus_sign:                                          | Configured route identifier.                                |
| `UpstreamModelID`                                           | `*string`                                                   | :heavy_minus_sign:                                          | First configured target's model, for existing list clients. |