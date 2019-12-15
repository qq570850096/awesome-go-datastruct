package CreativeType

import "fmt"

//所谓万丈高楼平地起，但是我们建造(Build) 高楼时，需要经历很多阶段，比如打地基、搭框
//架、浇筑水泥、封顶等,这些都是很难一气呵成的。 所以一般我们是先建造组成高楼的各个部
//分，然后将其-个个地组装起来，好比搭积木一般，分阶段拼接后组装成一个完整的物体。 还有
//个问题，就是同样的积木，同样的搭建过程，却能Build出不同的物体，这就叫做建造者模
//式。
//将一个复杂的对象的构建与它的表示相分离，使得同样的构建过程可以创建出不同的表示。建造
//者模式(Builder Pattern)也叫做生成器模式。


//建造者模式通常有以下几部分角色组成:
//●建造者(Builder) : Builder 角色负责定义用来生成实例的接口(API) ;
//●具体的建造者(ConcreateBuilder) : ConcreateBuilder角色负责实现Builder角色定义
//的接口的实现类。针对不同的商业逻辑，具体化复杂对象的各部分的创建。在建造完成之后
//提供产品的实例;
//●监工(Director) : Director角色负责使用Builder角色的接口API来生成实例。内部不涉
//及具体产品信息，只负责保证对象各部分完整创建或按照某种顺序进行创建。即Director是
//负责指挥如何build的，只负责调度，具体实施交给具体的建造者;
//●产品(Product) :即要创建的复杂对象;
//●使用者(Client) :实际使用Builder模式的角色，即下面的测试类。

type Product struct {
	ground string
	cement string
	roof string
}

func (p *Product) Cement() string {
	return p.cement
}

func (p *Product) SetCement(cement string) {
	p.cement = cement
}

func (p *Product) Roof() string {
	return p.roof
}

func (p *Product) SetRoof(roof string) {
	p.roof = roof
}

func (p *Product) Ground() string {
	return p.ground
}

func (p *Product) SetGround(ground string) {
	p.ground = ground
}

type Builder interface {
	BuildGround()
	BuildCement()
	BuildRoof()
	
	BuildProduct() *Product
}

type ConcreteBuilder struct {
	p *Product
}

func (this *ConcreteBuilder) BuildGround() {
	this.p.SetGround("建造地基")
	fmt.Println(this.p.ground)
}

func (this *ConcreteBuilder) BuildCement() {
	this.p.SetCement("建造房子")
	fmt.Println(this.p.Cement())
}
func (this *ConcreteBuilder) BuildRoof() {
	this.p.SetRoof("建造房顶")
	fmt.Println(this.p.Roof())
}

func (this *ConcreteBuilder) BuildProduct() *Product {
	fmt.Println("建造完毕")
	return this.p
}

type Director struct {
	builder Builder
}

func (this *Director) Construst() Product {
	this.builder.BuildGround()
	this.builder.BuildCement()
	this.builder.BuildRoof()

	return *this.builder.BuildProduct()
}
// 优缺点
//●封装性:客户端不必知道产品内部组合细节，只需关心我们要生成某个对象，具体对象产生
//细节不必知晓。Main 类并不知道Builder类,它只是调用了Director类的construct方法
//完成对象的获取;
//●建造者独立，易于拓展:上面我们只列举了ConcreteBuilder建造者类，如果需要其它建造
//者新建类即可。建造者之间彼此独立，系统拓展性好，符合开闭原则;
//●便于控制细节风险:由于具体建造者是独立的，因此可以对具体建造过程逐步细化，不会对
//其它模块产生影响。


