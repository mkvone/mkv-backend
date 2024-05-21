package apis

import "github.com/mkvone/mkv-backend/cmd/config"

type APIManager struct {
	Config *[]config.ChainConfig
}

func (m *APIManager) Init() {
	m.UpdateEvery5Seconds()
	m.UpdateEveryDay()
	m.UpdateEvery1Min()
	m.UpdateEvery5Min()

}

func (m *APIManager) UpdateEvery5Seconds() {
	refresh_Validator_Signing_Status(m.Config)
}

func (m *APIManager) UpdateEveryDay() {
	daily_Validator_and_Chain_Updates(m.Config)
}

func (m *APIManager) UpdateEvery1Min() {
	updateSnapshotInfo(m.Config)
}
func (m *APIManager) UpdateEvery5Min() {
	updateSymbolPrice(m.Config)
}
