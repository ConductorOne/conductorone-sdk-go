# GroupProvisioner

GroupProvisioner resolves to members of a specific group.


## Fields

| Field                                            | Type                                             | Required                                         | Description                                      |
| ------------------------------------------------ | ------------------------------------------------ | ------------------------------------------------ | ------------------------------------------------ |
| `AllowReassignment`                              | **bool*                                          | :heavy_minus_sign:                               | Whether the provisioner can reassign the task.   |
| `AppGroupID`                                     | **string*                                        | :heavy_minus_sign:                               | The app group ID (entitlement ID).               |
| `AppID`                                          | **string*                                        | :heavy_minus_sign:                               | The app ID containing the group.                 |
| `FallbackUserIds`                                | []*string*                                       | :heavy_minus_sign:                               | Fallback user IDs if no group members are found. |