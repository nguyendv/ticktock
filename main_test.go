package main

import (
    "testing"
)

// Unittest for parseConfig()
func TestParseConfig(t *testing.T) {
    var valid_data = `
tick: a
tock: b
bong: c
`
    config := parseConfig([]byte(valid_data))
    if config.Tick != "a" || config.Tock != "b" || config.Bong != "c" {
        t.Errorf("It should be able to parse a valid config")
    }

    var invalid_data = `
tic: a
tk: b
bon: c
`
    config = parseConfig([]byte(invalid_data))
    if config.Tick != "tick" || config.Tock != "tock" || config.Bong != "bong" {
        t.Errorf("If config data is invalid, use defaults")
        t.Errorf("Got %s,%s,%s. Expect tick,tock,bong", config.Tick, config.Tock, config.Bong)
    }
}
