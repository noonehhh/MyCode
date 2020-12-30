package main

import (
	"fmt"
)

/**
暂未理解。。。。
观察者模式
观察者模式（又被称为发布-订阅（Publish/Subscribe）模式，属于行为型模式的一种，它定义了一种一对多的依赖关系，让多个观察者对象同时监听某一个主体对象。
这个主体对象在状态变化时，会通知所有的观察者对象，使他们能够自动更新自己。
*/

type Customer interface {
	update()
}

type CustomerA struct {
}

func (*CustomerA) update() {
	fmt.Print("A 已收到")
}

type CustomerB struct {
}

func (*CustomerB) update() {
	fmt.Print("B 已收到")
}

// 报社 （被观察者)
type NewsOffice struct {
	customers []Customer
}

func (n *NewsOffice) addCustomer(customer Customer) {
	n.customers = append(n.customers, customer)
}

func (n *NewsOffice) newspaperCome() {
	// 通知所有客户
	n.notifyAllCustomer()
}

func (n *NewsOffice) notifyAllCustomer() {
	for _, customer := range n.customers {
		customer.update()
	}
}

func main() {

	customerA := &CustomerA{}
	customerB := &CustomerB{}

	office := &NewsOffice{}
	// 模拟客户订阅
	office.addCustomer(customerA)
	office.addCustomer(customerB)
	// 新的报纸
	office.newspaperCome()

}
