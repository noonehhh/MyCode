# DOCKER

#### 为什么会出现

* 解决配置部署问题
* 一次安装导出运行
* 容器虚拟化

#### 虚拟机的缺点

* 资源占用多
* 冗余步骤多
* 启动慢

#### 概念

* 镜像：类似Java中的类，本身是只读的
* 容器：镜像的运行实例，可以被开始、启动、暂停、删除、每个容器都是不想隔离的，把凭证安全的平台
* 仓库：集中存放镜像的场所，分为私有仓库和公开仓库

#### centos7安装docker

- 根据官方文档安装，一切的教学资料均来源于官方文档。
- 1.只能装所需软件包
- 2.设置stage镜像仓库，注意由于防火墙问题，需从阿里云镜像安装。

~~~shell
yum-config-manager --add-repo http://mirrors.aliyun.com/docker-ce/linux/centos/docker-ce.repo
~~~

- 3.更新yum软件包索引 yum makecache fast
- 4.安装docker CE yum -y install docker-ce
- 5.启动docker systemtcl start docker
- 6.配置镜像加速

- - 1.mkdir -p /etc/docker
  - 2.vim /etc/docker/daemon.json
  - 3.systemctl daemon-reload
  - 4.systemctl restart docker

- 7.卸载docker

- - 1.systemctl stop docker
  - 2.yum -y remove docker-ce

- - * rm -rf /var/lib/docker

#### 工作流程

![image](https://github.com/No8LaVine/MyCode/blob/master/images/docker2.png)

#### docker运行原理

- 1、docker是怎么工作的

docker是一个C-S结构的系统，docker守护进程运行在主机上，通过socket连接从客户端访问，守护进程从客户端接受命令并管理运行在主机上的容器，即运行时环境

- 2、为什么docker比vm快

- - 1.docker有更少的抽象层，由于docker不需要Hypervisor实现硬件资源虚拟化,运行在docker容器上的程序直接使用的都是实际物理机的硬件资源。因此在CPU、内存利用率上docker将会在效率上有明显优势。
  - 2.docker利用的是宿主机的内核,而不需要Guest OS。因此,当新建一个容器时,docker不需要和虚拟机一样重新加载一个操作系统内核。仍而避免引寻、加载操作系统内核返个比较费时费资源的过程,当新建一个虚拟机时,虚拟机软件需要加载Guest OS,返个新建过程是分钟级别的。而docker由于直接利用宿主机的操作系统,则省略了返个过程,因此新建一个docker容器只需要几秒钟。

- ##### **vm于docker的区别**

![img](.assets/untitle-1605667758231.png)



* 其他详见Docker.html    文件过大，下载下来才可以看
