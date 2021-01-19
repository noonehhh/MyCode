### **是什么？**

* 基于nginx和lua的高性能web平台，内部集成了大量的lua库、第三方模块以及大多数的依赖项。

### **能干什么？**

* 用于搭建能够处理超高并发、扩展极高的web应用、web服务和动态网关。

### **安装**

**step1**：安装依赖库

~~~shell
yum install readline-devel pcre-devel openssl-devel gcc
~~~

**step2**：下载安装openresty

~~~shell
step2：下载安装openresty
wget https://openresty.org/download/openresty-1.9.15.1.tar.gz
tar xvf openresty-1.9.15.1.tar.gz
cd openresty-1.9.15.1
./configure --with-luajit && make && make install
~~~

**step3：**安装完成

~~~shell
输入 whereis openresty
>openresty: /usr/local/openresty
~~~

### **启动**

  -p指定项目目录 -c指定配置文件

~~~shell
/usr/local/openresty/nginx/sbin/nginx -c /usr/local/openresty/nginx/conf/nginx.conf
/usr/local/openresty/nginx/sbin/nginx -p 'pwd' -c /usr/local/openresty/nginx/conf/nginx.conf
上面的一般不用 一般直接用以下就就可
--启动
/opt/openresty/nginx/sbin/nginx
--停止
/opt/openresty/nginx/sbin/nginx -s stop
--重启
/opt/openresty/nginx/sbin/nginx -s reload
--检验nginx配置是否正确
/opt/openresty/nginx/sbin/nginx -t
~~~

### **配置**

- 三方库模块

~~~shell
lua_package_path lua模块路径，多个之间用;分隔，其中;;表示默认搜索路径，默认到nginx安装路径下找
例子：
lua_package_path "/home/work/nginx/conf/lua/?.lua;;"";  #lua 模块  
lua_package_cpath "/usr/servers/lualib/?.so;;";  #c模块
~~~

- 安装第三方库

