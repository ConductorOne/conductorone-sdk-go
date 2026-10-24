# EdgeServiceListEgressDenialSummaryResponse

The EdgeServiceListEgressDenialSummaryResponse message.


## Fields

| Field                                                                                | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `Groups`                                                                             | [][shared.TBEgressDenialGroup](../../../pkg/models/shared/tbegressdenialgroup.md)    | :heavy_minus_sign:                                                                   | Ordered by count descending.                                                         |
| `Truncated`                                                                          | `*bool`                                                                              | :heavy_minus_sign:                                                                   | True when more groups matched than were returned; narrow the filters or<br/> the window. |