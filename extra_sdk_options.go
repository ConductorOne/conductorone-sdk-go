package conductoronesdkgo

import (
	"context"
	"crypto/tls"
	"net/http"

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

type CustomOptions struct {
	useWithServer bool
	useWithTenant bool

	ServerURL string
	Tenant    string

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
	if options.ServerURL != "" {
		serverURL = options.ServerURL
	} else {
		resp, err := NormalizeTenant(cred.ClientID)
		if err != nil {
			return nil, err
		}
		serverURL = resp.ServerURL
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
