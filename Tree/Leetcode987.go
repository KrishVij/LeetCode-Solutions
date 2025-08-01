package main

import(

	"sort"
)
type BnNode struct {

	Value int
	Left *BnNode
	Right *BnNode
}

type BnTree struct {

	Root *BnNode
}

type Level struct {

	X,Y int
}

func (bn *BnTree) Bfs() {

	if bn.Root == nil {

		return
	}

	minY := 0
	maxY := 0

	mpp := make(map[*BnNode]Level)
	queue := []*BnNode{}
	queue = append(queue,bn.Root)

	for len(queue) > 0 {

		Current := queue[0]

		if len(mpp) == 0 {

			mpp[Current] = Level{X: 0, Y: 0}
			minY = mpp[Current].Y
			maxY = mpp[Current].Y
		}
		queue = queue[1: ]

		
		if Current.Left != nil {

			queue = append(queue,Current.Left)
			mpp[Current.Left] = Level{X: mpp[Current].X + 1, Y: mpp[Current].Y - 1}

			if mpp[Current.Left].Y < minY {

				minY = mpp[Current.Left].Y
			}
			
			if mpp[Current.Left].Y > maxY {

				maxY = mpp[Current.Left].Y
			}
			
		}

		if Current.Right != nil {

			queue = append(queue,Current.Right)
			mpp[Current.Right] = Level{X: mpp[Current].X + 1, Y: mpp[Current].Y - 1}
			
			if mpp[Current.Right].Y < minY {

				minY = mpp[Current.Right].Y
			}

			if mpp[Current.Right].Y > maxY {

				maxY = mpp[Current.Right].Y
			}
		}
	}

	nodes := []*BnNode{}
	for node,_ := range mpp {

		nodes = append(nodes, node)
	}

	sort.Slice(nodes, (i, j int) bool {

		a,b :=  mpp[nodes[i]], mpp[nodes[j]]

		if a.Y != b.Y {

			return a.Y < b.Y
		}

		if a.X != b.X {

			return a.X < b.X
		}

		return nodes[i].Value < nodes[j].Value
	})

	columnMap := make(map[int][]int)
	for _, node := range nodes {
		y := mpp[node].Y
		columnMap[y] = append(columnMap[y], node.Value)
	}
	
	result := [][]int{}
	for y := minY; y <= maxY; y++ {
		if col, ok := columnMap[y]; ok {
			result = append(result, col)
		}
	}
}

// func (bn *BnTree) DsTreeverticalTraversal() [][]int {

// 	mpp := make(map[*BnNode]level)
// }