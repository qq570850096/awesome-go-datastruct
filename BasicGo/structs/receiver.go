package structdemo

// User 用于演示方法接收者和值/指针传参的差异。
type User struct {
	Name string
	Age  int
}

// RenameValue 使用值接收者，不会修改调用方持有的 User。
func (u User) RenameValue(newName string) {
	u.Name = newName
}

// RenamePointer 使用指针接收者，会直接修改调用方持有的 User。
func (u *User) RenamePointer(newName string) {
	u.Name = newName
}

// ChangeUserValue 使用值传递，调用者看到的 Age 不会变化。
func ChangeUserValue(u User) {
	u.Age++
}

// ChangeUserPointer 使用指针传递，调用者可见 Age 被修改。
func ChangeUserPointer(u *User) {
	u.Age++
}

