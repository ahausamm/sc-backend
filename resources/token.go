package token

import (
	"github.com/gin-gonic/gin"
	"crypto/sha256"
	"bytes"
	"reflect"
	"encoding/base64"
)

type Token struct {
	Token        string
}
type Error struct {
	Message		string
}

type Instance struct {
	InstanceId 	string
	UserId		string
	Ip 			string
}

type Instances []Instance

var allowedInstance = Instances{
	Instance{
		InstanceId: "1",
		UserId: "a",
		Ip: "192.168.0.102",
	},
	Instance{
		InstanceId: "2",
		UserId: "b",
		Ip: "192.168.0.2",
	},
}
func isInstanceAllowed(InstanceId string, UserId string) bool{
	for _, allInstances := range allowedInstance {
 		if allInstances.InstanceId == InstanceId && allInstances.UserId == UserId {
 			return true
 		}
 	}
 	return false
}

func CreateToken(c *gin.Context) {
	InstanceId := c.PostForm("InstanceId")
	UserId := c.PostForm("UserId")
	AuthorizationHeader := c.Request.Header.Get("Authorization")
	if(InstanceId != "" && UserId != "") {
		if(isInstanceAllowed(InstanceId,UserId)) {
			h256 := sha256.New()
			var EncryptedStringNew bytes.Buffer
			for _, allInstances := range allowedInstance {
		 		if allInstances.InstanceId == InstanceId {
		 			EncryptedStringNew.WriteString(allInstances.InstanceId)
		 			EncryptedStringNew.WriteString("-")
		 			EncryptedStringNew.WriteString(allInstances.Ip)
		 			EncryptedStringNew.WriteString("-")
		 			EncryptedStringNew.WriteString(allInstances.UserId)
		 		}
		 	}
		 	h256.Write([]byte(EncryptedStringNew.String()))
			if reflect.DeepEqual([]byte(AuthorizationHeader),[]byte(base64.StdEncoding.EncodeToString(h256.Sum(nil)))) {
				var token = Token{Token: "hallo"}
				c.JSON(200, token)
		 	} else {
				var error = Error{Message: "forbidden"}
				c.JSON(403,error)
		 	}
		} else {
			var error = Error{Message: "forbidden"}
			c.JSON(403,error)
		}
	} else {
		var error = Error{Message: "forbidden"}
		c.JSON(403,error)
	}
}
