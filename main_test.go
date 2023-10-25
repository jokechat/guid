package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/jokechat/guid/common"
	"testing"
)

func TestId(t *testing.T) {
	workId, err := common.GetWorkId()
	spew.Dump(workId)
	spew.Dump(err)
}
