## Go SDK Changes:
* `ConductoroneApi.AppUserMergeRules.ListSources()`: **Added**
* `ConductoroneApi.AutomationExecutionActions.RetryAutomationExecutions()`: **Added**
* `ConductoroneApi.VirtualMcpServer.Search()`: **Added**
* `ConductoroneApi.VirtualMcpServer.Summary()`: **Added**
* `ConductoroneApi.AccessConflict.CreateMonitor()`: 
  * `request.Request` **Changed**
    - `Rule` **Added**
    - `Severity` **Added**
  * `response` **Changed**
    - `Rule` **Added**
    - `Severity` **Added**
* `ConductoroneApi.AccessConflict.GetMonitor()`: `response` **Changed**
    - `Rule` **Added**
    - `Severity` **Added**
* `ConductoroneApi.AccessConflict.UpdateMonitor()`: 
  * `request.Request.ConflictMonitorUpdateRequest` **Changed**
    - `Rule` **Added**
    - `Severity` **Added**
  * `response` **Changed**
    - `Rule` **Added**
    - `Severity` **Added**
* `ConductoroneApi.AppUser.List()`:  `response.List[].AppUser.Sources` **Added**
* `ConductoroneApi.AppUser.ListAppUsersForUser()`:  `response.List[].AppUser.Sources` **Added**
* `ConductoroneApi.AppUser.ListOwnedServiceAccounts()`:  `response.List[].AppUser.Sources` **Added**
* `ConductoroneApi.AppUser.Merge()`: **Deprecated**
* `ConductoroneApi.AppUser.Merge()`:  `response.AppUserView.AppUser.Sources` **Added**
* `ConductoroneApi.AppUser.Search()`: 
  * `request.Request` **Changed**
    - `ExcludeSourceAppIds` **Added**
    - `ExcludeSourceConnectorIds` **Added**
    - `SourceAppIds` **Added**
    - `SourceConnectorIds` **Added**
  *  `response.List[].AppUser.Sources` **Added**
* `ConductoroneApi.AppUser.Update()`:  `response.AppUserView.AppUser.Sources` **Added**
* `ConductoroneApi.Edge.Create()`: `response.Edge` **Changed**
    - `EgressRules[].ResponseInspectionProfile` **Added**
    - `Inference.Targets[].VertexChat` **Added**
    - `InferenceConnections` **Added**
    - `ResponseInspectionProfile` **Added**
* `ConductoroneApi.Edge.Get()`: `response.Edge` **Changed**
    - `EgressRules[].ResponseInspectionProfile` **Added**
    - `Inference.Targets[].VertexChat` **Added**
    - `InferenceConnections` **Added**
    - `ResponseInspectionProfile` **Added**
* `ConductoroneApi.Edge.List()`: `response.List[]` **Changed**
    - `EgressRules[].ResponseInspectionProfile` **Added**
    - `Inference.Targets[].VertexChat` **Added**
    - `InferenceConnections` **Added**
    - `ResponseInspectionProfile` **Added**
* `ConductoroneApi.Edge.ListInferenceCatalog()`: `response` **Changed**
    - `DiscoveredConnectionModels` **Added**
    - `DiscoveredTargetModels` **Added**
* `ConductoroneApi.Edge.PreviewEgressHostDecision()`: 
  *  `request.Request.EdgeServicePreviewEgressHostDecisionRequest.DraftRules.Rules[].ResponseInspectionProfile` **Added**
  *  `response.ResultingRules.Rules[].ResponseInspectionProfile` **Added**
* `ConductoroneApi.Edge.PreviewEgressRulesUpdate()`: 
  *  `request.Request.EdgeServicePreviewEgressRulesUpdateRequest.ProposedRules.Rules[].ResponseInspectionProfile` **Added**
* `ConductoroneApi.Edge.Update()`: 
  * `request.Request.EdgeServiceUpdateRequest.Edge` **Changed**
    - `EgressRules[].ResponseInspectionProfile` **Added**
    - `Inference.Targets[].VertexChat` **Added**
    - `InferenceConnections` **Added**
    - `ResponseInspectionProfile` **Added**
  * `response.Edge` **Changed**
    - `EgressRules[].ResponseInspectionProfile` **Added**
    - `Inference.Targets[].VertexChat` **Added**
    - `InferenceConnections` **Added**
    - `ResponseInspectionProfile` **Added**
* `ConductoroneApi.AppEntitlements.ListUsers()`:  `response.List[].AppUser.AppUser.Sources` **Added**
* `ConductoroneApi.AppEntitlementSearch.SearchAppEntitlementsWithExpired()`:  `response.List[].AppUser.Sources` **Added**
* `ConductoroneApi.AppEntitlementSearch.SearchGrants()`:  `response.List[].AppEntitlementUserBinding.AppUser.AppUser.Sources` **Added**
* `ConductoroneApi.AutomationExecution.GetAutomationExecution()`:  `response.AutomationExecution.LastRetriedAt` **Added**
* `ConductoroneApi.AutomationExecution.ListAutomationExecutions()`:  `response.AutomationExecutions[].LastRetriedAt` **Added**
* `ConductoroneApi.AutomationExecutionSearch.SearchAllAutomationExecutions()`:  `response.List[].AutomationExecution.LastRetriedAt` **Added**
* `ConductoroneApi.AutomationExecutionSearch.SearchAutomationExecutions()`:  `response.List[].AutomationExecution.LastRetriedAt` **Added**
* `ConductoroneApi.RequestCatalogSearch.SearchRequestableCredentialOfferings()`: `response.List[]` **Changed**
    - `FixedArgs` **Added**
    - `FixedArgumentDisplayNames` **Added**
    - `InputConfigurationHash` **Added**
    - `RequesterArguments` **Added**
