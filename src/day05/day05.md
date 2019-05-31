# 1,WaitGroup

WaitGroup是一个计数信号量，可以用来记录并维护运行的 goroutine。如果WaitGroup的值大于0，Wait方法就会阻塞。
使用defer声明在函数退出时调用Done方法,就可以使WaitGroup数-1

# 2，defer关键字

关键字defer会修改函数调用时机，在正在执行的函数返回时才真正调用defer声明的函数。
对这里的示例程序来说，我们使用关键字defer保证，每个 goroutine 一旦完成其工作就调用Done方法。

# 3,runtime.GOMAXPROCS()方法

控制逻辑处理器数目（相当于NIO的selecter或者Java线程池，会提供一个运算，但是同时间只会处理一个goroutine）