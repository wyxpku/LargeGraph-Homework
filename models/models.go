package models

import (
	"errors"
	"fmt"
	"time"

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

	now := time.Now().Format("2006-01-02 15:04:05")
	if res, err := conn.ExecNeo("CREATE (u:User {email:{email}, name:{name}, password:{password}, time:{now}}) RETURN ID(u)", map[string]interface{}{
		"email":    email,
		"name":     name,
		"password": password,
		"now":      now,
	}); err != nil {
		return err
	} else {
		fmt.Println(res)
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

func GetUserMoment(id int64) (moments []interface{}, err error) {
	conn, err := driver.OpenNeo("bolt://neo4j:610@localhost:7687")
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	if data, _, _, err := conn.QueryNeoAll("MATCH (u:User)-[:Post]->(m) WHERE ID(u)={id} RETURN m", map[string]interface{}{"id": id}); err != nil {
		return nil, err
	} else {
		moments = make([]interface{}, 0, len(data))
		for _, d := range data {
			moments = append(moments, d[0])
		}
		return moments, nil
	}
}

func AddUserMoment(id int64, content string) error {
	conn, err := driver.OpenNeo("bolt://neo4j:610@localhost:7687")
	if err != nil {
		return err
	}
	defer conn.Close()

	now := time.Now().Format("2006-01-02 15:04:05")
	if _, err := conn.ExecNeo("MATCH (u:User) WHERE ID(u)={id} CREATE (m:Moment {content:{content}, time:{time}}), (u)-[:Post]->(m)", map[string]interface{}{
		"id":      id,
		"content": content,
		"time":    now,
	}); err != nil {
		return err
	} else {
		return nil
	}
}

func UserFollow(userId, targetId int64) error {
	if userId == targetId {
		return errors.New("Can't Follow Yourself")
	}
	conn, err := driver.OpenNeo("bolt://neo4j:610@localhost:7687")
	if err != nil {
		return err
	}
	defer conn.Close()

	if _, err := conn.ExecNeo("MATCH (a:User),(b:User) WHERE ID(a)={id1} AND ID(b)={id2} CREATE UNIQUE (a)-[:Follow]->(b)", map[string]interface{}{
		"id1": userId,
		"id2": targetId,
	}); err != nil {
		return err
	} else {
		return nil
	}
}

func UserUnfollow(userId, targetId int64) error {
	if userId == targetId {
		return errors.New("Can't Unfollow Yourself")
	}
	conn, err := driver.OpenNeo("bolt://neo4j:610@localhost:7687")
	if err != nil {
		return err
	}
	defer conn.Close()

	if _, err := conn.ExecNeo("MATCH (a:User)-[r:Follow]->(b:User) WHERE ID(a)={id1} AND ID(b)={id2} DELETE r", map[string]interface{}{
		"id1": userId,
		"id2": targetId,
	}); err != nil {
		return err
	} else {
		return nil
	}
}

func GetUserFollowing(id int64) []interface{} {
	conn, err := driver.OpenNeo("bolt://neo4j:610@localhost:7687")
	if err != nil {
		return nil
	}
	defer conn.Close()

	if data, _, _, err := conn.QueryNeoAll("MATCH (u:User)-[:Follow]->(a) WHERE ID(u)={id} RETURN a", map[string]interface{}{"id": id}); err != nil {
		return nil
	} else {
		users := make([]interface{}, 0, len(data))
		for _, d := range data {
			users = append(users, d[0])
		}
		return users
	}
}

func GetUserFollowed(id int64) []interface{} {
	conn, err := driver.OpenNeo("bolt://neo4j:610@localhost:7687")
	if err != nil {
		return nil
	}
	defer conn.Close()

	if data, _, _, err := conn.QueryNeoAll("MATCH (a:User)-[:Follow]->(u) WHERE ID(u)={id} RETURN a", map[string]interface{}{"id": id}); err != nil {
		return nil
	} else {
		users := make([]interface{}, 0, len(data))
		for _, d := range data {
			users = append(users, d[0])
		}
		return users
	}
}
