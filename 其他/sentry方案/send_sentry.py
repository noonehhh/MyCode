import tornado.web
import tornado.ioloop
import raven

client = raven.Client("http://wow:vod_be@sentry.quene.com/20")

class Sentry(tornado.web.RequestHandler):
    def get(self):
        client.captureMessage('sentry test')
        self.write("hello world")
if __name__ == "__main__":
    app = tornado.web.Application([
        (r"/", Sentry)
    ])
    app.listen(8000)
    tornado.ioloop.IOLoop.current().start()