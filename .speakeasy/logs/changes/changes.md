## Go SDK Changes:
* `ConductoroneApi.Functions.CreateFunction()`: `response.Function` **Changed** **Breaking** ⚠️
    - `EncryptedValues` **Removed** **Breaking** ⚠️
    - `Secret` **Added**
* `ConductoroneApi.Functions.GetFunction()`: `response.Function` **Changed** **Breaking** ⚠️
    - `EncryptedValues` **Removed** **Breaking** ⚠️
    - `Secret` **Added**
* `ConductoroneApi.Functions.ListFunctions()`: `response.List[]` **Changed** **Breaking** ⚠️
    - `EncryptedValues` **Removed** **Breaking** ⚠️
    - `Secret` **Added**
* `ConductoroneApi.Functions.UpdateFunction()`: 
  * `request.Request.Function` **Changed** **Breaking** ⚠️
    - `EncryptedValues` **Removed** **Breaking** ⚠️
    - `Secret` **Added**
  * `response.Function` **Changed** **Breaking** ⚠️
    - `EncryptedValues` **Removed** **Breaking** ⚠️
    - `Secret` **Added**
* `ConductoroneApi.FunctionsSearch.Search()`: `response.List[]` **Changed** **Breaking** ⚠️
    - `EncryptedValues` **Removed** **Breaking** ⚠️
    - `Secret` **Added**
* `ConductoroneApi.Automation.CreateAutomation()`: 
  * `request.Request.DraftTriggers[].FormTrigger.Form` **Changed**
    - `FieldRelationships[].DependentOn` **Added**
    - `Fields[].Required` **Added**
  * `response.Automation.DraftTriggers[].FormTrigger.Form` **Changed**
    - `FieldRelationships[].DependentOn` **Added**
    - `Fields[].Required` **Added**
* `ConductoroneApi.RequestCatalogManagement.Update()`: 
  *  `request.Request.RequestCatalogManagementServiceUpdateRequest.RequestCatalog.AccessEntitlements[].Purpose.Enum(appEntitlementPurposeValueOwnership)` **Added**
  *  `response.RequestCatalogView.RequestCatalog.AccessEntitlements[].Purpose.Enum(appEntitlementPurposeValueOwnership)` **Added**
* `ConductoroneApi.AccessReview.Create()`: 
  *  `request.Request.NotificationConfig.SendKickoff` **Added**
  * `response.AccessReviewView.AccessReview` **Changed**
    - `AccuracyIssueAction` **Added**
    - `AutoCloseCampaign` **Added**
    - `AutoCloseDecision` **Added**
    - `AutoStartCampaign` **Added**
    - `NotificationConfig.SendKickoff` **Added**
    - `ScheduledStartDate` **Added**
* `ConductoroneApi.AccessReview.Get()`: `response.AccessReviewView.AccessReview` **Changed**
    - `AccuracyIssueAction` **Added**
    - `AutoCloseCampaign` **Added**
    - `AutoCloseDecision` **Added**
    - `AutoStartCampaign` **Added**
    - `NotificationConfig.SendKickoff` **Added**
    - `ScheduledStartDate` **Added**
* `ConductoroneApi.AccessReview.List()`: `response.List[].AccessReview` **Changed**
    - `AccuracyIssueAction` **Added**
    - `AutoCloseCampaign` **Added**
    - `AutoCloseDecision` **Added**
    - `AutoStartCampaign` **Added**
    - `NotificationConfig.SendKickoff` **Added**
    - `ScheduledStartDate` **Added**
