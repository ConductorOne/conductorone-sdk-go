# UpdateConnectorScheduleRequest

The UpdateConnectorScheduleRequest message contains the fields required to update a connector's sync schedule.

This message contains a oneof named schedule. Only a single field of the following list may be set at a time:
  - cron



## Fields

| Field                                                                                | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `Cron`                                                                               | [*shared.ConnectorScheduleCron](../../../pkg/models/shared/connectorschedulecron.md) | :heavy_minus_sign:                                                                   | N/A                                                                                  |