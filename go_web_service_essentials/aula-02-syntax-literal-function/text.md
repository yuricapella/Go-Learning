oque entendi:
geralmente vamos declarar uma função e assignar ela a uma variavel para usar depois
defer roda apos tudo da func main rodar
var n int

declare an anonymous function and call it

func() {
    fmt.Println("Direct:", n)
}

now assign it to a variable
f:= func() {
    fmt.Println("Variable:", n)
}

call the anonymous function pela variavel
f()

Defer the call to the anonymous function till after main returns.
defer func(){
 fmt.println(defer1:", n)
}

set value n to 3 before the return
n = 3

f()

defer func(){
    fmt.println("defer 2:", n)
}

ordem:
1-func()0 
2- f()0
3- f() 3
4- defer 2: 3
5- defer 1: 3