package flow

import (
	"github.com/looplab/fsm"
)

var events fsm.Events

func init() {
	events = fsm.Events{
		{
			Name: ToBotNameSetting.String(),
			Src: []string{
				Default.String(),
			},
			Dst: SettingBotName.String(),
		},
		{
			Name: BotNameSet.String(),
			Src: []string{
				SettingBotName.String(),
			},
			Dst: Default.String(),
		},
		{
			Name: ToSpeakerNameSetting.String(),
			Src: []string{
				Default.String(),
			},
			Dst: SettingSpeakerName.String(),
		},
		{
			Name: SpeakerNameSet.String(),
			Src: []string{
				SettingSpeakerName.String(),
			},
			Dst: Default.String(),
		},
	}
}

func NewFlow() *fsm.FSM {
	return fsm.NewFSM(
		string(Default),
		events,
		fsm.Callbacks{},
	)
}
