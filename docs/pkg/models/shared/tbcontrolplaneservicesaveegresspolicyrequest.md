# TBControlPlaneServiceSaveEgressPolicyRequest

The TBControlPlaneServiceSaveEgressPolicyRequest message.


## Fields

| Field                                                                  | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `DefaultDenyReason`                                                    | `*string`                                                              | :heavy_minus_sign:                                                     | The defaultDenyReason field.                                           |
| `DefaultOutcome`                                                       | [*shared.DefaultOutcome](../../../pkg/models/shared/defaultoutcome.md) | :heavy_minus_sign:                                                     | The defaultOutcome field.                                              |
| `Rules`                                                                | [][shared.TBEgressRule](../../../pkg/models/shared/tbegressrule.md)    | :heavy_minus_sign:                                                     | The rules field.                                                       |
| `TbInstanceID`                                                         | `*string`                                                              | :heavy_minus_sign:                                                     | The tbInstanceId field.                                                |