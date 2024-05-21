package cmd

import (
	"sync"

	"github.com/mkvone/mkv-backend/cmd/apis"
	"github.com/mkvone/mkv-backend/cmd/config"
	"github.com/mkvone/mkv-backend/cmd/db"
	"github.com/mkvone/mkv-backend/cmd/server"
	"github.com/robfig/cron"
)

var mutex = &sync.Mutex{}

func handleBlockInfo(cfg *[]config.ChainConfig) {
	for bi := range blockInfoChan {
		mutex.Lock()
		for i := range *cfg {
			if (*cfg)[i].Name == bi.ChainName {
				if bi.Height == "" {
					continue
				}
				(*cfg)[i].Block.Height = bi.Height
				(*cfg)[i].Block.Time = bi.Time
				l("✅ Block info updated chain name : ", bi.ChainName, " Height : ", bi.Height)
			}
		}
		mutex.Unlock()
	}
}
func cronjob(cfg *[]config.ChainConfig) {
	apiManager := apis.APIManager{Config: cfg}
	apiManager.Init()
	c := cron.New()

	c.AddFunc("@every 5s", func() {
		apiManager.UpdateEvery5Seconds()
	})
	c.AddFunc("@every 1m", func() {
		apiManager.UpdateEvery1Min()
	})
	c.AddFunc("@every 5m", func() {
		apiManager.UpdateEvery5Min()
	})
	c.AddFunc("@every 1d", func() {
		apiManager.UpdateEveryDay()
	})
	c.Start()

	defer c.Stop()
	select {}

}
func Run(cfg *config.Config) {

	db.ConnectToMongoDB(cfg.Databases.MongoDB.URL)
	go cronjob(&cfg.Chains)
	go handleBlockInfo(&cfg.Chains)

	// ws current block height
	for i := range cfg.Chains {
		chain := &cfg.Chains[i]
		go SubscribeToNewBlocks(*chain)
	}

	// swagger
	if cfg.Swagger.Enable {
		server.Serve(cfg)
	}

	select {}
}
