### map

#### 内部结构

`golang map`的内部结构如下：

~~~go
type hmap struct {
    // 元素个数，调用 len(map) 时，直接返回此值
	count     int
	flags     uint8
	// buckets 的对数 log_2
	B         uint8
	// overflow 的 bucket 近似数
	noverflow uint16
	// 计算 key 的哈希的时候会传入哈希函数
	hash0     uint32
    // 指向 buckets 数组，大小为 2^B
    // 如果元素个数为0，就为 nil
	buckets    unsafe.Pointer
	// 扩容的时候，buckets 长度会是 oldbuckets 的两倍
	oldbuckets unsafe.Pointer
	// 指示扩容进度，小于此地址的 buckets 迁移完成
	nevacuate  uintptr
	extra *mapextra // optional fields
}
~~~

#### 字段释义

| 字段         | 解释                                                         |
| ------------ | ------------------------------------------------------------ |
| `count`      | 键值对的数量                                                 |
| `B`          | `2^B=len(buckets)` 表示当前哈希表持有的 `buckets` 数量，但是因为哈希表中桶的数量都 2 的倍数，所以该字段会存储对数 |
| `hash0`      | `hash`因子，它能为哈希函数的结果引入随机性，这个值在创建哈希表时确定，并在调用哈希函数时作为参数传入； |
| `buckets`    | 指向一个数组(连续内存空间)，数组的类型为`[]bmap`，`bmap`类型就是存在键值对的结构下面会详细介绍，这个字段我们可以称之为正常桶。**如下图所示** |
| `oldbuckets` | 是哈希在扩容时用于保存之前 `buckets` 的字段，它的大小是当前 `buckets` 的一半； |
| `extra`      | 溢出桶结构，正常桶里面某个`bmap`存满了，会使用这里面的内存空间存放键值对 |
| `noverflow`  | 溢出桶里`bmap`大致的数量                                     |
| `nevacuate`  | 分流次数，成倍扩容分流操作计数的字段(`Map`扩容相关字段)      |
| `flags`      | 状态标识，比如正在被写、`buckets`和`oldbuckets`在被遍历、等量扩容(`Map`扩容相关字段) |

#### 字段 buckets

