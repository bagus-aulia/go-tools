package general

import (
	"bytes"
	"errors"
	"io"
	"mime/multipart"
)

// GenFileHeaderFaker used to Generate File Header for unit test purpose
func GenFileHeaderFaker() (*multipart.FileHeader, error) {
	// Create a buffer to simulate an uploaded file
	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)

	// Create a form file field for the file
	fileWriter, err := writer.CreateFormFile("file", "test.png")
	if err != nil {
		return nil, err
	}

	// Write some content to the file
	_, err = io.Copy(fileWriter, bytes.NewReader([]byte("test image content")))
	if err != nil {
		return nil, err
	}

	// Add the additional parameter to the form
	err = writer.WriteField("additionalParam", "test-param")
	if err != nil {
		return nil, err
	}

	// Close the multipart writer to set the terminating boundary
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	// Parse the multipart form to get the file header
	req := multipart.NewReader(&buffer, writer.Boundary())
	form, err := req.ReadForm(1024)
	if err != nil {
		return nil, err
	}
	fileHeaders := form.File["file"]
	if len(fileHeaders) == 0 {
		err := errors.New("unprocessable entity")
		return nil, err
	}

	return fileHeaders[0], nil
}
