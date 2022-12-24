package main

type User struct {
	name string
}

type Users []User

func (u Users) Match(name string) (*User, error) {
	for i := 0; i < len(u); i++ {
		user := u[i]
		if user.name == name {
			return &user, nil
		}
	}
	return nil, nil
}

func main() {
	for i := 0; i < 100000; i++ {
		users := Users{
			User{name: "John"},
			User{name: "Jane"},
			User{name: "Josh"},
		}
		_, err := users.Match("Josh")
		if err != nil {
			panic(err)
		}
	}
}
