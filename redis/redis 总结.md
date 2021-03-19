## `redis` 常见总结

### `redis`单线程与多线程

#### 一、单线程

##### redis为什么快？

> 1. 完全基于内存，绝大部分请求是纯粹的内存操作，非常快速。数据存在内存中
> 2. 数据结构简单，对数据操作也简单，Redis 中的数据结构是专门进行设计的；
> 3. 采用单线程，避免了不必要的上下文切换和竞争条件，也不存在多进程或者多线程导致的切换而消耗 CPU，不用去考虑各种锁的问题，不存在加锁释放锁操作，没有因为可能出现死锁而导致的性能消耗；
> 4. 使用多路 I/O 复用模型，非阻塞 IO；
> 5. 使用底层模型不同，它们之间底层实现方式以及与客户端之间通信的应用协议不一样，Redis 直接自己构建了 VM机制 ，因为一般的系统调用系统函数的话，会浪费一定的时间去移动和请求；

https://segmentfault.com/a/1190000037637112

##### redis为什么是单线程的？

> 1. 单线程编程容易并且更容易维护；
>
> 2. Redis的性能瓶颈不再 CPU ，主要在内存和网络；
>
> 3. 多线程就会存在死锁、线程上下文切换等问题，甚至会影响性能。
>
>    因为 redis 是基于内存的操作，CPU 不是 redis 的瓶颈，redis 的瓶颈有可能是机器的内存大小，或者网络带宽，就没必要引入多线程，因为引入多线程则意味着复杂化，cpu 调度、上下文切换、加锁等问题。

##### redis单线程如何保证并发竞争问题？

单进程单线程模式，采用队列模式将并发访问变为串行访问。Redis 本身没有锁的概念，Redis 对于多个客户端连接并不存在竞争，利用 setnx 实现锁。

#### 二、多线程

##### redis 6.0 多线程

###### redis 6.0之前真的是单线程吗？

~~~
Redis  在处理客户端的请求时，包括获取(Socket 读)、解析、执行、内容返回(Socket  写)等都由一个顺序串行的主线程处理，
这就是所谓的“单线程”。
~~~

但如果严格来讲从 **Redis 4.0** 之后并不是单线程，除了主线程外，它也有后台线程在处理一些较为缓慢的操作，
例如**清理脏数据、无用连接的释放、大Key的删除等等。**

###### 如何开启

redis.conf

~~~shell
io-threads-do-reads yes # 默认no
~~~

###### 设置线程数

~~~shell
io-threads 4
~~~

###### redis 6.0多线程实现机制

> 总体来说就是：主线程主要负责接收客户端连接，并且分发到各个 IO线程 ，而 IO线程负责读取客户端命令。命令读取完成后，由 主线程执行命令。 主线程执行完命令后，再由 IO线程把回复数据发送给客户端。

~~~
主线程负责接收建立连接请求，获取 socket 放入全局等待读处理队列

主线程处理完读事件之后，通过 RR(Round Robin) 将这些连接分配给这些 IO 线程

主线程阻塞等待 IO 线程读取 socket 完毕

主线程通过单线程的方式执行请求命令，请求数据读取并解析完成，但并不执行

主线程阻塞等待 IO 线程将数据回写 socket 完毕

解除绑定，清空等待队列
~~~

###### 开启多线程后是否存在线程并发安全问题

~~~
从上面的实现机制可以看出，Redis 的多线程部分只是用来处理网络数据的读写和协议解析，执行命令仍然是单线程顺序执行。
所以我们不需要去考虑控制 key、lua、事务，LPUSH`/`LPOP等等的并发及线程安全问题。
~~~



#### 三、其他问题总结

##### `redis` 内存碎片

###### 查看

输入命令 `info memory`

