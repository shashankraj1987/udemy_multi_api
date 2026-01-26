package models

type Models interface {
	Save()
	GetAll()
	GetId()
	Update()
	Delete()
}

// GetModels returns all available model types
func GetModels() []interface{} {
	return []interface{}{
		&User{},
		&Event{},
	}
}
