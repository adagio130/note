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





Java's Remote Method Invocation API
1. Enables interprocess communication
2. Invoke methods on objects in another JVM instance