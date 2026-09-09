# TBControlPlaneServiceGetEgressPolicyResponse

The TBControlPlaneServiceGetEgressPolicyResponse message.


## Fields

| Field                                                                                        | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `CompiledPolicyYaml`                                                                         | `*string`                                                                                    | :heavy_minus_sign:                                                                           | The compiled TB policy YAML for the current generation, empty when<br/> `policy.rules` is empty. |
| `Policy`                                                                                     | [*shared.TBEgressPolicy](../../../pkg/models/shared/tbegresspolicy.md)                       | :heavy_minus_sign:                                                                           | N/A                                                                                          |