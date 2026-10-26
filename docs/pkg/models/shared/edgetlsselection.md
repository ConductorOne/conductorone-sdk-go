# EdgeTLSSelection

EdgeTlsSelection chooses the TLS interception material for an Edge's
 egress proxy.

This message contains a oneof named selection. Only a single field of the following list may be set at a time:
  - builtin



## Fields

| Field                                                                              | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `Builtin`                                                                          | [*shared.EdgeBuiltinTLSConfig](../../../pkg/models/shared/edgebuiltintlsconfig.md) | :heavy_minus_sign:                                                                 | N/A                                                                                |