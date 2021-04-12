import datetime
import time
import requests
import redis
import json
import sys
import logging
import raven

config = {
    "wow": {
        "url": "http://10.10.153.224:9001",
        "vod_be": {
            "key": "a6f431e198734105aef464f6e702351f",
            "secret": "563b78f7bc1941828df0e5ebd79cc17a",
            "num": "20"
        },
        "mcms": {
            "key": "35836757370643db9cf51fe2df425bc9",
            "secret": "8ca3a8825fd24fada3a47420e55e40f0",
            "num": "10"
        },
        "vod_li": {
            "key": "aead908e4b4e4480a700f85021d4849e",
            "secret": "cb28f0b991d64c31989bb4822a0310bf",
            "num": "23"
        },
        "wow_user": {
            "key": "3a598aef625243a49800294315600cee",
            "secret": "cac677ca2c014c969869570f60c9711d",
            "num": "2"
        }
    },
    "ktv": {
        "url": "http://10.10.153.224:9002",
        "vadd_stb1": {
            "key": "7ca45e28b66f4e86a13cdd9ff28638ac",
            "secret": "c2604aa199744814a1b84b595cbb9067",
            "num": "6"
        },
        "lscms": {
            "key": "ede4aa042e6e4da4bb650421eb44bd11",
            "secret": "497ade054d96464bbc84aaba96f1dbef",
            "num": "11"
        },
        "value_added": {
            "key": "1d631feadae24c808d2a4ba87010d005",
            "secret": "2597e93c26a046a49065183079748552",
            "num": "2"
        }
    },
    "default": {
        "url": "http://10.10.153.224:9001",
        "default": {
            "key": "a6f431e198734105aef464f6e702351f",
            "secret": "563b78f7bc1941828df0e5ebd79cc17a",
            "num": "20"
        }
    }

}

REDIS_CACHE = {
    'host':'10.10.134.128',
    'port': '6379',
    "db":6
}

r = redis.StrictRedis(host=REDIS_CACHE['host'], port=REDIS_CACHE['port'], db=REDIS_CACHE['db'])

class Redis():
    def connect(self):
        args = sys.argv[1:]
        if len(args) > 0:
            global r
            host = args[0]
            port = args[1]
            db = args[2]
            r = redis.StrictRedis(host=host, port=port, db=db)
    def get_info(self):
        keys = r.smembers("sentry_keys")
        return keys

        #a = r.rpop("sentry:header:uri")
        #if a is None:
        #       return None,None,None
        #data = json.loads(a.decode())
        #body_key = "sentry:body:{}".format(data["timestamp"])
        #body = r.get(body_key)
        #return data, body, body_key


class Sentry(object):
    def run(self, sentry_key, data, body, body_key):
        headers, u, limit_key = self.create_sentry(data["header"], data["uri"])
        self.ratelimit_sleep(limit_key)
        res = self.sentry_post(u, headers, body)

        if res is not None:
            if res.status_code == 200:
                r.delete(body_key)
            logging.info(res.text)
            print(res.text)
        else:
            r.lpush(sentry_key, json.dumps(data))
    def ratelimit_sleep(self,limit_key):
        times = r.get(limit_key)
        print(limit_key)
        args = limit_key.split(':')
        if times is not None:
            if int(times) == 10:
                print('休眠中')
                content = "msg: sentry 熔断，project：{}".format('{}:{}'.format(args[1], args[2]))
                self.gropu_msg(content=content)
                time.sleep(20)
            if int(times) == 5:
                content = {
                    "msg_type": "text",
                    "content": {
                        "text": '''
                        项目名称: {project}
                        告警时间: {time}
                        消息内容:一分钟内Senrty告警次数已达 5 次
                        '''.format(project=args[2], time=datetime.datetime.now().strftime("%Y-%m-%d %H:%M:%S"))
                    }
                }
                send_msg(content)
    def gropu_msg(self, content, bot="https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=dc56a968-27f1-479e-9939-ae3e15978e55", t = 4):
        if not content:
            return
        try:
            msg = {"msgtype": "text","text": {"content": content}}
            msg = json.dumps(msg)
            msg = msg.encode('utf-8')
            res = requests.post(bot, headers={'Content-Type': 'application/json'}, data=msg)
            print(res.content)
            logging.info(res.content)
        except Exception as e:
            logging.error(e)
            print(e)
            time.sleep(10)
            self.group_msg(content, bot, t=t-1)

    def create_sentry(self, header, uri):
        k_index = -1
        s_index = -1
        sentryInfo = header['x-sentry-auth'].split(' ')
        for i in range(len(sentryInfo)):
            if "key" in sentryInfo[i]:
                k_index = i
                flag = sentryInfo[i][sentryInfo[i].index("=") + 1:].replace(",","",-1)
            elif "secret" in sentryInfo[i]:
                s_index = i
                project = sentryInfo[i][sentryInfo[i].index("=") + 1:].replace(",","",-1)
        if flag not in config.keys():
            flag = "default"
            project = "default"
        sentry_key = "sentry_key=" + config[flag][project]["key"] + ","
        sentry_secret = "sentry_secret=" + config[flag][project]["secret"]
        if k_index == -1:
            sentryInfo.append(sentry_key)
        else:
            sentryInfo[k_index] = sentry_key
        if s_index == -1:
            sentryInfo.append(sentry_secret)
        else:
            sentryInfo[s_index] = sentry_secret
        header['x-sentry-auth'] = " ".join(sentryInfo)

        url = config[flag]["url"]
        uri_arr = uri.split("/")
        for i in range(len(uri_arr)):
            if uri_arr[i].isdigit():
                uri_arr[i] = config[flag][project]["num"]
        uri = "/".join(uri_arr)
        u = url + uri
        limit_key = "sentry:{}:{}:limit".format(flag,project)
        return header, u, limit_key

    def sentry_post(self, u, headers, body):
        try:
            res = requests.post(url=u, headers=headers, data=body)
            return res
        except Exception as e:
            logging.error("post to sentry err", e)


    # 飞书消息推送
def send_msg(data, t=4, feishu_url='https://open.feishu.cn/open-apis/bot/v2/hook/d3a90b72-aec1-486c-bf55-3e802483022a'):
    if t <= 0:
        return
    try:
        requests.post(feishu_url, json=data)
    except Exception as e:
        logging.error(e)
        time.sleep(10)
        send_msg(data, t=t - 1)


if __name__ == '__main__':
    Redis().connect()
    while True:
        keys = Redis().get_info()
        if keys is None:
            time.sleep(10)
        else:
            for i in keys:
                a = r.rpop(i)
                if a is None:
                    continue
                data = json.loads(a.decode())
                #if data.get('timestamp') is None:
                #       r.rpush(i,a)
                #       continue
                body_key = "sentry:body:{}".format(data["timestamp"])
                body = r.get(body_key)
                Sentry().run(i,data, body, body_key)
                if r.llen(i) == 0:
                    r.srem("sentry_keys", i)