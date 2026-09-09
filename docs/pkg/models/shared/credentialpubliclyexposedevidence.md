# CredentialPubliclyExposedEvidence

CredentialPubliclyExposedEvidence carries scanner attribution for a public exposure.


## Fields

| Field                                      | Type                                       | Required                                   | Description                                |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| `CredentialRevoked`                        | `*bool`                                    | :heavy_minus_sign:                         | The credentialRevoked field.               |
| `FingerprintPrefix`                        | `*string`                                  | :heavy_minus_sign:                         | The fingerprintPrefix field.               |
| `FirstObservedAt`                          | [*time.Time](https://pkg.go.dev/time#Time) | :heavy_minus_sign:                         | N/A                                        |
| `FirstScannerID`                           | `*string`                                  | :heavy_minus_sign:                         | The firstScannerId field.                  |
| `ReportingScanners`                        | []`string`                                 | :heavy_minus_sign:                         | The reportingScanners field.               |
| `RevokedAt`                                | [*time.Time](https://pkg.go.dev/time#Time) | :heavy_minus_sign:                         | N/A                                        |
| `SourceKind`                               | `*string`                                  | :heavy_minus_sign:                         | The sourceKind field.                      |
| `SourceURL`                                | `*string`                                  | :heavy_minus_sign:                         | The sourceUrl field.                       |