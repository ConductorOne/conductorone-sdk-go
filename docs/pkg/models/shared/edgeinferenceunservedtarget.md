# EdgeInferenceUnservedTarget

EdgeInferenceUnservedTarget is a target of a configured inference route that connected Edge processes report
 they cannot serve.


## Fields

| Field                                                       | Type                                                        | Required                                                    | Description                                                 |
| ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- |
| `Partial`                                                   | `*bool`                                                     | :heavy_minus_sign:                                          | True when some connected Edge processes can still serve it. |
| `Reason`                                                    | `*string`                                                   | :heavy_minus_sign:                                          | The reason field.                                           |
| `RouteID`                                                   | `*string`                                                   | :heavy_minus_sign:                                          | The routeId field.                                          |
| `TargetID`                                                  | `*string`                                                   | :heavy_minus_sign:                                          | The targetId field.                                         |