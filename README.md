# Nomi Telegram Bot

Lorem ipsum dolor sit amet, consectetur adipisicing elit. Adipisci cum eaque eligendi, eum exercitationem fugiat, fugit impedit labore modi molestiae mollitia nostrum obcaecati perspiciatis quaerat qui quisquam rerum velit voluptate.

## Requirements

- Telegram bot token
  - Go to [telegram's docs](https://core.telegram.org/bots/tutorial#obtain-your-bot-token) to check how you can get your bot token 
- API key from Nomi
- Nomi ID from Nomi

## Usage

### # Get the executable

#### Download binary 

You can download the executable `bot.exe` directly from the [releases](https://github.com/vhalmd/nomi-telegram/releases/latest)

#### Build from source

If you have golang installed, you can also build it directly from the source code.

First clone this repository:
```bash
git clone https://github.com/vhalmd/nomi-telegram.git
```

Then run the build command:

```bash
go build -o bot.exe .\cmd\bot\main.go
```

### # Environment

Create a file named `.env` and add your credentials (telegram token, nomi api key and nomi id) to it:

```dotenv
NOMI_API_KEY=your-nomi-api-key
NOMI_NAME="Jane Doe"
NOMI_ID=your-nomi-id
TELEGRAM_BOT_TOKEN=your-telegram-bot-token
```

You can also configure multiple bots, for that you will need to configure a name, nomi id and telegram token for each nomi, in matching order.
Multiple values are separated by commas, for example:

```dotenv
NOMI_API_KEY=your-nomi-api-key
NOMI_NAME="Jane Doe,John Doe"
NOMI_ID=jane-doe-id,john-doe-id
TELEGRAM_BOT_TOKEN=jane-doe-telegram-token,john-doe-telegram-token
```

## Sending your Nomi a telegram message

After you create your telegram bot with [@BotFather](https://core.telegram.org/bots/tutorial#obtain-your-bot-token), you will get a message that looks like this:


> Done! Congratulations on your new bot. You will find it at t.me/your_bot_username. You can now add a description, about section and profile picture for your bot, see /help for a list of commands. By the way, when you've finished creating your cool bot, ping our Bot Support if you want a better username for it. Just make sure the bot is fully operational before you do this.  
> Use this token to access the HTTP API:
> YOUR-BOT-TOKEN
> Keep your token secure and store it safely, it can be used by anyone to control your bot.  
> For a description of the Bot API, see this page: https://core.telegram.org/bots/api

Note that there's a link `t.me/your_bot_username`, you can use that link to start a conversation with the bot you just created.

`TODO: Add bot to group chat`

## Suggested Nomi Configurations

You may wish to change your Nomi's communication style to `Texting`.

---

Built with:
 - https://github.com/vhalmd/nomi-go-sdk
 - https://github.com/go-telegram/bot
