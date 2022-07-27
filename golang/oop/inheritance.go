package main

import "fmt"

func main() {
	p := Person{
		Name: "ahian",
	}
	p.print()
	lawyer := Worker{
		Person: Person{Name: "ahian"},
		Job:    "lawyer",
	}
	lawyer.print()
}

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
