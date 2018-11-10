/*
* USAGE: ./codewars <password>
*
* curl localhost:8080/<password>         // CORRECT
* curl localhost:8080/<not-the-password> // WRONG
*
*/

package main

import (
  "os"
  "fmt"
  // "log"
  "net/http"
  "github.com/labstack/echo"
)

func handler(c echo.Context) error {
  fmt.Println(c.Param("username"));
  if c.Param("guess") == os.Args[1] {
    return c.String(http.StatusOK, "CORRECT\n")
  } else {
    return c.String(http.StatusOK, "WRONG\n")
  }
}

func main() {
  e := echo.New()
  e.GET("/:username/:guess", handler)
  e.Logger.Fatal(e.Start(":8080"))
}
