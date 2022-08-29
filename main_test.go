package main
// import (
// 	"encoding/json"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"
// 	"upvote-crypto/database"
// 	"upvote-crypto/controllers"
// 	"github.com/gin-gonic/gin"
// )

// func SetUpRouter() *gin.Engine{
//     router := gin.Default()
//     return router
// }

// type Votes struct {
// 	id int
// 	Name string
// 	Vote int
// }

// func initDB(dbToUse string) {
// 	db := database.GetDatabase()
//  	db.Exec("USE "+dbToUse)
// }

// func TestAllCurrencies(t *testing.T) {
//     initDB("currencies")

//     req, err := http.NewRequest("GET", "/upvote/currencies", nil)
//     if err != nil {
//         t.Fatal(err)
//     }

//     rr := httptest.NewRecorder()
// 		handler := SetUpRouter()
// 		handler.GET("/upvote/currencies", controllers.ShowCurrencies)
//     handler.ServeHTTP(rr, req)

//     if status := rr.Code; status != http.StatusOK {
//         t.Errorf("handler returned wrong status code: got %v want %v",
//             status, http.StatusOK)
//     }

//     var response []Votes
//     if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
//         t.Errorf("got invalid response, expected list of currencies, got: %v", rr.Body.String())
//     }
// }
