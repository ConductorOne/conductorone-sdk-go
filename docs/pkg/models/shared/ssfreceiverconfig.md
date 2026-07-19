# SSFReceiverConfig

SSFReceiverConfig selects which inbound shared-signals streams this session
 trusts. Each stream's issuer, keys, expected audience, and per-event actions
 are configured on the stream itself; this policy just lists the stream IDs.


## Fields

| Field                                                                  | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `Enabled`                                                              | `*bool`                                                                | :heavy_minus_sign:                                                     | Whether inbound shared-signals consumption is enabled for this policy. |
| `SsfReceiverStreamIds`                                                 | []`string`                                                             | :heavy_minus_sign:                                                     | The inbound stream IDs this policy trusts.                             |