package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	
	// Configurar rotas de teste
	router.POST("/auth/register", Register)
	router.POST("/auth/login", Login)
	
	return router
}

func TestRegister(t *testing.T) {
	router := setupTestRouter()
	
	tests := []struct {
		name       string
		payload    map[string]interface{}
		wantStatus int
	}{
		{
			name: "Registro válido",
			payload: map[string]interface{}{
				"name":     "João Silva",
				"email":    "joao@test.com",
				"password": "123456",
			},
			wantStatus: http.StatusCreated,
		},
		{
			name: "Email inválido",
			payload: map[string]interface{}{
				"name":     "João Silva",
				"email":    "email-invalido",
				"password": "123456",
			},
			wantStatus: http.StatusBadRequest,
		},
		{
			name: "Senha muito curta",
			payload: map[string]interface{}{
				"name":     "João Silva",
				"email":    "joao@test.com",
				"password": "123",
			},
			wantStatus: http.StatusBadRequest,
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			payloadBytes, _ := json.Marshal(tt.payload)
			req, _ := http.NewRequest("POST", "/auth/register", bytes.NewBuffer(payloadBytes))
			req.Header.Set("Content-Type", "application/json")
			
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			
			assert.Equal(t, tt.wantStatus, w.Code)
		})
	}
}

func TestLogin(t *testing.T) {
	router := setupTestRouter()
	
	tests := []struct {
		name       string
		payload    map[string]interface{}
		wantStatus int
	}{
		{
			name: "Login válido",
			payload: map[string]interface{}{
				"email":    "joao@test.com",
				"password": "123456",
			},
			wantStatus: http.StatusOK,
		},
		{
			name: "Email inválido",
			payload: map[string]interface{}{
				"email":    "email-invalido",
				"password": "123456",
			},
			wantStatus: http.StatusBadRequest,
		},
		{
			name: "Senha vazia",
			payload: map[string]interface{}{
				"email":    "joao@test.com",
				"password": "",
			},
			wantStatus: http.StatusBadRequest,
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			payloadBytes, _ := json.Marshal(tt.payload)
			req, _ := http.NewRequest("POST", "/auth/login", bytes.NewBuffer(payloadBytes))
			req.Header.Set("Content-Type", "application/json")
			
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			
			assert.Equal(t, tt.wantStatus, w.Code)
		})
	}
} 