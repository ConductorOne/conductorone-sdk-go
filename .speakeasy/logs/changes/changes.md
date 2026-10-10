## Go SDK Changes:
* `ConductoroneApi.SpendInsights.GetDenial()`: `response.Episode` **Changed** (Breaking ⚠️)
    - `Reason.Enum(denyReasonAppPausedByYou)` **Removed** (Breaking ⚠️)
    - `Reason.Enum(denyReasonAppSuspended)` **Removed** (Breaking ⚠️)
    - `Reason.Enum(denyReasonNoSupply)` **Removed** (Breaking ⚠️)
    - `Reason.Enum(denyReasonSuspendedByAdmin)` **Removed** (Breaking ⚠️)
    - `Reason.Enum(denyReasonTenantFrozen)` **Removed** (Breaking ⚠️)
    - `Reason.Enum(denyReasonUnspecified)` **Removed** (Breaking ⚠️)
    - `Reason.Enum(spendDenyReasonAppPausedByYou)` **Added**
    - `Reason.Enum(spendDenyReasonAppSuspended)` **Added**
    - `Reason.Enum(spendDenyReasonNoSupply)` **Added**
    - `Reason.Enum(spendDenyReasonSuspendedByAdmin)` **Added**
    - `Reason.Enum(spendDenyReasonTenantFrozen)` **Added**
    - `Reason.Enum(spendDenyReasonUnspecified)` **Added**
    - `RequestTaskState.Enum(taskStateClosed)` **Added**
    - `RequestTaskState.Enum(taskStateOpen)` **Added**
    - `RequestTaskState.Enum(taskStateUnspecified)` **Added**
    - `RequestTaskState.Enum(ticketStateClosed)` **Removed** (Breaking ⚠️)
    - `RequestTaskState.Enum(ticketStateOpen)` **Removed** (Breaking ⚠️)
    - `RequestTaskState.Enum(ticketStateUnspecified)` **Removed** (Breaking ⚠️)
    - `ScopeKind.Enum(spendBlockScopeKindApp)` **Removed** (Breaking ⚠️)
    - `ScopeKind.Enum(spendBlockScopeKindSubject)` **Removed** (Breaking ⚠️)
    - `ScopeKind.Enum(spendBlockScopeKindSubjectApp)` **Removed** (Breaking ⚠️)
    - `ScopeKind.Enum(spendBlockScopeKindTenant)` **Removed** (Breaking ⚠️)
    - `ScopeKind.Enum(spendBlockScopeKindUnspecified)` **Removed** (Breaking ⚠️)
    - `ScopeKind.Enum(spendScopeKindApp)` **Added**
    - `ScopeKind.Enum(spendScopeKindSubject)` **Added**
    - `ScopeKind.Enum(spendScopeKindSubjectApp)` **Added**
    - `ScopeKind.Enum(spendScopeKindTenant)` **Added**
    - `ScopeKind.Enum(spendScopeKindUnspecified)` **Added**
* `ConductoroneApi.SpendInsights.GetMySpendStatus()`: `response.Block` **Changed** (Breaking ⚠️)
    - `AttemptedAppId` **Added**
    - `PeriodEnd` **Added**
    - `PersonalRequestSizing` **Added**
    - `Reason.Enum(denyReasonAppPausedByYou)` **Removed** (Breaking ⚠️)
    - `Reason.Enum(denyReasonAppSuspended)` **Removed** (Breaking ⚠️)
    - `Reason.Enum(denyReasonNoSupply)` **Removed** (Breaking ⚠️)
    - `Reason.Enum(denyReasonSuspendedByAdmin)` **Removed** (Breaking ⚠️)
    - `Reason.Enum(denyReasonTenantFrozen)` **Removed** (Breaking ⚠️)
    - `Reason.Enum(denyReasonUnspecified)` **Removed** (Breaking ⚠️)
    - `Reason.Enum(spendDenyReasonAppPausedByYou)` **Added**
    - `Reason.Enum(spendDenyReasonAppSuspended)` **Added**
    - `Reason.Enum(spendDenyReasonNoSupply)` **Added**
    - `Reason.Enum(spendDenyReasonSuspendedByAdmin)` **Added**
    - `Reason.Enum(spendDenyReasonTenantFrozen)` **Added**
    - `Reason.Enum(spendDenyReasonUnspecified)` **Added**
    - `RequestTaskId` **Added**
    - `ScopeAppId` **Removed** (Breaking ⚠️)
    - `ScopeKind` **Removed** (Breaking ⚠️)
    - `Scope` **Added**
* `ConductoroneApi.SpendInsights.ResolveEffectiveLimits()`: `response.DenialReason` **Changed** (Breaking ⚠️)
    - `Enum(denyReasonAppPausedByYou)` **Removed** (Breaking ⚠️)
    - `Enum(denyReasonAppSuspended)` **Removed** (Breaking ⚠️)
    - `Enum(denyReasonNoSupply)` **Removed** (Breaking ⚠️)
    - `Enum(denyReasonSuspendedByAdmin)` **Removed** (Breaking ⚠️)
    - `Enum(denyReasonTenantFrozen)` **Removed** (Breaking ⚠️)
    - `Enum(denyReasonUnspecified)` **Removed** (Breaking ⚠️)
    - `Enum(spendDenyReasonAppPausedByYou)` **Added**
    - `Enum(spendDenyReasonAppSuspended)` **Added**
    - `Enum(spendDenyReasonNoSupply)` **Added**
    - `Enum(spendDenyReasonSuspendedByAdmin)` **Added**
    - `Enum(spendDenyReasonTenantFrozen)` **Added**
    - `Enum(spendDenyReasonUnspecified)` **Added**
* `ConductoroneApi.SpendInsights.SearchDenials()`: 
  * `request.Request.Filters` **Changed** (Breaking ⚠️)
    - `Departments` **Added**
    - `ScopeKind.Enum(spendBlockScopeKindApp)` **Removed** (Breaking ⚠️)
    - `ScopeKind.Enum(spendBlockScopeKindSubject)` **Removed** (Breaking ⚠️)
    - `ScopeKind.Enum(spendBlockScopeKindSubjectApp)` **Removed** (Breaking ⚠️)
    - `ScopeKind.Enum(spendBlockScopeKindTenant)` **Removed** (Breaking ⚠️)
    - `ScopeKind.Enum(spendBlockScopeKindUnspecified)` **Removed** (Breaking ⚠️)
    - `ScopeKind.Enum(spendScopeKindApp)` **Added**
    - `ScopeKind.Enum(spendScopeKindSubject)` **Added**
    - `ScopeKind.Enum(spendScopeKindSubjectApp)` **Added**
    - `ScopeKind.Enum(spendScopeKindTenant)` **Added**
    - `ScopeKind.Enum(spendScopeKindUnspecified)` **Added**
  * `response.Episodes[]` **Changed** (Breaking ⚠️)
    - `Reason.Enum(denyReasonAppPausedByYou)` **Removed** (Breaking ⚠️)
    - `Reason.Enum(denyReasonAppSuspended)` **Removed** (Breaking ⚠️)
    - `Reason.Enum(denyReasonNoSupply)` **Removed** (Breaking ⚠️)
    - `Reason.Enum(denyReasonSuspendedByAdmin)` **Removed** (Breaking ⚠️)
    - `Reason.Enum(denyReasonTenantFrozen)` **Removed** (Breaking ⚠️)
    - `Reason.Enum(denyReasonUnspecified)` **Removed** (Breaking ⚠️)
    - `Reason.Enum(spendDenyReasonAppPausedByYou)` **Added**
    - `Reason.Enum(spendDenyReasonAppSuspended)` **Added**
    - `Reason.Enum(spendDenyReasonNoSupply)` **Added**
    - `Reason.Enum(spendDenyReasonSuspendedByAdmin)` **Added**
    - `Reason.Enum(spendDenyReasonTenantFrozen)` **Added**
    - `Reason.Enum(spendDenyReasonUnspecified)` **Added**
    - `RequestTaskState.Enum(taskStateClosed)` **Added**
    - `RequestTaskState.Enum(taskStateOpen)` **Added**
    - `RequestTaskState.Enum(taskStateUnspecified)` **Added**
    - `RequestTaskState.Enum(ticketStateClosed)` **Removed** (Breaking ⚠️)
    - `RequestTaskState.Enum(ticketStateOpen)` **Removed** (Breaking ⚠️)
    - `RequestTaskState.Enum(ticketStateUnspecified)` **Removed** (Breaking ⚠️)
    - `ScopeKind.Enum(spendBlockScopeKindApp)` **Removed** (Breaking ⚠️)
    - `ScopeKind.Enum(spendBlockScopeKindSubject)` **Removed** (Breaking ⚠️)
    - `ScopeKind.Enum(spendBlockScopeKindSubjectApp)` **Removed** (Breaking ⚠️)
    - `ScopeKind.Enum(spendBlockScopeKindTenant)` **Removed** (Breaking ⚠️)
    - `ScopeKind.Enum(spendBlockScopeKindUnspecified)` **Removed** (Breaking ⚠️)
    - `ScopeKind.Enum(spendScopeKindApp)` **Added**
    - `ScopeKind.Enum(spendScopeKindSubject)` **Added**
    - `ScopeKind.Enum(spendScopeKindSubjectApp)` **Added**
    - `ScopeKind.Enum(spendScopeKindTenant)` **Added**
    - `ScopeKind.Enum(spendScopeKindUnspecified)` **Added**
