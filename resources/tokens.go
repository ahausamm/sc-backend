package tokens

import (
	"github.com/gin-gonic/gin"
	"crypto/sha256"
	"bytes"
	"reflect"
)

type Token struct {
	Token        string
}
type Error struct {
	Message		string
}

type Instance struct {
	InstanceId 	string
	Ip 			string
}

type Instances []Instance

var allowedInstance = Instances{
	Instance{
		InstanceId: "1",
		Ip: "192.168.0.1",
	},
	Instance{
		InstanceId: "2",
		Ip: "192.168.0.2",
	},
}
func isInstanceAllowed(InstanceId string) bool{
	for _, allInstances := range allowedInstance {
 		if allInstances.InstanceId == InstanceId {
 			return true
 		}
 	}
 	return false
}

func CreateToken(c *gin.Context) {
	InstanceId := c.Params.ByName("InstanceId")
	EncryptedString := c.Params.ByName("EncryptedString")
	if(isInstanceAllowed(InstanceId)) {
		h256 := sha256.New()
		var EncryptedStringNew bytes.Buffer
		for _, allInstances := range allowedInstance {
	 		if allInstances.InstanceId == InstanceId {
	 			EncryptedStringNew.WriteString(allInstances.InstanceId)
	 			EncryptedStringNew.WriteString("-")
	 			EncryptedStringNew.WriteString(allInstances.Ip)
	 		}
	 	}
	 	h256.Write([]byte(EncryptedStringNew.String()))
		if reflect.DeepEqual([]byte(EncryptedString),h256.Sum(nil)) {
			var token = Token{Token: "hallo"}
			c.JSON(200, token)
	 	} else {
			var error = Error{Message: "not allowed"}
			c.JSON(403,error)
	 	}
	} else {
		var error = Error{Message: "not found"}
		c.JSON(404,error)
	}
}
