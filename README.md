# j2go-cheatsheet

java 转 go 快速对照表，建议搭配 [smart TOC](https://chrome.google.com/webstore/detail/smart-toc/lifgeihcfpkmmlfjbailfpfhbahhibba) 工具阅读

这个速查表是为了给 *高级 Java 工程师* 快速熟悉 Golang 语言特性准备的。 

语言只是工具，都有各自的优势和不足，没必要抬高这个贬低那个，请不要做语言的奴隶！


> 说明：目录中带星号 * 的代表这部分内容是针对于 java 侧来说的，在 golang 领域可能没有此概念


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
基础数据类型是CPU可以直接进行运算的类型
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

# 集合
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
list1 = make([]T, len(xx));  // array list
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


# I/O
## 文件读写
<table>
<thead><tr><th>java</th><th>go</th></tr></thead>
<tbody>
<tr><td>

```java

```

</td><td>

```go

```

</td></tr>
</tbody></table>

# 字符串处理
<table>
<thead><tr><th>java</th><th>go</th></tr></thead>
<tbody>
<tr><td>

```java

```

</td><td>

```go

```

</td></tr>
</tbody></table>

# 并发
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

```

</td></tr>
</tbody></table>

# *面向对象编程
<table>
<thead><tr><th>java</th><th>go</th></tr></thead>
<tbody>
<tr><td>

```java

```

</td><td>

```go

```

</td></tr>
</tbody></table>

# 服务器
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
