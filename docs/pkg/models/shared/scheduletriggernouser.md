# ScheduleTriggerNoUser

ScheduleTriggerNoUser fires on a cron schedule with no subject user (e.g. reports, syncs, orchestration).
 Minimum cron interval is enforced at 1 hour in validation.


## Fields

| Field                                      | Type                                       | Required                                   | Description                                |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| `Advanced`                                 | `*bool`                                    | :heavy_minus_sign:                         | The advanced field.                        |
| `CronSpec`                                 | `*string`                                  | :heavy_minus_sign:                         | The cronSpec field.                        |
| `Start`                                    | [*time.Time](https://pkg.go.dev/time#Time) | :heavy_minus_sign:                         | N/A                                        |
| `Timezone`                                 | `*string`                                  | :heavy_minus_sign:                         | The timezone field.                        |