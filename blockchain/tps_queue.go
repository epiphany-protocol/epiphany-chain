package blockchain

import "sync"

const (
	MinutesPerHour   uint8 = 60
	SecondsPerMinute uint8 = 60
)

//Struct where recent transaction count is stored
type TPSQueueImpl struct {
	recentTxCountPerMinute [MinutesPerHour]uint64
	head                   uint8
	tail                   uint8
	sumTxCount             uint64
	queueLock              sync.RWMutex
	currentTxn             uint64
}

//Add to current transaction count within a minute
func (q TPSQueueImpl) addcurrentTxn(count uint64) {
	q.queueLock.Lock()
	q.currentTxn += count
	q.queueLock.Unlock()
}

//Push and refresh transaction count within a minute
func (q TPSQueueImpl) pushTxCountPerMinute() {
	q.queueLock.Lock()
	q.tail = getNextPos(q.tail)
	if q.head == q.tail {
		q.sumTxCount -= q.recentTxCountPerMinute[q.head]
		q.head = getNextPos(q.head)
	}
	q.sumTxCount += q.currentTxn
	q.recentTxCountPerMinute[q.tail] = q.currentTxn
	q.currentTxn = 0
	q.queueLock.Unlock()
}

func getNextPos(pos uint8) uint8 {
	return (pos + 1) % MinutesPerHour
}

//Calculate TPS in recent minute and hour
func (q TPSQueueImpl) getTPSRecent() (float64, float64) {
	q.queueLock.RLock()
	tpsRecentMinute := float64(q.recentTxCountPerMinute[q.tail]) / float64(SecondsPerMinute)
	tpsRecentHour := float64(q.sumTxCount) / float64(SecondsPerMinute) / float64(MinutesPerHour)
	q.queueLock.RUnlock()
	return tpsRecentMinute, tpsRecentHour
}
