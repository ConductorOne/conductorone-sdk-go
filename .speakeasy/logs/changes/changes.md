## Go SDK Changes:
* `ConductoroneApi.AutomationSearch.SearchAutomationTemplateVersions()`: `response.List[]` **Changed** (Breaking ⚠️)
    - `AutomationSteps[].ConnectorCreateAccount.PasswordCel` **Added**
    - `AutomationSteps[].GeneratePassword.GeneratePasswordPolicy` **Added**
    - `AutomationSteps[].GeneratePassword.PasswordPolicyId` **Added**
    - `AutomationSteps[].GrantEntitlements.AppEntitlementRefsCel` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].GrantEntitlements.AppEntitlementRefs` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionCriteria` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionListCel` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionList` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionNone` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementInclusionCriteria` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementInclusionListCel` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementInclusionList` **Added**
    - `AutomationSteps[].SetCredential` **Added**
    - `AutomationSteps[].StoreCredential` **Added**
    - `Triggers[].FormTrigger` **Removed** (Breaking ⚠️)
    - `Triggers[].GrantDeletedTrigger.GrantTriggerFilter.EntitlementFilter` **Removed** (Breaking ⚠️)
    - `Triggers[].ManualAutomationTrigger` **Removed** (Breaking ⚠️)
    - `Triggers[].WebhookAutomationTrigger.WebhookListenerAuthCapabilityUrl` **Added**
* `ConductoroneApi.AutomationSearch.SearchAutomations()`: `response.List[]` **Changed** (Breaking ⚠️)
    - `AutomationSteps[].ConnectorCreateAccount.PasswordCel` **Added**
    - `AutomationSteps[].GeneratePassword.GeneratePasswordPolicy` **Added**
    - `AutomationSteps[].GeneratePassword.PasswordPolicyId` **Added**
    - `AutomationSteps[].GrantEntitlements.AppEntitlementRefsCel` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].GrantEntitlements.AppEntitlementRefs` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionCriteria` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionListCel` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionList` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionNone` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementInclusionCriteria` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementInclusionListCel` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementInclusionList` **Added**
    - `AutomationSteps[].SetCredential` **Added**
    - `AutomationSteps[].StoreCredential` **Added**
    - `DraftTriggers[].FormTrigger` **Removed** (Breaking ⚠️)
    - `DraftTriggers[].GrantDeletedTrigger.GrantTriggerFilter.EntitlementFilter` **Removed** (Breaking ⚠️)
    - `DraftTriggers[].ManualAutomationTrigger` **Removed** (Breaking ⚠️)
    - `DraftTriggers[].WebhookAutomationTrigger.WebhookListenerAuthCapabilityUrl` **Added**
* `ConductoroneApi.Automation.CreateAutomation()`: 
  * `request.Request` **Changed** (Breaking ⚠️)
    - `AutomationSteps[].ConnectorCreateAccount.PasswordCel` **Added**
    - `AutomationSteps[].GeneratePassword.GeneratePasswordPolicy` **Added**
    - `AutomationSteps[].GeneratePassword.PasswordPolicyId` **Added**
    - `AutomationSteps[].GrantEntitlements.AppEntitlementRefsCel` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].GrantEntitlements.AppEntitlementRefs` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionCriteria` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionListCel` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionList` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionNone` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementInclusionCriteria` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementInclusionListCel` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementInclusionList` **Added**
    - `AutomationSteps[].SetCredential` **Added**
    - `AutomationSteps[].StoreCredential` **Added**
    - `DraftTriggers[].FormTrigger` **Removed** (Breaking ⚠️)
    - `DraftTriggers[].GrantDeletedTrigger.GrantTriggerFilter.EntitlementFilter` **Removed** (Breaking ⚠️)
    - `DraftTriggers[].ManualAutomationTrigger` **Removed** (Breaking ⚠️)
    - `DraftTriggers[].WebhookAutomationTrigger.WebhookListenerAuthCapabilityUrl` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `Automation.AutomationSteps[].ConnectorCreateAccount.PasswordCel` **Added**
    - `Automation.AutomationSteps[].GeneratePassword.GeneratePasswordPolicy` **Added**
    - `Automation.AutomationSteps[].GeneratePassword.PasswordPolicyId` **Added**
    - `Automation.AutomationSteps[].GrantEntitlements.AppEntitlementRefsCel` **Removed** (Breaking ⚠️)
    - `Automation.AutomationSteps[].GrantEntitlements.AppEntitlementRefs` **Removed** (Breaking ⚠️)
    - `Automation.AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionCriteria` **Added**
    - `Automation.AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionListCel` **Added**
    - `Automation.AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionList` **Added**
    - `Automation.AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionNone` **Added**
    - `Automation.AutomationSteps[].GrantEntitlements.GrantEntitlementInclusionCriteria` **Added**
    - `Automation.AutomationSteps[].GrantEntitlements.GrantEntitlementInclusionListCel` **Added**
    - `Automation.AutomationSteps[].GrantEntitlements.GrantEntitlementInclusionList` **Added**
    - `Automation.AutomationSteps[].SetCredential` **Added**
    - `Automation.AutomationSteps[].StoreCredential` **Added**
    - `Automation.DraftTriggers[].FormTrigger` **Removed** (Breaking ⚠️)
    - `Automation.DraftTriggers[].GrantDeletedTrigger.GrantTriggerFilter.EntitlementFilter` **Removed** (Breaking ⚠️)
    - `Automation.DraftTriggers[].ManualAutomationTrigger` **Removed** (Breaking ⚠️)
    - `Automation.DraftTriggers[].WebhookAutomationTrigger.WebhookListenerAuthCapabilityUrl` **Added**
    - `WebhookCapabilityUrl` **Added**
* `ConductoroneApi.Automation.GetAutomation()`: `response.Automation` **Changed** (Breaking ⚠️)
    - `AutomationSteps[].ConnectorCreateAccount.PasswordCel` **Added**
    - `AutomationSteps[].GeneratePassword.GeneratePasswordPolicy` **Added**
    - `AutomationSteps[].GeneratePassword.PasswordPolicyId` **Added**
    - `AutomationSteps[].GrantEntitlements.AppEntitlementRefsCel` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].GrantEntitlements.AppEntitlementRefs` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionCriteria` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionListCel` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionList` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionNone` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementInclusionCriteria` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementInclusionListCel` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementInclusionList` **Added**
    - `AutomationSteps[].SetCredential` **Added**
    - `AutomationSteps[].StoreCredential` **Added**
    - `DraftTriggers[].FormTrigger` **Removed** (Breaking ⚠️)
    - `DraftTriggers[].GrantDeletedTrigger.GrantTriggerFilter.EntitlementFilter` **Removed** (Breaking ⚠️)
    - `DraftTriggers[].ManualAutomationTrigger` **Removed** (Breaking ⚠️)
    - `DraftTriggers[].WebhookAutomationTrigger.WebhookListenerAuthCapabilityUrl` **Added**
* `ConductoroneApi.Automation.ListAutomations()`: `response.List[]` **Changed** (Breaking ⚠️)
    - `AutomationSteps[].ConnectorCreateAccount.PasswordCel` **Added**
    - `AutomationSteps[].GeneratePassword.GeneratePasswordPolicy` **Added**
    - `AutomationSteps[].GeneratePassword.PasswordPolicyId` **Added**
    - `AutomationSteps[].GrantEntitlements.AppEntitlementRefsCel` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].GrantEntitlements.AppEntitlementRefs` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionCriteria` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionListCel` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionList` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionNone` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementInclusionCriteria` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementInclusionListCel` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementInclusionList` **Added**
    - `AutomationSteps[].SetCredential` **Added**
    - `AutomationSteps[].StoreCredential` **Added**
    - `DraftTriggers[].FormTrigger` **Removed** (Breaking ⚠️)
    - `DraftTriggers[].GrantDeletedTrigger.GrantTriggerFilter.EntitlementFilter` **Removed** (Breaking ⚠️)
    - `DraftTriggers[].ManualAutomationTrigger` **Removed** (Breaking ⚠️)
    - `DraftTriggers[].WebhookAutomationTrigger.WebhookListenerAuthCapabilityUrl` **Added**
* `ConductoroneApi.FunctionsSearch.Search()`: 
  *  `request.Request.FunctionTypes[].Enum(functionTypeCodeMode)` **Added**
  * `response.List[]` **Changed** (Breaking ⚠️)
    - `EncryptedValues` **Removed** (Breaking ⚠️)
    - `FunctionType.Enum(functionTypeCodeMode)` **Added**
    - `Secret` **Added**
