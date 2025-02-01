package handlers

import (
	"context"
	"gffbot/internal/game"
	"gffbot/internal/text"
	"log"
	"strings"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/go-telegram/ui/keyboard/inline"
)




func DefaultHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message == nil {
		return
	}
	
	index, exists := findUserInData(game.User{ChatID: update.Message.Chat.ID})
	if exists && users[index].SendingKey {
		u := users[index]

		if u.LobbyID != 0 || u.LobbyKey != "" {
			u.SendMessage(ctx, b, text.AlreadyInLobbyF, u.LobbyKey)
			return
		}

		key := update.Message.Text

		if lobby, exists := lobbies[key]; exists {

			if lobby.IsStarted {
				u.SendMessage(ctx, b, text.LobbyGameIsStarted)
				return
			}

			// Поменять

			memebersList := []string{u.Name}

			for _, memeber := range lobby.Members {
				memebersList = append(memebersList, memeber.Name)
			}

			newList := strings.Join(memebersList, "\n")

			// ==//==

			for _, member := range lobby.Members {
				member.SendMessage(ctx, b, text.PlayerJoinedLobbyF, u.Name, newList)
			}

			u.SendingKey = false
			u.LobbyKey = key
			u.LobbyID = lobby.ID
			users[index] = u

			lobby.Members = append(lobby.Members, u)

			lobbies[key] = lobby

			u.SendMessage(ctx, b, text.PlayerJoinedLobbyF, text.PlayerJoinedLobbyF, u.GetText(text.You), newList)
		} else {
			u.SendMessage(ctx, b, text.LobbyNotExists)
			return
		}

		//		users[index] = u

	} else {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   text.GetConvertToLang(update.Message.From.LanguageCode, text.UnknownCommand),
		})
	}
}

func StartHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	kb := inline.New(b).
		Row().
		Button(text.GetConvertToLang(update.Message.From.LanguageCode, text.Join), []byte("1"), onJoinLobbySelect).
		Row().
		Button(text.GetConvertToLang(update.Message.From.LanguageCode, text.Create), []byte("2"), onCreateLobbySelect)

	newUser := game.User{
		ChatID:	update.Message.Chat.ID,
		Name:	bot.EscapeMarkdown(update.Message.From.FirstName) + " " +
					bot.EscapeMarkdown(update.Message.From.LastName),
		Lang:	update.Message.From.LanguageCode,
	}

	users = append(users, newUser)

	log.Printf("New user: {Name: %s, ChatID: %d} is added", newUser.Name, newUser.ChatID)

	newUser.SendReplayMarkup(ctx, b, kb, text.StartCommandF, newUser.Name)
}

func GameStartHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	index, exists := findUserInData(game.User{ChatID: update.Message.Chat.ID})
	if !exists {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   text.GetConvertToLang(update.Message.From.LanguageCode, text.SomethingWentWrong),
		})
		return
	}

	lobby, exists := lobbies[users[index].LobbyKey]
	if !exists {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   text.GetConvertToLang(update.Message.From.LanguageCode, text.SomethingWentWrong),
		})
		return
	}

	if lobby.GameType == text.GameNotSelected {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   text.GetConvertToLang(update.Message.From.LanguageCode, text.CantStartGame),
		})
		return
	}

	if len(lobby.Members) < game.MINIMUM_MEMBERS_FOR_MAFIA {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   text.GetConvertToLang(update.Message.From.LanguageCode, text.AtLeastMembersF, game.MINIMUM_MEMBERS_FOR_MAFIA),
		})
		return
	}

	for _, member := range lobby.Members {
		member.SendMessage(ctx, b, text.GameStarted)
	}

	lobby.IsStarted = true

	lobbies[users[index].LobbyKey] = lobby

	// go lobby.StartGame(ctx, b) ???

	lobby.StartGame(ctx, b)
}