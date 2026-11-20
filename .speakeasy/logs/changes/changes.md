## Go SDK Changes:
* `ConductoroneApi.AutomationExecution.ListAutomationExecutions()`:  `request.Request` **Added** (Breaking ⚠️)
* `ConductoroneApi.Automation.ListAutomations()`: 
  *  `request.Request` **Added** (Breaking ⚠️)
  * `response.List[].AutomationSteps[]` **Changed**
    - `CreateRevokeTasksV2.ExclusionCriteria.ExcludedResourceTypeRefs` **Added**
    - `CreateRevokeTasksV2.InclusionCriteria.ResourceTypeRefs` **Added**
    - `SendSlackMessage.AllowPrivateChannelFallback` **Added**
* `ConductoroneApi.Decoy.List()`:  `request.Request` **Added** (Breaking ⚠️)
* `ConductoroneApi.Functions.CreateFunction()`: 
  *  `request.Request.BrowserEnabled` **Removed** (Breaking ⚠️)
  * `response.Function` **Changed** (Breaking ⚠️)
    - `BrowserEnabled` **Removed** (Breaking ⚠️)
    - `CompileError` **Added**
    - `CompileStatus` **Added**
    - `NetworkConfig` **Added**
* `ConductoroneApi.Functions.GetFunction()`: `response.Function` **Changed** (Breaking ⚠️)
    - `BrowserEnabled` **Removed** (Breaking ⚠️)
    - `CompileError` **Added**
    - `CompileStatus` **Added**
    - `NetworkConfig` **Added**
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
* `ConductoroneApi.Principal.List()`:  `request.Request` **Added** (Breaking ⚠️)
* `ConductoroneApi.WorkloadFederation.ListProviders()`: 
  *  `request.Request` **Added** (Breaking ⚠️)
  * `response.List[]` **Changed**
    - `C1Edge` **Added**
    - `WellKnownProvider.Enum(wellKnownWorkloadProviderC1Edge)` **Added**
* `ConductoroneApi.StepUpProvider.List()`: 
  *  `request.Request` **Added** (Breaking ⚠️)
  *  `response.List[].Microsoft.RequireFreshAuthentication` **Added**
* `ConductoroneApi.FunctionsSearch.Search()`: 
  *  `request.Request.SortOptions` **Added**
  * `response.List[]` **Changed** (Breaking ⚠️)
    - `BrowserEnabled` **Removed** (Breaking ⚠️)
    - `CompileError` **Added**
    - `CompileStatus` **Added**
    - `NetworkConfig` **Added**
* `ConductoroneApi.GatewayKey.List()`:  `request.Request` **Added** (Breaking ⚠️)
* `ConductoroneApi.PersonalClient.List()`:  `request.Request` **Added** (Breaking ⚠️)
* `ConductoroneApi.Functions.UpdateFunction()`: 
  * `request.Request.Function` **Changed** (Breaking ⚠️)
    - `BrowserEnabled` **Removed** (Breaking ⚠️)
    - `NetworkConfig` **Added**
  * `response.Function` **Changed** (Breaking ⚠️)
    - `BrowserEnabled` **Removed** (Breaking ⚠️)
    - `CompileError` **Added**
    - `CompileStatus` **Added**
    - `NetworkConfig` **Added**
* `ConductoroneApi.Functions.ListFunctions()`: 
  *  `request.Request` **Added** (Breaking ⚠️)
  * `response.List[]` **Changed** (Breaking ⚠️)
    - `BrowserEnabled` **Removed** (Breaking ⚠️)
    - `CompileError` **Added**
    - `CompileStatus` **Added**
    - `NetworkConfig` **Added**
* `ConductoroneApi.PaperSecretAdmin.Revoke()`: **Removed** (Breaking ⚠️)
* `ConductoroneApi.EdgeIntegration.Create()`: **Added**
* `ConductoroneApi.AppUserMergeRules.Update()`: **Added**
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
* `ConductoroneApi.McpPrompt.Get()`: **Added**
* `ConductoroneApi.McpPrompt.List()`: **Added**
* `ConductoroneApi.McpPrompt.ListHistory()`: **Added**
* `ConductoroneApi.McpPrompt.Search()`: **Added**
* `ConductoroneApi.McpPrompt.Update()`: **Added**
* `ConductoroneApi.ConnectorSyncDiagnostics.CompareSyncStats()`: **Added**
* `ConductoroneApi.ConnectorSyncDiagnostics.GetSyncChangeTimeline()`: **Added**
* `ConductoroneApi.ConnectorSyncDiagnostics.GetSyncErrorDigest()`: **Added**
* `ConductoroneApi.ConnectorSyncDiagnostics.GetSyncHistory()`: **Added**
* `ConductoroneApi.ConnectorSyncDiagnostics.GetSyncStatus()`: **Added**
* `ConductoroneApi.ConnectorSyncDiagnostics.ListDebuggableConnectors()`: **Added**
* `ConductoroneApi.AccessReview.List()`: `response.List[].AccessReview` **Changed**
    - `CampaignSchedule` **Added**
    - `ColumnConfig.Columns[].Enum(accessReviewTaskColumnResourceDescription)` **Added**
    - `ColumnConfig.OrderedColumns[].Builtin.Enum(accessReviewTaskColumnResourceDescription)` **Added**
    - `DefaultView.Enum(accessReviewViewTypeByAccessProfile)` **Added**
    - `ErrorState.Enum(accessReviewErrorStateDuplicationFailed)` **Added**
    - `ScopeType.Enum(accessReviewScopeTypeByAccessProfiles)` **Added**
    - `ScopeV2.AccessProfileReviewDimensions` **Added**
    - `ScopeV2.AppSelectionCriteria.AnnotationMatches` **Added**
    - `ScopeV2.AppSelectionCriteria.AppIds` **Added**
    - `ScopeV2.AppSelectionCriteria.ResourceProfileMatches` **Added**
    - `ScopeV2.AppSelectionCriteria.ResourceTypeRefs` **Added**
    - `ScopeV2.ExcludedAppSelectionCriteria` **Added**
    - `ScopeV2.GrantsByCriteria.UsedWithinLast` **Added**
    - `ScopeV2.SpecificAccessProfiles` **Added**
* `ConductoroneApi.EdgeIntegration.Delete()`: **Added**
* `ConductoroneApi.EdgeIntegration.Get()`: **Added**
* `ConductoroneApi.EdgeIntegration.List()`: **Added**
* `ConductoroneApi.EdgeIntegration.ListUsages()`: **Added**
* `ConductoroneApi.EdgeIntegration.Update()`: **Added**
* `ConductoroneApi.Edge.CheckInferenceVaultAccess()`: **Added**
* `ConductoroneApi.Edge.CheckVaultAccess()`: **Added**
* `ConductoroneApi.Edge.Create()`: **Added**
* `ConductoroneApi.Edge.Delete()`: **Added**
* `ConductoroneApi.Edge.Get()`: **Added**
* `ConductoroneApi.Edge.GetAuditInclusionProof()`: **Added**
* `ConductoroneApi.Edge.GetEgressTrafficEvent()`: **Added**
* `ConductoroneApi.Edge.GetEgressTrafficSummary()`: **Added**
* `ConductoroneApi.Edge.GetEgressTrafficSummaryRollup()`: **Added**
* `ConductoroneApi.Edge.GetInferenceUsageAttribution()`: **Added**
* `ConductoroneApi.AccessReview.Get()`: `response.AccessReview.AccessReview` **Changed**
    - `CampaignSchedule` **Added**
    - `ColumnConfig.Columns[].Enum(accessReviewTaskColumnResourceDescription)` **Added**
    - `ColumnConfig.OrderedColumns[].Builtin.Enum(accessReviewTaskColumnResourceDescription)` **Added**
    - `DefaultView.Enum(accessReviewViewTypeByAccessProfile)` **Added**
    - `ErrorState.Enum(accessReviewErrorStateDuplicationFailed)` **Added**
    - `ScopeType.Enum(accessReviewScopeTypeByAccessProfiles)` **Added**
    - `ScopeV2.AccessProfileReviewDimensions` **Added**
    - `ScopeV2.AppSelectionCriteria.AnnotationMatches` **Added**
    - `ScopeV2.AppSelectionCriteria.AppIds` **Added**
    - `ScopeV2.AppSelectionCriteria.ResourceProfileMatches` **Added**
    - `ScopeV2.AppSelectionCriteria.ResourceTypeRefs` **Added**
    - `ScopeV2.ExcludedAppSelectionCriteria` **Added**
    - `ScopeV2.GrantsByCriteria.UsedWithinLast` **Added**
    - `ScopeV2.SpecificAccessProfiles` **Added**