* `ConductoroneApi.AccessReviewSetupEntitlement.AddCampaignEntitlements()`: **Added**
* `ConductoroneApi.AccessReviewSetupEntitlement.RemoveCampaignEntitlements()`: **Added**
* `ConductoroneApi.AccessReviewTemplateSetupEntitlement.AddTemplateEntitlements()`: **Added**
* `ConductoroneApi.AccessReviewTemplateSetupEntitlement.RemoveTemplateEntitlements()`: **Added**
* `ConductoroneApi.AppSecretAdmin.Revoke()`: **Added**
* `ConductoroneApi.AuthzenServer.Create()`: **Added**
* `ConductoroneApi.AuthzenServer.CreateAuthzenPolicy()`: **Added**
* `ConductoroneApi.AuthzenServer.Delete()`: **Added**
* `ConductoroneApi.AuthzenServer.DeleteAuthzenPolicy()`: **Added**
* `ConductoroneApi.AuthzenServer.Get()`: **Added**
* `ConductoroneApi.AuthzenServer.GetAuthzenPolicy()`: **Added**
* `ConductoroneApi.AuthzenServer.List()`: **Added**
* `ConductoroneApi.AuthzenServer.ListAuthzenPolicies()`: **Added**
* `ConductoroneApi.AuthzenServer.TestAuthzenPolicy()`: **Added**
* `ConductoroneApi.AuthzenServer.Update()`: **Added**
* `ConductoroneApi.Edge.CheckInferenceVaultAccess()`: **Added**
* `ConductoroneApi.Edge.Create()`: **Added**
* `ConductoroneApi.Edge.Delete()`: **Added**
* `ConductoroneApi.Edge.Get()`: **Added**
* `ConductoroneApi.Edge.GetAuditInclusionProof()`: **Added**
* `ConductoroneApi.Edge.GetEgressTrafficEvent()`: **Added**
* `ConductoroneApi.Edge.GetEgressTrafficSummary()`: **Added**
* `ConductoroneApi.Edge.GetEgressTrafficSummaryRollup()`: **Added**
* `ConductoroneApi.Edge.GetInferenceUsageAttribution()`: **Added**
* `ConductoroneApi.Edge.GetInferenceUsageAttributionRollup()`: **Added**
* `ConductoroneApi.Edge.GetInferenceUsageSummary()`: **Added**
* `ConductoroneApi.Edge.GetInferenceUsageSummaryRollup()`: **Added**
* `ConductoroneApi.Edge.List()`: **Added**
* `ConductoroneApi.Edge.ListAuditAttestationKeys()`: **Added**
* `ConductoroneApi.Edge.ListAuditAttestationManifests()`: **Added**
* `ConductoroneApi.Edge.ListEdgesRollup()`: **Added**
* `ConductoroneApi.Edge.ListEgressDenialSummary()`: **Added**
* `ConductoroneApi.Edge.ListEgressRuleEntitlementIssues()`: **Added**
* `ConductoroneApi.Edge.ListEgressTopAgents()`: **Added**
* `ConductoroneApi.Edge.ListEgressTopAgentsRollup()`: **Added**
* `ConductoroneApi.Edge.ListEgressTopDestinations()`: **Added**
* `ConductoroneApi.Edge.ListEgressTopToolsRollup()`: **Added**
* `ConductoroneApi.Edge.ListEgressTrafficEvents()`: **Added**
* `ConductoroneApi.Edge.ListInferenceCatalog()`: **Added**
* `ConductoroneApi.Edge.ListInferenceRoutes()`: **Added**
* `ConductoroneApi.Edge.ListInferenceTopAgents()`: **Added**
* `ConductoroneApi.Edge.ListInferenceTopTools()`: **Added**
* `ConductoroneApi.Edge.ListInferenceTrafficEvents()`: **Added**
* `ConductoroneApi.Edge.ListUnattributedInferenceTrafficEvents()`: **Added**
* `ConductoroneApi.Edge.ListUsableEdges()`: **Added**
* `ConductoroneApi.Edge.PreviewEgressHostDecision()`: **Added**
* `ConductoroneApi.Edge.PreviewEgressRulesUpdate()`: **Added**
* `ConductoroneApi.Edge.Update()`: **Added**
* `ConductoroneApi.Artifact.Delete()`: **Added**
* `ConductoroneApi.Artifact.Get()`: **Added**
* `ConductoroneApi.Artifact.GetVersion()`: **Added**
* `ConductoroneApi.Artifact.ListVersions()`: **Added**
* `ConductoroneApi.Artifact.Search()`: **Added**
* `ConductoroneApi.Artifact.Update()`: **Added**
* `ConductoroneApi.RequestCatalogManagement.SearchScopeRoleBindingsPerCatalog()`: **Added**
* `ConductoroneApi.Classifier.CreateClassifier()`: **Added**
* `ConductoroneApi.Classifier.CreateClassifierBinding()`: **Added**
* `ConductoroneApi.Classifier.DeleteClassifier()`: **Added**
* `ConductoroneApi.Classifier.DeleteClassifierBinding()`: **Added**
* `ConductoroneApi.Classifier.GetClassifier()`: **Added**
* `ConductoroneApi.Classifier.ListClassifierBindings()`: **Added**
* `ConductoroneApi.Classifier.ListClassifiers()`: **Added**
* `ConductoroneApi.Classifier.UpdateClassifier()`: **Added**
* `ConductoroneApi.ClassifierTemplate.AddClassifierRuleFromTemplate()`: **Added**
* `ConductoroneApi.ClassifierTemplate.GetClassifierTemplate()`: **Added**
* `ConductoroneApi.ClassifierTemplate.InstantiateClassifierTemplate()`: **Added**
* `ConductoroneApi.ClassifierTemplate.ListClassifierTemplates()`: **Added**
* `ConductoroneApi.FindingSettings.GetEdgeFindingSettings()`: **Added**
* `ConductoroneApi.FindingSettings.UpdateEdgeFindingSettings()`: **Added**
* `ConductoroneApi.MyFundLimits.ClearTemporaryLimit()`: **Added**
* `ConductoroneApi.MyFundLimits.SetTemporaryLimit()`: **Added**
* `ConductoroneApi.GoLink.Create()`: **Added**
* `ConductoroneApi.GoLink.Delete()`: **Added**
* `ConductoroneApi.GoLink.Get()`: **Added**
* `ConductoroneApi.GoLink.List()`: **Added**
* `ConductoroneApi.GoLink.ListVersions()`: **Added**
* `ConductoroneApi.GoLink.Resolve()`: **Added**
* `ConductoroneApi.GoLink.Update()`: **Added**
* `ConductoroneApi.GoLinkSearch.Search()`: **Added**
* `ConductoroneApi.ModelApiKey.Create()`: **Added**
* `ConductoroneApi.ModelApiKey.Get()`: **Added**
* `ConductoroneApi.ModelApiKey.List()`: **Added**
* `ConductoroneApi.ModelApiKey.Revoke()`: **Added**
* `ConductoroneApi.RoleMiningManagement.CountEntitlementSelectionFilterAttributeValues()`: **Added**
* `ConductoroneApi.RoleMiningManagement.ListCohortFilterAttributeValues()`: **Added**
* `ConductoroneApi.RoleMiningManagement.ListCohortFilterAttributes()`: **Added**
* `ConductoroneApi.AppSecretSearch.SearchMyVendedCredentials()`: **Added**
* `ConductoroneApi.VirtualMcpServerMySearch.Search()`: **Added**
* `ConductoroneApi.RequestCatalogSearch.SearchRequestableCredentialOfferings()`: **Added**
* `ConductoroneApi.PaperSecretAdmin.Delete()`: **Added**
* `ConductoroneApi.PaperSecret.Delete()`: **Added**
* `ConductoroneApi.ShadowMcpOccurrence.Search()`: **Added**
* `ConductoroneApi.ToolGatesSearch.Search()`: **Added**
* `ConductoroneApi.AgentClassifier.GetAgentClassifierPolicy()`: **Added**
* `ConductoroneApi.AgentClassifier.UpdateAgentClassifierPolicy()`: **Added**
* `ConductoroneApi.SpendInsights.GetMySpendForecast()`: **Added**
* `ConductoroneApi.SpendInsights.GetMySpendHistory()`: **Added**
* `ConductoroneApi.SpendInsights.SearchMySpendUsage()`: **Added**
* `ConductoroneApi.Task.CreateSpendRemedyTask()`: **Added**
* `ConductoroneApi.Task.GetSpendRemedyReview()`: **Added**
* `ConductoroneApi.ToolGates.Create()`: **Added**
* `ConductoroneApi.ToolGates.Delete()`: **Added**
* `ConductoroneApi.ToolGates.Get()`: **Added**
* `ConductoroneApi.ToolGates.List()`: **Added**
* `ConductoroneApi.ToolGates.Update()`: **Added**
* `ConductoroneApi.UserAttributeManagement.Create()`: **Added**
* `ConductoroneApi.UserAttributeManagement.Delete()`: **Added**
* `ConductoroneApi.UserAttributeManagement.Get()`: **Added**
* `ConductoroneApi.UserAttributeManagement.List()`: **Added**
* `ConductoroneApi.UserAttributeManagement.ListAttributeTypes()`: **Added**
* `ConductoroneApi.UserAttributeManagement.Update()`: **Added**
* `ConductoroneApi.VirtualMcpServer.Create()`: **Added**
* `ConductoroneApi.VirtualMcpServer.Delete()`: **Added**
* `ConductoroneApi.VirtualMcpServer.Get()`: **Added**
* `ConductoroneApi.VirtualMcpServer.List()`: **Added**
* `ConductoroneApi.VirtualMcpServer.Update()`: **Added**
* `ConductoroneApi.VirtualMcpToolBinding.CreateBindings()`: **Added**
* `ConductoroneApi.VirtualMcpToolBinding.DeleteBindings()`: **Added**
* `ConductoroneApi.VirtualMcpToolBinding.List()`: **Added**
* `ConductoroneApi.VirtualMcpToolsetBinding.CreateBindings()`: **Added**
* `ConductoroneApi.VirtualMcpToolsetBinding.DeleteBindings()`: **Added**
* `ConductoroneApi.VirtualMcpToolsetBinding.List()`: **Added**
* `ConductoroneApi.PaperSecretAdmin.Revoke()`: **Removed** (Breaking ⚠️)
* `ConductoroneApi.PaperSecret.Revoke()`: **Removed** (Breaking ⚠️)
* `ConductoroneApi.TbControlPlane.GetDiscoverySnapshot()`: **Removed** (Breaking ⚠️)
* `ConductoroneApi.TbControlPlane.GetEgressPolicy()`: **Removed** (Breaking ⚠️)
* `ConductoroneApi.TbControlPlane.PushDiscovery()`: **Removed** (Breaking ⚠️)
* `ConductoroneApi.TbControlPlane.SaveEgressPolicy()`: **Removed** (Breaking ⚠️)
* `ConductoroneApi.A2Ui.ListSurfaces()`: `response.Surfaces[]` **Changed**
    - `Components[].C1AppUserPicker` **Added**
    - `Components[].C1CredentialOfferingPicker` **Added**
    - `SavedReportId` **Added**
