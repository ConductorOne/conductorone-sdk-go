# NotificationConfig

Controls which email notifications are sent during the access review lifecycle.


## Fields

| Field                                                                         | Type                                                                          | Required                                                                      | Description                                                                   |
| ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
| `SendClose`                                                                   | `*bool`                                                                       | :heavy_minus_sign:                                                            | Whether to send a notification when the campaign is closed.                   |
| `SendKickoff`                                                                 | `*bool`                                                                       | :heavy_minus_sign:                                                            | Whether to send a notification when the campaign is started.                  |
| `SendReminders`                                                               | `*bool`                                                                       | :heavy_minus_sign:                                                            | Whether to send periodic reminder emails to reviewers with outstanding tasks. |