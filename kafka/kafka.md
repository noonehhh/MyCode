### Kafka

##### 安装

* 下载安装包  `wget https://mirrors.bfsu.edu.cn/apache/kafka/2.4.1/kafka_2.11-2.4.1.tgz`
* 解压 `tar -zxvf kafka_2.11-2.4.1.tgz` `
* `kafka` 的启动需要依赖 `jdk` 和 `zookeeper` ，`zookeeper` 的安装见 [link]([MyCode/zookeeper.md at master · No8LaVine/MyCode (github.com)](https://github.com/No8LaVine/MyCode/blob/master/mydoc/zookeeper/zookeeper.md))

##### 启动

* 进入 `bin` 目录，执行 `zookeeper-server-start.sh -daemon config/zookeeper.properties`

##### 验证

进入bin目录执行以下

* 生产者生产消息 `./kafka-console-producer.sh --broker-list localhost:9092 --topic sun`
* 输入 `hello world`
* 消费者 `./kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic sun --from-beginning` ![](https://github.com/No8LaVine/MyCode/blob/master/images/kafka2.png))
* 此时消费者可以看到消息 ![](https://github.com/No8LaVine/MyCode/blob/master/images/kafka1.png)