![](https://github.com/No8LaVine/MyCode/blob/master/images/redis1.png)

~~~shell
1、used_memory：
已经使用了的内存大小，包括redis进程内部开销和你的cache的数据所占用的内存，单位byte。

2、used_memory_human：
用户数据所占用的内存，就是你缓存的数据的大小。

3、used_memory_rss：（rss for Resident Set Size）
表示redis物理内存的大小，即向OS申请了多少内存使用与used_memory的区别在后面解释。

4、used_memory_peak：
redis内存使用的峰值。

5、used_memory_peak_human：
用户cache数据的峰值大小。

6、used_memory_lua：
执行lua脚本所占用的内存。

7、mem_fragmentation_ratio：
内存碎片率
~~~

其中`mem_fragmentation_ratio`就是内存碎片率，计算方法是`mem_fragmentation_ratio = used_memory_rss / used_memory`，大于 1 表示有内存碎片，小于 1 表示正在使用虚拟内存也就是硬盘。

~~~
内存碎片率在1~1.5之间是正常的

大于1.5要考虑清理内存碎片

小于1说明已经开始使用交换内存，也就是使用硬盘了，正常的内存不够用了，需要考虑是否要进行内存的扩容
~~~

###### 如何产生的

~~~
Redis内部有自己的内存管理器，为了提高内存使用的效率，来对内存的申请和释放进行管理。

Redis中的值删除的时候，并没有把内存直接释放，交还给操作系统，而是交给了Redis内部有内存管理器。

Redis中申请内存的时候，也是先看自己的内存管理器中是否有足够的内存可用。

Redis的这种机制，提高了内存的使用率，但是会使Redis中有部分自己没在用，却不释放的内存，导致了内存碎片的发生。
~~~

###### 解决方法

**低于4.0版本的Redis**

~~~
如果你的Redis版本是4.0以下的，Redis服务器重启后，Redis会将没用的内存归还给操作系统，碎片率会降下来。
~~~

**高于4.0版本的Redis**

`Redis4.0`版本开始，可以在不重启的情况下，线上整理内存碎片。自动碎片清理，只要设置了如下的配置，内存就会自动清理了。

```bash
config set activedefrag yes
```

如果想把`Redis`的配置，写到配置文件中去。

```bash
config rewrite
```

如果你对自动清理的效果不满意，可以使用如下命令，直接试下手动碎片清理：

```bash
memory purge
```



##### redis过期策略

> 1. 定时删除：在这是键的过期时间的同时，创建一个定时器 Timer，让定时器在键过期时间来临时立即执行对过期键的删除。
>
>    对内存友好，对 CPU 不友好。如果过期删除的键比较多的时候，删除键这一行为会占用相当一部分 CPU 性能，会对 Redis 的吞吐量造成一定影响。
>
>    
>
> 2. 惰性删除：键过期后不管，每次读取该键时，判断该键是否过期，如果过期删除该键返回空。
>
>    对 CPU 友好，内存不友好。如果很多键过期了，但在将来很长一段时间内没有很多客户端访问该键导致过期键不会被删除，占用大量内存空间。
>
>    
>
> 3. 定期删除：每隔一段时间对数据库中的过期键进行一次检查。
>
>    是定时删除和惰性删除的一种折中。每隔一段时间执行一次删除过期键的操作，并且限制删除操作执行的时长和频率。



##### redis如何保证原子性

单线程



##### redis事务

**multi开启，exec执行，discard放弃    watch监听**

![](https://github.com/No8LaVine/MyCode/blob/master/images/redis2.png)

![](https://github.com/No8LaVine/MyCode/blob/master/images/redis3.png)

###### 事务的特性

>隔离性：事务中的所有命令都会序列化、按顺序地执行。事务在执行的过程中，不会被其他客户端发送来的命令请求所打断。
>
>原子性：事务中的命令要么全部被执行，要么全部都不执行

###### 事务中的错误

~~~
如过在事务中操作一个key的时候这个key过期了，就会造成事务出错
~~~

**全体连坐**

![](https://github.com/No8LaVine/MyCode/blob/master/images/redis4.png)

**冤头债主**

![](https://github.com/No8LaVine/MyCode/blob/master/images/redis5.png)



###### watch监听

~~~
监听一个或多个key，执行之前这个key被其他命令改变，事务将被打断。
~~~



##### 怎么保证 redis 和 db 中的数据一致？

经典        缓存+数据库读写模式

~~~
读的时候，先读缓存，缓存没有再读数据库，然后将数据放入缓存，返回响应。

更新的时候，先更新数据库，再删除缓存
~~~



###### 为什么是删除缓存而不是更新缓存？

~~~
复杂场景下，缓存可能不只是数据库中取出的值。

更新缓存代价较高。
~~~



###### 更新数据库和删除缓存的顺序问题

* 方案一，先删除缓存，再更新数据库
  * 如果第一步删除缓存成功，第二步更新数据库失败，此时再次查询缓存，最多会有一次`cache miss`
* 方案二，先更新数据库，再删除缓存
  * 如果第一步更新数据库成功，第二步删除缓存失败,则造成下面的缓存不一致问题



###### 缓存不一致问题

更新完数据库在删除缓存的时候，删除失败，导致数据库中是新数据，缓存中是旧数据。

可能出现的情况：

* 删除缓存失败，那么不会去执行更新操作。
* 删除缓存成功，更新失败，读请求还是会将旧值写回到`redis`中。
* 删除缓存成功，更新成功，读请求会将新值写回到`redis`中。



###### 解决缓存不一致方案

**采用延时双删策略**

> 1. 双延时策略
>
>    ​	在写库前后都删一次缓存，设置合理的超时时间
>
> 2. 设置缓存过期时间

缺点：双删策略+缓存超时设置，这样最差的情况就是在超时时间内数据存在不一致，而且又增加了写请求的耗时。



**异步更新缓存(基于订阅binlog的同步机制)**

MySQL binlog 增量订阅消费+消息队列+增量数据更新到redis

> 读`Redis`：热数据基本都在`Redis`
>
> 写`MySQL`:增删改都是操作`MySQL`
>
> 更新`Redis`数据：`MySQ`的数据操作`binlog`，来更新到`Redis`

操作步骤：

> 1. 读取`binlog`后分析 ，利用消息队列,推送更新各台的`redis`缓存数据。
> 2. 这样一旦`MySQL`中产生了新的写入、更新、删除等操作，就可以把`binlog`相关的消息推送至`Redis`，`Redis`再根据`binlog`中的记录，对`Redis`进行更新
> 3. 其实这种机制，很类似`MySQL`的主从备份机制，因为`MySQL`的主备也是通过`binlog`来实现的数据一致性。
> 4. 这里可以结合使用`canal`(阿里的一款开源框架)，通过该框架可以对`MySQL`的`binlog`进行订阅，而`canal`正是模仿了`mysql`的`slave`数据库的备份请求，使得`Redis`的数据更新达到了相同的效果。
> 5. 当然，这里的消息推送工具你也可以采用别的第三方：`kafka`、`rabbitMQ`等来实现推送更新`Redis`!

##### redis 跳表

https://juejin.cn/post/6844903955831619597

https://segmentfault.com/a/1190000022028505

##### redis 内存与性能优化

https://www.biugogo.com/2019/12/01/Redis%E7%B3%BB%E5%88%97(4)---%E5%86%85%E5%AD%98%E4%BC%98%E5%8C%96/

##### `SDS`

https://github.com/No8LaVine/MyCode/blob/master/mydoc/redis/redis%20SDS.md

##### redis 高并发、高可用
https://github.com/No8LaVine/MyCode/blob/master/mydoc/redis/redis%20%E9%AB%98%E5%B9%B6%E5%8F%91%E3%80%81%E9%AB%98%E5%8F%AF%E7%94%A8.md

##### redis rehash

https://juejin.cn/post/6844903680706101262

https://www.cnblogs.com/williamjie/p/11205593.html

##### setnx 加锁

~~~python
def lock_after_pay(self):
    if self.client.setnx(self.getKeyName(), 1):
        self.client.expire(self.getKeyName(), 600)
        return True
    return False
~~~

##### redis内存回收策略

* https://github.com/No8LaVine/MyCode/blob/master/mydoc/%E5%B7%A5%E5%85%B7/redis%E6%BB%A1%E8%BD%BD%E5%8E%8B%E6%B5%8B.md
* https://www.jianshu.com/p/1f8e36285539
* https://juejin.cn/post/6844904193052934151

##### redis 分布式锁

https://juejin.cn/post/6844903688088059912

https://crossoverjie.top/2018/03/29/distributed-lock/distributed-lock-redis/

https://xiaomi-info.github.io/2019/12/17/redis-distributed-lock/

https://www.cnblogs.com/chengxy-nds/p/12750502.html