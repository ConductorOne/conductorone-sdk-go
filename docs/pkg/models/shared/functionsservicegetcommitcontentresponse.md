# FunctionsServiceGetCommitContentResponse

FunctionsServiceGetCommitContentResponse contains a commit and all its file contents.


## Fields

| Field                                                                  | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `Commit`                                                               | [*shared.FunctionCommit](../../../pkg/models/shared/functioncommit.md) | :heavy_minus_sign:                                                     | N/A                                                                    |
| `Files`                                                                | map[string]`string`                                                    | :heavy_minus_sign:                                                     | Map of filename to file content bytes.                                 |