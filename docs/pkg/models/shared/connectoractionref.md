# ConnectorActionRef

ConnectorActionRef describes dispatch through a connector's built-in
 GrantManagerService Grant / Revoke RPC — i.e. the default connector
 operation, used for synthesized tickets like scope-role requests.


## Fields

| Field                                                        | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `AppID`                                                      | `*string`                                                    | :heavy_minus_sign:                                           | The app whose connector handles the operation.               |
| `ConnectorID`                                                | `*string`                                                    | :heavy_minus_sign:                                           | The connector that will execute the Grant / Revoke.          |
| `Operation`                                                  | [*shared.Operation](../../../pkg/models/shared/operation.md) | :heavy_minus_sign:                                           | Which connector RPC this dispatches to.                      |