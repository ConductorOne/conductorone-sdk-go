# ConnectorScheduleCron

A cron-based schedule definition for connector syncs.


## Fields

| Field                                                                       | Type                                                                        | Required                                                                    | Description                                                                 |
| --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- |
| `CronSpec`                                                                  | `*string`                                                                   | :heavy_minus_sign:                                                          | The cron expression defining the sync schedule.                             |
| `Timezone`                                                                  | `*string`                                                                   | :heavy_minus_sign:                                                          | The IANA timezone name for the cron schedule (e.g., "America/Los_Angeles"). |