* `ConductoroneApi.AccessReview.Update()`: 
  * `request.Request.AccessReviewServiceUpdateRequest.AccessReview` **Changed**
    - `AccuracyIssueAction` **Added**
    - `AutoCloseCampaign` **Added**
    - `AutoCloseDecision` **Added**
    - `AutoStartCampaign` **Added**
    - `NotificationConfig.SendKickoff` **Added**
    - `ScheduledStartDate` **Added**
  * `response.AccessReviewView.AccessReview` **Changed**
    - `AccuracyIssueAction` **Added**
    - `AutoCloseCampaign` **Added**
    - `AutoCloseDecision` **Added**
    - `AutoStartCampaign` **Added**
    - `NotificationConfig.SendKickoff` **Added**
    - `ScheduledStartDate` **Added**
* `ConductoroneApi.RequestCatalogManagement.Get()`:  `response.RequestCatalogView.RequestCatalog.AccessEntitlements[].Purpose.Enum(appEntitlementPurposeValueOwnership)` **Added**
* `ConductoroneApi.AccessReviewTemplate.Get()`: `response.AccessReviewTemplate` **Changed**
    - `AccuracyIssueAction` **Added**
    - `AutoCloseCampaign` **Added**
    - `AutoCloseDecision` **Added**
    - `AutoStartCampaign` **Added**
    - `NotificationConfig.SendKickoff` **Added**
* `ConductoroneApi.AccessReviewTemplate.Update()`: 
  * `request.Request.AccessReviewTemplateServiceUpdateRequest.AccessReviewTemplate` **Changed**
    - `AccuracyIssueAction` **Added**
    - `AutoCloseCampaign` **Added**
    - `AutoCloseDecision` **Added**
    - `AutoStartCampaign` **Added**
    - `NotificationConfig.SendKickoff` **Added**
  * `response.AccessReviewTemplate` **Changed**
    - `AccuracyIssueAction` **Added**
    - `AutoCloseCampaign` **Added**
    - `AutoCloseDecision` **Added**
    - `AutoStartCampaign` **Added**
    - `NotificationConfig.SendKickoff` **Added**
* `ConductoroneApi.Apps.Create()`:  `response.App.EnableConnectorSourcedOwnership` **Added**
* `ConductoroneApi.Apps.Get()`:  `response.App.EnableConnectorSourcedOwnership` **Added**
* `ConductoroneApi.Apps.List()`:  `response.List[].EnableConnectorSourcedOwnership` **Added**
* `ConductoroneApi.Apps.Update()`: 
  *  `request.Request.UpdateAppRequest.App.EnableConnectorSourcedOwnership` **Added**
  *  `response.App.EnableConnectorSourcedOwnership` **Added**
* `ConductoroneApi.Connector.Create()`:  `response.ConnectorView.Connector.ConnectorSyncCronSchedule` **Added**
* `ConductoroneApi.Automation.UpdateAutomation()`: 
  * `request.Request.UpdateAutomationRequest.Automation.DraftTriggers[].FormTrigger.Form` **Changed**
    - `FieldRelationships[].DependentOn` **Added**
    - `Fields[].Required` **Added**
  * `response.Automation.DraftTriggers[].FormTrigger.Form` **Changed**
    - `FieldRelationships[].DependentOn` **Added**
    - `Fields[].Required` **Added**
* `ConductoroneApi.Connector.Get()`:  `response.ConnectorView.Connector.ConnectorSyncCronSchedule` **Added**
* `ConductoroneApi.Connector.List()`:  `response.List[].Connector.ConnectorSyncCronSchedule` **Added**
* `ConductoroneApi.Connector.Update()`: 
  *  `request.Request.ConnectorServiceUpdateRequest.Connector.ConnectorSyncCronSchedule` **Added**
  *  `response.ConnectorView.Connector.ConnectorSyncCronSchedule` **Added**
* `ConductoroneApi.Connector.UpdateDelegated()`: 
  *  `request.Request.ConnectorServiceUpdateDelegatedRequest.Connector.ConnectorSyncCronSchedule` **Added**
  *  `response.ConnectorView.Connector.ConnectorSyncCronSchedule` **Added**
