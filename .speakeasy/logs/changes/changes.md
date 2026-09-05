## Go SDK Changes:
* `ConductoroneApi.AccessReviewTemplate.Create()`: 
  * `request.Request` **Changed** (Breaking ⚠️)
    - `ColumnConfig.OrderedColumns` **Added**
    - `RecurrenceRule.Frequency` **Changed** (Breaking ⚠️)
  * `response.AccessReviewTemplate` **Changed**
    - `ColumnConfig.OrderedColumns` **Added**
    - `MsTeamsChannel` **Added**
    - `RecurrenceRule.Frequency` **Changed**
* `ConductoroneApi.AccessReviewTemplate.Update()`: 
  * `request.Request.AccessReviewTemplateServiceUpdateRequest.AccessReviewTemplate` **Changed** (Breaking ⚠️)
    - `ColumnConfig.OrderedColumns` **Added**
    - `MsTeamsChannel` **Added**
    - `RecurrenceRule.Frequency` **Changed** (Breaking ⚠️)
  * `response.AccessReviewTemplate` **Changed**
    - `ColumnConfig.OrderedColumns` **Added**
    - `MsTeamsChannel` **Added**
    - `RecurrenceRule.Frequency` **Changed**
* `ConductoroneApi.A2Ui.GetSurfaceProvenance()`: **Added**
* `ConductoroneApi.AccessReviewReport.List()`: **Added**
* `ConductoroneApi.AccessReviewActions.GenerateReport()`: **Added**
* `ConductoroneApi.McpResource.Get()`: **Added**
* `ConductoroneApi.McpResource.List()`: **Added**
* `ConductoroneApi.McpResource.ListHistory()`: **Added**
* `ConductoroneApi.McpResource.Search()`: **Added**
* `ConductoroneApi.McpResource.Update()`: **Added**
* `ConductoroneApi.McpAccessProfile.SearchAccessProfiles()`: **Added**
* `ConductoroneApi.AppEntitlementSearch.SearchReachableResourcesForUser()`: **Added**
* `ConductoroneApi.AppManagedState.Get()`: **Added**
* `ConductoroneApi.AppManagedState.List()`: **Added**
* `ConductoroneApi.AppManagedState.Promote()`: **Added**
* `ConductoroneApi.SsoApplication.BatchDeleteSubjectCompatibility()`: **Added**
* `ConductoroneApi.SsoApplication.BatchImportSubjectCompatibility()`: **Added**
* `ConductoroneApi.SsoApplication.Create()`: **Added**
* `ConductoroneApi.SsoApplication.CreateClient()`: **Added**
* `ConductoroneApi.SsoApplication.Delete()`: **Added**
* `ConductoroneApi.SsoApplication.DeleteClient()`: **Added**
* `ConductoroneApi.SsoApplication.Get()`: **Added**
* `ConductoroneApi.SsoApplication.List()`: **Added**
* `ConductoroneApi.SsoApplication.ListClients()`: **Added**
* `ConductoroneApi.SsoApplication.ListHistory()`: **Added**
* `ConductoroneApi.SsoApplication.ParseSamlServiceProviderMetadata()`: **Added**
* `ConductoroneApi.SsoApplication.RotateClientSecret()`: **Added**
* `ConductoroneApi.SsoApplication.Search()`: **Added**
* `ConductoroneApi.SsoApplication.Update()`: **Added**
* `ConductoroneApi.SsoApplication.UpdateClient()`: **Added**
* `ConductoroneApi.RequestCatalogManagement.PlanTypeChange()`: **Added**
* `ConductoroneApi.UiConversations.EnsureOnboardingSession()`: **Added**
* `ConductoroneApi.Feedback.CreateFeedback()`: **Added**
* `ConductoroneApi.Finding.UpdateFindingAssignee()`: **Added**
* `ConductoroneApi.FindingSettings.ListFindingSettings()`: **Added**
* `ConductoroneApi.FindingSettings.UpdateFindingSettings()`: **Added**
* `ConductoroneApi.FunctionsInvocation.GetResultDownloadUrl()`: **Added**
* `ConductoroneApi.AppCap.Delete()`: **Added**
* `ConductoroneApi.AppCap.Get()`: **Added**
* `ConductoroneApi.AppCap.List()`: **Added**
* `ConductoroneApi.AppCap.ListHistory()`: **Added**
* `ConductoroneApi.AppCap.SetLimit()`: **Added**
* `ConductoroneApi.AppCap.Suspend()`: **Added**
* `ConductoroneApi.AppCap.Unsuspend()`: **Added**
* `ConductoroneApi.FundAssignment.ClearExtension()`: **Added**
* `ConductoroneApi.FundAssignment.Delete()`: **Added**
* `ConductoroneApi.FundAssignment.Get()`: **Added**
* `ConductoroneApi.FundAssignment.GrantExtension()`: **Added**
* `ConductoroneApi.FundAssignment.ListHistory()`: **Added**
* `ConductoroneApi.FundAssignment.Search()`: **Added**
* `ConductoroneApi.FundAssignment.SetLimit()`: **Added**
* `ConductoroneApi.FundAssignment.Suspend()`: **Added**
* `ConductoroneApi.FundAssignment.Unsuspend()`: **Added**
* `ConductoroneApi.MyFundLimits.Delete()`: **Added**
* `ConductoroneApi.MyFundLimits.List()`: **Added**
* `ConductoroneApi.MyFundLimits.ListHistory()`: **Added**
* `ConductoroneApi.MyFundLimits.Pause()`: **Added**
* `ConductoroneApi.MyFundLimits.Resume()`: **Added**
* `ConductoroneApi.MyFundLimits.SetLimit()`: **Added**
* `ConductoroneApi.FundPolicy.Create()`: **Added**
* `ConductoroneApi.FundPolicy.Delete()`: **Added**
* `ConductoroneApi.FundPolicy.FreezeTenant()`: **Added**
* `ConductoroneApi.FundPolicy.Get()`: **Added**
* `ConductoroneApi.FundPolicy.ListHistory()`: **Added**
* `ConductoroneApi.FundPolicy.SetOrgCeiling()`: **Added**
* `ConductoroneApi.FundPolicy.UnfreezeTenant()`: **Added**
* `ConductoroneApi.FundPolicy.Update()`: **Added**
* `ConductoroneApi.FundRule.Create()`: **Added**
* `ConductoroneApi.FundRule.Delete()`: **Added**
* `ConductoroneApi.FundRule.Get()`: **Added**
* `ConductoroneApi.FundRule.List()`: **Added**
* `ConductoroneApi.FundRule.ListHistory()`: **Added**
* `ConductoroneApi.FundRule.Search()`: **Added**
* `ConductoroneApi.FundRule.Update()`: **Added**
* `ConductoroneApi.SubjectAppLimit.Delete()`: **Added**
* `ConductoroneApi.SubjectAppLimit.Get()`: **Added**
* `ConductoroneApi.SubjectAppLimit.ListHistory()`: **Added**
* `ConductoroneApi.SubjectAppLimit.Search()`: **Added**
* `ConductoroneApi.SubjectAppLimit.SetLimit()`: **Added**
* `ConductoroneApi.SubjectAppLimit.Suspend()`: **Added**
* `ConductoroneApi.SubjectAppLimit.Unsuspend()`: **Added**
* `ConductoroneApi.GatewayKey.List()`: **Added**
* `ConductoroneApi.GatewayKey.Mint()`: **Added**
* `ConductoroneApi.GatewayKey.Revoke()`: **Added**
* `ConductoroneApi.ProviderCredential.Clear()`: **Added**
* `ConductoroneApi.ProviderCredential.Get()`: **Added**
* `ConductoroneApi.ProviderCredential.Set()`: **Added**
* `ConductoroneApi.Reporting.Delete()`: **Added**
* `ConductoroneApi.Reporting.Get()`: **Added**
* `ConductoroneApi.Reporting.GetProgram()`: **Added**
* `ConductoroneApi.Reporting.GetRunProvenance()`: **Added**
* `ConductoroneApi.Reporting.List()`: **Added**
* `ConductoroneApi.Reporting.Run()`: **Added**
* `ConductoroneApi.Reporting.Save()`: **Added**
* `ConductoroneApi.Reporting.Update()`: **Added**
* `ConductoroneApi.RoleMiningManagement.EvaluateEntitlementSelection()`: **Added**
* `ConductoroneApi.PaperSecret.SearchSecretsSharedWithMe()`: **Added**
* `ConductoroneApi.SessionPolicy.GetEffectiveSessionPolicy()`: **Added**
* `ConductoroneApi.SessionPolicy.ListUserPolicies()`: **Added**
* `ConductoroneApi.SessionPolicy.SearchPolicyUsers()`: **Added**
* `ConductoroneApi.SsoSettings.Get()`: **Added**
* `ConductoroneApi.SsoSettings.ListHistory()`: **Added**
* `ConductoroneApi.SsoSettings.Update()`: **Added**
* `ConductoroneApi.TaskActions.RetryProvisioning()`: **Added**
* `ConductoroneApi.TbControlPlane.GetDiscoverySnapshot()`: **Added**
* `ConductoroneApi.TbControlPlane.GetEgressPolicy()`: **Added**
* `ConductoroneApi.TbControlPlane.PushDiscovery()`: **Added**
* `ConductoroneApi.TbControlPlane.SaveEgressPolicy()`: **Added**
* `ConductoroneApi.A2Ui.ListSurfaces()`: `response.Surfaces[]` **Changed**
    - `Components[].C1MetricCards` **Added**
    - `Components[].C1Table` **Added**
    - `HasReportProgram` **Added**
    - `ReportEditTarget` **Added**
* `ConductoroneApi.AccessReview.Create()`:  `response.AccessReview.AccessReview.ColumnConfig.OrderedColumns` **Added**
* `ConductoroneApi.AccessReview.Get()`:  `response.AccessReview.AccessReview.ColumnConfig.OrderedColumns` **Added**
* `ConductoroneApi.AccessReview.List()`:  `response.List[].AccessReview.ColumnConfig.OrderedColumns` **Added**
* `ConductoroneApi.AccessReview.Update()`: 
  *  `request.Request.AccessReviewServiceUpdateRequest.AccessReview.ColumnConfig.OrderedColumns` **Added**
  *  `response.AccessReview.AccessReview.ColumnConfig.OrderedColumns` **Added**
* `ConductoroneApi.AccessReviewTemplate.Get()`: `response.AccessReviewTemplate` **Changed**
    - `ColumnConfig.OrderedColumns` **Added**
    - `MsTeamsChannel` **Added**
    - `RecurrenceRule.Frequency` **Changed**
* `ConductoroneApi.AppUser.List()`: `response.List[].AppUser` **Changed**
    - `AgentStatus` **Added**
    - `NhiDetail` **Added**
    - `NhiType` **Added**
* `ConductoroneApi.AppUser.ListAppUsersForUser()`: `response.List[].AppUser` **Changed**
    - `AgentStatus` **Added**
    - `NhiDetail` **Added**
    - `NhiType` **Added**
