package handlers

import (
	"github.com/bwmarrin/discordgo"
)

func (h *Handler) InteractionCreate(s *discordgo.Session, i *discordgo.InteractionCreate) {
	switch i.Type {
	case discordgo.InteractionApplicationCommand:
		if h, ok := commandsHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
			return
		}
		h.Log.Errorf("Unhandled interaction for commands: %s", i.ApplicationCommandData().Name)
	case discordgo.InteractionMessageComponent:
		if h, ok := componentsHandlers[i.MessageComponentData().CustomID]; ok {
			h(s, i)
			return
		}
		h.Log.Errorf("Unhandled interaction for components: %s", i.MessageComponentData().CustomID)
	case discordgo.InteractionModalSubmit:
		if h, ok := componentsHandlers[i.ModalSubmitData().CustomID]; ok {
			h(s, i)
			return
		}
		h.Log.Errorf("Unhandled interaction for modal submits: %s", i.ModalSubmitData().CustomID)
	default:
		h.Log.Errorf("Unhandled interaction type: %s", i.Type.String())
	}

}
