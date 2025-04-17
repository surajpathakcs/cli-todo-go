package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	for index , arg := range args{
		fmt.Printf("Arg %d : %s\n",index,arg)
	}

	if len(args) < 2 {
		 fmt.Println("Calculator: [add | sub] ")
		 return
	}

	cmd := args[1]
	
	arg1,err1 := strconv.Atoi(args[2])
	arg2,err2 := strconv.Atoi(args[3])
	if(err1 != nil){
		fmt.Println(err1)
	}
	if err2 != nil{
		fmt.Println(err2)
	}

	switch cmd {
	case "add":
		fmt.Printf("%d",calc(arg1,arg2,"+"))

	case "sub":
		fmt.Printf("%d",calc(arg1,arg2,"-"))

	default:
		fmt.Printf("Try again")
	}
}

func calc(num1 int ,num2 int , op string) int{
	if(op == "+"){
		return num1 + num2 
	}
	if (op == "-"){
		return num1 - num2
	}
	return 0
}
