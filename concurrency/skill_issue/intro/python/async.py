import asyncio
import random
import time

async def async_random_wait():
    work_time = random.randint(1, 5)
    await asyncio.sleep(work_time)
    return work_time

OPERATIONS_COUNT = 100

async def solution_via_gather():
    execution_time = 0
    total_time = 0
    
    start = time.time()
    
    tasks = [
        asyncio.create_task(async_random_wait())
        for _ in range(OPERATIONS_COUNT)
    ]
    results = await asyncio.gather(*tasks)
    
    total_time = sum(results)
    
    execution_time = time.time() - start
    
    print("execution time:", execution_time)
    print("total time:", total_time)
    
# async def solution_via_asyncio_queue():
#     execution_time = 0
#     total_time = 0
#     q = asyncio.Queue()
    
#     start = time.time()
    
#     async def producer():
#         for _ in range(OPERATIONS_COUNT):
#             work_time = await async_random_wait()
#             await q.put(work_time)
            
#     async def consumer():
#         nonlocal total_time
#         for _ in range(OPERATIONS_COUNT):
#             work_time = await q.get()
#             total_time += work_time
            
#     await asyncio.gather(producer(), consumer())
    
#     execution_time = time.time() - start
    
#     print("execution time:", execution_time)
#     print("total time:", total_time)

async def solution_via_asyncio_mutex():
    execution_time = 0
    total_time = 0
    lock = asyncio.Lock()
    
    start = time.time()
    
    async def worker():
        nonlocal total_time
        work_time = await async_random_wait()

        async with lock:
            total_time += work_time
            
    tasks = [
        asyncio.create_task(worker())
        for _ in range(OPERATIONS_COUNT)
    ]
    
    await asyncio.gather(*tasks)
    
    execution_time = time.time() - start
    
    print("execution time:", execution_time)
    print("total time:", total_time)
    
if __name__ == "__main__":
    print("Solution via gather:")
    asyncio.run(solution_via_gather())
    # print("Solution via queue:")
    # asyncio.run(solution_via_asyncio_queue())
    print("Solution via mutex:")
    asyncio.run(solution_via_asyncio_mutex())