* `ConductoroneApi.AccessReview.Create()`: 
  * `request.Request` **Changed**
    - `ScopeType.Enum(accessReviewScopeTypeByAccessProfiles)` **Added**
    - `ScopeV2.AccessProfileReviewDimensions` **Added**
    - `ScopeV2.SpecificAccessProfiles` **Added**
  * `response.AccessReview.AccessReview` **Changed**
    - `CampaignSchedule` **Added**
    - `ColumnConfig.Columns[].Enum(accessReviewTaskColumnResourceDescription)` **Added**
    - `ColumnConfig.OrderedColumns[].Builtin.Enum(accessReviewTaskColumnResourceDescription)` **Added**
    - `ScopeType.Enum(accessReviewScopeTypeByAccessProfiles)` **Added**
    - `ScopeV2.AccessProfileReviewDimensions` **Added**
    - `ScopeV2.SpecificAccessProfiles` **Added**
* `ConductoroneApi.AccessReview.Get()`: `response.AccessReview.AccessReview` **Changed**
    - `CampaignSchedule` **Added**
    - `ColumnConfig.Columns[].Enum(accessReviewTaskColumnResourceDescription)` **Added**
    - `ColumnConfig.OrderedColumns[].Builtin.Enum(accessReviewTaskColumnResourceDescription)` **Added**
    - `ScopeType.Enum(accessReviewScopeTypeByAccessProfiles)` **Added**
    - `ScopeV2.AccessProfileReviewDimensions` **Added**
    - `ScopeV2.SpecificAccessProfiles` **Added**
* `ConductoroneApi.AccessReview.List()`: `response.List[].AccessReview` **Changed**
    - `CampaignSchedule` **Added**
    - `ColumnConfig.Columns[].Enum(accessReviewTaskColumnResourceDescription)` **Added**
    - `ColumnConfig.OrderedColumns[].Builtin.Enum(accessReviewTaskColumnResourceDescription)` **Added**
    - `ScopeType.Enum(accessReviewScopeTypeByAccessProfiles)` **Added**
    - `ScopeV2.AccessProfileReviewDimensions` **Added**
    - `ScopeV2.SpecificAccessProfiles` **Added**
