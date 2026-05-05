# AppResourceRef

A reference to a specific app resource by its composite key.


## Fields

| Field                                                      | Type                                                       | Required                                                   | Description                                                |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `AppID`                                                    | `*string`                                                  | :heavy_minus_sign:                                         | The ID of the app that owns the resource.                  |
| `AppResourceTypeID`                                        | `*string`                                                  | :heavy_minus_sign:                                         | The ID of the resource type that classifies this resource. |
| `ID`                                                       | `*string`                                                  | :heavy_minus_sign:                                         | The unique ID of the app resource.                         |