* `ConductoroneApi.AppEntitlements.Create()`: 
  *  `request.Request.CreateAppEntitlementRequest.Purpose.Enum(appEntitlementPurposeValueOwnership)` **Added**
  * `response.AppEntitlementView` **Changed**
    - `ActorObjectPermissions` **Added**
    - `AppEntitlement.Purpose.Enum(appEntitlementPurposeValueOwnership)` **Added**
* `ConductoroneApi.AppEntitlements.Get()`: `response.AppEntitlementView` **Changed**
    - `ActorObjectPermissions` **Added**
    - `AppEntitlement.Purpose.Enum(appEntitlementPurposeValueOwnership)` **Added**
* `ConductoroneApi.AppEntitlements.List()`: `response.List[]` **Changed**
    - `ActorObjectPermissions` **Added**
    - `AppEntitlement.Purpose.Enum(appEntitlementPurposeValueOwnership)` **Added**
* `ConductoroneApi.AppEntitlements.ListForAppResource()`: `response.List[]` **Changed**
    - `ActorObjectPermissions` **Added**
    - `AppEntitlement.Purpose.Enum(appEntitlementPurposeValueOwnership)` **Added**
* `ConductoroneApi.AppEntitlements.ListForAppUser()`: `response.List[]` **Changed**
    - `ActorObjectPermissions` **Added**
    - `AppEntitlement.Purpose.Enum(appEntitlementPurposeValueOwnership)` **Added**
* `ConductoroneApi.AppEntitlements.Update()`: 
  *  `request.Request.UpdateAppEntitlementRequest.AppEntitlement.Purpose.Enum(appEntitlementPurposeValueOwnership)` **Added**
  * `response.AppEntitlementView` **Changed**
    - `ActorObjectPermissions` **Added**
    - `AppEntitlement.Purpose.Enum(appEntitlementPurposeValueOwnership)` **Added**
* `ConductoroneApi.AppEntitlementSearch.Search()`: `response.List[]` **Changed**
    - `ActorObjectPermissions` **Added**
    - `AppEntitlement.Purpose.Enum(appEntitlementPurposeValueOwnership)` **Added**
* `ConductoroneApi.AppEntitlementSearch.SearchAppEntitlementsForAppUser()`: `response.List[]` **Changed**
    - `ActorObjectPermissions` **Added**
    - `AppEntitlement.Purpose.Enum(appEntitlementPurposeValueOwnership)` **Added**
* `ConductoroneApi.AppEntitlementSearch.SearchGrants()`: 
  *  `request.Request.Purpose` **Added**
  * `response.List[].AppEntitlementView` **Changed**
    - `ActorObjectPermissions` **Added**
    - `AppEntitlement.Purpose.Enum(appEntitlementPurposeValueOwnership)` **Added**
* `ConductoroneApi.AppResource.CreateManuallyManagedAppResource()`: `response.AppResource` **Changed**
    - `AccessConfigId` **Added**
    - `Profile` **Added**
* `ConductoroneApi.AppResource.Get()`: `response.AppResourceView` **Changed**
    - `ActorObjectPermissions` **Added**
    - `AppResource.AccessConfigId` **Added**
    - `AppResource.Profile` **Added**
* `ConductoroneApi.AppResource.List()`: `response.List[]` **Changed**
    - `ActorObjectPermissions` **Added**
    - `AppResource.AccessConfigId` **Added**
    - `AppResource.Profile` **Added**
* `ConductoroneApi.AppResource.Update()`: 
  *  `request.Request.AppResourceServiceUpdateRequest.AppResource.AccessConfigId` **Added**
  * `response.AppResourceView` **Changed**
    - `ActorObjectPermissions` **Added**
    - `AppResource.AccessConfigId` **Added**
    - `AppResource.Profile` **Added**
