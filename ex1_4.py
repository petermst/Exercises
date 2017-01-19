from threading import Thread
from threading import Lock

i = 0
lock = Lock()

def someThreadFunction1():
	global lock
	global i
	for x in range(999999):
		lock.acquire()
		i += 1
		lock.release()

def someThreadFunction2():
	global lock
	global i
	for y in range(1000000):
		lock.acquire()
		i-=1
		lock.release()

def main():
	
	someThread1 = Thread(target = someThreadFunction1, args = (),)
	someThread1.start()
	someThread2 = Thread(target = someThreadFunction2, args = (),)
	someThread2.start()
	someThread1.join()
	someThread2.join()
	print "Tallet er: ", i

main()
