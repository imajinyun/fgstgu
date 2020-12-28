#include <stdio.h>
#include "hello.h"

int main(int argc, const char* argv[]) {
	int res = sum(10, 20);
	printf("result: %d\n", res);
	return 0;
}
