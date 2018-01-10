package racerHandlers

import (
  "log"
  "net/http"
  "github.com/gin-gonic/gin"

  "beardude_eventV2/dao"
  "beardude_eventV2/models"
)

type DepInj struct {
	Dao *dao.DAO
}

/*
 * post params: username, password
 */
func (dep *DepInj) CreateRacer(c *gin.Context) {
  var racerModel models.Racer

  if err := c.BindJSON(&racerModel);err == nil && len(racerModel.Username) > 0 {
    racerModel.Activate = 1

    if err := dep.Dao.CreateRacer(&racerModel); err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "error": err.Error()})
    } else {
      c.JSON(http.StatusOK, gin.H{"status": "success"})
    }
  } else {
    log.Println("CreateRacer: none infor input")
    c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "error": "CreateRacer: none email input"})
  }
}

/*
 * post params: username
 */
func (dep *DepInj) GetRacer(c *gin.Context) {
  var racerModel models.Racer
  if err := c.BindJSON(&racerModel);err == nil && len(racerModel.Username) > 0 {
    res, err := dep.Dao.GetRacerByName(racerModel.Username)
    if err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"status": "fail","error": err})
    } else {
      c.JSON(http.StatusOK, gin.H{"status": "success", "result": res})
    }
  } else {
    c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "error": "GetRacer: none info param input"})
  }
}

/*
 * post params: username, password, nickname
 */
func (dep *DepInj) UpdateRacer(c *gin.Context) {
  var racerModel models.Racer
  if err := c.BindJSON(&racerModel);err == nil && len(racerModel.Username) > 0 {
    err := dep.Dao.UpdateRacer(&racerModel)
    if err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"status": "fail","error": err})
    } else {
      c.JSON(http.StatusOK, gin.H{"status": "success"})
    }
  } else {
    c.JSON(http.StatusBadRequest, gin.H{"status": "fail","error": "UpdateRacer: none info param"})
  }
}

/*
 * post params: username
 */
func (dep *DepInj) RemoveRacer (c *gin.Context) {
  var racerModel models.Racer
  if err := c.BindJSON(&racerModel);err == nil && len(racerModel.Username) > 0 {
    racerModel.Activate = 0
    err := dep.Dao.RemoveRacer(&racerModel)

    if err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"status": "fail","error": err})
    } else {
      c.JSON(http.StatusOK, gin.H{"status": "success"})
    }
  } else {
    c.JSON(http.StatusBadRequest, gin.H{"status": "fail","error": "RemoveRacer: none info param"})
  }
}
