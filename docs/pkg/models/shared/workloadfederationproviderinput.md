# WorkloadFederationProviderInput

WorkloadFederationProvider represents a tenant-level workload identity
 issuer registration. Two issuer schemes are supported:

   - https://...   classic OIDC issuer; `settings.oidc` MUST be set.
   - spiffe://...  SPIFFE trust-domain URI; `settings.spiffe` MUST be set.

 The (well_known_provider, issuer_url scheme, settings oneof) tuple is a
 tri-invariant: SPIFFE wkp ⟺ spiffe:// issuer ⟺ settings.spiffe set; any
 other wkp ⟺ https:// issuer ⟺ settings.oidc set. Issuer URLs are unique
 within tenant.

This message contains a oneof named settings. Only a single field of the following list may be set at a time:
  - oidc
  - spiffe



## Fields

| Field                                                                            | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `Description`                                                                    | `*string`                                                                        | :heavy_minus_sign:                                                               | A description of what this provider is for.                                      |
| `Disabled`                                                                       | `*bool`                                                                          | :heavy_minus_sign:                                                               | Whether the provider is disabled. Disabled providers reject all token exchanges. |
| `DisplayName`                                                                    | `*string`                                                                        | :heavy_minus_sign:                                                               | The display name of the provider.                                                |
| `Oidc`                                                                           | [*shared.OIDCSettings](../../../pkg/models/shared/oidcsettings.md)               | :heavy_minus_sign:                                                               | N/A                                                                              |
| `Spiffe`                                                                         | [*shared.SPIFFESettings](../../../pkg/models/shared/spiffesettings.md)           | :heavy_minus_sign:                                                               | N/A                                                                              |