package main

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUploadTest(t *testing.T) {
	assert := assert.New(t)
	path := "C:/Users/hotqk/Downloads/my.jpg"
	file, _ := os.Open(path) //파일 오픈, 파일 없는 경우 에러반환
	defer file.Close()       //파일 열면 닫아줘야 함

	os.RemoveAll("./uploads")

	buf := &bytes.Buffer{}                                                  //버퍼를 만들어준다
	writer := multipart.NewWriter(buf)                                      //웹으로 파일같은 데이터를 전송할떄 포멧 MIME, 멀티파트 라이터를만들어야한다
	multi, err := writer.CreateFormFile("upload_file", filepath.Base(path)) //폼 파일하나를 만듬
	assert.NoError(err)

	io.Copy(multi, file)
	writer.Close()

	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/uploads", buf)
	req.Header.Set("Content-type", writer.FormDataContentType())

	uploadsHandler(res, req)
	assert.Equal(res.Code, http.StatusOK)

	uploadFilePath := "./uploads/" + filepath.Base(path)
	_, err = os.Stat(uploadFilePath) //파일 info를 준다 파일 있는지 확인
	assert.NoError(err)

	uploadFile, _ := os.Open(uploadFilePath)
	originFile, _ := os.Open(path)
	defer uploadFile.Close()
	defer originFile.Close()

	uploadData := []byte{}
	originData := []byte{}
	uploadFile.Read(uploadData)
	uploadFile.Read(originData)

	assert.Equal(originData, uploadData)
}
