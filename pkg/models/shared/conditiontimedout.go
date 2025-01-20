// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/utils"
	"time"
)

// The ConditionTimedOut message.
type ConditionTimedOut struct {
	TimedOutAt *time.Time `json:"timedOutAt,omitempty"`
}

func (c ConditionTimedOut) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *ConditionTimedOut) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ConditionTimedOut) GetTimedOutAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.TimedOutAt
}