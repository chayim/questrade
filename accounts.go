package questrade

import "encoding/json"

type Account struct {
	Type              string
	Number            string
	Status            string
	IsPrimary         bool
	IsBilling         bool
	ClientAccountType string
}

type Accounts struct {
	Accounts []Account
}

func (qt Questrade) Accounts() []Account {

	accounts := Accounts{}
	res, _ := qt.request("v1/accounts")
	if res == nil {
		return nil
	}

	json.Unmarshal(res, &accounts)

	return accounts.Accounts

}
