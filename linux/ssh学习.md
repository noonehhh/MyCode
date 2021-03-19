### ssh学习

##### 常用

* ssh 跳转的时候一般是，本地---跳板机---宿主机 ，`ssh xxx@host`

* `-t` 参数，可以在向目标机器发送命令，例如，

  `ssh -t work@host "ls"`

  ![](https://github.com/No8LaVine/MyCode/blob/master/images/ssh1.png)

  `-t` 参数可以实现从本地---宿主机，`ssh -t work@jump_host "ssh work@target_host"`

