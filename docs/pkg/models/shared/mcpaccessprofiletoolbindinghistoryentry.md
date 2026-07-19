# MCPAccessProfileToolBindingHistoryEntry

MCPAccessProfileToolBindingHistoryEntry is a single change-history record
 capturing the tool bindings added or removed in one transaction.


## Fields

| Field                                                                                                                   | Type                                                                                                                    | Required                                                                                                                | Description                                                                                                             |
| ----------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------- |
| `Items`                                                                                                                 | [][shared.MCPAccessProfileToolBindingHistoryItem](../../../pkg/models/shared/mcpaccessprofiletoolbindinghistoryitem.md) | :heavy_minus_sign:                                                                                                      | The bindings added or removed in this change.                                                                           |
| `Metadata`                                                                                                              | [*shared.ListHistoryEntryMetadata](../../../pkg/models/shared/listhistoryentrymetadata.md)                              | :heavy_minus_sign:                                                                                                      | N/A                                                                                                                     |