* `ConductoroneApi.AppUser.ListOwnedServiceAccounts()`: `response.List[].AppUser` **Changed**
    - `AgentStatus` **Added**
    - `NhiDetail` **Added**
    - `NhiType` **Added**
* `ConductoroneApi.AppUser.Search()`: 
  * `request.Request` **Changed**
    - `AgentStatuses` **Added**
    - `NhiTypes` **Added**
  * `response.List[].AppUser` **Changed**
    - `AgentStatus` **Added**
    - `NhiDetail` **Added**
    - `NhiType` **Added**
* `ConductoroneApi.AppUser.Update()`: `response.AppUserView.AppUser` **Changed**
    - `AgentStatus` **Added**
    - `NhiDetail` **Added**
    - `NhiType` **Added**
* `ConductoroneApi.Apps.Create()`: 
  *  `request.Request.MatchBatonRef` **Added**
  *  `response.App.MatchBatonRef` **Added**
* `ConductoroneApi.Apps.Get()`:  `response.App.MatchBatonRef` **Added**
* `ConductoroneApi.Apps.List()`:  `response.List[].MatchBatonRef` **Added**
* `ConductoroneApi.Apps.Update()`: 
  *  `request.Request.UpdateAppRequest.App.MatchBatonRef` **Added**
  *  `response.App.MatchBatonRef` **Added**
* `ConductoroneApi.McpTool.Get()`: `response.Tool` **Changed**
    - `RequestableViaToolset` **Added**
    - `Requestable` **Added**
* `ConductoroneApi.McpTool.List()`: `response.Tools[]` **Changed**
    - `RequestableViaToolset` **Added**
    - `Requestable` **Added**
* `ConductoroneApi.McpTool.ListHistory()`: `response.List[].Snapshot` **Changed**
    - `RequestableViaToolset` **Added**
    - `Requestable` **Added**
* `ConductoroneApi.McpTool.Search()`: 
  *  `request.Request.McpToolServiceSearchRequest.IncludeRequestable` **Added**
  * `response.List[]` **Changed**
    - `RequestableViaToolset` **Added**
    - `Requestable` **Added**
* `ConductoroneApi.McpTool.Update()`: `response.Tool` **Changed**
    - `RequestableViaToolset` **Added**
    - `Requestable` **Added**
* `ConductoroneApi.McpAccessProfile.Create()`: `response.Profile` **Changed**
    - `ConnectorDisplayName` **Added**
    - `Requestable` **Added**
* `ConductoroneApi.McpAccessProfile.Get()`: `response.Profile` **Changed**
    - `ConnectorDisplayName` **Added**
    - `Requestable` **Added**
* `ConductoroneApi.McpAccessProfile.GetByAppEntitlementId()`: `response.Profile` **Changed**
    - `ConnectorDisplayName` **Added**
    - `Requestable` **Added**
* `ConductoroneApi.McpAccessProfile.List()`: `response.Profiles[]` **Changed**
    - `ConnectorDisplayName` **Added**
    - `Requestable` **Added**
* `ConductoroneApi.McpAccessProfile.Update()`: `response.Profile` **Changed**
    - `ConnectorDisplayName` **Added**
    - `Requestable` **Added**
* `ConductoroneApi.McpAccessProfileToolBinding.GetAccessProfilesForTools()`: `response.AccessProfilesForTools[].AccessProfiles[]` **Changed**
    - `ConnectorDisplayName` **Added**
    - `Requestable` **Added**
* `ConductoroneApi.AppEntitlements.Create()`: 
  * `request.Request.CreateAppEntitlementRequest.ProvisionPolicy` **Changed**
    - `DevicePlacement` **Added**
    - `MultiStep.ProvisionSteps[].DevicePlacement` **Added**
  *  `response.AppEntitlementView.AppEntitlement.DeprovisionerPolicy.DevicePlacement` **Added**
* `ConductoroneApi.AppEntitlements.Get()`:  `response.AppEntitlementView.AppEntitlement.DeprovisionerPolicy.DevicePlacement` **Added**
* `ConductoroneApi.AppEntitlements.List()`: 
  * `request.Request` **Changed**
    - `AppUserId` **Added**
    - `Q` **Added**
  *  `response.List[].AppEntitlement.DeprovisionerPolicy.DevicePlacement` **Added**
* `ConductoroneApi.AppEntitlements.ListForAppResource()`:  `response.List[].AppEntitlement.DeprovisionerPolicy.DevicePlacement` **Added**
* `ConductoroneApi.AppEntitlements.ListForAppUser()`:  `response.List[].AppEntitlement.DeprovisionerPolicy.DevicePlacement` **Added**
* `ConductoroneApi.AppEntitlements.ListUsers()`: `response.List[].AppUser.AppUser` **Changed**
    - `AgentStatus` **Added**
    - `NhiDetail` **Added**
    - `NhiType` **Added**
* `ConductoroneApi.AppEntitlements.Update()`: 
  * `request.Request.UpdateAppEntitlementRequest.Entitlement.DeprovisionerPolicy` **Changed**
    - `DevicePlacement` **Added**
    - `MultiStep.ProvisionSteps[].DevicePlacement` **Added**
  *  `response.AppEntitlementView.AppEntitlement.DeprovisionerPolicy.DevicePlacement` **Added**
* `ConductoroneApi.AppEntitlementSearch.Search()`:  `response.List[].AppEntitlement.DeprovisionerPolicy.DevicePlacement` **Added**
* `ConductoroneApi.AppEntitlementSearch.SearchAppEntitlementsForAppUser()`:  `response.List[].AppEntitlement.DeprovisionerPolicy.DevicePlacement` **Added**
* `ConductoroneApi.AppEntitlementSearch.SearchAppEntitlementsWithExpired()`: `response.List[].AppUser` **Changed**
    - `AgentStatus` **Added**
    - `NhiDetail` **Added**
    - `NhiType` **Added**
* `ConductoroneApi.AppEntitlementSearch.SearchGrants()`: `response.List[]` **Changed**
    - `AppEntitlementUserBinding.AppUser.AppUser.AgentStatus` **Added**
    - `AppEntitlementUserBinding.AppUser.AppUser.NhiDetail` **Added**
    - `AppEntitlementUserBinding.AppUser.AppUser.NhiType` **Added**
    - `Entitlement.AppEntitlement.DeprovisionerPolicy.DevicePlacement` **Added**
* `ConductoroneApi.AppEntitlementUserBinding.SearchPastGrants()`:  `response.List[].History.Id` **Added**
* `ConductoroneApi.McpServer.Get()`:  `response.McpServer.EndpointUrlLocked` **Added**
* `ConductoroneApi.McpServer.GetCatalog()`: `response.CatalogEntry` **Changed**
    - `AuthModes[].ClientIdMode` **Added**
    - `AuthModes[].OptionalScopes` **Added**
    - `DefaultToolPrefix` **Added**
* `ConductoroneApi.McpServer.List()`: 
  *  `request.Request.Query` **Added**
  *  `response.List[].EndpointUrlLocked` **Added**
* `ConductoroneApi.McpServer.ListCatalog()`: `response.List[]` **Changed**
    - `AuthModes[].ClientIdMode` **Added**
    - `AuthModes[].OptionalScopes` **Added**
    - `DefaultToolPrefix` **Added**
* `ConductoroneApi.McpServer.Register()`: 
  *  `request.Request.McpServerServiceRegisterRequest.AccessProfileIds` **Added**
  * `response` **Changed**
    - `AccessProfilesAttached` **Added**
    - `McpServer.EndpointUrlLocked` **Added**
* `ConductoroneApi.McpServer.SearchWithToolCount()`:  `response.List[].McpServer.EndpointUrlLocked` **Added**
* `ConductoroneApi.McpServer.Update()`:  `response.McpServer.EndpointUrlLocked` **Added**
* `ConductoroneApi.McpServer.UpdateCredentials()`:  `response.McpServer.EndpointUrlLocked` **Added**
* `ConductoroneApi.AppResourceType.CreateManuallyManagedResourceType()`: 
  *  `request.Request.CreateManuallyManagedResourceTypeRequest.ResourceType.Enum(clawAgent)` **Added**
* `ConductoroneApi.Auth.Introspect()`:  `response.DisabledModules` **Added**
* `ConductoroneApi.Automation.CreateAutomation()`: 
  *  `request.Request.AutomationSteps[].CreateRevokeTasksV2.GrantSourceFilter` **Added**
  *  `response.Automation.AutomationSteps[].CreateRevokeTasksV2.GrantSourceFilter` **Added**
* `ConductoroneApi.Automation.GetAutomation()`:  `response.Automation.AutomationSteps[].CreateRevokeTasksV2.GrantSourceFilter` **Added**
* `ConductoroneApi.Automation.ListAutomations()`:  `response.List[].AutomationSteps[].CreateRevokeTasksV2.GrantSourceFilter` **Added**
* `ConductoroneApi.Automation.UpdateAutomation()`: 
  *  `request.Request.UpdateAutomationRequest.Automation.AutomationSteps[].CreateRevokeTasksV2.GrantSourceFilter` **Added**
  *  `response.Automation.AutomationSteps[].CreateRevokeTasksV2.GrantSourceFilter` **Added**
* `ConductoroneApi.RequestCatalogManagement.Create()`: 
  *  `request.Request.Type` **Added**
  * `response.RequestCatalogView.RequestCatalog` **Changed**
    - `AccessEntitlements[].DeprovisionerPolicy.DevicePlacement` **Added**
    - `Type` **Added**
* `ConductoroneApi.RequestCatalogManagement.Get()`: `response.RequestCatalogView.RequestCatalog` **Changed**
    - `AccessEntitlements[].DeprovisionerPolicy.DevicePlacement` **Added**
    - `Type` **Added**
* `ConductoroneApi.RequestCatalogManagement.List()`: `response.List[].RequestCatalog` **Changed**
    - `AccessEntitlements[].DeprovisionerPolicy.DevicePlacement` **Added**
    - `Type` **Added**
* `ConductoroneApi.RequestCatalogManagement.ListEntitlementsForAccess()`:  `response.List[].AppEntitlement.DeprovisionerPolicy.DevicePlacement` **Added**
* `ConductoroneApi.RequestCatalogManagement.ListEntitlementsPerCatalog()`:  `response.List[].AppEntitlement.DeprovisionerPolicy.DevicePlacement` **Added**
* `ConductoroneApi.RequestCatalogManagement.Update()`: 
  * `request.Request.RequestCatalogManagementServiceUpdateRequest.Catalog` **Changed**
    - `AccessEntitlements[].DeprovisionerPolicy.DevicePlacement` **Added**
    - `AccessEntitlements[].DeprovisionerPolicy.MultiStep.ProvisionSteps[].DevicePlacement` **Added**
    - `Type` **Added**
  * `response.RequestCatalogView.RequestCatalog` **Changed**
    - `AccessEntitlements[].DeprovisionerPolicy.DevicePlacement` **Added**
    - `Type` **Added**
