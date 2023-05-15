package router

import (
	"bytes"
	"fmt"
	"io"
	"k8s.io/klog/v2"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/huaweicloud/huaweicloud-sdk-go-obs/obs"
	"go-learning/gin-obs/types"
)

func Upload(r *gin.Context) {
	obsClient, err := obs.New(types.AccessKey, types.SecretKey, types.Endpoint)
	if err != nil {
		klog.Fatalf("unable to initialize obsClient : %v", err)
	}

	file, header, err := r.Request.FormFile("file")
	if err != nil {
		r.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}
	defer file.Close()

	buffer := bytes.NewBuffer(nil)
	_, err = io.Copy(buffer, file)
	if err != nil {
		panic(err)
	}
	reader := bytes.NewReader(buffer.Bytes())
	if err != nil {
		r.String(http.StatusBadRequest, fmt.Sprintf("read file err: %s", err.Error()))
		return
	}

	if err != nil {
		r.String(http.StatusBadRequest, fmt.Sprintf("create gin-obs client err: %s", err.Error()))
		return
	}

	input := &obs.PutObjectInput{}
	input.Bucket = "XXXX"
	input.Key = "test/" + header.Filename
	input.Body = reader

	_, err = obsClient.PutObject(input)
	if err != nil {
		r.String(http.StatusBadRequest, fmt.Sprintf("put object err: %s", err.Error()))
		return
	}

	r.String(http.StatusOK, fmt.Sprintf("upload file %s success", header.Filename))
}
