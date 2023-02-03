package data

type User struct {
	Id    string `json:"id"`
	Email string `json:"email"`
}

type UserModel struct{}

func (*UserModel) ById(id string) (User, error) {
	return User{
		Id:    id,
		Email: "tim@mck-p.com",
	}, nil
}

func (*UserModel) ByEmail(email string) (User, error) {
	return User{
		Id:    "1234",
		Email: email,
	}, nil
}

/*
Creates a new UserModel data object to be able to
manipulate and query the User Data.

TODO: Add in injection for Database
*/
func NewUserModel() UserModel {
	user := UserModel{}

	return user
}
