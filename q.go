package ebgo
import (
    "errors"
    //"time"
    "sync"
)

type Queue struct {
    mu sync.Mutex
    capacity int
    q        []string
}

// FifoQueue
type FifoQueue interface {
    Insert()
    Remove()
}

// Insert inserts the item into the queue
func (q *Queue) Insert(item string) error {
    q.mu.Lock()
         defer q.mu.Unlock()
    if len(q.q) < int(q.capacity) {
        q.q = append(q.q, item)
        return nil
    }
    return errors.New("Queue is full")
}

// Remove removes the oldest element from the queue
func (q *Queue) Remove() (string, error) {
    q.mu.Lock()
         defer q.mu.Unlock()
    if len(q.q) > 0 {
        item := q.q[0]
        q.q = q.q[1:]
        return item, nil
    }
    return "0", errors.New("Queue is empty")
}

// CreateQueue creates an empty queue with desired capacity
func CreateQueue(capacity int) *Queue {
    return &Queue{
        capacity: capacity,
        q:        make([]string, 0, capacity),
    }
}
