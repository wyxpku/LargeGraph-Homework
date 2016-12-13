package models

import (
	"errors"
	"fmt"

	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
)

var driver bolt.Driver

func init() {
	driver = bolt.NewDriver()
}

func AddUser(email, name, password string) error {
	conn, err := driver.OpenNeo("bolt://neo4j:610@localhost:7687")
	if err != nil {
		return err
	}
	defer conn.Close()

	if row, err := conn.QueryNeo("MATCH (u:User{email:{email}}) RETURN ID(u)", map[string]interface{}{"email": email}); err != nil {
		return err
	} else if ids, _, _ := row.NextNeo(); len(ids) > 0 {
		return errors.New("Duplicate Email")
	} else {
		row.Close()
	}

	if result, err := conn.ExecNeo("CREATE (u:User {email:{email}, name:{name}, password:{password}}) RETURN ID(n)", map[string]interface{}{
		"email":    email,
		"name":     name,
		"password": password,
	}); err != nil {
		fmt.Println(result.LastInsertId())
		return err
	} else {
		return nil
	}
}

func GetUserId(email, password string) (id int64, err error) {
	conn, err := driver.OpenNeo("bolt://neo4j:610@localhost:7687")
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	if row, err := conn.QueryNeo("MATCH (u:User{email:{email}, password:{password}}) RETURN ID(u)", map[string]interface{}{
		"email":    email,
		"password": password,
	}); err != nil {
		return 0, err
	} else if ids, _, _ := row.NextNeo(); len(ids) > 0 {
		return ids[0].(int64), nil
	} else {
		return 0, errors.New("email or password error")
	}
}

func GetUser(id int64) (user interface{}, err error) {
	conn, err := driver.OpenNeo("bolt://neo4j:610@localhost:7687")
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	if row, err := conn.QueryNeo("MATCH (u:User) WHERE ID(u)={id} RETURN u", map[string]interface{}{"id": id}); err != nil {
		return nil, err
	} else if ids, _, _ := row.NextNeo(); len(ids) == 0 {
		return nil, errors.New("no such user with id")
	} else {
		fmt.Println(ids[0])
		return ids[0], nil
	}
}

func GetUserAll() (users []interface{}, err error) {
	conn, err := driver.OpenNeo("bolt://neo4j:610@localhost:7687")
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	if data, _, _, err := conn.QueryNeoAll("MATCH (u:User) RETURN u", nil); err != nil {
		return nil, err
	} else {
		users = make([]interface{}, 0, len(data))
		for _, d := range data {
			users = append(users, d[0])
		}
		return users, nil
	}
}
