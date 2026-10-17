# CredentialExpiringType

CredentialExpiringType: a credential is inside the detector's expiry
 warning window, or already past it. Covers both ConductorOne-managed
 credentials and connector-synced secret-trait AppResources. Dedup is
 (credential arm, credential_id/target triple). Target: IdentityUserTarget
 for a service-principal credential, or AppResourceTarget for a
 connector-synced secret.

This message contains a oneof named credential. Only a single field of the following list may be set at a time:
  - userClientId
  - appResourceSecret



## Fields

| Field                                                                                                                                                              | Type                                                                                                                                                               | Required                                                                                                                                                           | Description                                                                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `AppResourceSecret`                                                                                                                                                | [*shared.AppResourceTarget](../../../pkg/models/shared/appresourcetarget.md)                                                                                       | :heavy_minus_sign:                                                                                                                                                 | N/A                                                                                                                                                                |
| `CredentialDisplayName`                                                                                                                                            | `*string`                                                                                                                                                          | :heavy_minus_sign:                                                                                                                                                 | The credentialDisplayName field.                                                                                                                                   |
| `UserClientID`                                                                                                                                                     | `*string`                                                                                                                                                          | :heavy_minus_sign:                                                                                                                                                 | Service-principal credential.<br/>This field is part of the `credential` oneof.<br/>See the documentation for `c1.api.finding.v1.CredentialExpiringType` for more details. |