package apis

import (
	"github.com/mkvone/mkv-backend/cmd/config"
	"github.com/mkvone/mkv-backend/cmd/db"
)

func updateSymbolPrice(cfg *[]config.ChainConfig) {
	for i := range *cfg {
		go func(chain *config.ChainConfig, symbol *config.Symbol) {
			coin := db.FindLatestCoinPrice(symbol.Ticker)
			if coin != nil {
				chain.Symbol.Price = float32(coin.USD)
			}

		}(&(*cfg)[i], &(*cfg)[i].Symbol)
	}
}
