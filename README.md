# j2go-cheatsheet

java 转 go 快速对照表，建议搭配 [smart TOC](https://chrome.google.com/webstore/detail/smart-toc/lifgeihcfpkmmlfjbailfpfhbahhibba) 工具阅读

这个速查表是为了给 _高级 Java 工程师_ 快速熟悉 Golang 语言特性准备的。

语言只是工具，都有各自的优势和不足，在合适的场景用合适的语言

> 说明：目录中带星号 \* 的代表这部分内容是针对于 java 侧来说的，在 golang 领域可能没有此概念

**如果有错误的地方欢迎提 PR 或者 issue**

# 三方依赖仓库地址

<table>
<thead><tr><th>java</th><th>go</th></tr></thead>
<tbody>
<tr><td>

https://mvnrepository.com/

</td><td>

https://pkg.go.dev/

</td></tr>
</tbody></table>

# 包管理工具

<table>
<thead><tr><th>java</th><th>go</th></tr></thead>
<tbody>
<tr><td>

Maven、Gradle

</td><td>

Go Modules (go 1.13 设为默认)

</td></tr>
</tbody></table>

# 基础数据类型

基础数据类型是 CPU 可以直接进行运算的类型

<table>
<thead>
<tr>
<th>java</th>
<th>go</th>
<th>占用内存</th>
</tr>
</thead>
<tbody>
<tr>
<td>byte</td>
<td>byte/uint8</td>
<td>1  byte</td>
 
</tr>

<tr>
<td>short</td>
<td>int16</td>
<td>2  byte</td>
 
</tr>

<tr>
<td>int</td>
<td>int/int32/int64</td>
<td>java 4 byte; go 4/8  byte(Platform dependent)</td>
 
</tr>

<tr>
<td>long</td>
<td>int64</td>
<td>8  byte</td>
 
</tr>

<tr>
<td>float</td>
<td>float32</td>
<td>4  byte</td>
 
</tr>

<tr>
<td>double</td>
<td>float64</td>
<td>8  byte</td>
 
</tr>

<tr>
<td>char</td>
<td>rune</td>
<td>java 2 byte; go 4 byte</td>
 
</tr>

<tr>
<td>boolean</td>
<td>bool</td>
<td>1 byte</td>
 
</tr>

</tbody></table>

# 声明方式及赋值

## 变量

type: 类型

identifier: 标识符、变量名

value: 值

<table>
<thead><tr><th>java</th><th>go</th></tr></thead>
<tbody>
<tr><td>

```java

type identifier;
type identifier = value;

```

</td><td>

```go
var identifier type
var identifier type = value
var identifier = value
identifier := value
```

</td></tr>
</tbody>
</table>

## 常量

<table>
<thead><tr><th>java</th><th>go</th></tr></thead>
<tbody>
<tr><td>

```java
private static int NUM = 0; //int
```

</td><td>

```go
const num = 0
```

</td></tr>
</tbody>
</table>

# 循环控制

## for

<table>
<thead><tr><th>java</th><th>go</th></tr></thead>
<tbody>
<tr><td>

```java
for(int i = 0; i < arr.length; i++){
    //...
}
```

</td><td>

```go
for i := 0; i < len(arr); i++ {
    //...
}
```

</td></tr>
</tbody>
</table>

## for 增强

<table>
<thead><tr><th>java</th><th>go</th></tr></thead>
<tbody>
<tr><td>

```java
String s = "abc";
for(int i : s.toCharArray()){
    //...
}
```

</td><td>

```go
s := "abc"

for i,v := range s {
    //...
}
```

</td></tr>
</tbody>
</table>

## while

<table>
<thead><tr><th>java</th><th>go</th></tr></thead>
<tbody>
<tr><td>

```java
int i = 0;
while(i < 10){
    //...
}
```

</td><td>

```go
i : = 0
for i < 10 {
    i++
}
```

</td></tr>
</tbody>
</table>

## switch

<table>
<thead><tr><th>java</th><th>go</th></tr></thead>
<tbody>
<tr><td>

```java
int i = 0;
switch(i){
case 1:
    System.out.println("1");
case 2:
    System.out.println("2");
case 3:
    System.out.println("3");
default:
    System.out.println("0");
}
```

</td><td>

```go
i := 0
switch i {
case 1:
    fmt.Println("1")
case 2:
    fmt.Println("2")
case 3:
    fmt.Println("3")
default:
    fmt.Println("0")
```

</td></tr>
</tbody>
</table>

# 集合

## Array

<table>
<thead><tr><th>java</th><th>go</th></tr></thead>
<tbody>
<tr><td>

```java
int[] arr = new int[10];
 Arrays.sort(arr);// sort 默认升序
```

</td><td>

```go
arr := [10]int{} // array 非常用
sli := make([]int, 10) // slice
sort.Ints(sli) // sort []int
sort.Slice(sli, func(a, b int) bool { return sli[a] < sli[b] })
sort.SliceStable(sli, func(a, b int) bool { return sli[a] < sli[b] }) // 如果元素相同保持当前位置不变


```

</td></tr>
</tbody></table>

## List

<table>
<thead><tr><th>java</th><th>go</th></tr></thead>
<tbody>
<tr><td>

```java
List<T> list1 = new ArrayList<>();
List<T> list2 = new LinkedList<>();
```

</td><td>

```go
list1 := make([]T, len(xx));  // array list
var list2 list.List // linked list
list2.pushBack()
for e := list2.Front(); e != nil; e=e.Next() {
        fmt.Println(e.Value.(xx))
    }
```

</td></tr>
</tbody></table>

## Map

<table>
<thead><tr><th>java</th><th>go</th></tr></thead>
<tbody>
<tr><td>

```java
Map<T,T> map = new HashMap<>();
map.put(T,T);
T t = map.get(T);
```

</td><td>

```go
maps := make(map[T]T)
maps[T] = T
x, ok := maps[T]
if ok {
    //do somethings...
}

```

</td></tr>
</tbody></table>

## Set

<table>
<thead><tr><th>java</th><th>go</th></tr></thead>
<tbody>
<tr><td>

```java
Set<T> set = new HashSet<>();
```

</td><td>

```go
type void struct{}
var x void
set := make(map[T]void)
```

</td></tr>
</tbody></table>

## Stack

<table>
<thead><tr><th>java</th><th>go</th></tr></thead>
<tbody>
<tr><td>

```java
      Stack<Integer> stack = new Stack<>();
       stack.push(1);
       if(!stack.isEmpty()){
           int top = stack.peek();// inspect
           int topEle = stack.pop();// pop
       }
```

</td><td>

```go
	stack := make([]int, 0)
	stack = append(stack, 1)
	stack = append(stack, 2)
	for len(stack) > 0 {
		topEle := stack[len(stack)-1] // inspect
		fmt.Println("top element is: ", topEle)
        stack[len(stack)-1] = "" // 避免内存泄漏
		stack = stack[:len(stack)-1] // pop
		fmt.Println("current statck size: ", len(stack))

	}
	stack = append(stack, 3)
	fmt.Println(len(stack))
	stack = stack[:0] // flush
	fmt.Println(len(stack))
```

</td></tr>
</tbody></table>

# I/O

## 文件操作

### 读取文件

<table>
<thead><tr><th>java</th><th>go</th></tr></thead>
<tbody>
<tr><td>

```java
 byte[] bytes = Files.readAllBytes(Paths.get(filePath));
```

</td><td>

```go
f, err := ioutil.ReadFile("file path")//f is []byte

```

</td></tr>
</tbody></table>

## 监听命令行输入
<table>
<thead><tr><th>java</th><th>go</th></tr></thead>
<tbody>
<tr><td>

```java
       Scanner scanner = new Scanner(System.in);
        while (scanner.hasNext()) {
            System.out.println(scanner.next());
        }
```

</td><td>

```go
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		fmt.Println(input.Text())
	}

```

</td></tr>
</tbody></table>

## 获取命令行参数
<table>
<thead><tr><th>java</th><th>go</th></tr></thead>
<tbody>
<tr><td>

```java
       Scanner scanner = new Scanner(System.in);
        while (scanner.hasNext()) {
            System.out.println(scanner.next());
        }
```

</td><td>

```go
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		fmt.Println(input.Text())
	}

```

</td></tr>
</tbody></table>

# String

## 分割字符串

<table>
<thead><tr><th>java</th><th>go</th></tr></thead>
<tbody>
<tr><td>

```java
String[] strs = "a,b,c".split(",");
```

</td><td>

```go
// []string
strs := strings.Split("a,b,c", ",")
```

</td></tr>
</tbody></table>

## 是否以 prefix 开头

<table>
<thead><tr><th>java</th><th>go</th></tr></thead>
<tbody>
<tr><td>

```java
 "".startsWith(prefix);
 "".endsWith("suffix");//查找后缀
```

</td><td>

```go
var s string
strings.HasPrefix(s, "prefix")
```

</td></tr>
</tbody></table>

## 是否以 suffix 开头

<table>
<thead><tr><th>java</th><th>go</th></tr></thead>
<tbody>
<tr><td>

```java
 "".endsWith("suffix");
```

</td><td>

```go
var s string
strings.HasSuffix(s,"suffix")
```

</td></tr>
</tbody></table>

## 判断是否相等

<table>
<thead><tr><th>java</th><th>go</th></tr></thead>
<tbody>
<tr><td>

```java
 "".equals("");
```

</td><td>

```go
str1 == str2
```

</td></tr>
</tbody></table>

# \*面向对象编程

面向对象是以对象为核心向外拓展的，可以理解为现实环境中的映射。在 java 中表示为 Class，在 golang 中表示为 struct。在设计时会将对象能力通过方法的形式整合在一起并通过权限控制来加强封装，使用接口 interface 来隐式实现多态。

## 函数

函数是能给调用者返回一些需要的值，可以在任何地方使用

<table>
<thead><tr><th>java</th><th>go</th></tr></thead>
<tbody>
<tr><td>

```java
class Door{
    private String size;
    public void open(){
        System.out.println("door opened");
    }

    public String getSize() {
        return size;
    }

    public void setSize(String size) {
        this.size = size;
    }
}

```

</td><td>

```go
type Door struct {
  size string
}
func open(){
  fmt.Println("door opened")

}


```

</td></tr>
</tbody></table>

## 方法

方法是被定义在类内部，能去修改类的属性，注重类的概念

<table>
<thead><tr><th>java</th><th>go</th></tr></thead>
<tbody>
<tr><td>

```java
class Door{
    private String size;
    public void open(){
        System.out.println("door opened");
    }

    public String getSize() {
        return size;
    }

    public void setSize(String size) {
        this.size = size;
    }
}
```

</td><td>

```go

type Door struct {
  size string
}
func (door *Door) setSize(){
 door.size = "big"
}
func (door *Door) getSize(){
  fmt.Println("door size is:",door.size)
}


```

</td></tr>
</tbody></table>

# 泛型
<table>
<thead><tr><th>java</th><th>go</th></tr></thead>
<tbody>
<tr><td>

```java
public static <D> void print(D data){
        System.out.println("your data is " + data);
    }
```

</td><td>

```go
func print[D any](data D) {
	fmt.Print("your data is: ", data)
}
```

</td></tr>
</tbody></table>

# 接口

golang 的接口是隐式实现

<table>
<thead><tr><th>java</th><th>go</th></tr></thead>
<tbody>
<tr><td>

```java
public class TestInterface {
    public static void main(String[] args) {
        System.out.println("light");
        CarLight carLight = new CarLight();
        carLight.turnOn();
        carLight.turnOff();

    }
}

interface Light {
    void turnOn();

    void turnOff();
}

class CarLight implements Light {

    @Override
    public void turnOn() {
        System.out.println("turn on");
    }

    @Override
    public void turnOff() {
        System.out.println("turn off");
    }
}
```

</td><td>

```golang
type light interface {
	TurnOn()
	TurnOff()
}

type CarLight struct{}

func (carLigth CarLight) TurnOn() {
	fmt.Println("trun on")
}
func (carLight CarLight) TurnOff() {
	fmt.Println("turn off")
}
func main() {
	var carLight CarLight
	fmt.Println("light")
	carLight.TurnOn()
	carLight.TurnOff()

}
```

</td></tr>
</tbody></table>

## 封装

java 和 go 都能够通过一定的规则来控制程序的可见性 scope。
可见性由低到高：

> java: default>protect>private>public

> go: 小写开头(private) > 大写开头(public)

golang 中大小写可以应用到常量、变量、类型、函数、结构体中
举例：

```go
type User struct{...} User 能够被其他包内的代码访问
type user struct{...} user 只能包内访问
```

<table>
<thead><tr><th>java</th><th>go</th></tr></thead>
<tbody>
<tr><td>

</td><td>

</td></tr>
</tbody></table>

## 继承

<table>
<thead><tr><th>java</th><th>go</th></tr></thead>
<tbody>
<tr><td>

```java
class Person {
    private String name;

    public Person(String name) {
        this.name = name;
    }

    public void print() {
        System.out.println("my name is " + name);
    }


}

class Worker extends Person {
    private final String job;

    public Worker(String name, String job) {
        super(name);
        this.job = job;
    }

    @Override
    public void print() {
        System.out.println("my job is " + job);
    }
}
```

</td><td>

```go
type Person struct {
	Name string
}

func (person *Person) print() {
	fmt.Println("my name is ", person.Name)
}

type Worker struct {
	Person        // 继承父类
	Job    string // 子类新增
}

// override
func (worker *Worker) print() {
	fmt.Println("my job is ", worker.Job)
}
```
</td></tr>
</tbody></table>

## 多态
多态借助于 interface 来实现
<table>
<thead><tr><th>java</th><th>go</th></tr></thead>
<tbody>
<tr><td>

</td><td>

</td></tr>
</tbody></table>

# 并发

## 等待所有线程工作完成

非原子操作，可使用 atomic 进行累加计算
<table>
<thead><tr><th>java</th><th>go</th></tr></thead>
<tbody>
<tr><td>

```java
 static int num = 0;
    public static void main(String[] args) throws InterruptedException {
        CountDownLatch countDownLatch = new CountDownLatch(100);
        for (int i = 0; i < 10; i++) {
            new Thread(() -> {
                for (int j = 0; j < 10; j++) {
                    new Thread(()->{
                        add(countDownLatch);
                    }).start();
                }
            }).start();
        }
        countDownLatch.await();
        System.out.println(num);
    }
    private static void add(CountDownLatch countDownLatch){
            num++;
            countDownLatch.countDown();
    }
```

</td><td>

```go
var num = 0
// 计数器
func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		wg.Add(1) // 计数器加 1
		go add(&wg)
	}
	wg.Wait() // 等待计数器归零
	fmt.Println(num)
}
func add(wg *sync.WaitGroup) {
	num++
	wg.Done() // 计数器减 1
}
```

</td></tr>
</tbody></table>

## 锁

<table>
<thead><tr><th>java</th><th>go</th></tr></thead>
<tbody>
<tr><td>

```java
Lock lock = new ReentrantLock();
if(lock.tryLock()){
    try{
        // do something...
    } finally{
        lock.unlock();
    }
}

```

</td><td>

```go
func main() {
	var lock sync.Mutex
	for i := 0; i < 10000; i++ {
		go addByLock(&lock)
	}
	time.Sleep(10)
	fmt.Println(n)
}
func addByLock(lock *sync.Mutex) {
	lock.Lock()
	n++
	lock.Unlock()
}
```

</td></tr>
</tbody></table>

## Channel

<table>
<thead><tr><th>java</th><th>go</th></tr></thead>
<tbody>
<tr><td>

```java
无
```
</td><td>

```go
channel := make(chan string, 1)
go func() {
channel <- "add"
}()
go func() {
val := <-channel
fmt.Println(val)
}()
time.Sleep(time.Second * 2)
```
</td></tr>
</tbody></table>


# 序列化
## json
<table>
<thead><tr><th>java</th><th>go</th></tr></thead>
<tbody>
<tr><td>

```java
// jackson api
ObjectMapper objectMapper = new ObjectMapper();  
String jsonStr = objectMapper.writeValueAsString(obj);
```
</td><td>

```go
	user := &user{Name: "ahian", Age: 18}
	b, err := json.Marshal(user)
	if err != nil {
		fmt.Errorf("could not marshal json: %s\n", err)
		return
	}
	fmt.Println(string(b))
```
</td></tr>
</tbody></table>

# 时间处理工具

<table>
<thead><tr><th>java</th><th>go</th></tr></thead>
<tbody>
<tr><td>

```java
        SimpleDateFormat simpleDateFormat = new SimpleDateFormat("yyyy-MM-dd HH:mm:ss");
        String yearMonthDay = simpleDateFormat.format(new Date());
```

</td><td>

```go
	now := time.Now() // 当前时间 默认格式：2006-01-02 15:04:05.999999999 -0700 MST
	fmt.Println(now)
	format := now.Format("2006-01-02 15:04:05") // 格式化 年月日时分秒
	fmt.Println(format)
```


</td></tr>
</tbody></table>

# 服务器
指对外提供服务的程序，一般是通过 REST 方式访问
<table>
<thead><tr><th>java</th><th>go</th></tr></thead>
<tbody>
<tr><td>

</td><td>

</td></tr>
</tbody></table>

# ORM 工具

<table>
<thead><tr><th>java</th><th>go</th></tr></thead>
<tbody>
<tr><td>

- MyBatis
- Spring Data

</td><td>

- Gorm

</td></tr>
</tbody></table>
