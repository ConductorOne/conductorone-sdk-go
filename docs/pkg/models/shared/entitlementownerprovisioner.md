# EntitlementOwnerProvisioner

EntitlementOwnerProvisioner resolves to entitlement owners.


## Fields

| Field                                                 | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `AllowReassignment`                                   | **bool*                                               | :heavy_minus_sign:                                    | Whether the provisioner can reassign the task.        |
| `FallbackUserIds`                                     | []*string*                                            | :heavy_minus_sign:                                    | Fallback user IDs if no entitlement owners are found. |