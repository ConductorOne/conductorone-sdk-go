# AppResource

 The app resource message is a single resource that can have entitlements.



## Fields

| Field                                                  | Type                                                   | Required                                               | Description                                            |
| ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ |
| `AppID`                                                | **string*                                              | :heavy_minus_sign:                                     |  The app that this resource belongs to.<br/>           |
| `AppResourceTypeID`                                    | **string*                                              | :heavy_minus_sign:                                     |  The resource type that this resource is.<br/>         |
| `CreatedAt`                                            | [*time.Time](https://pkg.go.dev/time#Time)             | :heavy_minus_sign:                                     | N/A                                                    |
| `CustomDescription`                                    | **string*                                              | :heavy_minus_sign:                                     |  A custom description that can be set for a resource.<br/> |
| `DeletedAt`                                            | [*time.Time](https://pkg.go.dev/time#Time)             | :heavy_minus_sign:                                     | N/A                                                    |
| `Description`                                          | **string*                                              | :heavy_minus_sign:                                     |  The description set for the resource.<br/>            |
| `DisplayName`                                          | **string*                                              | :heavy_minus_sign:                                     |  The display name for this resource.<br/>              |
| `GrantCount`                                           | **string*                                              | :heavy_minus_sign:                                     |  The number of grants to this resource.<br/>           |
| `ID`                                                   | **string*                                              | :heavy_minus_sign:                                     |  The id of the resource.<br/>                          |
| `UpdatedAt`                                            | [*time.Time](https://pkg.go.dev/time#Time)             | :heavy_minus_sign:                                     | N/A                                                    |