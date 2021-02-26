### **supervisor**

##### 是什么：

* Linux/Unix系统进程监控工具

##### 能干什么：

- 启动、重启、关闭进程，除了对单个进程的控制，还可以同时启动、关闭多个进程，同时启动所有进程；

##### 怎么干：

1. supervisor管理进程，就是通过fork/exec的方式把这些被管理的进程，当作supervisor的子进程来启动，我们只需在supervisor的配置文件里，把要管理的进程的可执行文件的路径写进去即可。
2. 被管理的进程作为supervisor的子进程，当子进程挂掉的时候，父进程就可以准确地获取子进程挂掉的信息，所以自然可以对挂掉的子进程进行自动重启，若不想重启，可将配置文件里的autostart = false。supervisor通过INI格式配置文件进行配置，为每个进程提供了很多多配置选项，可以很容易的重启进程或者自动的轮转日志。

##### 组件

- supervisord 主进程，负责管理进程的server，它会根据配置文件创建指定数量的应用程序的子程序，管理子程序的整个生命周期，对crash的进程重启，对进程变化发送事件通知等。同时内置web server和XML-RPC Interface,轻松实现进程管理，该配置文件位于/etc/supervisor/supervisor.conf
- supervisorctl客户端命令行工具，提供一个类似shell的操作接口，通过它可以连接到不同的supervisord进程上来管理他们各自的子程序，命令通过UNIX socket或者TCP来和服务通讯。命令通过命令行发送消息给supervisord，可以查看进程状态，加载配置文件，启动停止进程，查看进程标准输出和错误输出，远程操作等，服务端也可以要求客户端提供身份验证之后才能进行操作。
- web server supervisor提供了web server功能，可通过web控制进程（需要设置[inet httpserver]配置项）
- XML-RPC Interface XML-RPC接口，就像HTTP提供web UI一样，用来控制supervisor和由它运行的程序

##### **安装**

~~~shell
1.用python下载
pip install supervisor
2.生成配置文件
echo_supervisord_conf > /etc/supervisord.conf
~~~

* ps  为了不将所有新增配置信息全部放在同一配置文件里，每个程序设置一个文件夹，相互隔离；

##### **配置示例**

~~~shell
vim supervisord.conf

