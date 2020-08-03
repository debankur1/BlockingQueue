Blocking Queue Implementation in Golang using DoublyLinkedList and Read Write Locks

Operations:

1. Enqueue() --> Insert Data in the Queue.
2. Dequeue() --> Remove first element in the Queue and return the data.
3. Peek()    --> Top element in the Queue.
4. Size()    --> Number of elements in the Queue. 

Operations like Enqueue() and Dequeue() only one Thread will be allowed at a time as the underlying memory is 
getting updated.

For Peek() and Size() multiple thread can access at the same time.