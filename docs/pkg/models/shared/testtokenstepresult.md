# TestTokenStepResult

TestTokenStepResult represents the result of a single validation step.


## Fields

| Field                                                                          | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `Actual`                                                                       | **string*                                                                      | :heavy_minus_sign:                                                             | Actual value from the token.                                                   |
| `Detail`                                                                       | **string*                                                                      | :heavy_minus_sign:                                                             | Human-readable detail message.                                                 |
| `Expected`                                                                     | **string*                                                                      | :heavy_minus_sign:                                                             | Expected value (for comparison steps).                                         |
| `Passed`                                                                       | **bool*                                                                        | :heavy_minus_sign:                                                             | Whether this step passed.                                                      |
| `Skipped`                                                                      | **bool*                                                                        | :heavy_minus_sign:                                                             | Whether this step was skipped (e.g., CIDR check when no allowlist configured). |
| `StepName`                                                                     | **string*                                                                      | :heavy_minus_sign:                                                             | Step name for display (e.g., "JWT decode", "Issuer match").                    |