* `ConductoroneApi.Functions.UpdateFunction()`: 
  * `request.Request.Function` **Changed** (Breaking ⚠️)
    - `EncryptedValues` **Removed** (Breaking ⚠️)
    - `FunctionType.Enum(functionTypeCodeMode)` **Added**
    - `Secret` **Added**
  * `response.Function` **Changed** (Breaking ⚠️)
    - `EncryptedValues` **Removed** (Breaking ⚠️)
    - `FunctionType.Enum(functionTypeCodeMode)` **Added**
    - `Secret` **Added**
* `ConductoroneApi.Functions.ListFunctions()`: `response.List[]` **Changed** (Breaking ⚠️)
    - `EncryptedValues` **Removed** (Breaking ⚠️)
    - `FunctionType.Enum(functionTypeCodeMode)` **Added**
    - `Secret` **Added**
* `ConductoroneApi.Functions.GetFunction()`: `response.Function` **Changed** (Breaking ⚠️)
    - `EncryptedValues` **Removed** (Breaking ⚠️)
    - `FunctionType.Enum(functionTypeCodeMode)` **Added**
    - `Secret` **Added**
* `ConductoroneApi.Functions.CreateFunction()`: 
  *  `request.Request.FunctionType.Enum(functionTypeCodeMode)` **Added**
  * `response.Function` **Changed** (Breaking ⚠️)
    - `EncryptedValues` **Removed** (Breaking ⚠️)
    - `FunctionType.Enum(functionTypeCodeMode)` **Added**
    - `Secret` **Added**
* `ConductoroneApi.ConnectorCatalog.ConfigurationSchema()`: `response` **Changed** (Breaking ⚠️)
    - `ConfigSchema.Fields[].KeyValueField.SupportsFileUpload` **Added**
    - `ConfigSchema.Fields[].StringField.PickerField.C1UserFilter` **Added**
    - `ConfigSchema.Fields[].StringField.TextField.Suffix` **Added**
    - `ConfigSchema.Fields[].StringMapField.DefaultValue` **Added**
    - `ConfigSchema.Fields[].StringMapField.Optional` **Removed** (Breaking ⚠️)
    - `ConfigSchema.Fields[].StringMapField.StringMapRules` **Added**
    - `Form.FieldRelationships[].DependentOn` **Added**
    - `Form.Fields[].Required` **Added**
    - `Form.Fields[].StringMapField` **Added**
* `ConductoroneApi.Automation.UpdateAutomation()`: 
  * `request.Request.UpdateAutomationRequest.Automation` **Changed** (Breaking ⚠️)
    - `AutomationSteps[].ConnectorCreateAccount.PasswordCel` **Added**
    - `AutomationSteps[].GeneratePassword.GeneratePasswordPolicy` **Added**
    - `AutomationSteps[].GeneratePassword.PasswordPolicyId` **Added**
    - `AutomationSteps[].GrantEntitlements.AppEntitlementRefsCel` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].GrantEntitlements.AppEntitlementRefs` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionCriteria` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionListCel` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionList` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionNone` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementInclusionCriteria` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementInclusionListCel` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementInclusionList` **Added**
    - `AutomationSteps[].SetCredential` **Added**
    - `AutomationSteps[].StoreCredential` **Added**
    - `DraftTriggers[].FormTrigger` **Removed** (Breaking ⚠️)
    - `DraftTriggers[].GrantDeletedTrigger.GrantTriggerFilter.EntitlementFilter` **Removed** (Breaking ⚠️)
    - `DraftTriggers[].ManualAutomationTrigger` **Removed** (Breaking ⚠️)
    - `DraftTriggers[].WebhookAutomationTrigger.WebhookListenerAuthCapabilityUrl` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `Automation.AutomationSteps[].ConnectorCreateAccount.PasswordCel` **Added**
    - `Automation.AutomationSteps[].GeneratePassword.GeneratePasswordPolicy` **Added**
    - `Automation.AutomationSteps[].GeneratePassword.PasswordPolicyId` **Added**
    - `Automation.AutomationSteps[].GrantEntitlements.AppEntitlementRefsCel` **Removed** (Breaking ⚠️)
    - `Automation.AutomationSteps[].GrantEntitlements.AppEntitlementRefs` **Removed** (Breaking ⚠️)
    - `Automation.AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionCriteria` **Added**
    - `Automation.AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionListCel` **Added**
    - `Automation.AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionList` **Added**
    - `Automation.AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionNone` **Added**
    - `Automation.AutomationSteps[].GrantEntitlements.GrantEntitlementInclusionCriteria` **Added**
    - `Automation.AutomationSteps[].GrantEntitlements.GrantEntitlementInclusionListCel` **Added**
    - `Automation.AutomationSteps[].GrantEntitlements.GrantEntitlementInclusionList` **Added**
    - `Automation.AutomationSteps[].SetCredential` **Added**
    - `Automation.AutomationSteps[].StoreCredential` **Added**
    - `Automation.DraftTriggers[].FormTrigger` **Removed** (Breaking ⚠️)
    - `Automation.DraftTriggers[].GrantDeletedTrigger.GrantTriggerFilter.EntitlementFilter` **Removed** (Breaking ⚠️)
    - `Automation.DraftTriggers[].ManualAutomationTrigger` **Removed** (Breaking ⚠️)
    - `Automation.DraftTriggers[].WebhookAutomationTrigger.WebhookListenerAuthCapabilityUrl` **Added**
    - `WebhookCapabilityUrl` **Added**
* `ConductoroneApi.AccessReview.Get()`: `response.AccessReviewView.AccessReview` **Changed**
    - `AccessReviewColumnConfig` **Added**
    - `AccessReviewScopeV2.AppSelectionCriteriaScope.ComplianceFrameworkAttributeValueIds` **Added**
    - `AccessReviewScopeV2.AppSelectionCriteriaScope.RiskLevelAttributeValueIds` **Added**
    - `AccessReviewScopeV2.ResourceSelectionScope` **Added**
    - `AccuracyIssueAction` **Added**
    - `AutoCloseCampaign` **Added**
    - `AutoCloseDecision` **Added**
    - `AutoStartCampaign` **Added**
    - `CampaignInsights` **Added**
    - `DefaultView.Enum(accessReviewViewTypeByResource)` **Added**
    - `NotificationConfig.SendKickoff` **Added**
    - `ScheduledStartDate` **Added**
    - `ScopeType.Enum(accessReviewScopeTypeByResource)` **Added**