* `ConductoroneApi.AccessReview.Update()`: 
  * `request.Request.AccessReviewServiceUpdateRequest.AccessReview` **Changed**
    - `CampaignSchedule` **Added**
    - `ColumnConfig.Columns[].Enum(accessReviewTaskColumnResourceDescription)` **Added**
    - `ColumnConfig.OrderedColumns[].Builtin.Enum(accessReviewTaskColumnResourceDescription)` **Added**
    - `ScopeType.Enum(accessReviewScopeTypeByAccessProfiles)` **Added**
    - `ScopeV2.AccessProfileReviewDimensions` **Added**
    - `ScopeV2.SpecificAccessProfiles` **Added**
  * `response.AccessReview.AccessReview` **Changed**
    - `CampaignSchedule` **Added**
    - `ColumnConfig.Columns[].Enum(accessReviewTaskColumnResourceDescription)` **Added**
    - `ColumnConfig.OrderedColumns[].Builtin.Enum(accessReviewTaskColumnResourceDescription)` **Added**
    - `ScopeType.Enum(accessReviewScopeTypeByAccessProfiles)` **Added**
    - `ScopeV2.AccessProfileReviewDimensions` **Added**
    - `ScopeV2.SpecificAccessProfiles` **Added**
* `ConductoroneApi.AccessReviewSetupEntitlement.GetCampaignScopeAndEntitlements()`: `response.ScopeV2` **Changed**
    - `AccessProfileReviewDimensions` **Added**
    - `SpecificAccessProfiles` **Added**
* `ConductoroneApi.AccessReviewSetupEntitlement.SetCampaignScopeAndEntitlements()`: 
  * `request.Request.AccessReviewSetupEntitlementAndScopeServiceSetRequest.ScopeV2` **Changed**
    - `AccessProfileReviewDimensions` **Added**
    - `SpecificAccessProfiles` **Added**
  * `response.ScopeV2` **Changed**
    - `AccessProfileReviewDimensions` **Added**
    - `SpecificAccessProfiles` **Added**
* `ConductoroneApi.AccessReviewSetupEntitlement.SetCampaignScopeByResourceType()`: 
  * `request.Request.AccessReviewSetScopeByResourceTypeRequest.ScopeV2` **Changed**
    - `AccessProfileReviewDimensions` **Added**
    - `SpecificAccessProfiles` **Added**
* `ConductoroneApi.AccessReviewTemplate.Create()`: 
  * `request.Request` **Changed**
    - `ColumnConfig.Columns[].Enum(accessReviewTaskColumnResourceDescription)` **Added**
    - `ColumnConfig.OrderedColumns[].Builtin.Enum(accessReviewTaskColumnResourceDescription)` **Added**
    - `Scope.AccessProfileReviewDimensions` **Added**
    - `Scope.SpecificAccessProfiles` **Added**
    - `ScopeType.Enum(accessReviewScopeTypeByAccessProfiles)` **Added**
  * `response.AccessReviewTemplate` **Changed**
    - `CampaignSchedule` **Added**
    - `ColumnConfig.Columns[].Enum(accessReviewTaskColumnResourceDescription)` **Added**
    - `ColumnConfig.OrderedColumns[].Builtin.Enum(accessReviewTaskColumnResourceDescription)` **Added**
    - `Scope.AccessProfileReviewDimensions` **Added**
    - `Scope.SpecificAccessProfiles` **Added**
    - `ScopeType.Enum(accessReviewScopeTypeByAccessProfiles)` **Added**
* `ConductoroneApi.AccessReviewTemplate.Get()`: `response.AccessReviewTemplate` **Changed**
    - `CampaignSchedule` **Added**
    - `ColumnConfig.Columns[].Enum(accessReviewTaskColumnResourceDescription)` **Added**
    - `ColumnConfig.OrderedColumns[].Builtin.Enum(accessReviewTaskColumnResourceDescription)` **Added**
    - `Scope.AccessProfileReviewDimensions` **Added**
    - `Scope.SpecificAccessProfiles` **Added**
    - `ScopeType.Enum(accessReviewScopeTypeByAccessProfiles)` **Added**
* `ConductoroneApi.AccessReviewTemplate.Update()`: 
  * `request.Request.AccessReviewTemplateServiceUpdateRequest.AccessReviewTemplate` **Changed**
    - `CampaignSchedule` **Added**
    - `ColumnConfig.Columns[].Enum(accessReviewTaskColumnResourceDescription)` **Added**
    - `ColumnConfig.OrderedColumns[].Builtin.Enum(accessReviewTaskColumnResourceDescription)` **Added**
    - `Scope.AccessProfileReviewDimensions` **Added**
    - `Scope.SpecificAccessProfiles` **Added**
    - `ScopeType.Enum(accessReviewScopeTypeByAccessProfiles)` **Added**
  * `response.AccessReviewTemplate` **Changed**
    - `CampaignSchedule` **Added**
    - `ColumnConfig.Columns[].Enum(accessReviewTaskColumnResourceDescription)` **Added**
    - `ColumnConfig.OrderedColumns[].Builtin.Enum(accessReviewTaskColumnResourceDescription)` **Added**
    - `Scope.AccessProfileReviewDimensions` **Added**
    - `Scope.SpecificAccessProfiles` **Added**
    - `ScopeType.Enum(accessReviewScopeTypeByAccessProfiles)` **Added**
* `ConductoroneApi.AccessReviewTemplateSetupEntitlement.GetScopeAndEntitlements()`: `response.Scope` **Changed**
    - `AccessProfileReviewDimensions` **Added**
    - `SpecificAccessProfiles` **Added**
* `ConductoroneApi.AccessReviewTemplateSetupEntitlement.SetScopeAndEntitlements()`: 
  * `request.Request.AccessReviewTemplateSetupEntitlementServiceSetRequest.Scope` **Changed**
    - `AccessProfileReviewDimensions` **Added**
    - `SpecificAccessProfiles` **Added**
  * `response.Scope` **Changed**
    - `AccessProfileReviewDimensions` **Added**
    - `SpecificAccessProfiles` **Added**
* `ConductoroneApi.AccessReviewTemplateSetupEntitlement.SetScopeByResourceType()`: 
  * `request.Request.AccessReviewTemplateSetScopeByResourceTypeRequest.Scope` **Changed**
    - `AccessProfileReviewDimensions` **Added**
    - `SpecificAccessProfiles` **Added**
* `ConductoroneApi.Connector.Update()`: 
  *  `request.Request.ConnectorServiceUpdateRequest.Justification` **Added**
* `ConductoroneApi.Connector.UpdateConnectorSchedule()`: 
  *  `request.Request.UpdateConnectorScheduleRequest.Justification` **Added**
* `ConductoroneApi.AppEntitlementRoutingRule.ListAppEntitlementRoutingRules()`:  `response.MaxRulesPerApp` **Added**
* `ConductoroneApi.AppEntitlementSearch.Search()`: `request.Request` **Changed**
    - `AppResourceTypeRefs` **Added**
    - `OwnerUserIds` **Added**
* `ConductoroneApi.McpServer.GetCatalog()`:  `response.CatalogEntry.ConnectType` **Added**
* `ConductoroneApi.McpServer.ListCatalog()`:  `response.List[].ConnectType` **Added**
* `ConductoroneApi.McpServer.Register()`: 
  *  `request.Request.McpServerServiceRegisterRequest.AccessProvisioning` **Added**
* `ConductoroneApi.AppResource.CreateManuallyManagedAppResource()`:  `response.AppResource.IsManuallyManaged` **Added**
* `ConductoroneApi.AppResource.Get()`:  `response.AppResourceView.AppResource.IsManuallyManaged` **Added**
* `ConductoroneApi.AppResource.List()`:  `response.List[].AppResource.IsManuallyManaged` **Added**
* `ConductoroneApi.AppResource.Update()`:  `response.AppResourceView.AppResource.IsManuallyManaged` **Added**
* `ConductoroneApi.Automation.CreateAutomation()`: 
  * `request.Request.AutomationSteps[].CreateRevokeTasksV2` **Changed**
    - `ExclusionCriteria.ExcludedResourceTypeRefs` **Added**
    - `InclusionCriteria.ResourceTypeRefs` **Added**
  * `response.Automation.AutomationSteps[].CreateRevokeTasksV2` **Changed**
    - `ExclusionCriteria.ExcludedResourceTypeRefs` **Added**
    - `InclusionCriteria.ResourceTypeRefs` **Added**