* `ConductoroneApi.TaskSearch.Search()`: 
  * `request.Request` **Changed**
    - `AttentionCauses` **Added**
    - `NeedsAttention` **Added**
  * `response.List[]` **Changed**
    - `Attention` **Added**
    - `CredentialIssuePath` **Added**
    - `Task.Type.Action.CredentialIssue.FixedArgs` **Added**
* `ConductoroneApi.Task.CreateActionTask()`: `response.TaskView` **Changed**
    - `Attention` **Added**
    - `CredentialIssuePath` **Added**
    - `Task.Type.Action.CredentialIssue.FixedArgs` **Added**
* `ConductoroneApi.Task.CreateGrantTask()`: `response.TaskView` **Changed**
    - `Attention` **Added**
    - `CredentialIssuePath` **Added**
    - `Task.Type.Action.CredentialIssue.FixedArgs` **Added**
* `ConductoroneApi.Task.CreateOffboardingTask()`: `response.TaskView` **Changed**
    - `Attention` **Added**
    - `CredentialIssuePath` **Added**
    - `Task.Type.Action.CredentialIssue.FixedArgs` **Added**
* `ConductoroneApi.Task.CreateResourceActionTask()`: `response.TaskView` **Changed**
    - `Attention` **Added**
    - `CredentialIssuePath` **Added**
    - `Task.Type.Action.CredentialIssue.FixedArgs` **Added**
* `ConductoroneApi.Task.CreateRevokeTask()`: `response.TaskView` **Changed**
    - `Attention` **Added**
    - `CredentialIssuePath` **Added**
    - `Task.Type.Action.CredentialIssue.FixedArgs` **Added**
* `ConductoroneApi.Task.Get()`: `response.TaskView` **Changed**
    - `Attention` **Added**
    - `CredentialIssuePath` **Added**
    - `Task.Type.Action.CredentialIssue.FixedArgs` **Added**
* `ConductoroneApi.TaskActions.Approve()`: `response.TaskView` **Changed**
    - `Attention` **Added**
    - `CredentialIssuePath` **Added**
    - `Task.Type.Action.CredentialIssue.FixedArgs` **Added**
* `ConductoroneApi.TaskActions.ApproveWithStepUp()`: `response.TaskView` **Changed**
    - `Attention` **Added**
    - `CredentialIssuePath` **Added**
    - `Task.Type.Action.CredentialIssue.FixedArgs` **Added**
* `ConductoroneApi.TaskActions.Close()`: `response.TaskView` **Changed**
    - `Attention` **Added**
    - `CredentialIssuePath` **Added**
    - `Task.Type.Action.CredentialIssue.FixedArgs` **Added**
* `ConductoroneApi.TaskActions.Comment()`: `response.TaskView` **Changed**
    - `Attention` **Added**
    - `CredentialIssuePath` **Added**
    - `Task.Type.Action.CredentialIssue.FixedArgs` **Added**
* `ConductoroneApi.TaskActions.Deny()`: `response.TaskView` **Changed**
    - `Attention` **Added**
    - `CredentialIssuePath` **Added**
    - `Task.Type.Action.CredentialIssue.FixedArgs` **Added**
* `ConductoroneApi.TaskActions.EscalateToEmergencyAccess()`: `response.TaskView` **Changed**
    - `Attention` **Added**
    - `CredentialIssuePath` **Added**
    - `Task.Type.Action.CredentialIssue.FixedArgs` **Added**
* `ConductoroneApi.TaskActions.HardReset()`: `response.TaskView` **Changed**
    - `Attention` **Added**
    - `CredentialIssuePath` **Added**
    - `Task.Type.Action.CredentialIssue.FixedArgs` **Added**
* `ConductoroneApi.TaskActions.ProcessNow()`: `response.TaskView` **Changed**
    - `Attention` **Added**
    - `CredentialIssuePath` **Added**
    - `Task.Type.Action.CredentialIssue.FixedArgs` **Added**
* `ConductoroneApi.TaskActions.Reassign()`: `response.TaskView` **Changed**
    - `Attention` **Added**
    - `CredentialIssuePath` **Added**
    - `Task.Type.Action.CredentialIssue.FixedArgs` **Added**
* `ConductoroneApi.TaskActions.Restart()`: `response.TaskView` **Changed**
    - `Attention` **Added**
    - `CredentialIssuePath` **Added**
    - `Task.Type.Action.CredentialIssue.FixedArgs` **Added**
* `ConductoroneApi.TaskActions.RetryProvisioning()`: `response.TaskView` **Changed**
    - `Attention` **Added**
    - `CredentialIssuePath` **Added**
    - `Task.Type.Action.CredentialIssue.FixedArgs` **Added**
* `ConductoroneApi.TaskActions.SkipStep()`: `response.TaskView` **Changed**
    - `Attention` **Added**
    - `CredentialIssuePath` **Added**
    - `Task.Type.Action.CredentialIssue.FixedArgs` **Added**
* `ConductoroneApi.TaskActions.UpdateGrantDuration()`: `response.TaskView` **Changed**
    - `Attention` **Added**
    - `CredentialIssuePath` **Added**
    - `Task.Type.Action.CredentialIssue.FixedArgs` **Added**
* `ConductoroneApi.TaskActions.UpdateRequestData()`: `response.TaskView` **Changed**
    - `Attention` **Added**
    - `CredentialIssuePath` **Added**
    - `Task.Type.Action.CredentialIssue.FixedArgs` **Added**
