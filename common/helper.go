package common

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetWorkId() (uint64, error) {
	hostname := os.Getenv("HOSTNAME")
	if hostname == "" {
		return 1, nil
	}
	arr := strings.Split(hostname, "-")
	workId, err := strconv.ParseUint(arr[len(arr)-1], 10, 64)
	if err != nil {
		return 0, fmt.Errorf("parse workerId:%s, error:%s", hostname, err)
	}
	return workId, err
}
