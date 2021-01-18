### **1. struct 能不能被比较？**

##### 补充知识：

* golang中的可比较类型：**Integer**，**Floating-point**，**String**，**Boolean**，**Complex(复数型)**，**Pointer**，**Channel**，**Interface**，**Array**

* 不可比较类型：**Slice**，**Map**，**Function**

### 1.1同一结构体的不同实例

#### 测试

~~~go
type Test struct {
   a int
   b string
}

func main(){
   test1 := Test{a:1, b:"2"}
   test2 := Test{a:1, b:"2"}
   fmt.Println(test1 == test2)
}
~~~

#### 结果

* 输出：true



#### 添加不可比较变量

~~~go
type Test struct {
   a int
   b string
   arr []int
}

func main(){
   test1 := Test{a:1, b:"2"}
   test2 := Test{a:1, b:"2"}
   fmt.Println(test1 == test2)
}
~~~

#### 报错

* 不能编译

#### 总结

* 同一个结构体的两个实例可比较也不可比较，当结构不包含不可直接比较成员变量时可直接比较，否则不可直接比较。

### 如果需要对两个结构体进行比较，可以用reflect.DeepEqual 

reflect.DeepEqual 是如何对变量进行比较的呢？

不同类型的值永远不会深度相等

当两个数组的元素对应深度相等时，两个数组深度相等

当两个相同结构体的所有字段对应深度相等的时候，两个结构体深度相等

当两个函数都为nil时，两个函数深度相等，其他情况不相等（相同函数也不相等）

当两个interface的真实值深度相等时，两个interface深度相等

map的比较需要同时满足以下几个

- 两个map都为nil或者都不为nil，并且长度要相等
- 相同的map对象或者所有key要对应相同
- map对应的value也要深度相等

指针，满足以下其一即是深度相等

- 两个指针满足go的==操作符
- 两个指针指向的值是深度相等的

切片，需要同时满足以下几点才是深度相等

- 两个切片都为nil或者都不为nil，并且长度要相等
- 两个切片底层数据指向的第一个位置要相同或者底层的元素要深度相等
- 注意：空的切片跟nil切片是不深度相等的

其他类型的值（numbers, bools, strings, channels）如果满足go的==操作符，则是深度相等的。要注意不是所有的值都深度相等于自己，例如函数，以及嵌套包含这些值的结构体，数组等。



### **2.两个不同的struct能不能被比较？**

#### 可以比较，也不可以比较

#### 可通过强制转换来比较

~~~go
type T2 struct {
    Name  string
    Age   int
    Arr   [2]bool
    ptr   *int
}

type T3 struct {
    Name  string
    Age   int
    Arr   [2]bool
    ptr   *int
}

func main() {
    var ss1 T2
    var ss2 T3
    // Cannot use 'ss2' (type T3) as type T2 in assignment
    //ss1 = ss2     // 不同结构体之间是不可以赋值的
    ss3 := T2(ss2)
    fmt.Println(ss3==ss1) // true
}
~~~

#### 如果成员变量中含有不可比较成员变量，即使可以强制转换，也不可以比较

~~~go
type T2 struct {
    Name  string
    Age   int
    Arr   [2]bool
    ptr   *int
    map1  map[string]string
}

type T3 struct {
    Name  string
    Age   int
    Arr   [2]bool
    ptr   *int
    map1  map[string]string
}

func main() {
    var ss1 T2
    var ss2 T3
    
    ss3 := T2(ss2)
    fmt.Println(ss3==ss1)   // 含有不可比较成员变量
}
~~~

### 3.struct可以作为map的key吗？

struct必须是可比较的，才能作为key，否则编译时报错

~~~go
type T1 struct {
    Name  string
    Age   int
    Arr   [2]bool
    ptr   *int
    slice []int
    map1  map[string]string
}

type T2 struct {
    Name string
    Age  int
    Arr  [2]bool
    ptr  *int
}

func main() {
    // n := make(map[T2]string, 0) // 无报错
    // fmt.Print(n)                // map[]

    m := make(map[T1]string, 0)
    fmt.Println(m) // invalid map key type T1
}
~~~

