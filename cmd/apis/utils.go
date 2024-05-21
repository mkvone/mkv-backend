package apis

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"time"

	"github.com/mkvone/mkv-backend/cmd/config"
)

func convertConsensusPubkey(src ConsensusPubkey) config.ConsensusPubkey {
	return config.ConsensusPubkey{
		Type:  src.Type,
		Value: src.Key,
	}
}

func convertDescription(src Description) config.Description {
	return config.Description{
		Moniker:  src.Moniker,
		Identity: src.Identity,
		Website:  src.Website,
		Details:  src.Details,
	}
}

func convertCommission(src Commission) config.Commission {
	return config.Commission{
		Rate:          src.CommissionRates.Rate,
		MaxRate:       src.CommissionRates.MaxRate,
		MaxChangeRate: src.CommissionRates.MaxChangeRate,
	}
}
func extractRank(opAddr string, validators []Validator) int {

	sort.Slice(validators, func(i, j int) bool {
		tokensI, _ := strconv.ParseInt(validators[i].Tokens, 10, 64)
		tokensJ, _ := strconv.ParseInt(validators[j].Tokens, 10, 64)
		return tokensI > tokensJ
	})

	for rank, validator := range validators {
		if validator.OperatorAddress == opAddr {
			return rank + 1
		}
	}
	return -1
}

func updateSigningInfo(validator *config.Validator, info SigningInfo) {
	missedBlocksCounter, _ := strconv.Atoi(info.ValSigningInfo.MissedBlocks)
	validator.Uptime.MissedBlock = missedBlocksCounter
	percent := (100 - (float32(missedBlocksCounter) / float32(validator.Uptime.TotalBlock)))
	validator.Uptime.Percent = percent
}

func updateStakingInfo(validator *config.Validator, Validator Validator) {
	votingPower, _ := strconv.Atoi(Validator.Tokens)
	validator.VotingPower = votingPower / 1000000
	validator.OperatorAddr = Validator.OperatorAddress
	validator.ConsensusPubkey = convertConsensusPubkey(Validator.ConsensusPubkey)
	validator.Jailed = Validator.Jailed
	validator.Status = Validator.Status
	validator.Tokens = Validator.Tokens
	validator.DelegatorShares = Validator.DelegatorShares
	validator.Description = convertDescription(Validator.Description)
	validator.Commission = convertCommission(Validator.Commission)
}

func updateNodeInfo(chain *config.ChainConfig, nodeInfo NodeInfo) {
	chain.ChainID = nodeInfo.NodeInfo.Network
	chain.GoVersion = nodeInfo.ApplicationVersion.GoVersion
	chain.Version = nodeInfo.ApplicationVersion.Version
	chain.Appname = nodeInfo.ApplicationVersion.Appname
	chain.CosmosSDK = nodeInfo.ApplicationVersion.CosmosSDKVersion
}
func timeAgo(timeStr string) (string, error) {
	// Parse the time string assuming the format "01/02/2006 03:04:05 PM -0700"
	// Adjust the format based on your actual input
	t, err := time.Parse("01/02/2006 03:04:05 PM -07:00", timeStr)
	if err != nil {
		log.Println(err)
		return "", err
	}

	// Get current time
	now := time.Now()

	// Calculate duration
	duration := now.Sub(t)

	if duration.Hours() >= 24 {
		days := int(duration.Hours() / 24)
		return fmt.Sprintf("%d days ago", days), nil
	} else if duration.Hours() >= 1 {
		hours := int(duration.Hours())
		return fmt.Sprintf("%d hours ago", hours), nil
	} else {
		mins := int(duration.Minutes())
		return fmt.Sprintf("%d mins ago", mins), nil
	}
}