* `ConductoroneApi.ConnectorCatalog.ConfigurationSchema()`: `response.FormSchema.Fields[].StringField` **Changed**
    - `DateField` **Added**
    - `PickerField.C1UserPicker.ExcludeUserIds` **Added**
    - `PickerField.C1UserPicker.IncludeDeactivated` **Added**
    - `PickerField.C1UserPicker.UserIds` **Added**
* `ConductoroneApi.Finding.BulkCreateFindingTasks()`: 
  * `request.Request.SearchRequest` **Changed**
    - `AssigneeIdentityUserIds` **Added**
    - `FindingTypes[].Enum(findingTypeConnectorSyncFailing)` **Added**
    - `FindingTypes[].Enum(findingTypeCredentialExpiring)` **Added**
    - `FindingTypes[].Enum(findingTypeCredentialPubliclyExposed)` **Added**
    - `FindingTypes[].Enum(findingTypeDeactivatedOwner)` **Added**
    - `FindingTypes[].Enum(findingTypeDecoyPubliclyExposed)` **Added**
    - `FindingTypes[].Enum(findingTypeShadowMcp)` **Added**
    - `FindingTypes[].Enum(findingTypeUnusedSecret)` **Added**
* `ConductoroneApi.Finding.BulkUpdateFindingState()`: `request.Request` **Changed**
    - `AssignOwner.AssigneeIdentityUserId` **Added**
    - `Reprocess` **Added**
    - `SearchRequest.AssigneeIdentityUserIds` **Added**
    - `SearchRequest.FindingTypes[].Enum(findingTypeConnectorSyncFailing)` **Added**
    - `SearchRequest.FindingTypes[].Enum(findingTypeCredentialExpiring)` **Added**
    - `SearchRequest.FindingTypes[].Enum(findingTypeCredentialPubliclyExposed)` **Added**
    - `SearchRequest.FindingTypes[].Enum(findingTypeDeactivatedOwner)` **Added**
    - `SearchRequest.FindingTypes[].Enum(findingTypeDecoyPubliclyExposed)` **Added**
    - `SearchRequest.FindingTypes[].Enum(findingTypeShadowMcp)` **Added**
    - `SearchRequest.FindingTypes[].Enum(findingTypeUnusedSecret)` **Added**
* `ConductoroneApi.Finding.CreateFinding()`: `response.Finding` **Changed**
    - `Annotations` **Added**
    - `AssigneeIdentityUserId` **Added**
    - `ConnectorSyncFailingEvidence` **Added**
    - `ConnectorSyncFailing` **Added**
    - `CredentialExpiringEvidence` **Added**
    - `CredentialExpiring` **Added**
    - `CredentialPubliclyExposedEvidence` **Added**
    - `CredentialPubliclyExposed` **Added**
    - `DeactivatedOwnerEvidence` **Added**
    - `DeactivatedOwner` **Added**
    - `DecoyPubliclyExposedEvidence` **Added**
    - `DecoyPubliclyExposed` **Added**
    - `ObjectPermissions` **Added**
    - `ShadowMcpEvidence` **Added**
    - `ShadowMcp` **Added**
    - `UnusedSecretEvidence` **Added**
    - `UnusedSecret` **Added**
* `ConductoroneApi.Finding.CreateFindingTask()`: `response.Finding` **Changed**
    - `Annotations` **Added**
    - `AssigneeIdentityUserId` **Added**
    - `ConnectorSyncFailingEvidence` **Added**
    - `ConnectorSyncFailing` **Added**
    - `CredentialExpiringEvidence` **Added**
    - `CredentialExpiring` **Added**
    - `CredentialPubliclyExposedEvidence` **Added**
    - `CredentialPubliclyExposed` **Added**
    - `DeactivatedOwnerEvidence` **Added**
    - `DeactivatedOwner` **Added**
    - `DecoyPubliclyExposedEvidence` **Added**
    - `DecoyPubliclyExposed` **Added**
    - `ObjectPermissions` **Added**
    - `ShadowMcpEvidence` **Added**
    - `ShadowMcp` **Added**
    - `UnusedSecretEvidence` **Added**
    - `UnusedSecret` **Added**
* `ConductoroneApi.Finding.GetFinding()`: `response.Finding` **Changed**
    - `Annotations` **Added**
    - `AssigneeIdentityUserId` **Added**
    - `ConnectorSyncFailingEvidence` **Added**
    - `ConnectorSyncFailing` **Added**
    - `CredentialExpiringEvidence` **Added**
    - `CredentialExpiring` **Added**
    - `CredentialPubliclyExposedEvidence` **Added**
    - `CredentialPubliclyExposed` **Added**
    - `DeactivatedOwnerEvidence` **Added**
    - `DeactivatedOwner` **Added**
    - `DecoyPubliclyExposedEvidence` **Added**
    - `DecoyPubliclyExposed` **Added**
    - `ObjectPermissions` **Added**
    - `ShadowMcpEvidence` **Added**
    - `ShadowMcp` **Added**
    - `UnusedSecretEvidence` **Added**
    - `UnusedSecret` **Added**
* `ConductoroneApi.Finding.UpdateFindingState()`: `response.Finding` **Changed**
    - `Annotations` **Added**
    - `AssigneeIdentityUserId` **Added**
    - `ConnectorSyncFailingEvidence` **Added**
    - `ConnectorSyncFailing` **Added**
    - `CredentialExpiringEvidence` **Added**
    - `CredentialExpiring` **Added**
    - `CredentialPubliclyExposedEvidence` **Added**
    - `CredentialPubliclyExposed` **Added**
    - `DeactivatedOwnerEvidence` **Added**
    - `DeactivatedOwner` **Added**
    - `DecoyPubliclyExposedEvidence` **Added**
    - `DecoyPubliclyExposed` **Added**
    - `ObjectPermissions` **Added**
    - `ShadowMcpEvidence` **Added**
    - `ShadowMcp` **Added**
    - `UnusedSecretEvidence` **Added**
    - `UnusedSecret` **Added**
* `ConductoroneApi.FindingRoutingRule.CreateFindingRoutingRule()`: 
  * `request.Request.RoutingRule` **Changed**
    - `Dispatchers` **Added**
    - `FindingType` **Added**
  * `response.RoutingRule` **Changed**
    - `Dispatchers` **Added**
    - `FindingType` **Added**
* `ConductoroneApi.FindingRoutingRule.GetFindingRoutingRule()`: `response.RoutingRule` **Changed**
    - `Dispatchers` **Added**
    - `FindingType` **Added**
* `ConductoroneApi.FindingRoutingRule.ListFindingRoutingRules()`: `response.List[]` **Changed**
    - `Dispatchers` **Added**
    - `FindingType` **Added**
* `ConductoroneApi.FindingRoutingRule.UpdateFindingRoutingRule()`: 
  * `request.Request.UpdateFindingRoutingRuleRequest.RoutingRule` **Changed**
    - `Dispatchers` **Added**
    - `FindingType` **Added**
  * `response.RoutingRule` **Changed**
    - `Dispatchers` **Added**
    - `FindingType` **Added**
* `ConductoroneApi.FindingSearch.Search()`: 
  * `request.Request` **Changed**
    - `AssigneeIdentityUserIds` **Added**
    - `FindingTypes[].Enum(findingTypeConnectorSyncFailing)` **Added**
    - `FindingTypes[].Enum(findingTypeCredentialExpiring)` **Added**
    - `FindingTypes[].Enum(findingTypeCredentialPubliclyExposed)` **Added**
    - `FindingTypes[].Enum(findingTypeDeactivatedOwner)` **Added**
    - `FindingTypes[].Enum(findingTypeDecoyPubliclyExposed)` **Added**
    - `FindingTypes[].Enum(findingTypeShadowMcp)` **Added**
    - `FindingTypes[].Enum(findingTypeUnusedSecret)` **Added**
  * `response.List[]` **Changed**
    - `Annotations` **Added**
    - `AssigneeIdentityUserId` **Added**
    - `ConnectorSyncFailingEvidence` **Added**
    - `ConnectorSyncFailing` **Added**
    - `CredentialExpiringEvidence` **Added**
    - `CredentialExpiring` **Added**
    - `CredentialPubliclyExposedEvidence` **Added**
    - `CredentialPubliclyExposed` **Added**
    - `DeactivatedOwnerEvidence` **Added**
    - `DeactivatedOwner` **Added**
    - `DecoyPubliclyExposedEvidence` **Added**
    - `DecoyPubliclyExposed` **Added**
    - `ObjectPermissions` **Added**
    - `ShadowMcpEvidence` **Added**
    - `ShadowMcp` **Added**
    - `UnusedSecretEvidence` **Added**
    - `UnusedSecret` **Added**
* `ConductoroneApi.FindingTransformationRule.CreateFindingTransformationRule()`: 
  *  `request.Request.TransformationRule.FindingType` **Added**
  *  `response.TransformationRule.FindingType` **Added**
* `ConductoroneApi.FindingTransformationRule.GetFindingTransformationRule()`:  `response.TransformationRule.FindingType` **Added**
* `ConductoroneApi.FindingTransformationRule.ListFindingTransformationRules()`:  `response.List[].FindingType` **Added**
* `ConductoroneApi.FindingTransformationRule.UpdateFindingTransformationRule()`: 
  *  `request.Request.UpdateFindingTransformationRuleRequest.TransformationRule.FindingType` **Added**
  *  `response.TransformationRule.FindingType` **Added**
* `ConductoroneApi.Functions.CreateFunction()`: 
  * `request.Request` **Changed**
    - `BrowserEnabled` **Added**
    - `FunctionType.Enum(functionTypeConnector)` **Added**
  * `response.Function` **Changed**
    - `BrowserEnabled` **Added**
    - `FunctionType.Enum(functionTypeConnector)` **Added**
    - `HookRefs` **Added**
    - `WorkflowTemplateRefs` **Added**
* `ConductoroneApi.Functions.GetFunction()`: `response.Function` **Changed**
    - `BrowserEnabled` **Added**
    - `FunctionType.Enum(functionTypeConnector)` **Added**
    - `HookRefs` **Added**
    - `WorkflowTemplateRefs` **Added**
* `ConductoroneApi.Functions.ListFunctions()`: `response.List[]` **Changed**
    - `BrowserEnabled` **Added**
    - `FunctionType.Enum(functionTypeConnector)` **Added**
    - `HookRefs` **Added**
    - `WorkflowTemplateRefs` **Added**
* `ConductoroneApi.Functions.UpdateFunction()`: 
  * `request.Request` **Changed**
    - `CommitMessage` **Added**
    - `Content` **Added**
    - `Function.BrowserEnabled` **Added**
    - `Function.FunctionType.Enum(functionTypeConnector)` **Added**
  * `response` **Changed**
    - `Commit` **Added**
    - `Function.BrowserEnabled` **Added**
    - `Function.FunctionType.Enum(functionTypeConnector)` **Added**
    - `Function.HookRefs` **Added**
    - `Function.WorkflowTemplateRefs` **Added**
