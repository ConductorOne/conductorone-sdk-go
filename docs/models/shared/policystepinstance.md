# PolicyStepInstance

The PolicyStepInstance message.

This message contains a oneof named instance. Only a single field of the following list may be set at a time:
  - approval
  - provision



## Fields

| Field                                                                            | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ID`                                                                             | **string*                                                                        | :heavy_minus_sign:                                                               | The id field.                                                                    |
| `Instance`                                                                       | [*PolicyStepInstanceInstance](../../models/shared/policystepinstanceinstance.md) | :heavy_minus_sign:                                                               | N/A                                                                              |
| `State`                                                                          | [*PolicyStepInstanceState](../../models/shared/policystepinstancestate.md)       | :heavy_minus_sign:                                                               | The state field.                                                                 |