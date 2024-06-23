package user

import "time"

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
