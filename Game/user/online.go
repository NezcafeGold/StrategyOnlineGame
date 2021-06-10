package user

import (
	"errors"
	"sync"
)

// ListOnlineUsers - список онлайн пользователей.
type ListOnlineUsers struct {
	storage map[string]*User // Хранилище
	rwMutex *sync.RWMutex    // RWMutex
}

// NewListOnlineUsers - создание списка онлайн пользователей.
func NewListOnlineUsers() *ListOnlineUsers {
	return &ListOnlineUsers{
		storage: make(map[string]*User),
		rwMutex: new(sync.RWMutex),
	}
}

// Delete - удалить пользователя из онлайн списка.
func (listOnlineUsers *ListOnlineUsers) Delete(id string) error {
	listOnlineUsers.rwMutex.RLock()
	_, ok := listOnlineUsers.storage[id]
	listOnlineUsers.rwMutex.RUnlock()
	
	if ok {
		listOnlineUsers.rwMutex.Lock()
		delete(listOnlineUsers.storage, id)
		listOnlineUsers.rwMutex.Unlock()

		return nil
	}

	return errors.New("Пользователь не найден. ")
}

// Add - добавить пользователя в онлайн список.
func (listOnlineUsers *ListOnlineUsers) Add(u *User) error {
	listOnlineUsers.rwMutex.RLock()
	_, ok := listOnlineUsers.storage[u.ID]
	listOnlineUsers.rwMutex.RUnlock()
	
	if ok {
		return errors.New("Пользователь уже онлайн. ")
	}

	listOnlineUsers.rwMutex.Lock()
	listOnlineUsers.storage[u.ID] = u
	listOnlineUsers.rwMutex.Unlock()
	
	return nil
}

// Get - получить пользователя из онлайн списка.
func (listOnlineUsers *ListOnlineUsers) Get(id string) (*User, error) {
	listOnlineUsers.rwMutex.RLock()
	u, ok := listOnlineUsers.storage[id]
	listOnlineUsers.rwMutex.RUnlock()
	
	if ok {
		return u, nil
	}
	
	return nil, errors.New("Пользователь не найден. ")
}