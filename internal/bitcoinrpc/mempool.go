package bitcoinrpc

import "fmt"

func (br *BitcoinRpc) GetMempoolTransactions() error {

	response, err := br.request().SetQueryParams(map[string]string{"id": "1", "method": "getmempoolinfo", "params": "[ true ]"}).Get(br.url())
	if err != nil {

		return err
	}

	data := response.Body()

	fmt.Println(string(data))

	return nil
}
