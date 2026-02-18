package main

import(
	"log"
	"os"
)

func main() {
	cfg:= config{
		addr:":8080",
		db: dbConfig{},
	}

	api:= application{	
		config: cfg,
	}

	

	if err:= api.run(api.mount());err!=nil{
		log.Printf("server failed on mounting. err : %s",err)
		os.Exit(1)
	}

}