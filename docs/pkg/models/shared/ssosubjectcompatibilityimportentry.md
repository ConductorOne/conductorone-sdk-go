# SSOSubjectCompatibilityImportEntry

SSOSubjectCompatibilityImportEntry is one client-parsed source row.


## Fields

| Field                                                                  | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `Row`                                                                  | `*int`                                                                 | :heavy_minus_sign:                                                     | One-based row number in the source file, including its header.         |
| `Subject`                                                              | `*string`                                                              | :heavy_minus_sign:                                                     | Exact legacy subject. C1 preserves these UTF-8 bytes without trimming. |
| `UserID`                                                               | `*string`                                                              | :heavy_minus_sign:                                                     | ConductorOne user ID resolved by the client before this batch is sent. |