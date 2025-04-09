import threading
import queue
import time
import random


def random_wait():
    work_time = random.randint(1, 5)
    time.sleep(work_time)
    return work_time

OPERATIONS_COUNT = 100

def solution_via_queue():
    execution_time = 0
    total_time = 0
    threads = []
    q = queue.Queue()
    
    start = time.time()
    threads = [
        threading.Thread(target=lambda: q.put(random_wait()))
        for _ in range(OPERATIONS_COUNT)
    ]
    
    for thread in threads:
        thread.start()
    
    for thread in threads:
        thread.join()
        
    while not q.empty():
        total_time += q.get()
        
    execution_time = time.time() - start
    
    print("execution time:", execution_time)
    print("total time:", total_time)
    
def solution_via_buffered_queue():
    execution_time = 0
    total_time = 0
    q = queue.Queue(maxsize=OPERATIONS_COUNT)
    
    start = time.time()
    
    for _ in range(OPERATIONS_COUNT):
        threading.Thread(target=lambda: q.put(random_wait())).start()
        
    for _ in range(OPERATIONS_COUNT):
        total_time += q.get()
        
    execution_time = time.time() - start
    
    print("execution time:", execution_time)
    print("total time:", total_time)
    
    
def solution_via_mutex():
    execution_time = 0
    total_time = 0
    lock = threading.Lock()
    
    start = time.time()
    
    def worker():
        nonlocal total_time
        work_time = random_wait()
        
        lock.acquire()
        total_time += work_time
        lock.release()
        
    threads = [
        threading.Thread(target=worker)
        for _ in range(OPERATIONS_COUNT)
    ]
    
    for thread in threads:
        thread.start()
    
    for thread in threads:
        thread.join()
        
    execution_time = time.time() - start
    
    print("execution time:", execution_time)
    print("total time:", total_time)
    
if __name__ == "__main__":
    print("Solution via queue:")
    solution_via_queue()
    print("Solution via buffered queue:")
    solution_via_buffered_queue()
    print("Solution via mutex:")
    solution_via_mutex()