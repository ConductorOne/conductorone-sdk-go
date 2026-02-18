# ManagerProvisioner

ManagerProvisioner resolves to the user's manager.


## Fields

| Field                                          | Type                                           | Required                                       | Description                                    |
| ---------------------------------------------- | ---------------------------------------------- | ---------------------------------------------- | ---------------------------------------------- |
| `AllowReassignment`                            | **bool*                                        | :heavy_minus_sign:                             | Whether the provisioner can reassign the task. |
| `FallbackUserIds`                              | []*string*                                     | :heavy_minus_sign:                             | Fallback user IDs if no manager is found.      |