* `ConductoroneApi.AppResourceOwners.List()`:  `response.ImmutableUserIds` **Added**
* `ConductoroneApi.AutomationSearch.SearchAutomationTemplateVersions()`: `response.List[].Triggers[].FormTrigger.Form` **Changed**
    - `FieldRelationships[].DependentOn` **Added**
    - `Fields[].Required` **Added**
* `ConductoroneApi.AutomationSearch.SearchAutomations()`: `response.List[].DraftTriggers[].FormTrigger.Form` **Changed**
    - `FieldRelationships[].DependentOn` **Added**
    - `Fields[].Required` **Added**
* `ConductoroneApi.ExternalClientSearch.Search()`: **Added**
* `ConductoroneApi.Automation.GetAutomation()`: `response.Automation.DraftTriggers[].FormTrigger.Form` **Changed**
    - `FieldRelationships[].DependentOn` **Added**
    - `Fields[].Required` **Added**
* `ConductoroneApi.Automation.ListAutomations()`: `response.List[].DraftTriggers[].FormTrigger.Form` **Changed**
    - `FieldRelationships[].DependentOn` **Added**
    - `Fields[].Required` **Added**
* `ConductoroneApi.Connector.CreateDelegated()`:  `response.ConnectorView.Connector.ConnectorSyncCronSchedule` **Added**
* `ConductoroneApi.Functions.GetFunctionSecretEncryptionKey()`: **Removed** **Breaking** ⚠️
* `ConductoroneApi.AccessReviewTemplate.Create()`: 
  *  `request.Request.NotificationConfig.SendKickoff` **Added**
  * `response.AccessReviewTemplate` **Changed**
    - `AccuracyIssueAction` **Added**
    - `AutoCloseCampaign` **Added**
    - `AutoCloseDecision` **Added**
    - `AutoStartCampaign` **Added**
    - `NotificationConfig.SendKickoff` **Added**
* `ConductoroneApi.RequestCatalogManagement.List()`:  `response.List[].RequestCatalog.AccessEntitlements[].Purpose.Enum(appEntitlementPurposeValueOwnership)` **Added**
* `ConductoroneApi.RequestCatalogManagement.ListEntitlementsForAccess()`: `response.List[]` **Changed**
    - `ActorObjectPermissions` **Added**
    - `AppEntitlement.Purpose.Enum(appEntitlementPurposeValueOwnership)` **Added**
* `ConductoroneApi.RequestCatalogManagement.ListEntitlementsPerCatalog()`: `response.List[]` **Changed**
    - `ActorObjectPermissions` **Added**
    - `AppEntitlement.Purpose.Enum(appEntitlementPurposeValueOwnership)` **Added**
* `ConductoroneApi.RequestCatalogManagement.Create()`:  `response.RequestCatalogView.RequestCatalog.AccessEntitlements[].Purpose.Enum(appEntitlementPurposeValueOwnership)` **Added**
* `ConductoroneApi.ConnectorCatalog.ConfigurationSchema()`: `response.Form` **Changed**
    - `FieldRelationships[].DependentOn` **Added**
    - `Fields[].Required` **Added**
* `ConductoroneApi.Policies.Create()`: 
  * `request.Request.PolicySteps.Map<PolicySteps>.Steps[]` **Changed**
    - `Action.ActionTargetBatonResourceAction` **Added**
    - `Form.Form.FieldRelationships[].DependentOn` **Added**
    - `Form.Form.Fields[].Required` **Added**
  * `response.Policy.PolicySteps.Map<PolicySteps>.Steps[]` **Changed**
    - `Action.ActionTargetBatonResourceAction` **Added**
    - `Form.Form.FieldRelationships[].DependentOn` **Added**
    - `Form.Form.Fields[].Required` **Added**
* `ConductoroneApi.Policies.Get()`: `response.Policy.PolicySteps.Map<PolicySteps>.Steps[]` **Changed**
    - `Action.ActionTargetBatonResourceAction` **Added**
    - `Form.Form.FieldRelationships[].DependentOn` **Added**
    - `Form.Form.Fields[].Required` **Added**
