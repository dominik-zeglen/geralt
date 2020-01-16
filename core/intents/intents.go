package intents

// Basic intents - small talk
const Hello Intent = "basic.hello"
const Status Intent = "basic.status"

// Activation
const Start Intent = "conversation.on"
const Stop Intent = "conversation.off"

// Intents regarding speaker
const SetName Intent = "speaker.setName"
const GetName Intent = "speaker.getName"

// Intents regarding bot
const SetBotName Intent = "bot.setName"
const GetBotName Intent = "bot.getName"
const BotStatus Intent = "bot.status"

// Decision intents
const Repeat Intent = "decision.repeat"
const Yes Intent = "decision.yes"
const No Intent = "decision.no"
const Back Intent = "decision.back"
