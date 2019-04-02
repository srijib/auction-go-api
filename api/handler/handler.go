package handler
 
import (
	"fmt"
	
	"net/http"
	"github.com/jinzhu/gorm"
)
func CreateBidEndpoint(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	fmt.Println("HELLO")
	// decoder := json.NewDecoder(r.Body)
	// bid := model.Bid{}
	// err := decoder.Decode(&bid)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(bid.BidPrice)
}
