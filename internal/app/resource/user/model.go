package user

import (
	"time"
)

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Password  string    `json:"password"`
	DOB       time.Time `json:"dob"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateParams struct {
	Name     string `json:"name" form:"required,max=255"`
	Email    string `json:"email" form:"required,email"`
	Phone    string `json:"phone" form:"required,max=15"`
	Password string `json:"password" form:"required,min=8"`
	DOB      string `json:"dob" form:"required,datetime=2006-01-02"`
	Role     string `json:"role" form:"required,oneof=admin customer"`
}

type SerializedUser struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	DOB   string `json:"dob"`
	Role  string `json:"role"`
}

type Users []*User

func (u *User) Serialized() *SerializedUser {
	return &SerializedUser{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
		Phone: u.Phone,
		DOB:   u.DOB.Format("2006-01-02"),
		Role:  u.Role,
	}
}

func (us Users) Serialized() []*SerializedUser {
	serializedUsers := make([]*SerializedUser, len(us))
	for i, v := range us {
		serializedUsers[i] = v.Serialized()
	}

	return serializedUsers
}

func (c *CreateParams) ToUser() *User {
	dob, _ := time.Parse("2006-01-02", c.DOB)

	return &User{
		Name:      c.Name,
		Email:     c.Email,
		Phone:     c.Phone,
		Password:  c.Password,
		DOB:       dob,
		Role:      c.Role,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
