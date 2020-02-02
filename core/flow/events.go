//go:generate stringer -type=FlowEvent
package flow

type FlowEvent int

const (
	Back FlowEvent = iota

	ToBotNameSetting
	BotNameSet

	ToSpeakerNameSetting
	SpeakerNameSet
)
