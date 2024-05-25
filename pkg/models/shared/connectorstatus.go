// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/conductorone/conductorone-sdk-go/pkg/utils"
	"time"
)

// ConnectorStatusStatus - The status of the connector sync.
type ConnectorStatusStatus string

const (
	ConnectorStatusStatusSyncStatusUnspecified ConnectorStatusStatus = "SYNC_STATUS_UNSPECIFIED"
	ConnectorStatusStatusSyncStatusRunning     ConnectorStatusStatus = "SYNC_STATUS_RUNNING"
	ConnectorStatusStatusSyncStatusDone        ConnectorStatusStatus = "SYNC_STATUS_DONE"
	ConnectorStatusStatusSyncStatusError       ConnectorStatusStatus = "SYNC_STATUS_ERROR"
)

func (e ConnectorStatusStatus) ToPointer() *ConnectorStatusStatus {
	return &e
}
func (e *ConnectorStatusStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SYNC_STATUS_UNSPECIFIED":
		fallthrough
	case "SYNC_STATUS_RUNNING":
		fallthrough
	case "SYNC_STATUS_DONE":
		fallthrough
	case "SYNC_STATUS_ERROR":
		*e = ConnectorStatusStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectorStatusStatus: %v", v)
	}
}

// ConnectorStatus - The status field on the connector is used to track the status of the connectors sync, and when syncing last started, completed, or caused the connector to update.
type ConnectorStatus struct {
	CompletedAt *time.Time `json:"completedAt,omitempty"`
	// The last error encountered by the connector.
	LastError *string    `json:"lastError,omitempty"`
	StartedAt *time.Time `json:"startedAt,omitempty"`
	// The status of the connector sync.
	Status    *ConnectorStatusStatus `json:"status,omitempty"`
	UpdatedAt *time.Time             `json:"updatedAt,omitempty"`
}

func (c ConnectorStatus) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *ConnectorStatus) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ConnectorStatus) GetCompletedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CompletedAt
}

func (o *ConnectorStatus) GetLastError() *string {
	if o == nil {
		return nil
	}
	return o.LastError
}

func (o *ConnectorStatus) GetStartedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.StartedAt
}

func (o *ConnectorStatus) GetStatus() *ConnectorStatusStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *ConnectorStatus) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}