* `ConductoroneApi.StepUpProvider.List()`:  `response.List[].StepUpMicrosoftSettings.ValidationMode` **Added**
* `ConductoroneApi.AccessReviewTemplateSetupEntitlement.SetScopeByResourceType()`: **Added**
* `ConductoroneApi.TenantAuthConfig.Create()`: **Added**
* `ConductoroneApi.TenantAuthConfig.Delete()`: **Added**
* `ConductoroneApi.TenantAuthConfig.Get()`: **Added**
* `ConductoroneApi.TenantAuthConfig.List()`: **Added**
* `ConductoroneApi.TenantAuthConfig.Update()`: **Added**
* `ConductoroneApi.Finding.BulkCreateFindingTasks()`: **Added**
* `ConductoroneApi.Finding.BulkUpdateFindingState()`: **Added**
* `ConductoroneApi.Finding.CreateFindingTask()`: **Added**
* `ConductoroneApi.Finding.GetFinding()`: **Added**
* `ConductoroneApi.Finding.UpdateFindingState()`: **Added**
* `ConductoroneApi.FindingRoutingRule.CreateFindingRoutingRule()`: **Added**
* `ConductoroneApi.FindingRoutingRule.DeleteFindingRoutingRule()`: **Added**
* `ConductoroneApi.FindingRoutingRule.GetFindingRoutingRule()`: **Added**
* `ConductoroneApi.FindingRoutingRule.ListFindingRoutingRules()`: **Added**
* `ConductoroneApi.FindingRoutingRule.UpdateFindingRoutingRule()`: **Added**
* `ConductoroneApi.FindingSearch.Search()`: **Added**
* `ConductoroneApi.Functions.Test()`: **Added**
* `ConductoroneApi.FunctionsInvocationSearch.Search()`: **Added**
* `ConductoroneApi.LocalDirectoryConfig.Create()`: **Added**
* `ConductoroneApi.LocalDirectoryConfig.Delete()`: **Added**
* `ConductoroneApi.LocalDirectoryConfig.Get()`: **Added**
* `ConductoroneApi.LocalDirectoryConfig.List()`: **Added**
* `ConductoroneApi.LocalDirectoryConfig.Update()`: **Added**
* `ConductoroneApi.LocalUserInvitation.Create()`: **Added**
* `ConductoroneApi.LocalUserInvitation.Get()`: **Added**
* `ConductoroneApi.LocalUserInvitation.Revoke()`: **Added**
* `ConductoroneApi.LocalUserInvitation.Search()`: **Added**
* `ConductoroneApi.RoleMiningManagement.CreateAccessProfileFromCohort()`: **Added**
* `ConductoroneApi.RoleMiningManagement.GetLatestRun()`: **Added**
* `ConductoroneApi.RoleMiningManagement.GetRoleMiningConfig()`: **Added**
* `ConductoroneApi.RoleMiningManagement.GetSuggestion()`: **Added**
* `ConductoroneApi.RoleMiningManagement.ListRuns()`: **Added**
* `ConductoroneApi.RoleMiningManagement.ListSuggestions()`: **Added**
* `ConductoroneApi.RoleMiningManagement.SearchCohortUsers()`: **Added**
* `ConductoroneApi.RoleMiningManagement.TriggerAnalysis()`: **Added**
* `ConductoroneApi.RoleMiningManagement.UpdateRoleMiningConfig()`: **Added**
* `ConductoroneApi.RoleMiningManagement.UpdateSuggestionState()`: **Added**
* `ConductoroneApi.AutomationExecutionSearch.SearchAllAutomationExecutions()`: **Added**
* `ConductoroneApi.ExternalClientSearch.Search()`: **Added**
* `ConductoroneApi.RoleMiningManagementSearch.Search()`: **Added**
* `ConductoroneApi.PaperSecretAdmin.Get()`: **Added**
* `ConductoroneApi.PaperSecretAdmin.Revoke()`: **Added**
* `ConductoroneApi.PaperSecretAdmin.Search()`: **Added**
* `ConductoroneApi.PaperSecretAdmin.SearchAuditEvents()`: **Added**
* `ConductoroneApi.PaperSecret.CreateExternal()`: **Added**
* `ConductoroneApi.PaperSecret.CreateInternal()`: **Added**
* `ConductoroneApi.PaperSecret.Get()`: **Added**
* `ConductoroneApi.PaperSecret.GetByShareCode()`: **Added**
* `ConductoroneApi.PaperSecret.GetContent()`: **Added**
* `ConductoroneApi.PaperSecret.Revoke()`: **Added**
* `ConductoroneApi.PaperSecret.SearchAuditEvents()`: **Added**
* `ConductoroneApi.PaperSecret.SearchMySecrets()`: **Added**
* `ConductoroneApi.PaperSecret.SetTextContent()`: **Added**
* `ConductoroneApi.SsfReceiverEventSearch.Search()`: **Added**
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
* `ConductoroneApi.Principal.Create()`: **Added**
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
* `ConductoroneApi.OnboardingSettings.Get()`: **Added**
* `ConductoroneApi.OnboardingSettings.Update()`: **Added**
* `ConductoroneApi.SsfReceiverStream.Create()`: **Added**
* `ConductoroneApi.SsfReceiverStream.Delete()`: **Added**
* `ConductoroneApi.SsfReceiverStream.Get()`: **Added**
* `ConductoroneApi.SsfReceiverStream.GetStats()`: **Added**
* `ConductoroneApi.SsfReceiverStream.List()`: **Added**
* `ConductoroneApi.SsfReceiverStream.Test()`: **Added**
* `ConductoroneApi.SsfReceiverStream.Update()`: **Added**
* `ConductoroneApi.SsfReceiverEvent.List()`: **Added**
* `ConductoroneApi.Functions.GetFunctionSecretEncryptionKey()`: **Removed** (Breaking ⚠️)
* `ConductoroneApi.AccessReview.Create()`: 
  * `request.Request` **Changed**
    - `AccessReviewScopeV2` **Added**
    - `NotificationConfig.SendKickoff` **Added**
    - `ScopeType.Enum(accessReviewScopeTypeByResource)` **Added**
  * `response.AccessReviewView.AccessReview` **Changed**
    - `AccessReviewColumnConfig` **Added**
    - `AccessReviewScopeV2.AppSelectionCriteriaScope.ComplianceFrameworkAttributeValueIds` **Added**
    - `AccessReviewScopeV2.AppSelectionCriteriaScope.RiskLevelAttributeValueIds` **Added**
    - `AccessReviewScopeV2.ResourceSelectionScope` **Added**
    - `AccuracyIssueAction` **Added**
    - `AutoCloseCampaign` **Added**
    - `AutoCloseDecision` **Added**
    - `AutoStartCampaign` **Added**
    - `CampaignInsights` **Added**
    - `DefaultView.Enum(accessReviewViewTypeByResource)` **Added**
    - `NotificationConfig.SendKickoff` **Added**
    - `ScheduledStartDate` **Added**
    - `ScopeType.Enum(accessReviewScopeTypeByResource)` **Added**
* `ConductoroneApi.AccessReviewTemplateSetupEntitlement.GetScopeAndEntitlements()`: **Added**
* `ConductoroneApi.AccessReview.List()`: `response.List[].AccessReview` **Changed**
    - `AccessReviewColumnConfig` **Added**
    - `AccessReviewScopeV2.AppSelectionCriteriaScope.ComplianceFrameworkAttributeValueIds` **Added**
    - `AccessReviewScopeV2.AppSelectionCriteriaScope.RiskLevelAttributeValueIds` **Added**
    - `AccessReviewScopeV2.ResourceSelectionScope` **Added**
    - `AccuracyIssueAction` **Added**
    - `AutoCloseCampaign` **Added**
    - `AutoCloseDecision` **Added**
    - `AutoStartCampaign` **Added**
    - `CampaignInsights` **Added**
    - `DefaultView.Enum(accessReviewViewTypeByResource)` **Added**
    - `NotificationConfig.SendKickoff` **Added**
    - `ScheduledStartDate` **Added**
    - `ScopeType.Enum(accessReviewScopeTypeByResource)` **Added**
* `ConductoroneApi.AccessReview.Update()`: 
  * `request.Request.AccessReviewServiceUpdateRequest.AccessReview` **Changed**
    - `AccessReviewColumnConfig` **Added**
    - `AccessReviewScopeV2.AppSelectionCriteriaScope.ComplianceFrameworkAttributeValueIds` **Added**
    - `AccessReviewScopeV2.AppSelectionCriteriaScope.RiskLevelAttributeValueIds` **Added**
    - `AccessReviewScopeV2.ResourceSelectionScope` **Added**
    - `AccuracyIssueAction` **Added**
    - `AutoCloseCampaign` **Added**
    - `AutoCloseDecision` **Added**
    - `AutoStartCampaign` **Added**
    - `CampaignInsights` **Added**
    - `DefaultView.Enum(accessReviewViewTypeByResource)` **Added**
    - `NotificationConfig.SendKickoff` **Added**
    - `ScheduledStartDate` **Added**
    - `ScopeType.Enum(accessReviewScopeTypeByResource)` **Added**
  * `response.AccessReviewView.AccessReview` **Changed**
    - `AccessReviewColumnConfig` **Added**
    - `AccessReviewScopeV2.AppSelectionCriteriaScope.ComplianceFrameworkAttributeValueIds` **Added**
    - `AccessReviewScopeV2.AppSelectionCriteriaScope.RiskLevelAttributeValueIds` **Added**
    - `AccessReviewScopeV2.ResourceSelectionScope` **Added**
    - `AccuracyIssueAction` **Added**
    - `AutoCloseCampaign` **Added**
    - `AutoCloseDecision` **Added**
    - `AutoStartCampaign` **Added**
    - `CampaignInsights` **Added**
    - `DefaultView.Enum(accessReviewViewTypeByResource)` **Added**
    - `NotificationConfig.SendKickoff` **Added**
    - `ScheduledStartDate` **Added**
    - `ScopeType.Enum(accessReviewScopeTypeByResource)` **Added**
* `ConductoroneApi.AccessReviewTemplate.Create()`: 
  * `request.Request` **Changed**
    - `AccessReviewColumnConfig` **Added**
    - `AccessReviewScopeV2` **Added**
    - `AccuracyIssueAction` **Added**
    - `AutoCloseCampaign` **Added**
    - `AutoCloseDecision` **Added**
    - `AutoGenerateReport` **Added**
    - `AutoStartCampaign` **Added**
    - `DefaultView` **Added**
    - `ExemptCertifiedAccessConflicts` **Added**
    - `IsCampaignScheduleEnabled` **Added**
    - `NotificationConfig.SendKickoff` **Added**
    - `RecurrenceRule` **Added**
    - `ReviewInstructions` **Added**
    - `ReviewSignatureConfig` **Added**
    - `ScopeType.Enum(accessReviewScopeTypeByResource)` **Added**
    - `UsePolicyOverride` **Added**
  * `response.AccessReviewTemplate` **Changed**
    - `AccessReviewColumnConfig` **Added**
    - `AccessReviewScopeV2.AppSelectionCriteriaScope.ComplianceFrameworkAttributeValueIds` **Added**
    - `AccessReviewScopeV2.AppSelectionCriteriaScope.RiskLevelAttributeValueIds` **Added**
    - `AccessReviewScopeV2.ResourceSelectionScope` **Added**
    - `AccuracyIssueAction` **Added**
    - `AutoCloseCampaign` **Added**
    - `AutoCloseDecision` **Added**
    - `AutoStartCampaign` **Added**
    - `DefaultView.Enum(accessReviewViewTypeByResource)` **Added**
    - `NotificationConfig.SendKickoff` **Added**
    - `ScopeType.Enum(accessReviewScopeTypeByResource)` **Added**
