package blockchain

type Data struct {
	Sender   string
	Receiver string
	Amount   uint64
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
