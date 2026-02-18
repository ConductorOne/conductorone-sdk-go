# AppOwnerProvisioner

AppOwnerProvisioner resolves to app owners.


## Fields

| Field                                          | Type                                           | Required                                       | Description                                    |
| ---------------------------------------------- | ---------------------------------------------- | ---------------------------------------------- | ---------------------------------------------- |
| `AllowReassignment`                            | **bool*                                        | :heavy_minus_sign:                             | Whether the provisioner can reassign the task. |
| `FallbackUserIds`                              | []*string*                                     | :heavy_minus_sign:                             | Fallback user IDs if no app owners are found.  |