* `ConductoroneApi.AccessReviewTemplate.Get()`: `response.AccessReviewTemplate` **Changed**
    - `AccessReviewColumnConfig` **Added**
    - `AccessReviewScopeV2.AppSelectionCriteriaScope.ComplianceFrameworkAttributeValueIds` **Added**
    - `AccessReviewScopeV2.AppSelectionCriteriaScope.RiskLevelAttributeValueIds` **Added**
    - `AccessReviewScopeV2.ResourceSelectionScope` **Added**
    - `AccuracyIssueAction` **Added**
    - `AutoCloseCampaign` **Added**
    - `AutoCloseDecision` **Added**
    - `AutoStartCampaign` **Added**
    - `DefaultView.Enum(accessReviewViewTypeByResource)` **Added**
    - `NotificationConfig.SendKickoff` **Added**
    - `ScopeType.Enum(accessReviewScopeTypeByResource)` **Added**
* `ConductoroneApi.AccessReviewTemplate.Update()`: 
  * `request.Request.AccessReviewTemplateServiceUpdateRequest.AccessReviewTemplate` **Changed**
    - `AccessReviewColumnConfig` **Added**
    - `AccessReviewScopeV2.AppSelectionCriteriaScope.ComplianceFrameworkAttributeValueIds` **Added**
    - `AccessReviewScopeV2.AppSelectionCriteriaScope.RiskLevelAttributeValueIds` **Added**
    - `AccessReviewScopeV2.ResourceSelectionScope` **Added**
    - `AccuracyIssueAction` **Added**
    - `AutoCloseCampaign` **Added**
    - `AutoCloseDecision` **Added**
    - `AutoStartCampaign` **Added**
    - `DefaultView.Enum(accessReviewViewTypeByResource)` **Added**
    - `NotificationConfig.SendKickoff` **Added**
    - `ScopeType.Enum(accessReviewScopeTypeByResource)` **Added**
  * `response.AccessReviewTemplate` **Changed**
    - `AccessReviewColumnConfig` **Added**
    - `AccessReviewScopeV2.AppSelectionCriteriaScope.ComplianceFrameworkAttributeValueIds` **Added**
    - `AccessReviewScopeV2.AppSelectionCriteriaScope.RiskLevelAttributeValueIds` **Added**
    - `AccessReviewScopeV2.ResourceSelectionScope` **Added**
    - `AccuracyIssueAction` **Added**
    - `AutoCloseCampaign` **Added**
    - `AutoCloseDecision` **Added**
    - `AutoStartCampaign` **Added**
    - `DefaultView.Enum(accessReviewViewTypeByResource)` **Added**
    - `NotificationConfig.SendKickoff` **Added**
    - `ScopeType.Enum(accessReviewScopeTypeByResource)` **Added**
* `ConductoroneApi.RequestCatalogManagement.Get()`: `response.RequestCatalogView.RequestCatalog` **Changed**
    - `AccessEntitlements[].ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `AccessEntitlements[].Purpose.Enum(appEntitlementPurposeValueOwnership)` **Added**
    - `GrantPolicyId` **Added**
* `ConductoroneApi.AccessReviewTemplateSetupEntitlement.SetScopeAndEntitlements()`: **Added**
* `ConductoroneApi.Apps.Create()`: 
  *  `request.Request.IdentityMatching.Enum(appUserIdentityMatchingCustom)` **Added**
  * `response.App` **Changed**
    - `AppOwners[].Origin` **Added**
    - `AppUserMapper` **Added**
    - `EnableConnectorSourcedOwnership` **Added**
    - `IdentityMatching.Enum(appUserIdentityMatchingCustom)` **Added**
* `ConductoroneApi.Apps.Update()`: 
  * `request.Request.UpdateAppRequest.App` **Changed**
    - `AppUserMapper` **Added**
    - `EnableConnectorSourcedOwnership` **Added**
    - `IdentityMatching.Enum(appUserIdentityMatchingCustom)` **Added**
  * `response.App` **Changed**
    - `AppOwners[].Origin` **Added**
    - `AppUserMapper` **Added**
    - `EnableConnectorSourcedOwnership` **Added**
    - `IdentityMatching.Enum(appUserIdentityMatchingCustom)` **Added**
* `ConductoroneApi.Connector.Create()`: `response.ConnectorView.Connector` **Changed**
    - `ConfigUpdatedAt` **Added**
    - `ConnectorApiVersion` **Added**
    - `ConnectorSyncCronSchedule` **Added**
* `ConductoroneApi.Connector.CreateDelegated()`: `response.ConnectorView.Connector` **Changed**
    - `ConfigUpdatedAt` **Added**
    - `ConnectorApiVersion` **Added**
    - `ConnectorSyncCronSchedule` **Added**
* `ConductoroneApi.Connector.Get()`: `response.ConnectorView.Connector` **Changed**
    - `ConfigUpdatedAt` **Added**
    - `ConnectorApiVersion` **Added**
    - `ConnectorSyncCronSchedule` **Added**
* `ConductoroneApi.Connector.List()`: `response.List[].Connector` **Changed**
    - `ConfigUpdatedAt` **Added**
    - `ConnectorApiVersion` **Added**
    - `ConnectorSyncCronSchedule` **Added**
* `ConductoroneApi.Connector.Update()`: 
  *  `request.Request.ConnectorServiceUpdateRequest.Connector.ConnectorSyncCronSchedule` **Added**
  * `response.ConnectorView.Connector` **Changed**
    - `ConfigUpdatedAt` **Added**
    - `ConnectorApiVersion` **Added**
    - `ConnectorSyncCronSchedule` **Added**
* `ConductoroneApi.Connector.UpdateDelegated()`: 
  *  `request.Request.ConnectorServiceUpdateDelegatedRequest.Connector.ConnectorSyncCronSchedule` **Added**
  * `response.ConnectorView.Connector` **Changed**
    - `ConfigUpdatedAt` **Added**
    - `ConnectorApiVersion` **Added**
    - `ConnectorSyncCronSchedule` **Added**
* `ConductoroneApi.AppAccessRequestsDefaults.CancelAppAccessRequestsDefaults()`:  `response.RequestSchemaId` **Added**
* `ConductoroneApi.AppAccessRequestsDefaults.CreateAppAccessRequestsDefaults()`: 
  *  `request.Request.AppAccessRequestDefaults.RequestSchemaId` **Added**
  *  `response.RequestSchemaId` **Added**
* `ConductoroneApi.AppAccessRequestsDefaults.GetAppAccessRequestsDefaults()`:  `response.RequestSchemaId` **Added**
* `ConductoroneApi.AppEntitlements.AddManuallyManagedMembers()`:  `response.BulkActionId` **Added**
* `ConductoroneApi.AppEntitlements.Create()`: 
  * `request.Request.CreateAppEntitlementRequest` **Changed**
    - `ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Purpose.Enum(appEntitlementPurposeValueOwnership)` **Added**
  * `response.AppEntitlementView` **Changed**
    - `ActorObjectPermissions` **Added**
    - `AppEntitlement.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `AppEntitlement.Purpose.Enum(appEntitlementPurposeValueOwnership)` **Added**
* `ConductoroneApi.AppEntitlements.CreateAutomation()`:  `response.AppEntitlementAutomation.ManagedByRequestCatalogId` **Added**
* `ConductoroneApi.AppEntitlements.Get()`: `response.AppEntitlementView` **Changed**
    - `ActorObjectPermissions` **Added**
    - `AppEntitlement.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `AppEntitlement.Purpose.Enum(appEntitlementPurposeValueOwnership)` **Added**
* `ConductoroneApi.AppEntitlements.GetAutomation()`:  `response.AppEntitlementAutomation.ManagedByRequestCatalogId` **Added**
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
* `ConductoroneApi.AppEntitlements.ListUsers()`:  `response.List[].OriginatingTicketId` **Added**
* `ConductoroneApi.AppEntitlements.Update()`: 
  * `request.Request.UpdateAppEntitlementRequest.AppEntitlement` **Changed**
    - `ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Purpose.Enum(appEntitlementPurposeValueOwnership)` **Added**
  * `response.AppEntitlementView` **Changed**
    - `ActorObjectPermissions` **Added**
    - `AppEntitlement.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `AppEntitlement.Purpose.Enum(appEntitlementPurposeValueOwnership)` **Added**
