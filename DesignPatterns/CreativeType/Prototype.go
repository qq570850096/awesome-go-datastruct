package CreativeType


//设计一个类的时候，我们通常会使用到构造函数，这里类和对象的关系好比模具和构件的关系，
//对象总是从类中创建的。但是某些场景下是不允许类的调用者直接调用构造函数，也就说对象未
//必需要从类中衍生出来，现实生活中存在太多案例是通过直接“克隆” 来产生新的对象，而且
//克隆出来的本体和克隆体看不出任何区别。
//原型模式不单是一种设计模式，也是一种编程范型。 简单理解原型模式Prototype:不根据类
//来生成实例，而是根据实例生成新的实例。也就说，如果需要一个和某对象一 模一样的对象，那
//么就可以使用原型模式。
import "fmt"

type Role interface {
	Clone() Role
	SetHeadColor(string)
	SetEyesColor(string)
	SetTall(int64)
	Show()
}

type RoleChinese struct {
	HeadColor string
	EyesColor string
	Tall      int64
}

func (pR *RoleChinese) Clone() Role {
	var pChinese = &RoleChinese{HeadColor: pR.HeadColor, EyesColor: pR.EyesColor, Tall: pR.Tall}
	return pChinese
}

func (pR *RoleChinese) SetHeadColor(color string) {
	pR.HeadColor = color
}

func (pR *RoleChinese) SetEyesColor(color string) {
	pR.EyesColor = color
}

func (pR *RoleChinese) SetTall(tall int64) {
	pR.Tall = tall
}

func (pR *RoleChinese) Show() {
	fmt.Println("Role's headcolor is:", pR.HeadColor, " EyesColor is:", pR.EyesColor, " tall is:", pR.Tall)
}

