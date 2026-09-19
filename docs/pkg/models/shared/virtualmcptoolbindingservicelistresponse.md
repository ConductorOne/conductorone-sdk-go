# VirtualMCPToolBindingServiceListResponse

VirtualMCPToolBindingServiceListResponse contains one page of tool assignments.


## Fields

| Field                                                                                 | Type                                                                                  | Required                                                                              | Description                                                                           |
| ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- |
| `Bindings`                                                                            | [][shared.VirtualMCPToolBinding](../../../pkg/models/shared/virtualmcptoolbinding.md) | :heavy_minus_sign:                                                                    | Individual tool assignments on the Virtual MCP Server.                                |
| `NextPageToken`                                                                       | `*string`                                                                             | :heavy_minus_sign:                                                                    | Token to retrieve the next page. Empty when there are no more results.                |