* `ConductoroneApi.Edge.GetInferenceUsageSummary()`: **Added**
* `ConductoroneApi.Edge.GetInferenceUsageSummaryRollup()`: **Added**
* `ConductoroneApi.Edge.GetPublicTrust()`: **Added**
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
* `ConductoroneApi.AppResource.SearchAppEntitlementResourceBindings()`: **Added**
* `ConductoroneApi.McpServer.Disconnect()`: **Added**
* `ConductoroneApi.XaaResourceServer.ResyncScopes()`: **Added**
* `ConductoroneApi.Artifact.Delete()`: **Added**
* `ConductoroneApi.Artifact.Get()`: **Added**
* `ConductoroneApi.Artifact.GetVersion()`: **Added**
* `ConductoroneApi.Artifact.ListVersions()`: **Added**
* `ConductoroneApi.Artifact.Search()`: **Added**
* `ConductoroneApi.Artifact.Update()`: **Added**
* `ConductoroneApi.Automation.CancelPendingAutomationExecutions()`: **Added**
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
* `ConductoroneApi.ConnectorAuthoring.DeleteCatalogEntry()`: **Added**
* `ConductoroneApi.DirectoryMergeRule.ListDirectoryMergeRules()`: **Added**
* `ConductoroneApi.DirectoryMergeRule.SetDirectoryMergeRules()`: **Added**
* `ConductoroneApi.FindingSettings.GetEdgeFindingSettings()`: **Added**
* `ConductoroneApi.FindingSettings.UpdateEdgeFindingSettings()`: **Added**
* `ConductoroneApi.MyFundLimits.ClearTemporaryLimit()`: **Added**
* `ConductoroneApi.MyFundLimits.SetTemporaryLimit()`: **Added**
* `ConductoroneApi.ExternalOAuthClient.Create()`: **Added**
* `ConductoroneApi.ExternalOAuthClient.Revoke()`: **Added**
* `ConductoroneApi.ExternalOAuthClient.RotateSecret()`: **Added**
* `ConductoroneApi.ExternalOAuthClient.Update()`: **Added**
* `ConductoroneApi.ModelApiKey.Create()`: **Added**
* `ConductoroneApi.ModelApiKey.Get()`: **Added**
* `ConductoroneApi.ModelApiKey.List()`: **Added**
* `ConductoroneApi.ModelApiKey.Revoke()`: **Added**
* `ConductoroneApi.OAuthAuthorizationMy.Disconnect()`: **Added**
* `ConductoroneApi.OAuthAuthorizationMy.ListAvailable()`: **Added**
* `ConductoroneApi.OAuthAuthorizationMy.ListConnections()`: **Added**
* `ConductoroneApi.OAuthAuthorization.Create()`: **Added**
* `ConductoroneApi.OAuthAuthorization.Delete()`: **Added**
* `ConductoroneApi.OAuthAuthorization.Get()`: **Added**
* `ConductoroneApi.OAuthAuthorization.List()`: **Added**
* `ConductoroneApi.OAuthAuthorization.ListConsents()`: **Added**
* `ConductoroneApi.OAuthAuthorization.RevokeConsent()`: **Added**
* `ConductoroneApi.OAuthAuthorization.RotateSecret()`: **Added**
* `ConductoroneApi.OAuthAuthorization.SearchConsents()`: **Added**
* `ConductoroneApi.OAuthAuthorization.Update()`: **Added**
* `ConductoroneApi.OAuthAuthorization.UpdateClient()`: **Added**
* `ConductoroneApi.RateLimit.Get()`: **Added**
* `ConductoroneApi.RateLimit.GetUsageReport()`: **Added**
* `ConductoroneApi.ApiKeyUsage.BatchGet()`: **Added**
* `ConductoroneApi.MyRateLimit.Get()`: **Added**
* `ConductoroneApi.RoleMiningManagement.CountEntitlementSelectionFilterAttributeValues()`: **Added**
* `ConductoroneApi.RoleMiningManagement.ListCohortFilterAttributeValues()`: **Added**
* `ConductoroneApi.RoleMiningManagement.ListCohortFilterAttributes()`: **Added**
* `ConductoroneApi.AppSecretSearch.SearchMyVendedCredentials()`: **Added**
* `ConductoroneApi.ApiKeySearch.Search()`: **Added**
* `ConductoroneApi.VirtualMcpServerMySearch.Search()`: **Added**
* `ConductoroneApi.RequestCatalogSearch.SearchRequestableCredentialOfferings()`: **Added**
* `ConductoroneApi.PaperSecretAdmin.Delete()`: **Added**
* `ConductoroneApi.AccessReview.Update()`: 
  * `request.Request.AccessReviewServiceUpdateRequest.AccessReview` **Changed**
    - `CampaignSchedule` **Added**
    - `ColumnConfig.Columns[].Enum(accessReviewTaskColumnResourceDescription)` **Added**
    - `ColumnConfig.OrderedColumns[].Builtin.Enum(accessReviewTaskColumnResourceDescription)` **Added**
    - `DefaultView.Enum(accessReviewViewTypeByAccessProfile)` **Added**
    - `ScopeType.Enum(accessReviewScopeTypeByAccessProfiles)` **Added**
    - `ScopeV2.AccessProfileReviewDimensions` **Added**
    - `ScopeV2.AppSelectionCriteria.AnnotationMatches` **Added**
    - `ScopeV2.AppSelectionCriteria.AppIds` **Added**
    - `ScopeV2.AppSelectionCriteria.ResourceProfileMatches` **Added**
    - `ScopeV2.AppSelectionCriteria.ResourceTypeRefs` **Added**
    - `ScopeV2.ExcludedAppSelectionCriteria` **Added**
    - `ScopeV2.GrantsByCriteria.UsedWithinLast` **Added**
    - `ScopeV2.SpecificAccessProfiles` **Added**
  * `response.AccessReview.AccessReview` **Changed**
    - `CampaignSchedule` **Added**
    - `ColumnConfig.Columns[].Enum(accessReviewTaskColumnResourceDescription)` **Added**
    - `ColumnConfig.OrderedColumns[].Builtin.Enum(accessReviewTaskColumnResourceDescription)` **Added**
    - `DefaultView.Enum(accessReviewViewTypeByAccessProfile)` **Added**
    - `ErrorState.Enum(accessReviewErrorStateDuplicationFailed)` **Added**
    - `ScopeType.Enum(accessReviewScopeTypeByAccessProfiles)` **Added**
    - `ScopeV2.AccessProfileReviewDimensions` **Added**
    - `ScopeV2.AppSelectionCriteria.AnnotationMatches` **Added**
    - `ScopeV2.AppSelectionCriteria.AppIds` **Added**
    - `ScopeV2.AppSelectionCriteria.ResourceProfileMatches` **Added**
    - `ScopeV2.AppSelectionCriteria.ResourceTypeRefs` **Added**
    - `ScopeV2.ExcludedAppSelectionCriteria` **Added**
    - `ScopeV2.GrantsByCriteria.UsedWithinLast` **Added**
    - `ScopeV2.SpecificAccessProfiles` **Added**
* `ConductoroneApi.ShadowMcpOccurrence.Search()`: **Added**
* `ConductoroneApi.ToolGatesSearch.Search()`: **Added**
* `ConductoroneApi.AgentClassifier.GetAgentClassifierPolicy()`: **Added**
* `ConductoroneApi.AgentClassifier.UpdateAgentClassifierPolicy()`: **Added**
* `ConductoroneApi.EmailCapabilities.Get()`: **Added**
* `ConductoroneApi.GoLink.Create()`: **Added**
* `ConductoroneApi.GoLink.Delete()`: **Added**
* `ConductoroneApi.GoLink.Get()`: **Added**
* `ConductoroneApi.GoLink.List()`: **Added**
* `ConductoroneApi.GoLink.ListVersions()`: **Added**
* `ConductoroneApi.GoLink.Resolve()`: **Added**
* `ConductoroneApi.GoLink.Update()`: **Added**
* `ConductoroneApi.GoLinkSearch.Search()`: **Added**
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
* `ConductoroneApi.AppUserMergeRules.Get()`: **Added**
* `ConductoroneApi.PaperSecret.Revoke()`: **Removed** (Breaking ⚠️)
* `ConductoroneApi.TenantEmailProvider.GetEmailCapabilities()`: **Removed** (Breaking ⚠️)
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
    - `ScopeV2.AppSelectionCriteria.AnnotationMatches` **Added**
    - `ScopeV2.AppSelectionCriteria.AppIds` **Added**
    - `ScopeV2.AppSelectionCriteria.ResourceProfileMatches` **Added**
    - `ScopeV2.AppSelectionCriteria.ResourceTypeRefs` **Added**
    - `ScopeV2.ExcludedAppSelectionCriteria` **Added**
    - `ScopeV2.GrantsByCriteria.UsedWithinLast` **Added**
    - `ScopeV2.SpecificAccessProfiles` **Added**
  * `response.AccessReview.AccessReview` **Changed**
    - `CampaignSchedule` **Added**
    - `ColumnConfig.Columns[].Enum(accessReviewTaskColumnResourceDescription)` **Added**
    - `ColumnConfig.OrderedColumns[].Builtin.Enum(accessReviewTaskColumnResourceDescription)` **Added**
    - `DefaultView.Enum(accessReviewViewTypeByAccessProfile)` **Added**
    - `ErrorState.Enum(accessReviewErrorStateDuplicationFailed)` **Added**
    - `ScopeType.Enum(accessReviewScopeTypeByAccessProfiles)` **Added**
    - `ScopeV2.AccessProfileReviewDimensions` **Added**
    - `ScopeV2.AppSelectionCriteria.AnnotationMatches` **Added**
    - `ScopeV2.AppSelectionCriteria.AppIds` **Added**
    - `ScopeV2.AppSelectionCriteria.ResourceProfileMatches` **Added**
    - `ScopeV2.AppSelectionCriteria.ResourceTypeRefs` **Added**
    - `ScopeV2.ExcludedAppSelectionCriteria` **Added**
    - `ScopeV2.GrantsByCriteria.UsedWithinLast` **Added**
    - `ScopeV2.SpecificAccessProfiles` **Added**
* `ConductoroneApi.Edge.GetInferenceUsageAttributionRollup()`: **Added**
* `ConductoroneApi.AppUserMergeRules.ResetToDefaults()`: **Added**
* `ConductoroneApi.PaperSecret.Delete()`: **Added**
* `ConductoroneApi.AccessReviewActions.GenerateReport()`: 
  * `request.Request.AccessReviewActionsServiceGenerateReportRequest.ReportColumnConfig.Columns[]` **Changed**
    - `Enum(accessReviewReportColumnAccessProfile)` **Added**
    - `Enum(accessReviewReportColumnRemediationTask)` **Added**
    - `Enum(accessReviewReportColumnReviewItem)` **Added**
    - `Enum(accessReviewReportColumnReviewItemType)` **Added**
* `ConductoroneApi.AccessReviewSetupEntitlement.GetCampaignScopeAndEntitlements()`: `response` **Changed**
    - `ExcludedEntitlements` **Added**
    - `ExcludedResourceTypes` **Added**
    - `ScopeV2.AccessProfileReviewDimensions` **Added**
    - `ScopeV2.AppSelectionCriteria.AnnotationMatches` **Added**
    - `ScopeV2.AppSelectionCriteria.AppIds` **Added**
    - `ScopeV2.AppSelectionCriteria.ResourceProfileMatches` **Added**
    - `ScopeV2.AppSelectionCriteria.ResourceTypeRefs` **Added**
    - `ScopeV2.ExcludedAppSelectionCriteria` **Added**
    - `ScopeV2.GrantsByCriteria.UsedWithinLast` **Added**
    - `ScopeV2.SpecificAccessProfiles` **Added**
* `ConductoroneApi.AccessReviewSetupEntitlement.SetCampaignScopeAndEntitlements()`: 
  * `request.Request.AccessReviewSetupEntitlementAndScopeServiceSetRequest.ScopeV2` **Changed**
    - `AccessProfileReviewDimensions` **Added**
    - `AppSelectionCriteria.AnnotationMatches` **Added**
    - `AppSelectionCriteria.AppIds` **Added**
    - `AppSelectionCriteria.ResourceProfileMatches` **Added**
    - `AppSelectionCriteria.ResourceTypeRefs` **Added**
    - `ExcludedAppSelectionCriteria` **Added**
    - `GrantsByCriteria.UsedWithinLast` **Added**
    - `SpecificAccessProfiles` **Added**
  * `response` **Changed**
    - `ExcludedEntitlements` **Added**
    - `ExcludedResourceTypes` **Added**
    - `ScopeV2.AccessProfileReviewDimensions` **Added**
    - `ScopeV2.AppSelectionCriteria.AnnotationMatches` **Added**
    - `ScopeV2.AppSelectionCriteria.AppIds` **Added**
    - `ScopeV2.AppSelectionCriteria.ResourceProfileMatches` **Added**
    - `ScopeV2.AppSelectionCriteria.ResourceTypeRefs` **Added**
    - `ScopeV2.ExcludedAppSelectionCriteria` **Added**
    - `ScopeV2.GrantsByCriteria.UsedWithinLast` **Added**
    - `ScopeV2.SpecificAccessProfiles` **Added**
* `ConductoroneApi.AccessReviewSetupEntitlement.SetCampaignScopeByResourceType()`: 
  * `request.Request.AccessReviewSetScopeByResourceTypeRequest.ScopeV2` **Changed**
    - `AccessProfileReviewDimensions` **Added**
    - `AppSelectionCriteria.AnnotationMatches` **Added**
    - `AppSelectionCriteria.AppIds` **Added**
    - `AppSelectionCriteria.ResourceProfileMatches` **Added**
    - `AppSelectionCriteria.ResourceTypeRefs` **Added**
    - `ExcludedAppSelectionCriteria` **Added**
    - `GrantsByCriteria.UsedWithinLast` **Added**
    - `SpecificAccessProfiles` **Added**
* `ConductoroneApi.AccessReviewTemplate.Create()`: 
  * `request.Request` **Changed**
    - `ColumnConfig.Columns[].Enum(accessReviewTaskColumnResourceDescription)` **Added**
    - `ColumnConfig.OrderedColumns[].Builtin.Enum(accessReviewTaskColumnResourceDescription)` **Added**
    - `DefaultView.Enum(accessReviewViewTypeByAccessProfile)` **Added**
    - `Scope.AccessProfileReviewDimensions` **Added**
    - `Scope.AppSelectionCriteria.AnnotationMatches` **Added**
    - `Scope.AppSelectionCriteria.AppIds` **Added**
    - `Scope.AppSelectionCriteria.ResourceProfileMatches` **Added**
    - `Scope.AppSelectionCriteria.ResourceTypeRefs` **Added**
    - `Scope.ExcludedAppSelectionCriteria` **Added**
    - `Scope.GrantsByCriteria.UsedWithinLast` **Added**
    - `Scope.SpecificAccessProfiles` **Added**
    - `ScopeType.Enum(accessReviewScopeTypeByAccessProfiles)` **Added**
  * `response.AccessReviewTemplate` **Changed**
    - `CampaignSchedule` **Added**
    - `ColumnConfig.Columns[].Enum(accessReviewTaskColumnResourceDescription)` **Added**
    - `ColumnConfig.OrderedColumns[].Builtin.Enum(accessReviewTaskColumnResourceDescription)` **Added**
    - `DefaultView.Enum(accessReviewViewTypeByAccessProfile)` **Added**
    - `Scope.AccessProfileReviewDimensions` **Added**
    - `Scope.AppSelectionCriteria.AnnotationMatches` **Added**
    - `Scope.AppSelectionCriteria.AppIds` **Added**
    - `Scope.AppSelectionCriteria.ResourceProfileMatches` **Added**
    - `Scope.AppSelectionCriteria.ResourceTypeRefs` **Added**
    - `Scope.ExcludedAppSelectionCriteria` **Added**
    - `Scope.GrantsByCriteria.UsedWithinLast` **Added**
    - `Scope.SpecificAccessProfiles` **Added**
    - `ScopeType.Enum(accessReviewScopeTypeByAccessProfiles)` **Added**
    - `SlackChannel.AllowPrivateChannelFallback` **Added**
