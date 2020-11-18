### **shell脚本编程**

- \#! 是一个约定的标记，它告诉系统这个脚本需要什么解释器来执行
- 扩展名.sh
- echo用于向窗口输出文本

##### **变量**

运行shell时会同时存在三种变量

1. 局部变量：在脚本或命令中定义，仅在当前shell实例中有效，其他shell启动的程序不能访问局部变量；
2. 环境变量：所有的程序，包括shell启动的程序，都能访问环境变量，有些程序需要环境变量来保证其正常运行。必要的时候shell脚本也可以定义环境变量。
3. shell变量：shell变量是由shell程序设置的特殊变量。shell变量中有一部分是环境变量，有一部分是局部变量，这些变量保证了shell的正常运行。

~~~shell
#!/bin/bash
echo "Hello World !"
~~~

~~~shell
name="yoyo"   //定义变量
echo ${name}  //使用变量
readonly name //只读变量
unset name    //删除变量
~~~

##### **shell中的文件比较**

~~~shell
-e fileNane 如果fileName存在           为真
-d fileNane 如果fileName为目录         为真
-f fileNane 如果fileName为常规文件      为真
-L fileNane 如果fileName为符号链接      为真
-r fileNane 如果fileName可读           为真
-w fileNane 如果fileName可写           为真
-x fileNane 如果fileName可执行         为真
-s fileNane 如果fileName文件长度部位0   为真
-h fileNane 如果fileName文件为软连接    为真
fileName1 -nt fileName2 如果fileName1比fileName2新 为真
fileName1 -ot fileName2 如果fileName1比fileName2旧 为真 
~~~

##### **shell整数变量表达式**

~~~shell
-eq 等于
-ne 不等于
-gt 大于
-ge 大于等于
-lt 小于
-le 小于等于
~~~

##### **shell字符串变量表达式**

~~~shell
f  [ $a = $b ]                   如果string1等于string2，   则为真
                                 字符串允许使用赋值号做等号
if  [ $string1 !=  $string2 ]    如果string1不等于string2，则为真       
if  [ -n $string1  ]             如果string1 非空(非0），返回0(true)  
if  [ -z $string1  ]             如果string1 为空，则为真
if  [ $sting1 ]                  如果string1 非空，返回0 (和-n类似) 
~~~

##### **其他运算符**

~~~shell
逻辑非 !                   条件表达式的相反
if [ ! 表达式 ]
if [ ! -d $num ]               如果不存在目录$num


    逻辑与 –a                   条件表达式的并列
if [ 表达式1  –a  表达式2 ]


    逻辑或 -o                   条件表达式的或
if [ 表达式1  –o 表达式2 ]
~~~

##### **输入输出**

~~~shell
>> 和 >都是输出重定向
< 是输入重定向
~~~

| 命令            | 说明                                               |
| --------------- | -------------------------------------------------- |
| command > file  | 将输出重定向到 file。                              |
| command < file  | 将输入重定向到 file。                              |
| command >> file | 将输出以追加的方式重定向到 file。                  |
| n > file        | 将文件描述符为 n 的文件重定向到 file。             |
| n >> file       | 将文件描述符为 n 的文件以追加的方式重定向到 file。 |
| n >& m          | 将输出文件 m 和 n 合并。                           |
| n <& m          | 将输入文件 m 和 n 合并。                           |
| << tag          | 将开始标记 tag 和结束标记 tag 之间的内容作为输入。 |

##### **awk文本分析工具**

~~~shell
awk的作用是把文件逐行读入，（空格，制表符）为默认分隔符将每行切片，切开后再处理
命令格式：
awk [-F field-separator] 'commands' input-file(s)
[-F]分隔符可选
~~~