- - 方式一，安装luarocks

  - 方式二

  - - 从给github下载三方包，例如resty.[http,https://github.com/pintsized/lua-resty-http](http,https://github.com/pintsized/lua-resty-http)
    - 下载后解压,将 lua-resty-http/lib/resty/ 目录下的 http.lua 和 http_headers.lua 两个文件拷贝到 /usr/local/openresty/lualib/resty 目录下即可

- 在nginx配置文件中写lua代码

~~~shell
location /lua {  
        content_by_lua 'ngx.say("hello world")';  
}  
~~~

- 代码过长改用文件处理

~~~shell
location /lua {   
    content_by_lua_file conf/lua/test.lua; #相对于nginx安装目录  
}   
~~~

- lua_code_cache lua代码缓存，即每次lua代码变更必须reload nginx才生效，如果在开发阶段可以通过lua_code_cache off;关闭缓存，这样调试时每次修改lua代码不需要reload nginx；但是正式环境一定记得开启缓存。 

~~~shell
   location /lua {   
        lua_code_cache off;  
        content_by_lua_file conf/lua/test.lua;  
}  
~~~

### openresty中lua脚本常用指令

~~~shell
ngx.req.get_headers   获取请求头
~~~

~~~shell
ngx.req.get_uri_args  获取url请求参数
~~~

~~~shell
ngx.req.get_post_args 获取post请求体 
也可用
ngx.req.read_body()
ngx.req.get_body_data()
这两个前提在配置文件中开启lua_need_request_body on;
~~~

~~~shell
ngx.var.request_uri 获取uri
~~~

### nginx常见配置

##### upstream

~~~shell
keepalive
    upstream xxx {
        server 0.0.0.0:8080 max_fails=3 fail_timeout=30s weight=5;
        keepalive 32;
    }

设置到 upstream 服务器的空闲 keepalive 连接的最大数量,当这个数量被突破时，最近使用最少的连接将被关闭
特别提醒：keepalive 指令不会限制一个 nginx worker 进程到 upstream 服务器连接的总数量

upstream 相关参数
    server 服务ip:端口
    weight 权重
    max_fails  最多失败连接的次数，超过就认为主机挂掉了
    fail_timeout  重新连接的时间
    backup  备用服务
    max_conns  允许的最大连接数
    slow_start  节点恢复后，等待多少秒后再加入
~~~

##### nginx 负载均衡算法

###### ip_hash算法

~~~shell
对于访问的ip，他会做一次hash运算，并对当前的负载应用数量做一次取余运算，这种算法能保证同一个ip访问的是同一台应用服务器。
upstream backend {
    ip_hash;
    server 127.0.0.1:8081;
    server 127.0.0.1:8082;
}
~~~

###### **url_hash算法**

~~~shell
对于请求的url进行hash运算，这种算法能保证同一个url访问的是同一台应用服务器。
upstream backend {
    url_hash;
    server 127.0.0.1:8081;
    server 127.0.0.1:8082;
}
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

##### nginx 防盗链

###### 什么是防盗链？

​	客户端向服务器请求资源时，为了减少网络带宽，提升响应时间，服务器一般不会一次将所有  资源完整地传回给客户端。比如在请求一个网页时，首先会传回该网页的文本内容，当客户端  浏览器在解析文本的过程中发现有图片存在时，会再次向服务器发起对该图片资源的请求，服务器将存储的图片资源再发送给客户端。在这个过程中，如果该服务器上只包含了网页的文本内容，并没有存储相关的图片资源，而是将图片资源链接到其他站点的服务器上，就形成了盗链行为。

​	比如在我项目中，我引用了的是淘宝中的一张图片的话，那么当我们网站重新加载的时候，就会请求淘宝的服务器，那么这就很有可能造成淘宝服务器负担。因此这个就是盗链行为。因此我们要实现防盗链。

###### 实现防盗链

​	使用http协议中请求头部的 `Referer` 头域来判断当前访问的网页或文件的源地址。通过该头域的值，我们可以检测访问目标资源的源地址。如果目标源地址不是我们自己站内的 `URL` 的话，那么这种情况下，我们采取阻止措施，实现防盗链。但是注意的是：`Referer` 头域中的值是可以被更改的。因此该方法也不能完全安全阻止防盗链。

###### Nginx 服务器的 Rewrite 功能实现防盗链。

`Nginx` 中有一个指令 `valid_referers`。 该指令可以用来获取 `Referer` 头域中的值，并且根据该值的情况给 `Nginx` 全局变量 `invalidreferer` 赋值。如果`Referer` 头域中没有符合 `validreferers` 指令的值的话，，`invalid_referer` 变量将会赋值为1。 valid_referers 指令基本语法如下：

~~~shell
valid_referers  none | blocked | server_names | string
none: 		    检测Referer头域不存在的情况。
blocked：        检测Referer头域的值被防火墙或者代理服务器删除或伪装的情况。那么在这种情况下，该头域的值不以"http://" 或 "https://" 开头。
server_names:  	设置一个或多个URL，检测Referer头域的值是否是URL中的某个。
~~~

###### 完整代码

~~~shell
location ~* \.(gif|jpg|png|jpeg)$ {

expires     30d;
valid_referers *.hugao8.com www.hugao8.com m.hugao8.com *.baidu.com *.google.com;
if ($invalid_referer) {
	rewrite ^/ http://ww4.sinaimg.cn/bmiddle/051bbed1gw1egjc4xl7srj20cm08aaa6.jpg;
	#return 404;
  }

}
~~~



### 其他常用配置及指令见

https://juejin.cn/post/6844904144235413512#heading-14

https://www.jianshu.com/p/8e0877d69b39

https://www.cnblogs.com/jimodetiantang/p/9257819.html

https://www.cnblogs.com/tinywan/p/6526191.html

