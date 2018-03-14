#include <iostream>
using namespace std;

int main () {
    // an array with 5 elements.
    double balance[5] = {1000.0, 2.0, 3.4, 17.0, 50.0};
    double *p;

    p = balance;

    // output each array element's value
    cout << "Array values using using pointer " << endl;

    for ( int i = 0; i < 5; i++) {
        cout << "*(p + " << i << ") : ";
        cout << *(balance + i) << endl;
    }

    return 0;
}
