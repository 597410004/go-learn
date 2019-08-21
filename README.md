# go-learn
## 开发工具
GoLand
## char2
### constant_test
常量定义: 多个常量定义并递增
### fib
常量定义的多种方式:
1. 最常用方式为: a:=1
2. 方式有 var a int=1;var a=1
3. 一行可以定义多个变量 a,b=b,a
## char3
条件:
1. for n<15 相当于while(n<15)
2. if可以定义两个值,比如: if a,flag;flag 表示 a为一个变量,flag为条件判断
3. switch为常量时,可以有多个值
### char4
#### array_test
##### 数组定义方式
1. ``arr1:=[5]int{1,2,3,4,5}`` 定长数组
2. ``var arr [3]int``  初始化数组
2. ``arr2:=[...]int{1,2,3,4,5,6,7,8,9}`` 不定长数组
##### 数组遍历
1. 遍历索引和值

```
for index,value:=range arr2{
    fmt.Println(index,value) 输出index和value
}
```

2. 遍历值

```
for _,value:=range arr2{
          fmt.Println(index,value) 使用占位符
      }
```
##### 数组切片
1. 切片方式不支持负数,如``arr[-2]`` 会报错
2. 切片如果既有开头又有结尾,则包含结尾不包含开头,如``arr[2:3]`` 包含第四位不包含第三位
3. ``arr[:3]`` 返回从0为到第四位
4. ``arr[2:]`` 返回从第三位开始值
##### 切片
1. 切片与数组区别:
   1. 切片可以无限伸缩
   2. 切片无法相互比较
   3. 初始化时是否指定长度,指定长度为数组
2. 切片声明方式:
    1. ``s:=int[]{}``
    2. ``var s0 int[]``
    3. ``s2:=make([]int,2,3)``
3. 切片容量与长度区别:
    1. 长度相当于默认值个数
    2. 切片容量可以无限延伸,每次延伸个数是上一次个数的2倍
4. 切片容量内存共享: 
    1. 如果另一个切片引用原始切片,切片改变值会改变原始切片值          
### char5
1. map初始化方式:
   1. ``m:=map[int]int{1:1,2:2}``
   2. ``m2 := map[int]int{}``
   3. ``m3 := make(map[int]int, 10)``
2. print打印区别:
    1. Log 和 Println 为正常打印,用``,``分隔
    2. Logf和 Printlnf为格式化打印,一般用``%d``,通用为``%v``(可以自己推测类型) 
3. 判断map中是否有值
    1. map在初始化以后,即使不存在key,默认值也为0
    2. 判断方法: ```if v,isExsits:=map[3];isExists{}```
4. 遍历map

```
for k,v :=range m{

    fmt.Println(k,v)
    
}
```
5. map中可以添加函数模仿工厂模式
```
func TestMapFactory(t *testing.T) {
	m:=map[int]func(number int)int{}
	m[1]= func(number int) int {
		return number
	}
	m[2]= func(number int) int {
		return number*number
	}
	t.Log(m[1](5),m[2](5)
	)
}
```    
### char6
1. 函数别名

``type Fun func(opt int)int`` 则这个函数别名为Fun
2. 函数定义
    1. ``func(opt int )`` 没有返回值,只有一个int参数-
    2. ``func(opt int)(int,int)`` 多个返回值
    3. ``func(opts...int)`` 不定参数
    
3. 类似try,catch语法
    1. 经常与defer函数一起使用
    2. 一般用于宕机后恢复操作
    3. 
    ```
     defer func() { // 必须要先声明defer，否则不能捕获到panic异常
          fmt.Println("d")
          if err := recover(); err != nil {
             fmt.Println(err) // 这里的err其实就是panic传入的内容,可以在这里写恢复代码
          }
          fmt.Println("e")
       }()
    ```
4. 类似final函数   
即使系统报错,defer里面的函数也会执行完才退出,一般用于释放资源
```
  	defer func() {
  		fmt.Println("清空资源")
  	}()
  	fmt.Println("--------")
  	panic("err")
```       
### char7
对象的定义
#### part1
1. 定义
     ```
    type Student struct {
    	Id string
    	Name string
    }
    ```    
2. 初始化
    ```
    s:=Student("1","hdd")
    
    ```
3. 添加方法
    ```
    func (s*Student)String()string{
       return fmt.Sprintf("name is %v - value is %v",s.Name,s.Id)
    }

    ```
