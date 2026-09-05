# CredentialExpiringType

CredentialExpiringType: a ConductorOne-managed credential is inside the
 detector's expiry warning window, or already past it. Dedup is
 (credential arm, credential_id). Target: IdentityUserTarget -- the identity
 holding the credential.

This message contains a oneof named credential. Only a single field of the following list may be set at a time:
  - userClientId



## Fields

| Field                                                                                                                                                              | Type                                                                                                                                                               | Required                                                                                                                                                           | Description                                                                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `CredentialDisplayName`                                                                                                                                            | `*string`                                                                                                                                                          | :heavy_minus_sign:                                                                                                                                                 | The credentialDisplayName field.                                                                                                                                   |
| `UserClientID`                                                                                                                                                     | `*string`                                                                                                                                                          | :heavy_minus_sign:                                                                                                                                                 | Service-principal credential.<br/>This field is part of the `credential` oneof.<br/>See the documentation for `c1.api.finding.v1.CredentialExpiringType` for more details. |