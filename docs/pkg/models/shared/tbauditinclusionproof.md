# TBAuditInclusionProof

TBAuditInclusionProof proves one audit row belongs to a signed batch. The
 service checks the path against the manifest root before returning it. It
 carries no row content: auditors verify the row against leaf_hash offline
 with `tb audit verify-proof`.


## Fields

| Field                                                                                          | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `LeafHash`                                                                                     | `*string`                                                                                      | :heavy_minus_sign:                                                                             | The row's stored RFC 6962 leaf hash, hex.                                                      |
| `LeafIndex`                                                                                    | `*string`                                                                                      | :heavy_minus_sign:                                                                             | Index of the row's leaf in its batch.                                                          |
| `Manifest`                                                                                     | [*shared.TBAuditAttestationManifest](../../../pkg/models/shared/tbauditattestationmanifest.md) | :heavy_minus_sign:                                                                             | N/A                                                                                            |
| `Path`                                                                                         | []`string`                                                                                     | :heavy_minus_sign:                                                                             | RFC 6962 audit path, leaf to root, hex.                                                        |