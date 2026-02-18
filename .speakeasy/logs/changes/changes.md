## Go SDK Changes:
* `ConductoroneApi.AutomationSearch.SearchAutomationTemplateVersions()`: `response.List[]` **Changed** (Breaking ⚠️)
    - `AutomationSteps[].GrantEntitlements.AppEntitlementRefsCel` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].GrantEntitlements.AppEntitlementRefs` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionCriteria` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionListCel` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionList` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionNone` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementInclusionCriteria` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementInclusionListCel` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementInclusionList` **Added**
    - `Triggers[].FormTrigger` **Removed** (Breaking ⚠️)
    - `Triggers[].GrantDeletedTrigger.GrantTriggerFilter.EntitlementFilter` **Removed** (Breaking ⚠️)
    - `Triggers[].ManualAutomationTrigger` **Removed** (Breaking ⚠️)
* `ConductoroneApi.FunctionsSearch.Search()`: `response.List[]` **Changed** (Breaking ⚠️)
    - `EncryptedValues` **Removed** (Breaking ⚠️)
    - `Secret` **Added**
* `ConductoroneApi.Functions.UpdateFunction()`: 
  * `request.Request.Function` **Changed** (Breaking ⚠️)
    - `EncryptedValues` **Removed** (Breaking ⚠️)
    - `Secret` **Added**
  * `response.Function` **Changed** (Breaking ⚠️)
    - `EncryptedValues` **Removed** (Breaking ⚠️)
    - `Secret` **Added**
* `ConductoroneApi.Functions.ListFunctions()`: `response.List[]` **Changed** (Breaking ⚠️)
    - `EncryptedValues` **Removed** (Breaking ⚠️)
    - `Secret` **Added**
* `ConductoroneApi.Functions.GetFunction()`: `response.Function` **Changed** (Breaking ⚠️)
    - `EncryptedValues` **Removed** (Breaking ⚠️)
    - `Secret` **Added**
* `ConductoroneApi.Functions.CreateFunction()`: `response.Function` **Changed** (Breaking ⚠️)
    - `EncryptedValues` **Removed** (Breaking ⚠️)
    - `Secret` **Added**
* `ConductoroneApi.Automation.UpdateAutomation()`: 
  * `request.Request.UpdateAutomationRequest.Automation` **Changed** (Breaking ⚠️)
    - `AutomationSteps[].GrantEntitlements.AppEntitlementRefsCel` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].GrantEntitlements.AppEntitlementRefs` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionCriteria` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionListCel` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionList` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionNone` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementInclusionCriteria` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementInclusionListCel` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementInclusionList` **Added**
    - `DraftTriggers[].FormTrigger` **Removed** (Breaking ⚠️)
    - `DraftTriggers[].GrantDeletedTrigger.GrantTriggerFilter.EntitlementFilter` **Removed** (Breaking ⚠️)
    - `DraftTriggers[].ManualAutomationTrigger` **Removed** (Breaking ⚠️)
  * `response.Automation` **Changed** (Breaking ⚠️)
    - `AutomationSteps[].GrantEntitlements.AppEntitlementRefsCel` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].GrantEntitlements.AppEntitlementRefs` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionCriteria` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionListCel` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionList` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionNone` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementInclusionCriteria` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementInclusionListCel` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementInclusionList` **Added**
    - `DraftTriggers[].FormTrigger` **Removed** (Breaking ⚠️)
    - `DraftTriggers[].GrantDeletedTrigger.GrantTriggerFilter.EntitlementFilter` **Removed** (Breaking ⚠️)
    - `DraftTriggers[].ManualAutomationTrigger` **Removed** (Breaking ⚠️)
* `ConductoroneApi.Automation.ListAutomations()`: `response.List[]` **Changed** (Breaking ⚠️)
    - `AutomationSteps[].GrantEntitlements.AppEntitlementRefsCel` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].GrantEntitlements.AppEntitlementRefs` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionCriteria` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionListCel` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionList` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionNone` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementInclusionCriteria` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementInclusionListCel` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementInclusionList` **Added**
    - `DraftTriggers[].FormTrigger` **Removed** (Breaking ⚠️)
    - `DraftTriggers[].GrantDeletedTrigger.GrantTriggerFilter.EntitlementFilter` **Removed** (Breaking ⚠️)
    - `DraftTriggers[].ManualAutomationTrigger` **Removed** (Breaking ⚠️)
* `ConductoroneApi.Automation.GetAutomation()`: `response.Automation` **Changed** (Breaking ⚠️)
    - `AutomationSteps[].GrantEntitlements.AppEntitlementRefsCel` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].GrantEntitlements.AppEntitlementRefs` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionCriteria` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionListCel` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionList` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionNone` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementInclusionCriteria` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementInclusionListCel` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementInclusionList` **Added**
    - `DraftTriggers[].FormTrigger` **Removed** (Breaking ⚠️)
    - `DraftTriggers[].GrantDeletedTrigger.GrantTriggerFilter.EntitlementFilter` **Removed** (Breaking ⚠️)
    - `DraftTriggers[].ManualAutomationTrigger` **Removed** (Breaking ⚠️)
