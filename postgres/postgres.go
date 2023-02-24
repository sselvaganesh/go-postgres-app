package pg

import (
	"fmt"
	"log"
	"context"
	"time"
	"os"
	"database/sql"
	"net/http"
	"github.com/docker-app/postgres-app/postgres/ty"
	"github.com/docker-app/postgres-app/pgconf"
	"github.com/docker-app/postgres-app/postgres/internal"
)

func GetInitCfg() t.PgCfg {
		
	cfg := t.PgCfg {
		SslModule: 	pgconf.SslModule,
		HostAddr:	pgconf.HostAddrLocal,
		Port:		pgconf.Port,
		User:		pgconf.User,
		Passwd:		pgconf.Passwd,
		Schema:		pgconf.Schema,
	}
	
	return cfg
}

func checkPgConnectionPeriodically(ctx context.Context, pgConn *sql.Conn) {

	for {
		time.Sleep(15 * time.Second)
		isAlive := handler.IsConnAlive(ctx, pgConn)
		
		if !isAlive {
			log.Panic()
		}
	
	}

}

func getRemoteHost() string {

	var pgHostAddr string

	if os.Getenv("RUN_LOCAL") == "TRUE" {
		pgHostAddr = pgconf.HostAddrLocal
	} else {
		pgHostAddr = pgconf.HostAddrDocker
	}

	log.Println("PgHostAddr: ", pgHostAddr)
	//postgres://postgres:password@localhost/DB_1?sslmode=disable
	remoteHost := fmt.Sprintf("%v://%v:%v@%v:%v/%v?sslmode=%v", pgconf.Schema, pgconf.User, pgconf.Passwd, pgHostAddr, pgconf.Port, pgconf.DbName, pgconf.SslModule)
	log.Printf("Remote Host: [%v]\n", remoteHost)
	
	return remoteHost
}

func GetPgListener(ctx context.Context) *sql.Conn {

	remoteHost := getRemoteHost()
	if remoteHost == "" {
		log.Println("Remote host not found")
		return nil
	}


	pgDB, err := handler.GetNewPgHandler(remoteHost)
	if err != nil {		
		return nil
	}
	
	//defer handler.ClosePgDb(pgDB)
	
	pgConn, err := handler.GetNewPgConn(ctx, pgDB)
	if err != nil {
		return nil
	}
	
	go checkPgConnectionPeriodically(ctx, pgConn)
	
	return pgConn
	
}

func Handler(ctx context.Context, w http.ResponseWriter, req *http.Request, reqType string, pgConn *sql.Conn) string {
	
	switch reqType {
		case "READ":
			result := handler.Read(ctx, pgConn, req.URL)
			return result		
		default:
			log.Println("Method definition not found.")
	}
	
	return ""
}






