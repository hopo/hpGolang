#include <stdio.h>

int a = 0x12345678;
short b = 1;
char c = 'A';
float d = 3.14;
long long e;

int main() {
	printf("%p", &a);
	printf("\n%p", &b);
	printf("\n%p", &c);
	printf("\n%p", &d);
	printf("\n%p", &e);
}
