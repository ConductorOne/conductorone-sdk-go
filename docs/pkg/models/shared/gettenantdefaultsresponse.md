# GetTenantDefaultsResponse

GetTenantDefaultsResponse contains the tenant-default subset of AI governance
 settings applied to newly registered MCP servers and tools.


## Fields

| Field                                                                                         | Type                                                                                          | Required                                                                                      | Description                                                                                   |
| --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- |
| `RequireToolApproval`                                                                         | `*bool`                                                                                       | :heavy_minus_sign:                                                                            | Whether newly discovered tools require admin approval before they can be<br/> granted or invoked. |