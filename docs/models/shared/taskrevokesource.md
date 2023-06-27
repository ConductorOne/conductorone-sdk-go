# TaskRevokeSource

The TaskRevokeSource message.

This message contains a oneof named origin. Only a single field of the following list may be set at a time:
  - review
  - request
  - expired
  - nonUsage



## Fields

| Field                                                                    | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `Origin`                                                                 | [*TaskRevokeSourceOrigin](../../models/shared/taskrevokesourceorigin.md) | :heavy_minus_sign:                                                       | N/A                                                                      |