* `ConductoroneApi.Policies.List()`: `response.List[].PolicySteps.Map<PolicySteps>.Steps[]` **Changed**
    - `Action.ActionTargetBatonResourceAction` **Added**
    - `Form.Form.FieldRelationships[].DependentOn` **Added**
    - `Form.Form.Fields[].Required` **Added**
* `ConductoroneApi.Policies.Update()`: 
  * `request.Request.UpdatePolicyRequest.Policy.PolicySteps.Map<PolicySteps>.Steps[]` **Changed**
    - `Action.ActionTargetBatonResourceAction` **Added**
    - `Form.Form.FieldRelationships[].DependentOn` **Added**
    - `Form.Form.Fields[].Required` **Added**
  * `response.Policy.PolicySteps.Map<PolicySteps>.Steps[]` **Changed**
    - `Action.ActionTargetBatonResourceAction` **Added**
    - `Form.Form.FieldRelationships[].DependentOn` **Added**
    - `Form.Form.Fields[].Required` **Added**
* `ConductoroneApi.RequestSchema.Create()`: 
  *  `request.Request.Fields[].Required` **Added**
  * `response.RequestSchema.Form` **Changed**
    - `FieldRelationships[].DependentOn` **Added**
    - `Fields[].Required` **Added**
* `ConductoroneApi.RequestSchema.Get()`: `response.RequestSchema.Form` **Changed**
    - `FieldRelationships[].DependentOn` **Added**
    - `Fields[].Required` **Added**
* `ConductoroneApi.RequestSchema.Update()`: 
  * `request.Request.RequestSchemaServiceUpdateRequest.RequestSchema.Form` **Changed**
    - `FieldRelationships[].DependentOn` **Added**
    - `Fields[].Required` **Added**
  * `response.RequestSchema.Form` **Changed**
    - `FieldRelationships[].DependentOn` **Added**
    - `Fields[].Required` **Added**
* `ConductoroneApi.AppResourceSearch.SearchAppResources()`: `response.List[]` **Changed**
    - `ActorObjectPermissions` **Added**
    - `AppResource.AccessConfigId` **Added**
    - `AppResource.Profile` **Added**
* `ConductoroneApi.AppSearch.Search()`:  `response.List[].EnableConnectorSourcedOwnership` **Added**
* `ConductoroneApi.PolicySearch.Search()`: `response.List[].PolicySteps.Map<PolicySteps>.Steps[]` **Changed**
    - `Action.ActionTargetBatonResourceAction` **Added**
    - `Form.Form.FieldRelationships[].DependentOn` **Added**
    - `Form.Form.Fields[].Required` **Added**
* `ConductoroneApi.RequestCatalogSearch.SearchEntitlements()`: `response.List[].AppEntitlementView` **Changed**
    - `ActorObjectPermissions` **Added**
    - `AppEntitlement.Purpose.Enum(appEntitlementPurposeValueOwnership)` **Added**
* `ConductoroneApi.TaskSearch.Search()`: 
  * `request.Request` **Changed**
    - `PendingActionFilter` **Added**
    - `TaskTypes[].TaskTypeGrant.TaskGrantSource.IsExtension` **Added**
  * `response.List[].Task` **Changed**
    - `Form.FieldRelationships[].DependentOn` **Added**
    - `Form.Fields[].Required` **Added**
    - `Origin.Enum(taskOriginCascadeDelete)` **Added**
    - `PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Action.ActionTargetBatonResourceAction` **Added**
    - `PolicyInstance.PolicyStepInstance.ActionInstance.ActionTargetBatonResourceActionInstance` **Added**
    - `TaskType.TaskTypeGrant.TaskGrantSource.IsExtension` **Added**