* `ConductoroneApi.AppEntitlements.UpdateAutomation()`:  `response.AppEntitlementAutomation.ManagedByRequestCatalogId` **Added**
* `ConductoroneApi.AppEntitlementSearch.Search()`: `response.List[]` **Changed**
    - `ActorObjectPermissions` **Added**
    - `AppEntitlement.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `AppEntitlement.Purpose.Enum(appEntitlementPurposeValueOwnership)` **Added**
* `ConductoroneApi.AppEntitlementSearch.SearchAppEntitlementsForAppUser()`: `response.List[]` **Changed**
    - `ActorObjectPermissions` **Added**
    - `AppEntitlement.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `AppEntitlement.Purpose.Enum(appEntitlementPurposeValueOwnership)` **Added**
* `ConductoroneApi.AppEntitlementSearch.SearchAppEntitlementsWithExpired()`:  `response.List[].User.Origin` **Added**
* `ConductoroneApi.AppEntitlementSearch.SearchGrants()`: 
  *  `request.Request.Purpose` **Added**
  * `response.List[]` **Changed**
    - `AppEntitlementUserView.OriginatingTicketId` **Added**
    - `AppEntitlementView.ActorObjectPermissions` **Added**
    - `AppEntitlementView.AppEntitlement.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `AppEntitlementView.AppEntitlement.Purpose.Enum(appEntitlementPurposeValueOwnership)` **Added**
* `ConductoroneApi.AppEntitlementOwners.List()`:  `response.List[].Origin` **Added**
* `ConductoroneApi.AppOwners.List()`:  `response.List[].Origin` **Added**
* `ConductoroneApi.AccessReviewSetupEntitlement.SetCampaignScopeByResourceType()`: **Added**
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
* `ConductoroneApi.RequestCatalogManagement.Create()`: 
  *  `request.Request.GrantPolicyId` **Added**
  * `response.RequestCatalogView.RequestCatalog` **Changed**
    - `AccessEntitlements[].ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `AccessEntitlements[].Purpose.Enum(appEntitlementPurposeValueOwnership)` **Added**
    - `GrantPolicyId` **Added**
* `ConductoroneApi.RequestCatalogManagement.CreateBundleAutomation()`: 
  *  `request.Request.CreateBundleAutomationRequest.BundleAutomationRuleCel` **Added**
  * `response` **Changed**
    - `BundleAutomationLastRunState.BundleAutomationCelEvaluationState` **Added**
    - `BundleAutomationRuleCel` **Added**
* `ConductoroneApi.Apps.List()`: `response.List[]` **Changed**
    - `AppOwners[].Origin` **Added**
    - `AppUserMapper` **Added**
    - `EnableConnectorSourcedOwnership` **Added**
    - `IdentityMatching.Enum(appUserIdentityMatchingCustom)` **Added**
* `ConductoroneApi.RequestCatalogManagement.GetBundleAutomation()`: `response` **Changed**
    - `BundleAutomationLastRunState.BundleAutomationCelEvaluationState` **Added**
    - `BundleAutomationRuleCel` **Added**
* `ConductoroneApi.RequestCatalogManagement.List()`: `response.List[].RequestCatalog` **Changed**
    - `AccessEntitlements[].ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `AccessEntitlements[].Purpose.Enum(appEntitlementPurposeValueOwnership)` **Added**
    - `GrantPolicyId` **Added**
* `ConductoroneApi.RequestCatalogManagement.ListEntitlementsForAccess()`: `response.List[]` **Changed**
    - `ActorObjectPermissions` **Added**
    - `AppEntitlement.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `AppEntitlement.Purpose.Enum(appEntitlementPurposeValueOwnership)` **Added**
* `ConductoroneApi.RequestCatalogManagement.ListEntitlementsPerCatalog()`: `response.List[]` **Changed**
    - `ActorObjectPermissions` **Added**
    - `AppEntitlement.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `AppEntitlement.Purpose.Enum(appEntitlementPurposeValueOwnership)` **Added**
* `ConductoroneApi.RequestCatalogManagement.SetBundleAutomation()`: 
  *  `request.Request.SetBundleAutomationRequest.BundleAutomationRuleCel` **Added**
  * `response` **Changed**
    - `BundleAutomationLastRunState.BundleAutomationCelEvaluationState` **Added**
    - `BundleAutomationRuleCel` **Added**
* `ConductoroneApi.RequestCatalogManagement.Update()`: 
  * `request.Request.RequestCatalogManagementServiceUpdateRequest.RequestCatalog` **Changed**
    - `AccessEntitlements[].ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `AccessEntitlements[].Purpose.Enum(appEntitlementPurposeValueOwnership)` **Added**
    - `GrantPolicyId` **Added**
  * `response.RequestCatalogView.RequestCatalog` **Changed**
    - `AccessEntitlements[].ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `AccessEntitlements[].Purpose.Enum(appEntitlementPurposeValueOwnership)` **Added**
    - `GrantPolicyId` **Added**
* `ConductoroneApi.AccessReviewSetupEntitlement.SetCampaignScopeAndEntitlements()`: **Added**
* `ConductoroneApi.Directory.Create()`: 
  *  `request.Request.DirectoryMergeConfig` **Added**
  *  `response.DirectoryView.Directory.DirectoryMergeConfig` **Added**
* `ConductoroneApi.Directory.Get()`:  `response.DirectoryView.Directory.DirectoryMergeConfig` **Added**
* `ConductoroneApi.Directory.List()`:  `response.List[].Directory.DirectoryMergeConfig` **Added**
* `ConductoroneApi.Directory.Update()`: 
  *  `request.Request.DirectoryServiceUpdateRequest.DirectoryMergeConfig` **Added**
  *  `response.DirectoryView.Directory.DirectoryMergeConfig` **Added**
* `ConductoroneApi.AccessReviewSetupEntitlement.GetCampaignScopeAndEntitlements()`: **Added**
* `ConductoroneApi.A2Ui.SubmitAction()`: **Added**
* `ConductoroneApi.Functions.Invoke()`: 
  *  `request.Request.FunctionsServiceInvokeRequest.VfsId` **Added**
* `ConductoroneApi.A2Ui.ListSurfaces()`: **Added**
* `ConductoroneApi.A2Ui.ListSurfaceFeedback()`: **Added**
* `ConductoroneApi.Policies.Create()`: 
  * `request.Request.PolicySteps.Map<PolicySteps>.Steps[]` **Changed**
    - `Action.ActionTargetBatonResourceAction` **Added**
    - `Form.Form.FieldRelationships[].DependentOn` **Added**
    - `Form.Form.Fields[].Required` **Added**
    - `Form.Form.Fields[].StringField.PickerField.C1UserFilter` **Added**
    - `Form.Form.Fields[].StringField.TextField.Suffix` **Added**
    - `Form.Form.Fields[].StringMapField` **Added**
    - `Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
  * `response.Policy.PolicySteps.Map<PolicySteps>.Steps[]` **Changed**
    - `Action.ActionTargetBatonResourceAction` **Added**
    - `Form.Form.FieldRelationships[].DependentOn` **Added**
    - `Form.Form.Fields[].Required` **Added**
    - `Form.Form.Fields[].StringField.PickerField.C1UserFilter` **Added**
    - `Form.Form.Fields[].StringField.TextField.Suffix` **Added**
    - `Form.Form.Fields[].StringMapField` **Added**
    - `Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
* `ConductoroneApi.Policies.Get()`: `response.Policy.PolicySteps.Map<PolicySteps>.Steps[]` **Changed**
    - `Action.ActionTargetBatonResourceAction` **Added**
    - `Form.Form.FieldRelationships[].DependentOn` **Added**
    - `Form.Form.Fields[].Required` **Added**
    - `Form.Form.Fields[].StringField.PickerField.C1UserFilter` **Added**
    - `Form.Form.Fields[].StringField.TextField.Suffix` **Added**
    - `Form.Form.Fields[].StringMapField` **Added**
    - `Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
