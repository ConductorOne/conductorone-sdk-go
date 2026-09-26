# VirtualMCPServerServiceListResponse

VirtualMCPServerServiceListResponse contains one page of Virtual MCP Servers.


## Fields

| Field                                                                       | Type                                                                        | Required                                                                    | Description                                                                 |
| --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- |
| `List`                                                                      | [][shared.VirtualMCPServer](../../../pkg/models/shared/virtualmcpserver.md) | :heavy_minus_sign:                                                          | Virtual MCP Servers in this page.                                           |
| `NextPageToken`                                                             | `*string`                                                                   | :heavy_minus_sign:                                                          | Token to retrieve the next page. Empty when there are no more results.      |