* `ConductoroneApi.SessionSettings.Get()`: `response.SessionSettings` **Changed**
    - `CidrRestriction1` **Changed**
    - `CidrRestriction2` **Changed**
    - `CidrRestriction3` **Changed**
    - `CidrRestriction4` **Changed**
    - `CidrRestriction5` **Added**
    - `ExternalClientsEnabled` **Added**
* `ConductoroneApi.SessionSettings.Update()`: 
  * `request.Request.SessionSettings` **Changed**
    - `CidrRestriction1` **Changed**
    - `CidrRestriction2` **Changed**
    - `CidrRestriction3` **Changed**
    - `CidrRestriction4` **Changed**
    - `CidrRestriction5` **Added**
    - `ExternalClientsEnabled` **Added**
  * `response.SessionSettings` **Changed**
    - `CidrRestriction1` **Changed**
    - `CidrRestriction2` **Changed**
    - `CidrRestriction3` **Changed**
    - `CidrRestriction4` **Changed**
    - `CidrRestriction5` **Added**
    - `ExternalClientsEnabled` **Added**
* `ConductoroneApi.TaskAudit.List()`: `response.List[]` **Changed**
    - `TaskAuditActionInstanceCreated.ActionInstance.Action.ActionTargetBatonResourceAction` **Added**
    - `TaskAuditActionInstanceCreated.ActionInstance.ActionTargetBatonResourceActionInstance` **Added**
    - `TaskAuditCreatedReplacementExtensionGrantTask` **Added**
    - `TaskAuditNewTaskCreatedFrom` **Added**
    - `TaskAuditReassignmentFallbackToAdmin` **Added**
* `ConductoroneApi.Task.CreateGrantTask()`: 
  *  `request.Request.TaskGrantSource.IsExtension` **Added**
  * `response.TaskView.Task` **Changed**
    - `Form.FieldRelationships[].DependentOn` **Added**
    - `Form.Fields[].Required` **Added**
    - `Origin.Enum(taskOriginCascadeDelete)` **Added**
    - `PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Action.ActionTargetBatonResourceAction` **Added**
    - `PolicyInstance.PolicyStepInstance.ActionInstance.ActionTargetBatonResourceActionInstance` **Added**
    - `TaskType.TaskTypeGrant.TaskGrantSource.IsExtension` **Added**
* `ConductoroneApi.Task.CreateOffboardingTask()`: `response.TaskView.Task` **Changed**
    - `Form.FieldRelationships[].DependentOn` **Added**
    - `Form.Fields[].Required` **Added**
    - `Origin.Enum(taskOriginCascadeDelete)` **Added**
    - `PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Action.ActionTargetBatonResourceAction` **Added**
    - `PolicyInstance.PolicyStepInstance.ActionInstance.ActionTargetBatonResourceActionInstance` **Added**
    - `TaskType.TaskTypeGrant.TaskGrantSource.IsExtension` **Added**
* `ConductoroneApi.Task.CreateRevokeTask()`: `response.TaskView.Task` **Changed**
    - `Form.FieldRelationships[].DependentOn` **Added**
    - `Form.Fields[].Required` **Added**
    - `Origin.Enum(taskOriginCascadeDelete)` **Added**
    - `PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Action.ActionTargetBatonResourceAction` **Added**
    - `PolicyInstance.PolicyStepInstance.ActionInstance.ActionTargetBatonResourceActionInstance` **Added**
    - `TaskType.TaskTypeGrant.TaskGrantSource.IsExtension` **Added**
* `ConductoroneApi.Task.Get()`: `response.TaskView.Task` **Changed**
    - `Form.FieldRelationships[].DependentOn` **Added**
    - `Form.Fields[].Required` **Added**
    - `Origin.Enum(taskOriginCascadeDelete)` **Added**
    - `PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Action.ActionTargetBatonResourceAction` **Added**
    - `PolicyInstance.PolicyStepInstance.ActionInstance.ActionTargetBatonResourceActionInstance` **Added**
    - `TaskType.TaskTypeGrant.TaskGrantSource.IsExtension` **Added**