* `ConductoroneApi.Policies.List()`: `response.List[].PolicySteps.Map<PolicySteps>.Steps[]` **Changed**
    - `Action.ActionTargetBatonResourceAction` **Added**
    - `Form.Form.FieldRelationships[].DependentOn` **Added**
    - `Form.Form.Fields[].Required` **Added**
    - `Form.Form.Fields[].StringField.PickerField.C1UserFilter` **Added**
    - `Form.Form.Fields[].StringField.TextField.Suffix` **Added**
    - `Form.Form.Fields[].StringMapField` **Added**
    - `Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
* `ConductoroneApi.Policies.Update()`: 
  * `request.Request.UpdatePolicyRequest.Policy.PolicySteps.Map<PolicySteps>.Steps[]` **Changed**
    - `Action.ActionTargetBatonResourceAction` **Added**
    - `Form.Form.FieldRelationships[].DependentOn` **Added**
    - `Form.Form.Fields[].Required` **Added**
    - `Form.Form.Fields[].StringField.PickerField.C1UserFilter` **Added**
    - `Form.Form.Fields[].StringField.TextField.Suffix` **Added**
    - `Form.Form.Fields[].StringMapField` **Added**
    - `Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
  * `response.Policy.PolicySteps.Map<PolicySteps>.Steps[]` **Changed**
    - `Action.ActionTargetBatonResourceAction` **Added**
    - `Form.Form.FieldRelationships[].DependentOn` **Added**
    - `Form.Form.Fields[].Required` **Added**
    - `Form.Form.Fields[].StringField.PickerField.C1UserFilter` **Added**
    - `Form.Form.Fields[].StringField.TextField.Suffix` **Added**
    - `Form.Form.Fields[].StringMapField` **Added**
    - `Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
* `ConductoroneApi.RequestSchema.Create()`: 
  * `request.Request` **Changed**
    - `FieldGroups` **Added**
    - `FieldRelationships` **Added**
    - `Fields[].Required` **Added**
    - `Fields[].StringField.PickerField.C1UserFilter` **Added**
    - `Fields[].StringField.TextField.Suffix` **Added**
    - `Fields[].StringMapField` **Added**
    - `JustificationVisibility` **Added**
  * `response.RequestSchema.Form` **Changed**
    - `FieldRelationships[].DependentOn` **Added**
    - `Fields[].Required` **Added**
    - `Fields[].StringField.PickerField.C1UserFilter` **Added**
    - `Fields[].StringField.TextField.Suffix` **Added**
    - `Fields[].StringMapField` **Added**
* `ConductoroneApi.RequestSchema.Get()`: `response.RequestSchema.Form` **Changed**
    - `FieldRelationships[].DependentOn` **Added**
    - `Fields[].Required` **Added**
    - `Fields[].StringField.PickerField.C1UserFilter` **Added**
    - `Fields[].StringField.TextField.Suffix` **Added**
    - `Fields[].StringMapField` **Added**
* `ConductoroneApi.RequestSchema.Update()`: 
  * `request.Request.RequestSchemaServiceUpdateRequest.RequestSchema.Form` **Changed**
    - `FieldRelationships[].DependentOn` **Added**
    - `Fields[].Required` **Added**
    - `Fields[].StringField.PickerField.C1UserFilter` **Added**
    - `Fields[].StringField.TextField.Suffix` **Added**
    - `Fields[].StringMapField` **Added**
  * `response.RequestSchema.Form` **Changed**
    - `FieldRelationships[].DependentOn` **Added**
    - `Fields[].Required` **Added**
    - `Fields[].StringField.PickerField.C1UserFilter` **Added**
    - `Fields[].StringField.TextField.Suffix` **Added**
    - `Fields[].StringMapField` **Added**
* `ConductoroneApi.AppResourceSearch.SearchAppResources()`: `response.List[]` **Changed**
    - `ActorObjectPermissions` **Added**
    - `AppResource.AccessConfigId` **Added**
    - `AppResource.Profile` **Added**
* `ConductoroneApi.AppSearch.Search()`: `response.List[]` **Changed**
    - `AppOwners[].Origin` **Added**
    - `AppUserMapper` **Added**
    - `EnableConnectorSourcedOwnership` **Added**
    - `IdentityMatching.Enum(appUserIdentityMatchingCustom)` **Added**
* `ConductoroneApi.A2Ui.CreateSurfaceFeedback()`: **Added**
* `ConductoroneApi.PolicySearch.Search()`: `response.List[].PolicySteps.Map<PolicySteps>.Steps[]` **Changed**
    - `Action.ActionTargetBatonResourceAction` **Added**
    - `Form.Form.FieldRelationships[].DependentOn` **Added**
    - `Form.Form.Fields[].Required` **Added**
    - `Form.Form.Fields[].StringField.PickerField.C1UserFilter` **Added**
    - `Form.Form.Fields[].StringField.TextField.Suffix` **Added**
    - `Form.Form.Fields[].StringMapField` **Added**
    - `Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