* `ConductoroneApi.AccessReviewTemplate.Get()`: `response.AccessReviewTemplate` **Changed**
    - `CampaignSchedule` **Added**
    - `ColumnConfig.Columns[].Enum(accessReviewTaskColumnResourceDescription)` **Added**
    - `ColumnConfig.OrderedColumns[].Builtin.Enum(accessReviewTaskColumnResourceDescription)` **Added**
    - `DefaultView.Enum(accessReviewViewTypeByAccessProfile)` **Added**
    - `Scope.AccessProfileReviewDimensions` **Added**
    - `Scope.AppSelectionCriteria.AnnotationMatches` **Added**
    - `Scope.AppSelectionCriteria.AppIds` **Added**
    - `Scope.AppSelectionCriteria.ResourceProfileMatches` **Added**
    - `Scope.AppSelectionCriteria.ResourceTypeRefs` **Added**
    - `Scope.ExcludedAppSelectionCriteria` **Added**
    - `Scope.GrantsByCriteria.UsedWithinLast` **Added**
    - `Scope.SpecificAccessProfiles` **Added**
    - `ScopeType.Enum(accessReviewScopeTypeByAccessProfiles)` **Added**
    - `SlackChannel.AllowPrivateChannelFallback` **Added**
* `ConductoroneApi.AccessReviewTemplate.Update()`: 
  * `request.Request.AccessReviewTemplateServiceUpdateRequest.AccessReviewTemplate` **Changed**
    - `CampaignSchedule` **Added**
    - `ColumnConfig.Columns[].Enum(accessReviewTaskColumnResourceDescription)` **Added**
    - `ColumnConfig.OrderedColumns[].Builtin.Enum(accessReviewTaskColumnResourceDescription)` **Added**
    - `DefaultView.Enum(accessReviewViewTypeByAccessProfile)` **Added**
    - `Scope.AccessProfileReviewDimensions` **Added**
    - `Scope.AppSelectionCriteria.AnnotationMatches` **Added**
    - `Scope.AppSelectionCriteria.AppIds` **Added**
    - `Scope.AppSelectionCriteria.ResourceProfileMatches` **Added**
    - `Scope.AppSelectionCriteria.ResourceTypeRefs` **Added**
    - `Scope.ExcludedAppSelectionCriteria` **Added**
    - `Scope.GrantsByCriteria.UsedWithinLast` **Added**
    - `Scope.SpecificAccessProfiles` **Added**
    - `ScopeType.Enum(accessReviewScopeTypeByAccessProfiles)` **Added**
    - `SlackChannel.AllowPrivateChannelFallback` **Added**
  * `response.AccessReviewTemplate` **Changed**
    - `CampaignSchedule` **Added**
    - `ColumnConfig.Columns[].Enum(accessReviewTaskColumnResourceDescription)` **Added**
    - `ColumnConfig.OrderedColumns[].Builtin.Enum(accessReviewTaskColumnResourceDescription)` **Added**
    - `DefaultView.Enum(accessReviewViewTypeByAccessProfile)` **Added**
    - `Scope.AccessProfileReviewDimensions` **Added**
    - `Scope.AppSelectionCriteria.AnnotationMatches` **Added**
    - `Scope.AppSelectionCriteria.AppIds` **Added**
    - `Scope.AppSelectionCriteria.ResourceProfileMatches` **Added**
    - `Scope.AppSelectionCriteria.ResourceTypeRefs` **Added**
    - `Scope.ExcludedAppSelectionCriteria` **Added**
    - `Scope.GrantsByCriteria.UsedWithinLast` **Added**
    - `Scope.SpecificAccessProfiles` **Added**
    - `ScopeType.Enum(accessReviewScopeTypeByAccessProfiles)` **Added**
    - `SlackChannel.AllowPrivateChannelFallback` **Added**
* `ConductoroneApi.AccessReviewTemplateSetupEntitlement.GetScopeAndEntitlements()`: `response` **Changed**
    - `ExcludedEntitlements` **Added**
    - `ExcludedResourceTypes` **Added**
    - `Scope.AccessProfileReviewDimensions` **Added**
    - `Scope.AppSelectionCriteria.AnnotationMatches` **Added**
    - `Scope.AppSelectionCriteria.AppIds` **Added**
    - `Scope.AppSelectionCriteria.ResourceProfileMatches` **Added**
    - `Scope.AppSelectionCriteria.ResourceTypeRefs` **Added**
    - `Scope.ExcludedAppSelectionCriteria` **Added**
    - `Scope.GrantsByCriteria.UsedWithinLast` **Added**
    - `Scope.SpecificAccessProfiles` **Added**
* `ConductoroneApi.AccessReviewTemplateSetupEntitlement.SetScopeAndEntitlements()`: 
  * `request.Request.AccessReviewTemplateSetupEntitlementServiceSetRequest.Scope` **Changed**
    - `AccessProfileReviewDimensions` **Added**
    - `AppSelectionCriteria.AnnotationMatches` **Added**
    - `AppSelectionCriteria.AppIds` **Added**
    - `AppSelectionCriteria.ResourceProfileMatches` **Added**
    - `AppSelectionCriteria.ResourceTypeRefs` **Added**
    - `ExcludedAppSelectionCriteria` **Added**
    - `GrantsByCriteria.UsedWithinLast` **Added**
    - `SpecificAccessProfiles` **Added**
  * `response` **Changed**
    - `ExcludedEntitlements` **Added**
    - `ExcludedResourceTypes` **Added**
    - `Scope.AccessProfileReviewDimensions` **Added**
    - `Scope.AppSelectionCriteria.AnnotationMatches` **Added**
    - `Scope.AppSelectionCriteria.AppIds` **Added**
    - `Scope.AppSelectionCriteria.ResourceProfileMatches` **Added**
    - `Scope.AppSelectionCriteria.ResourceTypeRefs` **Added**
    - `Scope.ExcludedAppSelectionCriteria` **Added**
    - `Scope.GrantsByCriteria.UsedWithinLast` **Added**
    - `Scope.SpecificAccessProfiles` **Added**
* `ConductoroneApi.AccessReviewTemplateSetupEntitlement.SetScopeByResourceType()`: 
  * `request.Request.AccessReviewTemplateSetScopeByResourceTypeRequest.Scope` **Changed**
    - `AccessProfileReviewDimensions` **Added**
    - `AppSelectionCriteria.AnnotationMatches` **Added**
    - `AppSelectionCriteria.AppIds` **Added**
    - `AppSelectionCriteria.ResourceProfileMatches` **Added**
    - `AppSelectionCriteria.ResourceTypeRefs` **Added**
    - `ExcludedAppSelectionCriteria` **Added**
    - `GrantsByCriteria.UsedWithinLast` **Added**
    - `SpecificAccessProfiles` **Added**
* `ConductoroneApi.AccessConflict.CreateMonitor()`: 
  *  `request.Request.NotificationConfig.SlackNotifications.AllowPrivateChannelFallback` **Added**
  *  `response.NotificationConfig.SlackNotifications.AllowPrivateChannelFallback` **Added**
* `ConductoroneApi.AccessConflict.GetMonitor()`:  `response.NotificationConfig.SlackNotifications.AllowPrivateChannelFallback` **Added**
* `ConductoroneApi.AccessConflict.UpdateMonitor()`: 
  *  `request.Request.ConflictMonitorUpdateRequest.NotificationConfig.SlackNotifications.AllowPrivateChannelFallback` **Added**
  *  `response.NotificationConfig.SlackNotifications.AllowPrivateChannelFallback` **Added**
* `ConductoroneApi.Apps.Create()`:  `response.App.EntitlementsUseResourceDescription` **Added**
* `ConductoroneApi.Apps.Get()`:  `response.App.EntitlementsUseResourceDescription` **Added**
* `ConductoroneApi.Apps.List()`:  `response.List[].EntitlementsUseResourceDescription` **Added**
* `ConductoroneApi.Apps.Update()`: 
  *  `request.Request.UpdateAppRequest.App.EntitlementsUseResourceDescription` **Added**
  *  `response.App.EntitlementsUseResourceDescription` **Added**
* `ConductoroneApi.Connector.Create()`:  `response.ConnectorView.Connector.Status.ErrorKind` **Added**
* `ConductoroneApi.Connector.CreateDelegated()`:  `response.ConnectorView.Connector.Status.ErrorKind` **Added**
* `ConductoroneApi.Connector.Get()`:  `response.ConnectorView.Connector.Status.ErrorKind` **Added**
* `ConductoroneApi.Connector.List()`:  `response.List[].Connector.Status.ErrorKind` **Added**
* `ConductoroneApi.Connector.Update()`: 
  * `request.Request.ConnectorServiceUpdateRequest` **Changed**
    - `Connector.Status.ErrorKind` **Added**
    - `Justification` **Added**
  *  `response.ConnectorView.Connector.Status.ErrorKind` **Added**
* `ConductoroneApi.Connector.UpdateConnectorSchedule()`: 
  *  `request.Request.UpdateConnectorScheduleRequest.Justification` **Added**
* `ConductoroneApi.Connector.UpdateDelegated()`: 
  *  `request.Request.ConnectorServiceUpdateDelegatedRequest.Connector.Status.ErrorKind` **Added**
  *  `response.ConnectorView.Connector.Status.ErrorKind` **Added**
* `ConductoroneApi.AppAccessRequestsDefaults.CancelAppAccessRequestsDefaults()`:  `response.CatalogMode` **Added**
* `ConductoroneApi.AppAccessRequestsDefaults.CreateAppAccessRequestsDefaults()`: 
  *  `request.Request.AppAccessRequestDefaults.CatalogMode` **Added**
  *  `response.CatalogMode` **Added**
* `ConductoroneApi.AppAccessRequestsDefaults.GetAppAccessRequestsDefaults()`:  `response.CatalogMode` **Added**
* `ConductoroneApi.McpResource.Search()`: 
  *  `request.Request.McpResourceServiceSearchRequest.Refs` **Added**
* `ConductoroneApi.AppEntitlementRoutingRule.ListAppEntitlementRoutingRules()`:  `response.MaxRulesPerApp` **Added**
* `ConductoroneApi.AppEntitlements.Create()`: 
  *  `request.Request.CreateAppEntitlementRequest.DeprovisionerPolicy` **Added**
* `ConductoroneApi.AppEntitlements.ListAutomationExclusions()`: `request.Request` **Changed**
    - `PageSize` **Added**
    - `PageToken` **Added**
* `ConductoroneApi.AppEntitlementSearch.Search()`: `request.Request` **Changed**
    - `AllowAutomationEnabledC1Groups` **Added**
    - `AnnotationMatches` **Added**
    - `AppResourceTypeRefs` **Added**
    - `EntitlementSlugs` **Added**
    - `ExcludeAutomationEnabled` **Added**
    - `ExcludeNonRoleProfileMemberships` **Added**
    - `OwnerUserIds` **Added**
    - `ResourceProfileMatches` **Added**
* `ConductoroneApi.McpServer.GetCatalog()`: `response.CatalogEntry` **Changed**
    - `AuthValidationUrl` **Added**
    - `ConnectType` **Added**
