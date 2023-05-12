package middleware

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Metric() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		log.Println("[middleware][Metric] after next: ", c.Request.Host, c.Request.RequestURI, c.Request.Method, strconv.Itoa(c.Writer.Status()))
		httpCountVec.WithLabelValues(
			c.Request.Host,
			c.Request.RequestURI,
			c.Request.Method,
			strconv.Itoa(c.Writer.Status()),
		).Inc()
	}
}

var (
	httpCountVec = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_count_vec",
			Help: "The total number of handled requests",
		},
		[]string{"host", "path", "method", "status"},
	)
)

func Start() {
	httpCountVec.WithLabelValues("host", "path", "method", "status").Inc()
	if err := prometheus.Register(httpCountVec); err != nil {
		log.Fatal("[middleware][Start] Register error: ", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())
	server := &http.Server{Addr: ":9010", Handler: mux}
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal("[middleware][Start] ListenAndServe error: ", err)
	}
}
