package encapsulation

import (
	"fmt"
	"testing"
)

type Employee struct {
	Id   string
	Name string
	Age  int
}

func TestCreateEmployeeObj(t *testing.T) {
	e := Employee{"0", "Bob", 20}
	e1 := Employee{Name: "Mike", Age: 30}
	e2 := new(Employee) // 返回指针
	e2.Id = "2"
	e2.Name = "Rose"
	e2.Age = 22
	t.Log(e)
	t.Log(e1)
	t.Log(e1.Id)
	t.Log(e2)
	t.Logf("e is %T", &e)
	t.Logf("e2 is %T", e2)
}

func (e Employee) String() string {
	fmt.Printf("Address is %x\n", &e.Name)
	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
}

//func (e *Employee) String() string {
//	fmt.Printf("Address is %x\n", &e.Name)
//	return fmt.Sprintf("ID:%s/Name:%s/Age:%d", e.Id, e.Name, e.Age)
//}

func TestStructsOperations(t *testing.T) {
	e := Employee{"0", "Bob", 20}
	fmt.Printf("Address is %x\n", &e.Name)
	t.Log(e.String())
	t.Logf("%T", e)
	t.Logf("%T", 10)
}
