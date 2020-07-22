#include <stdio.h>
#include <string.h>
#include <unistd.h>

int main(){
	char *p = NULL;
	const int MB = 1024 * 100;
	while(1){
		p = malloc(100 * MB);
		memset( p, 0, 100 * MB);
		sleep(1);
	}
	return 0;
}
