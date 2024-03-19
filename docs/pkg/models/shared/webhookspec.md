# WebhookSpec

The WebhookSpec message.


## Fields

| Field                                                                                                          | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `Destination`                                                                                                  | **string*                                                                                                      | :heavy_minus_sign:                                                                                             | The destination field.                                                                                         |
| `Payload`                                                                                                      | [*shared.Payload](../../../pkg/models/shared/payload.md)                                                       | :heavy_minus_sign:                                                                                             | Contains an arbitrary serialized message along with a @type that describes the type of the serialized message. |