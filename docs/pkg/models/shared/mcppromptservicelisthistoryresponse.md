# MCPPromptServiceListHistoryResponse

MCPPromptServiceListHistoryResponse contains one page of prompt history.


## Fields

| Field                                                                                 | Type                                                                                  | Required                                                                              | Description                                                                           |
| ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- |
| `List`                                                                                | [][shared.MCPPromptHistoryEntry](../../../pkg/models/shared/mcpprompthistoryentry.md) | :heavy_minus_sign:                                                                    | Snapshots and change metadata, newest first.                                          |
| `NextPageToken`                                                                       | `*string`                                                                             | :heavy_minus_sign:                                                                    | Token for the next page, or empty for the final page.                                 |