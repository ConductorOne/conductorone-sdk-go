# MCPAccessProfileToolBindingServiceGetAccessProfilesForToolsRequest

MCPAccessProfileToolBindingServiceGetAccessProfilesForToolsRequest requests
 the access profiles bound to a batch of MCP tools.


## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `McpToolIds`                                                          | []`string`                                                            | :heavy_minus_sign:                                                    | MCP tool IDs to look up. Sized to match frontend MultiGet batch size. |