package config

import (
	"os"
	"path/filepath"
	"testing"
)

func withConfigPath(t *testing.T) string {
	t.Helper()
	original := yamlFile
	yamlFile = filepath.Join(t.TempDir(), "server.yaml")
	t.Cleanup(func() { yamlFile = original })
	return yamlFile
}

func TestGetServerConfigCreatesPrivateConfig(t *testing.T) {
	configPath := withConfigPath(t)

	cfg := GetServerConfig()
	if len(cfg.SignKey) != 32 {
		t.Fatalf("generated sign key length = %d, want 32", len(cfg.SignKey))
	}
	if cfg.HttpConfig.Port != "127.0.0.1:8080" {
		t.Fatalf("default listener = %q, want loopback", cfg.HttpConfig.Port)
	}
	if cfg.HttpConfig.StaticDir != "/app/dist" {
		t.Fatalf("default static directory = %q, want /app/dist", cfg.HttpConfig.StaticDir)
	}

	info, err := os.Stat(configPath)
	if err != nil {
		t.Fatal(err)
	}
	if got := info.Mode().Perm(); got != 0600 {
		t.Fatalf("config permissions = %04o, want 0600", got)
	}
}

func TestGetServerConfigRejectsMalformedConfig(t *testing.T) {
	configPath := withConfigPath(t)
	if err := os.WriteFile(configPath, []byte("db: ["), 0600); err != nil {
		t.Fatal(err)
	}

	defer func() {
		if recover() == nil {
			t.Fatal("GetServerConfig accepted malformed configuration")
		}
	}()
	GetServerConfig()
}
