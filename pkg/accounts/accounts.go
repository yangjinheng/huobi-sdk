package accounts

import (
	"encoding/json"
	"fmt"
	"huobi-sdk/pkg/utils"
	"io/ioutil"
	"net/http"
)

// GetAccountInfo 获取账户基本信息
func GetAccountInfo() (*GetAccountInfoResponse, error) {
	request, err := http.NewRequest("GET", "/v1/account/accounts", nil)
	if err != nil {
		return nil, err
	}
	utils.ReqBuilder.Sign(request)
	client := utils.DefaultHTTPClient(utils.ProxyAddr)
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	accounts := &GetAccountInfoResponse{}
	err = json.Unmarshal(body, accounts)
	if err != nil {
		return nil, err
	}
	return accounts, nil
}

// GetAccountBalance 根据账户 ID 获取账户余额
func GetAccountBalance(accountID int) (*GetAccountBalanceResponse, error) {
	request, err := http.NewRequest("GET", "/v1/account/accounts/" + fmt.Sprintf("%d", accountID) + "/balance", nil)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	utils.ReqBuilder.Sign(request)
	client := utils.DefaultHTTPClient(utils.ProxyAddr)
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	balances := &GetAccountBalanceResponse{}
	err = json.Unmarshal(body, balances)
	if err != nil {
		return nil, err
	}
	return balances, nil
}

// GetAccountValuation 获取账户估值
func GetAccountValuation(accountType string) (*GetAccountValuationResponse, error) {
	request, err := http.NewRequest("GET", "/v2/account/asset-valuation?valuationCurrency=CNY&accountType=" + accountType, nil)
	if err != nil {
		return nil,err
	}
	utils.ReqBuilder.Sign(request)
	client := utils.DefaultHTTPClient(utils.ProxyAddr)
	resp, err := client.Do(request)
	if err != nil {
		return nil,err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil,err
	}
	valuation := &GetAccountValuationResponse{}
	err = json.Unmarshal(body, valuation)
	if err != nil {
		return nil, err
	}
	return valuation, nil
}
