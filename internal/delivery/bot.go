package delivery

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rafaeltiribas/cardapio-uff/internal/usecase"
	"log"
)

func StartBot() {
	token := usecase.RetrieveToken("TELEGRAM_TOKEN")

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		switch update.Message.Command() {
		case "start":
			msg.Text = "Olá! Eu sou o bot do RU da UFF. Aqui estão algumas coisas que posso fazer:\n\n" +
				"🍽️ /cardapio - Veja o cardápio do dia\n" +
				"🕒 /horarios - Confira os horários de funcionamento\n" +
				"❓ /ajuda - Lista todos os comandos que eu entendo\n\n" +
				"Digite um comando para começar!"
			//adicionar o id do usuário em um bd para ter uma contagem de usuários do bot.
		case "ajuda":
			msg.Text = "Aqui estão os comandos que você pode usar:\n\n" +
				"🍽️ /cardapio - Veja o cardápio do dia\n" +
				"🕒 /horarios - Confira os horários de funcionamento\n" +
				"🏠 /start - Mostra a mensagem de boas-vindas e os comandos principais\n\n" +
				"Se precisar de mais alguma coisa, é só perguntar!"
		case "cardapio":
			msg.Text = Menu()
		case "horarios":
			msg.Text = "🕒 Horários de funcionamento do RU da UFF:\n\n" +
				"🍽️ Almoço: 11h30 - 14h00\n" +
				"🍲 Jantar: 17h30 - 19h30\n\n" +
				"Qualquer dúvida, é só usar o comando /ajuda!"
		default:
			msg.Text = "😕 Oops! Eu não reconheço esse comando.\n\n" +
				"Tente usar um dos comandos que eu conheço:\n" +
				"🍽️ /cardapio - Veja o cardápio do dia\n" +
				"🕒 /horarios - Confira os horários de funcionamento\n" +
				"❓ /ajuda - Lista todos os comandos disponíveis\n\n" +
				"Se precisar de ajuda, use /ajuda para ver a lista completa."
		}
		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}