* `ConductoroneApi.TaskActions.Approve()`: `response.TaskView.Task` **Changed**
    - `Form.FieldRelationships[].DependentOn` **Added**
    - `Form.Fields[].Required` **Added**
    - `Origin.Enum(taskOriginCascadeDelete)` **Added**
    - `PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Action.ActionTargetBatonResourceAction` **Added**
    - `PolicyInstance.PolicyStepInstance.ActionInstance.ActionTargetBatonResourceActionInstance` **Added**
    - `TaskType.TaskTypeGrant.TaskGrantSource.IsExtension` **Added**
* `ConductoroneApi.TaskActions.ApproveWithStepUp()`: `response.TaskView.Task` **Changed**
    - `Form.FieldRelationships[].DependentOn` **Added**
    - `Form.Fields[].Required` **Added**
    - `Origin.Enum(taskOriginCascadeDelete)` **Added**
    - `PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Action.ActionTargetBatonResourceAction` **Added**
    - `PolicyInstance.PolicyStepInstance.ActionInstance.ActionTargetBatonResourceActionInstance` **Added**
    - `TaskType.TaskTypeGrant.TaskGrantSource.IsExtension` **Added**
* `ConductoroneApi.TaskActions.Close()`: `response.TaskView.Task` **Changed**
    - `Form.FieldRelationships[].DependentOn` **Added**
    - `Form.Fields[].Required` **Added**
    - `Origin.Enum(taskOriginCascadeDelete)` **Added**
    - `PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Action.ActionTargetBatonResourceAction` **Added**
    - `PolicyInstance.PolicyStepInstance.ActionInstance.ActionTargetBatonResourceActionInstance` **Added**
    - `TaskType.TaskTypeGrant.TaskGrantSource.IsExtension` **Added**
* `ConductoroneApi.TaskActions.Comment()`: `response.TaskView.Task` **Changed**
    - `Form.FieldRelationships[].DependentOn` **Added**
    - `Form.Fields[].Required` **Added**
    - `Origin.Enum(taskOriginCascadeDelete)` **Added**
    - `PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Action.ActionTargetBatonResourceAction` **Added**
    - `PolicyInstance.PolicyStepInstance.ActionInstance.ActionTargetBatonResourceActionInstance` **Added**
    - `TaskType.TaskTypeGrant.TaskGrantSource.IsExtension` **Added**
* `ConductoroneApi.TaskActions.Deny()`: `response.TaskView.Task` **Changed**
    - `Form.FieldRelationships[].DependentOn` **Added**
    - `Form.Fields[].Required` **Added**
    - `Origin.Enum(taskOriginCascadeDelete)` **Added**
    - `PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Action.ActionTargetBatonResourceAction` **Added**
    - `PolicyInstance.PolicyStepInstance.ActionInstance.ActionTargetBatonResourceActionInstance` **Added**
    - `TaskType.TaskTypeGrant.TaskGrantSource.IsExtension` **Added**
* `ConductoroneApi.TaskActions.EscalateToEmergencyAccess()`: `response.TaskView.Task` **Changed**
    - `Form.FieldRelationships[].DependentOn` **Added**
    - `Form.Fields[].Required` **Added**
    - `Origin.Enum(taskOriginCascadeDelete)` **Added**
    - `PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Action.ActionTargetBatonResourceAction` **Added**
    - `PolicyInstance.PolicyStepInstance.ActionInstance.ActionTargetBatonResourceActionInstance` **Added**
    - `TaskType.TaskTypeGrant.TaskGrantSource.IsExtension` **Added**
