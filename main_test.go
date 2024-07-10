package main
import (
  "os"
  "testing"
)

func GetEnv(t *testing.T) {
  expectedPort := "8080"
  os.Setenv("PORT", expectedPort)

  port := os.Getenv("PORT")

  if port != expectedPort {
    t.Errorf("Expected PORT to be %s, but got %s", expectedPort, port)
  }
}
