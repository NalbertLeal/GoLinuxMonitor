package memInfo

import (
	"errors"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"time"
)

const memFile = "/proc/meminfo"

type RamMem struct {
	totalMem     int
	usedMem      int
	freeMem      int
	availableMem int
	buffers      int
	cached       int
	swapedCache  int
}

func NewRamMem() (*RamMem, error) {
	MemTotalPattern, err := regexp.Compile("MemTotal:        [0-9]+ kB")
	UsedMemPattern, err := regexp.Compile("MemTotal:        [0-9]+ kB")
	MemFreePattern, err := regexp.Compile("MemFree:          [0-9]+ kB")
	MemAvailablePattern, err := regexp.Compile("MemAvailable:    [0-9]+ kB")
	BuffersPattern, err := regexp.Compile("Buffers:          [0-9]+ kB")
	CachedPattern, err := regexp.Compile("Cached:          [0-9]+ kB")
	SwapCachedPattern, err := regexp.Compile("SwapCached:            [0-9]+ kB")

	if err != nil {
		fmt.Println(err.Error())
		return nil, errors.New("Error while creating the cpu object.")
	}

	file, err := ioutil.ReadFile("/proc/meminfo")
	if err != nil {
		return nil, errors.New("Error while reading the CPU MHz")
	}
	fileStr := string(file)

	arrayMemTotal := MemTotalPattern.FindAllString(fileStr, -1)
	arrayUsedMem := UsedMemPattern.FindAllString(fileStr, -1)
	arrayMemFree := MemFreePattern.FindAllString(fileStr, -1)
	arrayMemAvailable := MemAvailablePattern.FindAllString(fileStr, -1)
	arrayBuffers := BuffersPattern.FindAllString(fileStr, -1)
	arrayCached := CachedPattern.FindAllString(fileStr, -1)
	arraySwapCached := SwapCachedPattern.FindAllString(fileStr, -1)

	number, _ := regexp.Compile("[0-9]+ kB")

	var temp [][]string

	temp = make([][]string, 7)

	temp[0] = number.FindAllString(arrayMemTotal[0], -1)
	temp[1] = number.FindAllString(arrayUsedMem[0], -1)
	temp[2] = number.FindAllString(arrayMemFree[0], -1)
	temp[3] = number.FindAllString(arrayMemAvailable[0], -1)
	temp[4] = number.FindAllString(arrayBuffers[0], -1)
	temp[5] = number.FindAllString(arrayCached[0], -1)
	temp[6] = number.FindAllString(arraySwapCached[0], -1)

	var temp2 []int

	temp2 = make([]int, 7)

	temp2[0], _ = strconv.Atoi(temp[0][0])
	temp2[1], _ = strconv.Atoi(temp[1][0])
	temp2[2], _ = strconv.Atoi(temp[2][0])
	temp2[3], _ = strconv.Atoi(temp[3][0])
	temp2[4], _ = strconv.Atoi(temp[4][0])
	temp2[5], _ = strconv.Atoi(temp[5][0])
	temp2[6], _ = strconv.Atoi(temp[6][0])

	return &RamMem{
		totalMem:     temp2[0],
		usedMem:      temp2[1],
		freeMem:      temp2[2],
		availableMem: temp2[3],
		buffers:      temp2[4],
		cached:       temp2[5],
		swapedCache:  temp2[6],
	}, nil
}

func (self *RamMem) GetTotalMem() int {
	return self.totalMem
}

func (self *RamMem) GetUsedMem() int {
	return self.totalMem
}

func (self *RamMem) GetFreeMem() int {
	return self.totalMem
}

func (self *RamMem) GetAvailableMem() int {
	return self.totalMem
}

func (self *RamMem) GetBuffers() int {
	return self.totalMem
}

func (self *RamMem) GetCached() int {
	return self.cached
}

func (self *RamMem) GetSwapedCache() int {
	return self.totalMem
}

