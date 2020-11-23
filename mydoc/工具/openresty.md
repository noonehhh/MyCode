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

~~~shell
keepalive
upstream xxx {
    server 0.0.0.0:8080 max_fails=3 fail_timeout=30s weight=5;
    keepalive 32;
}

设置到upstream服务器的空闲keepalive连接的最大数量,当这个数量被突破时，最近使用最少的连接将被关闭

特别提醒：keepalive指令不会限制一个nginx worker进程到upstream服务器连接的总数量
~~~

~~~shell
expires 缓存
expires 30s;#30秒
expires 30m;#30分钟
expires 2h;#2个小时
expires 30d;#30天
expires -1;#不缓存
~~~

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



### 其他常用配置及指令见

https://www.jianshu.com/p/8e0877d69b39

https://www.cnblogs.com/jimodetiantang/p/9257819.html

https://www.cnblogs.com/tinywan/p/6526191.html

