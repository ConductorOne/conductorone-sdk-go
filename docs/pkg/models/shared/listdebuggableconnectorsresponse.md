# ListDebuggableConnectorsResponse

ListDebuggableConnectorsResponse contains eligible connectors from the inspected page.


## Fields

| Field                                                                                  | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `Connectors`                                                                           | [][shared.DebuggableConnector](../../../pkg/models/shared/debuggableconnector.md)      | :heavy_minus_sign:                                                                     | Connectors the caller is authorized to manage.                                         |
| `NextPageToken`                                                                        | `*string`                                                                              | :heavy_minus_sign:                                                                     | Continue while this token is nonempty, even when the filtered connector list is empty. |