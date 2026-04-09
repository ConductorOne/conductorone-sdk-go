# LocalDirectoryConfigServiceCreateRequest

The LocalDirectoryConfigServiceCreateRequest message.


## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `AllowSelfRegistration`                                            | `*bool`                                                            | :heavy_minus_sign:                                                 | The allowSelfRegistration field.                                   |
| `AppID`                                                            | `string`                                                           | :heavy_check_mark:                                                 | FK to the existing App that will back this local directory.        |
| `DefaultProfileTypeID`                                             | `*string`                                                          | :heavy_minus_sign:                                                 | The defaultProfileTypeId field.                                    |
| `DisplayName`                                                      | `string`                                                           | :heavy_check_mark:                                                 | The displayName field.                                             |
| `InvitationTTL`                                                    | `*string`                                                          | :heavy_minus_sign:                                                 | N/A                                                                |
| `IsDefault`                                                        | `*bool`                                                            | :heavy_minus_sign:                                                 | Whether this should be the default local directory for the tenant. |
| `OnboardingFlowID`                                                 | `*string`                                                          | :heavy_minus_sign:                                                 | The onboardingFlowId field.                                        |
| `OrganizationID`                                                   | `*string`                                                          | :heavy_minus_sign:                                                 | Optional FK to a ThirdPartyOrganization.                           |
| `SelfRegistrationDomains`                                          | []`string`                                                         | :heavy_minus_sign:                                                 | The selfRegistrationDomains field.                                 |