* `ConductoroneApi.FunctionsInvocation.Get()`: `response.Invocation` **Changed**
    - `ResultRef` **Added**
    - `Status.Enum(functionInvocationStatusCancellationRequested)` **Added**
    - `Status.Enum(functionInvocationStatusCancelled)` **Added**
    - `Status.Enum(functionInvocationStatusUnknown)` **Added**
* `ConductoroneApi.FunctionsInvocation.List()`: `response.List[]` **Changed**
    - `ResultRef` **Added**
    - `Status.Enum(functionInvocationStatusCancellationRequested)` **Added**
    - `Status.Enum(functionInvocationStatusCancelled)` **Added**
    - `Status.Enum(functionInvocationStatusUnknown)` **Added**
* `ConductoroneApi.FunctionsInvocationSearch.Search()`: `response.List[]` **Changed**
    - `ResultRef` **Added**
    - `Status.Enum(functionInvocationStatusCancellationRequested)` **Added**
    - `Status.Enum(functionInvocationStatusCancelled)` **Added**
    - `Status.Enum(functionInvocationStatusUnknown)` **Added**
* `ConductoroneApi.Hooks.Create()`: 
  * `request.Request` **Changed**
    - `BuiltinPattern.BlockOutput` **Added**
    - `BuiltinPattern.BlockToolCall` **Added**
    - `BuiltinPattern.EncodedContentGuard` **Added**
    - `BuiltinPattern.LinkFilter` **Added**
    - `BuiltinPattern.PreToolBlock` **Added**
    - `BuiltinPattern.PromptInjectionScan` **Added**
    - `BuiltinPattern.SecretsMasking` **Added**
    - `Event.Enum(hookEventTypePreOutput)` **Added**
    - `JsonPatch` **Added**
    - `ManagedByGuardrails` **Added**
  * `response.Hook` **Changed**
    - `BuiltinPattern.BlockOutput` **Added**
    - `BuiltinPattern.BlockToolCall` **Added**
    - `BuiltinPattern.EncodedContentGuard` **Added**
    - `BuiltinPattern.LinkFilter` **Added**
    - `BuiltinPattern.PreToolBlock` **Added**
    - `BuiltinPattern.PromptInjectionScan` **Added**
    - `BuiltinPattern.SecretsMasking` **Added**
    - `Event.Enum(hookEventTypePreOutput)` **Added**
    - `JsonPatch` **Added**
    - `ManagedByGuardrails` **Added**
* `ConductoroneApi.Hooks.Get()`: `response.Hook` **Changed**
    - `BuiltinPattern.BlockOutput` **Added**
    - `BuiltinPattern.BlockToolCall` **Added**
    - `BuiltinPattern.EncodedContentGuard` **Added**
    - `BuiltinPattern.LinkFilter` **Added**
    - `BuiltinPattern.PreToolBlock` **Added**
    - `BuiltinPattern.PromptInjectionScan` **Added**
    - `BuiltinPattern.SecretsMasking` **Added**
    - `Event.Enum(hookEventTypePreOutput)` **Added**
    - `JsonPatch` **Added**
    - `ManagedByGuardrails` **Added**
* `ConductoroneApi.Hooks.List()`: `response.List[]` **Changed**
    - `BuiltinPattern.BlockOutput` **Added**
    - `BuiltinPattern.BlockToolCall` **Added**
    - `BuiltinPattern.EncodedContentGuard` **Added**
    - `BuiltinPattern.LinkFilter` **Added**
    - `BuiltinPattern.PreToolBlock` **Added**
    - `BuiltinPattern.PromptInjectionScan` **Added**
    - `BuiltinPattern.SecretsMasking` **Added**
    - `Event.Enum(hookEventTypePreOutput)` **Added**
    - `JsonPatch` **Added**
    - `ManagedByGuardrails` **Added**
* `ConductoroneApi.Hooks.Update()`: 
  * `request.Request.HooksServiceUpdateRequest.Hook` **Changed**
    - `BuiltinPattern.BlockOutput` **Added**
    - `BuiltinPattern.BlockToolCall` **Added**
    - `BuiltinPattern.EncodedContentGuard` **Added**
    - `BuiltinPattern.LinkFilter` **Added**
    - `BuiltinPattern.PreToolBlock` **Added**
    - `BuiltinPattern.PromptInjectionScan` **Added**
    - `BuiltinPattern.SecretsMasking` **Added**
    - `Event.Enum(hookEventTypePreOutput)` **Added**
    - `JsonPatch` **Added**
    - `ManagedByGuardrails` **Added**
  * `response.Hook` **Changed**
    - `BuiltinPattern.BlockOutput` **Added**
    - `BuiltinPattern.BlockToolCall` **Added**
    - `BuiltinPattern.EncodedContentGuard` **Added**
    - `BuiltinPattern.LinkFilter` **Added**
    - `BuiltinPattern.PreToolBlock` **Added**
    - `BuiltinPattern.PromptInjectionScan` **Added**
    - `BuiltinPattern.SecretsMasking` **Added**
    - `Event.Enum(hookEventTypePreOutput)` **Added**
    - `JsonPatch` **Added**
    - `ManagedByGuardrails` **Added**
* `ConductoroneApi.Policies.Create()`: 
  * `request.Request` **Changed**
    - `BaselinePolicyId` **Added**
    - `PolicySteps.Map<PolicySteps>.Steps[].Form.Form.Fields[].StringField.DateField` **Added**
    - `PolicySteps.Map<PolicySteps>.Steps[].Form.Form.Fields[].StringField.PickerField.C1UserPicker.ExcludeUserIds` **Added**
    - `PolicySteps.Map<PolicySteps>.Steps[].Form.Form.Fields[].StringField.PickerField.C1UserPicker.IncludeDeactivated` **Added**
    - `PolicySteps.Map<PolicySteps>.Steps[].Form.Form.Fields[].StringField.PickerField.C1UserPicker.UserIds` **Added**
    - `PolicySteps.Map<PolicySteps>.Steps[].Provision.ProvisionPolicy.DevicePlacement` **Added**
    - `PolicySteps.Map<PolicySteps>.Steps[].Provision.ProvisionPolicy.MultiStep.ProvisionSteps[].DevicePlacement` **Added**
    - `Rules[].PolicyId` **Added**
    - `Rules[].StepKey` **Added**
    - `Scope` **Added**
  * `response.Policy` **Changed**
    - `BaselinePolicyId` **Added**
    - `PolicySteps.Map<PolicySteps>.Steps[].Form.Form.Fields[].StringField.DateField` **Added**
    - `PolicySteps.Map<PolicySteps>.Steps[].Form.Form.Fields[].StringField.PickerField.C1UserPicker.ExcludeUserIds` **Added**
    - `PolicySteps.Map<PolicySteps>.Steps[].Form.Form.Fields[].StringField.PickerField.C1UserPicker.IncludeDeactivated` **Added**
    - `PolicySteps.Map<PolicySteps>.Steps[].Form.Form.Fields[].StringField.PickerField.C1UserPicker.UserIds` **Added**
    - `PolicySteps.Map<PolicySteps>.Steps[].Provision.ProvisionPolicy.DevicePlacement` **Added**
    - `PolicySteps.Map<PolicySteps>.Steps[].Provision.ProvisionPolicy.MultiStep.ProvisionSteps[].DevicePlacement` **Added**
    - `Rules[].PolicyId` **Added**
    - `Rules[].StepKey` **Added**
    - `Scope` **Added**
* `ConductoroneApi.Policies.Get()`: `response.Policy` **Changed**
    - `BaselinePolicyId` **Added**
    - `PolicySteps.Map<PolicySteps>.Steps[].Form.Form.Fields[].StringField.DateField` **Added**
    - `PolicySteps.Map<PolicySteps>.Steps[].Form.Form.Fields[].StringField.PickerField.C1UserPicker.ExcludeUserIds` **Added**
    - `PolicySteps.Map<PolicySteps>.Steps[].Form.Form.Fields[].StringField.PickerField.C1UserPicker.IncludeDeactivated` **Added**
    - `PolicySteps.Map<PolicySteps>.Steps[].Form.Form.Fields[].StringField.PickerField.C1UserPicker.UserIds` **Added**
    - `PolicySteps.Map<PolicySteps>.Steps[].Provision.ProvisionPolicy.DevicePlacement` **Added**
    - `PolicySteps.Map<PolicySteps>.Steps[].Provision.ProvisionPolicy.MultiStep.ProvisionSteps[].DevicePlacement` **Added**
    - `Rules[].PolicyId` **Added**
    - `Rules[].StepKey` **Added**
    - `Scope` **Added**
* `ConductoroneApi.Policies.List()`: `response.List[]` **Changed**
    - `BaselinePolicyId` **Added**
    - `PolicySteps.Map<PolicySteps>.Steps[].Form.Form.Fields[].StringField.DateField` **Added**
    - `PolicySteps.Map<PolicySteps>.Steps[].Form.Form.Fields[].StringField.PickerField.C1UserPicker.ExcludeUserIds` **Added**
    - `PolicySteps.Map<PolicySteps>.Steps[].Form.Form.Fields[].StringField.PickerField.C1UserPicker.IncludeDeactivated` **Added**
    - `PolicySteps.Map<PolicySteps>.Steps[].Form.Form.Fields[].StringField.PickerField.C1UserPicker.UserIds` **Added**
    - `PolicySteps.Map<PolicySteps>.Steps[].Provision.ProvisionPolicy.DevicePlacement` **Added**
    - `PolicySteps.Map<PolicySteps>.Steps[].Provision.ProvisionPolicy.MultiStep.ProvisionSteps[].DevicePlacement` **Added**
    - `Rules[].PolicyId` **Added**
    - `Rules[].StepKey` **Added**
    - `Scope` **Added**
