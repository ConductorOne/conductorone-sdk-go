# MCPAccessProfileServiceSearchAccessProfilesResponse

MCPAccessProfileServiceSearchAccessProfilesResponse returns one page of
 tenant-wide MCP access profiles.


## Fields

| Field                                                                       | Type                                                                        | Required                                                                    | Description                                                                 |
| --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- |
| `NextPageToken`                                                             | `*string`                                                                   | :heavy_minus_sign:                                                          | Token for next page.                                                        |
| `Profiles`                                                                  | [][shared.MCPAccessProfile](../../../pkg/models/shared/mcpaccessprofile.md) | :heavy_minus_sign:                                                          | The page of matching MCP access profiles.                                   |