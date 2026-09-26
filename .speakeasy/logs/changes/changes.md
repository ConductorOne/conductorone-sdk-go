## Go SDK Changes:
* `ConductoroneApi.SpendInsights.GetMySpendStatus()`: `response.Block` **Changed** (Breaking ⚠️)
    - `AttemptedAppId` **Added**
    - `PeriodEnd` **Added**
    - `PersonalRequestSizing` **Added**
    - `RequestTaskId` **Added**
    - `ScopeAppId` **Removed** (Breaking ⚠️)
    - `ScopeKind` **Removed** (Breaking ⚠️)
    - `Scope` **Added**
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
* `ConductoroneApi.Edge.Create()`: **Added**
* `ConductoroneApi.Edge.Delete()`: **Added**
* `ConductoroneApi.Edge.Get()`: **Added**
* `ConductoroneApi.Edge.GetEgressTrafficSummary()`: **Added**
* `ConductoroneApi.Edge.GetEgressTrafficSummaryRollup()`: **Added**
* `ConductoroneApi.Edge.GetInferenceUsageAttribution()`: **Added**
* `ConductoroneApi.Edge.GetInferenceUsageAttributionRollup()`: **Added**
* `ConductoroneApi.Edge.GetInferenceUsageSummary()`: **Added**
* `ConductoroneApi.Edge.GetInferenceUsageSummaryRollup()`: **Added**
* `ConductoroneApi.Edge.List()`: **Added**
* `ConductoroneApi.Edge.ListEgressTopAgents()`: **Added**
* `ConductoroneApi.Edge.ListEgressTopAgentsRollup()`: **Added**
* `ConductoroneApi.Edge.ListEgressTopDestinations()`: **Added**
* `ConductoroneApi.Edge.ListEgressTopToolsRollup()`: **Added**
* `ConductoroneApi.Edge.ListEgressTrafficEvents()`: **Added**
* `ConductoroneApi.Edge.ListInferenceRoutes()`: **Added**
* `ConductoroneApi.Edge.ListInferenceTopAgents()`: **Added**
* `ConductoroneApi.Edge.ListInferenceTopTools()`: **Added**
* `ConductoroneApi.Edge.ListInferenceTrafficEvents()`: **Added**
* `ConductoroneApi.Edge.Update()`: **Added**
* `ConductoroneApi.TbControlPlane.GetTrafficSummary()`: **Added**
* `ConductoroneApi.TbControlPlane.GetUsageAttribution()`: **Added**
* `ConductoroneApi.TbControlPlane.GetUsageSummary()`: **Added**
* `ConductoroneApi.TbControlPlane.ListTopAgents()`: **Added**
* `ConductoroneApi.TbControlPlane.ListTopTools()`: **Added**
* `ConductoroneApi.TbControlPlane.ListTrafficEvents()`: **Added**
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
* `ConductoroneApi.RoleMiningManagement.CountEntitlementSelectionFilterAttributeValues()`: **Added**
* `ConductoroneApi.RoleMiningManagement.ListCohortFilterAttributeValues()`: **Added**
* `ConductoroneApi.RoleMiningManagement.ListCohortFilterAttributes()`: **Added**
* `ConductoroneApi.VirtualMcpServerMySearch.Search()`: **Added**
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
* `ConductoroneApi.A2Ui.ListSurfaces()`:  `response.Surfaces[].SavedReportId` **Added**
* `ConductoroneApi.AccessReview.Create()`: `response.AccessReview.AccessReview.ColumnConfig` **Changed**
    - `Columns[].Enum(accessReviewTaskColumnResourceDescription)` **Added**
    - `OrderedColumns[].Builtin.Enum(accessReviewTaskColumnResourceDescription)` **Added**
* `ConductoroneApi.AccessReview.Get()`: `response.AccessReview.AccessReview.ColumnConfig` **Changed**
    - `Columns[].Enum(accessReviewTaskColumnResourceDescription)` **Added**
    - `OrderedColumns[].Builtin.Enum(accessReviewTaskColumnResourceDescription)` **Added**
* `ConductoroneApi.AccessReview.List()`: `response.List[].AccessReview.ColumnConfig` **Changed**
    - `Columns[].Enum(accessReviewTaskColumnResourceDescription)` **Added**
    - `OrderedColumns[].Builtin.Enum(accessReviewTaskColumnResourceDescription)` **Added**
* `ConductoroneApi.AccessReview.Update()`: 
  * `request.Request.AccessReviewServiceUpdateRequest.AccessReview.ColumnConfig` **Changed**
    - `Columns[].Enum(accessReviewTaskColumnResourceDescription)` **Added**
    - `OrderedColumns[].Builtin.Enum(accessReviewTaskColumnResourceDescription)` **Added**
  * `response.AccessReview.AccessReview.ColumnConfig` **Changed**
    - `Columns[].Enum(accessReviewTaskColumnResourceDescription)` **Added**
    - `OrderedColumns[].Builtin.Enum(accessReviewTaskColumnResourceDescription)` **Added**
