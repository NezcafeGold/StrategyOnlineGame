package resources

// Order - ордер.
type Order struct {
	ID           string  `json:"id"             bson:"_id"`            // ID.
	CostPerPiece float32 `json:"cost_per_piece" bson:"cost_per_piece"` // Стоимость за штуку.
	Value        uint32  `json:"value"          bson:"value"`          // Кол-во.
	Owner        Owner   `json:"owner"          bson:"owner"`          // Владелец.
}

// Owner - владелец ордера.
type Owner struct {
	ID       string `json:"id"        bson:"_id"`        // ID.
	NickName string `json:"nick_name" bson:"nick_name" ` // Ник игрока
}

// newOrder - создание нового ордера.
func newOrder(costPerPiece float32, value uint32, owner Owner) *Order {
	return &Order{
		ID:           owner.ID,
		CostPerPiece: costPerPiece,
		Value:        value,
		Owner:        owner,
	}
}