# LocalUserInvitationServiceCreateRequest

The LocalUserInvitationServiceCreateRequest message.


## Fields

| Field                                                | Type                                                 | Required                                             | Description                                          |
| ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- |
| `DisplayName`                                        | `string`                                             | :heavy_check_mark:                                   | The displayName field.                               |
| `Email`                                              | `string`                                             | :heavy_check_mark:                                   | The email field.                                     |
| `InitialRoleIds`                                     | []`string`                                           | :heavy_minus_sign:                                   | Optional initial role IDs to assign upon acceptance. |
| `JobID`                                              | `*string`                                            | :heavy_minus_sign:                                   | Optional FK to a ThirdPartyJob.                      |
| `OnboardingFlowID`                                   | `*string`                                            | :heavy_minus_sign:                                   | Optional onboarding flow override.                   |
| `Purpose`                                            | `*string`                                            | :heavy_minus_sign:                                   | Human-readable reason for the invitation.            |
| `SponsorUserID`                                      | `*string`                                            | :heavy_minus_sign:                                   | Optional sponsor User override.                      |