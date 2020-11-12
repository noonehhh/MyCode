### 1.数据类型

- 数字类型：

  python有四种基本数字类型：int(只有一种，表示长整数)，bool，float，complex

- 字符串:

  单双引号使用完全相同，无单独字符类型

  转义符\   使用r可以让\不发生转义

  \+ 连接字符串 * 重复字符串

  字符串索引两种方式，从左以0开始.从右以-1开始

  字符串截取 [头下标:尾下标:步长]

- Tuple(元组)

  写在()里

  元组里的元素不能被改变

  元组支持不同的数据类型

  同字符串一样，元组索引值以 0 为开始值，-1 为从末尾的开始位置

---------------------------------以上三种数据类型为不可改变数据类型--------------------------------

- List(列表)

   写在[]里

   同字符串一样，列表索引值以 0 为开始值，-1 为从末尾的开始位置

   列表可以被截取和索引，被截取后返回一个所需元素的新列表

   拼接列表用+

   python的列表支持不同的数据类型

   列表中的元素可以被改变

 string、list 和 tuple 都属于 sequence（序列）

- Set(集合)

  不可重复序列

  可用{}或set()函数创建集合，创建空集合必须用set()

  集合可以进行交 &、并 |、差 -、补 ^运算

- Dictionary（字典）

  k,v型数据结构

  用{}标识，无序的k,v集合

  k必须是不可变类型

  k必须是唯一的

  创建空字典用{}

**Python数据类型转换**

有时候，我们需要对数据内置的类型进行转换，数据类型的转换，你只需要将数据类型作为函数名即可。

以下几个内置的函数可以执行数据类型之间的转换。这些函数返回一个新的对象，表示转换的值。

