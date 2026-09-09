# SSOApplicationServiceParseSAMLServiceProviderMetadataRequest

SSOApplicationServiceParseSAMLServiceProviderMetadataRequest carries one
 SAML service-provider metadata document to parse.


## Fields

| Field                                                                                                                                            | Type                                                                                                                                             | Required                                                                                                                                         | Description                                                                                                                                      |
| ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ |
| `MetadataXML`                                                                                                                                    | `string`                                                                                                                                         | :heavy_check_mark:                                                                                                                               | The SP metadata XML document, exactly as downloaded or exported from the<br/> service provider. Maximum 1 MiB. The document is parsed, never stored. |