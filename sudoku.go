package main

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"math/rand"
	"os"
	"sort"
)

var initdata = []uint32{
	0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0,
}

type OpNums struct {
	index    int
	backUp   []uint32
	nowIndex int
}

type Sudoku struct {
	fullData []uint32
}

func NewSudoku(initdata []uint32) *Sudoku {
	if len(initdata) != 81 {
		return nil
	}
	sudoku := &Sudoku{}
	sudoku.Init(initdata)
	return sudoku
}

func (this *Sudoku) Init(initdata []uint32) {
	this.fullData = initdata
}

func (this *Sudoku) Print() {
	data := make([][]string, 0, 9)
	temp := make([]string, 0, 9)
	for index, val := range this.fullData {
		numstr := " "
		if val != 0 {
			numstr = fmt.Sprintf("%d", val)
		}
		temp = append(temp, numstr)
		if (index+1)%9 == 0 {
			data = append(data, temp)
			temp = make([]string, 0, 9)
		}
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetRowLine(true)

	for _, v := range data {
		table.Append(v)
	}
	table.Render()
}

func (this *Sudoku) GetLineIndexs(index int) []int {
	indexs := make([]int, 0)
	line := index / 9
	for i := line * 9; i < line*9+9; i++ {
		indexs = append(indexs, i)
	}
	return indexs
}

func (this *Sudoku) GetCoumIndexs(index int) []int {
	indexs := make([]int, 0)
	coum := index % 9
	for j := coum; j < 81; j += 9 {
		indexs = append(indexs, j)
	}
	return indexs
}

func (this *Sudoku) GetLittleSudokuIndexs(index int) []int {
	indexs := make([]int, 0)
	line := index / 9
	coum := index % 9

	base := line/3*3*9 + coum/3*3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			tempIndex := base + i + j*9
			indexs = append(indexs, tempIndex)
		}
	}
	return indexs
}

// 相关的格子(排除index本身)
func (this *Sudoku) GetRelatedIndexs(index int) map[int]bool {
	indexs := make(map[int]bool, 0)
	for _, in := range this.GetLineIndexs(index) {
		if in != index {
			indexs[in] = true
		}
	}
	for _, in := range this.GetCoumIndexs(index) {
		if in != index {
			indexs[in] = true
		}
	}
	for _, in := range this.GetLittleSudokuIndexs(index) {
		if in != index {
			indexs[in] = true
		}
	}
	return indexs
}

// 能填的数字
func (this *Sudoku) CanFillNums(index int) []uint32 {
	nums := make([]uint32, 0)
	if this.fullData[index] != 0 {
		return nums
	}
	backUp := map[uint32]bool{1: true, 2: true, 3: true, 4: true, 5: true, 6: true, 7: true, 8: true, 9: true}

	relatedIndexs := this.GetRelatedIndexs(index)
	for in, _ := range relatedIndexs {
		if this.fullData[in] != 0 {
			delete(backUp, this.fullData[in])
		}
	}

	//fmt.Println(backUp)
	for num, _ := range backUp {
		nums = append(nums, num)
	}
	return nums
}

// 是否填满
func (this *Sudoku) CheckOver() bool {
	for _, val := range this.fullData {
		if val == 0 {
			return false
		}
	}
	checkData := make(map[uint32]bool)
	// 每行
	for i := 0; i < 73; i += 9 {
		for _, in := range this.GetLineIndexs(i) {
			checkData[this.fullData[in]] = true
		}
		if len(checkData) != 9 {
			return false
		}
	}
	checkData = make(map[uint32]bool)
	// 每列
	for i := 0; i < 9; i++ {
		for _, in := range this.GetCoumIndexs(i) {
			checkData[this.fullData[in]] = true
		}
		if len(checkData) != 9 {
			return false
		}
	}
	checkData = make(map[uint32]bool)
	// 小宫格
	little := []int{0, 27, 54}
	for _, in := range little {
		for _, i := range this.GetLittleSudokuIndexs(in) {
			checkData[this.fullData[i]] = true
		}
		if len(checkData) != 9 {
			return false
		}

		checkData = make(map[uint32]bool)
		for _, j := range this.GetLittleSudokuIndexs(in + 3) {
			checkData[this.fullData[j]] = true
		}
		if len(checkData) != 9 {
			return false
		}

		checkData = make(map[uint32]bool)
		for _, k := range this.GetLittleSudokuIndexs(in + 6) {
			checkData[this.fullData[k]] = true
		}
		if len(checkData) != 9 {
			return false
		}
	}
	return true
}

