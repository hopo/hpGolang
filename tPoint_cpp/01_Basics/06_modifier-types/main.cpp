#include <iostream>
using namespace std;

/* This program show the difference between
    * signed and unsigned intefers.
*/
int main() {
    short int i; // a sifned short integer
    short unsigned int j; // an unsigned short integer

    j = 50000;

    i = j;
    cout << i << " " << j;

    return 0;
}
