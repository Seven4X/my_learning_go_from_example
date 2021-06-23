package foreach

import "testing"

type student struct {
	Name string
	Age  int
}

func pase_student() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	//stu 是copy的地址，三个取址相同，最终取址指向最后一个
	for _, stu := range stus {
		println(stu.Name)
		m[stu.Name] = &stu
	}
	for k, v := range m {
		println(k, "=>", v.Name)
	}
}

func TestFor(t *testing.T) {
	pase_student()
}