* `ConductoroneApi.McpServer.ListCatalog()`: `response.List[]` **Changed**
    - `AuthValidationUrl` **Added**
    - `ConnectType` **Added**
* `ConductoroneApi.McpServer.Register()`: 
  * `request.Request.McpServerServiceRegisterRequest` **Changed**
    - `AccessProvisioning` **Added**
    - `ToolEnablement` **Added**
* `ConductoroneApi.McpServer.TestConnection()`: 
  *  `request.Request.HostedConfig` **Added**
  *  `response.CredentialVerified` **Added**
* `ConductoroneApi.AppResource.CreateManuallyManagedAppResource()`:  `response.AppResource.IsManuallyManaged` **Added**
* `ConductoroneApi.AppResource.Get()`:  `response.AppResourceView.AppResource.IsManuallyManaged` **Added**
* `ConductoroneApi.AppResource.List()`:  `response.List[].AppResource.IsManuallyManaged` **Added**
* `ConductoroneApi.AppResource.Update()`:  `response.AppResourceView.AppResource.IsManuallyManaged` **Added**
* `ConductoroneApi.XaaResourceServer.Get()`:  `response.ScopeDiscoveryStatus` **Added**
* `ConductoroneApi.Automation.CreateAutomation()`: 
  * `request.Request.AutomationSteps[]` **Changed**
    - `CreateRevokeTasksV2.ExclusionCriteria.ExcludedResourceTypeRefs` **Added**
    - `CreateRevokeTasksV2.InclusionCriteria.ResourceTypeRefs` **Added**
    - `SendSlackMessage.AllowPrivateChannelFallback` **Added**
  * `response.Automation.AutomationSteps[]` **Changed**
    - `CreateRevokeTasksV2.ExclusionCriteria.ExcludedResourceTypeRefs` **Added**
    - `CreateRevokeTasksV2.InclusionCriteria.ResourceTypeRefs` **Added**
    - `SendSlackMessage.AllowPrivateChannelFallback` **Added**
* `ConductoroneApi.Automation.GetAutomation()`: `response.Automation.AutomationSteps[]` **Changed**
    - `CreateRevokeTasksV2.ExclusionCriteria.ExcludedResourceTypeRefs` **Added**
    - `CreateRevokeTasksV2.InclusionCriteria.ResourceTypeRefs` **Added**
    - `SendSlackMessage.AllowPrivateChannelFallback` **Added**
* `ConductoroneApi.Automation.UpdateAutomation()`: 
  * `request.Request.UpdateAutomationRequest.Automation.AutomationSteps[]` **Changed**
    - `CreateRevokeTasksV2.ExclusionCriteria.ExcludedResourceTypeRefs` **Added**
    - `CreateRevokeTasksV2.InclusionCriteria.ResourceTypeRefs` **Added**
    - `SendSlackMessage.AllowPrivateChannelFallback` **Added**
  * `response` **Changed**
    - `Automation.AutomationSteps[].CreateRevokeTasksV2.ExclusionCriteria.ExcludedResourceTypeRefs` **Added**
    - `Automation.AutomationSteps[].CreateRevokeTasksV2.InclusionCriteria.ResourceTypeRefs` **Added**
    - `Automation.AutomationSteps[].SendSlackMessage.AllowPrivateChannelFallback` **Added**
    - `BulkActionId` **Added**
* `ConductoroneApi.RequestCatalogManagement.PlanTypeChange()`: `response.Impacts[].Code` **Changed**
    - `Enum(requestCatalogTypeChangeImpactCodeContainingCatalogs)` **Added**
    - `Enum(requestCatalogTypeChangeImpactCodeRoleMembershipEntitlements)` **Added**
* `ConductoroneApi.ConnectorAuthoringActivation.ActivateRevision()`:  `response.DeploymentEventId` **Added**
* `ConductoroneApi.ConnectorAuthoringActivation.RollbackRevision()`:  `response.DeploymentEventId` **Added**
* `ConductoroneApi.ConnectorCatalog.ConfigurationSchema()`: `response` **Changed**
    - `FormSchema.FieldRelationships[].DependentOn.Dependencies` **Added**
    - `FormSchema.Fields[].StringField.PasswordField.Multiline` **Added**
    - `FormSchema.Fields[].StringField.RegexErrorMessage` **Added**
    - `FormSchema.Fields[].StringMapField.Secret` **Added**
    - `FormSchema.Fields[].StringSliceField.MultiSelectField` **Added**
    - `Schema.Fields[].StringMap.Secret` **Added**
* `ConductoroneApi.Directory.Create()`: 
  *  `request.Request.MergeRules` **Added**
* `ConductoroneApi.Directory.Update()`: 
  *  `request.Request.DirectoryServiceUpdateRequest.MergeRules` **Added**
* `ConductoroneApi.Finding.BulkCreateFindingTasks()`: 
  * `request.Request.SearchRequest` **Changed**
    - `EdgeIds` **Added**
    - `ExcludeFindingTypes` **Added**
    - `FindingTypes[].Enum(findingTypeAccessConflict)` **Added**
    - `FindingTypes[].Enum(findingTypeConnectorHostOutdated)` **Added**
    - `FindingTypes[].Enum(findingTypeEdge)` **Added**
    - `FindingTypes[].Enum(findingTypeMcpGatewayToolCallRisk)` **Added**
    - `FindingTypes[].Enum(findingTypeShadowApp)` **Added**
    - `SourceDetectorIds` **Added**
    - `TargetIdentityUserIds` **Added**
* `ConductoroneApi.Finding.BulkUpdateFindingState()`: 
  * `request.Request.SearchRequest` **Changed**
    - `EdgeIds` **Added**
    - `ExcludeFindingTypes` **Added**
    - `FindingTypes[].Enum(findingTypeAccessConflict)` **Added**
    - `FindingTypes[].Enum(findingTypeConnectorHostOutdated)` **Added**
    - `FindingTypes[].Enum(findingTypeEdge)` **Added**
    - `FindingTypes[].Enum(findingTypeMcpGatewayToolCallRisk)` **Added**
    - `FindingTypes[].Enum(findingTypeShadowApp)` **Added**
    - `SourceDetectorIds` **Added**
    - `TargetIdentityUserIds` **Added**
* `ConductoroneApi.Finding.CreateFinding()`: `response.Finding` **Changed**
    - `AccessConflictEvidence` **Added**
    - `AccessConflict` **Added**
    - `ConnectorHostOutdatedEvidence` **Added**
    - `ConnectorHostOutdated` **Added**
    - `CredentialExpiring.AppResourceSecret` **Added**
    - `EdgeFindingEvidence` **Added**
    - `EdgeFinding` **Added**
    - `McpGatewayToolCallRiskEvidence` **Added**
    - `McpGatewayToolCallRisk` **Added**
    - `NhiUnownedEvidence` **Added**
    - `ShadowApp` **Added**
    - `ShadowMcp.Transport` **Added**
    - `UsageResourceTarget` **Added**
* `ConductoroneApi.Finding.CreateFindingTask()`: `response.Finding` **Changed**
    - `AccessConflictEvidence` **Added**
    - `AccessConflict` **Added**
    - `ConnectorHostOutdatedEvidence` **Added**
    - `ConnectorHostOutdated` **Added**
    - `CredentialExpiring.AppResourceSecret` **Added**
    - `EdgeFindingEvidence` **Added**
    - `EdgeFinding` **Added**
    - `McpGatewayToolCallRiskEvidence` **Added**
    - `McpGatewayToolCallRisk` **Added**
    - `NhiUnownedEvidence` **Added**
    - `ShadowApp` **Added**
    - `ShadowMcp.Transport` **Added**
    - `UsageResourceTarget` **Added**
* `ConductoroneApi.Finding.GetFinding()`: `response.Finding` **Changed**
    - `AccessConflictEvidence` **Added**
    - `AccessConflict` **Added**
    - `ConnectorHostOutdatedEvidence` **Added**
    - `ConnectorHostOutdated` **Added**
    - `CredentialExpiring.AppResourceSecret` **Added**
    - `EdgeFindingEvidence` **Added**
    - `EdgeFinding` **Added**
    - `McpGatewayToolCallRiskEvidence` **Added**
    - `McpGatewayToolCallRisk` **Added**
    - `NhiUnownedEvidence` **Added**
    - `ShadowApp` **Added**
    - `ShadowMcp.Transport` **Added**
    - `UsageResourceTarget` **Added**
* `ConductoroneApi.Finding.UpdateFindingAssignee()`: `response.Finding` **Changed**
    - `AccessConflictEvidence` **Added**
    - `AccessConflict` **Added**
    - `ConnectorHostOutdatedEvidence` **Added**
    - `ConnectorHostOutdated` **Added**
    - `CredentialExpiring.AppResourceSecret` **Added**
    - `EdgeFindingEvidence` **Added**
    - `EdgeFinding` **Added**
    - `McpGatewayToolCallRiskEvidence` **Added**
    - `McpGatewayToolCallRisk` **Added**
    - `NhiUnownedEvidence` **Added**
    - `ShadowApp` **Added**
    - `ShadowMcp.Transport` **Added**
    - `UsageResourceTarget` **Added**
* `ConductoroneApi.Finding.UpdateFindingState()`: `response.Finding` **Changed**
    - `AccessConflictEvidence` **Added**
    - `AccessConflict` **Added**
    - `ConnectorHostOutdatedEvidence` **Added**
    - `ConnectorHostOutdated` **Added**
    - `CredentialExpiring.AppResourceSecret` **Added**
    - `EdgeFindingEvidence` **Added**
    - `EdgeFinding` **Added**
    - `McpGatewayToolCallRiskEvidence` **Added**
    - `McpGatewayToolCallRisk` **Added**
    - `NhiUnownedEvidence` **Added**
    - `ShadowApp` **Added**
    - `ShadowMcp.Transport` **Added**
    - `UsageResourceTarget` **Added**
* `ConductoroneApi.FindingRoutingRule.CreateFindingRoutingRule()`: 
  * `request.Request.RoutingRule` **Changed**
    - `Dispatchers[].Notify.SlackChannel.AllowPrivateChannelFallback` **Added**
    - `FindingType.Enum(findingTypeAccessConflict)` **Added**
    - `FindingType.Enum(findingTypeConnectorHostOutdated)` **Added**
    - `FindingType.Enum(findingTypeEdge)` **Added**
    - `FindingType.Enum(findingTypeMcpGatewayToolCallRisk)` **Added**
    - `FindingType.Enum(findingTypeShadowApp)` **Added**
  * `response.RoutingRule` **Changed**
    - `Dispatchers[].Notify.SlackChannel.AllowPrivateChannelFallback` **Added**
    - `FindingType.Enum(findingTypeAccessConflict)` **Added**
    - `FindingType.Enum(findingTypeConnectorHostOutdated)` **Added**
    - `FindingType.Enum(findingTypeEdge)` **Added**
    - `FindingType.Enum(findingTypeMcpGatewayToolCallRisk)` **Added**
    - `FindingType.Enum(findingTypeShadowApp)` **Added**
* `ConductoroneApi.FindingRoutingRule.GetFindingRoutingRule()`: `response.RoutingRule` **Changed**
    - `Dispatchers[].Notify.SlackChannel.AllowPrivateChannelFallback` **Added**
    - `FindingType.Enum(findingTypeAccessConflict)` **Added**
    - `FindingType.Enum(findingTypeConnectorHostOutdated)` **Added**
    - `FindingType.Enum(findingTypeEdge)` **Added**
    - `FindingType.Enum(findingTypeMcpGatewayToolCallRisk)` **Added**
    - `FindingType.Enum(findingTypeShadowApp)` **Added**
* `ConductoroneApi.FindingRoutingRule.ListFindingRoutingRules()`: `response.List[]` **Changed**
    - `Dispatchers[].Notify.SlackChannel.AllowPrivateChannelFallback` **Added**
    - `FindingType.Enum(findingTypeAccessConflict)` **Added**
    - `FindingType.Enum(findingTypeConnectorHostOutdated)` **Added**
    - `FindingType.Enum(findingTypeEdge)` **Added**
    - `FindingType.Enum(findingTypeMcpGatewayToolCallRisk)` **Added**
    - `FindingType.Enum(findingTypeShadowApp)` **Added**
