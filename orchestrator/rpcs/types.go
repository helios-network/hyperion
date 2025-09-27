package rpcs

import "time"

type RpcUsage struct {
	Success bool      `json:"success"`
	Error   string    `json:"error"`
	Time    time.Time `json:"time"`
}

type Rpc struct {
	Url       string     `json:"url"`
	IsPrimary bool       `json:"is_primary"`
	Usages    []RpcUsage `json:"usages"`
	IsTested  bool       `json:"is_tested"`
}
