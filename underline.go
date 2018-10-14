/**
关于下划线发现的好玩的东西
 */
package main
/*
判断一个类型是否实现了某个接口
var _ 接口名 = Value
如果没实现就会编译报错
 */
type Dog struct {

}
func(d Dog)Say(){

}
type Foo interface {
	Say()
}
func main(){
	var _ Foo = Dog{}

}
/**
下划线可用于断言，忽略返回值，import不使用其中的数据时只执行init方法
 */