* `ConductoroneApi.Automation.CreateAutomation()`: 
  * `request.Request` **Changed** (Breaking ⚠️)
    - `AutomationSteps[].GrantEntitlements.AppEntitlementRefsCel` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].GrantEntitlements.AppEntitlementRefs` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionCriteria` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionListCel` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionList` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionNone` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementInclusionCriteria` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementInclusionListCel` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementInclusionList` **Added**
    - `DraftTriggers[].FormTrigger` **Removed** (Breaking ⚠️)
    - `DraftTriggers[].GrantDeletedTrigger.GrantTriggerFilter.EntitlementFilter` **Removed** (Breaking ⚠️)
    - `DraftTriggers[].ManualAutomationTrigger` **Removed** (Breaking ⚠️)
  * `response.Automation` **Changed** (Breaking ⚠️)
    - `AutomationSteps[].GrantEntitlements.AppEntitlementRefsCel` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].GrantEntitlements.AppEntitlementRefs` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionCriteria` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionListCel` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionList` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionNone` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementInclusionCriteria` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementInclusionListCel` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementInclusionList` **Added**
    - `DraftTriggers[].FormTrigger` **Removed** (Breaking ⚠️)
    - `DraftTriggers[].GrantDeletedTrigger.GrantTriggerFilter.EntitlementFilter` **Removed** (Breaking ⚠️)
    - `DraftTriggers[].ManualAutomationTrigger` **Removed** (Breaking ⚠️)
* `ConductoroneApi.AutomationSearch.SearchAutomations()`: `response.List[]` **Changed** (Breaking ⚠️)
    - `AutomationSteps[].GrantEntitlements.AppEntitlementRefsCel` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].GrantEntitlements.AppEntitlementRefs` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionCriteria` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionListCel` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionList` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionNone` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementInclusionCriteria` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementInclusionListCel` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementInclusionList` **Added**
    - `DraftTriggers[].FormTrigger` **Removed** (Breaking ⚠️)
    - `DraftTriggers[].GrantDeletedTrigger.GrantTriggerFilter.EntitlementFilter` **Removed** (Breaking ⚠️)
    - `DraftTriggers[].ManualAutomationTrigger` **Removed** (Breaking ⚠️)
