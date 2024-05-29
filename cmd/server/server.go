package server

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	_ "github.com/mkvone/mkv-backend/docs"
	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/mkvone/mkv-backend/cmd/config"
)

//	@title			Swagger Example API
//	@version		1.0
//	@description	This is a sample server celler server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

// @license.name	Apache 2.0
// @license.url	http://www.apache.org/licenses/LICENSE-2.0.html
func Serve(cfg *config.Config) {
	e := echo.New()
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/swagger", func(c echo.Context) error {
		return c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})

	e.GET("/chains", getChains(cfg))
	e.GET("/validators", getValidators(cfg))
	e.GET("/validator/:chain", getValidator(cfg))
	e.GET("/stats", getStats(cfg))
	e.GET("/endpoints", getEndpoints(cfg))
	e.GET("/snapshots", getSnapshots(cfg))
	e.GET("/snapshot/:chain", getSnapshot(cfg))
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.Swagger.Port)))
}

// GetChains godoc
// @Summary Get chains
// @Description Get chains
// @Tags chains
// @Accept json
// @Produce json
// @Success 200 {object} APIResponse{data=[]config.ChainConfig}
// @Router /chains [get]
func getChains(cfg *config.Config) echo.HandlerFunc {
	// return func(c echo.Context) error {
	// 	return c.JSON(http.StatusOK, cfg.Chains)
	// }
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, APIResponse{
			Message: "Success",
			Code:    http.StatusOK,
			Data:    cfg.Chains,
		})
	}
}

// GetValidators godoc
// @Summary Get validators
// @Description Get validators
// @Tags validators
// @Accept json
// @Produce json
// @Success 200 {object} APIResponse{data=[]ValidatorResponse}
// @Router /validators [get]
func getValidators(cfg *config.Config) echo.HandlerFunc {
	// return func(c echo.Context) error {
	// 	return c.JSON(http.StatusOK, cfg.Chains)
	// }

	return func(c echo.Context) error {
		var validators []ValidatorResponse
		for _, chain := range cfg.Chains {
			if !chain.Validator.Enable {
				continue
			}
			validators = append(validators, ValidatorResponse{
				Name:        chain.Name,
				ChainID:     chain.ChainID,
				Path:        chain.Path,
				Image:       chain.ImgURL,
				BlockHeight: chain.Block.Height,
				BlockTime:   chain.Block.Time,
				Price:       chain.Symbol.Price,
				Ticker:      chain.Symbol.Ticker,
				Validator: Validator{
					OperatorAddr:  chain.Validator.OperatorAddr,
					WalletAddress: chain.Validator.WalletAddress,
					ValconAddress: chain.Validator.ValconAddress,
					VotingPower:   chain.Validator.VotingPower,
					Uptime: Uptime{
						Percent:     chain.Validator.Uptime.Percent,
						MissedBlock: chain.Validator.Uptime.MissedBlock,
						TotalBlock:  chain.Validator.Uptime.TotalBlock,
						Tombstoned:  chain.Validator.Uptime.Tombstoned,
					},
					Rank:            chain.Validator.Rank,
					Jailed:          chain.Validator.Jailed,
					Status:          chain.Validator.Status,
					Tokens:          chain.Validator.Tokens,
					DelegatorShares: chain.Validator.DelegatorShares,
					Description: Description{
						Moniker:  chain.Validator.Description.Moniker,
						Identity: chain.Validator.Description.Identity,
						Website:  chain.Validator.Description.Website,
						Details:  chain.Validator.Description.Details,
					},
					Commission: Commission{
						Rate:          chain.Validator.Commission.Rate,
						MaxRate:       chain.Validator.Commission.MaxRate,
						MaxChangeRate: chain.Validator.Commission.MaxChangeRate,
					},
					TotalDelegationCounts: chain.Validator.TotalDelegationCounts,
				},
			})
		}

		return c.JSON(http.StatusOK, APIResponse{
			Message: "Success",
			Code:    http.StatusOK,
			Data:    validators,
		})
	}
}

// GetValidator godoc
// @Summary Get validator
// @Description Get validator
// @Tags validators
// @Accept json
// @Produce json
// @Param chain path string true "Chain ID"
// @Success 200 {object} APIResponse{data=Validator}
// @Router /validator/{chain} [get]
func getValidator(cfg *config.Config) echo.HandlerFunc {
	return func(c echo.Context) error {
		path := c.Param("chain")
		for _, chain := range cfg.Chains {
			if chain.Path == path {
				if !chain.Validator.Enable {
					return c.JSON(http.StatusOK, APIResponse{
						Message: "Success",
						Code:    http.StatusOK,
						Data:    nil,
					})
				}
				return c.JSON(http.StatusOK, APIResponse{
					Message: "Success",
					Code:    http.StatusOK,
					Data:    chain,
				})
			}
		}
		return c.JSON(http.StatusOK, APIResponse{
			Message: "Success",
			Code:    http.StatusOK,
			Data:    nil,
		})
	}
}