//加入以下配置信息
[include]
files = /etc/supervisord.d/*.conf

//在supervisord.conf中设置通过web可以查看管理的进程，加入以下代码
[inet_http_server] 
port=9001
username=user      
password=123

//启动supervisord
supervisord -c /etc/supervisrd.conf

启动后查看9001端口是否监听
通过ip:9001可以查看supervisor的web界面
~~~

##### 管理

命令行工具（supervisorctl）或者web管理界面

supervisorctl支持的命令：

- supervisorctl help查看

~~~shell
update 更新新的配置到supervisord（不会重启原来已运行的程序）
reload，载入所有配置文件，并按新的配置启动、管理所有进程（会重启原来已运行的程序）
start xxx: 启动某个进程
restart xxx: 重启某个进程
stop xxx: 停止某一个进程(xxx)，xxx为[program:theprogramname]里配置的值
stop groupworker: 重启所有属于名为groupworker这个分组的进程(start,restart同理)
stop all，停止全部进程，注：start、restart、stop都不会载入最新的配置文
reread，当一个服务由自动启动修改为手动启动时执行一下就ok
~~~

##### **supervisor配置参数介绍**

unix_http_server配置块

在该配置块的参数项表示的是一个监听在socket上的HTTP server

~~~powershell
- file：一个unix domain socket的文件路径，HTTP/XML-RPC会监听在这上面 
- chmod：在启动时修改unix domain socket的mode 
- chown：修改socket文件的属主
- username：HTTP server在认证时的用户名 
- password：认证密码
~~~

inet_http_server配置块

在该配置块的参数表示的是一个监听在TCP上的HTTP server

~~~shell
- port：TCP监听的地址和端口(ip:port)，这个地址会被HTTP/XML-RPC监听
- username：HTTP server在认证时的用户名
- password：认证密码
~~~

supervisord配置块

该配置项的参数是关于supervisod进程的全局配置块

~~~shell
- logfile：log文件路径
- logfile_maxbytes：log文件达到多少后自动进行轮转，单位是KB、MB、GB。如果设置为0则表示不限制日志文件大小
- logfile_backups：轮转日志备份的数量，默认是10，如果设置为0，则不备份 
- loglevel：error、warn、info、debug、trace、blather、critical 
- pidfile：pid文件路径 
- umask：umask值，默认022 
- nodaemon：如果设置为true，则supervisord在前台启动，而不是以守护进程启动 
- minfds：supervisord在成功启动前可用的最小文件描述符数量，默认1024 
- minprocs：supervisord在成功启动前可用的最小进程描述符数量，默认200 
- nocleanup：防止supervisord在启动的时候清除已经存在的子进程日志文件 
- childlogdir：自动启动的子进程的日志目录 
- user：supervisord的运行用户 
- directory：supervisord以守护进程运行的时候切换到这个目录 
- strip_ansi：消除子进程日志文件中的转义序列 
- environment：一个k/v对的list列表
~~~

program配置块

该配置块就是我们要监控的程序的配置项，该配置块的头部是有固定格式的，一个关键字program，后面跟着一个冒号，接下来才是程序名。例如：[program:foo]，foo就是程序名，在使用supervisorctl来操作程序的时候，就是以foo来标明的。该块的参数介绍如下

~~~shell
- command：启动程序使用的命令，可以是绝对路径或者相对路径 
- process_name：一个python字符串表达式，用来表示supervisor进程启动的这个的名称，默认值是%(program_name)s 
- numprocs：Supervisor启动这个程序的多个实例，如果numprocs>1，则process_name的表达式必须包含%(process_num)s，默认是1 
- numprocs_start：一个int偏移值，当启动实例的时候用来计算numprocs的值 
- priority：权重，可以控制程序启动和关闭时的顺序，权重越低：越早启动，越晚关闭。默认值是999 
- autostart：如果设置为true，当supervisord启动的时候，进程会自动重启。 
- autorestart：值可以是false、true、unexpected。false：进程不会自动重启，unexpected：当程序退出时的退出码不是exitcodes中定义的时，进程会重启，true：进程会无条件重启当退出的时候。 
- startsecs：程序启动后等待多长时间后才认为程序启动成功 
- startretries：supervisord尝试启动一个程序时尝试的次数。默认是3 
- exitcodes：一个预期的退出返回码，默认是0,2。 
- stopsignal：当收到stop请求的时候，发送信号给程序，默认是TERM信号，也可以是 HUP, INT, QUIT, KILL, USR1, or USR2。 
- stopwaitsecs：在操作系统给supervisord发送SIGCHILD信号时等待的时间 
- stopasgroup：如果设置为true，则会使supervisor发送停止信号到整个进程组 
- killasgroup：如果设置为true，则在给程序发送SIGKILL信号的时候，会发送到整个进程组，它的子进程也会受到影响。 
- user：如果supervisord以root运行，则会使用这个设置用户启动子程序 
- redirect_stderr：如果设置为true，进程则会把标准错误输出到supervisord后台的标准输出文件描述符。 
- stdout_logfile：把进程的标准输出写入文件中，如果stdout_logfile没有设置或者设置为AUTO，则supervisor会自动选择一个文件位置。 
- stdout_logfile_maxbytes：标准输出log文件达到多少后自动进行轮转，单位是KB、MB、GB。如果设置为0则表示不限制日志文件大小 
- stdout_logfile_backups：标准输出日志轮转备份的数量，默认是10，如果设置为0，则不备份 
- stdout_capture_maxbytes：当进程处于stderr capture mode模式的时候，写入FIFO队列的最大bytes值，单位可以是KB、MB、GB 
- stdout_events_enabled：如果设置为true，当进程在写它的stderr到文件描述符的时候，PROCESS_LOG_STDERR事件会被触发 
- stderr_logfile：把进程的错误日志输出一个文件中，除非redirect_stderr参数被设置为true 
- stderr_logfile_maxbytes：错误log文件达到多少后自动进行轮转，单位是KB、MB、GB。如果设置为0则表示不限制日志文件大小 
- stderr_logfile_backups：错误日志轮转备份的数量，默认是10，如果设置为0，则不备份 
- stderr_capture_maxbytes：当进程处于stderr capture mode模式的时候，写入FIFO队列的最大bytes值，单位可以是KB、MB、GB 
- stderr_events_enabled：如果设置为true，当进程在写它的stderr到文件描述符的时候，PROCESS_LOG_STDERR事件会被触发 
- environment：一个k/v对的list列表 
- directory：supervisord在生成子进程的时候会切换到该目录 
- umask：设置进程的umask 
- serverurl：是否允许子进程和内部的HTTP服务通讯，如果设置为AUTO，supervisor会自动的构造一个url
~~~

例如：

~~~shell
[program:test_http] 
command=python test_http.py 10000 ; 被监控的进程启动命令 
directory=/root/ ; 执行前要不要先cd到目录去，一般不用 
priority=1 ;数字越高，优先级越高 
numprocs=1 ; 启动几个进程 
autostart=true ; 随着supervisord的启动而启动 
autorestart=true ; 自动重启。。当然要选上了 
startretries=10 ; 启动失败时的最多重试次数 
exitcodes=0 ; 正常退出代码（是说退出代码是这个时就不再重启了吗？待确定） 
stopsignal=KILL ; 用来杀死进程的信号 
stopwaitsecs=10 ; 发送SIGKILL前的等待时间 
redirect_stderr=true ; 重定向stderr到stdout
~~~

参考：https://juejin.im/post/6844903945937092622#heading-2