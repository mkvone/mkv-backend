package apis

import (
	"strconv"

	"github.com/mkvone/mkv-backend/cmd/config"
	"github.com/mkvone/mkv-backend/cmd/utils"
)

func refresh_Validator_Signing_Status(cfg *[]config.ChainConfig) {
	for i, _ := range *cfg {
		go func(chain *config.ChainConfig, validator *config.Validator) {
			if !validator.Enable {
				l("❓ Skipping validator update for chain name : ", chain.Name)
				return
			}
			apiBaseURL := chain.Endpoints.Api

			// Fetch signing info
			consAddr, _ := utils.EncodeValidatorAddress(validator.ConsensusPubkey.Value, validator.OperatorAddr)
			validator.ValconAddress = consAddr
			signingInfoURL := apiBaseURL + "/cosmos/slashing/v1beta1/signing_infos/" + consAddr

			var signingInfo SigningInfo
			if err := fetchDataAndHandleError(signingInfoURL, &signingInfo); err != nil {
				l("❌ Error fetching", chain.Name, "'s signing info", err)
				return
			}
			updateSigningInfo(validator, signingInfo)
			l("✅ Signing info updated chain name : ", chain.Name)
		}(&(*cfg)[i], &(*cfg)[i].Validator)
	}
}

func ParsingNodeInfo(cfg *[]config.ChainConfig) {
	for i := range *cfg {
		go func(chain *config.ChainConfig, nodeInfo NodeInfo) {
			apiBaseURL := chain.Endpoints.Api
			// Fetch and update node info
			node_infoParamsURL := apiBaseURL + "/cosmos/base/tendermint/v1beta1/node_info"
			if err := fetchDataAndHandleError(node_infoParamsURL, &nodeInfo); err != nil {
				l("❌ Error fetching  ", chain.Name, "'s node info", err)
				return
			}
			updateNodeInfo(chain, nodeInfo)
			l("✅ Node info updated chain name : ", chain.Name)

		}(&(*cfg)[i], NodeInfo{})
	}
}
func daily_Validator_and_Chain_Updates(cfg *[]config.ChainConfig) {
	for i := range *cfg {
		go func(chain *config.ChainConfig, validator *config.Validator) {
			apiBaseURL := chain.Endpoints.Api
			// Fetch and update node info
			// node_infoParamsURL := apiBaseURL + "/cosmos/base/tendermint/v1beta1/node_info"
			// var nodeInfo NodeInfo
			// if err := fetchDataAndHandleError(node_infoParamsURL, &nodeInfo); err != nil {
			// 	l("❌ Error fetching  ", chain.Name, "'s node info", err)
			// 	return
			// }
			// updateNodeInfo(chain, nodeInfo)
			// l("✅ Node info updated chain name : ", chain.Name)
			if !validator.Enable {
				l("❓ Skipping validator update for chain name : ", chain.Name)
				return
			}

			slashingParamsURL := apiBaseURL + "/cosmos/slashing/v1beta1/params"
			var slashingParams SlashingParams
			if err := fetchDataAndHandleError(slashingParamsURL, &slashingParams); err != nil {
				l("❌ Error fetching ", chain.Name, "'s slashing params", err)
				return
			}

			if signedBlocksWindow, err := strconv.Atoi(slashingParams.Params.SignedBlocksWindow); err == nil {
				validator.Uptime.TotalBlock = signedBlocksWindow
				l("✅ SignedBlocksWindow updated chain name : ", chain.Name)
			} else {
				l(err)
				return
			}

			// Fetch validators info
			stakingURL := apiBaseURL + "/cosmos/staking/v1beta1/validators/" + validator.OperatorAddr
			var validatorset ValidatorSet
			if err := fetchDataAndHandleError(stakingURL, &validatorset); err != nil {
				l("❌ Error fetching ", chain.Name, "'s validator info", err)
				return
			}
			updateStakingInfo(validator, validatorset.Validator)
			l("✅ Validator info updated chain name : ", chain.Name)
			// Fetch delegation info
			delegationURL := stakingURL + "/delegations"
			var totalDelegations DelegationCounts
			if err := fetchDataAndHandleError(delegationURL, &totalDelegations); err != nil {
				l("❌ Error fetching ", chain.Name, "'s delegation info", err)
				return
			}

			validator.TotalDelegationCounts = totalDelegations.Pagination.Total

			stakingURL = apiBaseURL + "/cosmos/staking/v1beta1/validators?pagination.limit=500&status=BOND_STATUS_BONDED"
			if err := fetchDataAndHandleError(stakingURL, &validatorset); err != nil {
				l("❌ Error fetching ", chain.Name, "'s validators info", err)
				return
			}
			validator.Rank = extractRank(validator.OperatorAddr, validatorset.Validators)
			l("✅ Rank updated chain name : ", chain.Name)
		}(&(*cfg)[i], &(*cfg)[i].Validator)
	}
}
