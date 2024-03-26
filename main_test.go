package main

import (
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
)


func TestMainHandlerWhenGoodRequest(t *testing.T) {
    req := httptest.NewRequest("GET", "/cafe?city=moscow&count=1", nil)

    responseRecorder := httptest.NewRecorder()
    handler := http.HandlerFunc(mainHandle)
    handler.ServeHTTP(responseRecorder, req)

    require.Equal(t, responseRecorder.Code, http.StatusOK)
    assert.NotEmpty(t, responseRecorder.Body)
}


func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
    totalCount := 4

    req := httptest.NewRequest("GET", "/cafe?city=moscow&count=1001", nil)

    responseRecorder := httptest.NewRecorder()
    handler := http.HandlerFunc(mainHandle)
    handler.ServeHTTP(responseRecorder, req)

    sliceRespBody := strings.Split(responseRecorder.Body.String(), ",")

    require.Equal(t, responseRecorder.Code, http.StatusOK)
    assert.Len(t, sliceRespBody, totalCount)
}


func TestMainHandlerWhenNotFoundCity(t *testing.T) {
    requireRespBody := "wrong city value"
    req := httptest.NewRequest("GET", "/cafe?city=ufa&count=1", nil)

    responseRecorder := httptest.NewRecorder()
    handler := http.HandlerFunc(mainHandle)
    handler.ServeHTTP(responseRecorder, req)

    require.Equal(t, responseRecorder.Code, http.StatusBadRequest)
    assert.Equal(t, responseRecorder.Body.String(), requireRespBody)
}

