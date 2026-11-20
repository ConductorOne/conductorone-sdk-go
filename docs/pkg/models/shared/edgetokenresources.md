# EdgeTokenResources

EdgeTokenResources names the resource URI for each supported Edge service.
 A field is empty when its capability is not in enabled_kinds.


## Fields

| Field                                                                       | Type                                                                        | Required                                                                    | Description                                                                 |
| --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- |
| `Egress`                                                                    | `*string`                                                                   | :heavy_minus_sign:                                                          | Resource URI used with the egress.connect scope.                            |
| `Inference`                                                                 | `*string`                                                                   | :heavy_minus_sign:                                                          | Resource URI used with the inference.models.read / inference.invoke scopes. |