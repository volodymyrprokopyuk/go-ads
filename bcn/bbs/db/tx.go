package db

type Account string

type Tx struct {
  From Account `json:"from"`
  To Account `json:"to"`
  Value uint `json:"value"`
  Data string `json:"data"`
}

func (t *Tx) IsReward() bool {
  return t.Data == "reward"
}
