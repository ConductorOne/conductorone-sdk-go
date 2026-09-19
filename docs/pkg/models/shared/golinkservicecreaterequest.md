# GoLinkServiceCreateRequest

The GoLinkServiceCreateRequest message.


## Fields

| Field                                                                          | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `Description`                                                                  | `*string`                                                                      | :heavy_minus_sign:                                                             | The description field.                                                         |
| `DisplayName`                                                                  | `*string`                                                                      | :heavy_minus_sign:                                                             | The displayName field.                                                         |
| `Prerequisite`                                                                 | [*shared.GoLinkPrerequisite](../../../pkg/models/shared/golinkprerequisite.md) | :heavy_minus_sign:                                                             | N/A                                                                            |
| `Routes`                                                                       | [][shared.GoLinkRouteConfig](../../../pkg/models/shared/golinkrouteconfig.md)  | :heavy_minus_sign:                                                             | The routes field.                                                              |
| `Target`                                                                       | [*shared.GoLinkTargetInput](../../../pkg/models/shared/golinktargetinput.md)   | :heavy_minus_sign:                                                             | N/A                                                                            |