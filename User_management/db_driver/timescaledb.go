package db_driver

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	//"github.com/pelletier/go-toml/query"
)

type Timescaledb_conn struct {
	Tdb   *pgxpool.Pool
	Table pgconn.CommandTag
}

func (tscdb *Timescaledb_conn) Timescale_conn(config map[string]string) error {
	ctx := context.Background()

	connStr := fmt.Sprintf("%s://%s:postgres@%s:%s/%s", config["user_name"], config["db_password"], config["db_host"], config["db_port"], config["db_name"])
	dbpool, err := pgxpool.New(ctx, connStr)
	//tsc_db := Timescaledb_conn{}
	tscdb.Tdb = dbpool

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	//defer dbpool.Close()

	//run a simple query to check our connection
	var greeting string
	err = dbpool.QueryRow(ctx, "select 'Hello, Timescale (but concurrently)'").Scan(&greeting)
	if err != nil {
		//log.Printf(os.Stderr, "QueryRow failed: %v\n", err)
		//log.Println(err)
		return err
	}
	log.Println(greeting)
	return nil
}

func (db_instance *Timescaledb_conn) Timescale_create_ts_table(queryCreateTable string, queryCreateHypertable string) error {
	/********************************************/
	/* Create Hypertable                        */
	/********************************************/
	// Create hypertable of time-series data called sensor_data
	/*queryCreateTable := `CREATE TABLE sensor_data (
	        time TIMESTAMPTZ NOT NULL,
	        sensor_id INTEGER,
	        temperature DOUBLE PRECISION,
	        cpu DOUBLE PRECISION,
	        FOREIGN KEY (sensor_id) REFERENCES sensors (id));
	        `
	queryCreateHypertable := `SELECT create_hypertable('sensor_data', 'time');`
	*/
	//execute statement
	_, err := db_instance.Tdb.Exec(context.Background(), queryCreateTable+queryCreateHypertable)
	if err != nil {
		//log.Fatal("Unable to create hypertable the table exists already with the database")
		return err
	}
	//tscdb.Table = hyper_table
	//tscdb.Table.Insert()
	log.Println("Successfully created hypertable ")
	return nil
}

func (tscdb *Timescaledb_conn) Timescale_create_relational_table(queryCreateTable string) error {
	//queryCreateTable := `CREATE TABLE sensors (id SERIAL PRIMARY KEY, type VARCHAR(50), location VARCHAR(50));`
	relational_table, err := tscdb.Tdb.Exec(context.Background(), queryCreateTable)
	if err != nil {
		return err
	}
	tscdb.Table = relational_table
	log.Println("Successfully created relational table SENSORS")
	return err
}

func (tscdb *Timescaledb_conn) Timescale_drop_db(database_name string) error {
	query := fmt.Sprintf("DROP DATABASE %s", database_name)
	_, err := tscdb.Tdb.Exec(context.Background(), query)
	if err != nil {
		return err
	}
	log.Println("Database deleted successfully!")
	return err
}
func (tscdb *Timescaledb_conn) Timescale_drop_table(table_name string) error {
	query := fmt.Sprintf("DROP TABLE %s;", table_name)
	_, err := tscdb.Tdb.Exec(context.Background(), query)
	if err != nil {
		//log.Fatal(err)
		return err
	}
	log.Println("Table deleted successfully!")
	return err
}

func (tscdb *Timescaledb_conn) Timescale_write_data(query_data string) error {
	/* INSERT into  relational table            */
	/********************************************/
	//Insert data into relational table

	// Slices of sample data to insert
	// observation i has type sensorTypes[i] and location sensorLocations[i]
	//sensorTypes := []string{"a", "a", "b", "b"}
	//sensorLocations := []string{"floor", "ceiling", "floor", "ceiling"}

	//INSERT statement in SQL
	//queryInsertMetadata := `INSERT INTO sensors (type, location) VALUES ($1, $2);`
	//q := fmt.Sprintf("INSERT INTO %s %f (type, location) VALUES ($1, $2);", "sensors",45.6)
	//Execute INSERT command
	_, err := tscdb.Tdb.Exec(context.Background(), query_data)
	if err != nil {
		//log.Println("Unable to insert into database")
		return err
	}
	//fmt.Printf("Inserted sensor (%s, %s) into database \n", sensorTypes[i], sensorLocations[i])

	//fmt.Println("Successfully inserted all sensors into database")
	return nil
}
func (tscdb *Timescaledb_conn) Timescale_write_many_data(batch_insert []string) error {
	for k, row_data := range batch_insert {
		_, err := tscdb.Tdb.Exec(context.Background(), row_data)
		if err != nil {
			//fmt.Fprintf(os.Stderr, "Unable to insert data into database: %v\n", err)
			log.Println("Unable to insert data into database")
			return err
		}
		//fmt.Printf("Inserted sensor (%s, %s) into database \n", sensorTypes[i], sensorLocations[i])
		log.Println("Inserted data ", k)

	}
	log.Println("Successfully inserted all data into database")
	return nil
}
func (tscdb *Timescaledb_conn) Timescale_read_data(query_data string) ([]byte, error) {
	var jsonData []byte
	err := tscdb.Tdb.QueryRow(context.Background(), query_data).Scan(&jsonData)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %w", err)
	}

	//fmt.Println(string(jsonData))
	return jsonData, nil

	//var x []byte
	//fmt.Println(string(row.Scan(&x)))

}

func (tscdb *Timescaledb_conn) Timescale_update_data(query_data string) error {
	_, err := tscdb.Tdb.Exec(context.Background(), query_data)
	if err != nil {
		//log.Println("Unable to insert into database")
		return err
	}

	return nil
}

func (tscdb *Timescaledb_conn) Timescale_remove_data(query_data string) error {
	_, err := tscdb.Tdb.Exec(context.Background(), query_data)
	if err != nil {
		log.Println("unable to remove the data")
		return err
	}
	//fmt.Println("removed successfully")
	return nil
}

func (tscdb *Timescaledb_conn) Timescale_list_data(query_data string) ([]byte, error) {
	rows, err := tscdb.Tdb.Query(context.Background(), query_data)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %w", err)
	}
	defer rows.Close()

	var results []json.RawMessage

	for rows.Next() {
		var jsonData json.RawMessage
		err := rows.Scan(&jsonData)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}
		results = append(results, jsonData)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %w", err)
	}

	jsonData, err := json.Marshal(results)
	if err != nil {
		return nil, fmt.Errorf("error marshaling results to JSON: %w", err)
	}

	return jsonData, nil
}

func kmain() {
	ctx := context.Background()
	connStr := "postgres://postgres:postgres@localhost:5432/test_db"
	dbpool, err := pgxpool.New(ctx, connStr)
	tsc_db := Timescaledb_conn{}
	tsc_db.Tdb = dbpool
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	//run a simple query to check our connection
	var greeting string
	err = dbpool.QueryRow(ctx, "select 'Hello, Timescale (but concurrently)'").Scan(&greeting)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(greeting)
	//tsc_db.TimescaleH_write_data()
	//check_table := Timescaledb_conn{}
	//tsc_db.Timescale_create_relational_table()
	queryCreateTable := `CREATE TABLE iot_data (
		time TIMESTAMPTZ NOT NULL,
		sensor_id INTEGER,
		temperature DOUBLE PRECISION,
		cpu DOUBLE PRECISION,
		FOREIGN KEY (sensor_id) REFERENCES sensors (id));
		`
	queryCreateHypertable := `SELECT create_hypertable('iot_data', 'time');`
	tsc_db.Timescale_create_ts_table(queryCreateTable, queryCreateHypertable)
}
