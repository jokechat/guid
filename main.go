package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jokechat/guid/common"
	"github.com/jokechat/snowflake"
	"log"
	"time"
)

func main() {
	r := gin.Default()
	workId, err := common.GetWorkId()

	epoch := time.Date(2019, time.January, 1, 0, 0, 0, 0, time.Local)

	w, _ := snowflake.NewWorkerWithOpts(
		snowflake.WithEpoch(epoch),
		snowflake.WithWorkerId(1),
	)

	r.GET("/", func(context *gin.Context) {
		id := w.Next()
		d := common.GuidData{
			ID:     id.Uint64(),
			WorkId: workId,
			Base32: id.Base32Lower(),
			Time:   id.Time(w),
		}

		result := &common.Result{
			Code: 0,
			Msg:  "success",
			Data: d,
		}
		context.JSON(200, result)
	})
	err = r.Run("0.0.0.0:80")
	if err != nil {
		log.Fatalln(err)
	}
}
