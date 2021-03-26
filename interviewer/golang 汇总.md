### 汇总

#### array、slice

https://github.com/No8LaVine/MyCode/blob/master/interviewer/golang%20array%E5%92%8Cslice.md

#### map

https://github.com/No8LaVine/MyCode/blob/master/interviewer/golang%20map%E6%B5%85%E6%9E%90.md

#### struct能否被比较？

https://github.com/No8LaVine/MyCode/blob/master/interviewer/golang%20struct%E8%83%BD%E5%90%A6%E8%A2%AB%E6%AF%94%E8%BE%83%EF%BC%9F.md

#### make和new的区别

https://github.com/No8LaVine/MyCode/blob/master/interviewer/golang%20struct%E8%83%BD%E5%90%A6%E8%A2%AB%E6%AF%94%E8%BE%83%EF%BC%9F.md

#### 内存逃逸

https://mp.weixin.qq.com/s/aAwzo_LuiVP6dAKDl0SiOg

#### Golang交替打印问题

https://mp.weixin.qq.com/s/1GJ3dJSBkqYwFIUaDhcQLQ

#### 映客面经

https://mp.weixin.qq.com/s/WTBViFBJG4J3TR-yqyryjQ

#### GMP 模型，为什么要有 P？

https://mp.weixin.qq.com/s/an7dml9NLOhqOZjEGLdEEw

#### 应用层有哪些协议

域名系统DNS协议、FTP文件传输协议、telnet远程终端协议、HTTP超文本传送协议、SMTP电子邮件协议、POP3邮件读取协议、Telnet远程登录协议

#### GC

何时触发GC：

>1. 手动调用 runtime.GC
>2. 在堆上分配大于 32K byte 对象的时候进行检测此时是否满足垃圾回收条件，如果满足则进行垃圾回收。

GC出发的条件：

> 1. `forceTrigger || memstats.heap_live >= memstats.gc_trigger` 。forceTrigger 是 forceGC 的标志；后面半句的意思是当前堆上的活跃对象大于我们初始化时候设置的 GC 触发阈值

https://www.kancloud.cn/aceld/golang/1958308

#### go的栈空间管理

https://github.com/No8LaVine/MyCode/blob/master/interviewer/golang%20%E6%A0%88%E7%A9%BA%E9%97%B4%E7%AE%A1%E7%90%86.md

#### GMP

https://www.kancloud.cn/aceld/golang/1958305

#### 进程、线程、协程

https://github.com/No8LaVine/MyCode/blob/master/interviewer/%E8%BF%9B%E7%A8%8B%E3%80%81%E7%BA%BF%E7%A8%8B%E3%80%81%E5%8D%8F%E7%A8%8B.md

#### string、byte、rune的区别

https://cloudsjhan.github.io/2018/10/25/golang%E4%B8%ADstring-rune-byte%E7%9A%84%E5%85%B3%E7%B3%BB/