# C1APITaskV1TaskType

The TaskType message.

This message contains a oneof named task_type. Only a single field of the following list may be set at a time:
  - grant
  - revoke
  - certify



## Fields

| Field                                                                            | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `Certify`                                                                        | [*C1APITaskV1TaskTypeCertify](../../models/shared/c1apitaskv1tasktypecertify.md) | :heavy_minus_sign:                                                               | The TaskTypeCertify message.                                                     |
| `Grant`                                                                          | [*C1APITaskV1TaskTypeGrant](../../models/shared/c1apitaskv1tasktypegrant.md)     | :heavy_minus_sign:                                                               | The TaskTypeGrant message.                                                       |
| `Revoke`                                                                         | [*C1APITaskV1TaskTypeRevoke](../../models/shared/c1apitaskv1tasktyperevoke.md)   | :heavy_minus_sign:                                                               | The TaskTypeRevoke message.                                                      |