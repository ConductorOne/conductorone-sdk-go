# FunctionsServiceGetLockFileResponse

FunctionsServiceGetLockFileResponse returns the deno lock file content for a commit.


## Fields

| Field                                                       | Type                                                        | Required                                                    | Description                                                 |
| ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- |
| `Content`                                                   | `*string`                                                   | :heavy_minus_sign:                                          | The raw content of the deno lock file (empty if not found). |
| `Exists`                                                    | `*bool`                                                     | :heavy_minus_sign:                                          | Whether the lock file exists for this commit.               |