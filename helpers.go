package conductoronesdkgo

import (
	"net/url"
	"strings"
)

type NormalizedTenant struct {
	useWithServer bool
	useWithTenant bool

	ServerURL string
	Tenant    string
}

func NormalizeTenant(input string) (*NormalizedTenant, error) {
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

	normalize := &NormalizedTenant{}

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

func ParseClientID(input string) (*NormalizedTenant, error) {
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