* `ConductoroneApi.FindingRoutingRule.UpdateFindingRoutingRule()`: 
  * `request.Request.UpdateFindingRoutingRuleRequest.RoutingRule` **Changed**
    - `Dispatchers[].Notify.SlackChannel.AllowPrivateChannelFallback` **Added**
    - `FindingType.Enum(findingTypeAccessConflict)` **Added**
    - `FindingType.Enum(findingTypeConnectorHostOutdated)` **Added**
    - `FindingType.Enum(findingTypeEdge)` **Added**
    - `FindingType.Enum(findingTypeMcpGatewayToolCallRisk)` **Added**
    - `FindingType.Enum(findingTypeShadowApp)` **Added**
  * `response.RoutingRule` **Changed**
    - `Dispatchers[].Notify.SlackChannel.AllowPrivateChannelFallback` **Added**
    - `FindingType.Enum(findingTypeAccessConflict)` **Added**
    - `FindingType.Enum(findingTypeConnectorHostOutdated)` **Added**
    - `FindingType.Enum(findingTypeEdge)` **Added**
    - `FindingType.Enum(findingTypeMcpGatewayToolCallRisk)` **Added**
    - `FindingType.Enum(findingTypeShadowApp)` **Added**
* `ConductoroneApi.FindingSearch.Search()`: 
  * `request.Request` **Changed**
    - `EdgeIds` **Added**
    - `ExcludeFindingTypes` **Added**
    - `FindingTypes[].Enum(findingTypeAccessConflict)` **Added**
    - `FindingTypes[].Enum(findingTypeConnectorHostOutdated)` **Added**
    - `FindingTypes[].Enum(findingTypeEdge)` **Added**
    - `FindingTypes[].Enum(findingTypeMcpGatewayToolCallRisk)` **Added**
    - `FindingTypes[].Enum(findingTypeShadowApp)` **Added**
    - `SourceDetectorIds` **Added**
    - `TargetIdentityUserIds` **Added**
  * `response.List[]` **Changed**
    - `AccessConflictEvidence` **Added**
    - `AccessConflict` **Added**
    - `ConnectorHostOutdatedEvidence` **Added**
    - `ConnectorHostOutdated` **Added**
    - `CredentialExpiring.AppResourceSecret` **Added**
    - `EdgeFindingEvidence` **Added**
    - `EdgeFinding` **Added**
    - `McpGatewayToolCallRiskEvidence` **Added**
    - `McpGatewayToolCallRisk` **Added**
    - `NhiUnownedEvidence` **Added**
    - `ShadowApp` **Added**
    - `ShadowMcp.Transport` **Added**
    - `UsageResourceTarget` **Added**
* `ConductoroneApi.FindingSettings.ListFindingSettings()`: `response.List[]` **Changed**
    - `FindingType.Enum(findingTypeAccessConflict)` **Added**
    - `FindingType.Enum(findingTypeConnectorHostOutdated)` **Added**
    - `FindingType.Enum(findingTypeEdge)` **Added**
    - `FindingType.Enum(findingTypeMcpGatewayToolCallRisk)` **Added**
    - `FindingType.Enum(findingTypeShadowApp)` **Added**
    - `HasRemediationPlaybook` **Added**
* `ConductoroneApi.FindingSettings.UpdateFindingSettings()`: 
  * `request.Request.Settings[].FindingType` **Changed**
    - `Enum(findingTypeAccessConflict)` **Added**
    - `Enum(findingTypeConnectorHostOutdated)` **Added**
    - `Enum(findingTypeEdge)` **Added**
    - `Enum(findingTypeMcpGatewayToolCallRisk)` **Added**
    - `Enum(findingTypeShadowApp)` **Added**
  * `response.List[]` **Changed**
    - `FindingType.Enum(findingTypeAccessConflict)` **Added**
    - `FindingType.Enum(findingTypeConnectorHostOutdated)` **Added**
    - `FindingType.Enum(findingTypeEdge)` **Added**
    - `FindingType.Enum(findingTypeMcpGatewayToolCallRisk)` **Added**
    - `FindingType.Enum(findingTypeShadowApp)` **Added**
    - `HasRemediationPlaybook` **Added**
* `ConductoroneApi.FindingTransformationRule.CreateFindingTransformationRule()`: 
  * `request.Request.TransformationRule` **Changed**
    - `FindingType.Enum(findingTypeAccessConflict)` **Added**
    - `FindingType.Enum(findingTypeConnectorHostOutdated)` **Added**
    - `FindingType.Enum(findingTypeEdge)` **Added**
    - `FindingType.Enum(findingTypeMcpGatewayToolCallRisk)` **Added**
    - `FindingType.Enum(findingTypeShadowApp)` **Added**
    - `Transforms[].SetAssignee` **Added**
  * `response.TransformationRule` **Changed**
    - `FindingType.Enum(findingTypeAccessConflict)` **Added**
    - `FindingType.Enum(findingTypeConnectorHostOutdated)` **Added**
    - `FindingType.Enum(findingTypeEdge)` **Added**
    - `FindingType.Enum(findingTypeMcpGatewayToolCallRisk)` **Added**
    - `FindingType.Enum(findingTypeShadowApp)` **Added**
    - `Transforms[].SetAssignee` **Added**
* `ConductoroneApi.FindingTransformationRule.GetFindingTransformationRule()`: `response.TransformationRule` **Changed**
    - `FindingType.Enum(findingTypeAccessConflict)` **Added**
    - `FindingType.Enum(findingTypeConnectorHostOutdated)` **Added**
    - `FindingType.Enum(findingTypeEdge)` **Added**
    - `FindingType.Enum(findingTypeMcpGatewayToolCallRisk)` **Added**
    - `FindingType.Enum(findingTypeShadowApp)` **Added**
    - `Transforms[].SetAssignee` **Added**
* `ConductoroneApi.FindingTransformationRule.ListFindingTransformationRules()`: `response.List[]` **Changed**
    - `FindingType.Enum(findingTypeAccessConflict)` **Added**
    - `FindingType.Enum(findingTypeConnectorHostOutdated)` **Added**
    - `FindingType.Enum(findingTypeEdge)` **Added**
    - `FindingType.Enum(findingTypeMcpGatewayToolCallRisk)` **Added**
    - `FindingType.Enum(findingTypeShadowApp)` **Added**
    - `Transforms[].SetAssignee` **Added**
* `ConductoroneApi.FindingTransformationRule.UpdateFindingTransformationRule()`: 
  * `request.Request.UpdateFindingTransformationRuleRequest.TransformationRule` **Changed**
    - `FindingType.Enum(findingTypeAccessConflict)` **Added**
    - `FindingType.Enum(findingTypeConnectorHostOutdated)` **Added**
    - `FindingType.Enum(findingTypeEdge)` **Added**
    - `FindingType.Enum(findingTypeMcpGatewayToolCallRisk)` **Added**
    - `FindingType.Enum(findingTypeShadowApp)` **Added**
    - `Transforms[].SetAssignee` **Added**
  * `response.TransformationRule` **Changed**
    - `FindingType.Enum(findingTypeAccessConflict)` **Added**
    - `FindingType.Enum(findingTypeConnectorHostOutdated)` **Added**
    - `FindingType.Enum(findingTypeEdge)` **Added**
    - `FindingType.Enum(findingTypeMcpGatewayToolCallRisk)` **Added**
    - `FindingType.Enum(findingTypeShadowApp)` **Added**
    - `Transforms[].SetAssignee` **Added**
* `ConductoroneApi.Functions.ListCommits()`: `request.Request` **Changed**
    - `PageSize` **Added**
    - `PageToken` **Added**
* `ConductoroneApi.AppUser.Merge()`: **Added**
* `ConductoroneApi.AppSecretAdmin.Revoke()`: **Added**
* `ConductoroneApi.FunctionsInvocation.Get()`: `response.Invocation` **Changed**
    - `ConversationId` **Added**
    - `NetworkCallerUserId` **Added**
    - `NetworkConfig` **Added**
    - `NetworkPolicyRevision` **Added**
* `ConductoroneApi.FunctionsInvocation.List()`: 
  * `request.Request` **Changed**
    - `PageSize` **Added**
    - `PageToken` **Added**
  * `response.List[]` **Changed**
    - `ConversationId` **Added**
    - `NetworkCallerUserId` **Added**
    - `NetworkConfig` **Added**
    - `NetworkPolicyRevision` **Added**
* `ConductoroneApi.FunctionsInvocationSearch.Search()`: `response.List[]` **Changed**
    - `ConversationId` **Added**
    - `NetworkCallerUserId` **Added**
    - `NetworkConfig` **Added**
    - `NetworkPolicyRevision` **Added**
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
* `ConductoroneApi.Hooks.Create()`: 
  *  `request.Request.BuiltinPattern.WriteAuthorization.BlockedMutatingClassifications` **Added**
  *  `response.Hook.BuiltinPattern.WriteAuthorization.BlockedMutatingClassifications` **Added**
* `ConductoroneApi.Hooks.Get()`:  `response.Hook.BuiltinPattern.WriteAuthorization.BlockedMutatingClassifications` **Added**
* `ConductoroneApi.Hooks.List()`:  `response.List[].BuiltinPattern.WriteAuthorization.BlockedMutatingClassifications` **Added**
* `ConductoroneApi.Hooks.Update()`: 
  *  `request.Request.HooksServiceUpdateRequest.Hook.BuiltinPattern.WriteAuthorization.BlockedMutatingClassifications` **Added**
  *  `response.Hook.BuiltinPattern.WriteAuthorization.BlockedMutatingClassifications` **Added**
* `ConductoroneApi.AccessReviewTemplateSetupEntitlement.SetScopeByExcludedResourceType()`: **Added**
* `ConductoroneApi.PersonalDevice.ListDeviceClients()`: `request.Request` **Changed**
    - `PageSize` **Added**
    - `PageToken` **Added**
* `ConductoroneApi.AccessReviewTemplateSetupEntitlement.SetScopeAndExcludedEntitlements()`: **Added**
* `ConductoroneApi.LocalUserInvitation.Create()`:  `response.InvitationUrl` **Added**
* `ConductoroneApi.Policies.Create()`: 
  * `request.Request.PolicySteps.Map<PolicySteps>.Steps[]` **Changed**
    - `Approval.AccessProfileOwners` **Added**
    - `Form.Form.FieldRelationships[].DependentOn.Dependencies` **Added**
    - `Form.Form.Fields[].StringField.PasswordField.Multiline` **Added**
    - `Form.Form.Fields[].StringField.RegexErrorMessage` **Added**
    - `Form.Form.Fields[].StringMapField.Secret` **Added**
    - `Form.Form.Fields[].StringSliceField.MultiSelectField` **Added**
  * `response.Policy.PolicySteps.Map<PolicySteps>.Steps[]` **Changed**
    - `Approval.AccessProfileOwners` **Added**
    - `Form.Form.FieldRelationships[].DependentOn.Dependencies` **Added**
    - `Form.Form.Fields[].StringField.PasswordField.Multiline` **Added**
    - `Form.Form.Fields[].StringField.RegexErrorMessage` **Added**
    - `Form.Form.Fields[].StringMapField.Secret` **Added**
    - `Form.Form.Fields[].StringSliceField.MultiSelectField` **Added**
* `ConductoroneApi.Policies.Get()`: `response.Policy.PolicySteps.Map<PolicySteps>.Steps[]` **Changed**
    - `Approval.AccessProfileOwners` **Added**
    - `Form.Form.FieldRelationships[].DependentOn.Dependencies` **Added**
    - `Form.Form.Fields[].StringField.PasswordField.Multiline` **Added**
    - `Form.Form.Fields[].StringField.RegexErrorMessage` **Added**
    - `Form.Form.Fields[].StringMapField.Secret` **Added**
    - `Form.Form.Fields[].StringSliceField.MultiSelectField` **Added**