* `ConductoroneApi.Policies.Update()`: 
  * `request.Request.UpdatePolicyRequest.Policy` **Changed**
    - `BaselinePolicyId` **Added**
    - `PolicySteps.Map<PolicySteps>.Steps[].Form.Form.Fields[].StringField.DateField` **Added**
    - `PolicySteps.Map<PolicySteps>.Steps[].Form.Form.Fields[].StringField.PickerField.C1UserPicker.ExcludeUserIds` **Added**
    - `PolicySteps.Map<PolicySteps>.Steps[].Form.Form.Fields[].StringField.PickerField.C1UserPicker.IncludeDeactivated` **Added**
    - `PolicySteps.Map<PolicySteps>.Steps[].Form.Form.Fields[].StringField.PickerField.C1UserPicker.UserIds` **Added**
    - `PolicySteps.Map<PolicySteps>.Steps[].Provision.ProvisionPolicy.DevicePlacement` **Added**
    - `PolicySteps.Map<PolicySteps>.Steps[].Provision.ProvisionPolicy.MultiStep.ProvisionSteps[].DevicePlacement` **Added**
    - `Rules[].PolicyId` **Added**
    - `Rules[].StepKey` **Added**
    - `Scope` **Added**
  * `response.Policy` **Changed**
    - `BaselinePolicyId` **Added**
    - `PolicySteps.Map<PolicySteps>.Steps[].Form.Form.Fields[].StringField.DateField` **Added**
    - `PolicySteps.Map<PolicySteps>.Steps[].Form.Form.Fields[].StringField.PickerField.C1UserPicker.ExcludeUserIds` **Added**
    - `PolicySteps.Map<PolicySteps>.Steps[].Form.Form.Fields[].StringField.PickerField.C1UserPicker.IncludeDeactivated` **Added**
    - `PolicySteps.Map<PolicySteps>.Steps[].Form.Form.Fields[].StringField.PickerField.C1UserPicker.UserIds` **Added**
    - `PolicySteps.Map<PolicySteps>.Steps[].Provision.ProvisionPolicy.DevicePlacement` **Added**
    - `PolicySteps.Map<PolicySteps>.Steps[].Provision.ProvisionPolicy.MultiStep.ProvisionSteps[].DevicePlacement` **Added**
    - `Rules[].PolicyId` **Added**
    - `Rules[].StepKey` **Added**
    - `Scope` **Added**
* `ConductoroneApi.RequestSchema.Create()`: 
  * `request.Request.Fields[].StringField` **Changed**
    - `DateField` **Added**
    - `PickerField.C1UserPicker.ExcludeUserIds` **Added**
    - `PickerField.C1UserPicker.IncludeDeactivated` **Added**
    - `PickerField.C1UserPicker.UserIds` **Added**
  * `response.RequestSchema.Form.Fields[].StringField` **Changed**
    - `DateField` **Added**
    - `PickerField.C1UserPicker.ExcludeUserIds` **Added**
    - `PickerField.C1UserPicker.IncludeDeactivated` **Added**
    - `PickerField.C1UserPicker.UserIds` **Added**
* `ConductoroneApi.RequestSchema.Get()`: `response.RequestSchema.Form.Fields[].StringField` **Changed**
    - `DateField` **Added**
    - `PickerField.C1UserPicker.ExcludeUserIds` **Added**
    - `PickerField.C1UserPicker.IncludeDeactivated` **Added**
    - `PickerField.C1UserPicker.UserIds` **Added**
* `ConductoroneApi.RequestSchema.Update()`: 
  * `request.Request.RequestSchemaServiceUpdateRequest.RequestSchema.Form.Fields[].StringField` **Changed**
    - `DateField` **Added**
    - `PickerField.C1UserPicker.ExcludeUserIds` **Added**
    - `PickerField.C1UserPicker.IncludeDeactivated` **Added**
    - `PickerField.C1UserPicker.UserIds` **Added**
  * `response.RequestSchema.Form.Fields[].StringField` **Changed**
    - `DateField` **Added**
    - `PickerField.C1UserPicker.ExcludeUserIds` **Added**
    - `PickerField.C1UserPicker.IncludeDeactivated` **Added**
    - `PickerField.C1UserPicker.UserIds` **Added**
* `ConductoroneApi.RoleMiningManagement.GetCustomAnalysisResult()`:  `response.CutoffImpactPoints` **Added**
* `ConductoroneApi.AppSearch.Search()`:  `response.List[].MatchBatonRef` **Added**
* `ConductoroneApi.AutomationSearch.SearchAutomationTemplateVersions()`:  `response.List[].AutomationSteps[].CreateRevokeTasksV2.GrantSourceFilter` **Added**
* `ConductoroneApi.AutomationSearch.SearchAutomations()`:  `response.List[].AutomationSteps[].CreateRevokeTasksV2.GrantSourceFilter` **Added**
* `ConductoroneApi.FindingAudit.Search()`: 
  * `request.Request.EventTypes[]` **Changed**
    - `Enum(findingAuditEventTypeAssigneeChanged)` **Added**
    - `Enum(findingAuditEventTypeReprocessCompleted)` **Added**
    - `Enum(findingAuditEventTypeReprocessRequested)` **Added**
  * `response.List[].EventType` **Changed**
    - `Enum(findingAuditEventTypeAssigneeChanged)` **Added**
    - `Enum(findingAuditEventTypeReprocessCompleted)` **Added**
    - `Enum(findingAuditEventTypeReprocessRequested)` **Added**
* `ConductoroneApi.FunctionsSearch.Search()`: 
  *  `request.Request.FunctionTypes[].Enum(functionTypeConnector)` **Added**
  * `response.List[]` **Changed**
    - `BrowserEnabled` **Added**
    - `FunctionType.Enum(functionTypeConnector)` **Added**
    - `HookRefs` **Added**
    - `WorkflowTemplateRefs` **Added**
* `ConductoroneApi.HooksSearch.Search()`: `response.List[]` **Changed**
    - `BuiltinPattern.BlockOutput` **Added**
    - `BuiltinPattern.BlockToolCall` **Added**
    - `BuiltinPattern.EncodedContentGuard` **Added**
    - `BuiltinPattern.LinkFilter` **Added**
    - `BuiltinPattern.PreToolBlock` **Added**
    - `BuiltinPattern.PromptInjectionScan` **Added**
    - `BuiltinPattern.SecretsMasking` **Added**
    - `Event.Enum(hookEventTypePreOutput)` **Added**
    - `JsonPatch` **Added**
    - `ManagedByGuardrails` **Added**
* `ConductoroneApi.ExternalClientSearch.Search()`:  `response.List[].ClientIdType.Enum(clientIdTypeApp)` **Added**
* `ConductoroneApi.PolicySearch.Search()`: 
  * `request.Request` **Changed**
    - `ScopeAppEntitlementId` **Added**
    - `ScopeAppId` **Added**
    - `ScopeObjectType` **Added**
    - `ScopeSlot` **Added**
    - `ScopeView` **Added**
  * `response.List[]` **Changed**
    - `BaselinePolicyId` **Added**
    - `PolicySteps.Map<PolicySteps>.Steps[].Form.Form.Fields[].StringField.DateField` **Added**
    - `PolicySteps.Map<PolicySteps>.Steps[].Form.Form.Fields[].StringField.PickerField.C1UserPicker.ExcludeUserIds` **Added**
    - `PolicySteps.Map<PolicySteps>.Steps[].Form.Form.Fields[].StringField.PickerField.C1UserPicker.IncludeDeactivated` **Added**
    - `PolicySteps.Map<PolicySteps>.Steps[].Form.Form.Fields[].StringField.PickerField.C1UserPicker.UserIds` **Added**
    - `PolicySteps.Map<PolicySteps>.Steps[].Provision.ProvisionPolicy.DevicePlacement` **Added**
    - `PolicySteps.Map<PolicySteps>.Steps[].Provision.ProvisionPolicy.MultiStep.ProvisionSteps[].DevicePlacement` **Added**
    - `Rules[].PolicyId` **Added**
    - `Rules[].StepKey` **Added**
    - `Scope` **Added**
* `ConductoroneApi.RequestCatalogSearch.SearchEntitlements()`:  `response.List[].Entitlement.AppEntitlement.DeprovisionerPolicy.DevicePlacement` **Added**
* `ConductoroneApi.TaskSearch.Search()`: 
  * `request.Request` **Changed**
    - `AccountStatuses` **Added**
    - `TaskTypes[].Action.CredentialIssue` **Added**
  * `response.List[].Task` **Changed**
    - `Actions[].Enum(taskActionTypeRetryProvisioning)` **Added**
    - `Form.Fields[].StringField.DateField` **Added**
    - `Form.Fields[].StringField.PickerField.C1UserPicker.ExcludeUserIds` **Added**
    - `Form.Fields[].StringField.PickerField.C1UserPicker.IncludeDeactivated` **Added**
    - `Form.Fields[].StringField.PickerField.C1UserPicker.UserIds` **Added**
    - `Policy.Current.Provision.Provision.ProvisionPolicy.DevicePlacement` **Added**
    - `Policy.Current.Provision.Provision.ProvisionPolicy.MultiStep.ProvisionSteps[].DevicePlacement` **Added**
    - `Policy.Current.Provision.State.Enum(provisionInstanceStateDevicePlacement)` **Added**
    - `Policy.Current.Provision.WaitingOn` **Added**
    - `Policy.Policy.BaselinePolicyId` **Added**
    - `Policy.Policy.Rules[].PolicyId` **Added**
    - `Policy.Policy.Rules[].StepKey` **Added**
    - `Policy.Policy.Scope` **Added**
    - `Type.Action.ActionInstance.ConnectorActionRef.Operation.Enum(operationIssueCredential)` **Added**
    - `Type.Action.CredentialIssue` **Added**
    - `Type.Action.Type.Enum(typeCredentialIssue)` **Added**
* `ConductoroneApi.UserSearch.Search()`: 
  *  `request.Request.SourceAppIds` **Added**
* `ConductoroneApi.AiGovernanceSettings.Get()`:  `response.AiGovernanceSettings.UntrustedJudgeDisable` **Added**
* `ConductoroneApi.AiGovernanceSettings.ListHistory()`:  `response.List[].Snapshot.UntrustedJudgeDisable` **Added**
* `ConductoroneApi.AiGovernanceSettings.Update()`: 
  *  `request.Request.AiGovernanceSettings.UntrustedJudgeDisable` **Added**
  *  `response.AiGovernanceSettings.UntrustedJudgeDisable` **Added**
* `ConductoroneApi.OrgNotificationSettings.Get()`: `response.OrgNotificationSettings.ChannelSettings` **Changed**
    - `Email.RequestCreated` **Added**
    - `Email.System` **Added**
    - `Slack.RequestCreated` **Added**
    - `Slack.System` **Added**
    - `Teams.RequestCreated` **Added**
    - `Teams.System` **Added**
* `ConductoroneApi.OrgNotificationSettings.Update()`: 
  * `request.Request.ChannelSettings` **Changed**
    - `Email.RequestCreated` **Added**
    - `Email.System` **Added**
    - `Slack.RequestCreated` **Added**
    - `Slack.System` **Added**
    - `Teams.RequestCreated` **Added**
    - `Teams.System` **Added**
  * `response.OrgNotificationSettings.ChannelSettings` **Changed**
    - `Email.RequestCreated` **Added**
    - `Email.System` **Added**
    - `Slack.RequestCreated` **Added**
    - `Slack.System` **Added**
    - `Teams.RequestCreated` **Added**
    - `Teams.System` **Added**
* `ConductoroneApi.UserNotificationSettings.Get()`: `response.UserNotificationSettings.ChannelSettings` **Changed**
    - `Email.RequestCreated` **Added**
    - `Email.System` **Added**
    - `Slack.RequestCreated` **Added**
    - `Slack.System` **Added**
    - `Teams.RequestCreated` **Added**
    - `Teams.System` **Added**
