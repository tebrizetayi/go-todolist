package models

type TodoItem struct {
	MyGormModel
	Title       string `sql:"type:varchar(254);not null"`
	Description string `sql:"type:varchar(254);"`
	Done        bool
	UserID      int `sql:"type:int;not null"`
	User        *User
}

type User struct {
	MyGormModel
	Username  string `sql:"type:varchar(254);unique;not null"`
	Firstname string `sql:"type:varchar(254);not null"`
	Lastname  string `sql:"type:varchar(254);not null"`
	Password  string `sql:"type:varchar(60);"`
}

func NewUser(username string, firstname string, lastname string, password string) *User {
	user := &User{
		Username:  username,
		Firstname: firstname,
		Lastname:  lastname,
		Password:  password,
	}
	return user
}

func NewToDoItem(user *User, title string, description string) *TodoItem {
	toDoItem := &TodoItem{
		Description: description,
		Title:       title,
		UserID:      user.ID,
	}

	return toDoItem
}
