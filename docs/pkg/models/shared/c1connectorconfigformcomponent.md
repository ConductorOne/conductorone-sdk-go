# C1ConnectorConfigFormComponent

C1ConnectorConfigFormComponent renders the shared admin connector-settings form inside an
 A2UI surface. The frontend resolves the catalog, connector, and config schema itself from
 the ids below, keeping the configuration field values out of the agent's data model — the
 agent never receives API keys, passwords, or other secrets entered by the user.


## Fields

| Field                       | Type                        | Required                    | Description                 |
| --------------------------- | --------------------------- | --------------------------- | --------------------------- |
| `AppID`                     | `*string`                   | :heavy_minus_sign:          | The appId field.            |
| `ConnectorID`               | `*string`                   | :heavy_minus_sign:          | The connectorId field.      |
| `SkipActionName`            | `*string`                   | :heavy_minus_sign:          | The skipActionName field.   |
| `SubmitActionName`          | `*string`                   | :heavy_minus_sign:          | The submitActionName field. |