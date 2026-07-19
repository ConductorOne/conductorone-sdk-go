# AccessReviewScopeV2

The AccessReviewScopeV2 message.

This message contains a oneof named apps_and_resources_scope. Only a single field of the following list may be set at a time:
  - appAccess
  - specificResources
  - appSelectionCriteria
  - resourceTypeSelections


This message contains a oneof named users_scope. Only a single field of the following list may be set at a time:
  - allUsers
  - selectedUsers
  - userCriteria
  - celExpression


This message contains a oneof named accounts_scope. Only a single field of the following list may be set at a time:
  - allAccounts
  - accountCriteria
  - accountCelExpression


This message contains a oneof named grants_scope. Only a single field of the following list may be set at a time:
  - allGrants
  - grantsByCriteria


This message contains a oneof named access_conflicts_scope. Only a single field of the following list may be set at a time:
  - allAccessConflicts
  - specificAccessConflicts


This message contains a oneof named resource_scope. Only a single field of the following list may be set at a time:
  - resourceSelection


This message contains a oneof named excluded_apps_and_resources_scope. Only a single field of the following list may be set at a time:
  - excludedSpecificResources
  - excludedResourceTypeSelections



## Fields

| Field                                                                                              | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `AccountCelExpression`                                                                             | [*shared.CelExpressionScope](../../../pkg/models/shared/celexpressionscope.md)                     | :heavy_minus_sign:                                                                                 | N/A                                                                                                |
| `AccountCriteria`                                                                                  | [*shared.AccountCriteriaScope](../../../pkg/models/shared/accountcriteriascope.md)                 | :heavy_minus_sign:                                                                                 | N/A                                                                                                |
| `AllAccessConflicts`                                                                               | [*shared.AllAccessConflictsScope](../../../pkg/models/shared/allaccessconflictsscope.md)           | :heavy_minus_sign:                                                                                 | N/A                                                                                                |
| `AllAccounts`                                                                                      | [*shared.AllAccountsScope](../../../pkg/models/shared/allaccountsscope.md)                         | :heavy_minus_sign:                                                                                 | N/A                                                                                                |
| `AllGrants`                                                                                        | [*shared.AllGrantsScope](../../../pkg/models/shared/allgrantsscope.md)                             | :heavy_minus_sign:                                                                                 | N/A                                                                                                |
| `AllUsers`                                                                                         | [*shared.AllUsersScope](../../../pkg/models/shared/allusersscope.md)                               | :heavy_minus_sign:                                                                                 | N/A                                                                                                |
| `AppAccess`                                                                                        | [*shared.ApplicationAccessScope](../../../pkg/models/shared/applicationaccessscope.md)             | :heavy_minus_sign:                                                                                 | N/A                                                                                                |
| `AppSelectionCriteria`                                                                             | [*shared.AppSelectionCriteriaScope](../../../pkg/models/shared/appselectioncriteriascope.md)       | :heavy_minus_sign:                                                                                 | N/A                                                                                                |
| `CelExpression`                                                                                    | [*shared.CelExpressionScope](../../../pkg/models/shared/celexpressionscope.md)                     | :heavy_minus_sign:                                                                                 | N/A                                                                                                |
| `ExcludedResourceTypeSelections`                                                                   | [*shared.ResourceTypeSelectionScope](../../../pkg/models/shared/resourcetypeselectionscope.md)     | :heavy_minus_sign:                                                                                 | N/A                                                                                                |
| `ExcludedSpecificResources`                                                                        | [*shared.SpecificResourcesScope](../../../pkg/models/shared/specificresourcesscope.md)             | :heavy_minus_sign:                                                                                 | N/A                                                                                                |
| `GrantsByCriteria`                                                                                 | [*shared.GrantsByCriteriaScope](../../../pkg/models/shared/grantsbycriteriascope.md)               | :heavy_minus_sign:                                                                                 | N/A                                                                                                |
| `PrincipalTypeFilter`                                                                              | [*shared.PrincipalTypeFilter](../../../pkg/models/shared/principaltypefilter.md)                   | :heavy_minus_sign:                                                                                 | Filters principals included in the scope. Unspecified is treated as users.                         |
| `ResourceSelection`                                                                                | [*shared.ResourceSelectionScope](../../../pkg/models/shared/resourceselectionscope.md)             | :heavy_minus_sign:                                                                                 | N/A                                                                                                |
| `ResourceTypeSelections`                                                                           | [*shared.ResourceTypeSelectionScope](../../../pkg/models/shared/resourcetypeselectionscope.md)     | :heavy_minus_sign:                                                                                 | N/A                                                                                                |
| `ScopeRoleSelection`                                                                               | [*shared.ScopeRoleSelectionScope](../../../pkg/models/shared/scoperoleselectionscope.md)           | :heavy_minus_sign:                                                                                 | N/A                                                                                                |
| `SelectedUsers`                                                                                    | [*shared.SelectedUsersScope](../../../pkg/models/shared/selectedusersscope.md)                     | :heavy_minus_sign:                                                                                 | N/A                                                                                                |
| `SpecificAccessConflicts`                                                                          | [*shared.SpecificAccessConflictsScope](../../../pkg/models/shared/specificaccessconflictsscope.md) | :heavy_minus_sign:                                                                                 | N/A                                                                                                |
| `SpecificResources`                                                                                | [*shared.SpecificResourcesScope](../../../pkg/models/shared/specificresourcesscope.md)             | :heavy_minus_sign:                                                                                 | N/A                                                                                                |
| `UserCriteria`                                                                                     | [*shared.UserCriteriaScope](../../../pkg/models/shared/usercriteriascope.md)                       | :heavy_minus_sign:                                                                                 | N/A                                                                                                |