* `ConductoroneApi.UserNotificationSettings.Update()`: 
  * `request.Request.ChannelSettings` **Changed**
    - `Email.RequestCreated` **Added**
    - `Email.System` **Added**
    - `Slack.RequestCreated` **Added**
    - `Slack.System` **Added**
    - `Teams.RequestCreated` **Added**
    - `Teams.System` **Added**
  * `response.UserNotificationSettings.ChannelSettings` **Changed**
    - `Email.RequestCreated` **Added**
    - `Email.System` **Added**
    - `Slack.RequestCreated` **Added**
    - `Slack.System` **Added**
    - `Teams.RequestCreated` **Added**
    - `Teams.System` **Added**
* `ConductoroneApi.RequestSettings.Get()`:  `response.RequestSettings.MaxBulkEntitlementSelection` **Added**
* `ConductoroneApi.RequestSettings.Update()`: 
  *  `request.Request.RequestSettings.MaxBulkEntitlementSelection` **Added**
  *  `response.RequestSettings.MaxBulkEntitlementSelection` **Added**
* `ConductoroneApi.Task.CreateActionTask()`: `response.TaskView.Task` **Changed**
    - `Actions[].Enum(taskActionTypeRetryProvisioning)` **Added**
    - `Form.Fields[].StringField.DateField` **Added**
    - `Form.Fields[].StringField.PickerField.C1UserPicker.ExcludeUserIds` **Added**
    - `Form.Fields[].StringField.PickerField.C1UserPicker.IncludeDeactivated` **Added**
    - `Form.Fields[].StringField.PickerField.C1UserPicker.UserIds` **Added**
    - `Policy.Current.Provision.Provision.ProvisionPolicy.DevicePlacement` **Added**
    - `Policy.Current.Provision.Provision.ProvisionPolicy.MultiStep.ProvisionSteps[].DevicePlacement` **Added**
    - `Policy.Current.Provision.State.Enum(provisionInstanceStateDevicePlacement)` **Added**
    - `Policy.Current.Provision.WaitingOn` **Added**
    - `Policy.Policy.BaselinePolicyId` **Added**
    - `Policy.Policy.Rules[].PolicyId` **Added**
    - `Policy.Policy.Rules[].StepKey` **Added**
    - `Policy.Policy.Scope` **Added**
    - `Type.Action.ActionInstance.ConnectorActionRef.Operation.Enum(operationIssueCredential)` **Added**
    - `Type.Action.CredentialIssue` **Added**
    - `Type.Action.Type.Enum(typeCredentialIssue)` **Added**
* `ConductoroneApi.Task.CreateGrantTask()`: `response.TaskView.Task` **Changed**
    - `Actions[].Enum(taskActionTypeRetryProvisioning)` **Added**
    - `Form.Fields[].StringField.DateField` **Added**
    - `Form.Fields[].StringField.PickerField.C1UserPicker.ExcludeUserIds` **Added**
    - `Form.Fields[].StringField.PickerField.C1UserPicker.IncludeDeactivated` **Added**
    - `Form.Fields[].StringField.PickerField.C1UserPicker.UserIds` **Added**
    - `Policy.Current.Provision.Provision.ProvisionPolicy.DevicePlacement` **Added**
    - `Policy.Current.Provision.Provision.ProvisionPolicy.MultiStep.ProvisionSteps[].DevicePlacement` **Added**
    - `Policy.Current.Provision.State.Enum(provisionInstanceStateDevicePlacement)` **Added**
    - `Policy.Current.Provision.WaitingOn` **Added**
    - `Policy.Policy.BaselinePolicyId` **Added**
    - `Policy.Policy.Rules[].PolicyId` **Added**
    - `Policy.Policy.Rules[].StepKey` **Added**
    - `Policy.Policy.Scope` **Added**
    - `Type.Action.ActionInstance.ConnectorActionRef.Operation.Enum(operationIssueCredential)` **Added**
    - `Type.Action.CredentialIssue` **Added**
    - `Type.Action.Type.Enum(typeCredentialIssue)` **Added**
* `ConductoroneApi.Task.CreateOffboardingTask()`: `response.TaskView.Task` **Changed**
    - `Actions[].Enum(taskActionTypeRetryProvisioning)` **Added**
    - `Form.Fields[].StringField.DateField` **Added**
    - `Form.Fields[].StringField.PickerField.C1UserPicker.ExcludeUserIds` **Added**
    - `Form.Fields[].StringField.PickerField.C1UserPicker.IncludeDeactivated` **Added**
    - `Form.Fields[].StringField.PickerField.C1UserPicker.UserIds` **Added**
    - `Policy.Current.Provision.Provision.ProvisionPolicy.DevicePlacement` **Added**
    - `Policy.Current.Provision.Provision.ProvisionPolicy.MultiStep.ProvisionSteps[].DevicePlacement` **Added**
    - `Policy.Current.Provision.State.Enum(provisionInstanceStateDevicePlacement)` **Added**
    - `Policy.Current.Provision.WaitingOn` **Added**
    - `Policy.Policy.BaselinePolicyId` **Added**
    - `Policy.Policy.Rules[].PolicyId` **Added**
    - `Policy.Policy.Rules[].StepKey` **Added**
    - `Policy.Policy.Scope` **Added**
    - `Type.Action.ActionInstance.ConnectorActionRef.Operation.Enum(operationIssueCredential)` **Added**
    - `Type.Action.CredentialIssue` **Added**
    - `Type.Action.Type.Enum(typeCredentialIssue)` **Added**
* `ConductoroneApi.Task.CreateResourceActionTask()`: `response.TaskView.Task` **Changed**
    - `Actions[].Enum(taskActionTypeRetryProvisioning)` **Added**
    - `Form.Fields[].StringField.DateField` **Added**
    - `Form.Fields[].StringField.PickerField.C1UserPicker.ExcludeUserIds` **Added**
    - `Form.Fields[].StringField.PickerField.C1UserPicker.IncludeDeactivated` **Added**
    - `Form.Fields[].StringField.PickerField.C1UserPicker.UserIds` **Added**
    - `Policy.Current.Provision.Provision.ProvisionPolicy.DevicePlacement` **Added**
    - `Policy.Current.Provision.Provision.ProvisionPolicy.MultiStep.ProvisionSteps[].DevicePlacement` **Added**
    - `Policy.Current.Provision.State.Enum(provisionInstanceStateDevicePlacement)` **Added**
    - `Policy.Current.Provision.WaitingOn` **Added**
    - `Policy.Policy.BaselinePolicyId` **Added**
    - `Policy.Policy.Rules[].PolicyId` **Added**
    - `Policy.Policy.Rules[].StepKey` **Added**
    - `Policy.Policy.Scope` **Added**
    - `Type.Action.ActionInstance.ConnectorActionRef.Operation.Enum(operationIssueCredential)` **Added**
    - `Type.Action.CredentialIssue` **Added**
    - `Type.Action.Type.Enum(typeCredentialIssue)` **Added**
* `ConductoroneApi.Task.CreateRevokeTask()`: `response.TaskView.Task` **Changed**
    - `Actions[].Enum(taskActionTypeRetryProvisioning)` **Added**
    - `Form.Fields[].StringField.DateField` **Added**
    - `Form.Fields[].StringField.PickerField.C1UserPicker.ExcludeUserIds` **Added**
    - `Form.Fields[].StringField.PickerField.C1UserPicker.IncludeDeactivated` **Added**
    - `Form.Fields[].StringField.PickerField.C1UserPicker.UserIds` **Added**
    - `Policy.Current.Provision.Provision.ProvisionPolicy.DevicePlacement` **Added**
    - `Policy.Current.Provision.Provision.ProvisionPolicy.MultiStep.ProvisionSteps[].DevicePlacement` **Added**
    - `Policy.Current.Provision.State.Enum(provisionInstanceStateDevicePlacement)` **Added**
    - `Policy.Current.Provision.WaitingOn` **Added**
    - `Policy.Policy.BaselinePolicyId` **Added**
    - `Policy.Policy.Rules[].PolicyId` **Added**
    - `Policy.Policy.Rules[].StepKey` **Added**
    - `Policy.Policy.Scope` **Added**
    - `Type.Action.ActionInstance.ConnectorActionRef.Operation.Enum(operationIssueCredential)` **Added**
    - `Type.Action.CredentialIssue` **Added**
    - `Type.Action.Type.Enum(typeCredentialIssue)` **Added**
* `ConductoroneApi.Task.Get()`: `response.TaskView.Task` **Changed**
    - `Actions[].Enum(taskActionTypeRetryProvisioning)` **Added**
    - `Form.Fields[].StringField.DateField` **Added**
    - `Form.Fields[].StringField.PickerField.C1UserPicker.ExcludeUserIds` **Added**
    - `Form.Fields[].StringField.PickerField.C1UserPicker.IncludeDeactivated` **Added**
    - `Form.Fields[].StringField.PickerField.C1UserPicker.UserIds` **Added**
    - `Policy.Current.Provision.Provision.ProvisionPolicy.DevicePlacement` **Added**
    - `Policy.Current.Provision.Provision.ProvisionPolicy.MultiStep.ProvisionSteps[].DevicePlacement` **Added**
    - `Policy.Current.Provision.State.Enum(provisionInstanceStateDevicePlacement)` **Added**
    - `Policy.Current.Provision.WaitingOn` **Added**
    - `Policy.Policy.BaselinePolicyId` **Added**
    - `Policy.Policy.Rules[].PolicyId` **Added**
    - `Policy.Policy.Rules[].StepKey` **Added**
    - `Policy.Policy.Scope` **Added**
    - `Type.Action.ActionInstance.ConnectorActionRef.Operation.Enum(operationIssueCredential)` **Added**
    - `Type.Action.CredentialIssue` **Added**
    - `Type.Action.Type.Enum(typeCredentialIssue)` **Added**
* `ConductoroneApi.TaskAudit.List()`: 
  *  `request.Request.ExcludeComments` **Added**
  * `response` **Changed**
    - `List[].AccountDeleted` **Added**
    - `List[].ActionSubmitted.Action.ActionType.Enum(taskActionTypeRetryProvisioning)` **Added**
    - `List[].AutomationTriggered` **Added**
    - `List[].ConditionalPolicyExecutionResult.ChainDepth` **Added**
    - `List[].ConditionalPolicyExecutionResult.OutcomePolicyId` **Added**
    - `List[].ConditionalPolicyExecutionResult.PolicyId` **Added**
    - `List[].ProvisionEntitlementMergeCompleted` **Added**
    - `List[].ProvisionEntitlementMergeTimedOut` **Added**
    - `List[].ProvisionWaitingForEntitlementMerge` **Added**
    - `List[].WebhookSuccess.Comment` **Added**
    - `TotalCount` **Added**