* `ConductoroneApi.Automation.GetAutomation()`: `response.Automation.AutomationSteps[].CreateRevokeTasksV2` **Changed**
    - `ExclusionCriteria.ExcludedResourceTypeRefs` **Added**
    - `InclusionCriteria.ResourceTypeRefs` **Added**
* `ConductoroneApi.Automation.ListAutomations()`: `response.List[].AutomationSteps[].CreateRevokeTasksV2` **Changed**
    - `ExclusionCriteria.ExcludedResourceTypeRefs` **Added**
    - `InclusionCriteria.ResourceTypeRefs` **Added**
* `ConductoroneApi.Automation.UpdateAutomation()`: 
  * `request.Request.UpdateAutomationRequest.Automation.AutomationSteps[].CreateRevokeTasksV2` **Changed**
    - `ExclusionCriteria.ExcludedResourceTypeRefs` **Added**
    - `InclusionCriteria.ResourceTypeRefs` **Added**
  * `response.Automation.AutomationSteps[].CreateRevokeTasksV2` **Changed**
    - `ExclusionCriteria.ExcludedResourceTypeRefs` **Added**
    - `InclusionCriteria.ResourceTypeRefs` **Added**
* `ConductoroneApi.RequestCatalogManagement.PlanTypeChange()`: `response.Impacts[].Code` **Changed**
    - `Enum(requestCatalogTypeChangeImpactCodeContainingCatalogs)` **Added**
    - `Enum(requestCatalogTypeChangeImpactCodeRoleMembershipEntitlements)` **Added**
* `ConductoroneApi.ConnectorCatalog.ConfigurationSchema()`:  `response.FormSchema.Fields[].StringField.PasswordField.Multiline` **Added**
* `ConductoroneApi.Finding.BulkCreateFindingTasks()`: 
  * `request.Request.SearchRequest` **Changed**
    - `EdgeIds` **Added**
    - `FindingTypes[].Enum(findingTypeEdge)` **Added**
    - `FindingTypes[].Enum(findingTypeMcpGatewayToolCallRisk)` **Added**
    - `FindingTypes[].Enum(findingTypeShadowApp)` **Added**
* `ConductoroneApi.Finding.BulkUpdateFindingState()`: 
  * `request.Request.SearchRequest` **Changed**
    - `EdgeIds` **Added**
    - `FindingTypes[].Enum(findingTypeEdge)` **Added**
    - `FindingTypes[].Enum(findingTypeMcpGatewayToolCallRisk)` **Added**
    - `FindingTypes[].Enum(findingTypeShadowApp)` **Added**
* `ConductoroneApi.Finding.CreateFinding()`: `response.Finding` **Changed**
    - `CredentialExpiring.AppResourceSecret` **Added**
    - `EdgeFindingEvidence` **Added**
    - `EdgeFinding` **Added**
    - `McpGatewayToolCallRiskEvidence` **Added**
    - `McpGatewayToolCallRisk` **Added**
    - `ShadowApp` **Added**
    - `ShadowMcp.Transport` **Added**
    - `UsageResourceTarget` **Added**
* `ConductoroneApi.Finding.CreateFindingTask()`: `response.Finding` **Changed**
    - `CredentialExpiring.AppResourceSecret` **Added**
    - `EdgeFindingEvidence` **Added**
    - `EdgeFinding` **Added**
    - `McpGatewayToolCallRiskEvidence` **Added**
    - `McpGatewayToolCallRisk` **Added**
    - `ShadowApp` **Added**
    - `ShadowMcp.Transport` **Added**
    - `UsageResourceTarget` **Added**
* `ConductoroneApi.Finding.GetFinding()`: `response.Finding` **Changed**
    - `CredentialExpiring.AppResourceSecret` **Added**
    - `EdgeFindingEvidence` **Added**
    - `EdgeFinding` **Added**
    - `McpGatewayToolCallRiskEvidence` **Added**
    - `McpGatewayToolCallRisk` **Added**
    - `ShadowApp` **Added**
    - `ShadowMcp.Transport` **Added**
    - `UsageResourceTarget` **Added**
* `ConductoroneApi.Finding.UpdateFindingAssignee()`: `response.Finding` **Changed**
    - `CredentialExpiring.AppResourceSecret` **Added**
    - `EdgeFindingEvidence` **Added**
    - `EdgeFinding` **Added**
    - `McpGatewayToolCallRiskEvidence` **Added**
    - `McpGatewayToolCallRisk` **Added**
    - `ShadowApp` **Added**
    - `ShadowMcp.Transport` **Added**
    - `UsageResourceTarget` **Added**
* `ConductoroneApi.Finding.UpdateFindingState()`: `response.Finding` **Changed**
    - `CredentialExpiring.AppResourceSecret` **Added**
    - `EdgeFindingEvidence` **Added**
    - `EdgeFinding` **Added**
    - `McpGatewayToolCallRiskEvidence` **Added**
    - `McpGatewayToolCallRisk` **Added**
    - `ShadowApp` **Added**
    - `ShadowMcp.Transport` **Added**
    - `UsageResourceTarget` **Added**
* `ConductoroneApi.FindingRoutingRule.CreateFindingRoutingRule()`: 
  * `request.Request.RoutingRule.FindingType` **Changed**
    - `Enum(findingTypeEdge)` **Added**
    - `Enum(findingTypeMcpGatewayToolCallRisk)` **Added**
    - `Enum(findingTypeShadowApp)` **Added**
  * `response.RoutingRule.FindingType` **Changed**
    - `Enum(findingTypeEdge)` **Added**
    - `Enum(findingTypeMcpGatewayToolCallRisk)` **Added**
    - `Enum(findingTypeShadowApp)` **Added**
* `ConductoroneApi.FindingRoutingRule.GetFindingRoutingRule()`: `response.RoutingRule.FindingType` **Changed**
    - `Enum(findingTypeEdge)` **Added**
    - `Enum(findingTypeMcpGatewayToolCallRisk)` **Added**
    - `Enum(findingTypeShadowApp)` **Added**
* `ConductoroneApi.FindingRoutingRule.ListFindingRoutingRules()`: `response.List[].FindingType` **Changed**
    - `Enum(findingTypeEdge)` **Added**
    - `Enum(findingTypeMcpGatewayToolCallRisk)` **Added**
    - `Enum(findingTypeShadowApp)` **Added**
* `ConductoroneApi.FindingRoutingRule.UpdateFindingRoutingRule()`: 
  * `request.Request.UpdateFindingRoutingRuleRequest.RoutingRule.FindingType` **Changed**
    - `Enum(findingTypeEdge)` **Added**
    - `Enum(findingTypeMcpGatewayToolCallRisk)` **Added**
    - `Enum(findingTypeShadowApp)` **Added**
  * `response.RoutingRule.FindingType` **Changed**
    - `Enum(findingTypeEdge)` **Added**
    - `Enum(findingTypeMcpGatewayToolCallRisk)` **Added**
    - `Enum(findingTypeShadowApp)` **Added**
* `ConductoroneApi.FindingSearch.Search()`: 
  * `request.Request` **Changed**
    - `EdgeIds` **Added**
    - `FindingTypes[].Enum(findingTypeEdge)` **Added**
    - `FindingTypes[].Enum(findingTypeMcpGatewayToolCallRisk)` **Added**
    - `FindingTypes[].Enum(findingTypeShadowApp)` **Added**
  * `response.List[]` **Changed**
    - `CredentialExpiring.AppResourceSecret` **Added**
    - `EdgeFindingEvidence` **Added**
    - `EdgeFinding` **Added**
    - `McpGatewayToolCallRiskEvidence` **Added**
    - `McpGatewayToolCallRisk` **Added**
    - `ShadowApp` **Added**
    - `ShadowMcp.Transport` **Added**
    - `UsageResourceTarget` **Added**
