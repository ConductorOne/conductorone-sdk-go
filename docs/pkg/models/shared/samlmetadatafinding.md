# SAMLMetadataFinding

SAMLMetadataFinding is one thing ConductorOne noticed while parsing a service
 provider's metadata document.


## Fields

| Field                                                        | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `Component`                                                  | [*shared.Component](../../../pkg/models/shared/component.md) | :heavy_minus_sign:                                           | Where the finding fits in the parsed document.               |
| `Level`                                                      | [*shared.Level](../../../pkg/models/shared/level.md)         | :heavy_minus_sign:                                           | The severity of this finding.                                |
| `Reason`                                                     | `*string`                                                    | :heavy_minus_sign:                                           | Plain-language explanation of why the finding was raised.    |