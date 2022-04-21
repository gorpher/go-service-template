package service

type User struct {
	ID        int64  `json:"id,omitempty" pg:"id"`
	UserName  string `json:"username,omitempty" pg:"username"`
	Password  string `json:"-" pg:"password"`
	Name      string `json:"name,omitempty" pg:"name"`
	Email     string `json:"email,omitempty" pg:"email"`
	Phone     string `json:"phone,omitempty" pg:"phone"`
	Note      string `json:"note,omitempty" pg:"note"`
	CreatedAt int64  `json:"created_at,omitempty" pg:"created_at"`
	DeletedAt int64  `json:"deleted_at,omitempty" pg:"deleted_at"`
}