* `ConductoroneApi.AccessReviewTemplate.Create()`: 
  * `request.Request.ColumnConfig` **Changed**
    - `Columns[].Enum(accessReviewTaskColumnResourceDescription)` **Added**
    - `OrderedColumns[].Builtin.Enum(accessReviewTaskColumnResourceDescription)` **Added**
  * `response.AccessReviewTemplate.ColumnConfig` **Changed**
    - `Columns[].Enum(accessReviewTaskColumnResourceDescription)` **Added**
    - `OrderedColumns[].Builtin.Enum(accessReviewTaskColumnResourceDescription)` **Added**
* `ConductoroneApi.AccessReviewTemplate.Get()`: `response.AccessReviewTemplate.ColumnConfig` **Changed**
    - `Columns[].Enum(accessReviewTaskColumnResourceDescription)` **Added**
    - `OrderedColumns[].Builtin.Enum(accessReviewTaskColumnResourceDescription)` **Added**
* `ConductoroneApi.AccessReviewTemplate.Update()`: 
  * `request.Request.AccessReviewTemplateServiceUpdateRequest.AccessReviewTemplate.ColumnConfig` **Changed**
    - `Columns[].Enum(accessReviewTaskColumnResourceDescription)` **Added**
    - `OrderedColumns[].Builtin.Enum(accessReviewTaskColumnResourceDescription)` **Added**
  * `response.AccessReviewTemplate.ColumnConfig` **Changed**
    - `Columns[].Enum(accessReviewTaskColumnResourceDescription)` **Added**
    - `OrderedColumns[].Builtin.Enum(accessReviewTaskColumnResourceDescription)` **Added**
* `ConductoroneApi.Connector.Update()`: 
  *  `request.Request.ConnectorServiceUpdateRequest.Justification` **Added**
* `ConductoroneApi.Connector.UpdateConnectorSchedule()`: 
  *  `request.Request.UpdateConnectorScheduleRequest.Justification` **Added**
* `ConductoroneApi.AppEntitlementSearch.Search()`: `request.Request` **Changed**
    - `AppResourceTypeRefs` **Added**
    - `OwnerUserIds` **Added**
* `ConductoroneApi.McpServer.Register()`: 
  *  `request.Request.McpServerServiceRegisterRequest.AccessProvisioning` **Added**
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
  * `request.Request.SearchRequest.FindingTypes[]` **Changed**
    - `Enum(findingTypeMcpGatewayToolCallRisk)` **Added**
    - `Enum(findingTypeShadowApp)` **Added**
* `ConductoroneApi.Finding.BulkUpdateFindingState()`: 
  * `request.Request.SearchRequest.FindingTypes[]` **Changed**
    - `Enum(findingTypeMcpGatewayToolCallRisk)` **Added**
    - `Enum(findingTypeShadowApp)` **Added**
* `ConductoroneApi.Finding.CreateFinding()`: `response.Finding` **Changed**
    - `McpGatewayToolCallRiskEvidence` **Added**
    - `McpGatewayToolCallRisk` **Added**
    - `ShadowApp` **Added**
    - `ShadowMcp.Transport` **Added**
    - `UsageResourceTarget` **Added**
* `ConductoroneApi.Finding.CreateFindingTask()`: `response.Finding` **Changed**
    - `McpGatewayToolCallRiskEvidence` **Added**
    - `McpGatewayToolCallRisk` **Added**
    - `ShadowApp` **Added**
    - `ShadowMcp.Transport` **Added**
    - `UsageResourceTarget` **Added**
* `ConductoroneApi.Finding.GetFinding()`: `response.Finding` **Changed**
    - `McpGatewayToolCallRiskEvidence` **Added**
    - `McpGatewayToolCallRisk` **Added**
    - `ShadowApp` **Added**
    - `ShadowMcp.Transport` **Added**
    - `UsageResourceTarget` **Added**
* `ConductoroneApi.Finding.UpdateFindingAssignee()`: `response.Finding` **Changed**
    - `McpGatewayToolCallRiskEvidence` **Added**
    - `McpGatewayToolCallRisk` **Added**
    - `ShadowApp` **Added**
    - `ShadowMcp.Transport` **Added**
    - `UsageResourceTarget` **Added**
* `ConductoroneApi.Finding.UpdateFindingState()`: `response.Finding` **Changed**
    - `McpGatewayToolCallRiskEvidence` **Added**
    - `McpGatewayToolCallRisk` **Added**
    - `ShadowApp` **Added**
    - `ShadowMcp.Transport` **Added**
    - `UsageResourceTarget` **Added**
* `ConductoroneApi.FindingRoutingRule.CreateFindingRoutingRule()`: 
  * `request.Request.RoutingRule.FindingType` **Changed**
    - `Enum(findingTypeMcpGatewayToolCallRisk)` **Added**
    - `Enum(findingTypeShadowApp)` **Added**
  * `response.RoutingRule.FindingType` **Changed**
    - `Enum(findingTypeMcpGatewayToolCallRisk)` **Added**
    - `Enum(findingTypeShadowApp)` **Added**
