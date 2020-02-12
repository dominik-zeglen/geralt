//go:generate stringer -type=FlowState
package flow

type FlowState int

const (
	Default FlowState = iota

	SettingBotName

	SettingSpeakerName
)
