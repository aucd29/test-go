Go 언어 - golang.org
====================
>Writer 황대길(huang@hanwha.com)

## 설치

1. git  -- git-scm.com/downloads에서 내려받아 설치한다
2. go   -- golang.org/dl에서 내려받아 설치한다

## 환경변수 설정

for macOS
```
echo 'export GOPATH=$HOME/go'        >> ~/.bash_profile
echo 'export PATH=$PATH:$GOPATH/bin' >> ~/.bash_profile
```

for Windows
```
컴퓨터 - 속성 - 고급 시스템 설정 - 환경 변수 - 사용자 변수에서
GOPATH에 %HOMEDRIVE%%HOMEPATH%\go를 넣고,
Path에 %GOPATH%\bin을 추가한다
```

## 예제 프로그램
1. 콘솔을 새로 열고 예제 프로그램을 설치한다:
```
go get github.com/golang/example/hello
```

+ 폴더 구조  - golang.org/doc/code.html3
  + 콘솔에서 go 문서 서버를 시작한다: (반드시 해야 하는 항목은 아님)
```
godoc -http :6060
```

2. 브라우저에서 http://localhost:6060/doc/code.html 페이지를 열어 둔다
3. 콘솔을 하나 더 열고 예제 프로그램을 확인한다

```
hello                 # 설치된 것을 확인한다
open $GOPATH          # Mac
explorer %GOPATH%     # Windows
```
```
bin -- 실행 파일 - main 프로그램과 그것이 import하는 패키지들
pkg -- 라이브러리 트리
src -- 소스 트리
```

4. 소스 코드
+ main 프로그램  
```
$GOPATH/src/github.com/golang/example/hello/hello.go
$GOPATH/bin/hello
```

```go
package main         // 패키지 이름이 main이면 main 프로그램

import (
	"fmt"  // a standard library package
	"github.com/golang/example/stringutil"
)

func main() {
	fmt.Println(stringutil.Reverse("!selpmaxe oG ,olleH"))
}
```


+ 라이브러리
```
$GOPATH/src/github.com/golang/example/stringutil/reverse.go
$GOPATH/pkg/darwin_amd64/github.com/golang/example/stringutil.a
```

```go
package stringutil   // 패키지 이름 = 폴더 이름 = 라이브러리 이름

// Reverse returns its argument string reversed rune-wise left to right.
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
```


+ 테스트 (*_test.go)  
```
$GOPATH/src/github.com/golang/example/stringutil/reverse_test.go
```

```go    
package stringutil   // 패키지 이름 = 폴더 이름

import "testing"

func TestReverse(t *testing.T) {   // 이름이 Test로 시작하는 함수가 실행된다
	for _, c := range []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	} {
		got := Reverse(c.in)
		if got != c.want {
			t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
```

## run, build, install
개별 빌드 테스트를 위해 bin 과 pkg 를 삭제 한다.
```
rm -rf $GOPATH/bin $GOPATH/pkg    # Mac
rd/s/q %GOPATH%\bin %GOPATH%\pkg  # Windows
```

파일 삭제 후 hello 를 빌드 하기 위해 해당 경로로 이동 한다.
```
cd %GOPATH%/src/github.com/golang/example/hello
```

hello 를 빌드 한다. 빌드 후 바로 실행하려면 run 을 빌드만 할 거면 build 를 입력 하면 된다.
```
go run hello.go                   # 빌드하여 실행한다
go build                          # 빌드하여 이 폴더에 넣는다
```

빌드가 완료되면 아래와 같이 실행 할 수 있다.
```
./hello                           # Mac
.\hello                           # Windows
```

빌드 후 bin 이나 pkg 에 복사하려면 install 을 이용하면 된다.
```
go install                        # bin 폴더와 pkg 폴더에 설치한다
```

빌드 정보를 보기 위해서는 -v (verbose) 옵션을 사용 한다.
```
go install -v                     # 새로 빌드된 것이 없다
```

resvers.go 에 rune 을 byte 로 변경 시 어떠한 점 이 다른지 파악하기 위해 코드르 수정 한다.
```
vi ../stringutil/reverse.go       # rune을 byte로 바꾼다
go install -v                     # 새로 빌드된 것을 본다
vi hello.go                       # oG를 o G로 바꾼다
go install -v                     # 새로 빌드된 것을 본다
hello                             # 실행한다
```

## test
```
cd $GOPATH/src/github.com/golang/example/stringutil   # Mac
cd %GOPATH%/src/github.com/golang/example/stringutil  # Windows
```

