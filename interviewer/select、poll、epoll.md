### select、poll、epoll

#### select

~~~
select监视3类文件描述符 writefds readfds execptfds ，进程调用select后就会阻塞，直到有描述符阻塞或者超时返回，
返回以后可以通过遍历 fdset 来找到就绪的描述符
~~~

##### 数据结构

~~~c
int select (int n, fd_set *readfds, fd_set *writefds, fd_set *exceptfds, struct timeval *timeout);
~~~

##### 缺点

>1. 需要将描述符集合从用户态拷贝到内核态
>2. 监视的文件描述符有限制，linux 最多1024
>3.  需要遍历，文件描述符多了以后性能线性下降

#### poll

~~~
原理和select一样，不同的是 poll 用 通过一个pollfd数组向内核传递需要关注的事件，所以监视的文件描述符没有上限，
pollfd 里面是 fd 集合，需要监控的事件和已经发生的事件
~~~

##### 数据结构

~~~c
int poll (struct pollfd *fds, unsigned int nfds, int timeout);
~~~

##### poll 和 select 的异同（优缺点）

###### 相同

>1. 需要将描述符集合从用户态拷贝到内核态
>2.  需要遍历，文件描述符多了以后性能线性下降

###### 不同

> 1. 监视的文件描述符没有限制

#### epoll

~~~
用一个文件描述符去管理多个文件描述符，将用户关系到的文件描述符添加到一个事件表中，这样只需要复制一次，
然后监视这些事件的状态码变化，一旦有文件描述符状态就绪，就通知进程
~~~

##### 与select、poll的区别

>select/poll，进程只有在调用一定的方法后，内核才对所有监视的文件描述符进行扫描
>
>epoll,采用基于事件的就绪通知方式。epoll 同样只告知那些就绪的文件描述符，而且当我们调用 epoll_wait() 获得就绪文件描述符时，返回的不是实际的描述符，而是一个代表就绪描述符数量的值，你只需要去epoll指定的一个数组中依次取得相应数量的文件描述符即可

##### 步骤

>1. epoll事先通过epoll_ctl()来注册一个文件描述符，一旦基于某个文件描述符就绪时，将其加入到就绪链表
>2. 内核会采用类似callback的回调机制，迅速激活这个文件描述符，当进程调用 epoll_wait() 时便得到通知。

~~~
当某一进程调用epoll_create方法时，Linux内核会创建一个eventpoll结构体，这个结构体中有两个成员与epoll的使用方式密切相关


每一个epoll对象都有一个独立的eventpoll结构体，用于存放通过epoll_ctl方法向epoll对象中添加进来的事件。
这些事件都会挂载在红黑树中，
如此，重复添加的事件就可以通过红黑树而高效的识别出来(红黑树的插入时间效率是lgn，其中n为树的高度)。

而所有添加到epoll中的事件都会与设备(网卡)驱动程序建立回调关系，也就是说，当相应的事件发生时会调用这个回调方法。
这个回调方法在内核中叫ep_poll_callback，它会将发生的事件添加到rdlist双链表中。

当调用epoll_wait检查是否有事件发生时，只需要检查eventpoll对象中的rdlist双链表中是否有epitem元素即可
。如果rdlist不为空，则把发生的事件复制到用户态，同时将事件数量返回给用户。
~~~

##### 接口

~~~c
int epoll_create(int size)；
    
int epoll_ctl(int epfd, int op, int fd, struct epoll_event *event)；
    
int epoll_wait(int epfd, struct epoll_event * events, int maxevents, int timeout);
~~~

###### event_poll

~~~
每一个epoll对象都有一个独立的eventpoll结构体，用于存放通过epoll_ctl方法向epoll对象中添加进来的事件。
这些事件都会挂载在红黑树中，如此，重复添加的事件就可以通过红黑树而高效的识别出来(红黑树的插入时间效率是lgn，其中n为树的高度)。
~~~

###### 接口解析

>1. int epoll_create(int size)；
>
>创建一个epoll的句柄，size用来告诉内核这个监听的数目一共有多大。



>2. int epoll_ctl(int epfd, int op, int fd, struct epoll_event *event)；
>
>将epoll_event结构拷贝到内核空间中；
>
>参数：
>
>epfd：由 epoll_create 生成的epoll专用的文件描述符；
>op：要进行的操作例如注册事件，可能的取值EPOLL_CTL_ADD 注册、EPOLL_CTL_MOD 修 改、EPOLL_CTL_DEL 删除
>fd：关联的文件描述符；
>event：指向epoll_event的指针；