// GetStats godoc
// @Summary Get stats
// @Description Get stats
// @Tags stats
// @Accept json
// @Produce json
// @Success 200 {object} APIResponse{data=Stats}
// @Router /stats [get]
func getStats(cfg *config.Config) echo.HandlerFunc {
	return func(c echo.Context) error {
		var stats Stats
		count := 0
		staked := float32(0)
		uptime := float32(0)
		daligators := 0
		for _, chain := range cfg.Chains {
			if !chain.Validator.Enable {
				continue
			}
			count++
			staked += float32(chain.Validator.VotingPower) * chain.Symbol.Price
			uptime += chain.Validator.Uptime.Percent
			d, _ := strconv.Atoi(chain.Validator.TotalDelegationCounts)
			daligators += d
		}
		stats = Stats{
			Uptime:     uptime / float32(count),
			Staked:     staked,
			Chains:     count,
			Delegators: daligators,
		}
		return c.JSON(http.StatusOK, APIResponse{
			Message: "Success",
			Code:    http.StatusOK,
			Data:    stats,
		})
	}
}

// GetEndpoints godoc
// @Summary Get endpoints
// @Description Get endpoints
// @Tags endpoints
// @Accept json
// @Produce json
// @Success 200 {object} APIResponse{data=Endpoints}
// @Router /endpoints [get]
func getEndpoints(cfg *config.Config) echo.HandlerFunc {
	return func(c echo.Context) error {
		var endpoints []Endpoints
		for _, chain := range cfg.Chains {
			endpoints = append(endpoints, Endpoints{
				Name:    chain.Name,
				ChainID: chain.ChainID,
				Path:    chain.Path,
				Rpc:     chain.Endpoints.Rpc,
				Rest:    chain.Endpoints.Api,
				Grpc:    chain.Endpoints.Grpc,
				Img:     chain.ImgURL,
			})

		}
		return c.JSON(http.StatusOK, APIResponse{
			Message: "Success",
			Code:    http.StatusOK,
			Data:    endpoints,
		})
	}
}

// GetSnapshots godoc
// @Summary Get snapshots
// @Description Get snapshots
// @Tags snapshots
// @Accept json
// @Produce json
// @Success 200 {object} APIResponse{data=[]Snapshots}
// @Router /snapshots [get]
func getSnapshots(cfg *config.Config) echo.HandlerFunc {
	return func(c echo.Context) error {
		var snapshots []Snapshots
		for _, chain := range cfg.Chains {
			if !chain.Snapshot.Enable {
				continue
			}

			snapshots = append(snapshots, Snapshots{
				Name:    chain.Name,
				ChainID: chain.ChainID,
				Path:    chain.Path,
				App:     chain.Appname,
				Go:      chain.GoVersion,
				Img:     chain.ImgURL,
				Base:    chain.Snapshot.SnapshotURL,
				Files:   chain.Snapshot.Files,
			})
		}

		return c.JSON(http.StatusOK, APIResponse{
			Message: "Success",
			Code:    http.StatusOK,
			Data:    snapshots,
		})
	}
}

// GetSnapshot godoc
// @Summary Get snapshot
// @Description Get snapshot
// @Tags snapshots
// @Accept json
// @Produce json
// @Param chain path string true "Chain ID"
// @Success 200 {object} APIResponse{data=Snapshots}
// @Failure 404 {object} APIResponse{data=nil}
// @Router /snapshot/{chain} [get]
func getSnapshot(cfg *config.Config) echo.HandlerFunc {
	return func(c echo.Context) error {
		path := c.Param("chain")
		var snapshots Snapshots

		for _, chain := range cfg.Chains {
			if chain.Path == path {
				if !chain.Snapshot.Enable {
					return c.JSON(http.StatusNotFound, APIResponse{
						Message: "Invalid chain ID or chain not found",
						Code:    http.StatusNotFound,
						Data:    nil,
					})
				}
				snapshots = Snapshots{
					Name:    chain.Name,
					ChainID: chain.ChainID,
					Path:    chain.Path,
					App:     chain.Appname,
					Go:      chain.GoVersion,
					Img:     chain.ImgURL,
					Base:    chain.Snapshot.SnapshotURL,
					Files:   chain.Snapshot.Files,
				}

				return c.JSON(http.StatusOK, APIResponse{
					Message: "Success",
					Code:    http.StatusOK,
					Data:    snapshots,
				})
			}
		}
		// error invalid chain id
		// If no chain matches the provided ID, return an error response
		return c.JSON(http.StatusNotFound, APIResponse{
			Message: "Invalid chain ID or chain not found",
			Code:    http.StatusNotFound,
			Data:    nil,
		})

	}
}
