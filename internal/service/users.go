package service

type AccountData struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func LogIn(data AccountData) bool {
	// Process the JSON data
}

// Модель пользователей:
type UserModel struct {
	ID             int    `gorm:"primaryKey"` // (уникальный идентификатор пользователя)
	Email          string // (электронная почта пользователя)
	Username       string // (логин пользователя)
	Password       string // (хэш пароля пользователя)
	AvatarURL      string // (ссылка на аватар пользователя)
	CreatedAt      string // (дата создания аккаунта)
	RatingLikes    string // (количество лайков, полученных пользователями за его сообщения и треды)
	RatingDislikes string // (количество дизлайков, полученных пользователями за его сообщения и треды)
}

//// Модель админа
//type AdminModel struct {
//	ID     int `gorm:"primaryKey"` // уникальный идентификатор админа
//	UserID int // ссылка на пользователя, которому установлены права администратора
//}

// Модель треда
type ThreadModel struct {
	ID            int    `gorm:"primaryKey"` // уникальный идентификатор треда
	UserID        int    // ссылка на создателя треда (пользователя)
	TopicID       int    // ссылка на тему треда(ид темы)
	Title         string // название треда
	Content       string // содержание треда
	CreatedAt     string // дата создания треда
	LikesCount    int    // количество лайков, полученных за тред
	DislikesCount int    // количество дизлайков, полученных за тред
}

// Модель сообщения
type MessageModel struct {
	ID            int    `gorm:"primaryKey"` // уникальный идентификатор сообщения
	UserID        int    // ссылка на автора сообщения
	ThreadID      int    // ссылка на тред, в котором написано сообщение
	Content       string // содержание сообщения
	CreatedAt     string // дата создания сообщения
	LikesCount    int    // количество лайков, полученных за сообщение
	DislikesCount int    // количество дизлайков, полученных за сообщение
}

// Модель темы
type TopicModel struct {
	ID         int    `gorm:"primaryKey"` // уникальный идентификатор темы
	Title      string // название темы
	CreatedAt  string // дата создания темы
	Rules      string // правила темы, сохраненные в формате HTML
	ViewsCount int    // количество просмотров темы
}

// Модель пометок
type NoteModel struct {
	ID       int    `gorm:"primaryKey"` // уникальный идентификатор пометки
	UserID   int    // ссылка на пользователя, который оставил пометку
	ThreadID int    // ссылка на тред, которому присвоена пометка
	Content  string // содержание пометки
}
