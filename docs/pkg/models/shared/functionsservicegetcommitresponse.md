# FunctionsServiceGetCommitResponse

The FunctionsServiceGetCommitResponse message.


## Fields

| Field                                                                  | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `FunctionCommit`                                                       | [*shared.FunctionCommit](../../../pkg/models/shared/functioncommit.md) | :heavy_minus_sign:                                                     | FunctionCommit represents a single commit in a function's history      |
| `Content`                                                              | map[string]*string*                                                    | :heavy_minus_sign:                                                     | The content field.                                                     |