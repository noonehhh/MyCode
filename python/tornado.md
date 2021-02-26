Tornado

**是什么？**

​    是一个python web框架和异步网络库，通过使用非阻塞网络I/O，Tornado可以扩展到数万个开放连接，使其非常适合 long polling ， WebSockets 以及其他需要与每个用户建立长期连接的应用程序。

**四大模块**

1. web框架：包括RequestHandler，他是创建web应用程序和各种支持类的子类；
2. HTTP的客户端和服务器端的实现，HTTPServer和AsyncHTTPClient;
3. 包含异步网络库IOLoop和IOStream作为HTTP组件的构建块，也可以用于实现其他模块。
4. 协程库tornada.gen，允许异步代码以比链接回调更简单的方式写入；

**异步和非阻塞IO**

- 实时web功能要求每个用户都有一个长期的、主要是空闲的连接，每个用户一个线程，开销非常的大；
- tornado采用单线程事件循环，所有的程序都以异步和非阻塞为目标，因为一次只能有一个操作处于活动状态；

**认证和安全**

- 普通的cookie由于可修改并不安全，如需设置cookie，例如鉴权、认证等，则需要对cookie进行签名防止伪造。tornado支持通过set_secure_cookie和get_secure_cookie方法对cookie签名。使用这两个方法需要在创建应用的时候指定一个名为cookie_secret的密钥。

~~~py
application = tornado.web.Application([    (r"/", MainHandler), ], cookie_secret="__TODO:_GENERATE_YOUR_OWN_RANDOM_VALUE_HERE__")
~~~

**HelloWorld**

~~~py
class MainHandler(tornado.web.RequestHandler):
    def get(self):
        self.write("行到水穷处，坐看云起时")


if __name__ == '__main__':
    app = tornado.web.Application(r'/', MainHandler)
    app.listen(8080)
    tornado.ioloop.IOLoop.current().start()
~~~



**web框架**

**主要模块：**

1. tornado.web.application和RrquestHandler类处理HTTP请求
2. tornado.template模板渲染
3. tornado.routing处理路由

**tornado.template模板渲染**

~~~py
class MainHandler(tornado.web.RequestHandler):
    def get(self):
        self.write("行到水穷处，坐看云起时")
        self.render("<h1>hello worrld<h1>")
        self.render("hello.html")
        self.render("hello.html"，msg="欢迎登录系统")

if __name__ == '__main__':
    app = tornado.web.Application(
    [r'/', MainHandler],
    template_path=os.path.join(os.path.dirname(__file__),"templates")
    )
    app.listen(8080)
    tornado.ioloop.IOLoop.current().start()

方法
self.render("<h1>hello worrld<h1>")
文件渲染
self.render("hello.html")
传参的方式对文件渲染
self.render("hello.html"，msg="欢迎登录系统")
此方法需要hello.html中的某个变量为msg，例如:
    <h1>{{msg}}<h1>
~~~



**HTTP服务端和客户端**

~~~py
def main(self):
    http_client = httpclient.AsyncHTTPClient()
    url = url_concat(url, params)
    request = httpclient.HTTPRequest(url=url, method='GET', headers=headers)
tornado.ioop.IOLoop.current().run_sync(main)
~~~



**协程和并发**

tornado开启协程是用注解@gen.coroutine