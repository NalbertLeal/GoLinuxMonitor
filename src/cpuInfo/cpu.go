package cpuInfo

import (
	"errors"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
)

type Cpu struct {
	numberOfCores int
	cpuMHz        float64
	cores         []*Core
}

func NewCpu(numberOfCores int) (*Cpu, error) {
	tempCores := make([]*Core, numberOfCores)
	for i := 0; i < numberOfCores; i++ {
		tempCores[i] = NewCore(i)
	}

	regexpPattern, err := regexp.Compile("processor\t: .\n[A-Za-z\\_\t: ]+\n[A-Za-z\\_\t: 0-9]+\n[A-Za-z\\_\t: 0-9]+\n[A-Za-z\\_\t: 0-9\\(\\)\\-@\\.]+")
	if err != nil {
		fmt.Println(err.Error())
		return nil, errors.New("Error while creating the cpu object.")
	}

	file, err := ioutil.ReadFile("/proc/cpuinfo")
	if err != nil {
		return nil, errors.New("Error while reading the CPU MHz")
	}
	fileStr := string(file)
	arrayReturn := regexpPattern.FindAllString(fileStr, -1)

	mhz, _ := strconv.ParseFloat(arrayReturn[0][len(arrayReturn[0])-7:len(arrayReturn[0])-3], 64)

	return &Cpu{
		numberOfCores: numberOfCores,
		cpuMHz:        mhz * 1000,
		cores:         tempCores,
	}, nil
}

func (self Cpu) GetCpuMHz() float64 {
	return self.cpuMHz
}

func (self Cpu) GetCoresInfo() {
	for i := 0; i < self.numberOfCores; i++ {
		self.cores[i].UpdateCoreInfoNow()
	}
}

func (self Cpu) StartCoresInfoRealTime() {
	for i := 0; i < self.numberOfCores; i++ {
		self.cores[i].RealTimeUpdate()
	}
}

func (self Cpu) GetCoreUsage(coreNumber int) float64 {
	return self.cores[coreNumber].GetCoreUsage()
}