* `ConductoroneApi.TaskActions.Approve()`: `response.TaskView.Task` **Changed**
    - `Actions[].Enum(taskActionTypeRetryProvisioning)` **Added**
    - `Form.Fields[].StringField.DateField` **Added**
    - `Form.Fields[].StringField.PickerField.C1UserPicker.ExcludeUserIds` **Added**
    - `Form.Fields[].StringField.PickerField.C1UserPicker.IncludeDeactivated` **Added**
    - `Form.Fields[].StringField.PickerField.C1UserPicker.UserIds` **Added**
    - `Policy.Current.Provision.Provision.ProvisionPolicy.DevicePlacement` **Added**
    - `Policy.Current.Provision.Provision.ProvisionPolicy.MultiStep.ProvisionSteps[].DevicePlacement` **Added**
    - `Policy.Current.Provision.State.Enum(provisionInstanceStateDevicePlacement)` **Added**
    - `Policy.Current.Provision.WaitingOn` **Added**
    - `Policy.Policy.BaselinePolicyId` **Added**
    - `Policy.Policy.Rules[].PolicyId` **Added**
    - `Policy.Policy.Rules[].StepKey` **Added**
    - `Policy.Policy.Scope` **Added**
    - `Type.Action.ActionInstance.ConnectorActionRef.Operation.Enum(operationIssueCredential)` **Added**
    - `Type.Action.CredentialIssue` **Added**
    - `Type.Action.Type.Enum(typeCredentialIssue)` **Added**
* `ConductoroneApi.TaskActions.ApproveWithStepUp()`: `response.TaskView.Task` **Changed**
    - `Actions[].Enum(taskActionTypeRetryProvisioning)` **Added**
    - `Form.Fields[].StringField.DateField` **Added**
    - `Form.Fields[].StringField.PickerField.C1UserPicker.ExcludeUserIds` **Added**
    - `Form.Fields[].StringField.PickerField.C1UserPicker.IncludeDeactivated` **Added**
    - `Form.Fields[].StringField.PickerField.C1UserPicker.UserIds` **Added**
    - `Policy.Current.Provision.Provision.ProvisionPolicy.DevicePlacement` **Added**
    - `Policy.Current.Provision.Provision.ProvisionPolicy.MultiStep.ProvisionSteps[].DevicePlacement` **Added**
    - `Policy.Current.Provision.State.Enum(provisionInstanceStateDevicePlacement)` **Added**
    - `Policy.Current.Provision.WaitingOn` **Added**
    - `Policy.Policy.BaselinePolicyId` **Added**
    - `Policy.Policy.Rules[].PolicyId` **Added**
    - `Policy.Policy.Rules[].StepKey` **Added**
    - `Policy.Policy.Scope` **Added**
    - `Type.Action.ActionInstance.ConnectorActionRef.Operation.Enum(operationIssueCredential)` **Added**
    - `Type.Action.CredentialIssue` **Added**
    - `Type.Action.Type.Enum(typeCredentialIssue)` **Added**
* `ConductoroneApi.TaskActions.Close()`: `response.TaskView.Task` **Changed**
    - `Actions[].Enum(taskActionTypeRetryProvisioning)` **Added**
    - `Form.Fields[].StringField.DateField` **Added**
    - `Form.Fields[].StringField.PickerField.C1UserPicker.ExcludeUserIds` **Added**
    - `Form.Fields[].StringField.PickerField.C1UserPicker.IncludeDeactivated` **Added**
    - `Form.Fields[].StringField.PickerField.C1UserPicker.UserIds` **Added**
    - `Policy.Current.Provision.Provision.ProvisionPolicy.DevicePlacement` **Added**
    - `Policy.Current.Provision.Provision.ProvisionPolicy.MultiStep.ProvisionSteps[].DevicePlacement` **Added**
    - `Policy.Current.Provision.State.Enum(provisionInstanceStateDevicePlacement)` **Added**
    - `Policy.Current.Provision.WaitingOn` **Added**
    - `Policy.Policy.BaselinePolicyId` **Added**
    - `Policy.Policy.Rules[].PolicyId` **Added**
    - `Policy.Policy.Rules[].StepKey` **Added**
    - `Policy.Policy.Scope` **Added**
    - `Type.Action.ActionInstance.ConnectorActionRef.Operation.Enum(operationIssueCredential)` **Added**
    - `Type.Action.CredentialIssue` **Added**
    - `Type.Action.Type.Enum(typeCredentialIssue)` **Added**
* `ConductoroneApi.TaskActions.Comment()`: `response.TaskView.Task` **Changed**
    - `Actions[].Enum(taskActionTypeRetryProvisioning)` **Added**
    - `Form.Fields[].StringField.DateField` **Added**
    - `Form.Fields[].StringField.PickerField.C1UserPicker.ExcludeUserIds` **Added**
    - `Form.Fields[].StringField.PickerField.C1UserPicker.IncludeDeactivated` **Added**
    - `Form.Fields[].StringField.PickerField.C1UserPicker.UserIds` **Added**
    - `Policy.Current.Provision.Provision.ProvisionPolicy.DevicePlacement` **Added**
    - `Policy.Current.Provision.Provision.ProvisionPolicy.MultiStep.ProvisionSteps[].DevicePlacement` **Added**
    - `Policy.Current.Provision.State.Enum(provisionInstanceStateDevicePlacement)` **Added**
    - `Policy.Current.Provision.WaitingOn` **Added**
    - `Policy.Policy.BaselinePolicyId` **Added**
    - `Policy.Policy.Rules[].PolicyId` **Added**
    - `Policy.Policy.Rules[].StepKey` **Added**
    - `Policy.Policy.Scope` **Added**
    - `Type.Action.ActionInstance.ConnectorActionRef.Operation.Enum(operationIssueCredential)` **Added**
    - `Type.Action.CredentialIssue` **Added**
    - `Type.Action.Type.Enum(typeCredentialIssue)` **Added**
* `ConductoroneApi.TaskActions.Deny()`: `response.TaskView.Task` **Changed**
    - `Actions[].Enum(taskActionTypeRetryProvisioning)` **Added**
    - `Form.Fields[].StringField.DateField` **Added**
    - `Form.Fields[].StringField.PickerField.C1UserPicker.ExcludeUserIds` **Added**
    - `Form.Fields[].StringField.PickerField.C1UserPicker.IncludeDeactivated` **Added**
    - `Form.Fields[].StringField.PickerField.C1UserPicker.UserIds` **Added**
    - `Policy.Current.Provision.Provision.ProvisionPolicy.DevicePlacement` **Added**
    - `Policy.Current.Provision.Provision.ProvisionPolicy.MultiStep.ProvisionSteps[].DevicePlacement` **Added**
    - `Policy.Current.Provision.State.Enum(provisionInstanceStateDevicePlacement)` **Added**
    - `Policy.Current.Provision.WaitingOn` **Added**
    - `Policy.Policy.BaselinePolicyId` **Added**
    - `Policy.Policy.Rules[].PolicyId` **Added**
    - `Policy.Policy.Rules[].StepKey` **Added**
    - `Policy.Policy.Scope` **Added**
    - `Type.Action.ActionInstance.ConnectorActionRef.Operation.Enum(operationIssueCredential)` **Added**
    - `Type.Action.CredentialIssue` **Added**
    - `Type.Action.Type.Enum(typeCredentialIssue)` **Added**
* `ConductoroneApi.TaskActions.EscalateToEmergencyAccess()`: `response.TaskView.Task` **Changed**
    - `Actions[].Enum(taskActionTypeRetryProvisioning)` **Added**
    - `Form.Fields[].StringField.DateField` **Added**
    - `Form.Fields[].StringField.PickerField.C1UserPicker.ExcludeUserIds` **Added**
    - `Form.Fields[].StringField.PickerField.C1UserPicker.IncludeDeactivated` **Added**
    - `Form.Fields[].StringField.PickerField.C1UserPicker.UserIds` **Added**
    - `Policy.Current.Provision.Provision.ProvisionPolicy.DevicePlacement` **Added**
    - `Policy.Current.Provision.Provision.ProvisionPolicy.MultiStep.ProvisionSteps[].DevicePlacement` **Added**
    - `Policy.Current.Provision.State.Enum(provisionInstanceStateDevicePlacement)` **Added**
    - `Policy.Current.Provision.WaitingOn` **Added**
    - `Policy.Policy.BaselinePolicyId` **Added**
    - `Policy.Policy.Rules[].PolicyId` **Added**
    - `Policy.Policy.Rules[].StepKey` **Added**
    - `Policy.Policy.Scope` **Added**
    - `Type.Action.ActionInstance.ConnectorActionRef.Operation.Enum(operationIssueCredential)` **Added**
    - `Type.Action.CredentialIssue` **Added**
    - `Type.Action.Type.Enum(typeCredentialIssue)` **Added**
* `ConductoroneApi.TaskActions.HardReset()`: `response.TaskView.Task` **Changed**
    - `Actions[].Enum(taskActionTypeRetryProvisioning)` **Added**
    - `Form.Fields[].StringField.DateField` **Added**
    - `Form.Fields[].StringField.PickerField.C1UserPicker.ExcludeUserIds` **Added**
    - `Form.Fields[].StringField.PickerField.C1UserPicker.IncludeDeactivated` **Added**
    - `Form.Fields[].StringField.PickerField.C1UserPicker.UserIds` **Added**
    - `Policy.Current.Provision.Provision.ProvisionPolicy.DevicePlacement` **Added**
    - `Policy.Current.Provision.Provision.ProvisionPolicy.MultiStep.ProvisionSteps[].DevicePlacement` **Added**
    - `Policy.Current.Provision.State.Enum(provisionInstanceStateDevicePlacement)` **Added**
    - `Policy.Current.Provision.WaitingOn` **Added**
    - `Policy.Policy.BaselinePolicyId` **Added**
    - `Policy.Policy.Rules[].PolicyId` **Added**
    - `Policy.Policy.Rules[].StepKey` **Added**
    - `Policy.Policy.Scope` **Added**
    - `Type.Action.ActionInstance.ConnectorActionRef.Operation.Enum(operationIssueCredential)` **Added**
    - `Type.Action.CredentialIssue` **Added**
    - `Type.Action.Type.Enum(typeCredentialIssue)` **Added**
* `ConductoroneApi.TaskActions.ProcessNow()`: `response.TaskView.Task` **Changed**
    - `Actions[].Enum(taskActionTypeRetryProvisioning)` **Added**
    - `Form.Fields[].StringField.DateField` **Added**
    - `Form.Fields[].StringField.PickerField.C1UserPicker.ExcludeUserIds` **Added**
    - `Form.Fields[].StringField.PickerField.C1UserPicker.IncludeDeactivated` **Added**
    - `Form.Fields[].StringField.PickerField.C1UserPicker.UserIds` **Added**
    - `Policy.Current.Provision.Provision.ProvisionPolicy.DevicePlacement` **Added**
    - `Policy.Current.Provision.Provision.ProvisionPolicy.MultiStep.ProvisionSteps[].DevicePlacement` **Added**
    - `Policy.Current.Provision.State.Enum(provisionInstanceStateDevicePlacement)` **Added**
    - `Policy.Current.Provision.WaitingOn` **Added**
    - `Policy.Policy.BaselinePolicyId` **Added**
    - `Policy.Policy.Rules[].PolicyId` **Added**
    - `Policy.Policy.Rules[].StepKey` **Added**
    - `Policy.Policy.Scope` **Added**
    - `Type.Action.ActionInstance.ConnectorActionRef.Operation.Enum(operationIssueCredential)` **Added**
    - `Type.Action.CredentialIssue` **Added**
    - `Type.Action.Type.Enum(typeCredentialIssue)` **Added**
