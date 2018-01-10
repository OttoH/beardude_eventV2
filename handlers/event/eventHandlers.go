package eventHandlers

import (
  "log"
  "strconv"
  "net/http"
  "github.com/gin-gonic/gin"

  "beardude_eventV2/dao"
  "beardude_eventV2/models"
)

type DepInj struct {
	Dao *dao.DAO
}

// TODO: bind JSON from context params
func (dep *DepInj) CreateEvent(c *gin.Context) {
  ids, idErr := strconv.Atoi(c.PostForm("managerIds"))
  if idErr != nil {
    log.Println("ManagerIds to int error: ", idErr)
    return
  }

  eventModel := &models.Events {
    Event: &models.Event {
      ManagerIds: ids,
      UniqueName : c.PostForm("name"),
      StartTime : c.PostForm("startTime"),
      EndTime : c.PostForm("endTime"),
    },
    }

  err := dep.Dao.CreateEvent(eventModel)

  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"status": "fail","error": err})
  } else {
    c.JSON(http.StatusOK, gin.H{"status": "success"})
  }
}
