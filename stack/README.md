## 栈（Stack）

### 定义

**栈**（Stack）是一种遵循**后进先出**（LIFO, Last In First Out）原则的线性数据结构。栈只允许在一端（栈顶）进行插入和删除操作，就像一摞盘子，最后放上去的盘子最先被拿走。

栈的核心操作：
1. **Push**：将元素压入栈顶
2. **Pop**：弹出栈顶元素
3. **Peek/Top**：查看栈顶元素（不弹出）

### 为什么使用栈？

栈因其简单而强大的特性，在计算机科学中有广泛应用：

**应用场景：**
- 函数调用栈（递归的底层实现）
- 表达式求值（中缀转后缀、后缀表达式计算）
- 括号匹配验证
- 浏览器前进/后退功能
- 撤销操作（Undo）
- 深度优先搜索（DFS）

**生活类比：** 弹夹装弹——最后装入的子弹最先射出；叠放的书本——最上面的书最先被拿走。

### 特性

| 操作 | 时间复杂度 | 说明 |
|------|-----------|------|
| Push | O(1) | 压入栈顶 |
| Pop | O(1) | 弹出栈顶 |
| Peek | O(1) | 查看栈顶 |
| IsEmpty | O(1) | 判断是否为空 |
| 空间复杂度 | O(n) | n 为元素个数 |

### 数据结构

```go
// Stack 顺序栈实现
type Stack struct {
    container []byte // 用来装元素的切片
    top       int    // 栈顶标记位
    size      int    // 容量限制
}

// 初始化时，要传入容量限制
func NewStack(size int) Stack {
    return Stack{
        container: make([]byte, size),
        top:       0,
        size:      size,
    }
}
```

### 核心方法实现

#### Push 入栈操作

```go
// Push 将元素压入栈顶
// 如果栈满返回 false，否则返回 true
func (s *Stack) Push(e byte) bool {
    if s.IsFull() {
        return false
    }
    s.container[s.top] = e
    s.top++
    return true
}
```

#### Pop 出栈操作

```go
// Pop 弹出栈顶元素
// 返回操作是否成功以及弹出的元素
func (s *Stack) Pop() (flag bool, ret byte) {
    // 如果栈是空的，那么就不能继续 Pop 了
    if s.IsEmpty() {
        return false, ret
    }
    ret = s.container[s.top-1]
    s.top--
    return true, ret
}
```

#### 辅助方法

```go
// IsEmpty 判断栈是否为空
func (s *Stack) IsEmpty() bool {
    return s.top == 0
}

// IsFull 判断栈是否已满
func (s *Stack) IsFull() bool {
    return s.top == s.size
}
```

### 应用示例：括号匹配

```go
// IsValid 使用栈校验括号字符串是否合法
func IsValid(s string) bool {
    stack := NewStack(100)
    // 遍历括号字符串
    for _, v := range s {
        if v == '(' {
            // 左括号入栈
            stack.Push(byte(v))
        }
        if v == ')' {
            // 右括号尝试匹配栈顶的左括号
            if flag, top := stack.Pop(); flag == true && top == '(' {
                continue
            } else {
                return false
            }
        }
    }
    // 字符串遍历完后如果栈也空了，说明括号匹配
    return stack.IsEmpty()
}
```

### 栈操作示意图

```
初始状态：
┌───┬───┬───┬───┬───┐
│   │   │   │   │   │  top = 0
└───┴───┴───┴───┴───┘

Push('A')：
┌───┬───┬───┬───┬───┐
│ A │   │   │   │   │  top = 1
└───┴───┴───┴───┴───┘
  ↑
 栈顶

Push('B'), Push('C')：
┌───┬───┬───┬───┬───┐
│ A │ B │ C │   │   │  top = 3
└───┴───┴───┴───┴───┘
          ↑
         栈顶

Pop() -> 'C'：
┌───┬───┬───┬───┬───┐
│ A │ B │   │   │   │  top = 2
└───┴───┴───┴───┴───┘
      ↑
     栈顶
```

### 测试用例

```go
func TestStack(t *testing.T) {
    // 创建容量为 5 的栈
    st := NewStack(5)

    // 测试入栈
    st.Push('A')
    st.Push('B')
    st.Push('C')
    fmt.Println("栈是否为空:", st.IsEmpty()) // false

    // 测试出栈
    _, val := st.Pop()
    fmt.Printf("弹出元素: %c\n", val) // C

    // 测试括号匹配
    fmt.Println("'(())' 是否合法:", IsValid("(())"))     // true
    fmt.Println("'(()' 是否合法:", IsValid("(()"))       // false
    fmt.Println("'())' 是否合法:", IsValid("())"))       // false
    fmt.Println("'()()' 是否合法:", IsValid("()()"))     // true
}
```

### 运行方式

```bash
go test ./stack
```

### LeetCode 实战

#### [20. 有效的括号](https://leetcode-cn.com/problems/valid-parentheses/)

判断字符串中的括号是否有效（支持多种括号类型）：

```go
func isValid(s string) bool {
    stack := make([]byte, 0)
    pairs := map[byte]byte{
        ')': '(',
        ']': '[',
        '}': '{',
    }

    for i := 0; i < len(s); i++ {
        if s[i] == '(' || s[i] == '[' || s[i] == '{' {
            // 左括号入栈
            stack = append(stack, s[i])
        } else {
            // 右括号匹配
            if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
                return false
            }
            stack = stack[:len(stack)-1]
        }
    }
    return len(stack) == 0
}
```

#### [155. 最小栈](https://leetcode-cn.com/problems/min-stack/)

设计支持常数时间获取最小元素的栈：

```go
type MinStack struct {
    stack    []int
    minStack []int // 辅助栈，存储当前最小值
}

func Constructor() MinStack {
    return MinStack{
        stack:    []int{},
        minStack: []int{math.MaxInt64},
    }
}

func (this *MinStack) Push(val int) {
    this.stack = append(this.stack, val)
    // 同步更新最小值栈
    top := this.minStack[len(this.minStack)-1]
    this.minStack = append(this.minStack, min(val, top))
}

func (this *MinStack) Pop() {
    this.stack = this.stack[:len(this.stack)-1]
    this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
    return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
    return this.minStack[len(this.minStack)-1]
}
```

#### [150. 逆波兰表达式求值](https://leetcode-cn.com/problems/evaluate-reverse-polish-notation/)

使用栈计算后缀表达式：

```go
func evalRPN(tokens []string) int {
    stack := make([]int, 0)
    for _, token := range tokens {
        val, err := strconv.Atoi(token)
        if err == nil {
            // 数字入栈
            stack = append(stack, val)
        } else {
            // 运算符：弹出两个操作数计算
            b := stack[len(stack)-1]
            a := stack[len(stack)-2]
            stack = stack[:len(stack)-2]
            switch token {
            case "+":
                stack = append(stack, a+b)
            case "-":
                stack = append(stack, a-b)
            case "*":
                stack = append(stack, a*b)
            case "/":
                stack = append(stack, a/b)
            }
        }
    }
    return stack[0]
}
```

#### [739. 每日温度](https://leetcode-cn.com/problems/daily-temperatures/)

使用单调栈找出下一个更大元素：

```go
func dailyTemperatures(temperatures []int) []int {
    n := len(temperatures)
    ans := make([]int, n)
    stack := make([]int, 0) // 存储下标

    for i := 0; i < n; i++ {
        // 当前温度大于栈顶温度，栈顶出栈并记录天数
        for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
            top := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            ans[top] = i - top
        }
        stack = append(stack, i)
    }
    return ans
}
```
