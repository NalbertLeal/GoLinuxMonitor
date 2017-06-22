package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"time"

	ci "./cpuInfo"
	mer "./memInfo"
)

func main() {
	fmt.Println("Starting GoCPU ...")

	cpu, err := ci.NewCpu(runtime.NumCPU())
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	cpu.StartCoresInfoRealTime()

	mem, err := mer.NewRamMem()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	mem.StartRealTimeUpdate()

	outFileCpu, err := os.Create("GoCpuOutputCpu.dat")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	outFileMem, err := os.Create("GoCpuOutputMem.dat")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Running ...")

	var coreInfo float64      // real usage now Ghz
	var totalCoreInfo float64 // total Ghz
	var coreGHz float64

	coreGHz = cpu.GetCpuMHz()
	totalCoreInfo = 0
	for {
		// cpu
		time.Sleep(time.Second * 3)
		for i := 0; i < runtime.NumCPU(); i++ {
			coreInfo = cpu.GetCoreUsage(i)
			totalCoreInfo = totalCoreInfo + (coreInfo/coreGHz)*100
		}
		totalCoreInfo = totalCoreInfo / 2
		outFileCpu.Write([]byte("" + strconv.FormatFloat((totalCoreInfo/4), 'f', 1, 64) + "\n"))
		totalCoreInfo = 0

		// memory
		totalMem := mem.GetTotalMem()
		usedMem := mem.GetUsedMem()
		freeMem := mem.GetFreeMem()
		availableMem := mem.GetAvailableMem()
		buffers := mem.GetBuffers()
		cached := mem.GetCached()
		swapedCache := mem.GetSwapedCache()

		outFileMem.Write([]byte("" + strconv.Itoa(totalMem) + "," + strconv.Itoa(usedMem) + "," + strconv.Itoa(freeMem) + "," + strconv.Itoa(availableMem) + "," + strconv.Itoa(buffers) + "," + strconv.Itoa(cached) + "," + strconv.Itoa(swapedCache) + "\n"))

	}
}
