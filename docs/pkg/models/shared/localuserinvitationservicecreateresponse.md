# LocalUserInvitationServiceCreateResponse

The LocalUserInvitationServiceCreateResponse message.


## Fields

| Field                                                                                      | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `Invitation`                                                                               | [*shared.LocalUserInvitation](../../../pkg/models/shared/localuserinvitation.md)           | :heavy_minus_sign:                                                                         | N/A                                                                                        |
| `InvitationURL`                                                                            | `*string`                                                                                  | :heavy_minus_sign:                                                                         | Link the invitee opens to accept. Returned only here; C1 stores just a hash of its secret. |