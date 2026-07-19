# TunnelAppliance

TunnelAppliance is the live state of the customer-side appliance for one
 bridge.


## Fields

| Field                                                                                | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `AnnouncedServiceCount`                                                              | `*int64`                                                                             | :heavy_minus_sign:                                                                   | Number of services the appliance is currently announcing.                            |
| `LastSeenAt`                                                                         | [*time.Time](https://pkg.go.dev/time#Time)                                           | :heavy_minus_sign:                                                                   | N/A                                                                                  |
| `Links`                                                                              | [][shared.TunnelApplianceLink](../../../pkg/models/shared/tunnelappliancelink.md)    | :heavy_minus_sign:                                                                   | Wormhole relays currently holding a Link for this bridge. Typically of<br/> length 1. |
| `Status`                                                                             | [*shared.TunnelApplianceStatus](../../../pkg/models/shared/tunnelappliancestatus.md) | :heavy_minus_sign:                                                                   | The status field.                                                                    |