![](https://github.com/No8LaVine/MyCode/blob/master/images/map1.png)

`buckets` 是一个指针，最终它指向的是一个结构体 `bmap`：

~~~go
type bmap struct {
	tophash [bucketCnt]uint8
    data    byte[1]  //key value数据:key/key/key/.../value/value/value...
	overflow *bmap   //溢出bucket的地址
}
~~~

##### 整体图

![](https://github.com/No8LaVine/MyCode/blob/master/images/map2.png)

`bmap`编译期间会动态生成一个新结构

~~~go
type bmap struct {
    topbits  [8]uint8
    keys     [8]keytype
    values   [8]valuetype
    pad      uintptr
    overflow uintptr
}
~~~

##### bmap

`bmap` 就是我们常说的“桶”，桶里面会最多装 8 个 `key`，这些 `key` 之所以会落入同一个桶，是因为它们经过哈希计算后，哈希结果是“一类”的。在桶内，又会根据 `key` 计算出来的 `hash` 值的高 8 位来决定 `key` 到底落入桶内的哪个位置（一个桶内最多有8个位置）。

###### 结构图

![](https://github.com/No8LaVine/MyCode/blob/master/images/map3.png)

#### 扩容

##### 溢出桶

每个`bmap`最多放8个键值对，当 `buckets` 里的 `bmap` 放满了怎么办？

引入`hmap.extra`结构:

~~~go
type mapextra struct {
    overflow    *[]*bmap
    oldoverflow *[]*bmap
    nextOverflow *bmap
}
~~~

![](https://github.com/No8LaVine/MyCode/blob/master/images/map4.png)

###### 字段释义

| 字段           | 解释                                                         |
| -------------- | ------------------------------------------------------------ |
| `overflow`     | 称之为**溢出桶**。和`hmap.buckets`的类型一样也是数组`[]bmap`，当正常桶`bmap`存满了的时候就使用`hmap.extra.overflow`的`bmap`。所以这里有个问题正常桶`hmap.buckets`里的`bmap`是怎么关联上溢出桶`hmap.extra.overflow`的`bmap`呢？我们下面说。 |
| `oldoverflow`  | 扩容时存放之前的`overflow`(`Map`扩容相关字段)                |
| `nextoverflow` | 指向溢出桶里下一个可以使用的`bmap`                           |

`hmap.extra.overflow`就是溢出桶，用于存放`bmap`溢出后的键值对

###### 正常桶`hmap.buckets`里的`bmap`是**怎么关联上**溢出桶`hmap.extra.overflow`的`bmap`的？

回顾本文上边说的`bmap`结构体，`bmap`结构体里有一个字段`overflow`，为指针类型，存放了对应使用的溢出桶`hmap.extra.overflow`里的`bmap`的地址。

##### 负载因子

负载因子用于衡量一个哈希表冲突情况，公式为：

~~~
loadFactor := count / (2^B)
count 是键值对数量，2^B是 buckets 的数量，上面的公式也就是：
负载因子 = 键数量/bucket数量

例如，对于一个bucket数量为4，包含4个键值对的哈希表来说，这个哈希表的负载因子为1
~~~

##### 触发扩容的条件

>1. 负载因子超过阈值，源码里定义的阈值是 6.5
>2. 哈希使用了太多溢出桶，overflow 的 bucket 数量过多：
>
>​              当 B < 15，也就是 bucket 总数 2^B 小于 2^15 时，如果 overflow 的 bucket 数量超过 2^B；
>
>​              当 B >= 15，也就是 bucket 总数 2^B 大于等于 2^15，如果 overflow 的 bucket 数量超过 2^15。

##### 增量扩容

当负载因子过大时，就新建一个`bucket`，新的`bucket`长度是原来的2倍，然后旧`bucket`数据搬迁到新的`bucket`。考虑到如果`map`存储了数以亿计的

`key-value`，一次性搬迁将会造成比较大的延时，`Go`采用逐步搬迁策略，即每次访问`map`时都会触发一次搬迁，每次搬迁 2 个键值对。

###### 步骤

>1. 新建一个 bucket，容量是原来 bucket 的 2 倍
>2. hmap 数据结构中 oldbuckets 指向 原bucket，buckets字段指向新的bucket
>3. 迁移数据，将oldbuckets中的键值对逐步的搬迁过来。当 oldbuckets 中的键值对全部搬迁完毕后，清空 oldbuckets。

例子：一个`bucket`满载的`map`(为了描述方便，图中`bucket`省略了`value`区域)

![](https://github.com/No8LaVine/MyCode/blob/master/images/map6.jpg)

有1个`bucket`。此地负载因子为7。再次插入数据时将会触发扩容操作，扩容之后再将新插入键写入新的`bucket`。

![](https://github.com/No8LaVine/MyCode/blob/master/images/map7.jpg)

##### 等量扩容

等量扩容，实际上并不是扩大容量，`buckets`数量不变，重新做一遍类似增量扩容的搬迁动作，把松散的键值对重新排列一次，以使`bucket`的使用率更高，进而保证更快的存取。
在极端场景下，比如不断的增删，而键值对正好集中在一小部分的`bucket`，这样会造成`overflow`的`bucket`数量增多，但负载因子又不高，从而无法执行增量搬迁的情况，如下图所示：

![](https://github.com/No8LaVine/MyCode/blob/master/images/map5.jpg)

可见，`overflow`的`bucket`中大部分是空的，访问效率会很差。此时进行一次等量扩容，即`buckets`数量不变，经过重新组织后`overflow`的`bucket`数量会减少，即节省了空间又会提高访问效率。

#### 查找

>1. 根据 key 值算出 hash 值
>2. 取哈希值低位与hmpa.B取模确定bucket位置
>3. 取哈希值高位在tophash数组中查询
>4. 如果tophash[i]中存储值也哈希值相等，则去找到该bucket中的key值进行比较
>5. 当前bucket没有找到，则继续从下个overflow的bucket中查找。
>6. 如果当前处于搬迁过程，则优先从oldbuckets查找

#### 插入

~~~
核心还是一个双层循环，外层遍历 bucket 和它的 overflow bucket，内层遍历整个 bucket 的各个 cell
~~~

>1. 首先会检查 map 的标志位 flags。如果 flags 的写标志位此时被置 1 了，说明有其他协程在执行“写”操作，进而导致程序 panic。
>2. 根据 key 值算出哈希值
>3. 取哈希值低位与 hmap.B 取模确定 bucke 位置
>4. 如果 bucket 在 oldbucket中，确保oldbucket完成了迁移过程，将其重新散列到 new bucket中。
>5. 定位自己的位置，两个指针，一个（`inserti`）指向 key 的 hash 值在 tophash 数组所处的位置，另一个(`insertk`)指向 cell 的位置（也就是 key 最终放置的地址）
>6. 在循环的过程中，inserti 和 insertk 分别指向第一个找到的空闲的 cell。如果之后在 map 没有找到 key 的存在，也就是说原来 map 中没有此 key，这意味着插入新 key。那最终 key 的安置地址就是第一次发现的“空位”（tophash 是 empty）。
>7. 如果这个 bucket 的 8 个 key 都已经放置满了，这时候需要在 bucket 后面挂上 overflow bucket。当然，也有可能是在 overflow bucket 后面再挂上一个 overflow bucket。这就说明，太多 key hash 到了此 bucket。
>8. 在正式安置 key 之前，还要检查 map 的状态，看它是否需要进行扩容。如果满足扩容的条件，就主动触发一次扩容操作。
>9. 这之后，整个之前的查找定位 key 的过程，还得再重新走一次。因为扩容之后，key 的分布都发生了变化。
>10. hmap.count 加 1

#### 删除

计算 key 的哈希，找到落入的 bucket。检查此 map 如果正在扩容的过程中，直接触发一次搬迁操作。

删除操作同样是两层循环，核心还是找到 key 的具体位置。寻找过程都是类似的，在 bucket 中挨个 cell 寻找。最后，将 count 值减 1，将对应位置的 tophash 值置成 `Empty`。



参考：

https://my.oschina.net/renhc/blog/2208417

https://draveness.me/golang/docs/part2-foundation/ch03-datastructure/golang-hashmap/#%E6%89%A9%E5%AE%B9

https://qcrao91.gitbook.io/go/map/map-de-kuo-rong-guo-cheng-shi-zen-yang-de

https://tiancaiamao.gitbooks.io/go-internals/content/zh/02.3.html

https://segmentfault.com/a/1190000039101378

https://golang.design/go-questions/map/extend/