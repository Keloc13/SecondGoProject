package main 

import "fmt"

func main(){

	value := [6]int{1,5,4,4,2,5};
	val := [6]int{1,5,4,4,2,5};
	val2 := [10]int{1,5,4,4,2,5,7,8,3,1};
	InsertionSortA(value);
	InsertionSortB(val);
	fmt.Println();
	InsertionSortC(val2);
}

func InsertionSortA(a [6]int){
	//var size int = 3;
	
	for j:=2; j < 6; j++{
		i:=1
		for a[j] > a[i]{
			i++;
		}

		m:= a[j];

		for k:=0; k <= j-i-1; k++{
			a[j-k] = a[j-k-1];
		}
		a[i] = m;
	}

	fmt.Println(a);
}

func InsertionSortB(a [6]int){
	for r:=2; r <6; r++{
		
		j:=5-r+1;
		i:=5;

		for a[j] < a[i]{
			i--;
		}

		m := a[j];
		for k:= 0; k <= i-j-1;k++{
			a[j+k] = a[j+k+1];
		}
		a[i] = m;
	}
	fmt.Println(a);
}

func InsertionSortC(a[10] int){
	
	for j:=2; j < 10; j++{	
		i:=1;
		b:=j;

		for i < b{
			c:= (i+b)/2;
			if a[j] > a[c]{
				i = c + 1;
			}else{
				b=c;
			}
		}
		m := a[j];
		for k:=0; k <= j-i-1; k++{
			a[j-k] = a[j-k-1];
		}
		a[i] = m;
		fmt.Println(j-2,a);
	}
	
}