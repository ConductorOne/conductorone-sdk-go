# SSOSubjectCompatibilityImportIssue

SSOSubjectCompatibilityImportIssue describes one CSV row that cannot be
 imported.


## Fields

| Field                                               | Type                                                | Required                                            | Description                                         |
| --------------------------------------------------- | --------------------------------------------------- | --------------------------------------------------- | --------------------------------------------------- |
| `Reason`                                            | `*string`                                           | :heavy_minus_sign:                                  | Human-readable reason this row cannot be imported.  |
| `Row`                                               | `*int`                                              | :heavy_minus_sign:                                  | One-based CSV row number, including the header row. |
| `Subject`                                           | `*string`                                           | :heavy_minus_sign:                                  | Legacy subject supplied by the batch entry.         |
| `UserID`                                            | `*string`                                           | :heavy_minus_sign:                                  | ConductorOne user ID supplied by the batch entry.   |