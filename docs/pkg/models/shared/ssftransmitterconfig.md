# SSFTransmitterConfig

SSFTransmitterConfig selects which outbound shared-signals streams this
 session emits security events to. Each stream's delivery endpoint,
 authentication, and per-event allowlist are configured on the stream itself;
 this policy just lists the stream IDs and the event types to emit.


## Fields

| Field                                                                | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `Enabled`                                                            | `*bool`                                                              | :heavy_minus_sign:                                                   | Whether outbound shared-signals emission is enabled for this policy. |
| `EventTypes`                                                         | []`string`                                                           | :heavy_minus_sign:                                                   | The shared-signals event types to emit at the policy level.          |
| `SsfTransmitterStreamIds`                                            | []`string`                                                           | :heavy_minus_sign:                                                   | The outbound stream IDs this policy emits to.                        |