* `ConductoroneApi.Policies.List()`: `response.List[].PolicySteps.Map<PolicySteps>.Steps[]` **Changed**
    - `Approval.AccessProfileOwners` **Added**
    - `Form.Form.FieldRelationships[].DependentOn.Dependencies` **Added**
    - `Form.Form.Fields[].StringField.PasswordField.Multiline` **Added**
    - `Form.Form.Fields[].StringField.RegexErrorMessage` **Added**
    - `Form.Form.Fields[].StringMapField.Secret` **Added**
    - `Form.Form.Fields[].StringSliceField.MultiSelectField` **Added**
* `ConductoroneApi.Policies.Update()`: 
  * `request.Request.UpdatePolicyRequest.Policy.PolicySteps.Map<PolicySteps>.Steps[]` **Changed**
    - `Approval.AccessProfileOwners` **Added**
    - `Form.Form.FieldRelationships[].DependentOn.Dependencies` **Added**
    - `Form.Form.Fields[].StringField.PasswordField.Multiline` **Added**
    - `Form.Form.Fields[].StringField.RegexErrorMessage` **Added**
    - `Form.Form.Fields[].StringMapField.Secret` **Added**
    - `Form.Form.Fields[].StringSliceField.MultiSelectField` **Added**
  * `response.Policy.PolicySteps.Map<PolicySteps>.Steps[]` **Changed**
    - `Approval.AccessProfileOwners` **Added**
    - `Form.Form.FieldRelationships[].DependentOn.Dependencies` **Added**
    - `Form.Form.Fields[].StringField.PasswordField.Multiline` **Added**
    - `Form.Form.Fields[].StringField.RegexErrorMessage` **Added**
    - `Form.Form.Fields[].StringMapField.Secret` **Added**
    - `Form.Form.Fields[].StringSliceField.MultiSelectField` **Added**
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
  * `request.Request` **Changed**
    - `FieldRelationships[].DependentOn.Dependencies` **Added**
    - `Fields[].StringField.PasswordField.Multiline` **Added**
    - `Fields[].StringField.RegexErrorMessage` **Added**
    - `Fields[].StringMapField.Secret` **Added**
    - `Fields[].StringSliceField.MultiSelectField` **Added**
  * `response.RequestSchema.Form` **Changed**
    - `FieldRelationships[].DependentOn.Dependencies` **Added**
    - `Fields[].StringField.PasswordField.Multiline` **Added**
    - `Fields[].StringField.RegexErrorMessage` **Added**
    - `Fields[].StringMapField.Secret` **Added**
    - `Fields[].StringSliceField.MultiSelectField` **Added**
* `ConductoroneApi.RequestSchema.Get()`: `response.RequestSchema.Form` **Changed**
    - `FieldRelationships[].DependentOn.Dependencies` **Added**
    - `Fields[].StringField.PasswordField.Multiline` **Added**
    - `Fields[].StringField.RegexErrorMessage` **Added**
    - `Fields[].StringMapField.Secret` **Added**
    - `Fields[].StringSliceField.MultiSelectField` **Added**
* `ConductoroneApi.RequestSchema.Update()`: 
  * `request.Request.RequestSchemaServiceUpdateRequest.RequestSchema.Form` **Changed**
    - `FieldRelationships[].DependentOn.Dependencies` **Added**
    - `Fields[].StringField.PasswordField.Multiline` **Added**
    - `Fields[].StringField.RegexErrorMessage` **Added**
    - `Fields[].StringMapField.Secret` **Added**
    - `Fields[].StringSliceField.MultiSelectField` **Added**
  * `response.RequestSchema.Form` **Changed**
    - `FieldRelationships[].DependentOn.Dependencies` **Added**
    - `Fields[].StringField.PasswordField.Multiline` **Added**
    - `Fields[].StringField.RegexErrorMessage` **Added**
    - `Fields[].StringMapField.Secret` **Added**
    - `Fields[].StringSliceField.MultiSelectField` **Added**
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
* `ConductoroneApi.AppSearch.Search()`:  `response.List[].EntitlementsUseResourceDescription` **Added**
* `ConductoroneApi.AutomationSearch.SearchAutomationTemplateVersions()`: `response.List[].AutomationSteps[]` **Changed**
    - `CreateRevokeTasksV2.ExclusionCriteria.ExcludedResourceTypeRefs` **Added**
    - `CreateRevokeTasksV2.InclusionCriteria.ResourceTypeRefs` **Added**
    - `SendSlackMessage.AllowPrivateChannelFallback` **Added**
* `ConductoroneApi.AutomationSearch.SearchAutomations()`: `response.List[].AutomationSteps[]` **Changed**
    - `CreateRevokeTasksV2.ExclusionCriteria.ExcludedResourceTypeRefs` **Added**
    - `CreateRevokeTasksV2.InclusionCriteria.ResourceTypeRefs` **Added**
    - `SendSlackMessage.AllowPrivateChannelFallback` **Added**
* `ConductoroneApi.FindingAudit.Search()`: 
  *  `request.Request.EventTypes[].Enum(findingAuditEventTypeHistoryImported)` **Added**
  *  `response.List[].EventType.Enum(findingAuditEventTypeHistoryImported)` **Added**
* `ConductoroneApi.AccessReviewTemplateSetupEntitlement.RemoveTemplateEntitlements()`: **Added**
* `ConductoroneApi.HooksSearch.Search()`:  `response.List[].BuiltinPattern.WriteAuthorization.BlockedMutatingClassifications` **Added**
* `ConductoroneApi.ExternalClientSearch.Search()`:  `response.List[].ClientIdType.Enum(clientIdTypeStatic)` **Added**
* `ConductoroneApi.PolicySearch.Search()`: 
  *  `request.Request.IncludeAccessProfileOwnerPolicy` **Added**
  * `response.List[].PolicySteps.Map<PolicySteps>.Steps[]` **Changed**
    - `Approval.AccessProfileOwners` **Added**
    - `Form.Form.FieldRelationships[].DependentOn.Dependencies` **Added**
    - `Form.Form.Fields[].StringField.PasswordField.Multiline` **Added**
    - `Form.Form.Fields[].StringField.RegexErrorMessage` **Added**
    - `Form.Form.Fields[].StringMapField.Secret` **Added**
    - `Form.Form.Fields[].StringSliceField.MultiSelectField` **Added**
* `ConductoroneApi.RoleMiningManagementSearch.Search()`: 
  *  `request.Request.CatalogId` **Added**
* `ConductoroneApi.PaperSecret.SearchMySecrets()`: 
  *  `request.Request.AccessRelationships` **Added**
* `ConductoroneApi.PaperSecret.SearchSecretsSharedWithMe()`: 
  *  `request.Request.SharingMode` **Added**
* `ConductoroneApi.StepUpProvider.Create()`: 
  *  `request.Request.Microsoft.RequireFreshAuthentication` **Added**
  *  `response.StepUpProvider.Microsoft.RequireFreshAuthentication` **Added**
* `ConductoroneApi.StepUpProvider.Get()`:  `response.StepUpProvider.Microsoft.RequireFreshAuthentication` **Added**
* `ConductoroneApi.AccessReviewTemplateSetupEntitlement.AddTemplateEntitlements()`: **Added**
* `ConductoroneApi.StepUpProvider.Search()`:  `response.List[].Microsoft.RequireFreshAuthentication` **Added**
* `ConductoroneApi.StepUpProvider.Update()`: 
  *  `request.Request.UpdateStepUpProviderRequest.StepUpProvider.Microsoft.RequireFreshAuthentication` **Added**
  *  `response.StepUpProvider.Microsoft.RequireFreshAuthentication` **Added**
* `ConductoroneApi.StepUpProvider.UpdateSecret()`:  `response.StepUpProvider.Microsoft.RequireFreshAuthentication` **Added**
* `ConductoroneApi.TaskSearch.Search()`: 
  * `request.Request` **Changed**
    - `AccessProfileIds` **Added**
    - `AccessProfileReviewKinds` **Added**
    - `StepApprovalTypes[].Enum(stepApprovalTypeAccessProfileOwners)` **Added**
    - `TaskTypes[].Action.AccessProfileComponent` **Added**
    - `TaskTypes[].Action.Source` **Added**
    - `TaskTypes[].Certify.ReviewSubject` **Added**
  * `response.List[].Task` **Changed**
    - `Form.FieldRelationships[].DependentOn.Dependencies` **Added**
    - `Form.Fields[].StringField.PasswordField.Multiline` **Added**
    - `Form.Fields[].StringField.RegexErrorMessage` **Added**
    - `Form.Fields[].StringMapField.Secret` **Added**
    - `Form.Fields[].StringSliceField.MultiSelectField` **Added**
    - `Policy.Current.Action.BatonResourceActionInstance.PartialResult` **Added**
    - `Policy.Current.Action.BatonResourceActionInstance.UnaddedMemberCount` **Added**
    - `Policy.Current.Approval.Approval.AccessProfileOwners` **Added**
    - `Type.Action.AccessProfileComponent` **Added**
    - `Type.Action.ActionInstance.RequestInstructions` **Added**
    - `Type.Action.PartialResult` **Added**
    - `Type.Action.ScopeRole.AppUserId` **Added**
    - `Type.Action.Source` **Added**
    - `Type.Action.Type.Enum(typeSpendRemedy)` **Added**
    - `Type.Action.UnaddedMemberCount` **Added**
    - `Type.Certify.ReviewSubject` **Added**
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
* `ConductoroneApi.AccessReviewTemplateSetupEntitlement.AddDeleteExcludedEntitlements()`: **Added**
* `ConductoroneApi.WorkloadFederation.ListTrusts()`: 
  * `request.Request` **Changed**
    - `PageSize` **Added**
    - `PageToken` **Added**
  *  `response.List[].C1EdgeService` **Added**
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
* `ConductoroneApi.AccessReviewSetupEntitlement.SetCampaignScopeByExcludedResourceType()`: **Added**
* `ConductoroneApi.Principal.ListBindings()`: 
  * `request.Request.Subject` **Changed**
    - `AuthzenServer` **Added**
    - `Edge` **Added**
    - `SsoApplication` **Added**
* `ConductoroneApi.Principal.ListCredentials()`: `request.Request` **Changed**
    - `PageSize` **Added**
    - `PageToken` **Added**
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
* `ConductoroneApi.AccessReviewSetupEntitlement.SetCampaignScopeAndExcludedEntitlements()`: **Added**
* `ConductoroneApi.AccessReviewSetupEntitlement.RemoveCampaignEntitlements()`: **Added**
* `ConductoroneApi.AccessReviewSetupEntitlement.AddDeleteExcludedEntitlements()`: **Added**
* `ConductoroneApi.AccessReviewSetupEntitlement.AddCampaignEntitlements()`: **Added**
* `ConductoroneApi.Task.CreateActionTask()`: 
  *  `request.Request.GroupOnly` **Added**
  * `response.TaskView.Task` **Changed**
    - `Form.FieldRelationships[].DependentOn.Dependencies` **Added**
    - `Form.Fields[].StringField.PasswordField.Multiline` **Added**
    - `Form.Fields[].StringField.RegexErrorMessage` **Added**
    - `Form.Fields[].StringMapField.Secret` **Added**
    - `Form.Fields[].StringSliceField.MultiSelectField` **Added**
    - `Policy.Current.Action.BatonResourceActionInstance.PartialResult` **Added**
    - `Policy.Current.Action.BatonResourceActionInstance.UnaddedMemberCount` **Added**
    - `Policy.Current.Approval.Approval.AccessProfileOwners` **Added**
    - `Type.Action.AccessProfileComponent` **Added**
    - `Type.Action.ActionInstance.RequestInstructions` **Added**
    - `Type.Action.PartialResult` **Added**
    - `Type.Action.ScopeRole.AppUserId` **Added**
    - `Type.Action.Source` **Added**
    - `Type.Action.Type.Enum(typeSpendRemedy)` **Added**
    - `Type.Action.UnaddedMemberCount` **Added**
    - `Type.Certify.ReviewSubject` **Added**
