### 日常试用集结

##### **查询端口号相关信息**

~~~shell
根据端口号
lsof -i:7777
COMMAND   PID  USER   FD   TYPE DEVICE SIZE/OFF NODE NAME
testhttp  8491 root   3u   IPv6 152950 0t0      TCP  *:cbt (LISTEN)

netstat -nap|grep 7777
tcp6 0 0 :::7777 :::*  LISTEN      8491/./testhttp     
unix  3  [ ]  STREAM CONNECTED 17777 1/systemd  /run/systemd/journal/stdout
根据进程名
ps aux|grep testhttp
root 8491  0.0  0.1 219456  6272 pts/1    Sl+  14:13   0:00 ./testhttp
root 8665  0.0  0.0 112728   972 pts/0    S+   14:16   0:00 grep --color=a

1.  根据进程pid查端口：
lsof -i | grep pid

2.  根据端口port查进程（某次面试还考过）：
lsof  -i:port     

3. 根据进程pid查端口：
 netstat -nap | grep pid

4.  根据端口port查进程
netstat -nap | grep port
~~~

##### **查看cpu核数**

~~~shell
逻辑
cat /proc/cpuinfo | grep name | cut -f2 -d: | uniq -c
物理
cat /proc/cpuinfo | grep "physical id" | sort| uniq|wc -l
~~~

##### **查找文件**

~~~shell
在当前路径下查找demo开头的文件
find -name demo*
在根目录下查找demo开头的文件
find / -name demo*
在根目录下查找小写字母开头的文件
find / -name [a-z]*.txt
在当前目录下查找不是demo开头的txt文件
~~~

##### **查看进程运行目录**

~~~shell
ps -ef|grep 8080
>work  20791  0.0  0.7 432888 62936 pts/1    S+   14:19   0:00 python main.py --port=8080 --debug=False
ll /proc/20791/cwd
> /test

cwd符号链接的是进程运行目录；
exe符号连接就是执行程序的绝对路径；
cmdline就是程序运行时输入的命令行命令；
environ记录了进程运行时的环境变量；
fd目录下是进程打开或使用的文件的符号连接。
~~~

##### **显示历史命令的运行时间**

~~~shell
export HISTTIMEFORMAT='%F %T '
~~~

##### dig命令

~~~shell
用来从DNS域名服务器查询主机地址信息
~~~

[link](https://www.cnblogs.com/sparkdev/p/7777871.html)

##### **IO模式：select，poll,epoll解析**

