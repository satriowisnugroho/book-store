package response_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/satriowisnugroho/book-store/internal/response"
	"github.com/stretchr/testify/assert"
)

func TestCustomErrorError(t *testing.T) {
	assert.Equal(t, "Record not found", response.ErrNotFound.Error())
}

func TestErrorInfoError(t *testing.T) {
	assert.Contains(t, response.ErrorInfo{}.Error(), "error - msg:")
}

func TestErrorBodyError(t *testing.T) {
	assert.Contains(t, response.ErrorBody{Errors: []response.ErrorInfo{{}}}.Error(), "response - errors")
}

func TestErrorResponseError(t *testing.T) {
	assert.Contains(t, response.ErrorResponse{}.Error(), "response - errors")
}

func TestErrUnauthorized(t *testing.T) {
	assert.Equal(t, http.StatusUnauthorized, response.ErrUnauthorized("").HTTPCode)
}

func TestBuildSuccess(t *testing.T) {
	assert.Equal(t, "foo", response.BuildSuccess("", "foo", "").Message)
}

func TestOK(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	response.OK(c, "", "foo")

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestOKWithPagination(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	response.OKWithPagination(c, "", "foo", 0, 0, 0)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestBuildError(t *testing.T) {
	tests := []struct {
		name     string
		input    []error
		expected response.ErrorBody
	}{
		{
			name:     "No errors",
			input:    nil,
			expected: response.InternalServerErrorBody(),
		},
		{
			name: "CustomError",
			input: []error{response.CustomError{
				HTTPCode: http.StatusBadRequest,
			}},
			expected: response.ErrorBody{
				Errors: []response.ErrorInfo{{}},
				Meta: response.MetaInfo{
					HTTPStatus: http.StatusBadRequest,
				},
			},
		},
		{
			name:  "ErrorInfo",
			input: []error{response.ErrorInfo{}},
			expected: response.ErrorBody{
				Errors: []response.ErrorInfo{{}},
			},
		},
		{
			name: "ErrorBody",
			input: []error{response.ErrorBody{
				Errors: []response.ErrorInfo{{}},
			}},
			expected: response.ErrorBody{
				Errors: []response.ErrorInfo{{}},
			},
		},
		{
			name: "ErrorResponse",
			input: []error{response.ErrorResponse{
				ErrorBody: response.ErrorBody{
					Errors: []response.ErrorInfo{{}},
				},
			}},
			expected: response.ErrorBody{
				Errors: []response.ErrorInfo{{}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := response.BuildError(tt.input...)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestError(t *testing.T) {
	tests := []struct {
		name           string
		err            error
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "context canceled",
			err:            context.Canceled,
			expectedStatus: 0,
		},
		{
			name:           "context deadline exceeded",
			err:            context.DeadlineExceeded,
			expectedStatus: 0,
		},
		{
			name:           "custom error",
			err:            response.ErrNotFound,
			expectedStatus: http.StatusNotFound,
		},
		{
			name:           "custom error without meta",
			err:            response.ErrorInfo{},
			expectedStatus: http.StatusUnprocessableEntity,
		},
		{
			name:           "generic error with cause",
			err:            errors.New("generic error"),
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			response.Error(c, tt.err)

			if tt.expectedStatus != 0 {
				assert.Equal(t, tt.expectedStatus, w.Code)
			}
		})
	}
}