| 函数                                                         | 描述                                                |
| ------------------------------------------------------------ | --------------------------------------------------- |
| [int(x [,base\])](https://www.runoob.com/python3/python-func-int.html) | 将x转换为一个整数                                   |
| [float(x)](https://www.runoob.com/python3/python-func-float.html) | 将x转换到一个浮点数                                 |
| [complex(real [,imag\])](https://www.runoob.com/python3/python-func-complex.html) | 创建一个复数                                        |
| [str(x)](https://www.runoob.com/python3/python-func-str.html) | 将对象 x 转换为字符串                               |
| [repr(x)](https://www.runoob.com/python3/python-func-repr.html) | 将对象 x 转换为表达式字符串                         |
| [eval(str)](https://www.runoob.com/python3/python-func-eval.html) | 用来计算在字符串中的有效Python表达式,并返回一个对象 |
| [tuple(s)](https://www.runoob.com/python3/python3-func-tuple.html) | 将序列 s 转换为一个元组                             |
| [list(s)](https://www.runoob.com/python3/python3-att-list-list.html) | 将序列 s 转换为一个列表                             |
| [set(s)](https://www.runoob.com/python3/python-func-set.html) | 转换为可变集合                                      |
| [dict(d)](https://www.runoob.com/python3/python-func-dict.html) | 创建一个字典。d 必须是一个 (key, value)元组序列。   |
| [frozenset(s)](https://www.runoob.com/python3/python-func-frozenset.html) | 转换为不可变集合                                    |
| [chr(x)](https://www.runoob.com/python3/python-func-chr.html) | 将一个整数转换为一个字符                            |
| [ord(x)](https://www.runoob.com/python3/python-func-ord.html) | 将一个字符转换为它的整数值                          |
| [hex(x)](https://www.runoob.com/python3/python-func-hex.html) | 将一个整数转换为一个十六进制字符串                  |
| [oct(x)](https://www.runoob.com/python3/python-func-oct.html) | 将一个整数转换为一个八进制字符串                    |

  

**2.书写规则**

- 缩进

  python用缩进表示代码块。缩进空格数可变，但同一代码块必须包含相同的缩进空格数，缩进不一致会导致运行错误。

- 行

python通常一行为一条语句，太长用\连接

**3.导包**

- 导入整个模块 

  import project

- 导入某个函数

  from project import func1

- 导入多个函数

  from project import func1,func2,func3

**4.变量定义**

   python中变量不需要声明，每个变量在使用前必须先赋值，赋值后该变量才会被创建。 

   python中变量没有类型，我们所说类型是变量指向内存中对象二点类型。

**关键字end可以用于将结果输出到同一行，或者在输出的末尾添加不同的字符，实例如下：**

~~~py
a, b = 0, 1 
while b < 1000:    
	print(b, end=',')    
	a, b = b, a+b    
	
输出：1,1,2,3,5,8,13,21,34,55,89,144,233,377,610,987,
~~~

**运算符（部分）**

  python中是没有&&及||这两个运算符的，取而代之的是英文and和or

  in，not in判断元素在/不在字符串，列表或元组中

  is ，not is 判断两个元素是否引用自同一个对象，即是否指向同一内存地址

**条件控制**

  if，elif，else

**循环语句**

  for-else，while-else

 **循环数字 range()**

**迭代器**

  字符串，列表或元组对象都可用于创建迭代器：

~~~py
>>> list=[1,2,3,4]
>>> it = iter(list)   # 创建迭代器对象 
>>> print (next(it))  # 输出迭代器的下一个元素 
1 
>>> print (next(it)) 
2
~~~

  迭代器对象可以使用常规for语句进行遍历：

~~~py
list=[1,2,3,4] 
it = iter(list)    # 创建迭代器对象 
for x in it:    
	print (x, end=" ")
~~~

用next

~~~py
import sys         # 引入 sys 模块  
list=[1,2,3,4] 
it = iter(list)    # 创建迭代器对象  
while True:    
	try:
    	print (next(it))
    except StopIteration:
    	sys.exit()
~~~

**生成器**

创建一个generator，有很多种方法。第一种方法很简单，只要把一个列表生成式的[]改成()，就创建了一个generator：

~~~py
>>> L = [x * x for x in range(10)] 
>>> L 
[0, 1, 4, 9, 16, 25, 36, 49, 64, 81] 
>>> g = (x * x for x in range(10)) 
>>> g
~~~

创建了生成器之后，迭代它一般用for循环

使用了 yield 的函数也是生成器（generator）

**继承**

~~~py
父类
class Person:
  def __init__(self, fname, lname):
    self.firstname = fname
    self.lastname = lname

  def printname(self):
    print(self.firstname, self.lastname)
    
子类
Class Student(Person):
    pass
    
Student 类和 Person 类有相同的 属性和方法   
为子类添加__init__函数后，子类不再继承父类的__init__函数
如果要保持对父类__init__函数的调用，要添加调用
Person.__init__(self, fname, lname)

super() 函数
它会使子类从其父继承所有方法和属性：
class Student(Person):
  def __init__(self, fname, lname):
    super().__init__(fname, lname)
使用 super() 函数，不必使用父元素的名称，它将自动从其父元素继承方法和属性。   

子类和父类有同名的方法，则子类将父方法覆盖  
~~~

**访问限制**

在方法或属性前后加__，该属性或方法就变成了私有变量，只有内部可以访问

**self关键字**

~~~py
class Student(object):
    def __init__(self, name, score):
        self.name = name
        self.score = score
~~~

-    __init__方法的第一参数永远是self，表示创建的类实例本身，因此，在__init__方法内部，就可以把各种属性绑定到self，因为self就指向创建的实例本身。
-   有了__init__方法，在创建实例的时候，就不能传入空的参数了，必须传入与__init__方法匹配的参数，但self不需要传，Python解释器会自己把实例变量传进去

**输入输出**

.format()函数用来格式话输入输出

**文件读写**

- 打开文件

file = open('path','r')   第二个参数代表打开文件的形式，r只读，w只写，a追加，r+读写

rb+按字节读取

- 读取文件

~~~py
按字符读取全部 file.read()
   file2 = open('./2.txt','r+',encoding='UTF-8')
   str = file2.read()
   file1 = open('./1.txt','r+')
   file1.write(''.join(str))
   file1.close()
   file2.close()
按行读取 file.readline()
   file2 = open('./2.txt','r+',encoding='UTF-8')
   file1 = open('./1.txt','r+')
   for line in file2:
      file1.write(line)
   file1.close()
   file2.close()
读取所有行 file.readlines()
  file2 = open('./2.txt','r+',encoding='UTF-8')
  str = file2.readlines()
  file1 = open('./1.txt','r+')
  file1.write(''.join(str))
~~~

- 创建文件

~~~py
linux系统
os.mknod(file)
windows系统
由于Windows系统没有node的概念，所以Windows创建文件用以下两种方法
open('path','w')
open('./1.txt','x')
~~~

- 创建文件夹

~~~py
import os
if not os.path.exists('./path'):
    os.mkdir('path')
else:
    print(True)
~~~

- 创建多级文件夹

~~~py
import os
if not os.path.exists('./dir/childDir'):
    os.makedirs('./dir/childDir')
else:
    print(True)
~~~

- 修改文件名

~~~py
import os
os.rename('./dir','./dir1')
~~~

- 批量修改文件名

~~~py
import os
paths = os.listdir('./dir')
i = 0
for path in paths:
    os.rename('./dir/' + path,'./dir/' + path+str(i))
    i += 1
~~~

- 删除文件/文件夹

~~~py
删除文件
import os
os.remove('./1.txt')
删除文件夹
import os
os.rmdir('./dir1/childDir0')
~~~

- os包

~~~py
os.access() 验证权限
~~~

**时间日期**

~~~py
时间戳 time.time()
时间格式化 time.strftime("%Y-%m-%d %H:%M:%S", time.localtime())
~~~

**JSON**

- 把 JSON 转换为 Python：

~~~py
import json

# 一些 JSON:
x =  '{ "name":"Bill", "age":63, "city":"Seatle"}'

# 解析 x:
y = json.loads(x)

# 结果是 Python 字典：
print(y["age"])
~~~

- 把 Python 转换为 JSON

~~~py
import json

# Python 对象（字典）：
x = {
  "name": "Bill",
  "age": 63,
  "city": "Seatle"
}

# 转换为 JSON：
y = json.dumps(x)

# 结果是 JSON 字符串：
print(y)
~~~

- 使用 indent 参数定义缩进数：

`json.dumps(x, indent=4)`

- 使用 separators 参数来更改默认分隔符：

`json.dumps(x, indent=4, separators=(". ", " = "))`

- 使用 sort_keys 参数来指定是否应对结果进行排序：

`json.dumps(x, indent=4, sort_keys=True)`

**Try Except**

- try 块允许测试代码块以查找错误。
- except 块允许处理错误。
- finally 块允许执行代码，无论 try 和 except 块的结果如何。

**多个异常**

~~~py
try:
  print(x)
except NameError:
  print("Variable x is not defined")
except:
  print("Something else went wrong")
~~~

**抛出（引发）异常**

~~~py
x = -1

if x < 0:
  raise Exception("Sorry, no numbers below zero")
~~~



**sorted排序方法**

[**https://www.runoob.com/python/python-func-sorted.html**](https://www.runoob.com/python/python-func-sorted.html)

~~~py
数组、列表、集合排序
b = [1,0,2,3,4,6,5,9,7,8]
print(sorted(b))
字典排序
a = {1:1,3:3,2:2,4:4,5:5,6:6,0:0}
print(sorted(a.items(),key = lambda x:x[0]))
~~~

**多线程**

- **threading模块**
- **包含的类**

~~~py
Thread：基本线程类
Lock：互斥锁
RLock：可重入锁，使单一进程再次获得已持有的锁(递归锁)
Condition：条件锁，使得一个线程等待另一个线程满足特定条件，比如改变状态或某个值。
Semaphore：信号锁。为线程间共享的有限资源提供一个”计数器”，如果没有可用资源则会被阻塞。
Event：事件锁，任意数量的线程等待某个事件的发生，在该事件发生后所有线程被激活
Timer：一种计时器
Barrier：Python3.2新增的“阻碍”类，必须达到指定数量的线程后才可以继续执行
~~~



- **常见方法和属性**

~~~py
current_thread()	返回当前线程
active_count()	返回当前活跃的线程数，1个主线程+n个子线程
get_ident()	返回当前线程
enumerater()	返回当前活动 Thread 对象列表
main_thread()	返回主 Thread 对象
settrace(func)	为所有线程设置一个 trace 函数
setprofile(func)	为所有线程设置一个 profile 函数
stack_size([size])	返回新创建线程栈大小；或为后续创建的线程设定栈大小为 size
TIMEOUT_MAX	Lock.acquire(), RLock.acquire(), Condition.wait() 允许的最大超时时间
~~~



- 创建线程

~~~py
import threading
def show(arg):
    time.sleep(1)
    print('thread' + str(arg) + "running ......")
    
if __name__ == "__main__":
    for i in range(10):
        t = threading.Thread(target = show,args(i))
        t.start()    
~~~



- 多线程的一个特点是，各个线程只顾执行自己的任务，不等待其他线程，如下：

~~~py
import time 
import threading

def dowaiting():
    print("start waiting : ",time.strtime("%H:%M:%S"))
    time.sleep(3)
    print("end waiting : ",time.strtime("%H:%M:%S"))
    
t = threading.Thread(target = dowaiting)
t.start()
t.sleep(1)

print("start job")
print("end job")    
~~~



- 如果希望主线程等待子线程，使用join方法

~~~py
import time 
import threading

def dowaiting():
    print("start waiting : ",time.strtime("%H:%M:%S"))
    time.sleep(3)
    print("end waiting : ",time.strtime("%H:%M:%S"))
    
t = threading.Thread(target = dowaiting)
t.start()
t.sleep(1)

print("start job")
print("end job")  
~~~



**线程锁**

- Lock 互斥锁
- RLock 可重入锁
- Semaphore 信号
- Event 事件
- Condition 条件
- Barrier “阻碍”

没有线程锁的时候，线程同时访问共享数据，会产生脏数据，如下例：

~~~py
import threading
import time

number = 0

def plus():
    global number       # global声明此处的number是外面的全局变量number
    for _ in range(1000000):    # 进行一个大数级别的循环加一运算
        number += 1
    print("子线程%s运算结束后，number = %s" % (threading.current_thread().getName(), number))

for i in range(2):      # 用2个子线程，就可以观察到脏数据
    t = threading.Thread(target=plus)
    t.start()


time.sleep(2)       # 等待2秒，确保2个子线程都已经结束运算。
print("主线程执行完毕后，number = ", number)

结果：
子线程Thread-2运算结束后，number = 1144974
子线程Thread-1运算结束后，number = 1181608
主线程执行完毕后，number =  1181608
产生了脏数据
~~~



- 互斥锁

独占锁，同一时刻只有一个线程可以访问共享数据

~~~py
import threading
import time

number = 0
lock = threading.Lock()

def func(lock):
    global number
    lock.acquire()  #加锁
    for _ in range(10000):
        number += 1
    print("子线程",number)
    lock.release() #释放锁

if __name__ == "__main__":
    for i in range(2):
        t = threading.Thread(target=func,args=(lock,))
        t.start()
    time.sleep(2)
    print("主线程",number)
~~~



- 信号Semaphore

类名：BoundedSemaphore。这种锁允许一定数量的线程同时更改数据，它不是互斥锁。比如地铁安检，排队人很多，工作人员只允许一定数量的人进入安检区，其它的人继续排队。

~~~py
import time
import threading

def run(n, se):
    se.acquire()  #加锁
    print("run the thread: %s" % n)
    time.sleep(1)
    se.release() #释放锁

# 设置允许5个线程同时运行
semaphore = threading.BoundedSemaphore(5)
for i in range(20):
    t = threading.Thread(target=run, args=(i,semaphore))
    t.start()
~~~



- 事件Event

运行机制，全局定义一个Flag，当Flag为false的时候，执行wait()方法就会阻塞，Flag为true的时候，所有线程放行。类似红绿灯，红灯全部阻拦，绿灯全部放行。

方法：

set()、wait()、clear()和is_set()。

- set()方法将Flag设为true
- clear()方法将Flag设为false
- wait()方法等待
- is_set():判断当前是否"绿灯放行"状态

~~~py
#利用Event类模拟红绿灯
import threading
import time

event = threading.Event()

def lighter():
    green_time = 5       # 绿灯时间
    red_time = 5         # 红灯时间
    event.set()          # 初始设为绿灯
    while True:
        print("\33[32;0m 绿灯亮...\033[0m")
        time.sleep(green_time)
        event.clear()
        print("\33[31;0m 红灯亮...\033[0m")
        time.sleep(red_time)
        event.set()

def run(name):
    while True:
        if event.is_set():      # 判断当前是否"放行"状态
            print("一辆[%s] 呼啸开过..." % name)
            time.sleep(1)
        else:
            print("一辆[%s]开来，看到红灯，无奈的停下了..." % name)
            event.wait()
            print("[%s] 看到绿灯亮了，瞬间飞起....." % name)

if __name__ == '__main__':

    light = threading.Thread(target=lighter,)
    light.start()

    for name in ['奔驰', '宝马', '奥迪']:
        car = threading.Thread(target=run, args=(name,))
        car.start()
~~~



**断言** 

~~~py
关键字 assert

assert expression
等价于
if not expression:
    raise AssertionError
~~~



**大小写切换**

~~~py
upper() 方法将字符串中的小写字母转为大写字母。
lower()方法将字符串中的大写字母转为小写字母。
~~~



**async与await**

https://www.cnblogs.com/xinghun85/p/9937741.html

**self关键字**

**指对象本身，self指变量在整个类中作为全局变量**

**断言**

- 在测试用例中，执行完测试用例后，最后一步是判断测试结果是pass还是fail，自动化测试脚本里面一般把这种生成测试结果的方法称为断言（assert）
- 基本断言方法

| 序号 | 断言方法                                | 断言描述                               |
| ---- | --------------------------------------- | -------------------------------------- |
| 1    | assertEqual(arg1, arg2, msg=None)       | 验证arg1=arg2，不等则fail              |
| 2    | assertNotEqual(arg1, arg2, msg=None)    | 验证arg1 != arg2, 相等则fail           |
| 3    | assertTrue(expr, msg=None)              | 验证expr是true，如果为false，则fail    |
| 4    | assertFalse(expr,msg=None)              | 验证expr是false，如果为true，则fail    |
| 5    | assertIs(arg1, arg2, msg=None)          | 验证arg1、arg2是同一个对象，不是则fail |
| 6    | assertIsNot(arg1, arg2, msg=None)       | 验证arg1、arg2不是同一个对象，是则fail |
| 7    | assertIsNone(expr, msg=None)            | 验证expr是None，不是则fail             |
| 8    | assertIsNotNone(expr, msg=None)         | 验证expr不是None，是则fail             |
| 9    | assertIn(arg1, arg2, msg=None)          | 验证arg1是arg2的子串，不是则fail       |
| 10   | assertNotIn(arg1, arg2, msg=None)       | 验证arg1不是arg2的子串，是则fail       |
| 11   | assertIsInstance(obj, cls, msg=None)    | 验证obj是cls的实例，不是则fail         |
| 12   | assertNotIsInstance(obj, cls, msg=None) | 验证obj不是cls的实例，是则fail         |

例：

~~~py
@async_wrapper
async def test_getUserWithPhoneNum(self):
    # 匹配用户
    code = '091IZ5000XQidK1XMh200yxNek2IZ50b'
    res = service.api.get_user_data_by_code(code)
    if res and 'openid' in res:
        self.assertEqual(True if res['openid'] else False, True)
    else:
        self.assertEqual(res, {})

    # 不存在的用户
    code = '0237mFkl2Jngx54eTqol2bWAur37mFkd'
    res = service.api.get_user_data_by_code(code)
    if res:
        self.assertEqual(True if res['openid'] else False, True)
    else:
        self.assertEqual(res, {})
~~~



*** 和 \**用法**

- 在列表、元组、字典前加* 

~~~py
a = [1, 2, 3]
b = (1, 2, 3)
c = {1: 1, 2: 2, 3: 3}

print(*a)
print(*b)
print(*c)

>>>1 2 3
>>>1 2 3
>>>1 2 3
在列表、元组、字典前加*号，会将列表、元组、字典拆分成一个一个的独立元素
例一：
a = [1, 2, 3]
def add(*args):
    print(type(args))
    for i in args:
        print(i)
add(a)
add(*a)
例二：
def func(*a, **b):
    print(a)
    print(b)
a= 3
b = 4
arr = (a, b)
kv = {"m": 1, "n": 2}
func(*arr, **kv)
~~~

- *args 和 * kwargs

~~~py
*args：接收若干个位置参数，转换成元组tuple形式
**kwargs：接收若干个关键字参数，转换成字典dict形式
ps:需要注意的是位置参数*args，一定要在关键字参数**kwargs前
~~~



**常用方法**

~~~py
dict.get(key, default=None) 返回指定键的值，如果值不在字典中返回默认值。
  例：
  dict = {'Name': 'Runoob', 'Age': 27}
>>>27
>>>Never
~~~

~~~py
dict()函数  创建一个字典。
  dict()                        # 创建空字典
  dict(a='a', b='b', t='t')     # 传入关键字
>>>{'a': 'a', 'b': 'b', 't': 't'}
 
 dict(zip(['one', 'two', 'three'], [1, 2, 3]))   # 映射函数方式来构造字典
>>>{'three': 3, 'two': 2, 'one': 1} 
 
 dict([('one', 1), ('two', 2), ('three', 3)])    # 可迭代对象方式来构造字典
>>>{'three': 3, 'two': 2, 'one': 1}
~~~

~~~py
getattr()  返回一个对象属性值。
getattr(object, name[, default])
参数：
object  --对象
name    --字符串，对象属性
default -- 默认返回值，如果不提供该参数，在没有对应属性时，将触发 AttributeError

例：
class A(object):
     bar = 1
     
a = A()
getattr(a, 'bar')        # 获取属性 bar 值
>>> 1

getattr(a, 'bar2')       # 属性 bar2 不存在，触发异常
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
AttributeError: 'A' object has no attribute 'bar2'

getattr(a, 'bar2', 3)    # 属性 bar2 不存在，但设置了默认值
3
~~~

~~~py
zip() 将可迭代的对象作为参数，将对象中对应的元素打包成一个个元组，然后返回
由这些元组组成的列表。
如果各个迭代器的元素个数不一致，则返回列表长度与最短的对象相同，利用 * 号操作符，
可以将元组解压为列表。

例：
a = [1,2,3]
b = [4,5,6]
c = [4,5,6,7,8]
zipped = zip(a,b)     # 打包为元组的列表
>>>[(1, 4), (2, 5), (3, 6)]
zip(a,c)
>>>[(1, 4), (2, 5), (3, 6)]
zip(*zipped)          # 与 zip 相反，*zipped 可理解为解压，返回二维矩阵式
[(1, 2, 3), (4, 5, 6)]
~~~

~~~py
urlparse.urlparse() 解析url
将url分为6个部分，返回一个包含6个字符串项目的元组：
协议、位置、路径、参数、查询、片段

import urlparse
url = urlparse.urlparse('http://www.baidu.com/index.html;user?id=5#comment')
>>>ParseResscheme='http', netloc='www.baidu.com', path='/index.html, params='user', query='id=5', fragment='comment'
~~~

~~~py
pick.dumps() 将对象obj对象序列化并返回一个byte对象
pick.loads() 从字节对象中读取被封装的对象
例：
import pickle
dict1 = dict(name='八岐大蛇',
             age=1000,
             sex='男',
             addr='东方',
             enemy=['八神', '草薙京', '神乐千鹤'])

data_dumps = pickle.dumps(dict1)
print(data_dumps)

data=pickle.loads(data_dumps )#从字节对象中读取被封装的对象，并返回
print(data)
~~~

~~~py
decode()
decode() 方法以 encoding 指定的编码格式解码字符串。默认编码为字符串编码。
语法规则：
str.decode(encoding='UTF-8',errors='strict')

一般使用直接 str.decode()
~~~

~~~py
hasattr() 判断对象是否包含对应的属性
hasattr(object, name)
例：
class Coordinate:
    x = 10
    y = -5
    z = 0
 
point1 = Coordinate() 
print(hasattr(point1, 'x'))                  >>>True
print(hasattr(point1, 'y'))                  >>>True
print(hasattr(point1, 'z'))                  >>>True
print(hasattr(point1, 'no'))  # 没有该属性    >>>False
~~~

~~~py
合并数组
from functools import reduce
import operator
arr = [[1,2,3]]
arr = reduce(operator.add, arr)
~~~

~~~py
#等分数组
def chunks(arr, n):
    return [arr[i:i+n] for i in range(0, len(arr), n)]
~~~

~~~py
获取当天剩余时间
86400 - int(time.time() + 8 * 3600) % 86400
~~~

~~~py
isinstance() 函数来判断一个对象是否是一个已知的类型，类似 type()。
isinstance() 与 type() 区别：
    type() 不会认为子类是一种父类类型，不考虑继承关系。
    isinstance() 会认为子类是一种父类类型，考虑继承关系。
    如果要判断两个类型是否相同推荐使用 isinstance()。
例：
>>>a = 2
>>> isinstance (a,int)
True
>>> isinstance (a,str)
False
>>> isinstance (a,(str,int,list))    # 是元组中的一个返回 True
True

class A:
    pass
class B(A):
    pass
isinstance(A(), A)    # returns True
type(A()) == A        # returns True
isinstance(B(), A)    # returns True
type(B()) == A        # returns False 
~~~

**执行外部命令并获取它的输出**

~~~py
subprocess 启动一个新的进程，并连接到他们的输入/输出/错误管道，从而获取返回值
subprocess推荐使用他的run方法，高级用法科直接使用Popen方法
方法：
run方法
https://www.runoob.com/w3cnote/python3-subprocess.html

执行命令并将结果以字符串返回
如果执行的命令以非零码返回，则会报错，如下：
try:
     out_bytes = subprocess.check_output(['cmd','arg1','arg2'])
except subprocess.CalledProcessError as e:
    out_bytes = e.output       # Output generated before error
    code      = e.returncode   # Return cod
    
sterr参数
默认情况下，check_output只返回标准输入到标准输出的值，如需同时收集标准输出
和错误输出，使用sterr参数：
out_bytes = subprocess.check_output(['cmd','arg1','arg2'],
                                    stderr=subprocess.STDOUT)    
                                    
timeout参数
执行超时机制命令，用timeout参数
try:
    out_bytes = subprocess.check_output(['cmd','arg1','arg2'], timeout=5)
except subprocess.TimeoutExpired as e:
~~~

git add -u;git commit -m 'adust';git push