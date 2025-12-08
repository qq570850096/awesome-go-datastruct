## 排序算法合集（Sorting Algorithms）

### 概述

**排序**是计算机科学中最基础、最重要的算法之一。本模块实现了多种经典排序算法，涵盖了从简单到高效的各类算法，便于学习对比和实际应用。

**收录算法：**
- **O(n²) 算法**：冒泡排序、插入排序、选择排序
- **分治算法**：归并排序（递归+自底向上）、快速排序（单路/双路/三路）
- **其他算法**：希尔排序、桶排序

### 算法复杂度对比

| 算法 | 平均时间 | 最坏时间 | 最好时间 | 空间 | 稳定性 |
|------|---------|---------|---------|------|--------|
| 冒泡排序 | O(n²) | O(n²) | O(n) | O(1) | ✅ 稳定 |
| 选择排序 | O(n²) | O(n²) | O(n²) | O(1) | ❌ 不稳定 |
| 插入排序 | O(n²) | O(n²) | O(n) | O(1) | ✅ 稳定 |
| 希尔排序 | O(n^1.3) | O(n²) | O(n) | O(1) | ❌ 不稳定 |
| 归并排序 | O(n log n) | O(n log n) | O(n log n) | O(n) | ✅ 稳定 |
| 快速排序 | O(n log n) | O(n²) | O(n log n) | O(log n) | ❌ 不稳定 |
| 堆排序 | O(n log n) | O(n log n) | O(n log n) | O(1) | ❌ 不稳定 |
| 桶排序 | O(n+k) | O(n²) | O(n) | O(n+k) | ✅ 稳定 |

### 适用场景

| 场景 | 推荐算法 | 原因 |
|------|---------|------|
| 小规模数据 (n < 50) | 插入排序 | 常数因子小，简单高效 |
| 近乎有序数据 | 插入排序 | 最好情况 O(n) |
| 大规模数据 | 快速排序/归并排序 | O(n log n) |
| 重复元素多 | 三路快排 | 避免退化 |
| 需要稳定性 | 归并排序 | 稳定且高效 |
| 内存受限 | 堆排序/快速排序 | O(1) 或 O(log n) 空间 |
| 整数范围已知 | 计数排序/桶排序 | O(n) 线性时间 |

### 归并排序（Merge Sort）

分治思想：将数组分成两半，分别排序后合并。

```go
// merge 将 arr[l...mid] 和 arr[mid+1...r] 两部分进行归并
func merge(arr, aux []int, l, mid, r int) {
    // 复制到辅助数组
    aux = make([]int, r-l+1)
    for i := l; i <= r; i++ {
        aux[i-l] = arr[i]
    }

    // 归并过程
    i, j := l, mid+1
    for k := l; k <= r; k++ {
        if i > mid {
            arr[k] = aux[j-l]
            j++
        } else if j > r {
            arr[k] = aux[i-l]
            i++
        } else if aux[i-l] < aux[j-l] {
            arr[k] = aux[i-l]
            i++
        } else {
            arr[k] = aux[j-l]
            j++
        }
    }
}

// MergeSort 递归归并排序（自顶向下）
func MergeSort(arr []int, l, r int) {
    if l >= r {
        return
    }
    mid := (r + l) / 2
    MergeSort(arr, l, mid)
    MergeSort(arr, mid+1, r)
    // 优化：只有当左半部分最大值 > 右半部分最小值时才归并
    if arr[mid] > arr[mid+1] {
        var aux []int
        merge(arr, aux, l, mid, r)
    }
}

// MergeSortBU 自底向上归并排序（非递归）
func MergeSortBU(arr []int, n int) {
    aux := make([]int, n)
    // sz: 每次归并的子数组大小，从 1 开始翻倍
    for sz := 1; sz <= n; sz += sz {
        for i := 0; i < n-sz; i += sz + sz {
            if arr[i+sz-1] > arr[i+sz] {
                merge(arr, aux, i, i+sz-1, min(i+sz+sz-1, n-1))
            }
        }
    }
}
```

### 快速排序（Quick Sort）

分治思想：选择基准元素，将数组分为小于和大于基准的两部分。

#### 单路快排

```go
// partition 分区操作
// 返回 p，使得 arr[l...p-1] < arr[p] < arr[p+1...r]
func partition(arr []int, l, r int) int {
    // 随机选择基准，避免最坏情况
    rand.Seed(time.Now().Unix())
    randIndex := rand.Int()%(r-l+1) + l
    arr[l], arr[randIndex] = arr[randIndex], arr[l]

    v := arr[l]
    j := l
    for i := l + 1; i <= r; i++ {
        if arr[i] < v {
            j++
            arr[j], arr[i] = arr[i], arr[j]
        }
    }
    arr[l], arr[j] = arr[j], arr[l]
    return j
}

// QuickSort 快速排序
func QuickSort(arr []int, l, r int) {
    if l >= r {
        return
    }
    p := partition(arr, l, r)
    QuickSort(arr, l, p-1)
    QuickSort(arr, p+1, r)
}
```

