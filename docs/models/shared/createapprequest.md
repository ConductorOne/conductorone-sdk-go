# CreateAppRequest

 The CreateAppRequest message is used to create a new app.



## Fields

| Field                                              | Type                                               | Required                                           | Description                                        |
| -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- |
| `CertifyPolicyID`                                  | **string*                                          | :heavy_minus_sign:                                 |  Creates the app with this certify policy.<br/>    |
| `Description`                                      | **string*                                          | :heavy_minus_sign:                                 |  Creates the app with this description.<br/>       |
| `DisplayName`                                      | **string*                                          | :heavy_minus_sign:                                 |  Creates the app with this display name.<br/>      |
| `GrantPolicyID`                                    | **string*                                          | :heavy_minus_sign:                                 |  Creates the app with this grant policy.<br/>      |
| `MonthlyCostUsd`                                   | **float64*                                         | :heavy_minus_sign:                                 |  Creates the app with this monthly cost per seat.<br/> |
| `Owners`                                           | []*string*                                         | :heavy_minus_sign:                                 |  Creates the app with this array of owners.<br/>   |
| `RevokePolicyID`                                   | **string*                                          | :heavy_minus_sign:                                 |  Creates the app with this revoke policy.<br/>     |