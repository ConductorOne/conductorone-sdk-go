package conductoroneapi

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type ClientCredentials struct {
	ClientID     string
	ClientSecret string
}

type DeviceCodeResponse struct {
	DeviceCode      string `json:"device_code"`
	UserCode        string `json:"user_code"`
	VerificationURI string `json:"verification_uri"`
	ExpiresIn       int64  `json:"expires_in"`
	Interval        int64  `json:"interval"`
}

func LoginFlow(ctx context.Context, tenantName string, clientID string) (*ClientCredentials, error) {

	opts := []SDKOption{}
	// If they pass a URL, use the whole URL
	if strings.Contains(tenantName, ".") {
		opts = append(opts, WithServerURL(tenantName))
	} else {
		opts = append(opts, WithTenantDomain(tenantName))
	}
	client := New(opts...)
	httpClient := client.sdkConfiguration.DefaultClient

	deviceCodeURL := "https://" + client.sdkConfiguration.ServerURL + "/auth/v1/device_authorization"
	vals := url.Values{}
	vals.Add("client_id", clientID)

	req, err := http.NewRequest("POST", deviceCodeURL, strings.NewReader(vals.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Response: %s\n", data)

	codeResp := DeviceCodeResponse{}
	err = json.Unmarshal(data, &codeResp)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Open: %s\n", codeResp.VerificationURI)

	if _, _, err := doTokenRequest(ctx, clientID, client, codeResp); err != nil {
		return nil, err
	}

	return nil, errors.New("failure")
}

func doTokenRequest(ctx context.Context, clientID string, client *ConductoroneAPI, deviceCodeResp DeviceCodeResponse) (string, string, error) {
	httpClient := client.sdkConfiguration.DefaultClient

	tokenURL := "https://" + client.sdkConfiguration.ServerURL + "/auth/v1/token"
	vals := url.Values{}
	vals.Add("client_id", clientID)
	vals.Add("device_code", deviceCodeResp.DeviceCode)
	vals.Add("grant_type", "urn:ietf:params:oauth:grant-type:device_code")

	startLoop := time.Now()
	for {
		select {
		case <-time.After(time.Duration(deviceCodeResp.Interval) * time.Second):
			if time.Now().After(startLoop.Add(time.Duration(deviceCodeResp.ExpiresIn) * time.Second)) {
				return "", "", errors.New("timeout")
			}
		case <-ctx.Done():
			return "", "", nil
		}

		req, err := http.NewRequest("POST", tokenURL, strings.NewReader(vals.Encode()))
		if err != nil {
			return "", "", err
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		resp, err := httpClient.Do(req)
		if err != nil {
			return "", "", err
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return "", "", err
		}

		fmt.Printf("Token URL Resp: %s\n", body)
	}
}
