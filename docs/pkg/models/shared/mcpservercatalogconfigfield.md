# MCPServerCatalogConfigField

MCPServerCatalogConfigField describes a single extra configuration field for an MCP server catalog entry.


## Fields

| Field                                                                      | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `Default`                                                                  | `*string`                                                                  | :heavy_minus_sign:                                                         | Default value the registration form prefills. Ignored when secret.         |
| `Description`                                                              | `*string`                                                                  | :heavy_minus_sign:                                                         | Help text describing the field.                                            |
| `DisplayName`                                                              | `*string`                                                                  | :heavy_minus_sign:                                                         | Human-readable label for the field.                                        |
| `Name`                                                                     | `*string`                                                                  | :heavy_minus_sign:                                                         | Machine-readable field name (used as the map key in config_fields).        |
| `Placeholder`                                                              | `*string`                                                                  | :heavy_minus_sign:                                                         | Placeholder text shown in an empty input.                                  |
| `Required`                                                                 | `*bool`                                                                    | :heavy_minus_sign:                                                         | Whether this field must be provided.                                       |
| `Secret`                                                                   | `*bool`                                                                    | :heavy_minus_sign:                                                         | Whether the field value should be treated as a secret (e.g. masked in UI). |