# AuthzenTestRequest

AuthzenTestRequest is one AuthZEN access-evaluation request tuple to test
 a policy against.


## Fields

| Field                                                                            | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `Action`                                                                         | [*shared.AuthzenTestAction](../../../pkg/models/shared/authzentestaction.md)     | :heavy_minus_sign:                                                               | N/A                                                                              |
| `Context`                                                                        | map[string]`any`                                                                 | :heavy_minus_sign:                                                               | N/A                                                                              |
| `Resource`                                                                       | [*shared.AuthzenTestResource](../../../pkg/models/shared/authzentestresource.md) | :heavy_minus_sign:                                                               | N/A                                                                              |
| `Subject`                                                                        | [*shared.AuthzenTestSubject](../../../pkg/models/shared/authzentestsubject.md)   | :heavy_minus_sign:                                                               | N/A                                                                              |