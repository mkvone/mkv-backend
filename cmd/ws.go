package cmd

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/mkvone/mkv-backend/cmd/config"
	"github.com/sacOO7/gowebsocket"
)

func TransformToWebSocketURL(rpcURL string) string {
	return strings.Replace(rpcURL, "https://", "wss://", 1) + "/websocket"
}

type WebSocketMessage struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int64  `json:"id"`
	Result  Result `json:"result"`
}

type Result struct {
	Query  string              `json:"query"`
	Data   ResultData          `json:"data"`
	Events map[string][]string `json:"events"`
}

type ResultData struct {
	Type  string    `json:"type"`
	Value DataValue `json:"value"`
}

type DataValue struct {
	Block ValueBlock `json:"block"`
}

type ValueBlock struct {
	Header   Header        `json:"header"`
	Data     BlockData     `json:"data"`
	Evidence BlockEvidence `json:"evidence"`
}

type BlockData struct {
	Txs []string `json:"txs"`
}

type BlockEvidence struct {
	Evidence []interface{} `json:"evidence"`
}

type Header struct {
	ChainID            string `json:"chain_id"`
	Height             string `json:"height"`
	Time               string `json:"time"`
	LastCommitHash     string `json:"last_commit_hash"`
	DataHash           string `json:"data_hash"`
	ValidatorsHash     string `json:"validators_hash"`
	NextValidatorsHash string `json:"next_validators_hash"`
	ConsensusHash      string `json:"consensus_hash"`
	AppHash            string `json:"app_hash"`
	LastResultsHash    string `json:"last_results_hash"`
	EvidenceHash       string `json:"evidence_hash"`
	ProposerAddress    string `json:"proposer_address"`
}

func extractDataFromMessage(message string) (string, string, error) {
	var msg WebSocketMessage
	err := json.Unmarshal([]byte(message), &msg)
	if err != nil {
		return "", "", err
	}

	return msg.Result.Data.Value.Block.Header.Height, msg.Result.Data.Value.Block.Header.Time, nil
}
func SubscribeToNewBlocks(cc config.ChainConfig) {
	wsURL := TransformToWebSocketURL(cc.Endpoints.Rpc)
	socket := gowebsocket.New(wsURL)
	reconnectFunc := func() {
		time.Sleep(3 * time.Minute)
		socket.Connect()
		socket.SendText("{ \"jsonrpc\": \"2.0\", \"method\": \"subscribe\", \"params\": [\"tm.event='NewBlock'\"], \"id\": 1 }")

	}
	socket.OnConnected = func(socket gowebsocket.Socket) {
		l("🩷  Connected to ws server : ", cc.Name)
	}
	socket.OnTextMessage = func(message string, socket gowebsocket.Socket) {
		height, timestamp, err := extractDataFromMessage(message)
		if err != nil {
			l("❌ Error extracting data from message: ", err)
			return
		}
		// 채널을 통해 블록 정보를 전송
		blockInfoChan <- BlockInfo{
			Height:    height,
			Time:      timestamp,
			ChainName: cc.Name,
		}
	}

	socket.OnDisconnected = func(err error, socket gowebsocket.Socket) {
		// log.Printf("WebSocket disconnected: %v. Reconnecting...", err)
		err = fmt.Errorf("WebSocket disconnected(%s): %v. Reconnecting...", wsURL, err)
		l("❌ ", err)
		reconnectFunc()
	}
	socket.OnConnectError = func(err error, socket gowebsocket.Socket) {
		err = fmt.Errorf("WebSocket disconnected(%s): %v. Reconnecting...", wsURL, err)
		l("❌ ", err)
		reconnectFunc()
	}

	socket.Connect()
	// Subscribe to NewBlock events
	socket.SendText("{ \"jsonrpc\": \"2.0\", \"method\": \"subscribe\", \"params\": [\"tm.event='NewBlock'\"], \"id\": 1 }")

}
