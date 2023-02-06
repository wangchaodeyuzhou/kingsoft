package add_example

/*
#include <stdio.h>
#include <string.h>
#include <stdlib.h>

float c;

float addSum(float a, float b) {
    c = a + b;
    printf("add Sum in C %.2f\n", c);
    return c;
};

float printCInC() {
return c;
};

float GetNum() {
   printf("get num in C %.2f\n", c);
   return printCInC();
}

*/
import "C"
import "fmt"

func PrintCInGo() {
	c := C.addSum(1.2, 3.4)
	fmt.Println("in go c", c)

	getC := C.GetNum()
	fmt.Println("in go get c", getC)

}
