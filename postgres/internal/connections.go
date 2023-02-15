package handler

import (
	"database/sql"
	"log"
	"fmt"
	"context"
	"strings"
	"net/url"
	"github.com/docker-app/postgres-app/pgconf"
	_ "github.com/lib/pq"
	
)

func GetNewPgHandler(remoteHost string) (*sql.DB, error) {

	log.Println("Registered drivers: ", sql.Drivers())

	pg, err := sql.Open(pgconf.Driver, remoteHost)
	if err != nil {
		//log.Println("Actual error found here")
		log.Println(err)
		return nil, err
	}

	return pg, nil
}

func ClosePgDb(pg *sql.DB) {

	err := pg.Close()
	if err != nil {
		log.Println(err)		
	}
	
}

func IsDBReachable(pg *sql.DB) bool {

	err := pg.Ping()
	if err != nil {
		log.Println(err)	
		return false
	}

	return true
}

func GetNewPgConn(ctx context.Context, pg *sql.DB) (*sql.Conn, error) {

	pgConn, err := pg.Conn(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	log.Print("New connection is established.")
	return pgConn, nil
}

func ClosePgConn(ctx context.Context, pgConn *sql.Conn) error {

	err := pgConn.Close()
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func IsConnAlive(ctx context.Context, pgConn *sql.Conn) bool {

	err := pgConn.PingContext(ctx)
	if err != nil {
		log.Println(err)
		return false
	}
	
	log.Println("Connection is still active.")

	return true
}

// Inser into database
func Insert(ctx context.Context, pgConn *sql.Conn, url *url.URL) {

	
	//insert := fmt.Sprintf("INSERT INTO laptop (%v) VALUES")


}

// Read from database
func Read(ctx context.Context, pgConn *sql.Conn, url *url.URL) string {
	
	whereCond := GetReadQryCond(ctx, url)
	if whereCond == "" {
		log.Println("READ: WHERE clause is empty. Fetching all rows")		
	}

	qry := fmt.Sprintf("SELECT app_name FROM apps %v", whereCond)
	log.Printf("Executing query [%v]", qry)
	
	
	rows, err := pgConn.QueryContext(ctx, qry)
	if err != nil {
		log.Panic(err)
	}

	defer rows.Close()
	
	var result []string	
	for rows.Next() {
		var appName string
		if err := rows.Scan(&appName); err != nil {
			log.Println(err)
			continue
		}		
		result = append(result, appName)
	}
	
	resultStr := strings.Join(result, " \n")
	resultStr = resultStr + "\n"
	log.Printf("Query result: [%v]\n", resultStr)
	
	return resultStr
	
}

// Update row/rows record in database
func Update(ctx context.Context, pgConn *sql.Conn, url *url.URL) {


}

// Delete from database
func Delete(ctx context.Context, pgConn *sql.Conn, url *url.URL) {


}


