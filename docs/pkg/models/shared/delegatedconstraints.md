# DelegatedConstraints

DelegatedConstraints controls which third-party sign-in providers are
 accepted as proof of email ownership, and how they are scoped.


## Fields

| Field                                                                          | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `GoogleEnabled`                                                                | `*bool`                                                                        | :heavy_minus_sign:                                                             | Accept "Sign in with Google".                                                  |
| `GoogleHostedDomains`                                                          | []`string`                                                                     | :heavy_minus_sign:                                                             | Restrict Google sign-in to these Google Workspace domains. Empty = any domain. |
| `MicrosoftEnabled`                                                             | `*bool`                                                                        | :heavy_minus_sign:                                                             | Accept "Sign in with Microsoft".                                               |
| `MicrosoftTenantIds`                                                           | []`string`                                                                     | :heavy_minus_sign:                                                             | Restrict Microsoft sign-in to these Microsoft tenant IDs. Empty = any tenant.  |