#### 双路快排

处理大量重复元素时更高效：

```go
// partition2 双路快排分区
func partition2(arr []int, l, r int) int {
    rand.Seed(time.Now().Unix())
    randIndex := rand.Int()%(r-l+1) + l
    arr[l], arr[randIndex] = arr[randIndex], arr[l]

    v := arr[l]
    // 不变式：arr[l+1...i) <= v; arr(j...r] >= v
    i, j := l+1, r
    for {
        for i <= r && arr[i] < v {
            i++
        }
        for j >= l+1 && arr[j] > v {
            j--
        }
        if i > j {
            break
        }
        arr[j], arr[i] = arr[i], arr[j]
        i++
        j--
    }
    arr[l], arr[j] = arr[j], arr[l]
    return j
}
```

#### 三路快排

将数组分为三部分：小于、等于、大于基准：

```go
// quickSort3Ways 三路快速排序
// 将数组分为三部分：< v | == v | > v
func quickSort3Ways(arr []int, l, r int) {
    if l >= r {
        return
    }

    // 随机选择基准
    rand.Seed(time.Now().Unix())
    randIndex := rand.Int()%(r-l+1) + l
    arr[l], arr[randIndex] = arr[randIndex], arr[l]

    v := arr[l]
    // lt: 小于区间的右边界 arr[l+1...lt] < v
    // gt: 大于区间的左边界 arr[gt...r] > v
    // i: 当前遍历位置 arr[lt+1...i) == v
    lt, gt, i := l, r+1, l+1

    for i < gt {
        if arr[i] < v {
            arr[i], arr[lt+1] = arr[lt+1], arr[i]
            i++
            lt++
        } else if arr[i] > v {
            arr[i], arr[gt-1] = arr[gt-1], arr[i]
            gt--
        } else {
            i++
        }
    }
    arr[l], arr[lt] = arr[lt], arr[l]

    quickSort3Ways(arr, l, lt-1)
    quickSort3Ways(arr, gt, r)
}

// Quick3Ways 三路快排入口
func Quick3Ways(arr []int, n int) {
    quickSort3Ways(arr, 0, n-1)
}
```

### 插入排序（Insert Sort）

适合小规模或近乎有序的数据：

```go
func InsertSort(arr []int) {
    for i := 1; i < len(arr); i++ {
        // 将 arr[i] 插入到已排序部分的正确位置
        temp := arr[i]
        j := i
        for j > 0 && arr[j-1] > temp {
            arr[j] = arr[j-1]
            j--
        }
        arr[j] = temp
    }
}
```

### 希尔排序（Shell Sort）

插入排序的改进版，通过递减增量实现：

```go
func ShellSort(arr []int) {
    n := len(arr)
    // gap 从 n/2 开始，每次减半
    for gap := n / 2; gap > 0; gap /= 2 {
        // 对每个 gap 进行插入排序
        for i := gap; i < n; i++ {
            temp := arr[i]
            j := i
            for j >= gap && arr[j-gap] > temp {
                arr[j] = arr[j-gap]
                j -= gap
            }
            arr[j] = temp
        }
    }
}
```

### 冒泡排序（Bubble Sort）

通过相邻元素比较交换实现：

```go
func BubbleSort(arr []int) {
    n := len(arr)
    for i := 0; i < n-1; i++ {
        swapped := false
        for j := 0; j < n-1-i; j++ {
            if arr[j] > arr[j+1] {
                arr[j], arr[j+1] = arr[j+1], arr[j]
                swapped = true
            }
        }
        // 优化：如果没有交换，说明已经有序
        if !swapped {
            break
        }
    }
}
```

### 选择排序（Select Sort）

每次选择最小元素放到已排序部分末尾：

```go
func SelectSort(arr []int) {
    n := len(arr)
    for i := 0; i < n-1; i++ {
        minIdx := i
        for j := i + 1; j < n; j++ {
            if arr[j] < arr[minIdx] {
                minIdx = j
            }
        }
        arr[i], arr[minIdx] = arr[minIdx], arr[i]
    }
}
```

### 测试用例

