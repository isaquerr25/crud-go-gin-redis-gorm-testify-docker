package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	// Configurar o ambiente de teste (por exemplo, um banco de dados de teste separado)
	// Certifique-se de restaurar o ambiente ao seu estado original após os testes

	// Criar um servidor Gin com as mesmas rotas
	router := gin.Default()
	router.POST("/users", CreateUser)

	// Criar um request de teste para a rota de criação de usuário
	reqBody := strings.NewReader(`{"name": "John Doe", "email": "john@example.com"}`)
	req, err := http.NewRequest("POST", "/users", reqBody)
	assert.NoError(t, err)

	// Simular uma solicitação HTTP
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Verificar se a resposta HTTP é 200 OK
	assert.Equal(t, http.StatusOK, w.Code)

	// Verificar se o corpo da resposta contém os dados do usuário criado
	assert.Contains(t, w.Body.String(), "John Doe")
	assert.Contains(t, w.Body.String(), "john@example.com")

	// Adicionar mais verificações conforme necessário

	// Limpar o banco de dados após os testes
	// Substitua essa parte com a lógica específica do seu banco de dados
	// Por exemplo, para GORM, você pode usar db.DropTableIfExists(&User{})
	// Certifique-se de que essa lógica esteja adaptada para o seu caso específico
	// (não use em ambientes de produção sem cuidado)
}

func TestGetUsers(t *testing.T) {
	// Configurar o ambiente de teste (por exemplo, um banco de dados de teste separado)

	// Criar um servidor Gin com as mesmas rotas
	router := gin.Default()
	router.GET("/users/:id", GetUsers)

	// Criar um request de teste para a rota de obtenção de usuários
	req, err := http.NewRequest("GET", "/users/1", nil)
	assert.NoError(t, err)

	// Simular uma solicitação HTTP
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Verificar se a resposta HTTP é 200 OK
	assert.Equal(t, http.StatusOK, w.Code)

	// Adicionar mais verificações conforme necessário
}
