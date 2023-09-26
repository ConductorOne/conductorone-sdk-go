package conductoronesdkgo

import (
	"context"

	c1 "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

// This is a test to make sure that the generated code still builds, it attemps to use a variety of methods from the SDK.
func main() {
	ctx := context.Background()
	client := c1.New()
	getTest(ctx, client, "123")
	createTest(ctx, client, &shared.TaskServiceCreateGrantRequest{})
	updateTest(ctx, client, "123", "comment", "policyId")
	searchTest(ctx, client)
}

// Get
func getTest(ctx context.Context, client *c1.ConductoroneAPI, taskId string) (*shared.TaskServiceGetResponse, error) {
	resp, err := client.Task.Get(ctx, operations.C1APITaskV1TaskServiceGetRequest{ID: taskId})
	if err != nil {
		return nil, err
	}

	return resp.TaskServiceGetResponse, err
}

// Create
func createTest(ctx context.Context, client *c1.ConductoroneAPI, request *shared.TaskServiceCreateGrantRequest) (*operations.C1APITaskV1TaskServiceCreateGrantTaskResponse, error) {
	resp, err := client.Task.CreateGrantTask(ctx, request)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// Update
func updateTest(ctx context.Context, client *c1.ConductoroneAPI, taskId string, comment string, policyId string) (*shared.TaskActionsServiceApproveResponse, error) {
	resp, err := client.TaskActions.Approve(ctx, operations.C1APITaskV1TaskActionsServiceApproveRequest{
		TaskActionsServiceApproveRequest: &shared.TaskActionsServiceApproveRequest{
			Comment:      &comment,
			PolicyStepID: policyId,
		},
		TaskID: taskId,
	})
	if err != nil {
		return nil, err
	}

	return resp.TaskActionsServiceApproveResponse, nil
}

// Search
func searchTest(ctx context.Context, client *c1.ConductoroneAPI) (*operations.C1APIRequestcatalogV1RequestCatalogSearchServiceSearchEntitlementsResponse, error) {
	resp, err := client.RequestCatalogSearch.SearchEntitlements(ctx, shared.RequestCatalogSearchServiceSearchEntitlementsRequest{
		PageToken: nil,
	})
	if err != nil {
		return nil, err
	}

	return resp, err
}
