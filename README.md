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

# 声明方式及赋值


## 变量
<table>
<thead><tr><th>java</th><th>go</th></tr></thead>
<tbody>
<tr><td>

```java
int num = 0; //int
```

</td><td>

```go
var num int // int
num = 0

num1 := 0 // int
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
<td>int8</td>
<td>2  byte</td>
 
</tr>

<tr>
<td>boolean</td>
<td>bool</td>
<td>4  byte</td>
 
</tr>

</tbody></table>

# 集合
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

# I/O
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


- spring-boot-web


</td><td>


- Gin


</td></tr>
</tbody></table>

# ORM 工具
<table>
<thead><tr><th>java</th><th>go</th></tr></thead>
<tbody>
<tr><td>


- MyBatis


</td><td>


- Gorm


</td></tr>
</tbody></table>
