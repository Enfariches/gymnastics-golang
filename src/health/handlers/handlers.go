package handlers

import (
	"health/src/health/models"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
)

/*
Полностью повтораются методы из Controllers. Не используемый модуль
*/

func GetVideosListHandler(c *gin.Context) {
	videos := models.GetVideosList()
	c.JSON(200, videos)
}
func GetSheduleHandler(c *gin.Context) {

}
func UploadVideo(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(500, err)
		return
	}
	vid, err := models.AddVideo()
	if err != nil {
		c.JSON(500, err)
		return
	}
	if err := c.SaveUploadedFile(file, "upload/"+strconv.Itoa(int(vid.ID))+filepath.Ext(file.Filename)); err != nil {
		c.JSON(500, err)
		return
	}

}

func DeleteVideoHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(500, err)
		return
	}
	models.DeleteVideo(uint(id))
}

func AddSheduleHandler(c *gin.Context) {
	dayStr := c.Param("day")
	day, err := strconv.Atoi(dayStr)
	if err != nil {
		c.JSON(500, err)
		return
	}
	shed := models.AddSchedule(uint8(day))
	c.JSON(200, shed)
}

func ListSheduleHandler(c *gin.Context) {
	dayStr := c.Param("day")
	day, err := strconv.Atoi(dayStr)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, models.ListSchedule(uint8(day)))
}

func UpdSheduleHandler(c *gin.Context) {
	var shed models.Shedule
	err := c.BindJSON(&shed)
	if err != nil {
		c.JSON(500, err)
		return
	}
	shed.Update()
}

func DeleteSheduleHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(500, err)
		return
	}
	models.DeleteSchedule(uint(id))
}

func ToggleVideoHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(500, err)
		return
	}
	err = models.ToggleVid(uint(id))
	if err != nil {
		c.JSON(500, err)
		return
	}
}

func UpdVideoHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(500, err)
		return
	}
	var req struct {
		Name string
	}
	err = c.BindJSON(&req)
	if err != nil {
		c.JSON(500, err)
		return
	}
	vid, err := models.UpdVid(uint(id), req.Name)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, vid)
}

func AddStatHandler(c *gin.Context) {
	var req struct {
		VID  uint
		Name string
		Type string
	}
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(400, err)
		return
	}

	models.AddStat(req.VID, req.Name, req.Type)
}

func ListStatsHandler(c *gin.Context) {
	c.JSON(200, models.ListStats())
}
