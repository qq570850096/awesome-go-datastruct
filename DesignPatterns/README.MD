## 设计模式 —— Go 语言实现

> **生活的理想，就是为了理想的生活。 —— 张闻天**

### 概述

本模块实现了 GoF 23 种经典设计模式的 Go 语言版本，每个模式包含源代码、测试文件和详细文档。

### 模式概览

| 类型 | 模式 | 数量 |
|-----|------|-----|
| 创造型模式 | 单例、原型、简单工厂、抽象工厂、建造者 | 5 |
| 结构型模式 | 享元、代理、装饰器、外观、桥接、组合、过滤、适配器 | 8 |
| 行为型模式 | 中介者、命令、备忘录、模板方法、状态、策略、观察者、解释器、责任链、迭代器 | 10 |

### 项目结构

```
DesignPatterns/
├── CreativeType/       # 创造型模式
│   ├── Docs/           # 文档目录
│   └── *.go            # 实现代码
├── StructuralType/     # 结构型模式
│   ├── Docs/           # 文档目录
│   └── *.go            # 实现代码
├── BehavioralType/     # 行为型模式
│   ├── Docs/           # 文档目录
│   └── *.go            # 实现代码
├── 七大设计原则.md      # 设计原则文档
└── Compound.go         # 综合案例（农场模拟）
```

---

### [创造型模式](./CreativeType/Docs)

创造型模式关注对象的创建机制，将对象的创建与使用分离。

| 模式 | 代码 | 文档 | 应用场景 |
|-----|------|------|---------|
| 单例模式 | [Singleton.go](./CreativeType/Singleton.go) | [文档](./CreativeType/Docs/单例模式.md) | 全局唯一实例 |
| 原型模式 | [Prototype.go](./CreativeType/Prototype.go) | [文档](./CreativeType/Docs/原型模式.md) | 对象克隆 |
| 简单工厂 | [SimpleFactory.go](./CreativeType/SimpleFactory.go) | [文档](./CreativeType/Docs/简单工厂模式.md) | 统一创建入口 |
| 抽象工厂 | [AbstractFactory.go](./CreativeType/AbstractFactory.go) | [文档](./CreativeType/Docs/抽象工厂.md) | 产品族创建 |
| 建造者模式 | [Builder.go](./CreativeType/Builder.go) | [文档](./CreativeType/Docs/建造者模式.md) | 复杂对象构建 |

---

### [结构型模式](./StructuralType/Docs)

结构型模式关注类和对象的组合，形成更大的结构。

| 模式 | 代码 | 文档 | 应用场景 |
|-----|------|------|---------|
| 享元模式 | [Flyweight.go](./StructuralType/Flyweight.go) | [文档](./StructuralType/Docs/享元模式.md) | 共享细粒度对象 |
| 代理模式 | [Proxy.go](./StructuralType/Proxy.go) | [文档](./StructuralType/Docs/代理模式.md) | 访问控制 |
| 装饰器模式 | [Decorator.go](./StructuralType/Decorator.go) | [文档](./StructuralType/Docs/修饰器模式.md) | 动态扩展功能 |
| 外观模式 | [Facade.go](./StructuralType/Facade.go) | [文档](./StructuralType/Docs/外观模式.md) | 简化复杂接口 |
| 桥接模式 | [Bridge.go](./StructuralType/Bridge.go) | [文档](./StructuralType/Docs/桥接模式.md) | 抽象与实现分离 |
| 组合模式 | [Component.go](./StructuralType/Component.go) | [文档](./StructuralType/Docs/组合模式.md) | 树形结构 |
| 过滤模式 | [Filter.go](./StructuralType/Filter.go) | [文档](./StructuralType/Docs/过滤模式.md) | 条件筛选 |
| 适配器模式 | [Wrapper.go](./StructuralType/Wrapper.go) | [文档](./StructuralType/Docs/适配器模式.md) | 接口转换 |

---

### [行为型模式](./BehavioralType/Docs)

行为型模式关注对象之间的通信和职责分配。

| 模式 | 代码 | 文档 | 应用场景 |
|-----|------|------|---------|
| 中介者模式 | [Mediator.go](./BehavioralType/Mediator.go) | [文档](./BehavioralType/Docs/中介者模式.md) | 解耦多对象交互 |
| 命令模式 | [Command.go](./BehavioralType/Command.go) | [文档](./BehavioralType/Docs/命令模式.md) | 请求封装 |
| 备忘录模式 | [Memento.go](./BehavioralType/Memento.go) | [文档](./BehavioralType/Docs/备忘录模式.md) | 状态保存与恢复 |
| 模板方法模式 | [Template.go](./BehavioralType/Template.go) | [文档](./BehavioralType/Docs/模板方法模式.md) | 算法骨架 |
| 状态模式 | [State.go](./BehavioralType/State.go) | [文档](./BehavioralType/Docs/状态模式.md) | 状态驱动行为 |
| 策略模式 | [Strategy.go](./BehavioralType/Strategy.go) | [文档](./BehavioralType/Docs/策略模式.md) | 算法族切换 |
| 观察者模式 | [Observer.go](./BehavioralType/Observer.go) | [文档](./BehavioralType/Docs/观察者模式.md) | 事件通知 |
| 解释器模式 | [Interpreter.go](./BehavioralType/Interpreter.go) | [文档](./BehavioralType/Docs/解释器模式.md) | DSL 解析 |
| 责任链模式 | [Handler.go](./BehavioralType/Handler.go) | [文档](./BehavioralType/Docs/责任链模式.md) | 链式处理 |
| 迭代器模式 | [Iterator.go](./BehavioralType/Iterator.go) | [文档](./BehavioralType/Docs/迭代器模式.md) | 顺序访问 |

---

### 七大设计原则

详见 [七大设计原则.md](./七大设计原则.md)

| 原则 | 说明 |
|-----|------|
| 单一职责原则（SRP） | 一个类只负责一项职责 |
| 开闭原则（OCP） | 对扩展开放，对修改关闭 |
| 里氏替换原则（LSP） | 子类可以替换父类 |
| 接口隔离原则（ISP） | 使用多个专门接口 |
| 依赖倒置原则（DIP） | 依赖抽象而非具体 |
| 迪米特法则（LoD） | 最少知识原则 |
| 合成复用原则（CRP） | 优先使用组合而非继承 |

---

### 综合案例：农场模拟

[Compound.go](./Compound.go) 演示了在需求变更场景下如何灵活运用多种设计模式。

---

### 运行方式

```bash
# 运行所有设计模式测试
go test ./DesignPatterns/...

# 运行特定类型模式测试
go test ./DesignPatterns/CreativeType
go test ./DesignPatterns/StructuralType
go test ./DesignPatterns/BehavioralType

# 运行单个模式测试
go test ./DesignPatterns/CreativeType -run TestSingleton
```

### 学习建议

1. **先学原则**：阅读七大设计原则，理解设计模式的指导思想
2. **从简到繁**：建议按 创造型 → 结构型 → 行为型 顺序学习
3. **代码为主**：每个模式先看代码实现，再读文档解释
4. **动手实践**：运行测试用例，修改代码观察效果
