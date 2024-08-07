package fs

//package main
import (
	"log"

	"user_manager/fs_driver"
)

//	type db_cli_manager struct {
//		*fs_driver.Db_client_object
//	}
//
// var db_cli_manager *fs_driver.Db_client_objectvar
var Destro = fs_driver.Db_client_object{}

func Fs_open() {
	config := map[string]string{
		"user_name": "postgres", 
		"db_password": "postgres", 
		"db_host": "localhost", 
		"db_port": "5432", 
		"db_name": "task_app"}
	Destro.Fs_driver_open(config)

}
func Fs_create_relational_table(queryCreateTable string) error {

	/*
	   queryCreateTable_2 := `CREATE TABLE Service_identity(

	   	OWNER_ID   VARCHAR(36) NOT NULL,
	   	THING_ID   VARCHAR(36) NOT NULL,
	   	THING_KEY  VARCHAR(36) NOT NULL,
	   	PRIMARY KEY (OWNER_ID)

	   );`
	*/
	err := Destro.Fs_driver_create_relational_table(queryCreateTable)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func Fs_create_time_series_table(queryCreatetable string, queryCreateHyperTable string) error {
	err := Destro.Fs_driver_create_hyper_table(queryCreatetable, queryCreateHyperTable)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func Fs_write(query string) error {
	err := Destro.Fs_driver_write(query)
	if err != nil {
		//log.Fatal(err)
		return err
	}
	return nil

}
func Fs_read(query string) ([]byte, error) {
	//q := `SELECT row_to_json(Dot_channel) FROM Dot_channel WHERE thing_id ='p0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11';`
	data, err := Destro.Fs_driver_read(query)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func Fs_update(query_cmd string) error {
	//query := `UPDATE Dot_channel SET owner_id = 'g0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11' WHERE thing_id = 'a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11'`

	err := Destro.Fs_driver_update(query_cmd)
	if err != nil {
		return err
	}
	//fmt.Printf("Object  updated: %v\n", update_object)
	return nil

}
func Fs_delete(query_cmd string) error {

	err := Destro.Fs_driver_delete(query_cmd)
	if err != nil {
		return err
	}
	return nil
}
func Fs_list(query_cmd string) ([]byte, error) {
	//query := "SELECT json_agg(row_to_json(Dot_channel)) FROM Dot_channel"
	fetched, err := Destro.Fs_driver_list(query_cmd)
	if err != nil {
		//log.Fatal(err)
		return nil, err
	}

	return fetched, err
}
func main() {
}
