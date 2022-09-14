package blockchain

type Data struct {
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
	Amount   uint64 `json:"amount"`
}

func NewData(
	sender string,
	receiver string,
	amount uint64,
) *Data {
	return &Data{
		Sender:   sender,
		Receiver: receiver,
		Amount:   amount,
	}
}
