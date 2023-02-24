package main


import (
	"io"
	"log"
	"net/http"
	"context"
	"database/sql"
	"github.com/docker-app/postgres-app/postgres"
)

type postgres struct {
	ctx context.Context
	pgConn *sql.Conn
}

func(p *postgres) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	log.Print("Reaching this line.?")
	result := pg.Handler(p.ctx, w, r, "READ", p.pgConn )
	io.WriteString(w, result)
	
}

func main() {

	log.Println("	Main function	")

	ctx := context.TODO()	
	p := &postgres{
		ctx: ctx,
		pgConn: pg.GetPgListener(ctx),
	}

	// Create server
	svr := &http.Server {
		Addr: "0.0.0.0:8080",
		}
				
	log.Println("Server is listening in port 8080: ", svr)				
							
	http.Handle("/read", p)	
	if err := svr.ListenAndServe()	; err != nil {
		log.Fatal("ListenAndServer: ", err)
	}
}
