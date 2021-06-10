package materials

type TypeMaterial int

type Material struct {
	ID            string       `json:"id"              bson:"_id"`             // ID
	Title         string       `json:"title"           bson:"title"`           // Название
	TypeMaterial  TypeMaterial `json:"type_material"   bson:"type_material"`   // Тип материала
	Quantity      uint32       `json:"quantity"        bson:"quantity"`        // Кол-во
	CanBeSold     bool         `json:"can_be_sold "    bson:"can_be_sold "`    // Можно ли продать
	MaxValueStack uint32       `json:"max_value_stack" bson:"max_value_stack"` // Макс значение стака
}