>3.int epoll_wait(int epfd, struct epoll_event * events, int maxevents, int timeout);
>
>该函数用于轮询I/O事件的发生，返回需要处理的事件数目，如返回0表示已超时。
>
>参数：
>
>events：用来从内核得到事件的集合
>
>maxevents：告之内核这个events有多大，这个 maxevents的值不能大于创建epoll_create()时的size，
>
>timeout：是超时时间（毫秒，0会立即返回，-1将不确定，也有说法说是永久阻塞）。

###### event_wait工作机制

>1. 计算睡眠时间(如果有)，判断eventpoll对象的链表是否为空,不为空那就干活，不睡眠，并且初始化一个等待队列，把自己挂上去，设置自己的进程状态为可睡眠状态。判断是否有信号到来(有的话直接被中断醒来)，如果啥事都没有那就调用schedule_timeout进行睡眠，如果超时或者被唤醒，首先从自己初始化的等待队列删除 ，然后开始拷贝资源给用户空间了。
>
>2. 拷贝资源则是先把就绪事件链表转移到中间链表，然后挨个遍历拷贝到用户空间。并且挨个判断其是否为水平触发，是的话再次插入到就绪链表。

##### 机制

~~~
epoll采用基于事件的就绪通知方式。

在select/poll中，进程只有在调用一定的方法后，内核才对所有监视的文件描述符进行扫描，
而epoll事先通过epoll_ctl()来注册一个文件描述符，一旦基于某个文件描述符就绪时，
内核会采用类似callback的回调机制，迅速激活这个文件描述符，当进程调用epoll_wait()时便得到通知。
~~~

##### 工作模式

###### LT模式

~~~
默认模式，只要这个fd还有数据可读，每次 epoll_wait都会返回它的事件，提醒用户程序去操作
~~~

###### ET模式

~~~
它只会提示一次，直到下次再有数据流入之前都不会再提示了，无论fd中是否还有数据可读。
所以在ET模式下，read一个fd的时候一定要把它的buffer读完，或者遇到EAGAIN错误
~~~



##### epoll如何解决select和poll的缺点？

###### 缺点一：拷贝问题

> 对于第一个缺点，epoll的解决方案在epoll_ctl函数中。每次注册新的事件到epoll句柄中时（在epoll_ctl中指定 EPOLL_CTL_ADD），会把所有的fd拷贝进内核，而不是在epoll_wait的时候重复拷贝。epoll保证了每个fd在整个过程中只会拷贝一次。

###### 缺点二：遍历问题

> 对于第二个缺点，epoll的解决方案不像select或poll一样每次都把current轮流加入fd对应的设备等待队列中，而只在 epoll_ctl时把current挂一遍（这一遍必不可少）并为每个fd指定一个回调函数，当设备就绪，唤醒等待队列上的等待者时，就会调用这个回调 函数，而这个回调函数会把就绪的fd加入一个就绪链表）。epoll_wait的工作实际上就是在这个就绪链表中查看有没有就绪的fd（利用 schedule_timeout()实现睡一会，判断一会的效果，和select实现中的第7步是类似的）。

###### 缺点三：限制问题

> 对于第三个缺点，epoll没有这个限制，它所支持的FD上限是最大可以打开文件的数目，这个数字一般远大于2048,举个例子, 在1GB内存的机器上大约是10万左右，具体数目可以cat /proc/sys/fs/file-max察看,一般来说这个数目和系统内存关系很大。

#### 总结

>（1）select，poll实现需要自己不断轮询所有fd集合，直到设备就绪，期间可能要睡眠和唤醒多次交替。而epoll其实也需要调用 epoll_wait不断轮询就绪链表，期间也可能多次睡眠和唤醒交替，但是它是设备就绪时，调用回调函数，把就绪fd放入就绪链表中，并唤醒在 epoll_wait中进入睡眠的进程。虽然都要睡眠和交替，但是select和poll在“醒着”的时候要遍历整个fd集合，而epoll在“醒着”的 时候只要判断一下就绪链表是否为空就行了，这节省了大量的CPU时间，这就是回调机制带来的性能提升。
>
>（2）select，poll每次调用都要把fd集合从用户态往内核态拷贝一次，并且要把current往设备等待队列中挂一次，而epoll只要 一次拷贝，而且把current往等待队列上挂也只挂一次（在epoll_wait的开始，注意这里的等待队列并不是设备等待队列，只是一个epoll内 部定义的等待队列），这也能节省不少的开销
>
>

真牛：

https://www.bilibili.com/video/BV1F54y1U7jN

https://zhuanlan.zhihu.com/p/347451068

https://gitlib.com/page/linux-io-event.html

https://www.ktanx.com/blog/p/4706

参考：

https://segmentfault.com/a/1190000003063859

http://www.mianshigee.com/question/10256gqt

https://www.modb.pro/db/28290