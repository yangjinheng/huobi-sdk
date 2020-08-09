package accounts

// GetAccountInfoResponse 接口 GET /v1/account/accounts 的响应
type GetAccountInfoResponse struct {
	Status string `json:"status"`
	Data []AccountInfo `json:"data"`
}

// AccountInfo 账户 ID 及其相关信息
type AccountInfo struct {
	ID      int    `json:"id"`
	State   string `json:"state"`
	Subtype string `json:"subtype"`
	Type    string `json:"type"`
}



// GetAccountBalanceResponse 接口 GET /v1/account/accounts/{account-id}/balance 的响应
type GetAccountBalanceResponse struct {
	Status string `json:"status"`
	Data   BalanceData   `json:"data"`  
}

// BalanceData 某种账户下的余额数据
type BalanceData struct {
	ID    int64  `json:"id"`   
	Type  string `json:"type"` 
	State string `json:"state"`
	List  []CoinList `json:"list"` 
}

// CoinList 账户下每个币种的余额数据
type CoinList struct {
	Currency string `json:"currency"`
	Type     string `json:"type"`    
	Balance  string `json:"balance"` 
}



// GetAccountValuationResponse 获取 GET /v2/account/asset-valuation 的响应
type GetAccountValuationResponse struct {
	Code int64 `json:"code"`
	Ok   bool  `json:"ok"`  
	Data ValuationData  `json:"data"`
}

// ValuationData 某种账户按照某个币种估值的结果
type ValuationData struct {
	Balance   string `json:"balance"`  
	Timestamp int64  `json:"timestamp"`
}
