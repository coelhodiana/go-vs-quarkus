package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

// TestHealthHandler testa o nosso endpoint /health.
func TestHealthHandler(t *testing.T) {
    // 1. Setup: Cria um pedido HTTP 'fake' para o nosso endpoint.
    req, err := http.NewRequest("GET", "/health", nil)
    if err != nil {
        // Se não conseguirmos sequer criar o pedido, o teste falha catastroficamente.
        t.Fatal(err)
    }

    // Cria um 'ResponseRecorder', que é um gravador de respostas.
    // Ele vai atuar como o nosso 'ResponseWriter' para que possamos inspecionar a resposta.
    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(healthHandler)

    // 2. Execução: Chama o nosso handler diretamente, passando o gravador e o pedido.
    handler.ServeHTTP(rr, req)

    // 3. Verificação (Assertions): Verificamos se o resultado foi o esperado.

    // VERIFICAÇÃO 1: O código de status HTTP deve ser 200 (OK).
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler retornou o status code errado: obteve %v, esperava %v",
            status, http.StatusOK)
    }

    // VERIFICAÇÃO 2: O corpo da resposta deve ser exatamente o que esperamos.
    expected := "Go service is up and running!"
    if rr.Body.String() != expected {
        t.Errorf("handler retornou o corpo errado: obteve '%v', esperava '%v'",
            rr.Body.String(), expected)
    }
}