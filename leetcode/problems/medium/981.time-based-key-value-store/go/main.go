// https://leetcode.com/problems/time-based-key-value-store/?envType=problem-list-v2&envId=rab78cw1
// @@@@ 작은값 찾을때 조금 헤맸음. 
type TimeMap struct {
	store map[string][]record
}

type record struct {
	timestamp int
	value       string
}

func Constructor() TimeMap {
	return TimeMap{
		store: make(map[string][]record),
	}
}

func (t *TimeMap) Set(key string, value string, timestamp int) {
	newRecord := record{
		value:       value,
		timestamp: timestamp,
	}
	t.store[key] = append(t.store[key], newRecord)
}

func (t *TimeMap) Get(key string, timestamp int) string {
    store := t.store[key]
    
    if len(store) == 0 {
        return ""
    }

    left, right := 0, len(store)-1
    
    resultIdx := -1

    for left <= right {
        mid := (left + right) / 2
        
        if store[mid].timestamp <= timestamp {
            resultIdx = mid 
            left = mid + 1 
        } else {
            right = mid - 1
        }
    }

    if resultIdx == -1 {
        return ""
    }
    return store[resultIdx].value
}
