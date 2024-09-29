# WebSentry: Website Uptime and Performance Monitoring

WebSentry is a simple, lightweight Go application that monitors the uptime and performance of specified websites. It sends notifications to a Telegram chat when a website is down or responding slowly.

## Features
- Monitors multiple websites for uptime.
- Measures response time for each website.
- Sends notifications via Telegram when a site is down or slow.
- Configurable interval for checking websites.
- Easy to deploy and modify.

## How It Works
1. The program takes a list of websites as input.
2. It periodically checks the status and response time of each website.
3. If a website is down or responding too slowly, a message is sent to a specified Telegram chat.
4. The messages are sent using the Telegram Bot API, and notifications include information on the website status and response time.

## Installation

### Prerequisites
- Go 1.16+ installed on your machine.
- A Telegram bot token.
- The chat ID of the Telegram chat where you want to receive notifications.

### Clone the Repository
```bash
git clone https://github.com/CuneytKose/websentry.git
cd websentry
 
