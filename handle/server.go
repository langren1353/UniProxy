package handle

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/wyx2685/UniProxy/v2b"
)

var servers map[string]*v2b.ServerInfo
var updateTime time.Time

func GetServers(c *gin.Context) {
	if len(servers) != 0 {
		if time.Now().Before(updateTime) {
			c.JSON(200, Rsp{
				Success: true,
				Data:    servers,
			})
			return
		}
	}
	r, err := v2b.GetServers()
	if err != nil {
		log.Error("get server list error: ", err)
		c.JSON(400, Rsp{Success: false, Message: err.Error()})
		return
	}
	updateTime = time.Now().Add(180 * time.Hour)
	if len(r) != 0 {
		servers = make(map[string]*v2b.ServerInfo, len(r))
		for i := range r {
			servers[fmt.Sprintf(
				"%d_%s_%d",
				i + 1000,
				r[i].Type,
				r[i].Id,
			)] = &r[i]
		}
	}
	c.JSON(200, Rsp{
		Success: true,
		Data:    servers,
	})
	return
}
