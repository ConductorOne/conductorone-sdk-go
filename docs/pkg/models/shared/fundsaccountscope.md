# FundsAccountScope

AccountScope identifies the tenant, user, application, or user-application
 account against which spending was admitted.

This message contains a oneof named kind. Only a single field of the following list may be set at a time:
  - tenant
  - subject
  - app
  - subjectApp



## Fields

| Field                                                                              | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `App`                                                                              | [*shared.FundsAppScope](../../../pkg/models/shared/fundsappscope.md)               | :heavy_minus_sign:                                                                 | N/A                                                                                |
| `Subject`                                                                          | [*shared.FundsSubjectScope](../../../pkg/models/shared/fundssubjectscope.md)       | :heavy_minus_sign:                                                                 | N/A                                                                                |
| `SubjectApp`                                                                       | [*shared.FundsSubjectAppScope](../../../pkg/models/shared/fundssubjectappscope.md) | :heavy_minus_sign:                                                                 | N/A                                                                                |
| `Tenant`                                                                           | [*shared.FundsTenantScope](../../../pkg/models/shared/fundstenantscope.md)         | :heavy_minus_sign:                                                                 | N/A                                                                                |