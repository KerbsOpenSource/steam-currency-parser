# Steam currency parser
Automatic sending to telegeram. Example: [t.me/SteamCurrencyUpdate](https://t.me/s/SteamCurrencyUpdate)
## 🤔 What is this?
This app is used to check real-time currency adjustments for Steam and has a complete currency database. I don't see any other way to use it except notification in telegram, as it's not a stock exchange, it's still an open source project and you can do what you want.


At the time of writing the project I did not find any alternatives in open source. They read adjustments from third-party services or did it incorrectly calculating currency adjustments based on the cost of games.

>**My approach is based on real time adjustments based on the prices of exposed lots of items on steam market.
This approach makes it possible to check the whole currency base of Steam in real time.**

### Default currencies
+ EUR
+ CNY
+ JPY
+ AED
+ RUB
> The USD currency is used for validity checks, as steam uses USD as the anchor currency.

## 🏃‍♂️ How to run
### Use the startup flags to customize
> + ``-token`` — Telegram bot token ([@BotFather](https://telegram.me/BotFather))
> + ``-chatid`` — Telegram chat id ([@getmyid_bot](https://t.me/getmyid_bot))
> + ``-appid`` — Steam app id (Dota 2 - "570" СS2 - "730" TF2 - "440")
> + ``-hashname`` — Steam market item hash name (Lot name in URI encode)
#### example: `main -token "telegrambot_token" -chatid "1234567" -appid "570" -hashname "Ruling%20Staff%20of%20the%20Priest%20Kings"`

#### You can use the cron daemon to run at a specified time. For example, every hour. But it would be better to make it run from a bash file.

##### The currency base can be selected inside the main file. (It's a сringe, I know.)

## 🔨 How to compile
```
git clone https://github.com/KerbsOpenSource/steam-currency-parser
cd steam-currency-parser/cmd
go build -o steam-currency-parser main.go
```
