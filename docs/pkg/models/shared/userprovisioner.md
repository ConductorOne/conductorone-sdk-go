# UserProvisioner

UserProvisioner assigns specific users as provisioners.


## Fields

| Field                                          | Type                                           | Required                                       | Description                                    |
| ---------------------------------------------- | ---------------------------------------------- | ---------------------------------------------- | ---------------------------------------------- |
| `AllowReassignment`                            | **bool*                                        | :heavy_minus_sign:                             | Whether the provisioner can reassign the task. |
| `UserIds`                                      | []*string*                                     | :heavy_minus_sign:                             | The user IDs to assign as provisioners.        |