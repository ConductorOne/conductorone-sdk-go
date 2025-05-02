# UpdateUser

The UpdateUser message.

This message contains a oneof named user. Only a single field of the following list may be set at a time:
  - userIdCel
  - userRef


This message contains a oneof named user_status. Only a single field of the following list may be set at a time:
  - userStatusEnum
  - userStatusCel



## Fields

| Field                                                                                                                                                     | Type                                                                                                                                                      | Required                                                                                                                                                  | Description                                                                                                                                               |
| --------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `UserRef`                                                                                                                                                 | [*shared.UserRef](../../../pkg/models/shared/userref.md)                                                                                                  | :heavy_minus_sign:                                                                                                                                        | A reference to a user.                                                                                                                                    |
| `UserIDCel`                                                                                                                                               | **string*                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                        | The userIdCel field.<br/>This field is part of the `user` oneof.<br/>See the documentation for `c1.api.workflows.v1beta.UpdateUser` for more details.     |
| `UserStatusCel`                                                                                                                                           | **string*                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                        | The userStatusCel field.<br/>This field is part of the `user_status` oneof.<br/>See the documentation for `c1.api.workflows.v1beta.UpdateUser` for more details. |
| `UserStatusEnum`                                                                                                                                          | [*shared.UserStatusEnum](../../../pkg/models/shared/userstatusenum.md)                                                                                    | :heavy_minus_sign:                                                                                                                                        | The userStatusEnum field.<br/>This field is part of the `user_status` oneof.<br/>See the documentation for `c1.api.workflows.v1beta.UpdateUser` for more details. |