* `ConductoroneApi.FindingRoutingRule.GetFindingRoutingRule()`: `response.RoutingRule.FindingType` **Changed**
    - `Enum(findingTypeMcpGatewayToolCallRisk)` **Added**
    - `Enum(findingTypeShadowApp)` **Added**
* `ConductoroneApi.FindingRoutingRule.ListFindingRoutingRules()`: `response.List[].FindingType` **Changed**
    - `Enum(findingTypeMcpGatewayToolCallRisk)` **Added**
    - `Enum(findingTypeShadowApp)` **Added**
* `ConductoroneApi.FindingRoutingRule.UpdateFindingRoutingRule()`: 
  * `request.Request.UpdateFindingRoutingRuleRequest.RoutingRule.FindingType` **Changed**
    - `Enum(findingTypeMcpGatewayToolCallRisk)` **Added**
    - `Enum(findingTypeShadowApp)` **Added**
  * `response.RoutingRule.FindingType` **Changed**
    - `Enum(findingTypeMcpGatewayToolCallRisk)` **Added**
    - `Enum(findingTypeShadowApp)` **Added**
* `ConductoroneApi.FindingSearch.Search()`: 
  * `request.Request.FindingTypes[]` **Changed**
    - `Enum(findingTypeMcpGatewayToolCallRisk)` **Added**
    - `Enum(findingTypeShadowApp)` **Added**
  * `response.List[]` **Changed**
    - `McpGatewayToolCallRiskEvidence` **Added**
    - `McpGatewayToolCallRisk` **Added**
    - `ShadowApp` **Added**
    - `ShadowMcp.Transport` **Added**
    - `UsageResourceTarget` **Added**
* `ConductoroneApi.FindingSettings.ListFindingSettings()`: `response.List[].FindingType` **Changed**
    - `Enum(findingTypeMcpGatewayToolCallRisk)` **Added**
    - `Enum(findingTypeShadowApp)` **Added**
* `ConductoroneApi.FindingSettings.UpdateFindingSettings()`: 
  * `request.Request.Settings[].FindingType` **Changed**
    - `Enum(findingTypeMcpGatewayToolCallRisk)` **Added**
    - `Enum(findingTypeShadowApp)` **Added**
  * `response.List[].FindingType` **Changed**
    - `Enum(findingTypeMcpGatewayToolCallRisk)` **Added**
    - `Enum(findingTypeShadowApp)` **Added**
* `ConductoroneApi.FindingTransformationRule.CreateFindingTransformationRule()`: 
  * `request.Request.TransformationRule` **Changed**
    - `FindingType.Enum(findingTypeMcpGatewayToolCallRisk)` **Added**
    - `FindingType.Enum(findingTypeShadowApp)` **Added**
    - `Transforms[].SetAssignee` **Added**
  * `response.TransformationRule` **Changed**
    - `FindingType.Enum(findingTypeMcpGatewayToolCallRisk)` **Added**
    - `FindingType.Enum(findingTypeShadowApp)` **Added**
    - `Transforms[].SetAssignee` **Added**
* `ConductoroneApi.FindingTransformationRule.GetFindingTransformationRule()`: `response.TransformationRule` **Changed**
    - `FindingType.Enum(findingTypeMcpGatewayToolCallRisk)` **Added**
    - `FindingType.Enum(findingTypeShadowApp)` **Added**
    - `Transforms[].SetAssignee` **Added**
* `ConductoroneApi.FindingTransformationRule.ListFindingTransformationRules()`: `response.List[]` **Changed**
    - `FindingType.Enum(findingTypeMcpGatewayToolCallRisk)` **Added**
    - `FindingType.Enum(findingTypeShadowApp)` **Added**
    - `Transforms[].SetAssignee` **Added**
* `ConductoroneApi.FindingTransformationRule.UpdateFindingTransformationRule()`: 
  * `request.Request.UpdateFindingTransformationRuleRequest.TransformationRule` **Changed**
    - `FindingType.Enum(findingTypeMcpGatewayToolCallRisk)` **Added**
    - `FindingType.Enum(findingTypeShadowApp)` **Added**
    - `Transforms[].SetAssignee` **Added**
  * `response.TransformationRule` **Changed**
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
    - `LatestRun.SurfaceSnapshot.SavedReportId` **Added**
    - `Report.Description` **Added**
    - `Report.LatestRunSummary` **Added**
* `ConductoroneApi.Reporting.List()`: `response.List[]` **Changed**
    - `Description` **Added**
    - `LatestRunSummary` **Added**
* `ConductoroneApi.Reporting.Run()`:  `response.Run.SurfaceSnapshot.SavedReportId` **Added**
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
* `ConductoroneApi.SpendInsights.SearchDenials()`: 
  *  `request.Request.Filters.Departments` **Added**
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
