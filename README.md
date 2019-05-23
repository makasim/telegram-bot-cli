# telegram-bot-cli

A cli command to send message directly from the bash. 
Could be used on CI for sending build or deployment updates.  

```
TELEGRAM_TOKEN=telegramToken go run main.go -- userIdOrChannelId aMessage 
```

Set parse mode to `HTML` to send messages with inline links:

```
TELEGRAM_TOKEN=telegramToken go run main.go --parse-mode=HTML -- userIdOrChannelId aMessage
```

