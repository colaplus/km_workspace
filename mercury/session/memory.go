package session

import "sync"

type MemorySession struct {
	data   map[string]interface{}
	id     string
	rwlock sync.RWMutex
	flag   int
}

func NewMemorySession(id string) *MemorySession {
	m := &MemorySession{
		id:   id,
		data: make(map[string]interface{}, 8),
	}

	return m
}

func (m *MemorySession) Id() string {
	return m.id
}

func (m *MemorySession) IsModify() bool {
	if m.flag == SessionFlagModify {
		return true
	}

	return false
}

func (m *MemorySession) Set(key string, value interface{}) (err error) {
	m.rwlock.Lock()
	defer m.rwlock.Unlock()

	m.data[key] = value
	m.flag = SessionFlagModify
	return
}

func (m *MemorySession) Get(key string) (value interface{}, err error) {
	m.rwlock.RLock()
	defer m.rwlock.RUnlock()

	value, ok := m.data[key]
	if !ok {
		err = ErrSessionNotExist
		return
	}

	return
}

func (m *MemorySession) Del(key string) (err error) {
	m.rwlock.Lock()
	defer m.rwlock.Unlock()

	m.flag = SessionFlagModify
	delete(m.data, key)
	return
}

func (m *MemorySession) Save() (err error) {
	return
}