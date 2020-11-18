### 是什么？
sentry意为哨兵，用来监控生产环境，一旦出现的错误或异常，第一时间通过信息或邮件形式反馈给管理员。  
***
### 安装
两种方式：docker、python
1. docker安装
```
1. git clone https://github.com/getsentry/onpremise.git
    此项目包含使用Docker搭建sentry所需要的基本配置文件，常修改的有三个：
        sentry.conf.py 和 config.yml : 基础配置 
        requirements.txt : 配置安装插件
2. docker volume create --name=sentry-data 
   docker volume create --name=sentry-postgres
   创建两个数据卷，分别提供给sentry和数据库使用
3. cp -n .env.example .env
   将onpremise项目提供的env.example模板文件复制一份，并命名为.env，
  目的是让docker-compose从此文件中获取环境变量，此文件目前只是配置了sentry 秘钥 
  参考：docker-compose env使用
4. docker-compose build
   根据docker-compose.yml中的配置构建容器
5. docker-compose run --rm web config generate-secret-key
   生成秘钥，生成之后需要将秘钥复制到STEP 3中的.env中
6. docker-compose run --rm web upgrade
   构建数据库，用户根据向导可配置超级管理员用户
7. 按需要修改配置文件
8. docker-compose up -d 启动所有docker服务
9. 访问localhost:9000进入管理界面

排查错误
查看日志 docker-compose logs -f -t
保存输出 docker-compose logs -f -t >> myDockerCompose.log**
```
2. python安装  
https://learnku.com/articles/4295/centos6-install-python-based-on-sentry#reply22542
- - -

### sentry 架构
sentry服务端分为web、cron、wocker这几个部分，应用出现错误后将错误信息上报给web，web处理后放入消息队列或redis内存队列，wocker从队列中消费数据进行处理。  

* my-sentry：sentry的web服务  
* sentry-cron：sentry的定时任务，活性检测  
* sentry-worker：业务处理，数据持久化，报警  
  ![](.assets/clipboard1.png)

* 关于DSN

  在sentry添加一个项目后，将获得一个DSN，看起来像一个标准URL，但实际是由sentry sdk所需配置的标识符，由几个部分组成，包括协议、公钥、密钥、服务地址和项目标识符。

  ~~~json
  {PROTOCOL}://{PUBLIC_KEY}:{SECRET_KEY}@{HOST}/{PATH}{PROJECT}
  ~~~

  由5部分组成：

  1.  使用的协议: http或https;

  2.  验证sdk的公钥和密钥;

  3.  目标sentry服务器;

  4.  验证用户绑定的项目

  

  