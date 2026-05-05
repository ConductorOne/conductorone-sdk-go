# UpdateSessionSettingsRequest

The UpdateSessionSettingsRequest message.


## Fields

| Field                                                                                                        | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `SessionSettings`                                                                                            | [*shared.SessionSettings](../../../pkg/models/shared/sessionsettings.md)                                     | :heavy_minus_sign:                                                                                           | SessionSettings configures session security for the tenant, including timeouts and per-role IP restrictions. |
| `UpdateMask`                                                                                                 | `*string`                                                                                                    | :heavy_minus_sign:                                                                                           | N/A                                                                                                          |