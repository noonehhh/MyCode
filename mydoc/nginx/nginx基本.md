### nginx常见配置

~~~shell
main        # 全局配置，对全局生效
├── events  # 配置影响 Nginx 服务器或与用户的网络连接
├── http    # 配置代理，缓存，日志定义等绝大多数功能和第三方模块的配置
│   ├── upstream # 配置后端服务器具体地址，负载均衡配置不可或缺的部分
│   ├── server   # 配置虚拟主机的相关参数，一个 http 块中可以有多个 server 块
│   ├── server
│   │   ├── location  # server 块可以包含多个 location 块，location 指令用于匹配 uri
│   │   ├── location
│   │   └── ...
│   └── ...
└── ...
~~~

##### main

~~~shell
#user administrator administrators;  指定nginx进程使用什么用户启动

#worker_processes 2; 指定启动多少进程来处理请求，一般情况下设置成CPU的核数，如果开启了ssl和gzip更应该设置成与逻辑CPU数量一样甚至为2倍，可以减少I/O操作

#pid /nginx/pid/nginx.pid;   #指定nginx进程运行文件存放地址

#制定日志路径，级别。这个设置可以放入全局块，http块，server块，级别以此为：debug|info|notice|warn|error|crit|alert|emerg，，debug输出日志最为最详细，而crit输出日志最少。
error_log log/error.log debug;
~~~

##### events

~~~shell
events {
	#设置网路连接序列化，防止惊群现象发生，默认为on
	accept_mutex on;
	
	#设置一个进程是否同时接受多个网络连接，默认为off
	multi_accept on;
	
	#最大连接数，默认为512,每一个worker进程能并发处理（发起）的最大连接数（包含与客户端或后端被代理服务器间等所有连接数）
	#计算公式 最大连接数 = worker_processes * worker_connections/4
	worker_connections  1024; 
    
	#use是个事件模块指令，用来指定Nginx的工作模式，select、poll、kqueue、epoll、resig、/dev/poll、eventport，
	#其中select和poll都是标准的工作模式，kqueue和epoll是高效的工作模式，不同的是epoll用在Linux平台上，而kqueue用在BSD系统中
	#对于Linux系统，epoll工作模式是首选。在操作系统不支持这些高效模型时才使用select。
	use epoll;      
}
~~~

##### http

