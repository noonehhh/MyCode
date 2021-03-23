### array和slice

#### array

~~~
array是固定长度的数组，使用前必须确定数组长度
~~~

##### 创建

~~~go
arr1 := [5]int{1,2,3,4,5}

var arr2 [5]int
~~~

##### 底层

数组就是一片连续的内存

##### 特点

>1. `golang`中的数组是**`值类型`**,也就是说，如果你将一个数组赋值给另外一个数组，那么，实际上就是整个数组拷贝了一份
>2. 如果`golang`中的数组作为函数的参数，那么实际传递的参数是一份数组的拷贝，而不是数组的指针
>3. `array`的长度也是`Type`的一部分，这样就说明`[10]int`和`[20]int`是不一样的。

#### slice

~~~
slice是一个引用类型，是一个动态的指向数组切片的指针。
slice是一个不定长的，总是指向底层的数组array的数据结构。
~~~

##### 创建

~~~
var sli1 []int

sli2 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

sli3 := make([]int,10)
~~~

##### 底层

`slice` 实际上是一个结构体，包含三个字段：长度、容量、底层数组。

~~~go
//slice.go
type slice struct {
	array unsafe.Pointer // 元素指针
	len   int // 长度 
	cap   int // 容量
}
~~~

##### 注意

底层数组是可以被多个 `slice` 同时指向的，因此对一个 `slice` 的元素进行操作是有可能影响到其他 `slice` 的。

##### 扩容机制（重点）

>每一次扩容空间，都会重新申请一块区域，把旧空间里面的元素复制进来到新空间里，然后把新的元素追加进来。旧空间里面的元素等着垃圾回收。

###### 如何扩容？

网络上普遍结论

>1. 当原 slice 容量小于 `1024` 的时候，新 slice 容量变成原来的 `2` 倍；
>2. 原 slice 容量超过 `1024`，新 slice 容量变成原来的`1.25`倍。

###### 结论1

~~~go
func main(){
	sli := []int{1, 2}

	fmt.Println(len(sli),cap(sli))
	sli = append(sli, 3)
	fmt.Println(len(sli),cap(sli))
}

输出：
2，2
3，4
~~~

证实结论1正确，是这样吗？让我们再看一段代码

~~~
func main(){
	sli := []int{1, 2}

	fmt.Println(len(sli),cap(sli))
	sli = append(sli, 3, 4, 5)
	fmt.Println(len(sli),cap(sli))
}
输出：
2 2
5 6
~~~

结论1不正确吗？也不是，原因如下：

~~~
上面例子中newCap = 5,int数组所占字节为5*8 = 40
但 go 语言向内存管理模块向操作系统申请的内存容量却没有 40 大小的，只有48符合，于是newCap = 48/8 = 6

go语言内存管理模块是16bytes叠加的，8，16，32，48，64，80，96
~~~

###### 结论2

~~~go
package main

import "fmt"

func main() {
	s := make([]int, 0)

	oldCap := cap(s)

	for i := 0; i < 2048; i++ {
		s = append(s, i)

		newCap := cap(s)

		if newCap != oldCap {
			fmt.Printf("[%d -> %4d] cap = %-4d  |  after append %-4d  cap = %-4d\n", 0, i-1, oldCap, i, newCap)
			oldCap = newCap
		}
	}
}

输出：
1  [0 ->   -1] cap = 0     |  after append 0     cap = 1   
2  [0 ->    0] cap = 1     |  after append 1     cap = 2   
3  [0 ->    1] cap = 2     |  after append 2     cap = 4   
4  [0 ->    3] cap = 4     |  after append 4     cap = 8   
5  [0 ->    7] cap = 8     |  after append 8     cap = 16  
6  [0 ->   15] cap = 16    |  after append 16    cap = 32  
7  [0 ->   31] cap = 32    |  after append 32    cap = 64  
8  [0 ->   63] cap = 64    |  after append 64    cap = 128 
9  [0 ->  127] cap = 128   |  after append 128   cap = 256 
10 [0 ->  255] cap = 256   |  after append 256   cap = 512 
11 [0 ->  511] cap = 512   |  after append 512   cap = 1024
12 [0 -> 1023] cap = 1024  |  after append 1024  cap = 1280
13 [0 -> 1279] cap = 1280  |  after append 1280  cap = 1696
14 [0 -> 1695] cap = 1696  |  after append 1696  cap = 2304
~~~

>当向 slice 中添加元素 `1280` 的时候，老 slice 的容量为 `1280`，之后变成了 `1696`，两者并不是 `1.25` 倍的关系（1696/1280=1.325）。添加完 `1696` 后，新的容量 `2304` 当然也不是 `1696` 的 `1.25` 倍。

结论2也不正确吗？看源码，向 `slice` 追加元素的时候，若容量不够，会调用 `growslice` 函数，所以我们直接看它的代码。

~~~go
// go 1.9.5 src/runtime/slice.go:82
func growslice(et *_type, old slice, cap int) slice {
    // ……
    newcap := old.cap
	doublecap := newcap + newcap
	if cap > doublecap {
		newcap = cap
	} else {
		if old.len < 1024 {
			newcap = doublecap
		} else {
			for newcap < cap {
				newcap += newcap / 4
			}
		}
	}
	// ……
	
	capmem = roundupsize(uintptr(newcap) * ptrSize)
	newcap = int(capmem / ptrSize)
}
~~~

**解析**

~~~
如果只看前半部分，现在网上各种文章里说的 newcap 的规律是对的。现实是，后半部分还对 newcap 作了一个内存对齐，
这个和内存分配策略相关。进行内存对齐之后，新 slice 的容量是要 大于等于 老 slice 容量的 2倍或者1.25倍。
~~~

参考：

https://jodezer.github.io/2017/05/golangSlice%E7%9A%84%E6%89%A9%E5%AE%B9%E8%A7%84%E5%88%99

https://golang.design/go-questions/slice/grow/

https://www.cnblogs.com/Kingram/p/13630313.html