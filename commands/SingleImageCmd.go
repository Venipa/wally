package commands

import (
	"errors"
	"github.com/bwmarrin/discordgo"
	"github.com/toyz/wally/wallhaven"
)

func init() {
	Register("!v", singleImage)
}

func singleImage(s *discordgo.Session, m *discordgo.MessageCreate, args []string) error {
	if len(args) == 0 {
		return errors.New("invalid usage `!v <wallpaper id>` (example: `!v 94x38z`)")
	}

	image, err := wallhaven.SingleImage(args[0])
	if err != nil {
		return err
	}

	_, err = s.ChannelMessageSendEmbed(m.ChannelID, image.CreateEmbed())
	return err
}