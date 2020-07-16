package dataStruct

import "fmt"

type Heap struct {
	list []int
}

func (h *Heap) Push(v int) {
	h.list = append(h.list, v)

	idx := len(h.list) - 1
	for idx >= 0 {
		parentIdx := (idx - 1) / 2
		if parentIdx < 0 {
			break
		}
		if h.list[idx] > h.list[parentIdx] {
			h.list[idx], h.list[parentIdx] = h.list[parentIdx], h.list[idx]
			idx = parentIdx
		} else {
			break
		}
	}
}

func (h *Heap) Print() {
	fmt.Println(h.list)
}

func (h *Heap) Pop() int {
	if len(h.list) == 0 {
		return 0
	}

	top := h.list[0]
	last := h.list[len(h.list)-1]
	h.list = h.list[:len(h.list)-1]

	h.list[0] = last
	idx := 0
	for idx < len(h.list) {
		swapIdx := -1
		leftIdx := idx*2 + 1
		if leftIdx >= len(h.list) { //왼쪽 자식이 없을떄, 왼쪽이없다면 오른쪽도 당연히 없다
			break
		}
		if h.list[leftIdx] > h.list[idx] { //자식값이 크다면 인데스 교체
			swapIdx = leftIdx
		}

		rightIdx := idx*2 + 2
		if rightIdx < len(h.list) { //오른쪽 자식있는 경우
			if h.list[rightIdx] > h.list[idx] { //자식노드가 더클때 swapIdx를 이용 오른쪽,왼쪽중 큰 쪽구분
				if swapIdx < 0 || h.list[swapIdx] < h.list[rightIdx] {
					swapIdx = rightIdx
				}
			}
		}
		//스왑할 애가 없을때: 자식보다 자기가 더클때
		if swapIdx < 0 {
			break
		}

		h.list[idx], h.list[swapIdx] = h.list[swapIdx], h.list[idx]
		idx = swapIdx

	}
	return top

}
