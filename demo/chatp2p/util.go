package chatp2p

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
)

func Serialize(obj interface{}) string {
	marshal, _ := json.Marshal(obj)
	return string(marshal)
}

func UnSerialize(data []byte, obj interface{}) {
	json.Unmarshal(data, obj)
}

func Md5(arg string) string {
	hash := md5.New()
	hash.Write([]byte(arg))
	return fmt.Sprintf("%x", hash.Sum(nil))
}
