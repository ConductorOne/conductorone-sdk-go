# ScheduleTriggerAppUser

The ScheduleTriggerAppUser message.


## Fields

| Field                                      | Type                                       | Required                                   | Description                                |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| `AppID`                                    | **string*                                  | :heavy_minus_sign:                         | The appId field.                           |
| `Condition`                                | **string*                                  | :heavy_minus_sign:                         | The condition field.                       |
| `CronSpec`                                 | **string*                                  | :heavy_minus_sign:                         | The cronSpec field.                        |
| `Start`                                    | [*time.Time](https://pkg.go.dev/time#Time) | :heavy_minus_sign:                         | N/A                                        |
| `Timezone`                                 | **string*                                  | :heavy_minus_sign:                         | The timezone field.                        |