* `ConductoroneApi.TaskActions.Reassign()`: `response.TaskView.Task` **Changed**
    - `Actions[].Enum(taskActionTypeRetryProvisioning)` **Added**
    - `Form.Fields[].StringField.DateField` **Added**
    - `Form.Fields[].StringField.PickerField.C1UserPicker.ExcludeUserIds` **Added**
    - `Form.Fields[].StringField.PickerField.C1UserPicker.IncludeDeactivated` **Added**
    - `Form.Fields[].StringField.PickerField.C1UserPicker.UserIds` **Added**
    - `Policy.Current.Provision.Provision.ProvisionPolicy.DevicePlacement` **Added**
    - `Policy.Current.Provision.Provision.ProvisionPolicy.MultiStep.ProvisionSteps[].DevicePlacement` **Added**
    - `Policy.Current.Provision.State.Enum(provisionInstanceStateDevicePlacement)` **Added**
    - `Policy.Current.Provision.WaitingOn` **Added**
    - `Policy.Policy.BaselinePolicyId` **Added**
    - `Policy.Policy.Rules[].PolicyId` **Added**
    - `Policy.Policy.Rules[].StepKey` **Added**
    - `Policy.Policy.Scope` **Added**
    - `Type.Action.ActionInstance.ConnectorActionRef.Operation.Enum(operationIssueCredential)` **Added**
    - `Type.Action.CredentialIssue` **Added**
    - `Type.Action.Type.Enum(typeCredentialIssue)` **Added**
* `ConductoroneApi.TaskActions.Restart()`: `response.TaskView.Task` **Changed**
    - `Actions[].Enum(taskActionTypeRetryProvisioning)` **Added**
    - `Form.Fields[].StringField.DateField` **Added**
    - `Form.Fields[].StringField.PickerField.C1UserPicker.ExcludeUserIds` **Added**
    - `Form.Fields[].StringField.PickerField.C1UserPicker.IncludeDeactivated` **Added**
    - `Form.Fields[].StringField.PickerField.C1UserPicker.UserIds` **Added**
    - `Policy.Current.Provision.Provision.ProvisionPolicy.DevicePlacement` **Added**
    - `Policy.Current.Provision.Provision.ProvisionPolicy.MultiStep.ProvisionSteps[].DevicePlacement` **Added**
    - `Policy.Current.Provision.State.Enum(provisionInstanceStateDevicePlacement)` **Added**
    - `Policy.Current.Provision.WaitingOn` **Added**
    - `Policy.Policy.BaselinePolicyId` **Added**
    - `Policy.Policy.Rules[].PolicyId` **Added**
    - `Policy.Policy.Rules[].StepKey` **Added**
    - `Policy.Policy.Scope` **Added**
    - `Type.Action.ActionInstance.ConnectorActionRef.Operation.Enum(operationIssueCredential)` **Added**
    - `Type.Action.CredentialIssue` **Added**
    - `Type.Action.Type.Enum(typeCredentialIssue)` **Added**
* `ConductoroneApi.TaskActions.SkipStep()`: `response.TaskView.Task` **Changed**
    - `Actions[].Enum(taskActionTypeRetryProvisioning)` **Added**
    - `Form.Fields[].StringField.DateField` **Added**
    - `Form.Fields[].StringField.PickerField.C1UserPicker.ExcludeUserIds` **Added**
    - `Form.Fields[].StringField.PickerField.C1UserPicker.IncludeDeactivated` **Added**
    - `Form.Fields[].StringField.PickerField.C1UserPicker.UserIds` **Added**
    - `Policy.Current.Provision.Provision.ProvisionPolicy.DevicePlacement` **Added**
    - `Policy.Current.Provision.Provision.ProvisionPolicy.MultiStep.ProvisionSteps[].DevicePlacement` **Added**
    - `Policy.Current.Provision.State.Enum(provisionInstanceStateDevicePlacement)` **Added**
    - `Policy.Current.Provision.WaitingOn` **Added**
    - `Policy.Policy.BaselinePolicyId` **Added**
    - `Policy.Policy.Rules[].PolicyId` **Added**
    - `Policy.Policy.Rules[].StepKey` **Added**
    - `Policy.Policy.Scope` **Added**
    - `Type.Action.ActionInstance.ConnectorActionRef.Operation.Enum(operationIssueCredential)` **Added**
    - `Type.Action.CredentialIssue` **Added**
    - `Type.Action.Type.Enum(typeCredentialIssue)` **Added**
* `ConductoroneApi.TaskActions.UpdateGrantDuration()`: `response.TaskView.Task` **Changed**
    - `Actions[].Enum(taskActionTypeRetryProvisioning)` **Added**
    - `Form.Fields[].StringField.DateField` **Added**
    - `Form.Fields[].StringField.PickerField.C1UserPicker.ExcludeUserIds` **Added**
    - `Form.Fields[].StringField.PickerField.C1UserPicker.IncludeDeactivated` **Added**
    - `Form.Fields[].StringField.PickerField.C1UserPicker.UserIds` **Added**
    - `Policy.Current.Provision.Provision.ProvisionPolicy.DevicePlacement` **Added**
    - `Policy.Current.Provision.Provision.ProvisionPolicy.MultiStep.ProvisionSteps[].DevicePlacement` **Added**
    - `Policy.Current.Provision.State.Enum(provisionInstanceStateDevicePlacement)` **Added**
    - `Policy.Current.Provision.WaitingOn` **Added**
    - `Policy.Policy.BaselinePolicyId` **Added**
    - `Policy.Policy.Rules[].PolicyId` **Added**
    - `Policy.Policy.Rules[].StepKey` **Added**
    - `Policy.Policy.Scope` **Added**
    - `Type.Action.ActionInstance.ConnectorActionRef.Operation.Enum(operationIssueCredential)` **Added**
    - `Type.Action.CredentialIssue` **Added**
    - `Type.Action.Type.Enum(typeCredentialIssue)` **Added**
* `ConductoroneApi.TaskActions.UpdateRequestData()`: `response.TaskView.Task` **Changed**
    - `Actions[].Enum(taskActionTypeRetryProvisioning)` **Added**
    - `Form.Fields[].StringField.DateField` **Added**
    - `Form.Fields[].StringField.PickerField.C1UserPicker.ExcludeUserIds` **Added**
    - `Form.Fields[].StringField.PickerField.C1UserPicker.IncludeDeactivated` **Added**
    - `Form.Fields[].StringField.PickerField.C1UserPicker.UserIds` **Added**
    - `Policy.Current.Provision.Provision.ProvisionPolicy.DevicePlacement` **Added**
    - `Policy.Current.Provision.Provision.ProvisionPolicy.MultiStep.ProvisionSteps[].DevicePlacement` **Added**
    - `Policy.Current.Provision.State.Enum(provisionInstanceStateDevicePlacement)` **Added**
    - `Policy.Current.Provision.WaitingOn` **Added**
    - `Policy.Policy.BaselinePolicyId` **Added**
    - `Policy.Policy.Rules[].PolicyId` **Added**
    - `Policy.Policy.Rules[].StepKey` **Added**
    - `Policy.Policy.Scope` **Added**
    - `Type.Action.ActionInstance.ConnectorActionRef.Operation.Enum(operationIssueCredential)` **Added**
    - `Type.Action.CredentialIssue` **Added**
    - `Type.Action.Type.Enum(typeCredentialIssue)` **Added**
* `ConductoroneApi.ConnectorOwnersV2.CreateEntitlementOwner()`:  `response.ConnectorOwnerEntitlement.AppEntitlement.DeprovisionerPolicy.DevicePlacement` **Added**
* `ConductoroneApi.ConnectorOwnersV2.GetEntitlementOwner()`:  `response.ConnectorOwnerEntitlement.AppEntitlement.DeprovisionerPolicy.DevicePlacement` **Added**
* `ConductoroneApi.ConnectorOwnersV2.SearchEntitlementOwners()`:  `response.List[].AppEntitlement.DeprovisionerPolicy.DevicePlacement` **Added**
* `ConductoroneApi.AppEntitlementOwnersV2.CreateEntitlementOwner()`:  `response.AppEntitlementOwnerEntitlement.AppEntitlement.DeprovisionerPolicy.DevicePlacement` **Added**
* `ConductoroneApi.AppEntitlementOwnersV2.GetEntitlementOwner()`:  `response.AppEntitlementOwnerEntitlement.AppEntitlement.DeprovisionerPolicy.DevicePlacement` **Added**
* `ConductoroneApi.AppEntitlementOwnersV2.SearchEntitlementOwners()`:  `response.List[].AppEntitlement.DeprovisionerPolicy.DevicePlacement` **Added**
* `ConductoroneApi.AppOwnersV2.CreateEntitlementOwner()`:  `response.AppOwnerEntitlement.AppEntitlement.DeprovisionerPolicy.DevicePlacement` **Added**
* `ConductoroneApi.AppOwnersV2.GetEntitlementOwner()`:  `response.AppOwnerEntitlement.AppEntitlement.DeprovisionerPolicy.DevicePlacement` **Added**
* `ConductoroneApi.AppOwnersV2.SearchEntitlementOwners()`:  `response.List[].AppEntitlement.DeprovisionerPolicy.DevicePlacement` **Added**
* `ConductoroneApi.AppResourceOwnersV2.CreateEntitlementOwner()`:  `response.AppResourceOwnerEntitlement.AppEntitlement.DeprovisionerPolicy.DevicePlacement` **Added**
* `ConductoroneApi.AppResourceOwnersV2.GetEntitlementOwner()`:  `response.AppResourceOwnerEntitlement.AppEntitlement.DeprovisionerPolicy.DevicePlacement` **Added**
* `ConductoroneApi.AppResourceOwnersV2.SearchEntitlementOwners()`:  `response.List[].AppEntitlement.DeprovisionerPolicy.DevicePlacement` **Added**
* `ConductoroneApi.AppUserOwnersV2.CreateEntitlementOwner()`:  `response.AppUserOwnerEntitlement.AppEntitlement.DeprovisionerPolicy.DevicePlacement` **Added**
* `ConductoroneApi.AppUserOwnersV2.SearchEntitlementOwners()`:  `response.List[].AppEntitlement.DeprovisionerPolicy.DevicePlacement` **Added**
* `ConductoroneApi.UserOwnersV2.CreateEntitlementOwner()`:  `response.UserOwnerEntitlement.AppEntitlement.DeprovisionerPolicy.DevicePlacement` **Added**
* `ConductoroneApi.UserOwnersV2.SearchEntitlementOwners()`:  `response.List[].AppEntitlement.DeprovisionerPolicy.DevicePlacement` **Added**
