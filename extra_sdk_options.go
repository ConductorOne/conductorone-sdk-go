package conductoronesdkgo

import (
	"context"
	"crypto/tls"
	"net/http"
	"net/url"
	"strings"

	"go.uber.org/zap"

	"github.com/conductorone/conductorone-sdk-go/uhttp"
)

const c1TenantDomain = ".conductor.one"
const ClientIdGolangSDK = "2RCzHlak5q7CY14SdBc8HoZEJRf"

func WithTenant(input string) (SDKOption, error) {
	resp, err := NormalizeTenant(input)
	if err != nil {
		return nil, err
	}

	if resp.UseWithTenant() {
		return WithTenantDomain(resp.Tenant()), nil
	}

	if resp.UseWithServer() {
		return WithServerURL(resp.ServerURL()), nil
	}

	return func(api *ConductoroneAPI) {}, nil
}

type CustomSDKOption func(*CustomOptions)

func WithTenantCustom(input string) (CustomSDKOption, error) {
	resp, err := NormalizeTenant(input)
	if err != nil {
		return nil, err
	}

	return func(sdk *CustomOptions) {
		sdk.serverURL = resp.ServerURL()
		sdk.tenant = resp.Tenant()
	}, nil
}

func WithLog(logger *zap.Logger) CustomSDKOption {
	return func(sdk *CustomOptions) {
		sdk.logger = logger
	}
}

func WithUserAgent(userAgent string) CustomSDKOption {
	return func(sdk *CustomOptions) {
		sdk.userAgent = userAgent
	}
}
func WithTLSConfig(tlsConfig *tls.Config) CustomSDKOption {
	return func(sdk *CustomOptions) {
		sdk.tlsConfig = tlsConfig
	}
}

type ClientConfig struct {
	serverURL string
	tenant    string
}

func (c ClientConfig) UseWithServer() bool {
	return c.serverURL != ""
}

func (c ClientConfig) UseWithTenant() bool {
	return c.tenant != ""
}

func (c ClientConfig) Tenant() string {
	return c.tenant
}

// ServerURL returns the server URL if it is set, will NOT construct the server URL from the tenant
func (c ClientConfig) ServerURL() string {
	return c.serverURL
}

// GetServerURL returns the server URL if it is set, otherwise it constructs the server URL from the tenant, if neither are not set it will return an empty string
func (c ClientConfig) GetServerURL() string {
	if c.UseWithServer() {
		return c.serverURL
	}
	if c.UseWithTenant() {
		u := &url.URL{}
		tenant := strings.ToLower(c.Tenant())
		u.Host = tenant + c1TenantDomain
		u.Scheme = "https"
		return u.String()
	}
	return ""
}

type CustomOptions struct {
	ClientConfig

	withClient *http.Client
	logger     *zap.Logger
	tlsConfig  *tls.Config

	userAgent string
}

func NewWithCredentials(ctx context.Context, cred *ClientCredentials, opts ...CustomSDKOption) (*ConductoroneAPI, error) {
	options := &CustomOptions{}
	for _, opt := range opts {
		opt(options)
	}
	if options.GetServerURL() == "" {
		resp, err := ParseClientID(cred.ClientID)
		if err != nil {
			return nil, err
		}
		options.ClientConfig = *resp
	}

	serverURL := options.GetServerURL()
	tokenSource, err := NewTokenSource(ctx, cred.ClientID, cred.ClientSecret, serverURL)
	if err != nil {
		return nil, err
	}

	if options.userAgent == "" {
		options.userAgent = "conductorone-sdk-go"
	}

	uclient, err := uhttp.NewClient(
		ctx,
		uhttp.WithTokenSource(tokenSource),
		uhttp.WithLogger(options.logger != nil, options.logger),
		uhttp.WithUserAgent(options.userAgent),
		uhttp.WithTLSClientConfig(options.tlsConfig),
	)
	if err != nil {
		return nil, err
	}

	sdkOpts := []SDKOption{}
	if options.UseWithServer() {
		sdkOpts = append(sdkOpts, WithServerURL(options.ServerURL()))
	}
	if options.UseWithTenant() {
		sdkOpts = append(sdkOpts, WithTenantDomain(options.Tenant()))
	}
	sdkOpts = append(sdkOpts, WithClient(uclient))

	return New(sdkOpts...), nil
}

func NormalizeTenant(input string) (*ClientConfig, error) {
	input = strings.ToLower(input)

	var err error
	u := &url.URL{}
	if !strings.Contains(input, "//") {
		if !strings.Contains(input, ".") {
			input += c1TenantDomain
		}
		u.Host = input
	} else {
		u, err = url.Parse(input)
		if err != nil {
			return nil, err
		}
	}

	normalize := &ClientConfig{}

	parts := strings.Split(u.Host, ".")
	if len(parts) == 3 && parts[1] == "conductor" && parts[2] == "one" {
		normalize.tenant = parts[0]

		return normalize, nil
	}

	u.Scheme = "https"
	normalize.serverURL = u.String()
	return normalize, nil
}

func ParseClientID(input string) (*ClientConfig, error) {
	// split the input into 2 parts by @
	items := strings.SplitN(input, "@", 2)
	if len(items) != 2 {
		return nil, ErrInvalidClientID
	}

	// split the right part into 2 parts by /
	items = strings.SplitN(items[1], "/", 2)
	if len(items) != 2 {
		return nil, ErrInvalidClientID
	}

	resp, err := NormalizeTenant(items[0])
	if err != nil {
		return nil, err
	}

	if resp.GetServerURL() == "" {
		return nil, ErrInvalidClientID
	}

	return resp, nil
}
