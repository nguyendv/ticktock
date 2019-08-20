package main
import (
    "fmt"
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "time"
)

// configuration for the program
type Config struct {
    Tick string `yaml:"tick"`  // what to print on tick
    Tock string `yaml:"tock"`  // what to print on tock
    Bong string `yaml:"bong"`  // what to print on bong
}

func main() {
    secTicker := time.Tick(time.Second)

    // We use select in the main loop
    // If two tickers happen at the same time, select choose a random one
    // So 50 milliseconds here to make sure the secTicker and timeout won't
    // happen at the same time
    timeout := time.After(3 * time.Hour + 50 * time.Millisecond)

    nSecs := 0 // Counter for number of seconds
    minute := 60 // 60 seconds
    hour := 3600 // 3600 seconds

    for {
        configFile, err := ioutil.ReadFile("config.yaml")
        if err != nil {
            fmt.Println("Can't open config file")
        }
        config := parseConfig(configFile)
        select {
        case <- secTicker:
            nSecs += 1
            switch {
            case nSecs % hour == 0:
                fmt.Println(config.Bong)
            case nSecs % minute == 0:
                fmt.Println(config.Tock)
            default:
                fmt.Println(config.Tick)
            }
        case <- timeout:
            return
        default:
            time.Sleep(500 * time.Millisecond)
        }
    }
}

// parse yaml config
func parseConfig(bytes []byte) Config {
    var config Config
    err := yaml.Unmarshal(bytes, &config)

    if err != nil {
        fmt.Println("Failed to read config file")
        // use defaults
        config.Tick = "tick"
        config.Tock = "tock"
        config.Bong = "bong"
        fmt.Println(config)
    }

    //  use defaults when the field is not found
    if config.Tick == "" {
        config.Tick = "tick"
    }
    if config.Tock == "" {
        config.Tock = "tock"
    }
    if config.Bong == "" {
        config.Bong = "bong"
    }
    return config
}
