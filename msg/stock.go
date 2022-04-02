package msg

import (
	"fmt"
	"strings"

	"github.com/cbrissonCoveo/CoveStonks/api"
	"github.com/pterm/pterm"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func StockPrice(price float64, priceState bool) string {
	var priceColor pterm.Letters
	priceStr := fmt.Sprintf("%.2f", price)

	if priceState {
		priceColor = pterm.NewLettersFromStringWithStyle(
			priceStr,
			pterm.NewStyle(pterm.FgGreen),
		)

	} else {
		priceColor = pterm.NewLettersFromStringWithStyle(
			priceStr,
			pterm.NewStyle(pterm.BgLightRed),
		)
	}

	priceAsFinalStr, _ := pterm.DefaultBigText.WithLetters(priceColor).Srender()

	return pterm.DefaultCenter.Sprint(priceAsFinalStr)
}

func template(name string, value string) string {
	const MAX_LENGTH = 42
	pad := MAX_LENGTH - len(name)

	// Example output
	// [ Percentage                     0.35% ]
	return fmt.Sprintf("[ %s %*s  ]", name, pad, value)
}

func Statistics(stock api.QuoteStructure, shares float64) string {
	out := []string{
		template("Name", stock.Name()),
		template("Market", stock.State()),
		template("Percentage", stock.MarketChangePercent()),
		template("Price", stock.MarketChange()),
	}

	if stock.MarketVolume() > 0 {
		p := message.NewPrinter(language.English)

		out = append(out, template("Volume", p.Sprintf("%d", stock.MarketVolume())))
	}

	if shares > 0 {
		out = append(
			out,
			template("Your position", fmt.Sprintf("%.2f %s", shares*stock.Price(), stock.StockCurrency())),
		)
	}

	return pterm.DefaultCenter.Sprint(strings.Join(out, "\n"))
}
