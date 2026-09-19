# VirtualMCPToolsetBindingServiceListResponse

VirtualMCPToolsetBindingServiceListResponse contains one page of Toolset assignments.


## Fields

| Field                                                                                       | Type                                                                                        | Required                                                                                    | Description                                                                                 |
| ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- |
| `Bindings`                                                                                  | [][shared.VirtualMCPToolsetBinding](../../../pkg/models/shared/virtualmcptoolsetbinding.md) | :heavy_minus_sign:                                                                          | Toolset assignments on the Virtual MCP Server.                                              |
| `NextPageToken`                                                                             | `*string`                                                                                   | :heavy_minus_sign:                                                                          | Token to retrieve the next page. Empty when there are no more results.                      |