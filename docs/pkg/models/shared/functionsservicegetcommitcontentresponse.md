# FunctionsServiceGetCommitContentResponse

FunctionsServiceGetCommitContentResponse contains a commit and all its file contents.


## Fields

| Field                                                                  | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `FunctionCommit`                                                       | [*shared.FunctionCommit](../../../pkg/models/shared/functioncommit.md) | :heavy_minus_sign:                                                     | FunctionCommit represents a single commit in a function's history      |
| `Files`                                                                | map[string]`string`                                                    | :heavy_minus_sign:                                                     | Map of filename to file content bytes.                                 |