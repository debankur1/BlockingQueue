package Deque

//All the operations that deals with Memory
//Synchronization are Blocking in nature like Enqueue and Dequeue

import (
	"container/list"
	"sync"
)

type DQueue struct {
	sync.RWMutex
	linkedList list.List
}
//Basic Enqueue operation that a Queue provides with a mechanism of FIFO operation
//These operation will insert element and the end of the Queue and retrieve element from the head of the Queue
func(DataSyncDQueue *DQueue)EnqueueFromFirst(message interface{}){
	DataSyncDQueue.Lock()
	DataSyncDQueue.linkedList.PushFront(message)
	DataSyncDQueue.Unlock()
}
//These operation will add item at the Top of the Queue , before the Top element of the current Top.
//Could be a scenario a VIP comes in and stands at the first position !!!
func(DataSyncDQueue *DQueue)EnqueueFromLast(message interface{}){
	DataSyncDQueue.Lock()
	DataSyncDQueue.linkedList.PushBack(message)
	DataSyncDQueue.Unlock()
}
//Basic Dequeue operation where the data will be popped out form the Queue from the beginning.
func(DataSyncDQueue *DQueue)DequeueFromFirst()interface{}{
	var result interface{}
	DataSyncDQueue.Lock()
	element := DataSyncDQueue.linkedList.Back()
	if element!=nil {
		DataSyncDQueue.linkedList.Remove(element)
		result = element.Value
	}
	DataSyncDQueue.Unlock()
	return result
}
//Dequeue operation from the back side of the Queue.
//May be people get bored standing in a huge Queue and wants to come out !!!
func(DataSyncDQueue *DQueue)DequeueFromLast()interface{}{
	var result interface{}
	DataSyncDQueue.Lock()
	element := DataSyncDQueue.linkedList.Front()
	if element!=nil {
		DataSyncDQueue.linkedList.Remove(element)
		result = element.Value
	}
	DataSyncDQueue.Unlock()
	return result
}
//Operation will provide the top element in the Deque
func(DataSyncDQueue *DQueue)Top()interface{}{
	var result interface{}
	DataSyncDQueue.RLock()
	element := DataSyncDQueue.linkedList.Back()
	if element!=nil {
		result = element.Value
	}
	DataSyncDQueue.RUnlock()
	return result
}
//Operation will provide the last element in the Deque
func(DataSyncDQueue *DQueue)Bottom()interface{}{
	var result interface{}
	DataSyncDQueue.RLock()
	element := DataSyncDQueue.linkedList.Front()
	if element!=nil {
		result = element.Value
	}
	DataSyncDQueue.RUnlock()
	return result
}
//Size of Deque
func (DataSyncDQueue *DQueue)Size()int{
	return DataSyncDQueue.linkedList.Len()
}
func (DataSyncDQueue *DQueue)isEmpty()bool{
	if DataSyncDQueue.linkedList.Len() > 0{
		return true
	}
	return false
}