* `ConductoroneApi.FindingSettings.ListFindingSettings()`: `response.List[]` **Changed**
    - `FindingType.Enum(findingTypeEdge)` **Added**
    - `FindingType.Enum(findingTypeMcpGatewayToolCallRisk)` **Added**
    - `FindingType.Enum(findingTypeShadowApp)` **Added**
    - `HasRemediationPlaybook` **Added**
* `ConductoroneApi.FindingSettings.UpdateFindingSettings()`: 
  * `request.Request.Settings[].FindingType` **Changed**
    - `Enum(findingTypeEdge)` **Added**
    - `Enum(findingTypeMcpGatewayToolCallRisk)` **Added**
    - `Enum(findingTypeShadowApp)` **Added**
  * `response.List[]` **Changed**
    - `FindingType.Enum(findingTypeEdge)` **Added**
    - `FindingType.Enum(findingTypeMcpGatewayToolCallRisk)` **Added**
    - `FindingType.Enum(findingTypeShadowApp)` **Added**
    - `HasRemediationPlaybook` **Added**
* `ConductoroneApi.FindingTransformationRule.CreateFindingTransformationRule()`: 
  * `request.Request.TransformationRule` **Changed**
    - `FindingType.Enum(findingTypeEdge)` **Added**
    - `FindingType.Enum(findingTypeMcpGatewayToolCallRisk)` **Added**
    - `FindingType.Enum(findingTypeShadowApp)` **Added**
    - `Transforms[].SetAssignee` **Added**
  * `response.TransformationRule` **Changed**
    - `FindingType.Enum(findingTypeEdge)` **Added**
    - `FindingType.Enum(findingTypeMcpGatewayToolCallRisk)` **Added**
    - `FindingType.Enum(findingTypeShadowApp)` **Added**
    - `Transforms[].SetAssignee` **Added**
* `ConductoroneApi.FindingTransformationRule.GetFindingTransformationRule()`: `response.TransformationRule` **Changed**
    - `FindingType.Enum(findingTypeEdge)` **Added**
    - `FindingType.Enum(findingTypeMcpGatewayToolCallRisk)` **Added**
    - `FindingType.Enum(findingTypeShadowApp)` **Added**
    - `Transforms[].SetAssignee` **Added**
* `ConductoroneApi.FindingTransformationRule.ListFindingTransformationRules()`: `response.List[]` **Changed**
    - `FindingType.Enum(findingTypeEdge)` **Added**
    - `FindingType.Enum(findingTypeMcpGatewayToolCallRisk)` **Added**
    - `FindingType.Enum(findingTypeShadowApp)` **Added**
    - `Transforms[].SetAssignee` **Added**
* `ConductoroneApi.FindingTransformationRule.UpdateFindingTransformationRule()`: 
  * `request.Request.UpdateFindingTransformationRuleRequest.TransformationRule` **Changed**
    - `FindingType.Enum(findingTypeEdge)` **Added**
    - `FindingType.Enum(findingTypeMcpGatewayToolCallRisk)` **Added**
    - `FindingType.Enum(findingTypeShadowApp)` **Added**
    - `Transforms[].SetAssignee` **Added**
  * `response.TransformationRule` **Changed**
    - `FindingType.Enum(findingTypeEdge)` **Added**
    - `FindingType.Enum(findingTypeMcpGatewayToolCallRisk)` **Added**
    - `FindingType.Enum(findingTypeShadowApp)` **Added**
    - `Transforms[].SetAssignee` **Added**
* `ConductoroneApi.FundPolicy.Create()`: 
  *  `request.Request.SpendRemedyRequestPolicyId` **Added**
  *  `response.Policy.SpendRemedyRequestPolicyId` **Added**
* `ConductoroneApi.FundPolicy.FreezeTenant()`:  `response.Policy.SpendRemedyRequestPolicyId` **Added**
* `ConductoroneApi.FundPolicy.Get()`:  `response.Policy.SpendRemedyRequestPolicyId` **Added**
* `ConductoroneApi.FundPolicy.ListHistory()`:  `response.List[].Snapshot.SpendRemedyRequestPolicyId` **Added**
* `ConductoroneApi.FundPolicy.SetOrgCeiling()`:  `response.Policy.SpendRemedyRequestPolicyId` **Added**
* `ConductoroneApi.FundPolicy.UnfreezeTenant()`:  `response.Policy.SpendRemedyRequestPolicyId` **Added**
* `ConductoroneApi.FundPolicy.Update()`: 
  *  `request.Request.Policy.SpendRemedyRequestPolicyId` **Added**
  *  `response.Policy.SpendRemedyRequestPolicyId` **Added**
* `ConductoroneApi.Policies.Create()`: 
  *  `request.Request.PolicySteps.Map<PolicySteps>.Steps[].Form.Form.Fields[].StringField.PasswordField.Multiline` **Added**
  *  `response.Policy.PolicySteps.Map<PolicySteps>.Steps[].Form.Form.Fields[].StringField.PasswordField.Multiline` **Added**
* `ConductoroneApi.Policies.Get()`:  `response.Policy.PolicySteps.Map<PolicySteps>.Steps[].Form.Form.Fields[].StringField.PasswordField.Multiline` **Added**
* `ConductoroneApi.Policies.List()`:  `response.List[].PolicySteps.Map<PolicySteps>.Steps[].Form.Form.Fields[].StringField.PasswordField.Multiline` **Added**
* `ConductoroneApi.Policies.Update()`: 
  *  `request.Request.UpdatePolicyRequest.Policy.PolicySteps.Map<PolicySteps>.Steps[].Form.Form.Fields[].StringField.PasswordField.Multiline` **Added**
  *  `response.Policy.PolicySteps.Map<PolicySteps>.Steps[].Form.Form.Fields[].StringField.PasswordField.Multiline` **Added**
* `ConductoroneApi.Reporting.Get()`: `response` **Changed**
    - `LatestRun.SurfaceSnapshot.Components[].C1AppUserPicker` **Added**
    - `LatestRun.SurfaceSnapshot.Components[].C1CredentialOfferingPicker` **Added**
    - `LatestRun.SurfaceSnapshot.SavedReportId` **Added**
    - `Report.Description` **Added**
    - `Report.LatestRunSummary` **Added**
* `ConductoroneApi.Reporting.List()`: `response.List[]` **Changed**
    - `Description` **Added**
    - `LatestRunSummary` **Added**
* `ConductoroneApi.Reporting.Run()`: `response.Run.SurfaceSnapshot` **Changed**
    - `Components[].C1AppUserPicker` **Added**
    - `Components[].C1CredentialOfferingPicker` **Added**
    - `SavedReportId` **Added**
* `ConductoroneApi.Reporting.Save()`: 
  *  `request.Request.Description` **Added**
  * `response.Report` **Changed**
    - `Description` **Added**
    - `LatestRunSummary` **Added**
* `ConductoroneApi.Reporting.Update()`: 
  *  `request.Request.ReportingServiceUpdateRequest.Description` **Added**
  * `response.Report` **Changed**
    - `Description` **Added**
    - `LatestRunSummary` **Added**
* `ConductoroneApi.RequestSchema.Create()`: 
  *  `request.Request.Fields[].StringField.PasswordField.Multiline` **Added**
  *  `response.RequestSchema.Form.Fields[].StringField.PasswordField.Multiline` **Added**
* `ConductoroneApi.RequestSchema.Get()`:  `response.RequestSchema.Form.Fields[].StringField.PasswordField.Multiline` **Added**
* `ConductoroneApi.RequestSchema.Update()`: 
  *  `request.Request.RequestSchemaServiceUpdateRequest.RequestSchema.Form.Fields[].StringField.PasswordField.Multiline` **Added**
  *  `response.RequestSchema.Form.Fields[].StringField.PasswordField.Multiline` **Added**
