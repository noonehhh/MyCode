### redis 持久化

~~~
redis 有两种持久化方式，RDB 和 AOF，这两种方式可以同时使用。
~~~

#### RDB （快照）

##### 特点

>1. `RDB` 是 `redis` 默认的持久化方式，持久化可以在指定的时间间隔内生成数据集的时间点快照，也就是说它保存了某个时间点的`Redis`数据，它所生成的 `RDB` 文件是一个压缩的二进制文件。
>2. 性能更好，需要进行持久化的时候，主进程会 `fork` 一个子进程出来，把持久化的工作交给子进程，自己不会有相关的`I/O`操作。
>3. 相比于 `AOF`，在数据量比较大的时候，`RDB` 的启动速度更快。

##### 原理

>1. `redis` 调用 `fork()` ，产生一个子进程
>2. 子进程把数据写进一个临时文件
>3. 子进程写完新的 `RDB` 文件后，把旧的 `RDB` 文件替换掉

##### 缺点 

>1. 容易造成丢失数据，比如 `redis` 每 5 分钟备份一次，如果 `redis` 宕机了那么就会丢失部分数据
>
>2. 使用 `fork()` 子进程进行数据持久化，如果数据量比较大花费的时间比较多，`fork()` 可能会非常耗时，造成服务器在某某毫秒内停止处理客户端。

##### 用法

###### save

~~~
save 执行一个同步保存操作，将当前 Redis 实例的所有数据快照(snapshot)以 RDB 文件的形式保存到硬盘。阻塞主进程，
客户端无法连接 redis，等 save 完成后，主进程才开始工作，客户端可以连接
~~~

###### bgsave

> 1. `Redis`父进程判断当前是否存在正在执行的子进程，如`RDB/AOF`子进程，如果存在`bgsave`命令直接返回。
> 2. `redis`父进程`fork`出一个子进程，不影响父进程处理客户端请求，但`fork`操作过程中父进程被阻塞。
> 3. 父进程`fork`完成后，`bgsave`命令返回`Background saving started by pid xxx`信息，并不再阻塞父进程，可以继续响应其他命令。
> 4. 由于 `fork` 的写时复制(`copy on write`)机制，父进程、子进程共享相同的内存地址，只有当父进程发生写操作修改内存数据时，才会去分配内存空间。也就是说，`redis`的子进程使用共享的内存地址完成`RDB`备份,而父进程当收到写请求后，会将涉及到的内存页复制出一份副本，在此副本进行修改。当子进程`RDB`完成后，通知主进程更换副本，`RDB`就此完成。
> 5. 子进程将数据全部写入临时 `rdb` 文件，用临时 `rdb` 文件替换原来的快照文件，子进程退出。

 3. 配置文件

    ​	`redis.conf`

    ~~~shell
    # 900s内至少达到一条写命令
    save 900 1
    # 300s内至少达至10条写命令
    save 300 10
    # 60s内至少达到10000条写命令
    save 60 10000
    ~~~

#### AOF

~~~
通过记录写操作完成持久化，将写操作记录到 .aof 为后缀的文件中。
~~~

##### 特点

> `AOF`只是追加日志文件，因此对服务器性能影响较小，速度比`RDB`要快，消耗的内存较少。

##### 缺点

>1. AOF方式生成的日志文件太大，即使通过AFO重写，文件体积仍然很大。
>2. 恢复数据的速度比RDB慢。

##### 用法

###### 配置文件开启

~~~shell
# 开启aof机制
appendonly yes

# aof文件名
appendfilename "appendonly.aof"

# 写入策略,always表示每个写操作都保存到aof文件中,也可以是everysec或no
appendfsync always

# 默认不重写aof文件
no-appendfsync-on-rewrite no

# 保存目录
dir ~/redis/
~~~

###### 配置文件的问题

~~~
AOF 的机制是不断将写操作追加到文件的末尾，随着写命令的增加，AOF 文件也越来越大，例如，
对一个数据进行了 100 次 incr 操作，则有 100 条记录
但其实这条数据只需要一个 set 命令即可。为了处理这种问题，AOF 文件重写应用而生。‘
~~~

###### AOF 重写

~~~
执行 BGREWRITEAOF 命令，优化当前 AOF 文件体积，即使 BGREWRITEAOF 执行失败，也不会有任何数据丢失，
因为旧的 AO` 文件在 BGREWRITEAOF 成功之前不会被修改。AOF重写方式也是异步操作
~~~

>1. 如果要写入`AOF`文件，则`Redis`主进程会`fork`一个子进程来处理
>2. 如果在重写过程中有新的变动，主进程会把新的变动写道内存的缓存区，同时把这些变动写到老的 `AOF` 文件里，确保即使重写失败也可保证数据安全
>3. 子进程完成重写后，给父进程发送一个信号，主进程会获得一个信号，把缓存区里的操作追加到新的 `AOF` 文件中。

**预定（scheduled）**

~~~
如果Redis的子进程正在执行快照的保存工作，那么AOF重写的操作会被预定(scheduled)，等到保存工作完成之后再执行AOF重写。
在这种情况下，BGREWRITEAOF的返回值仍然是OK，但还会加上一条额外的信息，说明BGREWRITEAOF要等到保存操作完成之后才能执行。
在 Redis 2.6 或以上的版本，可以使用INFO命令查看BGREWRITEAOF是否被预定。
~~~

**重复重写**

~~~
如果已经有别的AOF文件重写在执行，那么BGREWRITEAOF返回一个错误，并且这个新的BGREWRITEAOF请求也不会被预定到下次执行。
~~~

###### 配置设置重写条件

~~~shell
# Redis会记住自从上一次重写后AOF文件的大小（如果自Redis启动后还没重写过，则记住启动时使用的AOF文件的大小）。
# 如果当前的文件大小比起记住的那个大小超过指定的百分比，则会触发重写。
# 同时需要设置一个文件大小最小值，只有大于这个值文件才会重写，以防文件很小，但是已经达到百分比的情况。

auto-aof-rewrite-percentage 100
auto-aof-rewrite-min-size 64mb

# 要禁用自动的日志重写功能，我们可以把百分比设置为0
auto-aof-rewrite-percentage 0
~~~

###### 三种写入策略

`appendfsync`选项指定写入策略,有三个选项

~~~shell
客户端的每一个写操作都保存到aof文件当，这种策略很安全，但是每个写请注都有IO操作，所以也很慢。
appendfsync always

appendfsync的默认写入策略，每秒写入一次aof文件，因此，最多可能会丢失1s的数据。
appendfsync everysec

Redis服务器不负责写入aof，而是交由操作系统来处理什么时候写入aof文件。更快，但也是最不安全的选择，不推荐使用
appendfsync no
~~~

###### 数据损坏修复

在写入`AOF`日志文件时，如果`Redis`服务器宕机，则`AOF`日志文件文件会出格式错误，在重启`Redis`服务器时，`Redis`服务器会拒绝载入这个`AOF`文件，可以通过以下步骤修复`AOF`并恢复数据。

~~~SHEL
# 修复aof日志文件
$ redis-check-aof -fix file.aof
~~~

参考：

https://ningg.top/computer-basic-theory-copy-on-write/

https://draveness.me/whys-the-design-redis-bgsave-fork/

https://www.yuque.com/justdoit-oriyu/lltri2/bvs6go