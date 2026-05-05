# PolicySteps

A named sequence of steps that execute in order within a policy.


## Fields

| Field                                                                                                                | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `Steps`                                                                                                              | [][shared.PolicyStep](../../../pkg/models/shared/policystep.md)                                                      | :heavy_minus_sign:                                                                                                   | Ordered array of steps. Each step is a oneof -- exactly one step type is<br/> set per entry. Steps execute sequentially. |