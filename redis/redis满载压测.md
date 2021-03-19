### redis满载压测

##### 原因

~~~
6G 内存的reids在一个月内扩容到24G，且存在重复数据，鉴于此情况，需要对redis进行一次压测，
看能否更改为其他策略，如先进先出策略或淘汰最少使用策略。
~~~

##### 前期调研：

~~~
调研发现，redis的内存回收策略可以更改，默认是 noeviction；
~~~

##### redis内存回收策略：

- **volatile-lru**：采用最近使用最少的淘汰策略，`Redis` 将回收那些超时的（仅仅是超时的）键值对，也就是它只淘汰那些超时的键值对。

- **allkeys-lru**：采用最近最少使用的淘汰策略，`Redis` 将对所有（不仅仅是超时的）的键值对采用最近最少使用的淘汰策略。

- **volatile-lfu**：采用最近最不常用的淘汰策略，所谓最近最不常用，也就是一定时期内被访问次数最少的。`Redis` 将回收超时的键值对。

- **allkeys-lfu**：采用最近最不常用的淘汰策略，`Redis` 将对所有的键值对采用最近最不常用的淘汰策略。

- **volatile-random**：采用随机淘汰策略删除超时的键值对。

- **allkeys-random**：采用随机淘汰策略删除所有的键值对，这个策略不常用。

- **volatile-ttl**：采用删除存活时间最短的键值对策略。

- **noeviction:**不淘汰任何键值对，当内存满时，如果进行读操作，例如 `get` 命令，它将正常工作，而做写操作，它将返回错误，也就是说，当 `Redis` 采用这个策略内存达到最大的时候，它就只能读不能写了。

  参考：https://www.jianshu.com/p/677930ffbff0

##### 压测方案：

- 将 `redis` 的最大内存调整为 1mb，通过 `shell` 脚本写入数据；

```
查看redis内存相关配置 
info memory 
设置redis最大内存 
config set maxmemory 1mb 
查看是否设置成功 
config get maxmemory
```

1. 测试是否是 `LRU`（最近最少使用）淘汰策略：分时间段对不同前缀的数据进行调用，再进行写入数据操作，然后查看最早被调用的数据是否存在；
2. 测试是否是先进先出淘汰策略：对满载的 `redis` 直接写入数据，看最早写入的数据是否存在；
3. 测试是否是 `LFU`（最近不经常使用）淘汰策略：对数据进行访问，不同前缀的数据访问频率不同，再写入数据，查看被调用频率最少的数据是否存在；

##### 压测过程：

1、回收策略：`noeviction` (不淘汰任何键值对)
`shell` 脚本：

```shell
#/bin/bash 
for((i=1;i<10000;i++))do 
	redis-cli set $i $i 
	echo "get $i"|redis-cli 
done;
```

结果：从4097之后开始报 `OOM`

![](https://github.com/No8LaVine/MyCode/blob/master/images/redis1.png)



2、回收策略：`allkeys-lru`（最近最少使用的淘汰）

`shell` 脚本：

```shell
#/bin/bash 
for((i=1;i<10000;i++))do 
	redis-cli set $i $i 
	echo "get $i"|redis-cli 
done
```

结果：没有出现报错情况，脚本执行完之后在 `redis` 里执行 `get` 操作，发现 `redis` 里只有 7014-9999 这段数据，说明 `redis` 执行了最近最少使用的淘汰策略。

![](https://github.com/No8LaVine/MyCode/blob/master/images/all-lru.png)



3、回收策略：`allkeys-lfu`（最近最不常用的淘汰）

`shell` 脚本：

```shell
#/bin/bash 
for((i=1;i<10000;i++))do 
	redis-cli set $i $i 
	num=`expr $i % 2` 
	if [ $num -eq 0 ] 
	then        
		echo "get $i"|redis-cli 
	fi 
done;
```

结果：未出现报错，当超出内存时，优先淘汰掉使用频率低的数据，写入数据超过最大内存时，先淘汰掉没被使用过的单数，然后继续写入数据，再次超过最大内存时，会将最先写入的双数淘汰。

![](https://github.com/No8LaVine/MyCode/blob/master/images/all-lfu.png)



4、回收策略：`volatile-ttl` 当内存不足以容纳新写入数据时，在设置了过期时间的键空间中，有更早过期时间的 `key` 优先移除

`shell` 脚本：

```shell
#/bin/bash 
for((i=1;i<5000;i++))do 
	redis-cli set $i $i 
	if [ $i -gt 1000 ] 
	then        
		redis-cli EXPIRE $i 3600 
	elif [ $i -gt 2000 ] 
	then        
		redis-cli EXPIRE $i 3000 
	elif [ $i -gt 3000 ] 
	then        
		redis-cli EXPIRE $i 2400 
	elif [ $i -gt 4000 ] 
	then        
		redis-cli EXPIRE $i 1800 
	else        
		redis-cli EXPIRE $i 1200 
	fi 
done;
```

结果：1-1000，1000-2000，2000-3000段内的数据没有，猜测 `volatile-ttl` 策略是优先淘汰过期时间较长的 `key`

![](https://github.com/No8LaVine/MyCode/blob/master/images/ttl1.png)

验证以上猜想：

`shell` 脚本：

```shell
#/bin/bash 
for((i=1;i<5000;i++))do 
	redis-cli set $i $i 
	if [ $i -gt 1000 ] 
	then        
		redis-cli EXPIRE $i 1200 
	elif [ $i -gt 2000 ] 
	then        
		redis-cli EXPIRE $i 1800 
	elif [ $i -gt 3000 ] 
	then        
		redis-cli EXPIRE $i 2400 
	elif [ $i -gt 4000 ] 
	then        r
		edis-cli EXPIRE $i 3000 
	else        
		redis-cli EXPIRE $i 3600 
	fi 
done;
```

结果：3000-5000 段内的数据不存在，1-3000 数据存在，验证 `volatile-ttl` 策略是优先淘汰过期时间较长的数据；

![](https://github.com/No8LaVine/MyCode/blob/master/images/ttl2.png)

最终采用 `volatile-lru` 策略