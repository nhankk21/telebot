package main

import (
    "log"
    "net/http"
    "encoding/json"

    tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
    // Create a new bot instance
    bot, err := tgbotapi.NewBotAPI("6159207883:AAHtZrwMbA41qQK1zuXpYs-8iaZj-L1M7OE")
    if err != nil {
        log.Fatal(err)
    }

    // Set up a webhook handler function
    http.HandleFunc("/webhook", func(w http.ResponseWriter, r *http.Request) {
        // Decode the incoming update
        var update tgbotapi.Update
        if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
            log.Println("Failed to decode update:", err)
            return
        }

        // Process the update and generate a response
        if update.Message != nil {
            message := update.Message
            response := "Hello, " + message.From.FirstName + "!"
            reply := tgbotapi.NewMessage(message.Chat.ID, response)
            _, err := bot.Send(reply)
            if err != nil {
                log.Println("Failed to send message:", err)
                return
            }
        }
    })

    // Start the server with HTTPS
    log.Println("Starting server...")
    if err := http.ListenAndServeTLS(":8443", "server.crt", "server.key", nil); err != nil {
        log.Fatal(err)
    }
}
