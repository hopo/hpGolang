#include <stdio.h>
#include <stdbool.h>
#include <stdlib.h>

int solution(int price[], int price_len, int money) {
    int answer = 0, total = 0;
    for(int i = 0; i < price_len; ++i)
        total += price[i];
    answer = (money < total)? -1 : money - total;
    return answer;
}

int main() {
	int price[] = {2100, 3200, 2100, 800};
	int price_len = 4;
	int money = 10000;
	int ex = solution(price, price_len, money);	// 1800
	printf("%d", ex);
}