* `ConductoroneApi.AppEntitlements.ListUsers()`:  `response.List[].OriginatingTicketId` **Added**
* `ConductoroneApi.PaperSecretAdmin.Get()`: **Added**
* `ConductoroneApi.AppEntitlementSearch.SearchGrants()`: 
  *  `request.Request.Purpose` **Added**
  * `response.List[]` **Changed**
    - `AppEntitlementUserView.OriginatingTicketId` **Added**
    - `AppEntitlementView.ActorObjectPermissions` **Added**
    - `AppEntitlementView.AppEntitlement.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `AppEntitlementView.AppEntitlement.Purpose.Enum(appEntitlementPurposeValueOwnership)` **Added**
* `ConductoroneApi.PaperSecret.GetByShareCode()`: **Added**
* `ConductoroneApi.PaperSecret.GetContent()`: **Added**
* `ConductoroneApi.PaperSecret.Revoke()`: **Added**
* `ConductoroneApi.PaperSecret.SearchAuditEvents()`: **Added**
* `ConductoroneApi.PaperSecret.SearchMySecrets()`: **Added**
* `ConductoroneApi.PaperSecret.SetTextContent()`: **Added**
* `ConductoroneApi.AppEntitlementOwners.List()`:  `response.List[].Origin` **Added**
* `ConductoroneApi.PaperSecretAdmin.Revoke()`: **Added**
* `ConductoroneApi.PaperSecretAdmin.Search()`: **Added**
* `ConductoroneApi.PaperSecretAdmin.SearchAuditEvents()`: **Added**
* `ConductoroneApi.WorkloadFederation.CreateProvider()`: **Added**
* `ConductoroneApi.WorkloadFederation.CreateTrust()`: **Added**
* `ConductoroneApi.WorkloadFederation.DeleteProvider()`: **Added**
* `ConductoroneApi.WorkloadFederation.DeleteTrust()`: **Added**
* `ConductoroneApi.WorkloadFederation.GetProvider()`: **Added**
* `ConductoroneApi.WorkloadFederation.GetTrust()`: **Added**
* `ConductoroneApi.WorkloadFederation.ListProviders()`: **Added**
* `ConductoroneApi.WorkloadFederation.ListTrusts()`: **Added**
* `ConductoroneApi.WorkloadFederation.SearchTrusts()`: **Added**
* `ConductoroneApi.WorkloadFederation.TestCel()`: **Added**
* `ConductoroneApi.WorkloadFederation.TestToken()`: **Added**
* `ConductoroneApi.WorkloadFederation.UpdateProvider()`: **Added**
* `ConductoroneApi.WorkloadFederation.UpdateTrust()`: **Added**
* `ConductoroneApi.AppOwners.List()`:  `response.List[].Origin` **Added**
* `ConductoroneApi.Principal.CreateCredential()`: **Added**
* `ConductoroneApi.Principal.Delete()`: **Added**
* `ConductoroneApi.Principal.Get()`: **Added**
* `ConductoroneApi.Principal.GetCredential()`: **Added**
* `ConductoroneApi.Principal.List()`: **Added**
* `ConductoroneApi.Principal.ListCredentials()`: **Added**
* `ConductoroneApi.Principal.RevokeCredential()`: **Added**
* `ConductoroneApi.Principal.Update()`: **Added**
* `ConductoroneApi.Principal.UpdateCredential()`: **Added**
* `ConductoroneApi.OrgNotificationSettings.Get()`: **Added**
* `ConductoroneApi.OrgNotificationSettings.Update()`: **Added**
* `ConductoroneApi.UserNotificationSettings.Get()`: **Added**
* `ConductoroneApi.UserNotificationSettings.Update()`: **Added**
* `ConductoroneApi.Functions.GetFunctionSecretEncryptionKey()`: **Removed** (Breaking ⚠️)
* `ConductoroneApi.AccessReview.Create()`: 
  * `request.Request` **Changed**
    - `AccessReviewScopeV2` **Added**
    - `NotificationConfig.SendKickoff` **Added**
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
* `ConductoroneApi.AccessReviewTemplate.Create()`: 
  * `request.Request` **Changed**
    - `AccessReviewScopeV2` **Added**
    - `NotificationConfig.SendKickoff` **Added**
  * `response.AccessReviewTemplate` **Changed**
    - `AccuracyIssueAction` **Added**
    - `AutoCloseCampaign` **Added**
    - `AutoCloseDecision` **Added**
    - `AutoStartCampaign` **Added**
    - `NotificationConfig.SendKickoff` **Added**
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
* `ConductoroneApi.Apps.Create()`: `response.App` **Changed**
    - `AppOwners[].Origin` **Added**
    - `EnableConnectorSourcedOwnership` **Added**
* `ConductoroneApi.Apps.Get()`: `response.App` **Changed**
    - `AppOwners[].Origin` **Added**
    - `EnableConnectorSourcedOwnership` **Added**
* `ConductoroneApi.Apps.List()`: `response.List[]` **Changed**
    - `AppOwners[].Origin` **Added**
    - `EnableConnectorSourcedOwnership` **Added**
* `ConductoroneApi.Apps.Update()`: 
  *  `request.Request.UpdateAppRequest.App.EnableConnectorSourcedOwnership` **Added**
  * `response.App` **Changed**
    - `AppOwners[].Origin` **Added**
    - `EnableConnectorSourcedOwnership` **Added**
* `ConductoroneApi.Connector.Create()`:  `response.ConnectorView.Connector.ConnectorSyncCronSchedule` **Added**
* `ConductoroneApi.Connector.CreateDelegated()`:  `response.ConnectorView.Connector.ConnectorSyncCronSchedule` **Added**
* `ConductoroneApi.Connector.Get()`:  `response.ConnectorView.Connector.ConnectorSyncCronSchedule` **Added**
* `ConductoroneApi.Connector.List()`:  `response.List[].Connector.ConnectorSyncCronSchedule` **Added**
* `ConductoroneApi.Connector.Update()`: 
  *  `request.Request.ConnectorServiceUpdateRequest.Connector.ConnectorSyncCronSchedule` **Added**
  *  `response.ConnectorView.Connector.ConnectorSyncCronSchedule` **Added**
* `ConductoroneApi.Connector.UpdateDelegated()`: 
  *  `request.Request.ConnectorServiceUpdateDelegatedRequest.Connector.ConnectorSyncCronSchedule` **Added**
  *  `response.ConnectorView.Connector.ConnectorSyncCronSchedule` **Added**
* `ConductoroneApi.AppAccessRequestsDefaults.CancelAppAccessRequestsDefaults()`:  `response.RequestSchemaId` **Added**
* `ConductoroneApi.AppAccessRequestsDefaults.CreateAppAccessRequestsDefaults()`: 
  *  `request.Request.AppAccessRequestDefaults.RequestSchemaId` **Added**
  *  `response.RequestSchemaId` **Added**
* `ConductoroneApi.AppAccessRequestsDefaults.GetAppAccessRequestsDefaults()`:  `response.RequestSchemaId` **Added**
* `ConductoroneApi.AppEntitlements.Create()`: 
  * `request.Request.CreateAppEntitlementRequest` **Changed**
    - `ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Purpose.Enum(appEntitlementPurposeValueOwnership)` **Added**
  * `response.AppEntitlementView` **Changed**
    - `ActorObjectPermissions` **Added**
    - `AppEntitlement.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `AppEntitlement.Purpose.Enum(appEntitlementPurposeValueOwnership)` **Added**
* `ConductoroneApi.AppEntitlements.Get()`: `response.AppEntitlementView` **Changed**
    - `ActorObjectPermissions` **Added**
    - `AppEntitlement.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `AppEntitlement.Purpose.Enum(appEntitlementPurposeValueOwnership)` **Added**
* `ConductoroneApi.AppEntitlements.List()`: `response.List[]` **Changed**
    - `ActorObjectPermissions` **Added**
    - `AppEntitlement.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `AppEntitlement.Purpose.Enum(appEntitlementPurposeValueOwnership)` **Added**
