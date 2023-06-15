package handlers

import (
	"github.com/bwmarrin/discordgo"
)

func (handler Handler) InteractionHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	switch i.Type {
	case discordgo.InteractionApplicationCommand: // for commands
		if h, ok := componentsHandlers[i.ApplicationCommandData().Name]; ok {
			h(handler, s, i)
			return
		}
		handler.Log.Errorf("Unhandled interaction for commands: %s", i.ApplicationCommandData().Name)
	case discordgo.InteractionMessageComponent: // for message component interaction
		if h, ok := componentsHandlers[i.MessageComponentData().CustomID]; ok {
			h(handler, s, i)
			return
		}
		handler.Log.Errorf("Unhandled interaction for components: %s", i.MessageComponentData().CustomID)
	case discordgo.InteractionModalSubmit: // for modal submit interaction
		if h, ok := componentsHandlers[i.ModalSubmitData().CustomID]; ok {
			h(handler, s, i)
			return
		}
		handler.Log.Errorf("Unhandled interaction for modal submits: %s", i.ModalSubmitData().CustomID)
	default:
		handler.Log.Errorf("Unhandled interaction type: %s", i.Type.String())
	}

}
