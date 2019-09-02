# Go Content


## 什么是Context ?

Context通常被译作上下文，它是一个比较抽象的概念。
一般理解为程序单元的一个运行状态、现场、快照，而翻译中上下又很好地诠释了其本质，上下上下则是存在上下层的传递，上会把内容传递给下。

在Go语言中，程序单元也就指的是Goroutine。

每个Goroutine在执行之前，都要先知道程序当前的执行状态，通常将这些执行状态封装在一个Context变量中，传递给要执行的Goroutine中。
上下文则几乎已经成为传递与请求同生存周期变量的标准方法。

## Context包 

context包不仅实现了在程序单元之间共享状态变量的方法，同时能通过简单的方法，使我们在被调用程序单元的外部，通过设置ctx变量值，
将过期或撤销这些信号传递给被调用的程序单元。

Go1.7(当前是RC2版本)已将原来的golang.org/x/net/context包挪入了标准库中，放在$GOROOT/src/context下面。   
标准库中net、net/http、os/exec都用到了context。

context包的核心就是Context接口，其定义如下：
```go
package context
type Context interface {
    Deadline() (deadline time.Time, ok bool)
    Done() <-chan struct{}
    Err() error
    Value(key interface{}) interface{}
}
```
- Deadline会返回一个超时时间，Goroutine获得了超时时间后，例如可以对某些io操作设定超时时间。
- Done方法返回一个信道（channel），当Context被撤销或过期时，该信道是关闭的，即它是一个表示Context是否已关闭的信号。
- 当Done信道关闭后，Err方法表明Context被撤的原因。
- Value可以让Goroutine共享一些数据，当然获得数据是协程安全的。但使用这些数据的时候要注意同步，比如返回了一个map，而这个map的读写则要加锁。

Context接口没有提供方法来设置其值和过期时间，也没有提供方法直接将其自身撤销。
也就是说，Context不能改变和撤销其自身。那么该怎么通过Context传递改变后的状态呢？

## Context使用

无论是Goroutine，他们的创建和调用关系总是像层层调用进行的，就像人的辈分一样，而更靠顶部的Goroutine应有办法主动关闭其下属
的Goroutine的执行（不然程序可能就失控了）。

为了实现这种关系，Context结构也应该像一棵树，叶子节点须总是由根节点衍生出来的。

要创建Context树，第一步就是要得到根节点，context.Background函数的返回值就是根节点：
```
func Background() Context
```
该函数返回空的Context，该Context一般由接收请求的第一个Goroutine创建，是与进入请求对应的Context根节点，它不能被取消、没有值、也没有过期时间。它常常作为处理Request的顶层context存在。

有了根节点，又该怎么创建其它的子节点，孙节点呢？context包为我们提供了多个函数来创建他们：
```
func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc)
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
func WithValue(parent Context, key interface{}, val interface{}) Context
```
数都接收一个Context类型的参数parent，并返回一个Context类型的值，这样就层层创建出不同的节点。
子节点是从复制父节点得到的，并且根据接收参数设定子节点的一些状态值，接着就可以将子节点传递给下层的Goroutine了。




## 使用原则

使用Context的程序包需要遵循如下的原则来满足接口的一致性以及便于静态分析。
- 不要把Context存在一个结构体当中，显式地传入函数。Context变量需要作为第一个参数使用，一般命名为ctx；
- 即使方法允许，也不要传入一个nil的Context，如果你不确定你要用什么Context的时候传一个context.TODO；
- 使用context的Value相关方法只应该用于在程序和接口中传递的和请求相关的元数据，不要用它来传递一些可选的参数；
- 同样的Context可以用来传递到不同的goroutine中，Context在多个goroutine中是安全的；

在子Context被传递到的goroutine中，应该对该子Context的Done信道（channel）进行监控，一旦该信道被关
闭（即上层运行环境撤销了本goroutine的执行），应主动终止对当前请求信息的处理，释放资源并返回。
