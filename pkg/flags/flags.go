package flags

import (
	"flag"
	"log"
	"os"
	"strconv"
)

//flags for
type Flags struct {
	PostgresHost  string
	PostgresUser  string
	PostgresPass  string
	PostgresPort  int
	AlgoliaKey    string
	AlgoliaSecret string
	ServerPort    int
}

func ParseFlags() Flags {
	var f Flags
	flag.IntVar(&f.ServerPort, "port", 8080, "what port the server listens on")
	//postgres
	flag.StringVar(&f.PostgresHost, "phost", "127.0.0.1", "postgres host")
	flag.StringVar(&f.PostgresUser, "puser", "postgres", "postgres user")
	flag.StringVar(&f.PostgresPass, "ppass", "postgres", "postgres password")
	flag.IntVar(&f.PostgresPort, "pport", 5432, "postgres port")
	//algolia
	flag.StringVar(&f.AlgoliaKey, "algoliakey", "", "algolia key")
	flag.StringVar(&f.AlgoliaSecret, "algoliasecret", "", "algolia secret")
	//parse from envrionment?
	env := flag.Bool("env", false, "if true, parses flags from environment(besides this, obviously). Uses the names of the other options, e.g.: -phost is read from PHOST")

	flag.Parse()
	if *env {
		return parseFromEnv()
	}
	return f
}

func parseFromEnv() Flags {
	return Flags{
		PostgresHost:  getFromEnv("PHOST", "127.0.0.1"),
		PostgresUser:  getFromEnv("PUSER", "postgres"),
		PostgresPass:  getFromEnv("PPASS", "postgres"),
		PostgresPort:  getIntFromEnv("PPORT", 5432),
		AlgoliaKey:    getFromEnv("ALGOLIAKEY", ""),
		AlgoliaSecret: getFromEnv("ALGOLIASECRET", ""),
	}
}

//parse flags from environment
func getFromEnv(v string, defaultValue string) string {
	s := os.Getenv(v)
	if s == "" {
		return defaultValue
	}
	return s
}

func getIntFromEnv(v string, defaultValue int) int {
	s := os.Getenv(v)
	if s == "" {
		return defaultValue
	}
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("Error, could not parse %s: %v", v, err)
	}
	return i
}