~~~shell
http {
	# 文件扩展名与类型映射表
	include       mime.types;
	
	 # 加载子配置项
	include /etc/nginx/conf.d/*.conf;
	
	#默认文件类型，默认为text/plain
	#default_type属于HTTP核心模块指令，这里设定默认类型为二进制流，也就是当文件类型未定义时使用这种方式，
	#例如在没有配置PHP环境时，Nginx是不予解析的，此时，用浏览器访问PHP文件就会出现下载窗口。
	default_type  application/octet-stream; 
	
	#access_log off; #取消服务日志    
	log_format myFormat '$remote_addr–$remote_user [$time_local] $request $status $body_bytes_sent $http_referer $http_user_agent $http_x_forwarded_for'; 
	
	# Nginx访问日志存放位置
	access_log log/access.log myFormat;
	
	#允许sendfile方式传输文件，默认为off，可以在http块，server块，location块。
	sendfile on;
	
	# 减少网络报文段的数量
	tcp_nopush          on;
	
	#每个进程每次调用传输数量不能大于设定的值，默认为0，即不设上限。
    sendfile_max_chunk 100k;
    
    #连接超时时间，默认为75s，可以在http，server，location块。
    keepalive_timeout 65;

    #错误页
	error_page 404 https://www.baidu.com; 
}
~~~

##### upstream

~~~shell
#负载均衡
upstream mysvr {   
    server 127.0.0.1:7878;
    server 192.168.10.121:3333 backup;  #热备
    server 0.0.0.0:8080 max_fails=3 fail_timeout=30s weight=5;
    
    #设置到 upstream 服务器的空闲 keepalive 连接的最大数量,当这个数量被突破时，最近使用最少的连接将被关闭
	#特别提醒：keepalive 指令不会限制一个 nginx worker 进程到 upstream 服务器连接的总数量
    keepalive 32;
}

upstream 相关参数
    server 服务ip:端口
    weight 权重
    max_fails  最多失败连接的次数，超过就认为主机挂掉了
    fail_timeout  重新连接的时间
    backup  备用服务
    max_conns  允许的最大连接数
    slow_start  节点恢复后，等待多少秒后再加入
~~~

##### server

~~~shell
    server {
        keepalive_requests 120; #单连接请求上限次数。
        listen       4545;   #监听端口
        server_name  127.0.0.1;   #监听地址       
        location  ~*^.+$ {       #请求的url过滤，正则匹配，~为区分大小写，~*为不区分大小写。
            #root path;  #根目录
            #index vv.txt;  #设置默认页
            proxy_pass  http://mysvr;  #请求转向mysvr 定义的服务器列表
            deny 127.0.0.1;  #拒绝的ip
            allow 172.18.5.54; #允许的ip           
        } 
    }
    
    #server 块可以包含多个 location 块，location 指令用于匹配 uri，语法：
    location [ = | ~ | ~* | ^~] uri {
		...
	}
	
	= 精确匹配路径，用于不含正则表达式的 uri 前，如果匹配成功，不再进行后续的查找；
	^~ 用于不含正则表达式的 uri； 前，表示如果该符号后面的字符是最佳匹配，采用该规则，不再进行后续的查找；
	~ 表示用该符号后面的正则去匹配路径，区分大小写；
	~* 表示用该符号后面的正则去匹配路径，不区分大小写。跟 ~ 优先级都比较低，如有多个location的正则能匹配的话，则使用正则表达式最长的那个；
	
	如果 uri 包含正则表达式，则必须要有 ~ 或 ~* 标志。
~~~

##### location

~~~shell
location / {
    root   /usr/share/nginx/html;  # 网站根目录
    index  index.html index.htm;   # 默认首页文件
    deny 172.168.22.11;   # 禁止访问的ip地址，可以为all
    allow 172.168.33.44； # 允许访问的ip地址，可以为all
}
~~~

##### 注意点

~~~shell
上面是nginx的基本配置，需要注意的有以下几点：

1.$remote_addr 与$http_x_forwarded_for 用以记录客户端的ip地址； 
2.$remote_user ：用来记录客户端用户名称； 
3.$time_local ： 用来记录访问时间与时区；
4.$request ： 用来记录请求的url与http协议；
5.$status ： 用来记录请求状态；成功是200， 
6.$body_bytes_s ent ：记录发送给客户端文件主体内容大小；
7.$http_referer ：用来记录从那个页面链接访问过来的； 
8.$http_user_agent ：记录客户端浏览器的相关信息；

惊群现象：一个网路连接到来，多个睡眠的进程被同事叫醒，但只有一个进程能获得链接，这样会影响系统性能。

每个指令必须有分号结束。
~~~

##### nginx 命令行常用命令

~~~shell
nginx # 启动 nginx

nginx -s reload # 向主进程发送信号，重新加载配置文件，热重启

nginx -s reopen # 重启 Nginx

nginx -s stop # 快速关闭

nginx -s quit # 等待工作进程处理完成后关闭

nginx -t # 查看当前 Nginx 配置是否有错误

nginx -t -c <配置路径> # 检查配置是否有问题，如果已经在配置目录，则不需要 - c
~~~

##### root和alias

~~~shell
nginx指定文件路径有两种方式root和alias，指令的使用方法和作用域。
root与alias主要区别在于nginx如何解释location后面的uri，这会使两者分别以不同的方式将请求映射到服务器文件上。
root的处理结果是：root路径＋location路径
alias的处理结果是：使用alias路径替换location路径
alias是一个目录别名的定义，root则是最上层目录的定义。
还有一个重要的区别是alias后面必须要用“/”结束，否则会找不到文件的。。。而root则可有可无
---------------------root-----------------------
location   ^ ~   / t /   {
      root   / www / root / html / ;
}
如果一个请求的URI是/t/a.html时，web服务器将会返回服务器上的/www/root/html/t/a.html的文件。

---------------------alias-----------------------
location   ^ ~   / t /   {
alias   / www / root / html / new_t / ;
}
如果一个请求的URI是/t/a.html时，web服务器将会返回服务器上的/www/root/html/new_t/a.html的文件。注意这里是new_t，因为alias会把location后面配置的路径丢弃掉，把当前匹配到的目录指向到指定的目录。


-------------------------------------------------
注意：
1. 使用alias时，目录名后面一定要加"/"。
3. alias在使用正则匹配时，必须捕捉要匹配的内容并在指定的内容处使用。
4. alias只能位于location块中。（root可以不放在location中）
~~~

##### rewrite

~~~shell
rewrite 地址重定向
* 使用位置：server location if
* 基本语法：rewrite regex replacement [flag];
	rewrite的含义：该指令是实现URL重写的指令。
	regex的含义：用于匹配URI的正则表达式。
	replacement：将regex正则匹配到的内容替换成 replacement。
	flag: flag标记。
		flag 有以下值
		last: 本条规则匹配完成后，继续向下匹配新的location URI 规则。(不常用)
		break: 本条规则匹配完成即终止，不再匹配后面的任何规则(不常用)。
        redirect: 返回302临时重定向，浏览器地址会显示跳转新的URL地址。
		permanent: 返回301永久重定向。浏览器地址会显示跳转新的URL地址。
~~~

##### nginx 缓存

###### http 块

~~~shell
proxy_cache_path /www/javayz/cache levels=1:2 keys_zone=cache_javayz:500m inactive=20d max_size=1g;

#proxy_cache_path  缓存存放的路径 
#levels  缓存层级及目录的位数，1:2表示两级目录，第一级目录用1位16进制表示，第二级目录用2位16进制表示
#keys_zone 缓存区内存大小
#inactive 有效期，如果缓存有效期内未使用，则删除
#max_size 存储缓存的硬盘大小
~~~

###### location 块

~~~shell
#指定缓存区，就是上面设置的key_zone
proxy_cache cache_javayz;

#缓存的key，这里用请求的全路径md5做为key
proxy_cache_key $host$uri$is_args$args;

#对不通的http状态码设置不同的缓存时间,下面的配置表示200时才进行缓存，缓存时间12小时
proxy_cache_valid 200 12h;
~~~

##### expires 

~~~shell
expires 缓存
expires 30s;#30秒
expires 30m;#30分钟
expires 2h;#2个小时
expires 30d;#30天
expires -1;#不缓存
~~~



#### 日常使用技巧

##### 日志切割脚本

~~~shell
#!/bin/bash
#设置你的日志存放的目录
log_files_path="/mnt/usr/logs/"
#日志以年/月的目录形式存放
log_files_dir=${log_files_path}"backup/"
#设置需要进行日志分割的日志文件名称，多个以空格隔开
log_files_name=(access.log error.log)
#设置nginx的安装路径
nginx_sbin="/mnt/usr/sbin/nginx -c /mnt/usr/conf/nginx.conf"
#Set how long you want to save
save_days=10

############################################
#Please do not modify the following script #
############################################
mkdir -p $log_files_dir

log_files_num=${#log_files_name[@]}
#cut nginx log files
for((i=0;i<$log_files_num;i++));do
    mv ${log_files_path}${log_files_name[i]} ${log_files_dir}${log_files_name[i]}_$(date -d "yesterday" +"%Y%m%d")
done
$nginx_sbin -s reload
~~~

##### 开启 gzip 压缩

解释

`gzip` 是一种常用的网页压缩技术，传输的网页经过 `gzip` 压缩之后大小通常可以变为原来的一半甚至更小（官网原话），更小的网页体积也就意味着带宽的节约和传输速度的提升，特别是对于访问量巨大大型网站来说，每一个静态资源体积的减小，都会带来可观的流量与带宽的节省。

百度可以找到很多检测站点来查看目标网页有没有开启 `gzip` 压缩， [<网页GZIP压缩检测>](http://tool.chinaz.com/Gzips/Default.aspx?q=juejin.im) 输入掘金 `juejin.im` 来偷窥下掘金有没有开启 gzip。

###### `nginx` 配置 `gzip`

使用 `gzip` 不仅需要 ~配置，浏览器端也需要配合，需要在请求消息头中包含 `Accept-Encoding: gzip`（IE5 之后所有的浏览器都支持了，是现代浏览器的默认设置）。一般在请求 `html` 和 `css` 等静态资源的时候，支持的浏览器在 `request` 请求静态资源的时候，会加上 `Accept-Encoding: gzip` 这个 `header`，表示自己支持 `gzip` 的压缩方式，`Nginx` 在拿到这个请求的时候，如果有相应配置，就会返回经过 `gzip` 压缩过的文件给浏览器，并在 `response` 相应的时候加上 `content-encoding: gzip` 来告诉浏览器自己采用的压缩方式（因为浏览器在传给服务器的时候一般还告诉服务器自己支持好几种压缩方式），浏览器拿到压缩的文件后，根据自己的解压方式进行解析。

~~~shell
# /etc/nginx/conf.d/gzip.conf

gzip on; # 默认off，是否开启gzip
gzip_types text/plain text/css application/json application/x-javascript text/xml application/xml application/xml+rss text/javascript;

# 上面两个开启基本就能跑起了，下面的愿意折腾就了解一下
gzip_static on;
gzip_proxied any;
gzip_vary on;
gzip_comp_level 6;
gzip_buffers 16 8k;
# gzip_min_length 1k;
gzip_http_version 1.1;

#解释
gzip_types：要采用 gzip 压缩的 MIME 文件类型，其中 text/html 被系统强制启用；
gzip_static：默认 off，该模块启用后，Nginx 首先检查是否存在请求静态文件的 gz 结尾的文件，如果有则直接返回该 .gz 文件内容；
gzip_proxied：默认 off，nginx做为反向代理时启用，用于设置启用或禁用从代理服务器上收到相应内容 gzip 压缩；
gzip_vary：用于在响应消息头中添加 Vary：Accept-Encoding，使代理服务器根据请求头中的 Accept-Encoding 识别是否启用 gzip 压缩；
gzip_comp_level：gzip 压缩比，压缩级别是 1-9，1 压缩级别最低，9 最高，级别越高压缩率越大，压缩时间越长，建议 4-6；
gzip_buffers：获取多少内存用于缓存压缩结果，16 8k 表示以 8k*16 为单位获得；
gzip_min_length：允许压缩的页面最小字节数，页面字节数从header头中的 Content-Length 中进行获取。默认值是 0，不管页面多大都压缩。建议设置成大于 1k 的字节数，小于 1k 可能会越压越大；
gzip_http_version：默认 1.1，启用 gzip 所需的 HTTP 最低版本；

注意，一般 gzip 的配置建议加上 gzip_min_length 1k，不加的话：
由于文件太小，gzip 压缩之后可能会得到负的的体积优化，压缩之后体积还比压缩之前体积大了，所以最好设置低于 1kb 的文件就不要 gzip 压缩了
~~~





### 其他常用配置及指令见

https://juejin.cn/post/6844904144235413512#heading-26

https://learnku.com/articles/46237