## Go SDK Changes:
* `ConductoroneApi.AuthzenServer.Create()`: **Added**
* `ConductoroneApi.AuthzenServer.Delete()`: **Added**
* `ConductoroneApi.AuthzenServer.Get()`: **Added**
* `ConductoroneApi.AuthzenServer.List()`: **Added**
* `ConductoroneApi.AuthzenServer.Update()`: **Added**
* `ConductoroneApi.Edge.Create()`: **Added**
* `ConductoroneApi.Edge.Delete()`: **Added**
* `ConductoroneApi.Edge.Get()`: **Added**
* `ConductoroneApi.Edge.List()`: **Added**
* `ConductoroneApi.Edge.Update()`: **Added**
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
* `ConductoroneApi.ShadowMcpOccurrence.Search()`: **Added**
* `ConductoroneApi.SpendInsights.SearchMySpendUsage()`: **Added**
* `ConductoroneApi.Task.CreateSpendRemedyTask()`: **Added**
* `ConductoroneApi.Task.GetSpendRemedyReview()`: **Added**
* `ConductoroneApi.Task.GetSpendRemedyStatus()`: **Added**
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
* `ConductoroneApi.AppEntitlementSearch.Search()`: 
  *  `request.Request.AppResourceTypeRefs` **Added**
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
* `ConductoroneApi.Finding.BulkCreateFindingTasks()`: 
  *  `request.Request.SearchRequest.FindingTypes[].Enum(findingTypeShadowApp)` **Added**
* `ConductoroneApi.Finding.BulkUpdateFindingState()`: 
  *  `request.Request.SearchRequest.FindingTypes[].Enum(findingTypeShadowApp)` **Added**
* `ConductoroneApi.Finding.CreateFinding()`: `response.Finding` **Changed**
    - `ShadowApp` **Added**
    - `UsageResourceTarget` **Added**
* `ConductoroneApi.Finding.CreateFindingTask()`: `response.Finding` **Changed**
    - `ShadowApp` **Added**
    - `UsageResourceTarget` **Added**
* `ConductoroneApi.Finding.GetFinding()`: `response.Finding` **Changed**
    - `ShadowApp` **Added**
    - `UsageResourceTarget` **Added**
* `ConductoroneApi.Finding.UpdateFindingAssignee()`: `response.Finding` **Changed**
    - `ShadowApp` **Added**
    - `UsageResourceTarget` **Added**
* `ConductoroneApi.Finding.UpdateFindingState()`: `response.Finding` **Changed**
    - `ShadowApp` **Added**
    - `UsageResourceTarget` **Added**
* `ConductoroneApi.FindingRoutingRule.CreateFindingRoutingRule()`: 
  *  `request.Request.RoutingRule.FindingType.Enum(findingTypeShadowApp)` **Added**
  *  `response.RoutingRule.FindingType.Enum(findingTypeShadowApp)` **Added**
* `ConductoroneApi.FindingRoutingRule.GetFindingRoutingRule()`:  `response.RoutingRule.FindingType.Enum(findingTypeShadowApp)` **Added**
* `ConductoroneApi.FindingRoutingRule.ListFindingRoutingRules()`:  `response.List[].FindingType.Enum(findingTypeShadowApp)` **Added**
* `ConductoroneApi.FindingRoutingRule.UpdateFindingRoutingRule()`: 
  *  `request.Request.UpdateFindingRoutingRuleRequest.RoutingRule.FindingType.Enum(findingTypeShadowApp)` **Added**
  *  `response.RoutingRule.FindingType.Enum(findingTypeShadowApp)` **Added**
* `ConductoroneApi.FindingSearch.Search()`: 
  *  `request.Request.FindingTypes[].Enum(findingTypeShadowApp)` **Added**
  * `response.List[]` **Changed**
    - `ShadowApp` **Added**
    - `UsageResourceTarget` **Added**
* `ConductoroneApi.FindingSettings.ListFindingSettings()`:  `response.List[].FindingType.Enum(findingTypeShadowApp)` **Added**
* `ConductoroneApi.FindingSettings.UpdateFindingSettings()`: 
  *  `request.Request.Settings[].FindingType.Enum(findingTypeShadowApp)` **Added**
  *  `response.List[].FindingType.Enum(findingTypeShadowApp)` **Added**
* `ConductoroneApi.FindingTransformationRule.CreateFindingTransformationRule()`: 
  *  `request.Request.TransformationRule.FindingType.Enum(findingTypeShadowApp)` **Added**
  *  `response.TransformationRule.FindingType.Enum(findingTypeShadowApp)` **Added**
