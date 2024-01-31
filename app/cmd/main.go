package main

import (
	"context"
	"github.com/poster-keisuke/sample-clearn-architecture/app/server"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	server.Run(ctx)

	//conf := config.GetConfig()
	//dbClinet, err := db.NewDB(conf.DB)
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "[ERROR] Failed to setup DB: %s\n", err)
	//	panic(err)
	//}

}
