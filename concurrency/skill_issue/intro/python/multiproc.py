import multiprocessing as mp
import random
import time


def random_wait():
    work_time = random.randint(1, 5)
    time.sleep(work_time)
    return work_time

OPERATIONS_COUNT = 100

def solution_via_queue():
    execution_time = 0
    total_time = 0
    q = mp.Queue()
    
    start = time.time()
    
    ctx = mp.get_context('fork')
    
    processes = [
        ctx.Process(target=lambda q: q.put(random_wait()), args=(q,))
        for _ in range(OPERATIONS_COUNT)
    ]
    
    for process in processes:
        process.start()
    
    for process in processes:
        process.join()
    
    while not q.empty():
        total_time += q.get()
        
    execution_time = time.time() - start
    
    print("execution time:", execution_time)
    print("total time:", total_time)
    
def solution_via_shared_memory():
    execution_time = 0
    total_time = mp.Value('i', 0, lock=True)
    
    start = time.time()
    
    def worker(shared_total):
        work_time = random_wait()
        with shared_total.get_lock():
            shared_total.value += work_time
    
    ctx = mp.get_context('fork')
    
    processes = [
        ctx.Process(target=worker, args=(total_time,))
        for _ in range(OPERATIONS_COUNT)
    ]
    
    for process in processes:
        process.start()
    
    for process in processes:
        process.join()
        
    execution_time = time.time() - start
    
    print("execution time:", execution_time)
    print("total time:", total_time.value)
    
    
def solution_via_mutex():
    execution_time = 0
    total_time = mp.Value('i', 0, lock=True)
    lock = mp.Lock()
    
    start = time.time()
    
    def worker(shared_lock):
        nonlocal total_time        
        work_time = random_wait()
        shared_lock.acquire()
        total_time.value += work_time
        shared_lock.release()
        
    ctx = mp.get_context('fork')
    
    processes = [
        ctx.Process(target=worker, args=(lock,))
        for _ in range(OPERATIONS_COUNT)
    ]
    
    for process in processes:
        process.start()
    
    for process in processes:
        process.join()
        
    execution_time = time.time() - start
    
    print("execution time:", execution_time)
    print("total time:", total_time.value)

if __name__ == "__main__":
    print("Solution via queue:")
    solution_via_queue()
    print("Solution via shared memory:")
    solution_via_shared_memory()
    print("Solution via mutex:")
    solution_via_mutex()    
