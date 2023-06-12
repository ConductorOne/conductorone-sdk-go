// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package conductoroneapi

import (
	"fmt"
	"github.com/conductorone/conductorone-sdk-go/pkg/utils"
	"net/http"
	"time"
)

// ServerList contains the list of servers available to the SDK
var ServerList = []string{
	// The server for c1.api.accessbundle.v1.AccessBundleSearchService.
	"/",
}

// HTTPClient provides an interface for suplying the SDK with a custom HTTP client
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// String provides a helper function to return a pointer to a string
func String(s string) *string { return &s }

// Bool provides a helper function to return a pointer to a bool
func Bool(b bool) *bool { return &b }

// Int provides a helper function to return a pointer to an int
func Int(i int) *int { return &i }

// Int64 provides a helper function to return a pointer to an int64
func Int64(i int64) *int64 { return &i }

// Float32 provides a helper function to return a pointer to a float32
func Float32(f float32) *float32 { return &f }

// Float64 provides a helper function to return a pointer to a float64
func Float64(f float64) *float64 { return &f }

type sdkConfiguration struct {
	DefaultClient  HTTPClient
	SecurityClient HTTPClient

	ServerURL         string
	ServerIndex       int
	Language          string
	OpenAPIDocVersion string
	SDKVersion        string
	GenVersion        string
}

func (c *sdkConfiguration) GetServerDetails() (string, map[string]string) {
	if c.ServerURL != "" {
		return c.ServerURL, nil
	}

	return ServerList[c.ServerIndex], nil
}

// ConductoroneAPI - API For c1.api.accessbundle.v1.AccessBundleSearchService: This is an auto-generated API for c1.api.accessbundle.v1.AccessBundleSearchService.
type ConductoroneAPI struct {
	AppEntitlementUserBinding *appEntitlementUserBinding
	AppEntitlements           *appEntitlements
	AppResource               *appResource
	AppResourceType           *appResourceType
	Apps                      *apps
	Auth                      *auth
	RequestCatalogSearch      *requestCatalogSearch
	Task                      *task
	TaskActions               *taskActions
	TaskSearch                *taskSearch
	User                      *user

	sdkConfiguration sdkConfiguration
}

type SDKOption func(*ConductoroneAPI)

// WithServerURL allows the overriding of the default server URL
func WithServerURL(serverURL string) SDKOption {
	return func(sdk *ConductoroneAPI) {
		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithTemplatedServerURL allows the overriding of the default server URL with a templated URL populated with the provided parameters
func WithTemplatedServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *ConductoroneAPI) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithServerIndex allows the overriding of the default server by index
func WithServerIndex(serverIndex int) SDKOption {
	return func(sdk *ConductoroneAPI) {
		if serverIndex < 0 || serverIndex >= len(ServerList) {
			panic(fmt.Errorf("server index %d out of range", serverIndex))
		}

		sdk.sdkConfiguration.ServerIndex = serverIndex
	}
}

// WithClient allows the overriding of the default HTTP client used by the SDK
func WithClient(client HTTPClient) SDKOption {
	return func(sdk *ConductoroneAPI) {
		sdk.sdkConfiguration.DefaultClient = client
	}
}

// New creates a new instance of the SDK with the provided options
func New(opts ...SDKOption) *ConductoroneAPI {
	sdk := &ConductoroneAPI{
		sdkConfiguration: sdkConfiguration{
			Language:          "go",
			OpenAPIDocVersion: "0.0.1",
			SDKVersion:        "1.0.2",
			GenVersion:        "2.39.2",
		},
	}
	for _, opt := range opts {
		opt(sdk)
	}

	// Use WithClient to override the default client if you would like to customize the timeout
	if sdk.sdkConfiguration.DefaultClient == nil {
		sdk.sdkConfiguration.DefaultClient = &http.Client{Timeout: 60 * time.Second}
	}
	if sdk.sdkConfiguration.SecurityClient == nil {
		sdk.sdkConfiguration.SecurityClient = sdk.sdkConfiguration.DefaultClient
	}

	sdk.AppEntitlementUserBinding = newAppEntitlementUserBinding(sdk.sdkConfiguration)

	sdk.AppEntitlements = newAppEntitlements(sdk.sdkConfiguration)

	sdk.AppResource = newAppResource(sdk.sdkConfiguration)

	sdk.AppResourceType = newAppResourceType(sdk.sdkConfiguration)

	sdk.Apps = newApps(sdk.sdkConfiguration)

	sdk.Auth = newAuth(sdk.sdkConfiguration)

	sdk.RequestCatalogSearch = newRequestCatalogSearch(sdk.sdkConfiguration)

	sdk.Task = newTask(sdk.sdkConfiguration)

	sdk.TaskActions = newTaskActions(sdk.sdkConfiguration)

	sdk.TaskSearch = newTaskSearch(sdk.sdkConfiguration)

	sdk.User = newUser(sdk.sdkConfiguration)

	return sdk
}
