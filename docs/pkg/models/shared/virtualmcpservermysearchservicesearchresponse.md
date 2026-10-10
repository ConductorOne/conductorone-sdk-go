# VirtualMCPServerMySearchServiceSearchResponse

VirtualMCPServerMySearchServiceSearchResponse contains one page of the
 caller's currently available Virtual MCP Servers.


## Fields

| Field                                                                           | Type                                                                            | Required                                                                        | Description                                                                     |
| ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- |
| `List`                                                                          | [][shared.MyVirtualMCPServer](../../../pkg/models/shared/myvirtualmcpserver.md) | :heavy_minus_sign:                                                              | Available Virtual MCP Servers in this page.                                     |
| `NextPageToken`                                                                 | `*string`                                                                       | :heavy_minus_sign:                                                              | Opaque token that continues the PostgreSQL-backed search.                       |