package resources

type Gold  float64  // Золото.
type Iron  float64  // Железо.
type Stone float64  // Камень.
type Wood  float64  // Дерево.
type Food  float64  // Еда.

// Resources - ресурсы.
type Resources struct {
	Gold  Gold  `json:"gold"  bson:"gold"`  // Золото.
	Iron  Iron  `json:"iron"  bson:"iron"`  // Железо.
	Stone Stone `json:"stone" bson:"stone"` // Камень.
	Wood  Wood  `json:"wood"  bson:"wood"`  // Дерево.
	Food  Food  `json:"food"  bson:"food"`  // Еда.
}

// New - создать ресурсы с 0 значением.
func New() *Resources {
	return &Resources{
		Gold:  0.0,
		Iron:  0.0,
		Stone: 0.0,
		Wood:  0.0,
		Food:  0.0,
	}
}

