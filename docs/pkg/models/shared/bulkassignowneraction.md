# BulkAssignOwnerAction

The BulkAssignOwnerAction message.


## Fields

| Field                                                                                                         | Type                                                                                                          | Required                                                                                                      | Description                                                                                                   |
| ------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------- |
| `AssigneeIdentityUserID`                                                                                      | `*string`                                                                                                     | :heavy_minus_sign:                                                                                            | Empty is allowed: rpc.go falls back to the deprecated owner field's<br/> identity_user_id arm when this is unset. |
| `Owner`                                                                                                       | [*shared.FindingOwnerRef](../../../pkg/models/shared/findingownerref.md)                                      | :heavy_minus_sign:                                                                                            | N/A                                                                                                           |