## 设计模式——golang实现

> **生活的理想，就是为了理想的生活。                ——张闻天**

[toc]

### 创造型模式

#### 简单工厂模式



##### **意图:**

定义一个创建对象的接口，让其子类自己决定实例化哪一个工厂类，工厂模式使其创建过程延迟到子类进行
主要解决接口选择问题

##### **类图：**

![](http://www.liuanqihappybirthday.top/uploads/big/9e71d19fd66298e5a404727f99dc941a.png)

**返回的实例都实现同一接口**

##### 应用实例:

[源代码](./SimpleFactory.go)。


#### 抽象工厂模式

[源代码](./AbstractFactory.go)

抽象工厂模式是工厂模式(简单工厂、工厂方法)中最具抽象和一般性的一种形态。抽象工厂模式可以向客户端提供一个接口， 使客户端在不必指定产品的具体的情况下，创建多个产品族中的产品对象。

##### **定义**

抽象工厂模式的定义:**为创建一组相关或相互依赖的对象提供一个接口，且无须指定它们的具体类。**

![](http://www.liuanqihappybirthday.top/uploads/big/cd467410c25cd99a624ad1eed695f90a.jpg)

**同产品等级是同一类产品，同产品族的产品是由一个工厂生成的。**

##### **模式特点**

抽象工厂模式使用时，一般具有如下特点:

●系统中有多个产品族，每个具体工厂负责创建同一族但属于不同产品等级(产品种类)的产品

●系统一次只能消费某一族产品，即相同产品族的产品是一起被使用的当系统需要新增一个产品族时，只需要增加新的工厂类即可，无需修改源代码;**但是如果需要产品族中增加一个新种类的产品时，则所有的工厂类都需要修改。**

抽象工厂模式中的**抽象工厂类的职责就是定义每个工厂要实现的功能**，即定义多个产品族的产品的创建。这里，同一产品族下有多个产品时，对应的抽象工厂就会有多个抽象方法用来提供创建这些产品的接口。

##### **抽象工厂的角色**

抽象工厂-般包含四种角色，分别是: 

●抽象工厂(Abstract Factory) :提供了创建产品的接口，包含多个创建产品的方法，即包含多个类似new Product 0的方法;

●具体工厂(Concrete Factory) :实现抽象工厂定义的接口，完成某个具体产品的创建;

●抽象产品(Abstract Product) :抽象产品定义，-般有多少抽象产品，抽象工厂中就包含多少个创建产品的方法;

●具体产品. (Concrete Product) :抽象产品的实现类。

实现抽象产品所定义的接口，和具体工厂是多对一的关系: 

![](http://www.liuanqihappybirthday.top/uploads/big/e29f55abf781438fc4c73d8f6c39512c.jpg)

##### **具体实例：**

美的和TCL都会生产TV和空调，这时候，美的和TCL就是两个产品族，TV和空调就是两个产品等级，按照上面的类图照葫芦画瓢就能得到如下的类图。

![](http://www.liuanqihappybirthday.top/uploads/big/1a8006b8f0aa5fb48f7b4bc02b44b59b.jpg)

加入某天，需要加入海尔的TV和空调，那么只需要对这里做出如下修改就可以了。

![](http://www.liuanqihappybirthday.top/uploads/big/2f69a7c6093fa49ab16740897f53001b.jpg)

##### **总结：**

![](http://www.liuanqihappybirthday.top/uploads/big/993a187228c2a3bd0c299982d80bc974.jpg)

#### 单例模式

比较简单的一种设计模式，就不单独开一个源码了。

单例模式，是一种常用的软件设计模式，在它的核心结构中只包含一个被称为单例的特殊类。通过单例模式可以保证系统中一个类只有一个实例且该实例易于外界访问，从而方便对实例个数的控制并节约系统资源。

##### 1. 懒汉模式（Lazy Loading）

懒汉模式是开源项目中使用最多的一种，最大的缺点是非线程安全的



```go
type singleton struct {
}

// private
var instance *singleton

// public
func GetInstance() *singleton {
    if instance == nil {
        instance = &singleton{}     // not thread safe
    }
    return instance
}
```

##### 2. 带锁的单例模式



```go
type singleton struct {
}

var instance *singleton
var mu sync.Mutex

func GetInstance() *singleton {
    mu.Lock()
    defer mu.Unlock()

    if instance == nil {
        instance = &singleton{}     // unnecessary locking if instance already created
    }
    return instance
}
```

这里使用了Go的sync.Mutex,其工作模型类似于Linux内核的futex对象，具体实现极其简单，性能也有保证
 初始化时填入的0值将mutex设定在未锁定状态，同时保证时间开销最小
 这一特性允许将mutex作为其它对象的子对象使用

##### 3. 带检查锁的单例模式



```go
    if instance == nil {     // <-- Not yet perfect. since it's not fully atomic
        mu.Lock()
        defer mu.Unlock()
 
        if instance == nil {
            instance = &singleton{}
        }
    }
    return instance
}
```

这是一个不错的方法，但是还并不是很完美。因为编译器优化没有检查实例存储状态。如果使用sync/atomic包的话 就可以自动帮我们加载和设置标记。



```go
import "sync"
import "sync/atomic"

var initialized uint32
...

func GetInstance() *singleton {

    if atomic.LoadUInt32(&initialized) == 1 {
        return instance
    }

    mu.Lock()
    defer mu.Unlock()

    if initialized == 0 {
         instance = &singleton{}
         atomic.StoreUint32(&initialized, 1)
    }

    return instance
}
```

##### 4.比较好的一种方式sync.Once



```go
import (
    "sync"
)
 
type singleton struct {
}
 
var instance *singleton
var once sync.Once
 
func GetInstance() *singleton {
    once.Do(func() {
        instance = &singleton{}
    })
    return instance
}
```

#### 建造者模式

**将一个复杂的对象的构建与它的表示相分离，使得同样的构建过程可以创建出不同的表示。建造者模式(Builder Pattern)也叫做生成器模式。**

##### 组成角色

建造者模式通常有以下几部分角色组成:

●建造者(Builder) : Builder 角色负责定义用来生成实例的接口(API) ;

●具体的建造者(ConcreateBuilder) : ConcreateBuilder 角色是负责实现Builder角色定义的接口的实现类。针对不同的商业逻辑，具体化复杂对象的各部分的创建。在建造完成之后提供产品的实例;

●监工(Director) : Director角色负责使用Builder色的接口API来生成实例。内部不涉及具体产品信息，只负责保证对象各部分完整创建或按照某种顺序进行创建。即Director是负责指挥如何build的，只负责调度，具体实施交给具体的建造者;

●产品(Product) :即要创建的复杂对象;

●使用者(Client) :实际使用Builder模式的角色，即下面的[测试类](./builder_test.go)。

##### UML类图

![](http://www.liuanqihappybirthday.top/uploads/big/c5b045f77f3cdacee4a1db9da8532715.jpg)

##### 应用场景

建造者模式的典型应用场景如下:

●产品类非常复杂，不同的调度产生不同的结果时，使用建造者模式比较适合;

●相同的组件或配件都可以装配到一个对象，但是产生的结果又不相同，可以使用建造者模式。

##### 建造者模式VS工厂方法模式

建造者模式关注的是零件类型和装配顺序(工艺)同为创建型模式，注重点不同。另外工厂模式只有一个建造方法，而**建造者模式有多个建造零部件的方法并且强调建造顺序,而工厂模式没有顺序的概念。**

##### 总结

![](http://www.liuanqihappybirthday.top/uploads/big/4822afb70940e7caab56390ee2a4c9e6.jpg)

#### 原型模式

原型模式不单是一种设计模式，也是一种编程范型。 简单理解原型模式Prototype:不根据类来生成实例，而是根据实例生成新的实例。也就说，如果需要一个和某个对象一模一样的对象，那么就可以使用原型模式。

##### 定义

从设计模式的角度讲，原型模式是一种创建型模式， 摆脱了类的构造模式，原型模式告诉我们，想要创建一个对象， 我们不必关心对象的具体类型，而是找到一个对象，然后通过克隆来创建一个一模一样的对象。

##### 扩展阅读: 深拷贝与浅拷贝

`go` 语言中的传递都是值传递，传递一个对象，就会把对象拷贝一份传入函数中，传递一个指针，就会把指针拷贝一份传入进去。

赋值的时候也是这样，`res:=*e` 就会把传递的 `Example` 对象拷贝一份，如果是 `res:=e` 的话，那么拷贝的就是对象的指针了.

而深拷贝和浅拷贝也可以这样理解，深拷贝就是拷贝整个对象，浅拷贝就是拷贝对象指针。

对于深度拷贝，`go`和其他语言还经常使用序列化后反序列化的形式进行拷贝:

```go
 func deepCopy(dst, src interface{}) error {
     var buf bytes.Buffer
     if err := gob.NewEncoder(&buf).Encode(src); err != nil {
         return err
     }
     return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
  }
复制代码
```

实际上`gob`包序列化的时候也是用到了 `reflect`包来实现拷贝的

**注意:** golang完全是按值传递，所以如果深度拷贝的对象中包含有指针的话，那么深度拷贝后，这些指针也会相同，会导致部分数据共享，要注意这一点.

##### 总结

![](https://img.mukewang.com/5db11db8000175f805570171.png)

### 结构型模式

#### 适配器模式

##### 定义

适配器，其实很好理解，生活中也随处可见，比如电源适配器、usb 适配器等等,那么适配器模式，也被称为Wrapper模式。

Wrapper有“包装器”的意思，适配器模式的定义是:**将一个接口转换成客户希望的另一个接口，使接口不兼容的那些类可以一起工作，解决的痛点便是因接口不兼容导致的类不能正常工作的问题。**

![](http://www.liuanqihappybirthday.top/uploads/big/1b3391b9b7e5100614206acbc732e90b.png)

如上图所示，A、B代表已经成型的物体，如果这两个物体想要安装在一起，因为两者接口不兼容，不可能直接安装在一起，这时候怎么办？我们就可以通过引入一个物体C，物体C既要适配A也要适配B的接口。

什么时候使用适配器模式，从上面的案例我们也可以看出一点端倪:

● 现有的类或接口不能满足需求，且一般无法直接修改现有类或接口。比方该类为三方提供，就无法修改，亦或者像A、B这种已经塑模成型的物件，可能已大规模在使用中，所以不允许修改。

● 想要建立一个可以重复使用的类，用于与-些彼此之间没有太大关联的一些类，包括-些可能在将来引进的类一起工作，这些源类不一定有一致的接口。

##### 组成角色

![](http://www.liuanqihappybirthday.top/uploads/big/38243f22a6520b11632da2da23d054dc.png)

●目标角色(Target) :该角色定义把其它类转换为何种接口，也就是我们的期望接口，可以是一个抽象类或接口，也可以是具体类。

●适配器角色(Adapter) :适配器可以调用另-一个接口，作为一个转换器，对Adaptee和Target进行适配，适配器类是适配器模式的核心，通常都是一个具体的类。

●源角色(被适配Adaptee) :你想把谁转换成目标角色，这个“谁”就是源角色，它是已经存在的、运行良好的类或对象，经过适配器角色的包装，它会成为一个崭新、靓丽的角色。

●请求者(Client) :该角色负责使用Target定义的方法进行具体处理。

##### 类适配器（使用继承）

![](http://www.liuanqihappybirthday.top/uploads/big/0feccc9fa922530048ee1d88bd06fcf5.png)

##### 优缺点

主要优点:
●将目标类和适配者类解耦，通过引入一个适配器类来重用现有的适配者类，无须修改原有结构;

●增加了类的透明性和复用性，将具体的业务实现过程封装在适配者类中，对于客户端类而言是透明的，而且提高了适配者的复用性，同一个适配者类可以在多个不同的系统中复用;

●可以将两个互不相干的类关联在一起;

●增强系统灵活性。

主要缺点:
●类适配器对于Java、C#等不支持多重类继承的语言，一次最多只能适配一个适配者类，不能同时适配多个适配者。

##### 总结

![](https://img.mukewang.com/5db14584000195ff06350366.png)

#### 桥接模式

##### 定义

桥接模式(Bridge Pattern):将抽象部分与它的实现部分分离，使它们都可以独立地变化。它是一种对象结构型模式， 又称为柄体(Handle and Body)模式或接口(Interface) 模式。

Bridge的意思就是“桥梁”， 好比现实生活中的桥梁， 它的存在就是将河流两侧东西给连接起来，应用到软件里面Bridge就是将**类的功能层次结构与实现层次结构连接起来。**

##### UML类图

![](https://img1.sycdn.imooc.com/5db146f20001437210430651.png)

桥接模式中包含了几种角色，分别是: 

●抽象化(Abstraction) :该角色位于属于"类的功能层次结构” 的最上层，用于定义抽象接口，一般是抽象类而不是抽象接口。其内部往往包含一个实现类接口实例(Implementor)，使用委托方式进行内部调用;

●改善后的抽象化,或者叫补充抽象类(RefinedAbstraction) :该角色用于补充
Abstraction功能而存在，通常情况下不再是抽象类而是具体的实现类，在内部可以直接调用Implementor中的业务方法;

●实现者(Implementor) :该角色位于"类的实现层次结构” 的最上层，定义了用于实现Abstraction角色的接口(API) ，这里的接口并非要和Abstraction中定义的完全一致,Implementor只对这些接口进行声明，具体实现还是要交给子类。通过委托，在
Abstraction中,不仅可以调用自己方法,还可以调用到Implementor中定义的方法;

●具体实现者(Concretelmplementor) :该角色用于实现Implementor角色中定义的接口，不同的实现类提供不同的业务处理方法，程序运行时，Concretelmplementor 将替换Abstraction中的Implementor,提供给抽象类具体的业务操作方法。

##### 实例

我们的手机上都安装了各种各样的软件，我们如何表示“在手机上运行软件”这一活动？手机的品牌和软件是两个维度上的东西，我们将其类比成功能层级和实现层级。类图如下

![](https://img1.sycdn.imooc.com/5dbac89600014e8710240355.png)

测试方法

```go
package StructuralType

import "testing"

func TestMobilePhone_Run(t *testing.T) {
	h1 := HuaWei{MobilePhone{ChatSoft{}}}
	h2 := HuaWei{MobilePhone{GameSoft{}}}
	h1.Run()
	h2.Run()
	h2.GPUTurbo()
	m1 := XiaoMi{MobilePhone{ChatSoft{}}}
	m2 := XiaoMi{MobilePhone{GameSoft{}}}
	m1.Run()
	m2.Run()
	m2.GameTurbo()
}
```

测试结果

```
=== RUN   TestMobilePhone_Run
聊天软件启动
游戏软件启动
GPUTurbo started
聊天软件启动
游戏软件启动
GameTurbo started.
--- PASS: TestMobilePhone_Run (0.00s)
PASS
```

由于m1、m2都属于MobilePhoneAbstraction的类的实例，因此我们可以调用它们的run方法，而m1、m2又都属于HuaWeiMobilePhone的实例，所以我们还能调用其gpuTurborRun方法，比如后期我们run实现有所变动，那么只需要修改main方法和Softlmplementor的实现类即可，其它代码就可以保持不动了。

通过MobilePhoneAbstraction中的impl桥接，就实现了抽象与行为实现的分离，这种就是桥接模式的存在意义。

##### 优缺点

桥接模式的优缺点总结如下:
●抽象与实现相分离:抽象与实现相分离，从而让抽象与实现分别独立开来，分别定义接口，有助于系统分层及产生更好的结构化系统;

●更好的拓展性:系统拓展时，因为抽象与实现已经分别独立，所以可以进行分别拓展不会相互影响，从而大大提高系统拓展性。

##### 总结

![](https://img1.sycdn.imooc.com/5db1471a00016b9007460219.png)

### 行为型模式