* `ConductoroneApi.FindingTransformationRule.GetFindingTransformationRule()`:  `response.TransformationRule.FindingType.Enum(findingTypeShadowApp)` **Added**
* `ConductoroneApi.FindingTransformationRule.ListFindingTransformationRules()`:  `response.List[].FindingType.Enum(findingTypeShadowApp)` **Added**
* `ConductoroneApi.FindingTransformationRule.UpdateFindingTransformationRule()`: 
  *  `request.Request.UpdateFindingTransformationRuleRequest.TransformationRule.FindingType.Enum(findingTypeShadowApp)` **Added**
  *  `response.TransformationRule.FindingType.Enum(findingTypeShadowApp)` **Added**
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
* `ConductoroneApi.AutomationSearch.SearchAutomationTemplateVersions()`: `response.List[].AutomationSteps[].CreateRevokeTasksV2` **Changed**
    - `ExclusionCriteria.ExcludedResourceTypeRefs` **Added**
    - `InclusionCriteria.ResourceTypeRefs` **Added**
* `ConductoroneApi.AutomationSearch.SearchAutomations()`: `response.List[].AutomationSteps[].CreateRevokeTasksV2` **Changed**
    - `ExclusionCriteria.ExcludedResourceTypeRefs` **Added**
    - `InclusionCriteria.ResourceTypeRefs` **Added**
* `ConductoroneApi.PaperSecret.SearchSecretsSharedWithMe()`: 
  *  `request.Request.SharingMode` **Added**
* `ConductoroneApi.TaskSearch.Search()`:  `response.List[].Task.Type.Action.Type.Enum(typeSpendRemedy)` **Added**
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
* `ConductoroneApi.Task.CreateActionTask()`:  `response.TaskView.Task.Type.Action.Type.Enum(typeSpendRemedy)` **Added**
* `ConductoroneApi.Task.CreateGrantTask()`:  `response.TaskView.Task.Type.Action.Type.Enum(typeSpendRemedy)` **Added**
* `ConductoroneApi.Task.CreateOffboardingTask()`:  `response.TaskView.Task.Type.Action.Type.Enum(typeSpendRemedy)` **Added**
* `ConductoroneApi.Task.CreateResourceActionTask()`:  `response.TaskView.Task.Type.Action.Type.Enum(typeSpendRemedy)` **Added**
* `ConductoroneApi.Task.CreateRevokeTask()`:  `response.TaskView.Task.Type.Action.Type.Enum(typeSpendRemedy)` **Added**
* `ConductoroneApi.Task.Get()`:  `response.TaskView.Task.Type.Action.Type.Enum(typeSpendRemedy)` **Added**
* `ConductoroneApi.TaskActions.Approve()`:  `response.TaskView.Task.Type.Action.Type.Enum(typeSpendRemedy)` **Added**
* `ConductoroneApi.TaskActions.ApproveWithStepUp()`:  `response.TaskView.Task.Type.Action.Type.Enum(typeSpendRemedy)` **Added**
* `ConductoroneApi.TaskActions.Close()`:  `response.TaskView.Task.Type.Action.Type.Enum(typeSpendRemedy)` **Added**
* `ConductoroneApi.TaskActions.Comment()`:  `response.TaskView.Task.Type.Action.Type.Enum(typeSpendRemedy)` **Added**
* `ConductoroneApi.TaskActions.Deny()`:  `response.TaskView.Task.Type.Action.Type.Enum(typeSpendRemedy)` **Added**
* `ConductoroneApi.TaskActions.EscalateToEmergencyAccess()`:  `response.TaskView.Task.Type.Action.Type.Enum(typeSpendRemedy)` **Added**
* `ConductoroneApi.TaskActions.HardReset()`:  `response.TaskView.Task.Type.Action.Type.Enum(typeSpendRemedy)` **Added**
* `ConductoroneApi.TaskActions.ProcessNow()`:  `response.TaskView.Task.Type.Action.Type.Enum(typeSpendRemedy)` **Added**
* `ConductoroneApi.TaskActions.Reassign()`:  `response.TaskView.Task.Type.Action.Type.Enum(typeSpendRemedy)` **Added**
* `ConductoroneApi.TaskActions.Restart()`:  `response.TaskView.Task.Type.Action.Type.Enum(typeSpendRemedy)` **Added**
* `ConductoroneApi.TaskActions.RetryProvisioning()`:  `response.TaskView.Task.Type.Action.Type.Enum(typeSpendRemedy)` **Added**
* `ConductoroneApi.TaskActions.SkipStep()`:  `response.TaskView.Task.Type.Action.Type.Enum(typeSpendRemedy)` **Added**
* `ConductoroneApi.TaskActions.UpdateGrantDuration()`:  `response.TaskView.Task.Type.Action.Type.Enum(typeSpendRemedy)` **Added**
* `ConductoroneApi.TaskActions.UpdateRequestData()`:  `response.TaskView.Task.Type.Action.Type.Enum(typeSpendRemedy)` **Added**
