package api

import (
        "database/sql"
        "fmt"
        "github.com/gin-gonic/gin"
        _ "github.com/go-sql-driver/mysql"
        "net/http"
)

type Param struct {
        Name   string
        Salary int
        Age    int
}

type Employee struct {
        ID     int    `json:"id"`
        Name   string `json:"name"`
        Salary int    `json:"salary"`
        Age    int    `json:"age"`
}

func GetAPI(db *sql.DB, router *gin.Engine) {

        router.GET("/employee/:id", func(c *gin.Context) {
                var (
                        result   gin.H
                        employee Employee
                )
                id := c.Param("id")
                row := db.QueryRow("select * from employee where id=?;", id)
                err := row.Scan(&employee.ID, &employee.Name, &employee.Salary, &employee.Age)

                if err != nil {
                        result = gin.H{"count": 0, "result": sql.ErrNoRows}
                } else {
                        result = gin.H{"result": employee, "count": 1}
                }

                c.JSON(http.StatusOK, result)
        })
          router.GET("/employees", func(c *gin.Context) {

                var (
                        results   gin.H
                        employee  Employee
                        employees []Employee
                )
                rows, err := db.Query("select * from employee")
                if err != nil {
                        results = gin.H{"count": 0, "result": "Failed to query datas"}
                } else {
                        for rows.Next() {
                                err = rows.Scan(&employee.ID, &employee.Name, &employee.Salary, &employee.Age)
                                if err != nil {
                                        continue
                                } else {
                                        employees = append(employees, employee)
                                }
                        }
                }
                results = gin.H{"count": len(employees), "result": employees}
                c.JSON(http.StatusOK, results)
        })
}

func PostAPI(db *sql.DB, router *gin.Engine) {
        router.POST("/employee/add", func(c *gin.Context) {
                var e Param
                var result gin.H

                err := c.Bind(&e)
                if err != nil || e.Name == "" {
                        result = gin.H{"count": 0, "result": "failed to bind parameters"}
                } else {

                        stmt, err := db.Prepare("insert into employee (name,salary,age) values(?,?,?);")
                        if err != nil {
                                result = gin.H{"count": 0, "result": "failed to prepare insert command"}
                        } else {

                                _, err = stmt.Exec(e.Name, e.Salary, e.Age)
                                if err != nil {
                                        result = gin.H{"count": 0, "result": "failed to execute insert command"}
                                                                  } else {
                                        result = gin.H{"count": 1, "result": fmt.Sprintf("create user %v successfully", e.Name)}
                                }
                        }
                }
                c.JSON(http.StatusOK, result)
        })
}

func DeleteAPI(db *sql.DB, router *gin.Engine) {
        router.DELETE("/employee/:id", func(c *gin.Context) {
                var result gin.H
                id := c.Param("id")
                stmt, err := db.Prepare("delete from employee where id=?")
                if err != nil || id == "" {
                        result = gin.H{"count": 0, "result": "failed to prepare delete command"}
                } else {
                        _, err = stmt.Exec(id)
                        if err != nil {
                                result = gin.H{"count": 0, "result": "failed to exec delete command"}
                        } else {
                                result = gin.H{"count": 1, "result": fmt.Sprintf("delete id %v successfully", id)}
                        }
                }
                c.JSON(http.StatusOK, result)
        })
}

func PutAPI(db *sql.DB, router *gin.Engine) {
        router.PUT("/employee/:id", func(c *gin.Context) {
                var result gin.H
                var e Param

                id := c.Param("id")

                err := c.Bind(&e)
                if err != nil || e.Name == "" {
                        fmt.Println(err)
                        result = gin.H{"count": 0, "result": "failed to bind parameters"}
                } else {
                        stmt, err := db.Prepare("update employee set name=?, salary=?, age=? where id=?;")
                        if err != nil {
                                fmt.Println(err)
                                                          result = gin.H{"count": 0, "result": "failed to prepare update command"}
                        } else {
                                _, err = stmt.Exec(e.Name, e.Salary, e.Age, id)
                                if err != nil {
                                        fmt.Println(err)
                                        result = gin.H{"count": 0, "result": "failed to execute update command"}
                                } else {
                                        result = gin.H{"count": 1, "result": fmt.Sprintf("update id %v successfully", id)}
                                }
                        }
                }
                c.JSON(http.StatusOK, result)
        })
}
