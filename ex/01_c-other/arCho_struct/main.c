#include <stdio.h>

typedef struct books_s
{
	char title[64];
	char author[32];
	int price;
} book;

int main() {
	book book1 = {"How Google Works", "Eric Schmidt", 15000};
	book *book_ptr = &book1;

	// struct var access
	printf("(1) Book Title: %s, Author: %s, Price: %d", book1.title, book1.author, book1.price );

	// struct pointer access
	printf("\n(2) Book Title: %s, Author: %s, Price: %d", book_ptr->title, book_ptr->author, book_ptr->price);

	return 0;
}
