package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/cucumber/godog"
	"net/http"
	"net/url"
	"testing"
)

type doroTestSuite struct {
	resp      *http.Response
	ctx       context.Context
	userEmail string
}

func (s *doroTestSuite) unUsuarioNoRegistradoConEmail(email string) error {
	s.ctx = context.WithValue(context.Background(), "email", email)
	return nil
}

func (s *doroTestSuite) envíoUnaPeticiónPOSTAUsersConLosDatosDelUsuario() error {
	client := &http.Client{}
	userData := map[string]string{
		"email":    s.ctx.Value("email").(string),
		"password": "testPassword123",
		"name":     "Test User",
	}
	jsonData, err := json.Marshal(userData)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(
		"POST",
		"http://localhost:8080/users",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	s.resp, err = client.Do(req)
	return err
}

func (s *doroTestSuite) dadoUnUsuarioRegistradoConEmail(email string) error {
	s.userEmail = email
	client := &http.Client{}
	userData := map[string]string{
		"email":    email,
		"password": "testPassword123",
		"name":     "Test User",
	}
	jsonData, err := json.Marshal(userData)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(
		"POST",
		"http://localhost:8080/users",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	_, err = client.Do(req)
	return err
}

func (s *doroTestSuite) cuandoBuscoElUsuarioPorEmail() error {
	client := &http.Client{}
	url := fmt.Sprintf("http://localhost:8080/users?email=%s", url.QueryEscape(s.userEmail))
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	s.resp, err = client.Do(req)
	return err
}

func (s *doroTestSuite) elSistemaDebeResponderConCódigo(codigo int) error {
	if s.resp.StatusCode != codigo {
		return fmt.Errorf("se esperaba código %d pero se obtuvo %d", codigo, s.resp.StatusCode)
	}
	return nil
}

func (s *doroTestSuite) realizoUnaPeticiónGETABacklogs() error {
	client := &http.Client{}
	req, err := http.NewRequest(
		"GET",
		"http://localhost:8080/backlogs",
		nil,
	)
	if err != nil {
		return err
	}
	req.Header.Set("X-User-Email", s.userEmail)
	s.resp, err = client.Do(req)
	return err
}

func (s *doroTestSuite) realizoUnaPeticiónGETATasks() error {
	client := &http.Client{}
	req, err := http.NewRequest(
		"GET",
		"http://localhost:8080/tasks",
		nil,
	)
	if err != nil {
		return err
	}
	req.Header.Set("X-User-Email", s.userEmail)
	s.resp, err = client.Do(req)
	return err
}

func (s *doroTestSuite) debeRetornarListadoDeTareas() error {
	var tasks []struct {
		ID        string `json:"id"`
		Title     string `json:"title"`
		Status    string `json:"status"`
		Priority  string `json:"priority"`
		DueDate   string `json:"due_date"`
		CreatedAt string `json:"created_at"`
	}
	if err := json.NewDecoder(s.resp.Body).Decode(&tasks); err != nil {
		return fmt.Errorf("error decodificando tareas: %v", err)
	}
	if len(tasks) < 0 {
		return fmt.Errorf("se esperaba al menos una tarea")
	}
	for _, task := range tasks {
		if task.ID == "" {
			return fmt.Errorf("tarea sin ID")
		}
		if task.Status == "" {
			return fmt.Errorf("tarea sin estado")
		}
		if task.Priority == "" {
			return fmt.Errorf("tarea sin prioridad")
		}
	}
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	suite := &doroTestSuite{}
	// Escenario 1: Crear usuario
	ctx.Given(`^un usuario no registrado con email "([^"]*)"$`, suite.unUsuarioNoRegistradoConEmail)
	ctx.When(`^envío una petición POST a "/users" con los datos del usuario$`, suite.envíoUnaPeticiónPOSTAUsersConLosDatosDelUsuario)
	// Escenario 2: Obtener usuario por email
	ctx.Given(`^dado un usuario registrado con email "([^"]*)"$`, suite.dadoUnUsuarioRegistradoConEmail)
	ctx.When(`^cuando busco el usuario por email$`, suite.cuandoBuscoElUsuarioPorEmail)
	// Escenario 3: Obtener Backlogs
	ctx.When(`^realizo una petición GET a "/backlogs"$`, suite.realizoUnaPeticiónGETABacklogs)
	// Escenario 4: Obtener Tasks
	ctx.When(`^realizo una petición GET a "/tasks"$`, suite.realizoUnaPeticiónGETATasks)
	// Step común
	ctx.Then(`^el sistema debe responder con código (\d+)$`, suite.elSistemaDebeResponderConCódigo)
}

func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features"},
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}
