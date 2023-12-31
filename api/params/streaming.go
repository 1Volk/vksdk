package params // import "github.com/1Volk/vksdk/api/params"

import (
	"github.com/1Volk/vksdk/api"
)

// StreamingSetSettingsBuilder builder.
//
// https://vk.com/dev/streaming.setSettings
type StreamingSetSettingsBuilder struct {
	api.Params
}

// NewStreamingSetSettingsBuilder func.
func NewStreamingSetSettingsBuilder() *StreamingSetSettingsBuilder {
	return &StreamingSetSettingsBuilder{api.Params{}}
}

// MonthlyTier parameter.
func (b *StreamingSetSettingsBuilder) MonthlyTier(v string) *StreamingSetSettingsBuilder {
	b.Params["monthly_tier"] = v
	return b
}