테스트를 위해서는 원하는 파일 경로로 이동한 뒤 go test 를 입력하면 된다.
코드는 Prefix 로 Test 를 붙여야 된다는걸 기억 하자.
```
go test                           # 테스트 실패
vi reverse.go                     # byte를 rune으로 바꾼다
go test                           # 테스트 성공
```

---
## A Tour of Go  - tour.golang.org
go tool tour

- Basics - Packages, variables, and functions
- Exported names  -- 다른 패키지에서 참조하는 이름은 대문자로 시작한다
```
type  이름 타입
const 이름 타입 = 상수
var   이름 타입 = 값
func  이름(. . .) . . . { . . . }   // can only be declared at the top level
```

##### Functions, Functions continued, Multiple results, Named return values

```go
func init() { . . . }   // import되는 패키지에 있는 init() 함수들이 실행된다
func main() { . . . }   // init() 함수들이 모두 실행된 후에 실행된다

func add(x, y int) int { return x + y }         // 값을 하나 넘긴다

func swap(x, y int) (int, int) { return y, x }	// 값을 두 개 넘긴다

func root(n int) (r int) {                      // Named return value
	for i := 1; i * i <= n; i++ {
		r = i
	}
	return
}
```

##### Variables, Variables with initializers, Short variable declarations
```go
var i int           // 값이 0이 된다
var i float64 = 5
var i = 5
i := 5              // Short variable declarations -- 함수 안에서 사용한다
```
##### Basic types, Zero values, Type conversions, Type inference
###### Basic types
```go
bool                                     -- false (기본값)
int  int8  int16  int32  int64           -- 0
uint uint8 uint16 uint32 uint64 uintptr  -- 0
float32 float64                          -- 0
complex64 complex128                     -- 0
string(UTF-8)                            -- ""
byte(= uint8) rune(=int32 for UTF-32)
```
###### Struct types
```go
	struct {
		이름 타입           // field
		.  .  .
	}
```  
###### Interface types
```go
	interface {
		이름(. . .) . . .   // method
		.  .  .
	}
```  
###### Function types   -- func(. . .) . . .
###### Pointer types    -- *타입
###### Array types      -- [...]타입
###### Slice types      -- []타입
###### Map types        -- map[키타입]값타입
###### Channel types    -- chan 타입
```go
i := 42           // Type inference   -- int
f := float64(i)   // Type conversion  -- int --> float64
u := uint(f)      // Type conversion  -- float64 --> uint

// Constants, Enumerations, iota, _
const Truth, World, c = true, "世界", 2 + 3i
const (SUN = iota; MON; TUE; WED; THR; FRI; SAT)  // iota = 0, 1, 2, . . .
const (_ = iota; K = 1 << (10 * iota); M; G)      // _ = blank identifier
fmt.Println(Truth, World, c, SUN, MON, TUE, K, M/K, G/M)

// String literals, Rune literals
st1, st2 := "C:\\go\\src", `C:\go\src`   // string 타입 (글자열)
r1, r2 := 'A', '가'                      // rune 타입 (글자)
fmt.Printf("`%s` %q %T %v %q %T %v %c\n", st1, st2, r1, r1, r1, r2, r2, r2)

// Struct variables, struct literals
type rectangle struct{ a, b int }
r1 := rectangle{1, 2}
r2 := rectangle{b: 10, a: 20}
r3 := rectangle{b: 100}
r4 := rectangle{}
q  := &r1     // q points to r1, q = address of r1
fmt.Println(r1, r2, r3, r4, q.a, r1.b)
```

