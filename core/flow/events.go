//go:generate stringer -type=FlowEvent
package flow

type FlowEvent int

const (
	ToBotNameSetting FlowEvent = iota
	BotNameSet

	ToSpeakerNameSetting
	SpeakerNameSet
)
