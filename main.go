package main

import (
	"flag"
	"fmt"
	"log"
	"steam-currency-parser-telegram/regex"
	"steam-currency-parser-telegram/steam"
	"steam-currency-parser-telegram/telegram"
	"strconv"
	"time"
)

func mustFlags() (tgBotToken, chatID, appID, marketHashName string, priceCorrertor int) {
	flag.StringVar(&tgBotToken, "token", "", "Telegram bot token")
	flag.StringVar(&chatID, "chatid", "", "Telegram chat id")
	flag.StringVar(&appID, "appid", "", "Steam appID")
	flag.StringVar(&marketHashName, "hashname", "", "Steam market item hash name")
	flag.IntVar(&priceCorrertor, "pricecor", 1000, "Price corrector to get 1 unite price")

	flag.Parse()
	if tgBotToken == "" {
		log.Fatal("Telegram token not found")
	}
	if chatID == "" {
		log.Fatal("Telegram chat id not found flags")
	}

	if appID == "" {
		log.Fatal("Steam app not found flags")
	}

	if marketHashName == "" {
		log.Fatal("Steam market hash name not found flags")
	}
	return
}

func currencyDictionary() map[string]string {
	dictionary := make(map[string]string)

	// dictionary["USD"] = "1" // United States Dollar
	// dictionary["GBP"] = "2"  // United Kingdom Pound
	dictionary["EUR"] = "3" // European Union Euro
	// dictionary["CHF"] = "4"  // Swiss Francs
	dictionary["RUB"] = "5" // Russian Rouble
	// dictionary["PLN"] = "6"  // Polish Złoty
	// dictionary["BRL"] = "7"  // Brazilian Reals
	dictionary["JPY"] = "8" // Japanese Yen
	// dictionary["NOK"] = "9"  // Norwegian Krone
	// dictionary["IDR"] = "10" // Indonesian Rupiah
	// dictionary["MYR"] = "11" // Malaysian Ringgit
	// dictionary["PHP"] = "12" // Philippine Peso
	// dictionary["SGD"] = "13" // Singapore Dollar
	// dictionary["THB"] = "14" // Thai Baht
	// dictionary["VND"] = "15" // Vietnamese Dong
	// dictionary["KRW"] = "16" // South Korean Won
	// dictionary["TRY"] = "17" // Turkish Lira
	// dictionary["UAH"] = "18" // Ukrainian Hryvnia
	// dictionary["MXN"] = "19" // Mexican Peso
	// dictionary["CAD"] = "20" // Canadian Dollars
	// dictionary["AUD"] = "21" // Australian Dollars
	// dictionary["NZD"] = "22" // New Zealand Dollar
	dictionary["CNY"] = "23" // Chinese Renminbi (yuan)
	// dictionary["INR"] = "24" // Indian Rupee
	// dictionary["CLP"] = "25" // Chilean Peso
	// dictionary["PEN"] = "26" // Peruvian Sol
	// dictionary["COP"] = "27" // Colombian Peso
	// dictionary["ZAR"] = "28" // South African Rand
	// dictionary["HKD"] = "29" // Hong Kong Dollar
	// dictionary["TWD"] = "30" // New Taiwan Dollar
	// dictionary["SAR"] = "31" // Saudi Riyal
	dictionary["AED"] = "32" // United Arab Emirates Dirham
	// dictionary["ARS"] = "34" // Argentine Peso
	// dictionary["ILS"] = "35" // Israeli New Shekel
	// dictionary["KZT"] = "37" // Kazakhstani Tenge
	// dictionary["KWD"] = "38" // Kuwaiti Dinar
	// dictionary["QAR"] = "39" // Qatari Riyal
	// dictionary["CRC"] = "40" // Costa Rican Colón
	// dictionary["UYU"] = "41" // Uruguayan Peso

	// I'm dead inside

	return dictionary
}

func textMessageСurrency(currencies map[string]float64) string {
	textMessage := fmt.Sprintf("🇪🇺 EUR: %.2f€\n🇨🇳 CNY: ¥ %.2f\n🇯🇵 JPY: ¥ %.2f\n🇦🇪 AED: د %.2f\n🇷🇺 RUB: %.2f ₽", currencies["EUR"], currencies["CNY"], currencies["JPY"], currencies["AED"], currencies["RUB"])
	return textMessage
}

// You need to find an item on the steam market that will serve as an item for a comparable price
// You can write down how much you need to divide by to get values ​​for 1 unit
// Example priceItem = 200000 USD, priceCorrertor = 1000, unitPrice = (200000/100) / (1000/100) = 200 = 2 / 100  = 2.00 USD
// priceCorrertor := 1000.00
func main() {
	tgBotToken, chatID, appID, marketHashName, priceCorrertor := mustFlags()
	currencies := currencyDictionary()
	resultCurrencies := make(map[string]float64)
	for key, value := range currencies {
		time.Sleep(5 * time.Second)
		marketPrice := steam.LowestPrice(appID, value, marketHashName)
		priceInt, err := strconv.Atoi(regex.OnlyInt(marketPrice))
		if err != nil {
			log.Fatal(err)
			return
		}
		resultCurrencies[key] = float64(priceInt/priceCorrertor) / 100
	}
	textMessage := textMessageСurrency(resultCurrencies)
	telegram.SendMessage(textMessage, tgBotToken, chatID)
}