###### Embedded structs
```go
type user struct {
	name, email string
}

type admin struct {
	user
	level string
}

u := user{"길동", "gildong@abc.com"}
a := admin{u, "대리"}

fmt.Println(u, a)
fmt.Println(u.name, u.email)
fmt.Println(a.name, a.email, a.level)
fmt.Println(a.user.name, a.user.email, a.level)

// Function variables, function literals
var f func(int, int) int
add := func(a, b int) int { return a + b }
f = add
fmt.Printf("%T -- %d %d\n", add, add(1, 2), f(3, 4))

// Closures  -- function body 밖의 변수를 참조하는 함수
fibonacci := func() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a + b
		return a
	}
}()

for i := 0; i < 15; i++ {
	fmt.Print(fibonacci(), " ")
}
fmt.Println()

goldenRatio := func() func() float64 {
	a, b := 0., 1.
	return func() float64 {
		a, b = b, a + b
		return b / a
	}
}()

for i := 0; i < 50; i++ {
	fmt.Println(goldenRatio())
}

// Pointers
var p *int  // nil, integer pointer
i := 42     // 42, integer
p = &i      // p points i, &i = address of i, *p == 42
fmt.Println(i, p, *p, &i)

// Arrays
var a1 [3]int                        // [0 0 0]
a2 := [3]int{5, 2}                   // [5 2 0]
a3 := [...]int{10, 20, 30, 40, 50}   // [10 20 30 40 50]
a4 := [3][2]int{{1, 2}, {10, 20}, {100, 200}}
fmt.Println(a1, a2, a3, a4, a4[2][0])

// Slices  -- references to underlying arrays
var s0 []int             // s0 == nil, len(s0) == 0, cap(s0) == 0
s1 := a3[1:4]            // [20 30 40], len(s1) == 3, cap(s1) == 4
s2 := []int{1, 2, 3}     // slice literal -- len(s2) == 3, cap(s2) == 3
s3 := make([]int, 2, 5)  // [0 0], len(s3) == 2, cap(s3) == 5
s4 := append(s3, 5, 8)   // [0 0 5 8], len(s4) == 4, cap(s4) == 5
fmt.Println(s0, s1, s2, s3, s4)

// Maps  -- map keys to values
var m0 map[int]int                     // m0 == nil
m0 = make(map[int]int)                 // m0 != nil
m0[5] = 100                            // [5:100]
m0[8] = 200                            // [5:100 8:200]
m1 := map[int]int{ 3:10, 6:15, 9:30 }  // map literal
delete(m1, 6)                          // [3:10 9:30]
e, ok := m0[8]                         // e == 200, ok = true
fmt.Println(m0, m1, e, ok, m0[5])

// Range  -- iterator: 인덱스와 값, 키와 값을 넘긴다.
sum := 0
for _, v := range s2 {  // slice
	sum += v
}
fmt.Println(sum)

for k, v := range m1 {  // map
	fmt.Printf("%d:%d ", k, v)
}
fmt.Println()

s := "A가나B다C"
for i, r := range s {   // string
	fmt.Printf("%d%q ", i, r)
}
fmt.Println(len(s))

for i, len := 0, len(s); i < len; i++ {
	fmt.Printf("%x ", s[i])
}
fmt.Println()
```
_____________

유니코드 관련 정보
- http://www.unicode.org/charts/
- https://ko.wikipedia.org/wiki/UTF-8


##### Basics - Flow control statements: for, if, else, switch and defer
```go
for i := 0; i < 10; i++ { . . . }
for sum < 100 { . . . }
for { . . . }

if x < 0 { . . . }
if v, ok = getValue(); ok { . . . }
if ... { . . . } else { . . . }

switch 식 {
case 값1, 값2:
	. . .
case 값3:
	. . .
	fallthrough
default:
	. . .
}

switch {
case n > 0:
	. . .
case n < 0:
	. . .
default:
	. . .
}
```

##### Methods and interfaces
```go
package main

import (
	"fmt"
	"math"
)

type circle struct {
	r float64
}

// Method with the value receiver
func (c circle) String() string {
	return fmt.Sprintf("circle(%v)", c.r)
}

func (c circle) area() float64 {
	return math.Pi * c.r * c.r
}

// Method with the pointer receiver
func (c *circle) scale(f float64) {
	c.r *= f
}

type rectangle struct {
	a, b float64
}

func (r rectangle) String() string {
	return fmt.Sprintf("rectangle(%v, %v)", r.a, r.b)
}

func (r rectangle) area() float64 {
	return r.a * r.b
}

func (r *rectangle) scale(f float64) {
	r.a *= f
	r.b *= f
}

type scalar float64

func (s scalar) String() string {
	return fmt.Sprintf("scalar(%v)", float64(s))
}

func (s scalar) norm() float64 {
	return math.Abs(float64(s))
}

func (s *scalar) scale(f float64) {
	*s = scalar(f * float64(*s))
}

type vector struct {
	x, y float64
}

func (v vector) String() string {
	return fmt.Sprintf("vector(%v %v)", v.x, v.y)
}

func (v vector) norm() float64 {
	return math.Sqrt(v.x * v.x + v.y * v.y)
}

func (v *vector) scale(f float64) {
	v.x *= f
	v.y *= f
}

type ipAddr [4]byte

func (a ipAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", a[0], a[1], a[2], a[3])
}

// Method는 receiver type이 정의된 패키지 안에서 정의한다
```

