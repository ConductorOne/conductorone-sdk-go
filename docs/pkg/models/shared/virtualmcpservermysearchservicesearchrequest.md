# VirtualMCPServerMySearchServiceSearchRequest

VirtualMCPServerMySearchServiceSearchRequest searches Virtual MCP Servers
 the caller's own principal has a live `use` grant for. Only user principals
 may call this API.


## Fields

| Field                                                             | Type                                                              | Required                                                          | Description                                                       |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `PageSize`                                                        | `*int`                                                            | :heavy_minus_sign:                                                | Maximum number of matching servers to return. The maximum is 100. |
| `PageToken`                                                       | `*string`                                                         | :heavy_minus_sign:                                                | Opaque token that continues the PostgreSQL-backed search.         |