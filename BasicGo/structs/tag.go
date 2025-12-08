package structdemo

import "encoding/json"

// Account 展示导出字段、json tag 与字段忽略规则。
type Account struct {
	ID    int    `json:"id"`
	Name  string `json:"name,omitempty"`
	Token string `json:"-"`
}

func EncodeAccount(a Account) ([]byte, error) {
	return json.Marshal(a)
}

func DecodeAccount(data []byte) (Account, error) {
	var a Account
	err := json.Unmarshal(data, &a)
	return a, err
}