* `ConductoroneApi.Task.CreateGrantTask()`: `response.TaskView.Task` **Changed**
    - `Form.FieldRelationships[].DependentOn.Dependencies` **Added**
    - `Form.Fields[].StringField.PasswordField.Multiline` **Added**
    - `Form.Fields[].StringField.RegexErrorMessage` **Added**
    - `Form.Fields[].StringMapField.Secret` **Added**
    - `Form.Fields[].StringSliceField.MultiSelectField` **Added**
    - `Policy.Current.Action.BatonResourceActionInstance.PartialResult` **Added**
    - `Policy.Current.Action.BatonResourceActionInstance.UnaddedMemberCount` **Added**
    - `Policy.Current.Approval.Approval.AccessProfileOwners` **Added**
    - `Type.Action.AccessProfileComponent` **Added**
    - `Type.Action.ActionInstance.RequestInstructions` **Added**
    - `Type.Action.PartialResult` **Added**
    - `Type.Action.ScopeRole.AppUserId` **Added**
    - `Type.Action.Source` **Added**
    - `Type.Action.Type.Enum(typeSpendRemedy)` **Added**
    - `Type.Action.UnaddedMemberCount` **Added**
    - `Type.Certify.ReviewSubject` **Added**
* `ConductoroneApi.Task.CreateOffboardingTask()`: `response.TaskView.Task` **Changed**
    - `Form.FieldRelationships[].DependentOn.Dependencies` **Added**
    - `Form.Fields[].StringField.PasswordField.Multiline` **Added**
    - `Form.Fields[].StringField.RegexErrorMessage` **Added**
    - `Form.Fields[].StringMapField.Secret` **Added**
    - `Form.Fields[].StringSliceField.MultiSelectField` **Added**
    - `Policy.Current.Action.BatonResourceActionInstance.PartialResult` **Added**
    - `Policy.Current.Action.BatonResourceActionInstance.UnaddedMemberCount` **Added**
    - `Policy.Current.Approval.Approval.AccessProfileOwners` **Added**
    - `Type.Action.AccessProfileComponent` **Added**
    - `Type.Action.ActionInstance.RequestInstructions` **Added**
    - `Type.Action.PartialResult` **Added**
    - `Type.Action.ScopeRole.AppUserId` **Added**
    - `Type.Action.Source` **Added**
    - `Type.Action.Type.Enum(typeSpendRemedy)` **Added**
    - `Type.Action.UnaddedMemberCount` **Added**
    - `Type.Certify.ReviewSubject` **Added**
* `ConductoroneApi.Task.CreateResourceActionTask()`: `response.TaskView.Task` **Changed**
    - `Form.FieldRelationships[].DependentOn.Dependencies` **Added**
    - `Form.Fields[].StringField.PasswordField.Multiline` **Added**
    - `Form.Fields[].StringField.RegexErrorMessage` **Added**
    - `Form.Fields[].StringMapField.Secret` **Added**
    - `Form.Fields[].StringSliceField.MultiSelectField` **Added**
    - `Policy.Current.Action.BatonResourceActionInstance.PartialResult` **Added**
    - `Policy.Current.Action.BatonResourceActionInstance.UnaddedMemberCount` **Added**
    - `Policy.Current.Approval.Approval.AccessProfileOwners` **Added**
    - `Type.Action.AccessProfileComponent` **Added**
    - `Type.Action.ActionInstance.RequestInstructions` **Added**
    - `Type.Action.PartialResult` **Added**
    - `Type.Action.ScopeRole.AppUserId` **Added**
    - `Type.Action.Source` **Added**
    - `Type.Action.Type.Enum(typeSpendRemedy)` **Added**
    - `Type.Action.UnaddedMemberCount` **Added**
    - `Type.Certify.ReviewSubject` **Added**
* `ConductoroneApi.Task.CreateRevokeTask()`: `response.TaskView.Task` **Changed**
    - `Form.FieldRelationships[].DependentOn.Dependencies` **Added**
    - `Form.Fields[].StringField.PasswordField.Multiline` **Added**
    - `Form.Fields[].StringField.RegexErrorMessage` **Added**
    - `Form.Fields[].StringMapField.Secret` **Added**
    - `Form.Fields[].StringSliceField.MultiSelectField` **Added**
    - `Policy.Current.Action.BatonResourceActionInstance.PartialResult` **Added**
    - `Policy.Current.Action.BatonResourceActionInstance.UnaddedMemberCount` **Added**
    - `Policy.Current.Approval.Approval.AccessProfileOwners` **Added**
    - `Type.Action.AccessProfileComponent` **Added**
    - `Type.Action.ActionInstance.RequestInstructions` **Added**
    - `Type.Action.PartialResult` **Added**
    - `Type.Action.ScopeRole.AppUserId` **Added**
    - `Type.Action.Source` **Added**
    - `Type.Action.Type.Enum(typeSpendRemedy)` **Added**
    - `Type.Action.UnaddedMemberCount` **Added**
    - `Type.Certify.ReviewSubject` **Added**
* `ConductoroneApi.Task.Get()`: `response.TaskView.Task` **Changed**
    - `Form.FieldRelationships[].DependentOn.Dependencies` **Added**
    - `Form.Fields[].StringField.PasswordField.Multiline` **Added**
    - `Form.Fields[].StringField.RegexErrorMessage` **Added**
    - `Form.Fields[].StringMapField.Secret` **Added**
    - `Form.Fields[].StringSliceField.MultiSelectField` **Added**
    - `Policy.Current.Action.BatonResourceActionInstance.PartialResult` **Added**
    - `Policy.Current.Action.BatonResourceActionInstance.UnaddedMemberCount` **Added**
    - `Policy.Current.Approval.Approval.AccessProfileOwners` **Added**
    - `Type.Action.AccessProfileComponent` **Added**
    - `Type.Action.ActionInstance.RequestInstructions` **Added**
    - `Type.Action.PartialResult` **Added**
    - `Type.Action.ScopeRole.AppUserId` **Added**
    - `Type.Action.Source` **Added**
    - `Type.Action.Type.Enum(typeSpendRemedy)` **Added**
    - `Type.Action.UnaddedMemberCount` **Added**
    - `Type.Certify.ReviewSubject` **Added**
* `ConductoroneApi.TaskAudit.List()`: `response.List[]` **Changed**
    - `ActionInstanceCreated.Instance.BatonResourceActionInstance.PartialResult` **Added**
    - `ActionInstanceCreated.Instance.BatonResourceActionInstance.UnaddedMemberCount` **Added**
    - `ActionOutcome` **Added**
    - `ApprovalInstanceChange.Instance.Approval.AccessProfileOwners` **Added**
    - `ProvisionError.Comment` **Added**
* `ConductoroneApi.TaskActions.Approve()`: `response.TaskView.Task` **Changed**
    - `Form.FieldRelationships[].DependentOn.Dependencies` **Added**
    - `Form.Fields[].StringField.PasswordField.Multiline` **Added**
    - `Form.Fields[].StringField.RegexErrorMessage` **Added**
    - `Form.Fields[].StringMapField.Secret` **Added**
    - `Form.Fields[].StringSliceField.MultiSelectField` **Added**
    - `Policy.Current.Action.BatonResourceActionInstance.PartialResult` **Added**
    - `Policy.Current.Action.BatonResourceActionInstance.UnaddedMemberCount` **Added**
    - `Policy.Current.Approval.Approval.AccessProfileOwners` **Added**
    - `Type.Action.AccessProfileComponent` **Added**
    - `Type.Action.ActionInstance.RequestInstructions` **Added**
    - `Type.Action.PartialResult` **Added**
    - `Type.Action.ScopeRole.AppUserId` **Added**
    - `Type.Action.Source` **Added**
    - `Type.Action.Type.Enum(typeSpendRemedy)` **Added**
    - `Type.Action.UnaddedMemberCount` **Added**
    - `Type.Certify.ReviewSubject` **Added**
* `ConductoroneApi.TaskActions.ApproveWithStepUp()`: `response.TaskView.Task` **Changed**
    - `Form.FieldRelationships[].DependentOn.Dependencies` **Added**
    - `Form.Fields[].StringField.PasswordField.Multiline` **Added**
    - `Form.Fields[].StringField.RegexErrorMessage` **Added**
    - `Form.Fields[].StringMapField.Secret` **Added**
    - `Form.Fields[].StringSliceField.MultiSelectField` **Added**
    - `Policy.Current.Action.BatonResourceActionInstance.PartialResult` **Added**
    - `Policy.Current.Action.BatonResourceActionInstance.UnaddedMemberCount` **Added**
    - `Policy.Current.Approval.Approval.AccessProfileOwners` **Added**
    - `Type.Action.AccessProfileComponent` **Added**
    - `Type.Action.ActionInstance.RequestInstructions` **Added**
    - `Type.Action.PartialResult` **Added**
    - `Type.Action.ScopeRole.AppUserId` **Added**
    - `Type.Action.Source` **Added**
    - `Type.Action.Type.Enum(typeSpendRemedy)` **Added**
    - `Type.Action.UnaddedMemberCount` **Added**
    - `Type.Certify.ReviewSubject` **Added**
* `ConductoroneApi.TaskActions.Close()`: `response.TaskView.Task` **Changed**
    - `Form.FieldRelationships[].DependentOn.Dependencies` **Added**
    - `Form.Fields[].StringField.PasswordField.Multiline` **Added**
    - `Form.Fields[].StringField.RegexErrorMessage` **Added**
    - `Form.Fields[].StringMapField.Secret` **Added**
    - `Form.Fields[].StringSliceField.MultiSelectField` **Added**
    - `Policy.Current.Action.BatonResourceActionInstance.PartialResult` **Added**
    - `Policy.Current.Action.BatonResourceActionInstance.UnaddedMemberCount` **Added**
    - `Policy.Current.Approval.Approval.AccessProfileOwners` **Added**
    - `Type.Action.AccessProfileComponent` **Added**
    - `Type.Action.ActionInstance.RequestInstructions` **Added**
    - `Type.Action.PartialResult` **Added**
    - `Type.Action.ScopeRole.AppUserId` **Added**
    - `Type.Action.Source` **Added**
    - `Type.Action.Type.Enum(typeSpendRemedy)` **Added**
    - `Type.Action.UnaddedMemberCount` **Added**
    - `Type.Certify.ReviewSubject` **Added**
* `ConductoroneApi.TaskActions.Comment()`: `response.TaskView.Task` **Changed**
    - `Form.FieldRelationships[].DependentOn.Dependencies` **Added**
    - `Form.Fields[].StringField.PasswordField.Multiline` **Added**
    - `Form.Fields[].StringField.RegexErrorMessage` **Added**
    - `Form.Fields[].StringMapField.Secret` **Added**
    - `Form.Fields[].StringSliceField.MultiSelectField` **Added**
    - `Policy.Current.Action.BatonResourceActionInstance.PartialResult` **Added**
    - `Policy.Current.Action.BatonResourceActionInstance.UnaddedMemberCount` **Added**
    - `Policy.Current.Approval.Approval.AccessProfileOwners` **Added**
    - `Type.Action.AccessProfileComponent` **Added**
    - `Type.Action.ActionInstance.RequestInstructions` **Added**
    - `Type.Action.PartialResult` **Added**
    - `Type.Action.ScopeRole.AppUserId` **Added**
    - `Type.Action.Source` **Added**
    - `Type.Action.Type.Enum(typeSpendRemedy)` **Added**
    - `Type.Action.UnaddedMemberCount` **Added**
    - `Type.Certify.ReviewSubject` **Added**
