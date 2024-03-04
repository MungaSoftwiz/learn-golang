# Sycn package

The Sync package in Go is used for synchronization.
It provides primitives sucha as mutexes(`sync.Mutes`) & wait groups(`sync.WaitGroup`).

## Mutex
In Go Mutex(short for mutual exclusion) is a synchronization primitive provided
by the `sync` package.
It's used to control access to the resource at any given time.
This prevents data races and ensures that concurrent access to shared DS is safe.

Mutexes provide two main methods:
1.`Lock()`: This method acquires the lock of the mutex.
It lock is already acquired by another goroutine, calling goroutine will block
until lock is available.
2.`Unlock()`: This method releases the lock. Called after critical section of code
has been executed to allow other goroutines to acquire the lock.
