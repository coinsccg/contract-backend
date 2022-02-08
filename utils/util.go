package utils

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"auction/constant"
	"auction/logs"
)

func RequestBSCHoldersListByContractAddress(contractAddress, apiKey string, page, offset int) ([]byte, error) {
	url := fmt.Sprintf("%s&contractaddress=%s&page=%d&offset=%d&apikey=%s", constant.BSC_HOLDER_LIST_BY_CONTRACT_ADDRESS_URL, contractAddress, page, offset, apiKey)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		err = errors.New(string(content))
		return nil, err
	}
	return content, nil
}
