package main 

import (
"fmt")

func main(){
	fmt.Println("Hello world");
	testFor();
	testLoopCount(15);
	lessThanFive();
	checkProbability(10);
	checkProbability(50);

	var sum float64 = 0.0;
	for i:=0; i<10; i++{
		sum += checkProbability(i);
	}

	fmt.Println("The total probability is", sum);
}

func testFor(){
	i := 0;

	for i < 100{
		i++;
		fmt.Println("This is test:", i );
	}
}

func testLoopCount(n int){
	for i := 0; i < n; i++{
		j := n;
		for j > 1{
			fmt.Println("(", i,j, ")");
			j = j/2;
		}
		fmt.Println();
	}
}

func sqrt(number int) int{
	return 0;
}

func lessThanFive(){
	
	var counter int = 0;

	for i:=1; i <=6; i++{
		for k:=1; k <=6; k++{
			if i != k && i+k < 5{
				counter++;
			}
		}
	} 

	fmt.Println("Counter number is ", counter);
}

func checkProbability(x int) (y float64){
	var number float64 = (48.0/52.0);

	for i:=1; i <= x;i++{
		var value float64 = float64(i);
		number *= ((48.0-value)/52.0);
		//fmt.Println("Using count", x, "This number is ", number);
	} 


	fmt.Println("Using count", x, "This number is ", number);
	return number;
}