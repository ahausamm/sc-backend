package token

import (
	"github.com/gin-gonic/gin"
	/*"crypto/sha256"
	"bytes"
	"reflect"
	"encoding/base64"*/
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Token struct {
	Token        string
}
type Error struct {
	Message		string
}

func isInstanceAllowed(InstanceId string, UserId string) bool{
	db, err := sql.Open("mysql", "scmanage:scmanage@tcp(localhost:3306)/scmanage")
	if err != nil {
        return false
	}
	defer db.Close()
	instanceOut, err := db.Prepare("SELECT instance_id FROM instances WHERE instance_id = ? AND user_id = ?")
    if err != nil {
        return false
    }
    defer instanceOut.Close()
    var InstanceIdFromDb string
    err = instanceOut.QueryRow(InstanceId,UserId).Scan(&InstanceIdFromDb)
    if err != nil {
        return false
    }
    return true
}

func CreateToken(c *gin.Context) {
	InstanceId := c.PostForm("InstanceId")
	UserId := c.PostForm("UserId")
	//AuthorizationHeader := c.Request.Header.Get("Authorization")
	if(InstanceId != "" && UserId != "") {
		if(isInstanceAllowed(InstanceId,UserId)) {
			/*h256 := sha256.New()
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
		 	}*/
		} else {
			var error = Error{Message: "forbidden"}
			c.JSON(403,error)
		}
	} else {
		var error = Error{Message: "forbidden"}
		c.JSON(403,error)
	}
}
