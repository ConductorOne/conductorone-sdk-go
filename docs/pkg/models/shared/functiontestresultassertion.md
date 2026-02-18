# FunctionTestResultAssertion

A single assertion within a test.


## Fields

| Field                                       | Type                                        | Required                                    | Description                                 |
| ------------------------------------------- | ------------------------------------------- | ------------------------------------------- | ------------------------------------------- |
| `Actual`                                    | **string*                                   | :heavy_minus_sign:                          | The actual value.                           |
| `At`                                        | **string*                                   | :heavy_minus_sign:                          | Source location of the assertion.           |
| `Description`                               | **string*                                   | :heavy_minus_sign:                          | Description of the assertion.               |
| `Expected`                                  | **string*                                   | :heavy_minus_sign:                          | The expected value.                         |
| `Operator`                                  | **string*                                   | :heavy_minus_sign:                          | The comparison operator (e.g., "==", "!="). |
| `Pass`                                      | **bool*                                     | :heavy_minus_sign:                          | Whether the assertion passed.               |