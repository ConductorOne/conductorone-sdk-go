## Go SDK Changes:
* `ConductoroneApi.AppUserMergeRules.ListSources()`: **Added**
* `ConductoroneApi.AutomationExecutionActions.RetryAutomationExecutions()`: **Added**
* `ConductoroneApi.Edge.Create()`: `response.Edge` **Changed**
    - `EgressRules[].ResponseInspectionProfile` **Added**
    - `InferenceConnections` **Added**
    - `ResponseInspectionProfile` **Added**
* `ConductoroneApi.Edge.Get()`: `response.Edge` **Changed**
    - `EgressRules[].ResponseInspectionProfile` **Added**
    - `InferenceConnections` **Added**
    - `ResponseInspectionProfile` **Added**
* `ConductoroneApi.Edge.List()`: `response.List[]` **Changed**
    - `EgressRules[].ResponseInspectionProfile` **Added**
    - `InferenceConnections` **Added**
    - `ResponseInspectionProfile` **Added**
* `ConductoroneApi.Edge.PreviewEgressHostDecision()`: 
  *  `request.Request.EdgeServicePreviewEgressHostDecisionRequest.DraftRules.Rules[].ResponseInspectionProfile` **Added**
  *  `response.ResultingRules.Rules[].ResponseInspectionProfile` **Added**
* `ConductoroneApi.Edge.PreviewEgressRulesUpdate()`: 
  *  `request.Request.EdgeServicePreviewEgressRulesUpdateRequest.ProposedRules.Rules[].ResponseInspectionProfile` **Added**
* `ConductoroneApi.Edge.Update()`: 
  * `request.Request.EdgeServiceUpdateRequest.Edge` **Changed**
    - `EgressRules[].ResponseInspectionProfile` **Added**
    - `InferenceConnections` **Added**
    - `ResponseInspectionProfile` **Added**
  * `response.Edge` **Changed**
    - `EgressRules[].ResponseInspectionProfile` **Added**
    - `InferenceConnections` **Added**
    - `ResponseInspectionProfile` **Added**
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
