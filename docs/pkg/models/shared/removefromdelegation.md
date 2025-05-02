# RemoveFromDelegation

RemoveFromDelegation: find all users that have the target user as their delegated user, and clear the delegation.


## Fields

| Field                                                    | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `UserRef`                                                | [*shared.UserRef](../../../pkg/models/shared/userref.md) | :heavy_minus_sign:                                       | A reference to a user.                                   |
| `UserIDCel`                                              | **string*                                                | :heavy_minus_sign:                                       | The userIdCel field.                                     |