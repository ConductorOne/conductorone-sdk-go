# SSOApplicationServiceParseSAMLServiceProviderMetadataResponse

SSOApplicationServiceParseSAMLServiceProviderMetadataResponse returns the
 SAML configuration derived from one metadata document and every finding the
 parser raised about it.


## Fields

| Field                                                                                                              | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `Config`                                                                                                           | [*shared.SSOApplicationSAMLConfig](../../../pkg/models/shared/ssoapplicationsamlconfig.md)                         | :heavy_minus_sign:                                                                                                 | N/A                                                                                                                |
| `Findings`                                                                                                         | [][shared.SAMLMetadataFinding](../../../pkg/models/shared/samlmetadatafinding.md)                                  | :heavy_minus_sign:                                                                                                 | Everything the parser noticed about the document, including requirements<br/> it could not map into the configuration. |