package qry

/*

	1. Read
		localhost:<Port>/read?<db-field-name>=<value>

	2. Insert
		localhost:<Port>/insert?<db-field-name>=<value>&<db-field-name>=<value>&<db-field-name>=<value>...
		
	3. Update
		localhost:<Port>/update?<select-db-field-name>=<value>&<select-db-field-name>=<value>..&<update-db-field-name>=<value>&<update-db-field-name>=<value>
		
	4. Delete
		localhost:<Port>/delete?<db-field-name>=<value>
		
*/


const ReadQryStr = `SELECT * FROM laptop WHERE `
const InsertQryStr = `INSERT INTO laptop () `
const UpdateQryStr = ``
const DeleteQryStr = ``


