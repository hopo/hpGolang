#include <stdio.h>
#include <string.h>

void append(char* s, char c) {
        int len = strlen(s);
        s[len] = c;
        s[len+1] = '\0';
}

int main(void) {
        char str[256] = "hello";
        char c = 'o';

        append(str, c);

        printf("%s\n", str);
        return 0;
}
