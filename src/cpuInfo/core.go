package cpuInfo

import (
	"fmt"
	"time"
	"regexp"
	"strconv"
	"io/ioutil"
)

type Core struct {
	coreNumber	int
	coreUsage	float64
	coreUsagePercentage	int
}

func NewCore(coreNumber int) *Core {
	return &Core{
		coreNumber: coreNumber,
		coreUsagePercentage: 0,
	}
}

func (self *Core) GetCoreNumber() int {
	return self.coreNumber
}

func (self *Core) GetCoreUsage() float64 {
	return self.coreUsage
}

func (self *Core) UpdateCoreInfoNow() {
	regexpPattern, err := regexp.Compile("processor\t: " + strconv.Itoa(self.coreNumber) + "\n[A-Za-z\\_\t: ]+\n[A-Za-z\\_\t: 0-9]+\n[A-Za-z\\_\t: 0-9]+\n[A-Za-z\\_\t: 0-9\\(\\)\\-@\\.]+\n[A-Za-z\\_\t: 0-9]+\n[A-Za-z\\_\t: 0-9]+\ncpu MHz\t\t: [0-9\\.]+")
	if err != nil {
		fmt.Println(err.Error())
	}

	file, err := ioutil.ReadFile("/proc/cpuinfo")
	fileStr := string(file)
	ArrayCoreUsage := regexpPattern.FindAllString(fileStr, -1)
	self.coreUsage, _ = strconv.ParseFloat(ArrayCoreUsage[0][len(ArrayCoreUsage[0])-7:len(ArrayCoreUsage[0])], 64)
}

func (self *Core) RealTimeUpdate() {
	go self.realTimeUpdate()
}

func (self *Core) realTimeUpdate() {
	regexpPattern, err := regexp.Compile("processor\t: " + strconv.Itoa(self.coreNumber) + "\n[A-Za-z\\_\t: ]+\n[A-Za-z\\_\t: 0-9]+\n[A-Za-z\\_\t: 0-9]+\n[A-Za-z\\_\t: 0-9\\(\\)\\-@\\.]+\n[A-Za-z\\_\t: 0-9]+\n[A-Za-z\\_\t: 0-9]+\ncpu MHz\t\t: [0-9\\.]+")
	if err != nil {
		fmt.Println(err.Error())
	}
	for {
		file, err := ioutil.ReadFile("/proc/cpuinfo")
		if err != nil {
			return
		}
		fileStr := string(file)
		ArrayCoreUsage := regexpPattern.FindAllString(fileStr, -1)
		self.coreUsage, _ = strconv.ParseFloat(ArrayCoreUsage[0][len(ArrayCoreUsage[0])-7:len(ArrayCoreUsage[0])], 64)
		time.Sleep(time.Second * 2)
	}
}