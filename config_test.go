package main

import "testing"

func TestPostgresConfig_ConnectionString(t *testing.T) {
	pc := PostgresConfig{
		Host:         "localhost",
		Port:         "5432",
		User:         "pg",
		Password:     "pg",
		DatabaseName: "pg",
	}

	got := pc.ConnectionString()
	expected := "host=localhost port=5432 user=pg password=pg dbname=pg"

	if got != expected {
		t.Errorf("Invalid ConnectionString result: \n got: \t\t \"%s\" \n expected: \t \"%s\"", got, expected)
	}
}
