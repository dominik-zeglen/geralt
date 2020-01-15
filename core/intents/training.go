package intents

type trainingDataset []trainingData
type trainingData struct {
	sentences []string
	intent    Intent
}

func (t trainingDataset) count() int {
	counter := 0
	for _, intent := range t {
		counter += len(intent.sentences)
	}

	return counter
}

// This is some kickstart data
// Geralt should be learning
func getTrainingData() trainingDataset {
	return []trainingData{
		trainingData{
			intent: Hello,
			sentences: []string{
				"Hi",
				"Hey",
				"Hello",
				"Sup",
				"What's up",
				"Morning",
				"Good morning",
				"Good evening",
				"Good afternoon",
			},
		},

		trainingData{
			intent: Status,
			sentences: []string{
				"How are you",
				"How's going",
				"How's your day",
				"Everything is okay",
			},
		},

		trainingData{
			intent: Start,
			sentences: []string{
				"Hey {{.Botname}}",
				"Wake up {{.Botname}}",
			},
		},

		trainingData{
			intent: Stop,
			sentences: []string{
				"That's all",
				"Sleep now",
				"Good night",
				"I need to go",
				"Bye {{.Botname}}",
			},
		},

		trainingData{
			intent: SetName,
			sentences: []string{
				"I'm Alfred",
				"My name is Todd",
				"I'm Monica",
				"Call me Jack",
			},
		},

		trainingData{
			intent: GetName,
			sentences: []string{
				"Who are you speaking to",
				"{{.Botname}}, what's my name",
			},
		},

		trainingData{
			intent: SetBotName,
			sentences: []string{
				"Your name is {{.Botname}} since now",
				"You are {{.Botname}}",
				"You're {{.Botname}}",
			},
		},

		trainingData{
			intent: GetBotName,
			sentences: []string{
				"Who are you",
				"What's your name",
				"And your name is",
				"How they call you",
				"How should I call you",
			},
		},

		trainingData{
			intent: BotStatus,
			sentences: []string{
				"Get your status",
				"Is technically ok with you",
				"Everything works",
			},
		},

		trainingData{
			intent: Repeat,
			sentences: []string{
				"Repeat, please",
				"Could you repeat",
				"Once again, please",
				"Sorry, what",
				"What",
				"Repeat, dude",
			},
		},

		trainingData{
			intent: Yes,
			sentences: []string{
				"Yes",
				"Yup",
				"Ok",
				"Naturally",
				"Of course",
				"Yeah",
				"Sure",
			},
		},

		trainingData{
			intent: No,
			sentences: []string{
				"No",
				"Nope",
				"Nah",
			},
		},

		trainingData{
			intent: Back,
			sentences: []string{
				"Nevermind",
				"Doesn't matter",
				"Forget it",
				"Go back",
			},
		},
	}
}
