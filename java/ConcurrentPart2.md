# Condition Variable
Avoid busy waiting (spinning)

1. Queue of threads waiting for a certain condition.
2. Associated with a mutex.
3. Protect critical section of code with exclusion.
4. Provide ability for threads to wait until a condition occurs.
5. There are three main steps:
   * wait
   * signal (notify or wake)
   * broadcast
6. In Java -- Condition
```
mutex.lock();
// do something change that condition
conditionVaraiable.signal();
mutex.unlock();
```

# Consumer and Producer

1. Enforce mutual exclusion of producers adn consumers.
2. Prevent producers from trying to add data to full queue.
3. Prevent consumers from trying to remove data from empty queue.
4. Average rate of production < Average rate of consumption. (Or they will lose data)
5. In Java -- ArrayBlockQueue<E> (bounded queue and thread safe queue).

# Semaphore

1. Can be used by multiple threads at the same time.
2. Include a counter to track availability.
3. If counter is positive, every thread can acquire semaphore, then decrement counter.
4. If counter is zero, wait until available.
5. Release semaphore, increment counter.
6. Signal waited thread.

## Binary Semaphore

1. 0 represents locked.
2. 1 represents unlocked.
3. used similar to mutes by acquiring/releasing
4. Mutex v.s. Semaphore:
   * Mutex can only be acquired/released by the same thread.
   * Semaphore can acquired/released by different threads.

# Race condition

data races - two or more threads concurrently access the same memory location.
   * need to ensure mutual exclusion for shared resource.

race condition - is a flaw in the timing or ordering of a program's execution that causes incorrect behavior.

1. race conditions really hard to find.
2. Search for Race Conditions
   * use sleep statements to modify timing and execution order.
   * Also called Heisenbug - a software bug that disappears when you tru to study it.

# Barrier

prevents a group of threads from proceeding until enough threads have reached the barrier.

* CyclicBarrier
* CountDownLatch

# Computational graph

1. Spawn label: can run at any time after the spawn operation.
2. Sync: not be able to execute until the dependency tasks complete.
3. Directed Acyclic Graph(DAG)
4. Critical Path: represent the longer series of sequential operations through the diagram. It could indicate the shortest possible execution time if this program was parallelized as much as possible.
5. Minimizing the Critical Path is important in designing parallel algorithms.

# Thread pool
1. Taking task from queue.
2. The time it takes to execute the task is less than the time required to create a new thread.
3. Executor Service (Java)

# Future
1. Placeholder for a result that will be available later.
2. Mechanism to access the result of an asynchronous operation.
3. It's read-only
4. Resolving or Fulfilling, it means the Future is finished.
5. Callable<V> class

# ForkJoinPool
1. ExecutorService that executes ForkJoinTasks
2. fork() - Asynchronously execute task in ForkJoinPool
3. join() - Return result of computation when it is done
4. RecursiveTask<V> - Return a result
5. RecursiveAction - Does not return a result

# Parallel Reason
## Weak Scaling
1. Variable number of processors with fixed problem size per processor
2. Accomplish more work in the same time
## Strong Scaling
1. Variable number of processors with fixed total problem size
2. Accomplish same work in less time
## Throughput
tasks/per time
## Latency
time/per task
## Speedup
sequential execution time/parallel execution time with N workers
## Amdahl's law
Overall Speedup = 1/((1-P)+(1-P)/2)
P is parallel percent in one program


# Coarse-grained parallelism
high computation-to-communication ratio
# Fine-grained parallelism
good load-balancing


Java's Remote Method Invocation API
1. Enables interprocess communication
2. Invoke methods on objects in another JVM instance