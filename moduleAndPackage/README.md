# moduleAndPackage

#### 介绍
这是一个简单介绍Go语言模块和包的例子，模块是一个个.go的源文件，下面的命令中如果packagename相同，则可以理解为这些.go的文件处于同一个包下，一个包中往往包含着很多的模块。

```go
package packagename
```

当前目录下的go.mod用来管理moduleAndPackage这个项目，整个项目到架构如下

![](img/img001.jpg)

我们可以看到go.mod和main.go处于同一级目录下，当main.go中要调用apple.go中的函数时，需要用import语句声明该函数所处的项目位置和包的名字。

从上面的描述中我们也可以看出，项目名称和报的名称在声明的时候不一定要和文件夹的名称一致，但是为了方便，我们一般还是会给它们取相同的名字。