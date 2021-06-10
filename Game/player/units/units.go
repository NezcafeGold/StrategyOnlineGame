package units

import (
	"../../objects/units"
	"../../utils/config"
	"sync"
)

// Storage - хранилище юнитов пользователя.
type Storage struct {
	ID      string                 `json:"id"       bson:"_id"`      // ID
	MaxSize uint16                 `json:"max_size" bson:"max_size"` // Макс размер
	Size    uint16                 `json:"size"     bson:"size"`     // Текущий размер
	Units   map[string]*units.Unit `json:"units"    bson:"units"`    // Список юнитов
	RwMutex *sync.RWMutex                                            // RWMutex
}

// New - создать новое хранилище юнитов.
func New(id string) *Storage {
	return &Storage{
		ID:      id,
		MaxSize: config.MaxValueUnits,
		Size:    0,
		Units:   make(map[string]*units.Unit),
		RwMutex: new(sync.RWMutex),
	}
}

// Add - добавить юнита в хранилище.
func (storage *Storage) Add(un *units.Unit) {
	if storage.Size <= storage.MaxSize {
		storage.RwMutex.Lock()
		storage.Units[un.ID] = un
		storage.RwMutex.Unlock()

		storage.UpdateSize()
	}
}

// Get - получить юнита из хранилище.
func (storage *Storage) Get(id string) *units.Unit {
	storage.RwMutex.RLock()
	un := storage.Units[id]
	storage.RwMutex.RUnlock()

	return un
}

// Remove - удалить юнита из хранилища.
func (storage *Storage) Remove(unitID string) {
	storage.RwMutex.Lock()
	delete(storage.Units, unitID)
	storage.RwMutex.Unlock()

	storage.UpdateSize()
}

// UpdateSize - обновить инвентарь.
func (storage *Storage) UpdateSize() {
	storage.RwMutex.RLock()
	storage.Size = uint16(len(storage.Units))
	storage.RwMutex.RUnlock()
}