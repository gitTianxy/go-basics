package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "io"
    "log"
    "net/http"
    "os"
)

func init() {
    // logger custom
    f, _ := os.OpenFile("web/gin.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
    log.SetOutput(gin.DefaultWriter)

}

func main() {
    r := gin.Default()
    //r := gin.New()
    //r.Use(gin.Logger())
    //r.Use(gin.Recovery())


    r.GET("/ping", func(c *gin.Context) {
       c.JSON(200, gin.H{
           "message": "pong",
       })
    })
    r.GET("/query", querying)

    crudApi(r)

    r.Run("localhost:8080")
}

func querying(c *gin.Context) {
    params := map[string]interface{}{}
    k := "required_param"
    v := c.Query(k)
    params[k] = v
    k = "optional_param"
    v = c.DefaultQuery(k, "default value")
    params[k] = v
    c.JSON(http.StatusOK, gin.H{
        "message": fmt.Sprintf("query with params: %v", params),
    })
}

func crudApi(r *gin.Engine) {
    api := r.Group("/api")
    {
        api.GET("/get/:id", getting)
        api.POST("/post", posting)
        api.PUT("/put/:id", putting)
        api.DELETE("/delete/:id", deleting)
    }
}

func getting(c *gin.Context) {
    id := c.Param("id")
    c.JSON(http.StatusOK, gin.H{
        "message": fmt.Sprintf("%v got", id),
    })
}

type RequestBody struct {
    Id   int    `json:"id" form:"id"`
    Name string `json:"name" form:"name" binding:"required"`
    Age  int    `json:"age" form:"age"`
}

func posting(c *gin.Context) {
    var body RequestBody
    err := c.Bind(&body)
    if err != nil {
        //c.JSON(http.StatusBadRequest, gin.H{
        //    "message": "wrong body",
        //})
        log.Println("post err:", err.Error())
        return
    }
    /* get body data
    x, _ := c.GetRawData()
    fmt.Println("raw body:", string(x))*/
    log.Println("post:", body)
    c.JSON(http.StatusCreated, gin.H{
        "message": fmt.Sprintf("%v created", body),
    })
}

func putting(c *gin.Context) {
    id := c.Param("id")
    var body RequestBody
    c.Bind(&body)
    c.JSON(http.StatusCreated, gin.H{
        "message": fmt.Sprintf("%v updated", id),
    })
}
func deleting(c *gin.Context) {
    id := c.Param("id")
    c.JSON(http.StatusOK, gin.H{
        "message": fmt.Sprintf("%v deleted", id),
    })
}
