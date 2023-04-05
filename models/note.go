package models

import (
	"gin_notes/database"
	"time"

	"gorm.io/gorm"
)

type Note struct {
	ID        uint64         `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string         `gorm:"size:255;not null" json:"name"`
	Content   string         `gorm:"type:text;not null" json:"content"`
	UserID    uint64         `gorm:"index" json:"user_id"`
	CreatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time      `gorm:"index;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func NotesAll(user *User) *[]Note {
	var notes []Note
	database.Database.Where("deleted_at IS NULL and user_id = ?", user.ID).Order("updated_at desc").Find(&notes)
	return &notes
}

func NoteCreate(user *User, name string, content string) *Note {
	entry := Note{Name: name, Content: content, UserID: user.ID}
	database.Database.Create(&entry)
	return &entry
}

func NotesFind(user *User, id uint64) *Note {
	var note Note
	database.Database.Where("id = ? and user_id =?", id, user.ID).First(&note)
	return &note
}

func (note *Note) Update(name string, content string) {
	note.Name = name
	note.Content = content
	database.Database.Save(&note)
}

func NotesMarkDelete(user *User, id uint64) {
	// Update notes set deleted_at==Current Time> WHERE id = <id>
	database.Database.Where("id = ? and user_id =?", id, user.ID).Delete(&Note{})
}
