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

##### ls -ahl

~~~shell
a 显示隐藏的文件
l 显示详细列表模式
h 大小显示便于查看（例如G，K，M）
~~~

##### 查看连接你服务器 top10 用户端的 IP 地址

~~~shell
netstat -nat | awk '{print $5}' | awk -F ':' '{print $1}' | sort | uniq -c | sort -rn | head -n 10
~~~

##### 查看你最常用的10个命令

~~~shell
cat .bash_history | sort | uniq -c | sort -rn | head -n 10 (or cat .zhistory | sort | uniq -c | sort -rn | head -n 10
~~~

##### z系列

~~~
zless      查看压缩文件
zcat       用于不真正解压缩文件，就能显示压缩包中文件的内容的场合
		   -S：指定gzip格式的压缩包的后缀。当后缀不是标准压缩包后缀时使用此选项；
		   -c：将文件内容写到标注输出；
		   -d：执行解压缩操作；
		   -l：显示压缩包中文件的列表；
		   -L：显示软件许可信息；
		   -q：禁用警告信息；
		   -r：在目录上执行递归操作；
		   -t：测试压缩文件的完整性；
		   -V：显示指令的版本信息；
		   -l：更快的压缩速度；
		   -9：更高的压缩比。
zmore      使用zmore命令可以查看gzip、zip、compress压缩文件。
zgrep      命令可以在压缩文件中调用grep按正则表达式来搜索
~~~

##### linux文件时间截点：atime、ctime与mtime

~~~shell
1.atime是指access time，即文件被读取或者执行的时间，修改文件是不会改变access time的。

Time whenfile data was last accessed. Changedby  the following  functions:  creat(),  mknod(),  pipe(),utime(2), and read(2).

2.mtime即modify time，指文件内容被修改的时间，是在写入文件时随文件内容的更改而更改的。

Time whendata was last modified. Changed bythe  fol- lowing  functions:  creat(),mknod(), pipe(), utime(), andwrite(2).

3.ctime即change time文件状态改变时间，是在写入文件、更改所有者、权限或链接设置时随 Inode 的内容更改而更改的。

Time whenfile status was last changed. Changed by the following  functions:  chmod(),  chown(),  creat(), link(2),  mknod(),  pipe(),  unlink(2),  utime(),  and write().
~~~

##### linux查看文件ctime、atime、mtime命令

~~~shell
ls -lc test :查看test文件的ctime

ls -lu test :查看test文件的atime

ls -l test:查看test文件的mtime
~~~

##### –mtime中的参数n

~~~shell
–mtime n中的n指的是24*n(即n天), +n、-n、n分别表示：

+n：大于n，操作发在n+1天以前

-n：小于n，操作发生在n天以内

 n：等于n，操作刚好在n天时
~~~

##### 找到目录下所有的txt文件并删除

`find ./ -name "*.txt" -exec rm -rf {} \;`

##### **IO模式：select，poll,epoll解析**

