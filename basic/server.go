package basic

// import (
// 	"fmt"
// 	"net/http"
// 	"strconv"

// 	"github.com/ethereum/go-ethereum/common"
// 	"github.com/gin-gonic/gin"
// 	"github.com/syndtr/goleveldb/leveldb"
// )

// // Server transactions
// type Server struct {
// 	Db   leveldb.DB
// 	Port int
// 	Cgin gin.Engine
// }

// // CreateServer Create and initialize server
// func CreateServer(pathDb string, port int) (Server, error) {
// 	c := gin.Default()
// 	db, err := leveldb.OpenFile("./tmp", nil)
// 	if err != nil {
// 		return Server{}, err
// 	}

// 	s := Server{*db, port, *c}

// 	s.Init()

// 	return s, nil
// }

// // WithServer wrapper to encapsulate http calls
// func WithServer(srv *Server, handler func(c *gin.Context, srv *Server)) func(c *gin.Context) {
// 	return func(c *gin.Context) {
// 		handler(c, srv)
// 	}
// }

// // Init all routes
// func (s *Server) Init() {
// 	// s.Cgin.POST("/mint/:address/:coins", mint)

// 	// s.Cgin.POST("/mint/:address/:coins", WithServer(s, mint))
// 	s.Cgin.POST("/mint/:address/:coins", func(c *gin.Context) { mint(c, s) })

// 	s.Cgin.POST("/tx", postTx)
// 	s.Cgin.POST("/txs/:address", getTxs)
// 	s.Cgin.POST("/balance/:address", getBalance)

// 	portNumber := strconv.Itoa(s.Port)
// 	s.Cgin.Run(":" + portNumber)
// }

// func mint(c *gin.Context, s *Server) {
// 	paramAddress := c.Param("address")

// 	paramCoins := c.Param("coins")
// 	coins, err := strconv.Atoi(paramCoins)
// 	if err != nil {
// 		c.String(http.StatusNotAcceptable, "Fail")
// 	}
// 	fmt.Println(paramAddress)
// 	fmt.Println(paramCoins)

// 	address := common.HexToAddress(paramAddress)

// 	data, err := db.Get([]byte("key"), nil)
// }

// func postTx(c *gin.Context) {
// 	var tx Tx

// 	err := c.ShouldBind(&tx)
// 	if err != nil {
// 		c.String(http.StatusNotAcceptable, "Fail")
// 	}
// 	// Validate signature
// 	// Check addresses exist on Db
// 	// Modify address balances
// }

// func getTxs(c *gin.Context) {

// }

// func getBalance(c *gin.Context) {

// }
