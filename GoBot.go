package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func checkWebsite(website string) (bool, time.Duration, error) {
	start := time.Now()
	resp, err := http.Get(website)
	if err != nil {
		return false, 0, err
	}
	defer resp.Body.Close()

	duration := time.Since(start)
	return resp.StatusCode == 200, duration, nil
}

func sendTelegramMessage(botToken string, chatID string, message string) error {
	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", botToken)
	data := url.Values{}
	data.Set("chat_id", chatID)
	data.Set("text", message)

	resp, err := http.PostForm(apiURL, data)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("Telegram mesajı gönderilemedi, durum kodu: %d", resp.StatusCode)
	}
	return nil
}

func formatURL(website string) string {

	parts := strings.Split(website, ".")
	return fmt.Sprintf("%s[.]%s[.]%s", parts[0], parts[1], parts[2])
}

func monitorWebsites(websites []string, botToken string, chatID string) {
	for {
		for _, website := range websites {
			isUp, responseTime, err := checkWebsite(website)
			formattedURL := formatURL(website) // Site ismini formatla
			if err != nil {
				sendTelegramMessage(botToken, chatID, fmt.Sprintf("Hata: %v, Site: %s", err, formattedURL))
			} else if !isUp {
				sendTelegramMessage(botToken, chatID, fmt.Sprintf("Site erişilemez durumda: %s", formattedURL))
			} else {
				if responseTime > 100*time.Millisecond {
					sendTelegramMessage(botToken, chatID, fmt.Sprintf("Site yavaş abi,  %s,  yanıt süresi: %v", formattedURL, responseTime))
				} else {
					sendTelegramMessage(botToken, chatID, fmt.Sprintf("Site ayakta abi, %s,  yanıt süresi: %v", formattedURL, responseTime))
				}
			}
		}
		time.Sleep(5 * time.Minute)
	}
}

func main() {
	websites := []string{
		"<URL1>",
		"<URL2>",
		"<URL3>",
	}
	botToken := "<yourBotToken>"
	chatID := "<yourChatId>"

	go monitorWebsites(websites, botToken, chatID)

	select {}
}
