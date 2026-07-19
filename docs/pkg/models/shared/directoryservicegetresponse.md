# DirectoryServiceGetResponse

The Directory Service Get Response returns a directory view with a directory and JSONPATHs indicating the
 location in the expanded array that items are expanded as indicated by the expand mask in the request.


## Fields

| Field                                                                                                             | Type                                                                                                              | Required                                                                                                          | Description                                                                                                       |
| ----------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------- |
| `DirectoryView`                                                                                                   | [*shared.DirectoryView](../../../pkg/models/shared/directoryview.md)                                              | :heavy_minus_sign:                                                                                                | N/A                                                                                                               |
| `Expanded`                                                                                                        | [][shared.DirectoryServiceGetResponseExpanded](../../../pkg/models/shared/directoryservicegetresponseexpanded.md) | :heavy_minus_sign:                                                                                                | List of serialized related objects.                                                                               |