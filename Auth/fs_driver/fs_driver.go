package fs_driver

import (
	"fmt"
	"log"

	"tele_auth/db_driver"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Db_client_object struct {
	db_driver.Timescaledb_conn
	Client_tsdb *pgxpool.Pool
	Data        *pgx.Rows
}

type Dot_channel_model struct {
	Owner_id   string `json:"owner_id"`
	Thing_id   string `json:"thing_id"`
	Thing_key  string `json:"thing_key"`
	Channel_id string `json:"channel_id"`
}

func (open *Db_client_object) Fs_driver_open(config map[string]string) {

	err := open.Timescale_conn(config)
	if err != nil {
		log.Println(err)
	}
	open.Client_tsdb = open.Tdb
	log.Println("Successful connection to Time-scaleDB!")
}
func (rel_table *Db_client_object) Fs_driver_create_relational_table(query_create_table string) error {

	err := rel_table.Timescale_create_relational_table(query_create_table)
	if err != nil {
		log.Fatal(err)
		return err
	}
	log.Println("Table created successfully.")
	return nil
}

func (hyper_table *Db_client_object) Fs_driver_create_hyper_table(queryCreateTable string, queryCreateHypertable string) error{
	/*
	queryCreateTable := `CREATE TABLE Telemetry_data (
		time TIMESTAMPTZ NOT NULL,
		thing_id string,
		tempera DOUBLE PRECISION,
		cpu DOUBLE PRECISION,
		FOREIGN KEY (sensor_id) REFERENCES sensors (id));
		`
	queryCreateHypertable := `SELECT create_hypertable('vehicle_data', 'time');`
	*/

	//db_driver.Timescaledb_conn.Timescale_create_ts_table()
	err := hyper_table.Timescale_create_ts_table(queryCreateTable, queryCreateHypertable)
	if err != nil {
		return err
	}
	fmt.Println("Hyper table created successfully!")
	return nil

}

func (rel_table *Db_client_object) Delete_relational_table() {}

func (rel_table *Db_client_object) Delete_hyper_table() {
	err := rel_table.Timescale_drop_table("vehicle_data")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Table successfully deleted")
}

// CRUD operations to perform create, read, update and delete specific data from the table
func (hyper_table *Db_client_object) Fs_driver_write(query string) error {
	//table_name := "vehicle_data"

	err := hyper_table.Timescale_write_data(query)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("Data written successfully to database.")
	return nil

}
func (hyper_table *Db_client_object) Fs_driver_write_batch() error {
	query := fmt.Sprintf("INSERT INTO %s %f (owner_id, thing_id, thin) VALUES ($1, $2);", "Dot_channel", 45.6)
	fmt.Println(query)
	err := hyper_table.Timescale_write_many_data([]string{"1"})
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("Data batch written successfully")
	return nil
}

func (hyper_table *Db_client_object) Fs_driver_read(q string) ([]byte, error){
	//q := `SELECT * FROM Dot_channel WHERE thing_id ='a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11';`
	//fmt.Println(q)
	//var data_model Dot_channel_model

	data, err := hyper_table.Timescale_read_data(q)
	if err!= nil{
		return nil, err
	}
	return data, nil
	// if err != nil {
	//log.Println(err)
	// return data, err
	// }

	//data.FieldDescriptions()
	//hyper_table.Data = &data
	/*
		for data.Next() {
			var (
				thingID   string
				ownerID   string
				thingKey  string
				channelID string
			)

			err := data.Scan(&thingID, &ownerID, &thingKey, &channelID)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("Thing ID: %s, Owner ID: %s, Thing Key: %s, Channel ID: %s\n", thingID, ownerID, thingKey, channelID)
		}

	*/

}
func (hyper_table *Db_client_object) Fs_driver_update(query string) error{
	err := hyper_table.Timescale_update_data(query)
	if err != nil {
		log.Println(err)
		return err
	}
	fmt.Println("updated successfully")
	return nil
}
func (hyper_table *Db_client_object) Fs_driver_delete(query string) error{
	//query := `DELETE FROM Dot_channel WHERE thing_id='a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11'`
	err := hyper_table.Timescale_remove_data(query)
	if err != nil {
		log.Println(err)
		return err
	}
	fmt.Println("data removed successfully")
	return nil

}
func (hyper_table *Db_client_object) Fs_driver_list(query string) ([]byte,error){
	//query := "SELECT json_agg(row_to_json(Dot_channel)) FROM Dot_channel"
	data,err := hyper_table.Timescale_list_data(query)
	if err!= nil{
		//log.Println(err)
		return nil,err
	}
	return data, nil
	/*
		var result []Dot_channel_model
		for all_data.Next() {
			var dot Dot_channel_model
			err := all_data.Scan(&dot.Thing_id, &dot.Owner_id, &dot.Thing_key, &dot.Channel_id)
			if err != nil {
				log.Println("unable to read/scan the data from database")
			}
			fmt.Printf("Thing ID: %s, Owner ID: %s, Thing Key: %s, Channel ID: %s\n", dot.Thing_id, dot.Owner_id, dot.Thing_key, dot.Channel_id)
			result = append(result, dot)
		}
		for _, res := range result {
			fmt.Println(res.Thing_id)
		}
	*/
}

func jmain() {
	db := Db_client_object{}
	config := map[string]string{"user_name": "postgres", "db_password": "postgres", "db_host": "localhost", "db_port": "5432", "db_name": "timeseries"}

	db.Fs_driver_open(config)
	//db.Fs_driver_create_hyper_table()
	//db.Delete_hyper_table()
	//db.Fs_driver_create_relational_table()

	//db.Fs_driver_write()
	//db.Fs_driver_update()
	//q := `SELECT row_to_json(Dot_channel) FROM Dot_channel WHERE thing_id ='p0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11';`
	//db.Fs_driver_read(q)
	//db.Fs_driver_delete()

	//db.Fs_driver_list()
}