#include <typeinfo>
#include <iostream>

class someClass { };

int main(int argc, char* argv[]) {
    int a;
    someClass b;
    std::cout<<"a is of type: "<<typeid(a).name()<<std::endl; // Output 'a is of type int'
    std::cout<<"b is of type: "<<typeid(b).name()<<std::endl; // Output 'b is of type someClass'
    return 0;
}
