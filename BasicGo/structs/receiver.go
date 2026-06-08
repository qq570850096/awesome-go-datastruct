package structdemo

type User struct {
	Name string
	Age  int
}

func (u User) RenameValue(newName string) {
	u.Name = newName
}

func (u *User) RenamePointer(newName string) {
	u.Name = newName
}

func ChangeUserValue(u User) {
	u.Age++
}

func ChangeUserPointer(u *User) {
	u.Age++
}
