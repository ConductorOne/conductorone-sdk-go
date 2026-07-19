# AnnouncedTunnelService

AnnouncedTunnelService is one service entry the appliance declared in its
 wormhole HELLO frame. Read live from the discovery store; not persisted.


## Fields

| Field                                                         | Type                                                          | Required                                                      | Description                                                   |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| `Name`                                                        | `*string`                                                     | :heavy_minus_sign:                                            | Logical name of the service as declared by the appliance.     |
| `Port`                                                        | `*int64`                                                      | :heavy_minus_sign:                                            | TCP port the service listens on inside the appliance network. |
| `ServicePath`                                                 | `*string`                                                     | :heavy_minus_sign:                                            | Optional URL path prefix for the service.                     |
| `ServiceType`                                                 | `*string`                                                     | :heavy_minus_sign:                                            | Application-level protocol type (e.g. "http", "grpc").        |
| `TransportType`                                               | `*string`                                                     | :heavy_minus_sign:                                            | Transport protocol used by the wormhole tunnel (e.g. "tcp").  |