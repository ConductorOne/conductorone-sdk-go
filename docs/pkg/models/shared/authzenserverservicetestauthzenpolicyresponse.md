# AuthzenServerServiceTestAuthzenPolicyResponse

AuthzenServerServiceTestAuthzenPolicyResponse returns the evaluation
 outcome for authoring feedback.


## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `Decision`                                                            | `*bool`                                                               | :heavy_minus_sign:                                                    | The policy's decision. Only meaningful when `error` is empty.         |
| `Error`                                                               | `*string`                                                             | :heavy_minus_sign:                                                    | Compile or runtime error, if evaluation could not produce a decision. |
| `Variables`                                                           | map[string]`any`                                                      | :heavy_minus_sign:                                                    | N/A                                                                   |