package handler

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestHandler_jsonError(t *testing.T) {
	msg := "Hello JSON"
	result := jsonError(msg)
	require.Equal(t, string([]byte(`{"Message":"Hello JSON"}`)), string(result))
}