func (self *RamMem) UpdateMems() error {
	MemTotalPattern, err := regexp.Compile("MemTotal:        [0-9]+ kB")
	UsedMemPattern, err := regexp.Compile("MemTotal:        [0-9]+ kB")
	MemFreePattern, err := regexp.Compile("MemFree:          [0-9]+ kB")
	MemAvailablePattern, err := regexp.Compile("MemAvailable:    [0-9]+ kB")
	BuffersPattern, err := regexp.Compile("Buffers:          [0-9]+ kB")
	CachedPattern, err := regexp.Compile("Cached:          [0-9]+ kB")
	SwapCachedPattern, err := regexp.Compile("SwapCached:            [0-9]+ kB")

	file, err := ioutil.ReadFile("/proc/cpuinfo")
	if err != nil {
		return errors.New("Error while reading the memory")
	}
	fileStr := string(file)

	arrayMemTotal := MemTotalPattern.FindAllString(fileStr, -1)
	arrayUsedMem := UsedMemPattern.FindAllString(fileStr, -1)
	arrayMemFree := MemFreePattern.FindAllString(fileStr, -1)
	arrayMemAvailable := MemAvailablePattern.FindAllString(fileStr, -1)
	arrayBuffers := BuffersPattern.FindAllString(fileStr, -1)
	arrayCached := CachedPattern.FindAllString(fileStr, -1)
	arraySwapCached := SwapCachedPattern.FindAllString(fileStr, -1)

	number, _ := regexp.Compile("[0-9]+ kB")

	var temp [7][]string

	temp[0] = number.FindAllString(arrayMemTotal[0], -1)
	temp[1] = number.FindAllString(arrayUsedMem[0], -1)
	temp[2] = number.FindAllString(arrayMemFree[0], -1)
	temp[3] = number.FindAllString(arrayMemAvailable[0], -1)
	temp[4] = number.FindAllString(arrayBuffers[0], -1)
	temp[5] = number.FindAllString(arrayCached[0], -1)
	temp[6] = number.FindAllString(arraySwapCached[0], -1)

	var temp2 [7]int

	temp2[0], _ = strconv.Atoi(temp[0][0])
	temp2[1], _ = strconv.Atoi(temp[1][0])
	temp2[2], _ = strconv.Atoi(temp[2][0])
	temp2[3], _ = strconv.Atoi(temp[3][0])
	temp2[4], _ = strconv.Atoi(temp[4][0])
	temp2[5], _ = strconv.Atoi(temp[5][0])
	temp2[6], _ = strconv.Atoi(temp[6][0])

	self.totalMem = temp2[0]
	self.usedMem = temp2[1]
	self.freeMem = temp2[2]
	self.availableMem = temp2[3]
	self.buffers = temp2[4]
	self.cached = temp2[5]
	self.swapedCache = temp2[6]

	return nil
}

func (self *RamMem) StartRealTimeUpdate() {
	go self.startRealTime()
}

func (self *RamMem) startRealTime() {
	number, _ := regexp.Compile("[0-9]+")

	MemTotalPattern, _ := regexp.Compile("MemTotal:        [0-9]+ kB")
	UsedMemPattern, _ := regexp.Compile("MemTotal:        [0-9]+ kB")
	MemFreePattern, _ := regexp.Compile("MemFree:          [0-9]+ kB")
	MemAvailablePattern, _ := regexp.Compile("MemAvailable:    [0-9]+ kB")
	BuffersPattern, _ := regexp.Compile("Buffers:          [0-9]+ kB")
	CachedPattern, _ := regexp.Compile("Cached:          [0-9]+ kB")
	SwapCachedPattern, _ := regexp.Compile("SwapCached:            [0-9]+ kB")

	for {
		time.Sleep(time.Second * 2)

		file, err := ioutil.ReadFile("/proc/meminfo")
		if err != nil {
			return
		}
		fileStr := string(file)

		arrayMemTotal := MemTotalPattern.FindAllString(fileStr, -1)
		arrayUsedMem := UsedMemPattern.FindAllString(fileStr, -1)
		arrayMemFree := MemFreePattern.FindAllString(fileStr, -1)
		arrayMemAvailable := MemAvailablePattern.FindAllString(fileStr, -1)
		arrayBuffers := BuffersPattern.FindAllString(fileStr, -1)
		arrayCached := CachedPattern.FindAllString(fileStr, -1)
		arraySwapCached := SwapCachedPattern.FindAllString(fileStr, -1)

		var temp [][]string

		temp = make([][]string, 7)

		temp[0] = number.FindAllString(arrayMemTotal[0], -1)
		temp[1] = number.FindAllString(arrayUsedMem[0], -1)
		temp[2] = number.FindAllString(arrayMemFree[0], -1)
		temp[3] = number.FindAllString(arrayMemAvailable[0], -1)
		temp[4] = number.FindAllString(arrayBuffers[0], -1)
		temp[5] = number.FindAllString(arrayCached[0], -1)
		temp[6] = number.FindAllString(arraySwapCached[0], -1)

		var temp2 []int

		temp2 = make([]int, 7)

		temp2[0], _ = strconv.Atoi(temp[0][0])
		temp2[1], _ = strconv.Atoi(temp[1][0])
		temp2[2], _ = strconv.Atoi(temp[2][0])
		temp2[3], _ = strconv.Atoi(temp[3][0])
		temp2[4], _ = strconv.Atoi(temp[4][0])
		temp2[5], _ = strconv.Atoi(temp[5][0])
		temp2[6], _ = strconv.Atoi(temp[6][0])

		self.totalMem = temp2[0]
		self.usedMem = temp2[1]
		self.freeMem = temp2[2]
		self.availableMem = temp2[3]
		self.buffers = temp2[4]
		self.cached = temp2[5]
		self.swapedCache = temp2[6]
	}
}
