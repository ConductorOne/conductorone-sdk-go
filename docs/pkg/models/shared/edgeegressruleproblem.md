# EdgeEgressRuleProblem

EdgeEgressRuleProblem is one reason Update would reject a rule list.


## Fields

| Field                                                                                        | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `Kind`                                                                                       | [*shared.EdgeEgressRuleProblemKind](../../../pkg/models/shared/edgeegressruleproblemkind.md) | :heavy_minus_sign:                                                                           | The kind field.                                                                              |
| `Message`                                                                                    | `*string`                                                                                    | :heavy_minus_sign:                                                                           | The same text Update returns for this problem.                                               |
| `RuleIndex`                                                                                  | `*int`                                                                                       | :heavy_minus_sign:                                                                           | Zero-based rule index, or -1 for a problem with the whole list.                              |