```go
func TestSortAlgorithms(t *testing.T) {
    // 测试数据
    arr := []int{64, 34, 25, 12, 22, 11, 90}

    // 测试归并排序
    arr1 := make([]int, len(arr))
    copy(arr1, arr)
    MergeSort(arr1, 0, len(arr1)-1)
    fmt.Println("归并排序:", arr1) // [11 12 22 25 34 64 90]

    // 测试快速排序
    arr2 := make([]int, len(arr))
    copy(arr2, arr)
    QuickSort(arr2, 0, len(arr2)-1)
    fmt.Println("快速排序:", arr2) // [11 12 22 25 34 64 90]

    // 测试三路快排
    arr3 := make([]int, len(arr))
    copy(arr3, arr)
    Quick3Ways(arr3, len(arr3))
    fmt.Println("三路快排:", arr3) // [11 12 22 25 34 64 90]
}

// BenchmarkSort 性能基准测试
func BenchmarkSort(b *testing.B) {
    n := 10000
    arr := make([]int, n)
    for i := 0; i < n; i++ {
        arr[i] = rand.Intn(n)
    }

    b.Run("MergeSort", func(b *testing.B) {
        for i := 0; i < b.N; i++ {
            tmp := make([]int, n)
            copy(tmp, arr)
            MergeSort(tmp, 0, n-1)
        }
    })

    b.Run("QuickSort", func(b *testing.B) {
        for i := 0; i < b.N; i++ {
            tmp := make([]int, n)
            copy(tmp, arr)
            QuickSort(tmp, 0, n-1)
        }
    })
}
```

### 运行方式

```bash
# 运行测试
go test ./Sorts

# 运行基准测试
go test ./Sorts -bench .
```

### LeetCode 实战

#### [912. 排序数组](https://leetcode-cn.com/problems/sort-an-array/)

给定整数数组，按升序排序：

```go
// 使用三路快排实现
func sortArray(nums []int) []int {
    quickSort3Ways(nums, 0, len(nums)-1)
    return nums
}

func quickSort3Ways(arr []int, l, r int) {
    if l >= r {
        return
    }

    // 随机选择基准
    randIdx := l + rand.Intn(r-l+1)
    arr[l], arr[randIdx] = arr[randIdx], arr[l]

    v := arr[l]
    lt, gt, i := l, r+1, l+1

    for i < gt {
        if arr[i] < v {
            arr[i], arr[lt+1] = arr[lt+1], arr[i]
            i++
            lt++
        } else if arr[i] > v {
            arr[i], arr[gt-1] = arr[gt-1], arr[i]
            gt--
        } else {
            i++
        }
    }
    arr[l], arr[lt] = arr[lt], arr[l]

    quickSort3Ways(arr, l, lt-1)
    quickSort3Ways(arr, gt, r)
}
```

#### [215. 数组中的第K个最大元素](https://leetcode-cn.com/problems/kth-largest-element-in-an-array/)

使用快速选择算法（Quick Select）：

```go
func findKthLargest(nums []int, k int) int {
    // 第 k 大 = 第 n-k+1 小 = 索引 n-k
    targetIdx := len(nums) - k
    return quickSelect(nums, 0, len(nums)-1, targetIdx)
}

func quickSelect(arr []int, l, r, k int) int {
    if l == r {
        return arr[l]
    }

    // 随机选择基准
    randIdx := l + rand.Intn(r-l+1)
    arr[l], arr[randIdx] = arr[randIdx], arr[l]

    pivot := arr[l]
    i, j := l, r
    for i < j {
        for i < j && arr[j] >= pivot {
            j--
        }
        arr[i] = arr[j]
        for i < j && arr[i] <= pivot {
            i++
        }
        arr[j] = arr[i]
    }
    arr[i] = pivot

    if i == k {
        return arr[i]
    } else if i < k {
        return quickSelect(arr, i+1, r, k)
    } else {
        return quickSelect(arr, l, i-1, k)
    }
}
```

#### [148. 排序链表](https://leetcode-cn.com/problems/sort-list/)

使用归并排序对链表排序（O(n log n) 时间，O(1) 空间）：

```go
func sortList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    // 找中点（快慢指针）
    slow, fast := head, head.Next
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }

    // 分割
    mid := slow.Next
    slow.Next = nil

    // 递归排序
    left := sortList(head)
    right := sortList(mid)

    // 合并
    return mergeLists(left, right)
}

func mergeLists(l1, l2 *ListNode) *ListNode {
    dummy := &ListNode{}
    curr := dummy

    for l1 != nil && l2 != nil {
        if l1.Val <= l2.Val {
            curr.Next = l1
            l1 = l1.Next
        } else {
            curr.Next = l2
            l2 = l2.Next
        }
        curr = curr.Next
    }

    if l1 != nil {
        curr.Next = l1
    } else {
        curr.Next = l2
    }

    return dummy.Next
}
```

#### [75. 颜色分类](https://leetcode-cn.com/problems/sort-colors/)

三路快排思想（荷兰国旗问题）：

```go
func sortColors(nums []int) {
    // p0: 0 的右边界
    // p2: 2 的左边界
    // curr: 当前遍历位置
    p0, curr, p2 := 0, 0, len(nums)-1

    for curr <= p2 {
        if nums[curr] == 0 {
            nums[p0], nums[curr] = nums[curr], nums[p0]
            p0++
            curr++
        } else if nums[curr] == 2 {
            nums[p2], nums[curr] = nums[curr], nums[p2]
            p2--
            // 注意：curr 不前进，因为交换来的元素还需要判断
        } else {
            curr++
        }
    }
}
```