* `ConductoroneApi.RoleMiningManagement.CreateAccessProfileFromCohort()`: 
  *  `request.Request.AccessScope` **Added**
* `ConductoroneApi.RoleMiningManagement.EvaluateEntitlementSelection()`: `response.CoreHolderFacets[]` **Changed**
    - `ValuesTruncated` **Added**
    - `ValuesUnavailable` **Added**
* `ConductoroneApi.RoleMiningManagement.GetCustomAnalysisResult()`: `response` **Changed**
    - `AccessScope` **Added**
    - `Facets[].ValuesTruncated` **Added**
    - `Facets[].ValuesUnavailable` **Added**
    - `ResourceScope` **Added**
* `ConductoroneApi.RoleMiningManagement.GetLatestCustomAnalysisResult()`: `response.Result` **Changed**
    - `AccessScope` **Added**
    - `ResourceScope` **Added**
* `ConductoroneApi.RoleMiningManagement.ListCustomAnalysisResults()`: `response.List[]` **Changed**
    - `AccessScope` **Added**
    - `ResourceScope` **Added**
* `ConductoroneApi.RoleMiningManagement.SearchCohortUsers()`: 
  *  `request.Request.SearchCohortUsersRequest.AccessScope` **Added**
  *  `response.TotalCount` **Added**
* `ConductoroneApi.RoleMiningManagement.TriggerCustomAnalysis()`: `request.Request` **Changed**
    - `AccessScope` **Added**
    - `ResourceScope` **Added**
* `ConductoroneApi.AppResourceSearch.SearchAppResources()`:  `response.List[].AppResource.IsManuallyManaged` **Added**
* `ConductoroneApi.AutomationSearch.SearchAutomationTemplateVersions()`: `response.List[].AutomationSteps[].CreateRevokeTasksV2` **Changed**
    - `ExclusionCriteria.ExcludedResourceTypeRefs` **Added**
    - `InclusionCriteria.ResourceTypeRefs` **Added**
* `ConductoroneApi.AutomationSearch.SearchAutomations()`: `response.List[].AutomationSteps[].CreateRevokeTasksV2` **Changed**
    - `ExclusionCriteria.ExcludedResourceTypeRefs` **Added**
    - `InclusionCriteria.ResourceTypeRefs` **Added**
* `ConductoroneApi.FunctionsSearch.Search()`: 
  *  `request.Request.SortOptions` **Added**
* `ConductoroneApi.PolicySearch.Search()`:  `response.List[].PolicySteps.Map<PolicySteps>.Steps[].Form.Form.Fields[].StringField.PasswordField.Multiline` **Added**
* `ConductoroneApi.PaperSecret.SearchMySecrets()`: 
  *  `request.Request.AccessRelationships` **Added**
* `ConductoroneApi.PaperSecret.SearchSecretsSharedWithMe()`: 
  *  `request.Request.SharingMode` **Added**
* `ConductoroneApi.TaskSearch.Search()`: `response.List[].Task` **Changed**
    - `Form.Fields[].StringField.PasswordField.Multiline` **Added**
    - `Type.Action.ScopeRole.AppUserId` **Added**
    - `Type.Action.Type.Enum(typeSpendRemedy)` **Added**
* `ConductoroneApi.WorkloadFederation.CreateProvider()`: 
  *  `request.Request.WellKnownProvider.Enum(wellKnownWorkloadProviderC1Edge)` **Added**
  * `response.Provider` **Changed**
    - `C1Edge` **Added**
    - `WellKnownProvider.Enum(wellKnownWorkloadProviderC1Edge)` **Added**
* `ConductoroneApi.WorkloadFederation.CreateTrust()`:  `response.Trust.C1EdgeService` **Added**
* `ConductoroneApi.WorkloadFederation.GetProvider()`: `response.Provider` **Changed**
    - `C1Edge` **Added**
    - `WellKnownProvider.Enum(wellKnownWorkloadProviderC1Edge)` **Added**
* `ConductoroneApi.WorkloadFederation.GetTrust()`:  `response.Trust.C1EdgeService` **Added**
* `ConductoroneApi.WorkloadFederation.ListProviders()`: `response.List[]` **Changed**
    - `C1Edge` **Added**
    - `WellKnownProvider.Enum(wellKnownWorkloadProviderC1Edge)` **Added**
* `ConductoroneApi.WorkloadFederation.ListTrusts()`:  `response.List[].C1EdgeService` **Added**
* `ConductoroneApi.WorkloadFederation.SearchTrusts()`:  `response.List[].C1EdgeService` **Added**
* `ConductoroneApi.WorkloadFederation.UpdateProvider()`: 
  *  `request.Request.WorkloadFederationServiceUpdateProviderRequest.Provider.C1Edge` **Added**
  * `response.Provider` **Changed**
    - `C1Edge` **Added**
    - `WellKnownProvider.Enum(wellKnownWorkloadProviderC1Edge)` **Added**
* `ConductoroneApi.WorkloadFederation.UpdateTrust()`: 
  *  `request.Request.WorkloadFederationServiceUpdateTrustRequest.Trust.C1EdgeService` **Added**
  *  `response.Trust.C1EdgeService` **Added**
* `ConductoroneApi.Principal.AddBinding()`: 
  * `request.Request.Subject` **Changed**
    - `AuthzenServer` **Added**
    - `Edge` **Added**
    - `SsoApplication` **Added**
* `ConductoroneApi.Principal.DeleteBinding()`: 
  * `request.Request.Subject` **Changed**
    - `AuthzenServer` **Added**
    - `Edge` **Added**
    - `SsoApplication` **Added**
* `ConductoroneApi.Principal.ListBindings()`: 
  * `request.Request.Subject` **Changed**
    - `AuthzenServer` **Added**
    - `Edge` **Added**
    - `SsoApplication` **Added**
* `ConductoroneApi.AiGovernanceSettings.Get()`: `response.AiGovernanceSettings` **Changed**
    - `UsageAttributionCostCenterAttribute` **Added**
    - `UsageAttributionTeamAttribute` **Added**
* `ConductoroneApi.AiGovernanceSettings.ListHistory()`: `response.List[].Snapshot` **Changed**
    - `UsageAttributionCostCenterAttribute` **Added**
    - `UsageAttributionTeamAttribute` **Added**
* `ConductoroneApi.AiGovernanceSettings.Update()`: 
  * `request.Request.AiGovernanceSettings` **Changed**
    - `UsageAttributionCostCenterAttribute` **Added**
    - `UsageAttributionTeamAttribute` **Added**
  * `response.AiGovernanceSettings` **Changed**
    - `UsageAttributionCostCenterAttribute` **Added**
    - `UsageAttributionTeamAttribute` **Added**
* `ConductoroneApi.OrgNotificationSettings.Get()`: `response.OrgNotificationSettings.ChannelSettings` **Changed**
    - `Email.Findings` **Added**
    - `Slack.Findings` **Added**
    - `Teams.Findings` **Added**
* `ConductoroneApi.OrgNotificationSettings.Update()`: 
  * `request.Request.ChannelSettings` **Changed**
    - `Email.Findings` **Added**
    - `Slack.Findings` **Added**
    - `Teams.Findings` **Added**
  * `response.OrgNotificationSettings.ChannelSettings` **Changed**
    - `Email.Findings` **Added**
    - `Slack.Findings` **Added**
    - `Teams.Findings` **Added**
* `ConductoroneApi.UserNotificationSettings.Get()`: `response.UserNotificationSettings.ChannelSettings` **Changed**
    - `Email.Findings` **Added**
    - `Slack.Findings` **Added**
    - `Teams.Findings` **Added**