* `ConductoroneApi.RequestCatalogSearch.SearchEntitlements()`: `response.List[].AppEntitlementView` **Changed**
    - `ActorObjectPermissions` **Added**
    - `AppEntitlement.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `AppEntitlement.Purpose.Enum(appEntitlementPurposeValueOwnership)` **Added**
* `ConductoroneApi.StepUpProvider.Create()`: 
  *  `request.Request.StepUpMicrosoftSettings.ValidationMode` **Added**
  *  `response.StepUpProvider.StepUpMicrosoftSettings.ValidationMode` **Added**
* `ConductoroneApi.StepUpProvider.Get()`:  `response.StepUpProvider.StepUpMicrosoftSettings.ValidationMode` **Added**
* `ConductoroneApi.Apps.Get()`: `response.App` **Changed**
    - `AppOwners[].Origin` **Added**
    - `AppUserMapper` **Added**
    - `EnableConnectorSourcedOwnership` **Added**
    - `IdentityMatching.Enum(appUserIdentityMatchingCustom)` **Added**
* `ConductoroneApi.StepUpProvider.Search()`:  `response.List[].StepUpMicrosoftSettings.ValidationMode` **Added**
* `ConductoroneApi.StepUpProvider.Update()`: 
  *  `request.Request.UpdateStepUpProviderRequest.StepUpProvider.StepUpMicrosoftSettings.ValidationMode` **Added**
  *  `response.StepUpProvider.StepUpMicrosoftSettings.ValidationMode` **Added**
* `ConductoroneApi.StepUpProvider.UpdateSecret()`:  `response.StepUpProvider.StepUpMicrosoftSettings.ValidationMode` **Added**
* `ConductoroneApi.TaskSearch.Search()`: 
  * `request.Request` **Changed**
    - `ExcludeApplicationIds` **Added**
    - `PendingActionFilter` **Added**
    - `RequireApprovalReason` **Added**
    - `RequireDenialReason` **Added**
    - `TaskTypes[].TaskTypeFinding` **Added**
    - `TaskTypes[].TaskTypeGrant.TaskGrantSource.IsExtension` **Added**
  * `response.List[]` **Changed**
    - `ApproversPath` **Added**
    - `ResourceBindingsPath` **Added**
    - `Task.ApproverIds` **Added**
    - `Task.Form.FieldRelationships[].DependentOn` **Added**
    - `Task.Form.Fields[].Required` **Added**
    - `Task.Form.Fields[].StringField.PickerField.C1UserFilter` **Added**
    - `Task.Form.Fields[].StringField.TextField.Suffix` **Added**
    - `Task.Form.Fields[].StringMapField` **Added**
    - `Task.Origin.Enum(taskOriginCascadeDelete)` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Action.ActionTargetBatonResourceAction` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ActionInstance.ActionTargetBatonResourceActionInstance` **Added**
    - `Task.RevocationTargets` **Added**
    - `Task.TaskType.TaskTypeFinding` **Added**
    - `Task.TaskType.TaskTypeGrant.TaskGrantSource.IsExtension` **Added**
* `ConductoroneApi.UserSearch.Search()`: 
  * `request.Request` **Changed**
    - `DelegateStatus` **Added**
    - `DelegatedUserIds` **Added**
    - `ExcludeOrigins` **Added**
    - `IsDelegate` **Added**
    - `Origins` **Added**
  *  `response.List[].User.Origin` **Added**
* `ConductoroneApi.WebhooksSearch.Search()`:  `response.List[].CallbackTimeout` **Added**
* `ConductoroneApi.SessionSettings.Get()`: `response.SessionSettings` **Changed**
    - `CidrRestriction1` **Changed**
    - `CidrRestriction2` **Changed**
    - `CidrRestriction3` **Changed**
    - `CidrRestriction4` **Changed**
    - `CidrRestriction5` **Added**
    - `ClientIdApprovalRequestPolicyId` **Added**
    - `ClientIdMetadataDocumentPolicy` **Added**
    - `ExternalClientsEnabled` **Added**
* `ConductoroneApi.SessionSettings.Update()`: 
  * `request.Request.SessionSettings` **Changed**
    - `CidrRestriction1` **Changed**
    - `CidrRestriction2` **Changed**
    - `CidrRestriction3` **Changed**
    - `CidrRestriction4` **Changed**
    - `CidrRestriction5` **Added**
    - `ClientIdApprovalRequestPolicyId` **Added**
    - `ClientIdMetadataDocumentPolicy` **Added**
    - `ExternalClientsEnabled` **Added**
  * `response.SessionSettings` **Changed**
    - `CidrRestriction1` **Changed**
    - `CidrRestriction2` **Changed**
    - `CidrRestriction3` **Changed**
    - `CidrRestriction4` **Changed**
    - `CidrRestriction5` **Added**
    - `ClientIdApprovalRequestPolicyId` **Added**
    - `ClientIdMetadataDocumentPolicy` **Added**
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
    - `ResourceBindingsPath` **Added**
    - `Task.ApproverIds` **Added**
    - `Task.Form.FieldRelationships[].DependentOn` **Added**
    - `Task.Form.Fields[].Required` **Added**
    - `Task.Form.Fields[].StringField.PickerField.C1UserFilter` **Added**
    - `Task.Form.Fields[].StringField.TextField.Suffix` **Added**
    - `Task.Form.Fields[].StringMapField` **Added**
    - `Task.Origin.Enum(taskOriginCascadeDelete)` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Action.ActionTargetBatonResourceAction` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ActionInstance.ActionTargetBatonResourceActionInstance` **Added**
    - `Task.RevocationTargets` **Added**
    - `Task.TaskType.TaskTypeFinding` **Added**
    - `Task.TaskType.TaskTypeGrant.TaskGrantSource.IsExtension` **Added**
* `ConductoroneApi.Task.CreateOffboardingTask()`: `response.TaskView` **Changed**
    - `ApproversPath` **Added**
    - `ResourceBindingsPath` **Added**
    - `Task.ApproverIds` **Added**
    - `Task.Form.FieldRelationships[].DependentOn` **Added**
    - `Task.Form.Fields[].Required` **Added**
    - `Task.Form.Fields[].StringField.PickerField.C1UserFilter` **Added**
    - `Task.Form.Fields[].StringField.TextField.Suffix` **Added**
    - `Task.Form.Fields[].StringMapField` **Added**
    - `Task.Origin.Enum(taskOriginCascadeDelete)` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Action.ActionTargetBatonResourceAction` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ActionInstance.ActionTargetBatonResourceActionInstance` **Added**
    - `Task.RevocationTargets` **Added**
    - `Task.TaskType.TaskTypeFinding` **Added**
    - `Task.TaskType.TaskTypeGrant.TaskGrantSource.IsExtension` **Added**
* `ConductoroneApi.Task.CreateRevokeTask()`: `response.TaskView` **Changed**
    - `ApproversPath` **Added**
    - `ResourceBindingsPath` **Added**
    - `Task.ApproverIds` **Added**
    - `Task.Form.FieldRelationships[].DependentOn` **Added**
    - `Task.Form.Fields[].Required` **Added**
    - `Task.Form.Fields[].StringField.PickerField.C1UserFilter` **Added**
    - `Task.Form.Fields[].StringField.TextField.Suffix` **Added**
    - `Task.Form.Fields[].StringMapField` **Added**
    - `Task.Origin.Enum(taskOriginCascadeDelete)` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Action.ActionTargetBatonResourceAction` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ActionInstance.ActionTargetBatonResourceActionInstance` **Added**
    - `Task.RevocationTargets` **Added**
    - `Task.TaskType.TaskTypeFinding` **Added**
    - `Task.TaskType.TaskTypeGrant.TaskGrantSource.IsExtension` **Added**
* `ConductoroneApi.Task.Get()`: `response.TaskView` **Changed**
    - `ApproversPath` **Added**
    - `ResourceBindingsPath` **Added**
    - `Task.ApproverIds` **Added**
    - `Task.Form.FieldRelationships[].DependentOn` **Added**
    - `Task.Form.Fields[].Required` **Added**
    - `Task.Form.Fields[].StringField.PickerField.C1UserFilter` **Added**
    - `Task.Form.Fields[].StringField.TextField.Suffix` **Added**
    - `Task.Form.Fields[].StringMapField` **Added**
    - `Task.Origin.Enum(taskOriginCascadeDelete)` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Action.ActionTargetBatonResourceAction` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ActionInstance.ActionTargetBatonResourceActionInstance` **Added**
    - `Task.RevocationTargets` **Added**
    - `Task.TaskType.TaskTypeFinding` **Added**
    - `Task.TaskType.TaskTypeGrant.TaskGrantSource.IsExtension` **Added**
* `ConductoroneApi.TaskActions.Approve()`: `response.TaskView` **Changed**
    - `ApproversPath` **Added**
    - `ResourceBindingsPath` **Added**
    - `Task.ApproverIds` **Added**
    - `Task.Form.FieldRelationships[].DependentOn` **Added**
    - `Task.Form.Fields[].Required` **Added**
    - `Task.Form.Fields[].StringField.PickerField.C1UserFilter` **Added**
    - `Task.Form.Fields[].StringField.TextField.Suffix` **Added**
    - `Task.Form.Fields[].StringMapField` **Added**
    - `Task.Origin.Enum(taskOriginCascadeDelete)` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Action.ActionTargetBatonResourceAction` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ActionInstance.ActionTargetBatonResourceActionInstance` **Added**
    - `Task.RevocationTargets` **Added**
    - `Task.TaskType.TaskTypeFinding` **Added**
    - `Task.TaskType.TaskTypeGrant.TaskGrantSource.IsExtension` **Added**
* `ConductoroneApi.TaskActions.ApproveWithStepUp()`: `response.TaskView` **Changed**
    - `ApproversPath` **Added**
    - `ResourceBindingsPath` **Added**
    - `Task.ApproverIds` **Added**
    - `Task.Form.FieldRelationships[].DependentOn` **Added**
    - `Task.Form.Fields[].Required` **Added**
    - `Task.Form.Fields[].StringField.PickerField.C1UserFilter` **Added**
    - `Task.Form.Fields[].StringField.TextField.Suffix` **Added**
    - `Task.Form.Fields[].StringMapField` **Added**
    - `Task.Origin.Enum(taskOriginCascadeDelete)` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Action.ActionTargetBatonResourceAction` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ActionInstance.ActionTargetBatonResourceActionInstance` **Added**
    - `Task.RevocationTargets` **Added**
    - `Task.TaskType.TaskTypeFinding` **Added**
    - `Task.TaskType.TaskTypeGrant.TaskGrantSource.IsExtension` **Added**
* `ConductoroneApi.TaskActions.Close()`: `response.TaskView` **Changed**
    - `ApproversPath` **Added**
    - `ResourceBindingsPath` **Added**
    - `Task.ApproverIds` **Added**
    - `Task.Form.FieldRelationships[].DependentOn` **Added**
    - `Task.Form.Fields[].Required` **Added**
    - `Task.Form.Fields[].StringField.PickerField.C1UserFilter` **Added**
    - `Task.Form.Fields[].StringField.TextField.Suffix` **Added**
    - `Task.Form.Fields[].StringMapField` **Added**
    - `Task.Origin.Enum(taskOriginCascadeDelete)` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Action.ActionTargetBatonResourceAction` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ActionInstance.ActionTargetBatonResourceActionInstance` **Added**
    - `Task.RevocationTargets` **Added**
    - `Task.TaskType.TaskTypeFinding` **Added**
    - `Task.TaskType.TaskTypeGrant.TaskGrantSource.IsExtension` **Added**
* `ConductoroneApi.TaskActions.Comment()`: `response.TaskView` **Changed**
    - `ApproversPath` **Added**
    - `ResourceBindingsPath` **Added**
    - `Task.ApproverIds` **Added**
    - `Task.Form.FieldRelationships[].DependentOn` **Added**
    - `Task.Form.Fields[].Required` **Added**
    - `Task.Form.Fields[].StringField.PickerField.C1UserFilter` **Added**
    - `Task.Form.Fields[].StringField.TextField.Suffix` **Added**
    - `Task.Form.Fields[].StringMapField` **Added**
    - `Task.Origin.Enum(taskOriginCascadeDelete)` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Action.ActionTargetBatonResourceAction` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ActionInstance.ActionTargetBatonResourceActionInstance` **Added**
    - `Task.RevocationTargets` **Added**
    - `Task.TaskType.TaskTypeFinding` **Added**
    - `Task.TaskType.TaskTypeGrant.TaskGrantSource.IsExtension` **Added**
* `ConductoroneApi.TaskActions.Deny()`: `response.TaskView` **Changed**
    - `ApproversPath` **Added**
    - `ResourceBindingsPath` **Added**
    - `Task.ApproverIds` **Added**
    - `Task.Form.FieldRelationships[].DependentOn` **Added**
    - `Task.Form.Fields[].Required` **Added**
    - `Task.Form.Fields[].StringField.PickerField.C1UserFilter` **Added**
    - `Task.Form.Fields[].StringField.TextField.Suffix` **Added**
    - `Task.Form.Fields[].StringMapField` **Added**
    - `Task.Origin.Enum(taskOriginCascadeDelete)` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Action.ActionTargetBatonResourceAction` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ActionInstance.ActionTargetBatonResourceActionInstance` **Added**
    - `Task.RevocationTargets` **Added**
    - `Task.TaskType.TaskTypeFinding` **Added**
    - `Task.TaskType.TaskTypeGrant.TaskGrantSource.IsExtension` **Added**
* `ConductoroneApi.TaskActions.EscalateToEmergencyAccess()`: `response.TaskView` **Changed**
    - `ApproversPath` **Added**
    - `ResourceBindingsPath` **Added**
    - `Task.ApproverIds` **Added**
    - `Task.Form.FieldRelationships[].DependentOn` **Added**
    - `Task.Form.Fields[].Required` **Added**
    - `Task.Form.Fields[].StringField.PickerField.C1UserFilter` **Added**
    - `Task.Form.Fields[].StringField.TextField.Suffix` **Added**
    - `Task.Form.Fields[].StringMapField` **Added**
    - `Task.Origin.Enum(taskOriginCascadeDelete)` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Action.ActionTargetBatonResourceAction` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ActionInstance.ActionTargetBatonResourceActionInstance` **Added**
    - `Task.RevocationTargets` **Added**
    - `Task.TaskType.TaskTypeFinding` **Added**
    - `Task.TaskType.TaskTypeGrant.TaskGrantSource.IsExtension` **Added**
* `ConductoroneApi.TaskActions.HardReset()`: `response.TaskView` **Changed**
    - `ApproversPath` **Added**
    - `ResourceBindingsPath` **Added**
    - `Task.ApproverIds` **Added**
    - `Task.Form.FieldRelationships[].DependentOn` **Added**
    - `Task.Form.Fields[].Required` **Added**
    - `Task.Form.Fields[].StringField.PickerField.C1UserFilter` **Added**
    - `Task.Form.Fields[].StringField.TextField.Suffix` **Added**
    - `Task.Form.Fields[].StringMapField` **Added**
    - `Task.Origin.Enum(taskOriginCascadeDelete)` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Action.ActionTargetBatonResourceAction` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ActionInstance.ActionTargetBatonResourceActionInstance` **Added**
    - `Task.RevocationTargets` **Added**
    - `Task.TaskType.TaskTypeFinding` **Added**
    - `Task.TaskType.TaskTypeGrant.TaskGrantSource.IsExtension` **Added**
* `ConductoroneApi.TaskActions.ProcessNow()`: `response.TaskView` **Changed**
    - `ApproversPath` **Added**
    - `ResourceBindingsPath` **Added**
    - `Task.ApproverIds` **Added**
    - `Task.Form.FieldRelationships[].DependentOn` **Added**
    - `Task.Form.Fields[].Required` **Added**
    - `Task.Form.Fields[].StringField.PickerField.C1UserFilter` **Added**
    - `Task.Form.Fields[].StringField.TextField.Suffix` **Added**
    - `Task.Form.Fields[].StringMapField` **Added**
    - `Task.Origin.Enum(taskOriginCascadeDelete)` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Action.ActionTargetBatonResourceAction` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ActionInstance.ActionTargetBatonResourceActionInstance` **Added**
    - `Task.RevocationTargets` **Added**
    - `Task.TaskType.TaskTypeFinding` **Added**
    - `Task.TaskType.TaskTypeGrant.TaskGrantSource.IsExtension` **Added**
* `ConductoroneApi.TaskActions.Reassign()`: `response.TaskView` **Changed**
    - `ApproversPath` **Added**
    - `ResourceBindingsPath` **Added**
    - `Task.ApproverIds` **Added**
    - `Task.Form.FieldRelationships[].DependentOn` **Added**
    - `Task.Form.Fields[].Required` **Added**
    - `Task.Form.Fields[].StringField.PickerField.C1UserFilter` **Added**
    - `Task.Form.Fields[].StringField.TextField.Suffix` **Added**
    - `Task.Form.Fields[].StringMapField` **Added**
    - `Task.Origin.Enum(taskOriginCascadeDelete)` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Action.ActionTargetBatonResourceAction` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ActionInstance.ActionTargetBatonResourceActionInstance` **Added**
    - `Task.RevocationTargets` **Added**
    - `Task.TaskType.TaskTypeFinding` **Added**
    - `Task.TaskType.TaskTypeGrant.TaskGrantSource.IsExtension` **Added**
* `ConductoroneApi.TaskActions.Restart()`: `response.TaskView` **Changed**
    - `ApproversPath` **Added**
    - `ResourceBindingsPath` **Added**
    - `Task.ApproverIds` **Added**
    - `Task.Form.FieldRelationships[].DependentOn` **Added**
    - `Task.Form.Fields[].Required` **Added**
    - `Task.Form.Fields[].StringField.PickerField.C1UserFilter` **Added**
    - `Task.Form.Fields[].StringField.TextField.Suffix` **Added**
    - `Task.Form.Fields[].StringMapField` **Added**
    - `Task.Origin.Enum(taskOriginCascadeDelete)` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Action.ActionTargetBatonResourceAction` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ActionInstance.ActionTargetBatonResourceActionInstance` **Added**
    - `Task.RevocationTargets` **Added**
    - `Task.TaskType.TaskTypeFinding` **Added**
    - `Task.TaskType.TaskTypeGrant.TaskGrantSource.IsExtension` **Added**
* `ConductoroneApi.TaskActions.SkipStep()`: `response.TaskView` **Changed**
    - `ApproversPath` **Added**
    - `ResourceBindingsPath` **Added**
    - `Task.ApproverIds` **Added**
    - `Task.Form.FieldRelationships[].DependentOn` **Added**
    - `Task.Form.Fields[].Required` **Added**
    - `Task.Form.Fields[].StringField.PickerField.C1UserFilter` **Added**
    - `Task.Form.Fields[].StringField.TextField.Suffix` **Added**
    - `Task.Form.Fields[].StringMapField` **Added**
    - `Task.Origin.Enum(taskOriginCascadeDelete)` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Action.ActionTargetBatonResourceAction` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ActionInstance.ActionTargetBatonResourceActionInstance` **Added**
    - `Task.RevocationTargets` **Added**
    - `Task.TaskType.TaskTypeFinding` **Added**
    - `Task.TaskType.TaskTypeGrant.TaskGrantSource.IsExtension` **Added**
* `ConductoroneApi.TaskActions.UpdateGrantDuration()`: `response.TaskView` **Changed**
    - `ApproversPath` **Added**
    - `ResourceBindingsPath` **Added**
    - `Task.ApproverIds` **Added**
    - `Task.Form.FieldRelationships[].DependentOn` **Added**
    - `Task.Form.Fields[].Required` **Added**
    - `Task.Form.Fields[].StringField.PickerField.C1UserFilter` **Added**
    - `Task.Form.Fields[].StringField.TextField.Suffix` **Added**
    - `Task.Form.Fields[].StringMapField` **Added**
    - `Task.Origin.Enum(taskOriginCascadeDelete)` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Action.ActionTargetBatonResourceAction` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ActionInstance.ActionTargetBatonResourceActionInstance` **Added**
    - `Task.RevocationTargets` **Added**
    - `Task.TaskType.TaskTypeFinding` **Added**
    - `Task.TaskType.TaskTypeGrant.TaskGrantSource.IsExtension` **Added**
* `ConductoroneApi.TaskActions.UpdateRequestData()`: `response.TaskView` **Changed**
    - `ApproversPath` **Added**
    - `ResourceBindingsPath` **Added**
    - `Task.ApproverIds` **Added**
    - `Task.Form.FieldRelationships[].DependentOn` **Added**
    - `Task.Form.Fields[].Required` **Added**
    - `Task.Form.Fields[].StringField.PickerField.C1UserFilter` **Added**
    - `Task.Form.Fields[].StringField.TextField.Suffix` **Added**
    - `Task.Form.Fields[].StringMapField` **Added**
    - `Task.Origin.Enum(taskOriginCascadeDelete)` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Action.ActionTargetBatonResourceAction` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ActionInstance.ActionTargetBatonResourceActionInstance` **Added**
    - `Task.RevocationTargets` **Added**
    - `Task.TaskType.TaskTypeFinding` **Added**
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
