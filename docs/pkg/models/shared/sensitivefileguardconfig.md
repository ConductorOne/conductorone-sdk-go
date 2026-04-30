# SensitiveFileGuardConfig

SensitiveFileGuardConfig blocks tool calls that reference sensitive file
 paths or directories.


## Fields

| Field                         | Type                          | Required                      | Description                   |
| ----------------------------- | ----------------------------- | ----------------------------- | ----------------------------- |
| `BlockedDirectories`          | []`string`                    | :heavy_minus_sign:            | The blockedDirectories field. |
| `BlockedPatterns`             | []`string`                    | :heavy_minus_sign:            | The blockedPatterns field.    |