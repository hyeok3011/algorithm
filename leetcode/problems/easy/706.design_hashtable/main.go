package main

/*
https://leetcode.com/problems/design-hashmap/
해쉬 테이블 구현 문제이다.
기본적으로 레코드 배열을 사용하였고 collision을 대비해서 체이닝 작업을하였다.
근데 문제에 loadfactor나 table doubling에 대한 언급은 없어서 대충 간략하게만 구현해봤는데
관련된 해쉬테이블 문제가 따로 있으면 그 문제를 풀어보고 아니면 이 문제를 다시 풀어보는것도 좋을듯 하다.
Runtime 96 ms Beats 60.43% Memory 7.6 MB Beats 95.65%
*/

type Record struct {
	value int
	key   int
	next  *Record
}
type MyHashMap struct {
	records []*Record
	size    int
}

func Constructor() MyHashMap {
	return MyHashMap{
		records: make([]*Record, 10000),
		size:    10000,
	}
}

func (this *MyHashMap) Put(key int, value int) {
	hashKey := this.hash(key)

	if this.records[hashKey] == nil {
		this.records[hashKey] = &Record{
			value: value,
			key:   key,
		}
		return
	}

	var targetRecord, preRecord *Record
	targetRecord = this.records[hashKey]

	for targetRecord != nil {
		if targetRecord.key == key {
			targetRecord.value = value
			return
		}

		preRecord = targetRecord
		targetRecord = targetRecord.next
	}

	preRecord.next = &Record{
		key:   key,
		value: value,
	}
}

func (this *MyHashMap) hash(key int) int {
	return key % this.size
}

func (this *MyHashMap) Get(key int) int {
	record := this.records[this.hash(key)]

	for record != nil && record.key != key {
		record = record.next
	}

	if record == nil {
		return -1
	}

	return record.value
}

func (this *MyHashMap) Remove(key int) {
	hashKey := this.hash(key)

	if this.records[hashKey] == nil {
		return
	}

	targetRecord := this.records[hashKey]
	if targetRecord.key == key {
		this.records[hashKey] = targetRecord.next
		return
	}

	preRecord := targetRecord
	targetRecord = targetRecord.next

	for targetRecord != nil {
		if targetRecord.key == key {
			preRecord.next = targetRecord.next
			break
		}

		preRecord = targetRecord
		targetRecord = targetRecord.next
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