#### part2
1. 匿名字段
对象类,可以是基本类型如int
2. 修改匿名字段属性
直接修改
#### part3
匿名字段中参数与对象参数相同时,修改匿名参数字段需使用. ,参考part3代码
#### part4
自定义匿名字段,参考代码     
### char8
#### 扩展与复用
1. 扩展

    1. 接口定义
   
   ```
   type Pet struct{
    } 
   ```  
   
     2. 接口方法
   
   ```
    func (p *Pet)Speak(){
    fmt.Println("....")
    }
    ```
    
     3. 扩展接口
   
   ```
    type Dog struc{
     p*Pet
    }
    ```
    
    4. 复用接口方法
    ```
    func(d *Dog)Speak(){
      d.p.Speak()
    }
    ``` 
    5. 扩展接口方法
    ```
    func(d *Dog)Speak(){
        fmt.Println("Wang Wang")
    }
 
    ```
2. 复用

    1. 定义
    ```
    type Dog struc{
     Pet
    }
    ``` 
    2. 可以直接复用父类方法

3. 扩展和复用本质区别:
    1. 扩展必须实现父类方法,复用可以直接引用父类方法 
    2. 扩展和复用都可以复写父类方法,如果复用时不重写父类方法,则引用的是父类原始方法  
#### 接口

1. 定义

```
type Base interface{
Connect() string
}
```

2. 实现接口

```
type BaseImpl struc{
}
```   

3. 实现接口方法

```
func(b *BaseImpl)Connect()string{
    return "connect"
}
```  
### char9
错误处理机制:
1. 定义:`` var GreaterError=errors.New(text) ``
2. 参考代码

### char10
包导进导出:
1. 导进导出的包必须在环境变量中设置
2. 引用的包必须是go文件,不能是测试文件
3. 导出的变量或者方法首字母必须大写

### char11
glide管理依赖包:
1. 依赖: `` go get-t github.com/*** ``
2. 初始化 ``glide init``
3. 下载依赖 `` glide install ``
4. windows使用glide bug解决:参考笔记

### char12

1. 协程(一个协诚可以管理多个线程或一个线程)
    ```
    go func(){
    }()
    ```
2. 协程在操作同一对象时是不安全的
    ```
    func TestCount(t *testing.T) {
    	count := 0
    	for i := 0; i < 500; i++ {
    		go func() {
    			count++
    		}()
    	}
    	// count不等于499就会退出
    	t.Logf("count is :%v", count)
    }
    ```
3. 协程加锁
    ```
    func TestSafeCount(t *testing.T) {
    	// 同步锁
    	var mul sync.Mutex
    	count := 0
    	for i := 0; i < 900; i++ {
    		go func() {
    			defer func() {
    				// 最后释放锁
    				mul.Unlock()
    			}()
    			// 加锁
    			mul.Lock()
    			count++
    		}()
    	}
    	//如果去掉等待时间则不会显示正常结果
    	time.Sleep(1 * time.Second)
    	t.Logf("count is :%v", count)
    }
    ```
4. 协程加锁并等待任务完成
    ```
    func TestSafeCountWait(t *testing.T) {
    	// 同步锁
    	var mul sync.Mutex
    	var group sync.WaitGroup
    	count := 0
    	for i := 0; i < 900; i++ {
    		group.Add(1)
    		go func() {
    			defer func() {
    				// 最后释放锁
    				mul.Unlock()
    			}()
    			// 加锁
    			mul.Lock()
    			count++
    			group.Done()
    		}()
    	}
    	group.Wait()
    	//如果去掉等待时间则不会显示正常结果
    	t.Logf("count is :%v", count)
    }
    ```
5. WaitGroup用于等待一组线程的结束。父线程调用Add方法来设定应等待的线程的数量。
每个被等待的线程在结束时应调用Done方法。同时，主线程里可以调用Wait方法阻塞至所有线程结束。     
### char13
channel 通道用来传递消息
1. 初始化 `` retCh:=make(ch,string)``
2. 有buffer的channel`` reCh:=make(ch,string,1)``  
3. 缓冲区区别:在channel中去消息时如果有缓冲区不用等待,没有缓冲区则必须等待channel消息取出来才执行别的   
4. 读写操作
```
// write to channel
ch <- x

// read from channel
x <- ch

// another way to read
x = <- ch
```
### go-cli
go 命令行编写工具
#### part1
基本使用
#### part2
自定义命令
#### part3
利用变量来自定义命令
