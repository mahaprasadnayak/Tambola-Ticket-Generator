package global

import (
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
)

var (
	Port, Db_Creds string
	PostgresConn   *pgxpool.Pool
)

func init() {
	loadErr := godotenv.Load("./app/config.env")
	if loadErr != nil {
		fmt.Println("Error loading.env file ", loadErr)
		log.Fatal("Error loading.env file ", loadErr)
	}
	Port = os.ExpandEnv("$application_port")
	Db_Creds = os.ExpandEnv("$db_creds")

	//postgres connection

	// config, err := pgxpool.ParseConfig(Db_Creds + "?")
	// if err != nil {
	// 	fmt.Println("Error parsing connection string ", err)
	// 	log.Fatal("Error parsing connection string ", err)
	// }
	// PostgresConn, err = pgxpool.ConnectConfig(context.Background(), config)
	// if err != nil {
	// 	fmt.Println("Error connecting to database ", err)
	// 	log.Fatal("Error connecting to database ", err)
	// }
}
