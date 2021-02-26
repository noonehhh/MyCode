### nginx 配置跨域

由于浏览器同源策略的存在使得一个源中加载来自其它源中资源的行为受到了限制。即会出现跨域请求禁止。

##### 同源策略

​	域名、协议、端口相同。

~~~shell
server {
        listen       80;
        server_name  www.phpblog.com.cn;
        root   /Users/shiwenyuan/blog/public;
        index  index.html index.htm index.php;
        location / {
            try_files $uri $uri/ /index.php?$query_string;
        }
        add_header 'Access-Control-Allow-Origin' "$http_origin";
        add_header 'Access-Control-Allow-Credentials' 'true';
        add_header 'Access-Control-Allow-Methods' 'GET, POST, OPTIONS, DELETE, PUT, PATCH';
        add_header 'Access-Control-Allow-Headers' 'DNT,X-CustomHeader,Keep-Alive,User-Agent,X-Requested-With,If-Modified-Since,Cache-Control,Content-Type,X-XSRF-TOKEN';

        location ~ \.php$ {
            fastcgi_pass   127.0.0.1:9000;
            fastcgi_index  index.php;
            fastcgi_param  SCRIPT_FILENAME $document_root$fastcgi_script_name;
            include        fastcgi_params;
        }
        error_page  404              /404.html;
        error_page   500 502 503 504  /50x.html;
        location = /50x.html {
            root   html;
        }
}
~~~

* `Access-Control-Allow-Origin`：允许的域名，只能填 *（通配符）或者单域名。
* `Access-Control-Allow-Methods`: 允许的方法，多个方法以逗号分隔。
* `Access-Control-Allow-Headers`: 允许的头部，多个方法以逗号分隔。
* `Access-Control-Allow-Credentials`: 是否允许发送 `Cookie`。