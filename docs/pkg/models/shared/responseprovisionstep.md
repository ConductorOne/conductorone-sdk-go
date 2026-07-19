# ResponseProvisionStep

The ResponseProvisionStep message.

This message contains a oneof named outcome. Only a single field of the following list may be set at a time:
  - complete
  - errored



## Fields

| Field                                                                                                           | Type                                                                                                            | Required                                                                                                        | Description                                                                                                     |
| --------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------- |
| `Complete`                                                                                                      | [*shared.ResponseProvisionStepComplete](../../../pkg/models/shared/responseprovisionstepcomplete.md)            | :heavy_minus_sign:                                                                                              | N/A                                                                                                             |
| `Errored`                                                                                                       | [*shared.ResponseProvisionStepErrored](../../../pkg/models/shared/responseprovisionsteperrored.md)              | :heavy_minus_sign:                                                                                              | N/A                                                                                                             |
| `Version`                                                                                                       | `*string`                                                                                                       | :heavy_minus_sign:                                                                                              | version contains the constant value "v1". Future versions of the Webhook Response<br/> will use a different string. |