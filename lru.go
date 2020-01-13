package main

import "fmt"

type ListNode struct {
	Val  int
	Key  int
	Pre  *ListNode
	Next *ListNode
}

type LRUCache struct {
	MapKeyValueAddr map[int]*ListNode
	Lens            int
	NowLens         int
	Head            *ListNode
	Last            *ListNode
}

func Constructor(capacity int) LRUCache {
	r := make(map[int]*ListNode, capacity)
	lru := LRUCache{
		MapKeyValueAddr: r,
		Lens:            capacity,
		NowLens:         0,
		Head:            nil,
		Last:            nil,
	}
	return lru
}

func (this *LRUCache) Get(key int) int {
	if this.NowLens == 0 {
		return -1
	}

	res, ok := this.MapKeyValueAddr[key]
	if !ok {
		return -1
	}

	if this.Head == res {
		return res.Val
	}

	newHead := &ListNode{
		Val: res.Val,
		Key: key,
		Pre: nil,
	}
	// 分析和移动

	this.move(res, newHead, key)
	return res.Val
}

func (this *LRUCache) move(res *ListNode, newHead *ListNode, key int) {

	if this.Last == res {
		oldHead := this.Head
		oldHead.Pre = newHead
		newHead.Next = oldHead
		res.Pre.Next = nil
		this.MapKeyValueAddr[res.Pre.Key] = res.Pre
		this.Last = res.Pre
		this.MapKeyValueAddr[key] = newHead
		this.Head = newHead
		return
	}

	oldHead := this.Head
	oldHead.Pre = newHead
	newHead.Next = oldHead

	this.MapKeyValueAddr[key].Pre.Next = this.MapKeyValueAddr[key].Next

	this.MapKeyValueAddr[key].Next.Pre = this.MapKeyValueAddr[key].Pre
	this.MapKeyValueAddr[this.MapKeyValueAddr[key].Next.Key] = this.MapKeyValueAddr[key].Next
	this.MapKeyValueAddr[oldHead.Key] = oldHead
	this.MapKeyValueAddr[this.MapKeyValueAddr[key].Pre.Key] = this.MapKeyValueAddr[key].Pre
	this.MapKeyValueAddr[key] = newHead
	this.Head = newHead
	return

}

func (this *LRUCache) Put(key int, value int) {

	if this.Lens <= 0 {
		return
	}

	if this.NowLens == 0 {
		newHead := &ListNode{
			Val:  value,
			Key:  key,
			Pre:  nil,
			Next: nil,
		}
		this.Head = newHead
		this.Last = newHead
		this.MapKeyValueAddr[key] = newHead
		this.NowLens++
		return
	}

	if res, ok := this.MapKeyValueAddr[key]; ok {
		if this.Head == res {
			res.Val = value
			this.MapKeyValueAddr[key] = res
			return
		}

		newHead := &ListNode{
			Val:  value,
			Key:  key,
			Pre:  nil,
			Next: nil,
		}
		this.move(res, newHead, key)
		return
	}

	this.NowLens++
	newHead := &ListNode{
		Val: value,
		Key: key,
		Pre: nil,
	}

	oldHead := this.Head
	oldHead.Pre = newHead
	newHead.Next = oldHead
	this.Head = newHead
	this.MapKeyValueAddr[oldHead.Key] = oldHead
	this.MapKeyValueAddr[key] = newHead
	if this.NowLens > this.Lens {
		this.Last.Pre.Next = nil
		delete(this.MapKeyValueAddr, this.Last.Key)
		this.Last = this.Last.Pre
		this.NowLens--
	}
	return
}

func main() {

	cache := Constructor(2)

	cache.Put(1, 1)
	cache.Put(2, 2)

	a := cache.Get(1)
	fmt.Println(a)
	/*cache.put(2, 2)
	cache.get(1)      // 返回  1
	cache.put(3, 3)    // 该操作会使得密钥 2 作废
	cache.get(2)      // 返回 -1 (未找到)
	cache.put(4, 4)    // 该操作会使得密钥 1 作废
	cache.get(1)    // 返回 -1 (未找到)
	cache.get(3)      // 返回  3
	cache.get(4)       // 返回  4*/
}
