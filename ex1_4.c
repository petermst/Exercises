
#include <pthread.h>
#include <stdio.h>
#include<unistd.h>

static int i = 0;
pthread_mutex_t lock;

void* someThreadFunction1(){

	for(int j = 0; j < 999999; j++){
		pthread_mutex_lock(&lock);
		i++;
		pthread_mutex_unlock(&lock);
	}
	
	return NULL;
}

void* someThreadFunction2(){
	

	for(int k = 0; k < 1000000; k++){
		pthread_mutex_lock(&lock);
		i--;
		pthread_mutex_unlock(&lock);
	}
	
	return NULL;
	
	
}

int main(void){

	if (pthread_mutex_init(&lock, NULL) != 0){
        	printf("\n mutex init failed\n");
        	return 1;
    	}	
	
	pthread_t someThread1;
	pthread_create(&someThread1, NULL, someThreadFunction1, NULL);
	pthread_t someThread2;
	pthread_create(&someThread2, NULL, someThreadFunction2, NULL);
	pthread_join(someThread1, NULL);
	pthread_join(someThread2, NULL);
	pthread_mutex_destroy(&lock);
	printf("Dette er tallet: %d\n", i);

	return 0;
}

