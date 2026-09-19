# SetAssignee

SetAssignee sets the finding's assignee. Each arm names WHO, so later arms can
 derive the assignee from the finding's context without a schema break.

 Resolves to exactly one identity user id, or to none. There is deliberately no
 RemoveAssignee kind and no arm naming nobody: a transformation rule never
 clears an assignee, which stays a human action via UpdateFindingAssignee.

This message contains a oneof named typ. Only a single field of the following list may be set at a time:
  - directAssignee



## Fields

| Field                                                                  | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `DirectAssignee`                                                       | [*shared.DirectAssignee](../../../pkg/models/shared/directassignee.md) | :heavy_minus_sign:                                                     | N/A                                                                    |