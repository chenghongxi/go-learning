package router

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/huaweicloud/huaweicloud-sdk-go-obs/obs"
	"k8s.io/klog/v2"

	"go-learning/gin-obs/types"
)

func Download(r *gin.Context) {
	// 下载文件
	obsClient, err := obs.New(types.AccessKey, types.SecretKey, types.Endpoint)
	if err != nil {
		klog.Fatalf("unable to initialize obsClient : %v", err)
	}

	filename := r.Query("filename")
	input := &obs.GetObjectInput{}
	input.Bucket = "XXXXX"
	input.Key = filename

	output, err := obsClient.GetObject(input)
	if err != nil {
		r.String(http.StatusBadRequest, fmt.Sprintf("get object err: %s", err.Error()))
		return
	}
	defer output.Body.Close()

	data, _ := io.ReadAll(output.Body)

	r.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	r.Writer.Header().Add("Content-Type", "application/octet-stream")
	r.Writer.Write(data)

}
