package blockchain

import (
	"sync"
)

const (
	MinutesPerHour   uint8 = 60
	SecondsPerMinute uint8 = 60
)

type ExecSummary struct {
	timePeriodBlock uint64 //microseconds
	countBlock      uint64
	timePeriodTx    uint64 //microseconds
	countTx         uint64
}

//Struct where recent transaction count is stored
type ExecRecordQueueImpl struct {
	execPerMin         [MinutesPerHour]ExecSummary
	head               uint8
	tail               uint8
	sumTimePeriodBlock uint64 //microseconds
	sumCountBlock      uint64
	sumTimePeriodTx    uint64 //microseconds
	sumCountTx         uint64
	currentExec        *ExecSummary
	hasPushed          bool
	queueLock          sync.RWMutex
}

//Add to current Block execution period and count within 5 minutes
func (q *ExecRecordQueueImpl) addcurrentExecBlock(timePeriod uint64) {
	q.queueLock.Lock()
	q.currentExec.countBlock++
	q.currentExec.timePeriodBlock += timePeriod
	q.queueLock.Unlock()
}

//Add to current Block execution period and count within 5 minutes
func (q *ExecRecordQueueImpl) addcurrentExecTx(timePeriod uint64, count uint64) {
	q.queueLock.Lock()
	q.currentExec.countTx += count
	q.currentExec.timePeriodTx += timePeriod
	q.queueLock.Unlock()
}

//Push and refresh block and transaction execution period and count
func (q *ExecRecordQueueImpl) pushExecPerMinute() {
	q.queueLock.Lock()
	if q.head == q.tail && q.hasPushed {
		q.sumTimePeriodBlock -= q.execPerMin[q.head].timePeriodBlock
		q.sumCountBlock -= q.execPerMin[q.head].countBlock
		q.sumTimePeriodTx -= q.execPerMin[q.head].timePeriodTx
		q.sumCountTx -= q.execPerMin[q.head].countTx
		q.head = getNextPos(q.head)
	}
	q.hasPushed = true
	q.sumTimePeriodBlock += q.currentExec.timePeriodBlock
	q.sumCountBlock += q.currentExec.countBlock
	q.sumTimePeriodTx += q.currentExec.timePeriodTx
	q.sumCountTx += q.currentExec.countTx
	q.execPerMin[q.tail].timePeriodBlock = q.currentExec.timePeriodBlock
	q.execPerMin[q.tail].countBlock = q.currentExec.countBlock
	q.execPerMin[q.tail].timePeriodTx = q.currentExec.timePeriodTx
	q.execPerMin[q.tail].countTx = q.currentExec.countTx
	q.tail = getNextPos(q.tail)
	q.currentExec.timePeriodBlock = 0
	q.currentExec.countBlock = 0
	q.currentExec.timePeriodTx = 0
	q.currentExec.countTx = 0
	q.queueLock.Unlock()
}

//Calculate TPS in recent minute and hour
func (q *ExecRecordQueueImpl) getMetricsRecent() (float64, float64, float64, float64, float64, float64) {
	q.queueLock.RLock()
	length := getLength(q.head, q.tail)
	prevExecResult := q.execPerMin[getPrevPos(q.tail)]
	tpsRecentMinute := float64(prevExecResult.countTx) / float64(SecondsPerMinute)
	tpsRecentHour := float64(q.sumCountTx) / float64(SecondsPerMinute) / float64(length)
	blockPeriodRecent5Min, countBlockRecent5Min, txPeriodRecent5Min, countTxRecent5Min := getRecentNMinSummary(q, q.head, q.tail, 5)
	var avrgBlockPeriodRecent5Min, avrgBlockPeriodRecentHour, avrgTxPeriodRecent5Min, avrgTxPeriodRecentHour float64 = 0, 0, 0, 0
	if countBlockRecent5Min != 0 {
		avrgBlockPeriodRecent5Min = float64(blockPeriodRecent5Min) / float64(countBlockRecent5Min)
	}
	if q.sumCountBlock != 0 {
		avrgBlockPeriodRecentHour = float64(q.sumTimePeriodBlock) / float64(q.sumCountBlock)
	}
	if countTxRecent5Min != 0 {
		avrgTxPeriodRecent5Min = float64(txPeriodRecent5Min) / float64(countTxRecent5Min)
	}
	if q.sumCountTx != 0 {
		avrgTxPeriodRecentHour = float64(q.sumTimePeriodTx) / float64(q.sumCountTx)
	}
	q.queueLock.RUnlock()
	return tpsRecentMinute, tpsRecentHour, avrgBlockPeriodRecent5Min, avrgBlockPeriodRecentHour, avrgTxPeriodRecent5Min, avrgTxPeriodRecentHour
}

func getNextPos(pos uint8) uint8 {
	return (pos + 1) % MinutesPerHour
}

func getPrevPos(pos uint8) uint8 {
	return (pos + MinutesPerHour - 1) % MinutesPerHour
}

func getLength(head uint8, tail uint8) uint8 {
	if tail != head {
		return tail - head
	} else {
		return MinutesPerHour
	}
}

func getRecentNMinSummary(q *ExecRecordQueueImpl, head uint8, tail uint8, n uint8) (uint64, uint64, uint64, uint64) {
	if tail > head && tail-head > n {
		n = tail - head
	}
	var blockPeriodRecentNMin, txPeriodRecentNMin, countBlockRecentNMin, countTxRecentNMin uint64 = 0, 0, 0, 0
	start := getPrevPos(tail)
	for i := uint8(0); i < n; i++ {
		record := q.execPerMin[start]
		blockPeriodRecentNMin += record.timePeriodBlock
		txPeriodRecentNMin += record.timePeriodTx
		countBlockRecentNMin += record.countBlock
		countTxRecentNMin += record.countTx
		start = getPrevPos(start)
	}
	return blockPeriodRecentNMin, countBlockRecentNMin, txPeriodRecentNMin, countTxRecentNMin
}
