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
  // "fmt"
  "log"
  "time"
  "net/http"
  "github.com/labstack/echo"
)

func handler(c echo.Context) error {
  if c.Param("guess") == os.Args[1] {
    storeLog(c.Param("username"));
    return c.String(http.StatusOK, "CORRECT\n")
  } else {
    return c.String(http.StatusUnauthorized, "WRONG\n")
  }
}

func storeLog(msg string) {
  // If the file doesn't exist, create it, or append to the file
	f, err := os.OpenFile("winners.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

  t := time.Now().Format("2006-01-02 15:04:05.23")

	if _, err := f.Write([]byte(t + " -- " + msg + "\n")); err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

func main() {
  e := echo.New()
  e.GET("/:username/:guess", handler)
  e.HEAD("/:username/:guess", handler)
  e.Logger.Fatal(e.Start(":8080"))
}
