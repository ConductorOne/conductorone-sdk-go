# AccountScope

AccountScope names one of the four accounting scopes a spend is admitted
 against. The arm is the scope kind — no separate kind enum to keep in step.

 It lives in funds/v1, not beside the ledger row that stores it, because both
 sides need the same declaration: resolution produces a scope, the ledger
 stores it on BudgetAccount, and the account key is derived from it.

This message contains a oneof named kind. Only a single field of the following list may be set at a time:
  - tenant
  - subject
  - app
  - subjectApp



## Fields

| Field                                                                    | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `App`                                                                    | [*shared.AppScope](../../../pkg/models/shared/appscope.md)               | :heavy_minus_sign:                                                       | N/A                                                                      |
| `Subject`                                                                | [*shared.SubjectScope](../../../pkg/models/shared/subjectscope.md)       | :heavy_minus_sign:                                                       | N/A                                                                      |
| `SubjectApp`                                                             | [*shared.SubjectAppScope](../../../pkg/models/shared/subjectappscope.md) | :heavy_minus_sign:                                                       | N/A                                                                      |
| `Tenant`                                                                 | [*shared.TenantScope](../../../pkg/models/shared/tenantscope.md)         | :heavy_minus_sign:                                                       | N/A                                                                      |