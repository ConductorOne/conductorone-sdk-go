# PaperSecretServiceSearchAuditEventsRequest

PaperSecretServiceSearchAuditEventsRequest searches audit events for a secret
 owned by the calling user. Only the secret creator may query events. Results
 are sanitized to include only time, event type, and actor information.


## Fields

| Field                                                                | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `PageSize`                                                           | **int*                                                               | :heavy_minus_sign:                                                   | Maximum number of results per page (0 uses server default, max 100). |
| `PageToken`                                                          | **string*                                                            | :heavy_minus_sign:                                                   | Pagination token from a previous response's next_page_token.         |
| `VaultID`                                                            | **string*                                                            | :heavy_minus_sign:                                                   | Required. The vault ID of the secret whose audit events to retrieve. |