* `ConductoroneApi.TaskActions.HardReset()`: `response.TaskView.Task` **Changed**
    - `Form.FieldRelationships[].DependentOn` **Added**
    - `Form.Fields[].Required` **Added**
    - `Origin.Enum(taskOriginCascadeDelete)` **Added**
    - `PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Action.ActionTargetBatonResourceAction` **Added**
    - `PolicyInstance.PolicyStepInstance.ActionInstance.ActionTargetBatonResourceActionInstance` **Added**
    - `TaskType.TaskTypeGrant.TaskGrantSource.IsExtension` **Added**
* `ConductoroneApi.TaskActions.ProcessNow()`: `response.TaskView.Task` **Changed**
    - `Form.FieldRelationships[].DependentOn` **Added**
    - `Form.Fields[].Required` **Added**
    - `Origin.Enum(taskOriginCascadeDelete)` **Added**
    - `PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Action.ActionTargetBatonResourceAction` **Added**
    - `PolicyInstance.PolicyStepInstance.ActionInstance.ActionTargetBatonResourceActionInstance` **Added**
    - `TaskType.TaskTypeGrant.TaskGrantSource.IsExtension` **Added**
* `ConductoroneApi.TaskActions.Reassign()`: `response.TaskView.Task` **Changed**
    - `Form.FieldRelationships[].DependentOn` **Added**
    - `Form.Fields[].Required` **Added**
    - `Origin.Enum(taskOriginCascadeDelete)` **Added**
    - `PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Action.ActionTargetBatonResourceAction` **Added**
    - `PolicyInstance.PolicyStepInstance.ActionInstance.ActionTargetBatonResourceActionInstance` **Added**
    - `TaskType.TaskTypeGrant.TaskGrantSource.IsExtension` **Added**
* `ConductoroneApi.TaskActions.Restart()`: `response.TaskView.Task` **Changed**
    - `Form.FieldRelationships[].DependentOn` **Added**
    - `Form.Fields[].Required` **Added**
    - `Origin.Enum(taskOriginCascadeDelete)` **Added**
    - `PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Action.ActionTargetBatonResourceAction` **Added**
    - `PolicyInstance.PolicyStepInstance.ActionInstance.ActionTargetBatonResourceActionInstance` **Added**
    - `TaskType.TaskTypeGrant.TaskGrantSource.IsExtension` **Added**
* `ConductoroneApi.TaskActions.SkipStep()`: `response.TaskView.Task` **Changed**
    - `Form.FieldRelationships[].DependentOn` **Added**
    - `Form.Fields[].Required` **Added**
    - `Origin.Enum(taskOriginCascadeDelete)` **Added**
    - `PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Action.ActionTargetBatonResourceAction` **Added**
    - `PolicyInstance.PolicyStepInstance.ActionInstance.ActionTargetBatonResourceActionInstance` **Added**
    - `TaskType.TaskTypeGrant.TaskGrantSource.IsExtension` **Added**
* `ConductoroneApi.TaskActions.UpdateGrantDuration()`: `response.TaskView.Task` **Changed**
    - `Form.FieldRelationships[].DependentOn` **Added**
    - `Form.Fields[].Required` **Added**
    - `Origin.Enum(taskOriginCascadeDelete)` **Added**
    - `PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Action.ActionTargetBatonResourceAction` **Added**
    - `PolicyInstance.PolicyStepInstance.ActionInstance.ActionTargetBatonResourceActionInstance` **Added**
    - `TaskType.TaskTypeGrant.TaskGrantSource.IsExtension` **Added**
* `ConductoroneApi.TaskActions.UpdateRequestData()`: `response.TaskView.Task` **Changed**
    - `Form.FieldRelationships[].DependentOn` **Added**
    - `Form.Fields[].Required` **Added**
    - `Origin.Enum(taskOriginCascadeDelete)` **Added**
    - `PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Action.ActionTargetBatonResourceAction` **Added**
    - `PolicyInstance.PolicyStepInstance.ActionInstance.ActionTargetBatonResourceActionInstance` **Added**
    - `TaskType.TaskTypeGrant.TaskGrantSource.IsExtension` **Added**
