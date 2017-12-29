#include <stdio.h>

main() {
	FILE *fp;

	fp = fopen("/Users/HPMBA/workspace/test.txt", "w+");
	
	fprintf(fp, "This is testing for fprintf...\n");
	fputs("This is testting for fputs...\n", fp);
	fclose(fp);
}
