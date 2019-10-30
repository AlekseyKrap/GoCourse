package fifo

import "fmt"

type Fifo struct {
	queue []string
	test  func()
}

func (f *Fifo) push(name string) {
	f.queue = append(f.queue, name)
}

func (f *Fifo) pop() (elem string) {
	if len(f.queue) < 1 {
		return
	}

	elem = f.queue[0]
	f.queue = f.queue[1:]
	return
}

func test() {
	fmt.Println("test")
}

func ExecFifo() {

	var anyFifo = Fifo{test: test}
	anyFifo.test()
	anyFifo.push("Vasya")
	anyFifo.push("Johan")
	anyFifo.pop()

	fmt.Println("anyFifo", anyFifo)

}
