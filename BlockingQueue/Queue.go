package BlockingQueue

import (
	"container/list"
	"sync"
)
//Our Queue will backed by a Linked List following Operations : Enqueue(),Dequeue(),Peek()
type SyncQueue struct {
	sync.RWMutex
    linkedList list.List
}
//Allowing Only one thread at a time to add values to the queue, hence the normal blocking Lock is used.
func(DataSyncQueue *SyncQueue)Enqueue(message interface{}){
	DataSyncQueue.Lock()
	DataSyncQueue.linkedList.PushFront(message)
	DataSyncQueue.Unlock()
}
//Allowing Only one thread at a time to delete from the queue, hence the normal blocking Lock is used.
func(DataSyncQueue *SyncQueue)Dequeue()interface{}{
	var result interface{}
	DataSyncQueue.Lock()
	element := DataSyncQueue.linkedList.Back()
	if element!=nil {
		DataSyncQueue.linkedList.Remove(element)
		result = element.Value
	}
	DataSyncQueue.Unlock()
	return result
}
//Allowing multiple threads to read at the same time , hence the ReadLock is used.
func(DataSyncQueue *SyncQueue)Peek()interface{}{
	var result interface{}
	DataSyncQueue.RLock()
	element := DataSyncQueue.linkedList.Back()
	if element!=nil {
		result = element.Value
	}
	DataSyncQueue.RUnlock()
	return result
}
//Size of the LinkedList
func (DataSyncQueue *SyncQueue)Size()int{
	return DataSyncQueue.linkedList.Len()
}
//Boolean method that will return true when Deque has no items
func (DataSyncQueue *SyncQueue)isEmpty()bool{
	if DataSyncQueue.linkedList.Len() < 1{
		return true
	}
	return false
}