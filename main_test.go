package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Запрос сформирован корректно, сервис возвращает код ответа 200 и тело ответа не пустое
func TestMainHandlerWhenRequestCorrect(t *testing.T) {
	cafeCount := 3
	req := httptest.NewRequest("GET", "/cafe?count=3&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// код ответа = 200
	require.Equal(t, http.StatusOK, responseRecorder.Code)

	// тело ответа не пустое
	body := responseRecorder.Body.String()
	require.NotEmpty(t, body)

	// проверяем количество кафе, оно должно быть равным cafeCount
	list := strings.Split(body, ",")
	require.Len(t, list, cafeCount)
}

// Если в параметре count указано больше, чем есть всего, должны вернуться все доступные кафе.
func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=100&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// код ответа = 200
	require.Equal(t, http.StatusOK, responseRecorder.Code)

	// тело ответа не пустое
	body := responseRecorder.Body.String()
	require.NotEmpty(t, body)

	// проверяем количество кафе, оно должно быть равным totalCount
	list := strings.Split(body, ",")
	require.Len(t, list, totalCount)
}

// Город, который передаётся в параметре city, не поддерживается. Сервис возвращает код ответа 400 и ошибку wrong city value в теле ответа.
func TestMainHandlerWhenWrongCityValue(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=100&city=chaikovsky", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// код ответа = 400
	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
	// ошибка wrong city value
	assert.Equal(t, "wrong city value", responseRecorder.Body.String())
}