* `ConductoroneApi.UserNotificationSettings.Update()`: 
  * `request.Request.ChannelSettings` **Changed**
    - `Email.Findings` **Added**
    - `Slack.Findings` **Added**
    - `Teams.Findings` **Added**
  * `response.UserNotificationSettings.ChannelSettings` **Changed**
    - `Email.Findings` **Added**
    - `Slack.Findings` **Added**
    - `Teams.Findings` **Added**
* `ConductoroneApi.SessionSettings.Get()`: `response.SessionSettings` **Changed**
    - `EnableIdleTimeout` **Added**
    - `IdleTimeout` **Added**
* `ConductoroneApi.SessionSettings.Update()`: 
  * `request.Request.SessionSettings` **Changed**
    - `EnableIdleTimeout` **Added**
    - `IdleTimeout` **Added**
  * `response.SessionSettings` **Changed**
    - `EnableIdleTimeout` **Added**
    - `IdleTimeout` **Added**
* `ConductoroneApi.SpendInsights.GetAttributionRollups()`: 
  *  `request.Request.Filters` **Added**
* `ConductoroneApi.Task.CreateActionTask()`: `response.TaskView.Task` **Changed**
    - `Form.Fields[].StringField.PasswordField.Multiline` **Added**
    - `Type.Action.ScopeRole.AppUserId` **Added**
    - `Type.Action.Type.Enum(typeSpendRemedy)` **Added**
* `ConductoroneApi.Task.CreateGrantTask()`: `response.TaskView.Task` **Changed**
    - `Form.Fields[].StringField.PasswordField.Multiline` **Added**
    - `Type.Action.ScopeRole.AppUserId` **Added**
    - `Type.Action.Type.Enum(typeSpendRemedy)` **Added**
* `ConductoroneApi.Task.CreateOffboardingTask()`: `response.TaskView.Task` **Changed**
    - `Form.Fields[].StringField.PasswordField.Multiline` **Added**
    - `Type.Action.ScopeRole.AppUserId` **Added**
    - `Type.Action.Type.Enum(typeSpendRemedy)` **Added**
* `ConductoroneApi.Task.CreateResourceActionTask()`: `response.TaskView.Task` **Changed**
    - `Form.Fields[].StringField.PasswordField.Multiline` **Added**
    - `Type.Action.ScopeRole.AppUserId` **Added**
    - `Type.Action.Type.Enum(typeSpendRemedy)` **Added**
* `ConductoroneApi.Task.CreateRevokeTask()`: `response.TaskView.Task` **Changed**
    - `Form.Fields[].StringField.PasswordField.Multiline` **Added**
    - `Type.Action.ScopeRole.AppUserId` **Added**
    - `Type.Action.Type.Enum(typeSpendRemedy)` **Added**
* `ConductoroneApi.Task.Get()`: `response.TaskView.Task` **Changed**
    - `Form.Fields[].StringField.PasswordField.Multiline` **Added**
    - `Type.Action.ScopeRole.AppUserId` **Added**
    - `Type.Action.Type.Enum(typeSpendRemedy)` **Added**
* `ConductoroneApi.TaskActions.Approve()`: `response.TaskView.Task` **Changed**
    - `Form.Fields[].StringField.PasswordField.Multiline` **Added**
    - `Type.Action.ScopeRole.AppUserId` **Added**
    - `Type.Action.Type.Enum(typeSpendRemedy)` **Added**
* `ConductoroneApi.TaskActions.ApproveWithStepUp()`: `response.TaskView.Task` **Changed**
    - `Form.Fields[].StringField.PasswordField.Multiline` **Added**
    - `Type.Action.ScopeRole.AppUserId` **Added**
    - `Type.Action.Type.Enum(typeSpendRemedy)` **Added**
* `ConductoroneApi.TaskActions.Close()`: `response.TaskView.Task` **Changed**
    - `Form.Fields[].StringField.PasswordField.Multiline` **Added**
    - `Type.Action.ScopeRole.AppUserId` **Added**
    - `Type.Action.Type.Enum(typeSpendRemedy)` **Added**
* `ConductoroneApi.TaskActions.Comment()`: `response.TaskView.Task` **Changed**
    - `Form.Fields[].StringField.PasswordField.Multiline` **Added**
    - `Type.Action.ScopeRole.AppUserId` **Added**
    - `Type.Action.Type.Enum(typeSpendRemedy)` **Added**
* `ConductoroneApi.TaskActions.Deny()`: `response.TaskView.Task` **Changed**
    - `Form.Fields[].StringField.PasswordField.Multiline` **Added**
    - `Type.Action.ScopeRole.AppUserId` **Added**
    - `Type.Action.Type.Enum(typeSpendRemedy)` **Added**
* `ConductoroneApi.TaskActions.EscalateToEmergencyAccess()`: `response.TaskView.Task` **Changed**
    - `Form.Fields[].StringField.PasswordField.Multiline` **Added**
    - `Type.Action.ScopeRole.AppUserId` **Added**
    - `Type.Action.Type.Enum(typeSpendRemedy)` **Added**
* `ConductoroneApi.TaskActions.HardReset()`: `response.TaskView.Task` **Changed**
    - `Form.Fields[].StringField.PasswordField.Multiline` **Added**
    - `Type.Action.ScopeRole.AppUserId` **Added**
    - `Type.Action.Type.Enum(typeSpendRemedy)` **Added**
* `ConductoroneApi.TaskActions.ProcessNow()`: `response.TaskView.Task` **Changed**
    - `Form.Fields[].StringField.PasswordField.Multiline` **Added**
    - `Type.Action.ScopeRole.AppUserId` **Added**
    - `Type.Action.Type.Enum(typeSpendRemedy)` **Added**
* `ConductoroneApi.TaskActions.Reassign()`: `response.TaskView.Task` **Changed**
    - `Form.Fields[].StringField.PasswordField.Multiline` **Added**
    - `Type.Action.ScopeRole.AppUserId` **Added**
    - `Type.Action.Type.Enum(typeSpendRemedy)` **Added**
* `ConductoroneApi.TaskActions.Restart()`: `response.TaskView.Task` **Changed**
    - `Form.Fields[].StringField.PasswordField.Multiline` **Added**
    - `Type.Action.ScopeRole.AppUserId` **Added**
    - `Type.Action.Type.Enum(typeSpendRemedy)` **Added**
* `ConductoroneApi.TaskActions.RetryProvisioning()`: `response.TaskView.Task` **Changed**
    - `Form.Fields[].StringField.PasswordField.Multiline` **Added**
    - `Type.Action.ScopeRole.AppUserId` **Added**
    - `Type.Action.Type.Enum(typeSpendRemedy)` **Added**
* `ConductoroneApi.TaskActions.SkipStep()`: `response.TaskView.Task` **Changed**
    - `Form.Fields[].StringField.PasswordField.Multiline` **Added**
    - `Type.Action.ScopeRole.AppUserId` **Added**
    - `Type.Action.Type.Enum(typeSpendRemedy)` **Added**
* `ConductoroneApi.TaskActions.UpdateGrantDuration()`: `response.TaskView.Task` **Changed**
    - `Form.Fields[].StringField.PasswordField.Multiline` **Added**
    - `Type.Action.ScopeRole.AppUserId` **Added**
    - `Type.Action.Type.Enum(typeSpendRemedy)` **Added**
* `ConductoroneApi.TaskActions.UpdateRequestData()`: `response.TaskView.Task` **Changed**
    - `Form.Fields[].StringField.PasswordField.Multiline` **Added**
    - `Type.Action.ScopeRole.AppUserId` **Added**
    - `Type.Action.Type.Enum(typeSpendRemedy)` **Added**
* `ConductoroneApi.Webhooks.Test()`: `response.Webhook` **Changed**
    - `State.Enum(webhookStateCanceled)` **Added**
    - `SupersededReason` **Added**
