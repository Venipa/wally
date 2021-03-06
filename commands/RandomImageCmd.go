package commands

import (
	"errors"
	"github.com/bwmarrin/discordgo"
	"github.com/toyz/wally/wallhaven"
	"math/rand"
	"time"
)

func init() {
	// Random General,Anime,Person
	Register("!r", randomImage(""))
	// Random General
	Register("!rg", randomImage("100"))
	// Random Anime
	Register("!ra", randomImage("010"))
	// Random Person
	Register("!rp", randomImage("001"))
}

func randomImage(category string) CommandFunc {
	return func(s *discordgo.Session, m *discordgo.MessageCreate, args []string) error {
		resolution := ""
		if len(args) > 0 {
			resolution = args[0]
		}

		papers, err := getWallPapers(category, resolution)
		if err != nil {
			return err
		}

		paper := randomWallpaper(papers)
		paper, err = wallhaven.SingleImage(paper.ID)
		if err != nil {
			return err
		}

		_, err = s.ChannelMessageSendEmbed(m.ChannelID, paper.CreateEmbed())
		return err
	}
}

// 111 (general/anime/people)
func getWallPapers(category string, resolution string) ([]wallhaven.Wallpaper, error) {
	if category == "" {
		category = "111"
	}
	papers, err := wallhaven.RandomImage(category, resolution)
	if err != nil {
		return nil, err
	}
	if len(papers) == 0 {
		return nil, errors.New("no wallpapers found")
	}

	return papers, nil
}

func randomWallpaper(papers []wallhaven.Wallpaper) wallhaven.Wallpaper {
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s) // initialize local pseudorandom generator

	return papers[r.Intn(len(papers))]
}