* `ConductoroneApi.AppEntitlements.ListAutomationExclusions()`:  `response.List[].User.Origin` **Added**
* `ConductoroneApi.AppEntitlements.ListForAppResource()`: `response.List[]` **Changed**
    - `ActorObjectPermissions` **Added**
    - `AppEntitlement.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `AppEntitlement.Purpose.Enum(appEntitlementPurposeValueOwnership)` **Added**
* `ConductoroneApi.AppEntitlements.ListForAppUser()`: `response.List[]` **Changed**
    - `ActorObjectPermissions` **Added**
    - `AppEntitlement.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `AppEntitlement.Purpose.Enum(appEntitlementPurposeValueOwnership)` **Added**
* `ConductoroneApi.A2Ui.CreateSurfaceFeedback()`: **Added**
* `ConductoroneApi.AppEntitlements.Update()`: 
  * `request.Request.UpdateAppEntitlementRequest.AppEntitlement` **Changed**
    - `ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Purpose.Enum(appEntitlementPurposeValueOwnership)` **Added**
  * `response.AppEntitlementView` **Changed**
    - `ActorObjectPermissions` **Added**
    - `AppEntitlement.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `AppEntitlement.Purpose.Enum(appEntitlementPurposeValueOwnership)` **Added**
* `ConductoroneApi.AppEntitlementSearch.Search()`: `response.List[]` **Changed**
    - `ActorObjectPermissions` **Added**
    - `AppEntitlement.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `AppEntitlement.Purpose.Enum(appEntitlementPurposeValueOwnership)` **Added**
* `ConductoroneApi.AppEntitlementSearch.SearchAppEntitlementsForAppUser()`: `response.List[]` **Changed**
    - `ActorObjectPermissions` **Added**
    - `AppEntitlement.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `AppEntitlement.Purpose.Enum(appEntitlementPurposeValueOwnership)` **Added**
* `ConductoroneApi.AppEntitlementSearch.SearchAppEntitlementsWithExpired()`:  `response.List[].User.Origin` **Added**
* `ConductoroneApi.PaperSecret.Get()`: **Added**
* `ConductoroneApi.PaperSecret.CreateInternal()`: **Added**
* `ConductoroneApi.Principal.Create()`: **Added**
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
* `ConductoroneApi.AppResourceOwners.List()`: `response` **Changed**
    - `ImmutableUserIds` **Added**
    - `List[].Origin` **Added**
* `ConductoroneApi.PaperSecret.CreateExternal()`: **Added**
* `ConductoroneApi.ExternalClientSearch.Search()`: **Added**
* `ConductoroneApi.FunctionsInvocationSearch.Search()`: **Added**
* `ConductoroneApi.Functions.Test()`: **Added**
* `ConductoroneApi.AccessReviewTemplateSetupEntitlement.SetScopeByResourceType()`: **Added**
* `ConductoroneApi.AccessReviewTemplateSetupEntitlement.SetScopeAndEntitlements()`: **Added**
* `ConductoroneApi.RequestCatalogManagement.Create()`: `response.RequestCatalogView.RequestCatalog.AccessEntitlements[]` **Changed**
    - `ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Purpose.Enum(appEntitlementPurposeValueOwnership)` **Added**
* `ConductoroneApi.RequestCatalogManagement.Get()`: `response.RequestCatalogView.RequestCatalog.AccessEntitlements[]` **Changed**
    - `ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Purpose.Enum(appEntitlementPurposeValueOwnership)` **Added**
* `ConductoroneApi.RequestCatalogManagement.List()`: `response.List[].RequestCatalog.AccessEntitlements[]` **Changed**
    - `ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Purpose.Enum(appEntitlementPurposeValueOwnership)` **Added**
* `ConductoroneApi.RequestCatalogManagement.ListEntitlementsForAccess()`: `response.List[]` **Changed**
    - `ActorObjectPermissions` **Added**
    - `AppEntitlement.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `AppEntitlement.Purpose.Enum(appEntitlementPurposeValueOwnership)` **Added**
* `ConductoroneApi.RequestCatalogManagement.ListEntitlementsPerCatalog()`: `response.List[]` **Changed**
    - `ActorObjectPermissions` **Added**
    - `AppEntitlement.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `AppEntitlement.Purpose.Enum(appEntitlementPurposeValueOwnership)` **Added**
* `ConductoroneApi.RequestCatalogManagement.Update()`: 
  * `request.Request.RequestCatalogManagementServiceUpdateRequest.RequestCatalog.AccessEntitlements[]` **Changed**
    - `ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Purpose.Enum(appEntitlementPurposeValueOwnership)` **Added**
  * `response.RequestCatalogView.RequestCatalog.AccessEntitlements[]` **Changed**
    - `ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Purpose.Enum(appEntitlementPurposeValueOwnership)` **Added**
* `ConductoroneApi.ConnectorCatalog.ConfigurationSchema()`: `response` **Changed**
    - `ConfigSchema.Fields[].StringField.PickerField.C1UserFilter` **Added**
    - `Form.FieldRelationships[].DependentOn` **Added**
    - `Form.Fields[].Required` **Added**
* `ConductoroneApi.AccessReviewSetupEntitlement.SetCampaignScopeByResourceType()`: **Added**
* `ConductoroneApi.AccessReviewSetupEntitlement.SetCampaignScopeAndEntitlements()`: **Added**
* `ConductoroneApi.A2Ui.SubmitAction()`: **Added**
* `ConductoroneApi.A2Ui.ListSurfaces()`: **Added**
* `ConductoroneApi.Policies.Create()`: 
  * `request.Request.PolicySteps.Map<PolicySteps>.Steps[]` **Changed**
    - `Action.ActionTargetBatonResourceAction` **Added**
    - `Form.Form.FieldRelationships[].DependentOn` **Added**
    - `Form.Form.Fields[].Required` **Added**
    - `Form.Form.Fields[].StringField.PickerField.C1UserFilter` **Added**
    - `Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
  * `response.Policy.PolicySteps.Map<PolicySteps>.Steps[]` **Changed**
    - `Action.ActionTargetBatonResourceAction` **Added**
    - `Form.Form.FieldRelationships[].DependentOn` **Added**
    - `Form.Form.Fields[].Required` **Added**
    - `Form.Form.Fields[].StringField.PickerField.C1UserFilter` **Added**
    - `Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
* `ConductoroneApi.Policies.Get()`: `response.Policy.PolicySteps.Map<PolicySteps>.Steps[]` **Changed**
    - `Action.ActionTargetBatonResourceAction` **Added**
    - `Form.Form.FieldRelationships[].DependentOn` **Added**
    - `Form.Form.Fields[].Required` **Added**
    - `Form.Form.Fields[].StringField.PickerField.C1UserFilter` **Added**
    - `Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
