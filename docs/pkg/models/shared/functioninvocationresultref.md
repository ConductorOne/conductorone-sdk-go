# FunctionInvocationResultRef

FunctionInvocationResultRef describes an invocation result held outside the
 invocation object. It never carries a storage URI or a download URL.


## Fields

| Field                                                    | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ExpiresAt`                                              | [*time.Time](https://pkg.go.dev/time#Time)               | :heavy_minus_sign:                                       | N/A                                                      |
| `MediaType`                                              | `*string`                                                | :heavy_minus_sign:                                       | The mediaType field.                                     |
| `Path`                                                   | `*string`                                                | :heavy_minus_sign:                                       | The path field.                                          |
| `Sha256`                                                 | `*string`                                                | :heavy_minus_sign:                                       | Base64url SHA-256 of the result bytes.                   |
| `SizeBytes`                                              | `*int64`                                                 | :heavy_minus_sign:                                       | The sizeBytes field.                                     |
| `Storage`                                                | [*shared.Storage](../../../pkg/models/shared/storage.md) | :heavy_minus_sign:                                       | The storage field.                                       |