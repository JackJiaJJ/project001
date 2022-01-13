package main

import (
        "api"
        //"database/sql"
        "encoding/json"
        "fmt"
        "io"
        "net/http"
)

func main() {
        fmt.Println("start http server")

        http.HandleFunc("/employees", GetEmployees)
        err := http.ListenAndServe("x.xx.xxx.xxx:8081", nil)
        if err != nil {
                fmt.Println("Failed to ListenAndServe on :8081, error is -", err)
                return
        }
}

func GetEmployees(w http.ResponseWriter, r *http.Request) {
        var e api.Employee
        type Employees struct {
                Count          int
                EmployeeInfors []api.Employee `json:"employees"`
        }

        //var employees Employees
        var employeeInfors []api.Employee
        var result []byte

        db, err := api.DBOpen("mysql", "root:Zhu88jie!@tcp(x.xx.xxx.xxx:3306)/employee")
        if err != nil {
                fmt.Println("Failed to open db, error is -", err)
                return
        }
        fmt.Println("Connect db successfully")

        //w.Write([]byte("This is the first http script\n"))
        rows, err := db.Query("select * from employee")
        if err != nil {
                          io.WriteString(w, "Failed to get employee informations")
        } else {
                for rows.Next() {
                        err = rows.Scan(&e.ID, &e.Name, &e.Salary, &e.Age)
                        if err != nil {
                                fmt.Println("info: failed to run rows.Scan, try next line data, error is -", err)
                                continue
                        } else {
                                fmt.Println("employee is -", e)
                                employeeInfors = append(employeeInfors, e)
                        }
                }
                employees := Employees{
                        Count:          len(employeeInfors),
                        EmployeeInfors: employeeInfors,
                }

                fmt.Println(employees)

                result, err = json.MarshalIndent(employees, "", "\t")
                if err != nil {
                        w.Write([]byte("Failed to marshal result, error is -" + err.Error()))
                }
                w.Write(result)
        }
        defer rows.Close()
}
