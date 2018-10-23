package file

import (
	"fmt"
	"github.com/pkg/errors"
	"gopkg.in/gin-gonic/gin.v1"
	"io"
	"mime/multipart"
	"os"
	"path"
	"strings"
	"time"
)

func GetFileUrl(shortPath, uploadkey, html, static string, c *gin.Context) (fileUrl, fileName string, err error){
	var file multipart.File
	var header *multipart.FileHeader
	if html == "" && c != nil {
		file, header, err = c.Request.FormFile(uploadkey)
		if err != nil {
			return "", "", err
		}
		if file == nil {
			return "", "", errors.New("file 为空，请检查")
		}
		hz := strings.Split(header.Filename, ".")
		if len(hz) == 2 {
			fileName = fmt.Sprintf("%v_%v.%v", hz[0], time.Now().Nanosecond(), hz[1])
		} else {
			return "", "", errors.New("无法识别的文件类型，请检查")
		}
		
		filePath := path.Join(static,"upload", shortPath )
		fileUrl = path.Join(filePath, fileName)
		if !PathExists(filePath) {
			os.MkdirAll(filePath, os.ModePerm)
		}
		out, err := os.Create(fileUrl)
		if err != nil {
			return "", "", err
		}
		defer out.Close()
		_, err = io.Copy(out, file)
		if err != nil {
			return "", "", err
		}
		
	}
	
	
	return fileUrl, fileName, err
}

func PathExists(filePath string) bool {
	_,err := os.Stat(filePath)
	if err != nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