* `ConductoroneApi.Policies.List()`: `response.List[].PolicySteps.Map<PolicySteps>.Steps[]` **Changed**
    - `Action.ActionTargetBatonResourceAction` **Added**
    - `Form.Form.FieldRelationships[].DependentOn` **Added**
    - `Form.Form.Fields[].Required` **Added**
    - `Form.Form.Fields[].StringField.PickerField.C1UserFilter` **Added**
    - `Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
* `ConductoroneApi.Policies.Update()`: 
  * `request.Request.UpdatePolicyRequest.Policy.PolicySteps.Map<PolicySteps>.Steps[]` **Changed**
    - `Action.ActionTargetBatonResourceAction` **Added**
    - `Form.Form.FieldRelationships[].DependentOn` **Added**
    - `Form.Form.Fields[].Required` **Added**
    - `Form.Form.Fields[].StringField.PickerField.C1UserFilter` **Added**
    - `Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
  * `response.Policy.PolicySteps.Map<PolicySteps>.Steps[]` **Changed**
    - `Action.ActionTargetBatonResourceAction` **Added**
    - `Form.Form.FieldRelationships[].DependentOn` **Added**
    - `Form.Form.Fields[].Required` **Added**
    - `Form.Form.Fields[].StringField.PickerField.C1UserFilter` **Added**
    - `Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
* `ConductoroneApi.RequestSchema.Create()`: 
  * `request.Request.Fields[]` **Changed**
    - `Required` **Added**
    - `StringField.PickerField.C1UserFilter` **Added**
  * `response.RequestSchema.Form` **Changed**
    - `FieldRelationships[].DependentOn` **Added**
    - `Fields[].Required` **Added**
    - `Fields[].StringField.PickerField.C1UserFilter` **Added**
* `ConductoroneApi.RequestSchema.Get()`: `response.RequestSchema.Form` **Changed**
    - `FieldRelationships[].DependentOn` **Added**
    - `Fields[].Required` **Added**
    - `Fields[].StringField.PickerField.C1UserFilter` **Added**
* `ConductoroneApi.RequestSchema.Update()`: 
  * `request.Request.RequestSchemaServiceUpdateRequest.RequestSchema.Form` **Changed**
    - `FieldRelationships[].DependentOn` **Added**
    - `Fields[].Required` **Added**
    - `Fields[].StringField.PickerField.C1UserFilter` **Added**
  * `response.RequestSchema.Form` **Changed**
    - `FieldRelationships[].DependentOn` **Added**
    - `Fields[].Required` **Added**
    - `Fields[].StringField.PickerField.C1UserFilter` **Added**
* `ConductoroneApi.AppResourceSearch.SearchAppResources()`: `response.List[]` **Changed**
    - `ActorObjectPermissions` **Added**
    - `AppResource.AccessConfigId` **Added**
    - `AppResource.Profile` **Added**
* `ConductoroneApi.AppSearch.Search()`: `response.List[]` **Changed**
    - `AppOwners[].Origin` **Added**
    - `EnableConnectorSourcedOwnership` **Added**
* `ConductoroneApi.A2Ui.ListSurfaceFeedback()`: **Added**
* `ConductoroneApi.PolicySearch.Search()`: `response.List[].PolicySteps.Map<PolicySteps>.Steps[]` **Changed**
    - `Action.ActionTargetBatonResourceAction` **Added**
    - `Form.Form.FieldRelationships[].DependentOn` **Added**
    - `Form.Form.Fields[].Required` **Added**
    - `Form.Form.Fields[].StringField.PickerField.C1UserFilter` **Added**
    - `Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
