package conductoronesdkgo

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"go.uber.org/zap"

	"github.com/conductorone/conductorone-sdk-go/uhttp"
)

const ClientIdGolangSDK = "2RCzHlak5q7CY14SdBc8HoZEJRf"

func WithTenant(input string) (SDKOption, error) {
	resp, err := NormalizeTenant(input)
	if err != nil {
		return nil, err
	}

	if resp.useWithTenant {
		return WithTenantDomain(resp.Tenant), nil
	}

	if resp.useWithServer {
		return WithServerURL(resp.ServerURL), nil
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
		sdk.useWithServer = resp.useWithServer
		sdk.useWithTenant = resp.useWithTenant
		sdk.ServerURL = resp.ServerURL
		sdk.Tenant = resp.Tenant
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
	useWithServer bool
	useWithTenant bool

	ServerURL string
	Tenant    string
}

func (c ClientConfig) GetServerURL() string {
	if c.useWithServer {
		return c.ServerURL
	}
	if c.useWithTenant {
		return fmt.Sprintf("https://%s.conductor.one", c.Tenant)
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

	var serverURL string
	if options.GetServerURL() != "" {
		serverURL = options.GetServerURL()
	} else {
		resp, err := ParseClientID(cred.ClientID)
		if err != nil {
			return nil, err
		}
		serverURL = resp.GetServerURL()
		if serverURL == "" {
			return nil, ErrInvalidClientID
		}
	}

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
	if options.useWithServer {
		sdkOpts = append(sdkOpts, WithServerURL(options.ServerURL))
	}
	if options.useWithTenant {
		sdkOpts = append(sdkOpts, WithTenantDomain(options.Tenant))
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
			input += ".conductor.one"
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
		normalize.useWithTenant = true
		normalize.Tenant = parts[0]

		return normalize, nil
	}

	u.Scheme = "https"
	normalize.useWithServer = true
	normalize.ServerURL = u.String()
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

	return resp, nil
}
