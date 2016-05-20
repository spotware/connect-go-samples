package main

// TradingAccount represents a trading account of user
type TradingAccount struct {
	AccountId                   *int64   `json:"accountId,omitempty"`
	AccountNumber               *int64   `json:"accountNumber,omitempty"`
	Live                        *bool    `json:"live,omitempty"`
	BrokerName                  *string  `json:"brokerName,omitempty"`
	BrokerTitle                 *string  `json:"brokerTitle,omitempty"`
	BrokerCode                  *int64   `json:"brokerCode,omitempty"`
	DepositCurrency             *string  `json:"depositCurrency,omitempty"`
	TraderRegistrationTimestamp *int64   `json:"traderRegistrationTimestamp,omitempty"`
	TraderAccountType           *string  `json:"traderAccountType,omitempty"`
	Leverage                    *int     `json:"leverage,omitempty"`
	Balance                     *int64   `json:"balance,omitempty"`
	Deleted                     *bool    `json:"deleted,omitempty"`
	AccountStatus               *string  `json:"accountStatus,omitempty"`
}

func (c *AccountsAPI) ListTradingAccounts() ([]TradingAccount, error) {
	u := "tradingaccounts"
	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}
	tradingAccounts := new([]TradingAccount)
	message := Message{}
	message.Data = tradingAccounts
	_, err = c.Do(req, message)
	if err != nil {
		return nil, err
	}
	if message.Error != nil {
		return nil, message.Error
	}

	return *tradingAccounts, err
}

func (k TradingAccount) String() string {
	return Stringify(k)
}