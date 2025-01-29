package database

// import (
// 	"database/sql"
// 	"log"
// )

// func InitDB() *sql.DB {
// 	var err error
// 	dsn := "postgres://admin2:12345@postgres-db:5432/manage_rooms?sslmode=disable"

// 	db, err := sql.Open("pgx", dsn)
// 	if err != nil {
// 		log.Fatalf("Failed to connect to the database: %v", err)
// 	}

// 	_, err = db.Exec(`
//         CREATE TABLE IF NOT EXISTS rooms (
//             code INTEGER PRIMARY KEY,
//             is_available BOOLEAN
//         );
//     `)
// 	if err != nil {
// 		log.Fatalf("Failed to create table: %v", err)
// 	}

// 	// Check if the rooms table is empty
// 	var rowCount int
// 	err = db.QueryRow("SELECT COUNT(*) FROM rooms").Scan(&rowCount)
// 	if err != nil {
// 		log.Fatalf("Failed to check row count: %v", err)
// 	}

// 	// If the table is empty, insert rows
// 	if rowCount == 0 {
// 		_, err = db.Exec(`
// 			INSERT INTO rooms (code, is_available)
// 			SELECT generate_series(0, 1679615), true;
// 		`)
// 		if err != nil {
// 			log.Fatalf("Failed to insert room data: %v", err)
// 		}
// 	}

// 	return db
// }
