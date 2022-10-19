package domain

type UserRepositoryStub struct {
	users []User
}

func (s UserRepositoryStub) FindAll() ([]User, error) {
	return s.users, nil
}

func NewUserRepositoryStub() UserRepositoryStub {
	users := []User{
		{ID: "1", Name: "Ramon Ferreira", Birthdate: "1990-01-01", Password: "123", Email: "rfnascimento@ibm.com", Document: "12196183067", City: "Brasília", Zipcode: "70000000", Phone: "5561999991111", Status: 1},
	}

	return UserRepositoryStub{users: users}
}
