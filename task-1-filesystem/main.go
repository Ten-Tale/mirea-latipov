package main

import (
	"fmt"
	"log"

	"github.com/shirou/gopsutil/disk"
)

func main() {

	partitionList, err := disk.Partitions(true)

	checkError(err)

	for _, partition := range partitionList {
		partitionInfo, err := disk.Usage(partition.Mountpoint)

		checkError(err)

		fmt.Printf("Disk name: %s\n", partition.Device)
		fmt.Printf("File system type: %s\n", partition.Fstype)

		fmt.Printf("Mount path: %s\n", partitionInfo.Path)
		fmt.Printf("Total disk size: %d\n", partitionInfo.Total)
		fmt.Printf("Used disk space: %d\n", partitionInfo.Used)
		fmt.Printf("Free disk space: %d\n", partitionInfo.Free)

		fmt.Printf("\n-----------------------------------------\n")
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
