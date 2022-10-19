package main

import (
	// "encoding/json"
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/goccy/go-json"
)

func updateScript() {

	var users Users

	db := Database{}.InitDatabase()

	defer db.Close()

	jsonfile := "random_users.json"
	jsondata, err := ioutil.ReadFile(jsonfile)

	if err != nil {

		CheckErr(err)

	}

	err = json.Unmarshal(jsondata, &users)

	if err != nil {
		CheckErr(err)
	}

	lines, err := readLines("usernames_unique.txt")
	log.Printf("starting....")
	if err != nil {
		CheckErr(err)
	}
	// 	UPDATE table_name
	// SET column1 = value1, column2 = value2, ...
	// WHERE condition;
	statement, err := db.Prepare("update users set username = ? where id = ?")

	if err != nil {
		CheckErr(err)
	}
	defer statement.Close()
	log.Printf("starting....")
	for key, user := range users {

		_, err = statement.Exec(lines[key], key)
		if err != nil {
			CheckErr(err)
			continue
		}
		fmt.Println(key, ": Sucessfully Updated Username from", user.Login.Username, "to", lines[key])
		// fmt.Printf("user%v: = %s \n", key, user.Name.First)
	}

}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
