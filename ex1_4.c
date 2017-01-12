
#include <pthread.h>
#include <stdio.h>

static int i = 0;

void* someThreadFunction1(){
	for(int j = 0; j < 1000000; j++){
		i++;
	}
	return NULL;
}

void* someThreadFunction2(){
	for(int k = 0; k < 1000000; k++){
		i--;
	}
	return NULL;
}

int main(void){
	pthread_t someThread1;
	pthread_create(&someThread1, NULL, someThreadFunction1, NULL);
	pthread_t someThread2;
	pthread_create(&someThread2, NULL, someThreadFunction2, NULL);
	pthread_join(someThread1, NULL);
	pthread_join(someThread2, NULL);
	printf("Dette er tallet: %d\n", i);
	return 0;
}

