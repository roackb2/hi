package hi

import (
	"fmt"
	"github.com/roackb2/hi/utils"
	"io"
	"math"
	"math/cmplx"
	"math/rand"
	"net/http"
	"runtime"
	"time"
)

var (
	fb, twitter, yo bool       = true, true, false
	MaxInt          uint64     = 1<<64 - 1
	z               complex128 = cmplx.Sqrt(-5 + 12i)
	i8              int8
	i16             int16
	i32             int32
	i64             int64
	u8              uint8
	u16             uint16
	u32             uint32
	u64             uint64
	uptr            uintptr
	bt              byte
	ru              rune
	f32             float32
	f64             float64
	c64             complex64 = 5 - 2i
	c128            complex128
	s               string
)

func main() {
	fmt.Println("available CPU:", runtime.NumCPU())
	runtime.GOMAXPROCS(runtime.NumCPU())
	// runtime.GOMAXPROCS(1)
	var one, two int = 1, 2
	var hello string = "hello world"
	const pi = 3.1415926
	tryTree()
	tryExecutor()
	person := NewPerson()
	person.Print()
	person.ListSkills()
	worker := Worker(Person{name: "bob", age: 0, skills: []string{"java", "c++"}})
	worker.ListSkills()
	tryDefer(5)
	fmt.Println("OS", trySwitch())
	fmt.Println("Sqrt: ", Sqrt(24858))
	fmt.Println("for result: ", tryFor(10))
	fmt.Println("while result: ", tryWhile(10))
	fmt.Println("if result: ", tryIf(19))
	fmt.Println("complex:", c64)
	fmt.Println("pi: ", pi)
	fmt.Println("string: ", hello)
	fmt.Println("numbers: ", one, two)
	// fmt.Println(person{"Jay", 20})
	fmt.Printf(utils.Reverse("\n!dlrow olleh"))
	fmt.Println("The time is: ", time.Now())
	fmt.Println("Random: ", rand.Intn(10))
	fmt.Println("pi: ", math.Pi)
	fmt.Println("math.Nextafter(2,3): ", math.Nextafter(2, 3))
	fmt.Println("3 + 5: ", add(3, 5))
	fmt.Println("swap [uptown] [funk you up]:")
	fmt.Println(swap("uptown", "funk you up"))
	fmt.Println("sum and product of 3, 5: ")
	fmt.Println(sumAndProduct(3, 5))
	vars()
	types()
	trySwitchingMsg()
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello World!~")
	})

	http.ListenAndServe(":8080", nil)
}

func add(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func sumAndProduct(x, y int) (sum, product int) {
	sum = x + y
	product = x * y
	return
}

func vars() {
	hi, jay := 5, "jay"
	fmt.Println(fb, twitter, yo, hi, jay)
	fmt.Println()
}

func types() {
	const f = "%T(%v)\n"
	fmt.Printf(f, yo, yo)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, z, z)
}

func tryFor(max int) (result int) {
	result = 0
	for i := 0; i < max; i++ {
		result += i
	}
	return
}

func tryWhile(max int) (result int) {
	result = 0
	i := 0
	for i < max {
		result += i
		i++
	}
	return
}

func tryIf(input int) (result string) {
	if input < 10 {
		result = "thie result is less than ten"
	} else {
		result = "the result is greater than ten"
	}
	return
}

func trySwitch() (result string) {
	switch os := runtime.GOOS; os {
	case "darwin":
		result = "OS X"
	case "linux":
		result = "Linux"
	default:
		result = os
	}
	return
}

func tryDefer(x int) {
	defer fmt.Println("deferred x: ", x)
	fmt.Println("entering function")
	x++
	fmt.Println("after x++, x: ", x)
	for i := 0; i < 10; i++ {
		x++
		defer fmt.Println("x: ", x)
	}
}

func tryExecutor() {
	executor := NewExecutor("/")
	executor.exec("/bin/bash", "-c", "echo $HOME")
}

func trySwitchingMsg() {
	fstStop := make(chan bool)
	sndStop := make(chan bool)
	go tickMsg("hello", fstStop)
	go tickMsg("world", sndStop)
	time.Sleep(time.Second * 3)
	fstStop <- true
	time.Sleep(time.Second * 2)
	sndStop <- true
}

func tryTree() {
	ts := RandomPopulate(10)
	t0 := NewTree(3)
	t1 := NewTree(2)
	t2 := NewTree(4)
	t3 := NewTree(9)
	t4 := NewTree(14)
	t5 := NewTree(6)
	t0.appendLeft(t1)
	t0.appendRight(t2)
	t1.appendLeft(t3)
	t1.appendRight(t4)
	t2.appendLeft(t5)
	// target := t5.getRoot()
	// fmt.Println(target.value)
	rChan := make(chan string)
	start := time.Now()
	go ts.intJoin(rChan)
	<-rChan
	duration := time.Since(start)
	// fmt.Println(result)
	fmt.Println("join takes:", duration)
	t0.traverse(nil, func(v interface{}) {
		fmt.Println(v)
	}, nil)
	valueChan := make(chan interface{})
	go t0.chanTraverse(valueChan, nil, nil)
	for {
		value := <-valueChan
		if value == "done" {
			return
		} else {
			fmt.Println(value)
		}
	}
}
