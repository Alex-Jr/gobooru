package controllers_test

import (
	"bytes"
	"context"
	"encoding/json"
	"gobooru/internal/controllers"
	"gobooru/internal/dtos"
	"gobooru/internal/fixtures/fakes"
	"gobooru/internal/mocks"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/textproto"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPostControllerCreate(t *testing.T) {
	postService := mocks.NewMockPostService(t)

	postController := controllers.NewPostController(controllers.PostControllerConfig{
		PostService: postService,
	})

	args := struct {
		Description string
	}{
		Description: "test description",
	}

	want := struct {
		statusCode int
		dto        dtos.CreatePostResponseDTO
	}{
		statusCode: http.StatusOK,
		dto: dtos.CreatePostResponseDTO{
			Post: fakes.Post1,
		},
	}

	requestBody := new(bytes.Buffer)
	writer := multipart.NewWriter(requestBody)

	err := writer.WriteField("description", args.Description)
	require.NoError(t, err)

	fileHeader := make(textproto.MIMEHeader)
	fileHeader.Set("Content-Disposition", `form-data; name="file"; filename="example.jpg"`)
	fileHeader.Set("Content-Type", "text/plain")

	file, err := writer.CreatePart(fileHeader)
	require.NoError(t, err)

	file.Write([]byte("test file"))

	writer.Close()

	req, err := http.NewRequest(
		http.MethodPost,
		"/posts",
		requestBody,
	)
	require.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, writer.FormDataContentType())

	rec := httptest.NewRecorder()

	e := echo.New()
	c := e.NewContext(req, rec)

	formFile, _ := c.FormFile("file")

	postService.On(
		"Create",
		context.Background(),
		dtos.CreatePostDTO{
			Description: args.Description,
			File:        formFile,
		},
	).Return(
		want.dto,
		nil,
	)

	err = postController.Create(c)
	require.NoError(t, err)

	var responseDTO dtos.CreatePostResponseDTO
	err = json.Unmarshal(rec.Body.Bytes(), &responseDTO)
	require.NoError(t, err)

	assert.Equal(t, want.statusCode, rec.Code)
	assert.EqualValues(t, want.dto, responseDTO)
}
