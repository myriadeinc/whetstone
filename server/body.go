package server

import (
	"encoding/json"
)

type JSONRpcReq struct {
	Id     *json.RawMessage `json:"id"`
	Method string           `json:"method"`
	Params *json.RawMessage `json:"params"`
}

type JSONRpcResp struct {
	Id      *json.RawMessage `json:"id"`
	Version string           `json:"jsonrpc"`
	Result  interface{}      `json:"result"`
	Error   interface{}      `json:"error"`
}

// Push Job to XMRig
type pushMessage struct {
	Version string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  *json.RawMessage `json:"params"`
}

type LoginRequest struct {
	Login string `json:"login"`
	// Password string
}

type StatusReply struct {
	Status string `json:"status"`
}

type ErrorReply struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type JobData struct {
	Blob     string `json:"blob"`
	JobID    string `json:"job_id"`
	Target   string `json:"target"`
	ID       string `json:"id"`
	Height   int    `json:"height"`
	SeedHash string `json:"seed_hash"`
	Algo     string `json:"algo"`
}

type LoginReply struct {
	Id     string   `json:"id"`
	Job    *JobData `json:"job"`
	Status string   `json:"status"`
}

type LoginData struct {
	ID     string  `json:"id"`
	Job    JobData `json:"job"`
	Status string  `json:"status"`
}

type JobPushData struct {
	Jsonrpc string  `json:"jsonrpc"`
	Method  string  `json:"method"`
	Params  JobData `json:"params"`
}