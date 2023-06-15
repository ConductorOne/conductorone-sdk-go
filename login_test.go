package conductoroneapi

import (
	"fmt"
	"testing"
)

func TestLoginFlow(t *testing.T) {
	creds, err := LoginFlow("c1dev.logan.dev.ductone.com:2443", "2RCzHlak5q7CY14SdBc8HoZEJRf")
	fmt.Printf("Creds: %+v, Err: %+v", creds, err)
}
