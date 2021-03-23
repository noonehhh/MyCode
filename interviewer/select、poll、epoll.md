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

##### 接口

~~~c
int epoll_create(int size)；//创建一个epoll的句柄，size用来告诉内核这个监听的数目一共有多大
    
int epoll_ctl(int epfd, int op, int fd, struct epoll_event *event)；
    
int epoll_wait(int epfd, struct epoll_event * events, int maxevents, int timeout);
~~~

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

###### ET模式的实现机制

~~~
epoll 同样只告知那些就绪的文件描述符，而且当我们调用epoll_wait()获得就绪文件描述符时，返回的不是实际的描述符，
而是一个代表就绪描述符数量的值，你只需要去epoll指定的一个数组中依次取得相应数量的文件描述符即可

这里也使用了内存映射（mmap）技术，这样便彻底省掉了这些文件描述符在系统调用时复制的开销。

另一个本质的改进在于epoll采用基于事件的就绪通知方式。在select/poll中，进程只有在调用一定的方法后，
内核才对所有监视的文件描述符进行扫描

而epoll事先通过epoll_ctl()来注册一个文件描述符，一旦基于某个文件描述符就绪时，内核会采用类似callback的回调机制，
迅速激活这个文件描述符，当进程调用epoll_wait()时便得到通知。
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

https://gitlib.com/page/linux-io-event.html

https://www.ktanx.com/blog/p/4706

参考：

https://segmentfault.com/a/1190000003063859

http://www.mianshigee.com/question/10256gqt

https://www.modb.pro/db/28290