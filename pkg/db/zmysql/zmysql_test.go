package zmysql_test

import (
	"database/sql"
	"log"
	"os"
	"testing"
	"time"

	"github.com/zsp108/dbbmsql/pkg/db/zmysql"
)

var db *sql.DB

// TestMain 用于初始化全局资源（如数据库连接）并运行测试
func TestMain(m *testing.M) {
	// 初始化数据库连接
	ops := zmysql.Options{
		Host:                  "10.10.180.143",
		Username:              "root@mysqlt1",
		Password:              "baAA11__",
		Port:                  2881,
		Database:              "test",
		MaxIdleConnections:    10,
		MaxOpenConnections:    100,
		MaxConnectionLifeTime: 10 * time.Minute,
	}

	var err error
	db, err = zmysql.Connect(ops)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// 运行测试
	code := m.Run()

	// 清理资源
	db.Close()

	// 退出
	os.Exit(code)
}

// TestNew 测试数据库连接
func TestNew(t *testing.T) {
	if db == nil {
		t.Fatal("Database connection is not initialized")
	}

	// 简单检查数据库连接是否有效
	err := db.Ping()
	if err != nil {
		t.Errorf("Failed to ping database: %v", err)
	}

	t.Log("Database connection is healthy")
}

func TestCreateTable(t *testing.T) {
	if db == nil {
		t.Fatal("Database connection is not initialized")
	}

	// 创建测试表
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS test_table (id INT, name VARCHAR(255))")
	if err != nil {
		t.Errorf("Failed to create table: %v", err)
	}

	t.Log("Table created")
}

// TestSelect 测试查询功能
func TestSelect(t *testing.T) {
	if db == nil {
		t.Fatal("Database connection is not initialized")
	}

	// 使用子测试组织测试用例
	t.Run("SelectVersion", func(t *testing.T) {
		rows, err := zmysql.Select(db, "SELECT version()")
		if err != nil {
			t.Errorf("Failed to execute query: %v", err)
		}
		defer rows.Close()

		for rows.Next() {
			var version string
			err = rows.Scan(&version)
			if err != nil {
				t.Errorf("Failed to scan row: %v", err)
			}
			t.Logf("Database version: %s", version)
		}

		if err = rows.Err(); err != nil {
			t.Errorf("Row iteration error: %v", err)
		}
	})
}

func TestInsert(t *testing.T) {
	if db == nil {
		t.Fatal("Database connection is not initialized")
	}

	// 插入测试数据
	_, err := db.Exec("INSERT INTO test_table (id, name) VALUES (1, 'test1')")
	if err != nil {
		t.Errorf("Failed to insert data: %v", err)
	}

	t.Log("Data inserted")
}

func TestSelectData(t *testing.T) {
	if db == nil {
		t.Fatal("Database connection is not initialized")
	}

	// 查询测试数据
	rows, err := db.Query("SELECT id, name FROM test_table")
	if err != nil {
		t.Errorf("Failed to select data: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			t.Errorf("Failed to scan row: %v", err)
		}
		t.Logf("Data: id=%d, name=%s", id, name)
	}

	if err = rows.Err(); err != nil {
		t.Errorf("Row iteration error: %v", err)
	}
}
