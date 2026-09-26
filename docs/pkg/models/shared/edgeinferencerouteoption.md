# EdgeInferenceRouteOption

EdgeInferenceRouteOption is a server-approved inference route a caller with
 application management access may select. It never includes endpoint or
 credential material.


## Fields

| Field                                | Type                                 | Required                             | Description                          |
| ------------------------------------ | ------------------------------------ | ------------------------------------ | ------------------------------------ |
| `DisplayName`                        | `*string`                            | :heavy_minus_sign:                   | Human-readable route name.           |
| `RouteID`                            | `*string`                            | :heavy_minus_sign:                   | Server-approved route identifier.    |
| `UpstreamModelID`                    | `*string`                            | :heavy_minus_sign:                   | Upstream model paired with route_id. |