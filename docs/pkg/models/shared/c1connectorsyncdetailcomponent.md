# C1ConnectorSyncDetailComponent

C1ConnectorSyncDetailComponent renders the same live card as
 C1ConnectorSyncProgressComponent but pre-expanded with the phase checklist,
 live count tiles, and "What's happening" explainer visible from the first
 paint. Intended for message-body placement — emit one after each
 `submit_app_config` so the transcript carries a clear "this is what just
 happened" receipt for the connector the user just connected.


## Fields

| Field                  | Type                   | Required               | Description            |
| ---------------------- | ---------------------- | ---------------------- | ---------------------- |
| `AppID`                | `*string`              | :heavy_minus_sign:     | The appId field.       |
| `ConnectorID`          | `*string`              | :heavy_minus_sign:     | The connectorId field. |
| `Title`                | `*string`              | :heavy_minus_sign:     | The title field.       |