* `ConductoroneApi.RequestCatalogSearch.SearchEntitlements()`: `response.List[].AppEntitlementView` **Changed**
    - `ActorObjectPermissions` **Added**
    - `AppEntitlement.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `AppEntitlement.Purpose.Enum(appEntitlementPurposeValueOwnership)` **Added**
* `ConductoroneApi.StepUpProvider.Create()`: 
  *  `request.Request.StepUpMicrosoftSettings.ValidationMode` **Added**
  *  `response.StepUpProvider.StepUpMicrosoftSettings.ValidationMode` **Added**
* `ConductoroneApi.StepUpProvider.Get()`:  `response.StepUpProvider.StepUpMicrosoftSettings.ValidationMode` **Added**
* `ConductoroneApi.StepUpProvider.List()`:  `response.List[].StepUpMicrosoftSettings.ValidationMode` **Added**
* `ConductoroneApi.StepUpProvider.Search()`:  `response.List[].StepUpMicrosoftSettings.ValidationMode` **Added**
* `ConductoroneApi.StepUpProvider.Update()`: 
  *  `request.Request.UpdateStepUpProviderRequest.StepUpProvider.StepUpMicrosoftSettings.ValidationMode` **Added**
  *  `response.StepUpProvider.StepUpMicrosoftSettings.ValidationMode` **Added**
* `ConductoroneApi.StepUpProvider.UpdateSecret()`:  `response.StepUpProvider.StepUpMicrosoftSettings.ValidationMode` **Added**
* `ConductoroneApi.TaskSearch.Search()`: 
  * `request.Request` **Changed**
    - `ExcludeApplicationIds` **Added**
    - `PendingActionFilter` **Added**
    - `TaskTypes[].TaskTypeGrant.TaskGrantSource.IsExtension` **Added**
  * `response.List[]` **Changed**
    - `ApproversPath` **Added**
    - `Task.ApproverIds` **Added**
    - `Task.Form.FieldRelationships[].DependentOn` **Added**
    - `Task.Form.Fields[].Required` **Added**
    - `Task.Form.Fields[].StringField.PickerField.C1UserFilter` **Added**
    - `Task.Origin.Enum(taskOriginCascadeDelete)` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Action.ActionTargetBatonResourceAction` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ActionInstance.ActionTargetBatonResourceActionInstance` **Added**
    - `Task.TaskType.TaskTypeGrant.TaskGrantSource.IsExtension` **Added**
* `ConductoroneApi.UserSearch.Search()`: 
  * `request.Request` **Changed**
    - `ExcludeOrigins` **Added**
    - `Origins` **Added**
  *  `response.List[].User.Origin` **Added**
* `ConductoroneApi.WebhooksSearch.Search()`:  `response.List[].CallbackTimeout` **Added**
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
* `ConductoroneApi.SystemLog.ListEvents()`: 
  *  `request.Request.UntilEventUid` **Added**
* `ConductoroneApi.TaskAudit.List()`: `response.List[]` **Changed**
    - `TaskAuditActionInstanceCreated.ActionInstance.Action.ActionTargetBatonResourceAction` **Added**
    - `TaskAuditActionInstanceCreated.ActionInstance.ActionTargetBatonResourceActionInstance` **Added**
    - `TaskAuditCreatedReplacementExtensionGrantTask` **Added**
    - `TaskAuditMetaData.User.Origin` **Added**
    - `TaskAuditNewTaskCreatedFrom` **Added**
    - `TaskAuditReassignmentFallbackToAdmin` **Added**
* `ConductoroneApi.Task.CreateGrantTask()`: 
  *  `request.Request.TaskGrantSource.IsExtension` **Added**
  * `response.TaskView` **Changed**
    - `ApproversPath` **Added**
    - `Task.ApproverIds` **Added**
    - `Task.Form.FieldRelationships[].DependentOn` **Added**
    - `Task.Form.Fields[].Required` **Added**
    - `Task.Form.Fields[].StringField.PickerField.C1UserFilter` **Added**
    - `Task.Origin.Enum(taskOriginCascadeDelete)` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Action.ActionTargetBatonResourceAction` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ActionInstance.ActionTargetBatonResourceActionInstance` **Added**
    - `Task.TaskType.TaskTypeGrant.TaskGrantSource.IsExtension` **Added**
* `ConductoroneApi.Task.CreateOffboardingTask()`: `response.TaskView` **Changed**
    - `ApproversPath` **Added**
    - `Task.ApproverIds` **Added**
    - `Task.Form.FieldRelationships[].DependentOn` **Added**
    - `Task.Form.Fields[].Required` **Added**
    - `Task.Form.Fields[].StringField.PickerField.C1UserFilter` **Added**
    - `Task.Origin.Enum(taskOriginCascadeDelete)` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Action.ActionTargetBatonResourceAction` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ActionInstance.ActionTargetBatonResourceActionInstance` **Added**
    - `Task.TaskType.TaskTypeGrant.TaskGrantSource.IsExtension` **Added**
* `ConductoroneApi.Task.CreateRevokeTask()`: `response.TaskView` **Changed**
    - `ApproversPath` **Added**
    - `Task.ApproverIds` **Added**
    - `Task.Form.FieldRelationships[].DependentOn` **Added**
    - `Task.Form.Fields[].Required` **Added**
    - `Task.Form.Fields[].StringField.PickerField.C1UserFilter` **Added**
    - `Task.Origin.Enum(taskOriginCascadeDelete)` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Action.ActionTargetBatonResourceAction` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ActionInstance.ActionTargetBatonResourceActionInstance` **Added**
    - `Task.TaskType.TaskTypeGrant.TaskGrantSource.IsExtension` **Added**
* `ConductoroneApi.Task.Get()`: `response.TaskView` **Changed**
    - `ApproversPath` **Added**
    - `Task.ApproverIds` **Added**
    - `Task.Form.FieldRelationships[].DependentOn` **Added**
    - `Task.Form.Fields[].Required` **Added**
    - `Task.Form.Fields[].StringField.PickerField.C1UserFilter` **Added**
    - `Task.Origin.Enum(taskOriginCascadeDelete)` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Action.ActionTargetBatonResourceAction` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ActionInstance.ActionTargetBatonResourceActionInstance` **Added**
    - `Task.TaskType.TaskTypeGrant.TaskGrantSource.IsExtension` **Added**
* `ConductoroneApi.TaskActions.Approve()`: `response.TaskView` **Changed**
    - `ApproversPath` **Added**
    - `Task.ApproverIds` **Added**
    - `Task.Form.FieldRelationships[].DependentOn` **Added**
    - `Task.Form.Fields[].Required` **Added**
    - `Task.Form.Fields[].StringField.PickerField.C1UserFilter` **Added**
    - `Task.Origin.Enum(taskOriginCascadeDelete)` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Action.ActionTargetBatonResourceAction` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ActionInstance.ActionTargetBatonResourceActionInstance` **Added**
    - `Task.TaskType.TaskTypeGrant.TaskGrantSource.IsExtension` **Added**
* `ConductoroneApi.TaskActions.ApproveWithStepUp()`: `response.TaskView` **Changed**
    - `ApproversPath` **Added**
    - `Task.ApproverIds` **Added**
    - `Task.Form.FieldRelationships[].DependentOn` **Added**
    - `Task.Form.Fields[].Required` **Added**
    - `Task.Form.Fields[].StringField.PickerField.C1UserFilter` **Added**
    - `Task.Origin.Enum(taskOriginCascadeDelete)` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Action.ActionTargetBatonResourceAction` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ActionInstance.ActionTargetBatonResourceActionInstance` **Added**
    - `Task.TaskType.TaskTypeGrant.TaskGrantSource.IsExtension` **Added**
* `ConductoroneApi.TaskActions.Close()`: `response.TaskView` **Changed**
    - `ApproversPath` **Added**
    - `Task.ApproverIds` **Added**
    - `Task.Form.FieldRelationships[].DependentOn` **Added**
    - `Task.Form.Fields[].Required` **Added**
    - `Task.Form.Fields[].StringField.PickerField.C1UserFilter` **Added**
    - `Task.Origin.Enum(taskOriginCascadeDelete)` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Action.ActionTargetBatonResourceAction` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ActionInstance.ActionTargetBatonResourceActionInstance` **Added**
    - `Task.TaskType.TaskTypeGrant.TaskGrantSource.IsExtension` **Added**
* `ConductoroneApi.TaskActions.Comment()`: `response.TaskView` **Changed**
    - `ApproversPath` **Added**
    - `Task.ApproverIds` **Added**
    - `Task.Form.FieldRelationships[].DependentOn` **Added**
    - `Task.Form.Fields[].Required` **Added**
    - `Task.Form.Fields[].StringField.PickerField.C1UserFilter` **Added**
    - `Task.Origin.Enum(taskOriginCascadeDelete)` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Action.ActionTargetBatonResourceAction` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ActionInstance.ActionTargetBatonResourceActionInstance` **Added**
    - `Task.TaskType.TaskTypeGrant.TaskGrantSource.IsExtension` **Added**
* `ConductoroneApi.TaskActions.Deny()`: `response.TaskView` **Changed**
    - `ApproversPath` **Added**
    - `Task.ApproverIds` **Added**
    - `Task.Form.FieldRelationships[].DependentOn` **Added**
    - `Task.Form.Fields[].Required` **Added**
    - `Task.Form.Fields[].StringField.PickerField.C1UserFilter` **Added**
    - `Task.Origin.Enum(taskOriginCascadeDelete)` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Action.ActionTargetBatonResourceAction` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ActionInstance.ActionTargetBatonResourceActionInstance` **Added**
    - `Task.TaskType.TaskTypeGrant.TaskGrantSource.IsExtension` **Added**
* `ConductoroneApi.TaskActions.EscalateToEmergencyAccess()`: `response.TaskView` **Changed**
    - `ApproversPath` **Added**
    - `Task.ApproverIds` **Added**
    - `Task.Form.FieldRelationships[].DependentOn` **Added**
    - `Task.Form.Fields[].Required` **Added**
    - `Task.Form.Fields[].StringField.PickerField.C1UserFilter` **Added**
    - `Task.Origin.Enum(taskOriginCascadeDelete)` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Action.ActionTargetBatonResourceAction` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ActionInstance.ActionTargetBatonResourceActionInstance` **Added**
    - `Task.TaskType.TaskTypeGrant.TaskGrantSource.IsExtension` **Added**
* `ConductoroneApi.TaskActions.HardReset()`: `response.TaskView` **Changed**
    - `ApproversPath` **Added**
    - `Task.ApproverIds` **Added**
    - `Task.Form.FieldRelationships[].DependentOn` **Added**
    - `Task.Form.Fields[].Required` **Added**
    - `Task.Form.Fields[].StringField.PickerField.C1UserFilter` **Added**
    - `Task.Origin.Enum(taskOriginCascadeDelete)` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Action.ActionTargetBatonResourceAction` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ActionInstance.ActionTargetBatonResourceActionInstance` **Added**
    - `Task.TaskType.TaskTypeGrant.TaskGrantSource.IsExtension` **Added**
* `ConductoroneApi.TaskActions.ProcessNow()`: `response.TaskView` **Changed**
    - `ApproversPath` **Added**
    - `Task.ApproverIds` **Added**
    - `Task.Form.FieldRelationships[].DependentOn` **Added**
    - `Task.Form.Fields[].Required` **Added**
    - `Task.Form.Fields[].StringField.PickerField.C1UserFilter` **Added**
    - `Task.Origin.Enum(taskOriginCascadeDelete)` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Action.ActionTargetBatonResourceAction` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ActionInstance.ActionTargetBatonResourceActionInstance` **Added**
    - `Task.TaskType.TaskTypeGrant.TaskGrantSource.IsExtension` **Added**
* `ConductoroneApi.TaskActions.Reassign()`: `response.TaskView` **Changed**
    - `ApproversPath` **Added**
    - `Task.ApproverIds` **Added**
    - `Task.Form.FieldRelationships[].DependentOn` **Added**
    - `Task.Form.Fields[].Required` **Added**
    - `Task.Form.Fields[].StringField.PickerField.C1UserFilter` **Added**
    - `Task.Origin.Enum(taskOriginCascadeDelete)` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Action.ActionTargetBatonResourceAction` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ActionInstance.ActionTargetBatonResourceActionInstance` **Added**
    - `Task.TaskType.TaskTypeGrant.TaskGrantSource.IsExtension` **Added**
* `ConductoroneApi.TaskActions.Restart()`: `response.TaskView` **Changed**
    - `ApproversPath` **Added**
    - `Task.ApproverIds` **Added**
    - `Task.Form.FieldRelationships[].DependentOn` **Added**
    - `Task.Form.Fields[].Required` **Added**
    - `Task.Form.Fields[].StringField.PickerField.C1UserFilter` **Added**
    - `Task.Origin.Enum(taskOriginCascadeDelete)` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Action.ActionTargetBatonResourceAction` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ActionInstance.ActionTargetBatonResourceActionInstance` **Added**
    - `Task.TaskType.TaskTypeGrant.TaskGrantSource.IsExtension` **Added**
* `ConductoroneApi.TaskActions.SkipStep()`: `response.TaskView` **Changed**
    - `ApproversPath` **Added**
    - `Task.ApproverIds` **Added**
    - `Task.Form.FieldRelationships[].DependentOn` **Added**
    - `Task.Form.Fields[].Required` **Added**
    - `Task.Form.Fields[].StringField.PickerField.C1UserFilter` **Added**
    - `Task.Origin.Enum(taskOriginCascadeDelete)` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Action.ActionTargetBatonResourceAction` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ActionInstance.ActionTargetBatonResourceActionInstance` **Added**
    - `Task.TaskType.TaskTypeGrant.TaskGrantSource.IsExtension` **Added**
* `ConductoroneApi.TaskActions.UpdateGrantDuration()`: `response.TaskView` **Changed**
    - `ApproversPath` **Added**
    - `Task.ApproverIds` **Added**
    - `Task.Form.FieldRelationships[].DependentOn` **Added**
    - `Task.Form.Fields[].Required` **Added**
    - `Task.Form.Fields[].StringField.PickerField.C1UserFilter` **Added**
    - `Task.Origin.Enum(taskOriginCascadeDelete)` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Action.ActionTargetBatonResourceAction` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ActionInstance.ActionTargetBatonResourceActionInstance` **Added**
    - `Task.TaskType.TaskTypeGrant.TaskGrantSource.IsExtension` **Added**
* `ConductoroneApi.TaskActions.UpdateRequestData()`: `response.TaskView` **Changed**
    - `ApproversPath` **Added**
    - `Task.ApproverIds` **Added**
    - `Task.Form.FieldRelationships[].DependentOn` **Added**
    - `Task.Form.Fields[].Required` **Added**
    - `Task.Form.Fields[].StringField.PickerField.C1UserFilter` **Added**
    - `Task.Origin.Enum(taskOriginCascadeDelete)` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Action.ActionTargetBatonResourceAction` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ActionInstance.ActionTargetBatonResourceActionInstance` **Added**
    - `Task.TaskType.TaskTypeGrant.TaskGrantSource.IsExtension` **Added**
* `ConductoroneApi.User.Get()`:  `response.UserView.User.Origin` **Added**
* `ConductoroneApi.User.List()`:  `response.List[].User.Origin` **Added**
* `ConductoroneApi.Webhooks.Create()`: 
  *  `request.Request.CallbackTimeout` **Added**
  *  `response.Webhook.CallbackTimeout` **Added**
* `ConductoroneApi.Webhooks.Get()`:  `response.Webhook.CallbackTimeout` **Added**
* `ConductoroneApi.Webhooks.List()`:  `response.List[].CallbackTimeout` **Added**
* `ConductoroneApi.Webhooks.Update()`: 
  *  `request.Request.WebhooksServiceUpdateRequest.Webhook.CallbackTimeout` **Added**
  *  `response.Webhook.CallbackTimeout` **Added**