* `ConductoroneApi.TaskActions.Deny()`: `response.TaskView.Task` **Changed**
    - `Form.FieldRelationships[].DependentOn.Dependencies` **Added**
    - `Form.Fields[].StringField.PasswordField.Multiline` **Added**
    - `Form.Fields[].StringField.RegexErrorMessage` **Added**
    - `Form.Fields[].StringMapField.Secret` **Added**
    - `Form.Fields[].StringSliceField.MultiSelectField` **Added**
    - `Policy.Current.Action.BatonResourceActionInstance.PartialResult` **Added**
    - `Policy.Current.Action.BatonResourceActionInstance.UnaddedMemberCount` **Added**
    - `Policy.Current.Approval.Approval.AccessProfileOwners` **Added**
    - `Type.Action.AccessProfileComponent` **Added**
    - `Type.Action.ActionInstance.RequestInstructions` **Added**
    - `Type.Action.PartialResult` **Added**
    - `Type.Action.ScopeRole.AppUserId` **Added**
    - `Type.Action.Source` **Added**
    - `Type.Action.Type.Enum(typeSpendRemedy)` **Added**
    - `Type.Action.UnaddedMemberCount` **Added**
    - `Type.Certify.ReviewSubject` **Added**
* `ConductoroneApi.TaskActions.EscalateToEmergencyAccess()`: `response.TaskView.Task` **Changed**
    - `Form.FieldRelationships[].DependentOn.Dependencies` **Added**
    - `Form.Fields[].StringField.PasswordField.Multiline` **Added**
    - `Form.Fields[].StringField.RegexErrorMessage` **Added**
    - `Form.Fields[].StringMapField.Secret` **Added**
    - `Form.Fields[].StringSliceField.MultiSelectField` **Added**
    - `Policy.Current.Action.BatonResourceActionInstance.PartialResult` **Added**
    - `Policy.Current.Action.BatonResourceActionInstance.UnaddedMemberCount` **Added**
    - `Policy.Current.Approval.Approval.AccessProfileOwners` **Added**
    - `Type.Action.AccessProfileComponent` **Added**
    - `Type.Action.ActionInstance.RequestInstructions` **Added**
    - `Type.Action.PartialResult` **Added**
    - `Type.Action.ScopeRole.AppUserId` **Added**
    - `Type.Action.Source` **Added**
    - `Type.Action.Type.Enum(typeSpendRemedy)` **Added**
    - `Type.Action.UnaddedMemberCount` **Added**
    - `Type.Certify.ReviewSubject` **Added**
* `ConductoroneApi.TaskActions.HardReset()`: `response.TaskView.Task` **Changed**
    - `Form.FieldRelationships[].DependentOn.Dependencies` **Added**
    - `Form.Fields[].StringField.PasswordField.Multiline` **Added**
    - `Form.Fields[].StringField.RegexErrorMessage` **Added**
    - `Form.Fields[].StringMapField.Secret` **Added**
    - `Form.Fields[].StringSliceField.MultiSelectField` **Added**
    - `Policy.Current.Action.BatonResourceActionInstance.PartialResult` **Added**
    - `Policy.Current.Action.BatonResourceActionInstance.UnaddedMemberCount` **Added**
    - `Policy.Current.Approval.Approval.AccessProfileOwners` **Added**
    - `Type.Action.AccessProfileComponent` **Added**
    - `Type.Action.ActionInstance.RequestInstructions` **Added**
    - `Type.Action.PartialResult` **Added**
    - `Type.Action.ScopeRole.AppUserId` **Added**
    - `Type.Action.Source` **Added**
    - `Type.Action.Type.Enum(typeSpendRemedy)` **Added**
    - `Type.Action.UnaddedMemberCount` **Added**
    - `Type.Certify.ReviewSubject` **Added**
* `ConductoroneApi.TaskActions.ProcessNow()`: `response.TaskView.Task` **Changed**
    - `Form.FieldRelationships[].DependentOn.Dependencies` **Added**
    - `Form.Fields[].StringField.PasswordField.Multiline` **Added**
    - `Form.Fields[].StringField.RegexErrorMessage` **Added**
    - `Form.Fields[].StringMapField.Secret` **Added**
    - `Form.Fields[].StringSliceField.MultiSelectField` **Added**
    - `Policy.Current.Action.BatonResourceActionInstance.PartialResult` **Added**
    - `Policy.Current.Action.BatonResourceActionInstance.UnaddedMemberCount` **Added**
    - `Policy.Current.Approval.Approval.AccessProfileOwners` **Added**
    - `Type.Action.AccessProfileComponent` **Added**
    - `Type.Action.ActionInstance.RequestInstructions` **Added**
    - `Type.Action.PartialResult` **Added**
    - `Type.Action.ScopeRole.AppUserId` **Added**
    - `Type.Action.Source` **Added**
    - `Type.Action.Type.Enum(typeSpendRemedy)` **Added**
    - `Type.Action.UnaddedMemberCount` **Added**
    - `Type.Certify.ReviewSubject` **Added**
* `ConductoroneApi.TaskActions.Reassign()`: `response.TaskView.Task` **Changed**
    - `Form.FieldRelationships[].DependentOn.Dependencies` **Added**
    - `Form.Fields[].StringField.PasswordField.Multiline` **Added**
    - `Form.Fields[].StringField.RegexErrorMessage` **Added**
    - `Form.Fields[].StringMapField.Secret` **Added**
    - `Form.Fields[].StringSliceField.MultiSelectField` **Added**
    - `Policy.Current.Action.BatonResourceActionInstance.PartialResult` **Added**
    - `Policy.Current.Action.BatonResourceActionInstance.UnaddedMemberCount` **Added**
    - `Policy.Current.Approval.Approval.AccessProfileOwners` **Added**
    - `Type.Action.AccessProfileComponent` **Added**
    - `Type.Action.ActionInstance.RequestInstructions` **Added**
    - `Type.Action.PartialResult` **Added**
    - `Type.Action.ScopeRole.AppUserId` **Added**
    - `Type.Action.Source` **Added**
    - `Type.Action.Type.Enum(typeSpendRemedy)` **Added**
    - `Type.Action.UnaddedMemberCount` **Added**
    - `Type.Certify.ReviewSubject` **Added**
* `ConductoroneApi.TaskActions.Restart()`: `response.TaskView.Task` **Changed**
    - `Form.FieldRelationships[].DependentOn.Dependencies` **Added**
    - `Form.Fields[].StringField.PasswordField.Multiline` **Added**
    - `Form.Fields[].StringField.RegexErrorMessage` **Added**
    - `Form.Fields[].StringMapField.Secret` **Added**
    - `Form.Fields[].StringSliceField.MultiSelectField` **Added**
    - `Policy.Current.Action.BatonResourceActionInstance.PartialResult` **Added**
    - `Policy.Current.Action.BatonResourceActionInstance.UnaddedMemberCount` **Added**
    - `Policy.Current.Approval.Approval.AccessProfileOwners` **Added**
    - `Type.Action.AccessProfileComponent` **Added**
    - `Type.Action.ActionInstance.RequestInstructions` **Added**
    - `Type.Action.PartialResult` **Added**
    - `Type.Action.ScopeRole.AppUserId` **Added**
    - `Type.Action.Source` **Added**
    - `Type.Action.Type.Enum(typeSpendRemedy)` **Added**
    - `Type.Action.UnaddedMemberCount` **Added**
    - `Type.Certify.ReviewSubject` **Added**
* `ConductoroneApi.TaskActions.RetryProvisioning()`: `response.TaskView.Task` **Changed**
    - `Form.FieldRelationships[].DependentOn.Dependencies` **Added**
    - `Form.Fields[].StringField.PasswordField.Multiline` **Added**
    - `Form.Fields[].StringField.RegexErrorMessage` **Added**
    - `Form.Fields[].StringMapField.Secret` **Added**
    - `Form.Fields[].StringSliceField.MultiSelectField` **Added**
    - `Policy.Current.Action.BatonResourceActionInstance.PartialResult` **Added**
    - `Policy.Current.Action.BatonResourceActionInstance.UnaddedMemberCount` **Added**
    - `Policy.Current.Approval.Approval.AccessProfileOwners` **Added**
    - `Type.Action.AccessProfileComponent` **Added**
    - `Type.Action.ActionInstance.RequestInstructions` **Added**
    - `Type.Action.PartialResult` **Added**
    - `Type.Action.ScopeRole.AppUserId` **Added**
    - `Type.Action.Source` **Added**
    - `Type.Action.Type.Enum(typeSpendRemedy)` **Added**
    - `Type.Action.UnaddedMemberCount` **Added**
    - `Type.Certify.ReviewSubject` **Added**
* `ConductoroneApi.TaskActions.SkipStep()`: `response.TaskView.Task` **Changed**
    - `Form.FieldRelationships[].DependentOn.Dependencies` **Added**
    - `Form.Fields[].StringField.PasswordField.Multiline` **Added**
    - `Form.Fields[].StringField.RegexErrorMessage` **Added**
    - `Form.Fields[].StringMapField.Secret` **Added**
    - `Form.Fields[].StringSliceField.MultiSelectField` **Added**
    - `Policy.Current.Action.BatonResourceActionInstance.PartialResult` **Added**
    - `Policy.Current.Action.BatonResourceActionInstance.UnaddedMemberCount` **Added**
    - `Policy.Current.Approval.Approval.AccessProfileOwners` **Added**
    - `Type.Action.AccessProfileComponent` **Added**
    - `Type.Action.ActionInstance.RequestInstructions` **Added**
    - `Type.Action.PartialResult` **Added**
    - `Type.Action.ScopeRole.AppUserId` **Added**
    - `Type.Action.Source` **Added**
    - `Type.Action.Type.Enum(typeSpendRemedy)` **Added**
    - `Type.Action.UnaddedMemberCount` **Added**
    - `Type.Certify.ReviewSubject` **Added**
* `ConductoroneApi.TaskActions.UpdateGrantDuration()`: `response.TaskView.Task` **Changed**
    - `Form.FieldRelationships[].DependentOn.Dependencies` **Added**
    - `Form.Fields[].StringField.PasswordField.Multiline` **Added**
    - `Form.Fields[].StringField.RegexErrorMessage` **Added**
    - `Form.Fields[].StringMapField.Secret` **Added**
    - `Form.Fields[].StringSliceField.MultiSelectField` **Added**
    - `Policy.Current.Action.BatonResourceActionInstance.PartialResult` **Added**
    - `Policy.Current.Action.BatonResourceActionInstance.UnaddedMemberCount` **Added**
    - `Policy.Current.Approval.Approval.AccessProfileOwners` **Added**
    - `Type.Action.AccessProfileComponent` **Added**
    - `Type.Action.ActionInstance.RequestInstructions` **Added**
    - `Type.Action.PartialResult` **Added**
    - `Type.Action.ScopeRole.AppUserId` **Added**
    - `Type.Action.Source` **Added**
    - `Type.Action.Type.Enum(typeSpendRemedy)` **Added**
    - `Type.Action.UnaddedMemberCount` **Added**
    - `Type.Certify.ReviewSubject` **Added**
* `ConductoroneApi.TaskActions.UpdateRequestData()`: `response.TaskView.Task` **Changed**
    - `Form.FieldRelationships[].DependentOn.Dependencies` **Added**
    - `Form.Fields[].StringField.PasswordField.Multiline` **Added**
    - `Form.Fields[].StringField.RegexErrorMessage` **Added**
    - `Form.Fields[].StringMapField.Secret` **Added**
    - `Form.Fields[].StringSliceField.MultiSelectField` **Added**
    - `Policy.Current.Action.BatonResourceActionInstance.PartialResult` **Added**
    - `Policy.Current.Action.BatonResourceActionInstance.UnaddedMemberCount` **Added**
    - `Policy.Current.Approval.Approval.AccessProfileOwners` **Added**
    - `Type.Action.AccessProfileComponent` **Added**
    - `Type.Action.ActionInstance.RequestInstructions` **Added**
    - `Type.Action.PartialResult` **Added**
    - `Type.Action.ScopeRole.AppUserId` **Added**
    - `Type.Action.Source` **Added**
    - `Type.Action.Type.Enum(typeSpendRemedy)` **Added**
    - `Type.Action.UnaddedMemberCount` **Added**
    - `Type.Certify.ReviewSubject` **Added**
* `ConductoroneApi.Webhooks.Test()`: `response.Webhook` **Changed**
    - `State.Enum(webhookStateCanceled)` **Added**
    - `SupersededReason` **Added**
