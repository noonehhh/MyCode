### nginx 防盗链

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
valid_referers none blocked *.hugao8.com www.hugao8.com m.hugao8.com *.baidu.com *.google.com;
if ($invalid_referer) {
	rewrite ^/ http://ww4.sinaimg.cn/bmiddle/051bbed1gw1egjc4xl7srj20cm08aaa6.jpg;
	#return 404;
  }
}
~~~

