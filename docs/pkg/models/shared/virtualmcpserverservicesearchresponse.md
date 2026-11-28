# VirtualMCPServerServiceSearchResponse

VirtualMCPServerServiceSearchResponse contains one page of matching Virtual
 MCP Servers, ordered by display name and then ID.


## Fields

| Field                                                                       | Type                                                                        | Required                                                                    | Description                                                                 |
| --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- |
| `List`                                                                      | [][shared.VirtualMCPServer](../../../pkg/models/shared/virtualmcpserver.md) | :heavy_minus_sign:                                                          | Matching Virtual MCP Servers in this page.                                  |
| `NextPageToken`                                                             | `*string`                                                                   | :heavy_minus_sign:                                                          | Token to retrieve the next page. Empty when there are no more results.      |