##### Interfaces -- set of method signatures
```go
//	fmt.Stringer interface
//	type Stringer interface {
//		String() string
//	}

type shape interface {
	area() float64
}

type normer interface {
	norm() float64
}

type scalable interface {
	scale(float64)
}

// Veriadic function, Empty interface, Type switch
func put(interfaces ...interface{}) {
	for _, i := range interfaces {
		switch v := i.(type) {
		case shape:
			fmt.Println(v, v.area())
		case normer:
			fmt.Println(v, v.norm())
		default:
			fmt.Println(v)
		}
	}
}

// Type assertion  -- 인터페이스.(타입)
func putA(interfaces ...interface{}) {
	for _, i := range interfaces {
		if v, ok := i.(shape); ok {
			fmt.Println(v, v.area())
		} else if v, ok := i.(normer); ok {
			fmt.Println(v, v.norm())
		} else {
			fmt.Println(i)
		}
	}
}

func scale(f float64, scalables ...scalable) {
	for _, s := range scalables {
		s.scale(f)
	}
}

func main() {
	c, r := circle{1}, rectangle{1, 2}
	s, v := scalar(-5), vector{3, 4}
	a    := ipAddr{127, 0, 0, 1}
	put(c, r, s, v, a)
	scale(10, &c, &r, &s, &v)
	putA(c, r, s, v, a)
}
```

##### Built-in error interface
```go
//	type error interface {
//		Error() string
//	}
```

##### Defer (문장) -- 함수 안에서 남은 문장이 모두 실행된 다음에 (문장)이 실행된다
```
func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)    // src는 새 변수
	if err != nil {
		return
	}
	defer src.Close()               // src.Close()는 남은 문장 실행 후 실행된다
	dst, err := os.Create(dstName)  // dst는 새 변수
	if err != nil {
		return
	}
	defer dst.Close()               // dst.Close()는 남은 문장 실행 후 실행된다
	return io.Copy(dst, src)
}	// Short variable declaration -- := 앞에 새 변수가 있으면 된다
```

##### Defer, Panic, Recover
######	Built-in functions: panic, recover
```go
//	func panic(v interface{})     // panic의 argument가
//	func recover() interface{}    // recover의 return 값이 된다

package main

import (
	"fmt"
	"math"
)

func f(n float64) {
	if n < 0 {
		panic(fmt.Sprintf("Invalid argument: %v", n))
	}
	fmt.Println(n, math.Sqrt(n))
	f(n - 1)
}

func main() {
	defer func() {
		if v := recover(); v != nil {
			fmt.Println(v)
		}
	}()
	f(3)
}
```

##### Concurrency
###### Goroutines and Channels <-- Communicating Sequential Processes (Tony Hoare 1978)
```go
fib  := make(chan int)      // fibonacci 수 채널
quit := make(chan int)      // 끝 알림 채널
go func() {                 // Goroutine
	a, b := 0, 1
	for {
		select {
		case fib <- b:      // fibonacci 수 보낸다
			a, b = b, a+b
		case <-quit:        // 끝 알림을 받는다
			return
		}
	}
}()
for i := 0; i < 10; i++ {
	fmt.Print(<-fib, " ")   // 받은 fibonacci 수를 출력한다
}
quit <- 0                   // 끝 알림을 보낸다
```

###### Close and Range
```go
goldenRatio := make(chan float64)   // 채널을 만든다
go func(tolerance float64) {        // Goroutine
	a, b := 0., 1.
	for g0, g := 1., 2.; tolerance < math.Abs(g0 - g); {
		a, b = b, a+b
		g0, g = g, b / a
		goldenRatio <- g
	}
	close(goldenRatio)              // 채널을 닫는다
}(.000000000000000001)
for g := range goldenRatio {        // 채널에서 받는다
	fmt.Println(g)
}
```

###### Buffered Channels
```go
ch := make(chan int, 2)
ch <- 1
ch <- 2
//ch <- 3   // deadlock
fmt.Println(<-ch)
fmt.Println(<-ch)
```
-----

### Writing Web Applications  - golang.org/doc/articles/wiki/

```go
package main

import (
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/name/", serveEcho)
	http.ListenAndServe(":8080", nil)
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

func serveEcho(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello " + strings.TrimPrefix(r.URL.Path, "/name/") + "!"))
}
```

### 강연, 블로그, 책

- Another Go at Language Design - Rob Pike
  - https://www.youtube.com/watch_popup?v=7VcArS4Wpqk

- Go Proverbs - Rob Pike
  - https://www.youtube.com/watch_popup?v=PAAkCSZUG1c

- Farewell Node.js: Leaving node.js land
  - https://medium.com/@tjholowaychuk/farewell-node-js-4ba9e7f3e52b#.hjv2xy4dy

- http://www.gopl.io/ch1.pdf
- http://www.allitebooks.com/the-way-to-go-a-thorough-introduction-to-the-go-programming-language/
- http://www.allitebooks.com/go-in-action/
- http://www.allitebooks.com/go-in-practice/
- http://www.allitebooks.com/go-web-programming/
