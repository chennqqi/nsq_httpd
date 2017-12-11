package main

import (
	"container/list"
	"sync"
)

type Queue struct {
	list   *list.List
	mutext sync.Mutex
	wg     sync.WaitGroup
}

func NewQueue() *Queue {
	var q Queue
	l := list.New()
	l.Init()
	q.list = l
	return &q
}

func (q *Queue) Len() int {
	q.mutext.Lock()
	defer q.mutext.Unlock()
	l := q.list
	return l.Len()
}

func (q *Queue) PopFront() interface{} {
	q.mutext.Lock()
	defer q.mutext.Unlock()
	l := q.list
	f := l.Front()
	if f != nil {
		q.wg.Done()
		return l.Remove(f)
	}
	return nil
}

func (q *Queue) PopBack() interface{} {
	q.mutext.Lock()
	defer q.mutext.Unlock()
	l := q.list
	f := l.Back()
	if f != nil {
		q.wg.Done()
		return l.Remove(f)
	}
	return nil
}

func (q *Queue) PushFront(element interface{}) {
	q.mutext.Lock()
	defer q.mutext.Unlock()
	l := q.list
	l.PushFront(element)
	q.wg.Add(1)
}

func (q *Queue) PushBack(element interface{}) {
	q.mutext.Lock()
	defer q.mutext.Unlock()
	l := q.list
	l.PushBack(element)
	q.wg.Add(1)
}