func (this *Sudoku) CanFillNum(num uint32, index int) bool {
	for _, in := range this.GetLineIndexs(index) {
		if in != index && this.fullData[in] == num {
			return false
		}
	}

	for _, in := range this.GetCoumIndexs(index) {
		if in != index && this.fullData[in] == num {
			return false
		}
	}

	for _, in := range this.GetLittleSudokuIndexs(index) {
		if in != index && this.fullData[in] == num {
			return false
		}
	}
	return true
}

func (this *Sudoku) Compute() bool {
	sortData := make([]*OpNums, 0)
	for index, val := range this.fullData {
		if val == 0 {
			nums := this.CanFillNums(index)
			if len(nums) <= 0 {
				fmt.Println("出错", index)
				return false
			}
			opNums := &OpNums{
				index:    index,
				backUp:   nums,
				nowIndex: -1,
			}
			sortData = append(sortData, opNums)
		}
	}
	sort.SliceStable(sortData, func(i, j int) bool {
		if len(sortData[i].backUp) < len(sortData[j].backUp) {
			return true
		}
		return false
	})
	var tarindex int
	var con bool
	for !this.CheckOver() {
		if !con {
			if tarindex < 0 {
				return false
			}
			con, tarindex = this.computeBase(sortData, tarindex)
		}
	}
	return true
}

func (this *Sudoku) computeBase(sortData []*OpNums, sortIndex int) (bool, int) {
	for i := sortIndex; i < len(sortData); i++ {
		this.fullData[sortData[i].index] = 0
		if i != sortIndex {
			sortData[i].nowIndex = -1
		}
	}

	for i := sortIndex; i < len(sortData); i++ {
		if len(sortData[i].backUp) == 1 {
			this.fullData[sortData[i].index] = sortData[i].backUp[0]
			sortData[i].nowIndex = 0
		} else if len(sortData[i].backUp) > 1 {
			sortData[i].nowIndex++
			if sortData[i].nowIndex >= len(sortData[i].backUp) {
				return false, i - 1
			}

			if !this.CanFillNum(sortData[i].backUp[sortData[i].nowIndex], sortData[i].index) {
				return false, i
			} else {
				this.fullData[sortData[i].index] = sortData[i].backUp[sortData[i].nowIndex]
			}
		}
	}

	return true, 0
}

// 生成数独
func GenerateSudoku() {
	head := []uint32{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < 9; i++ {
		index1 := rand.Int31n(9)
		index2 := rand.Int31n(9)
		head[index1], head[index2] = head[index2], head[index1]
	}
	for i := 0; i < 9; i++ {
		initdata[i] = head[i]
	}

	sudoku := NewSudoku(initdata)
	if sudoku.Compute() {
		count := uint32(0)
		for count < 50 {
			in := rand.Int31n(81)
			if sudoku.fullData[in] != 0 {
				sudoku.fullData[in] = 0
				count++
			}
		}
		sudoku.Print()
		fmt.Println("生成成功")
	} else {
		fmt.Println("生成失败")
	}
}

// 解决数独
func SettleSudoku(initdata []uint32) bool {
	sudoku := NewSudoku(initdata)
	if sudoku.Compute() {
		sudoku.Print()
		fmt.Println("解决成功")
	} else {
		fmt.Println("解决失败")
		return false
	}
	return true
}
