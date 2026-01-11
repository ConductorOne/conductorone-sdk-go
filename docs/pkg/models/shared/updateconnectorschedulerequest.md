# UpdateConnectorScheduleRequest

The UpdateConnectorScheduleRequest message.

This message contains a oneof named schedule. Only a single field of the following list may be set at a time:
  - cron



## Fields

| Field                                                                                | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ConnectorScheduleCron`                                                              | [*shared.ConnectorScheduleCron](../../../pkg/models/shared/connectorschedulecron.md) | :heavy_minus_sign:                                                                   | The ConnectorScheduleCron message.                                                   |