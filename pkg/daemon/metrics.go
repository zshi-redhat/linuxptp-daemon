package daemon

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang/glog"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	utilwait "k8s.io/apimachinery/pkg/util/wait"
)

// Message defines log message structure transmit
// via MessageChannel.
type Message struct {
	Name	string
	Type	string
	Content	string
}

// Parser is the main structure for parse instance.
// It contains all the necessary data to parse linuxptp process log messages.
type Parser struct {
	// node name where daemon is running
	nodeName	string
	namespace	string

	ptpUpdate	*LinuxPTPConfUpdate
	// channel ensure Parser.Run() exit when main function exits.
	// stopCh is created by main function and passed by Daemon via NewParser()
	stopCh		<-chan struct{}
	messageChannel	<-chan Message
}

const (
	metricsPath = "/metrics"
	healthzPath = "/healthz"
)

var (
	ptp4lRootMeanSquareValue = 0
	Ptp4lRootMeanSquareValue = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "ptp4l_root_mean_square_value",
			Help: "Metric to get root mean square of master offset for ptp4l instance.",
		}, []string{"node", "network"})
)

func setPtp4lRootMeanSquareValue(node string, ptpnetwork string, value int) {
	Ptp4lRootMeanSquareValue.With(prometheus.Labels{
		"node": node,"network": ptpnetwork}).Set(float64(value))
}

func updatePtp4lMetrics(node string, ptpnetwork string, value int) {
	setPtp4lRootMeanSquareValue(node, ptpnetwork, value)
}

// NewParser is called by daemon to create new parser instance
func NewParser(
	nodeName	string,
	namespace	string,
	ptpUpdate	*LinuxPTPConfUpdate,
	stopCh		<-chan struct{},
	messageChannel	<-chan Message,
) *Parser {
	return &Parser{
		nodeName:	nodeName,
		namespace:	namespace,
		ptpUpdate:	ptpUpdate,
		stopCh:		stopCh,
		messageChannel:	messageChannel,
	}
}

// Run in a for loop to listen for log messages in MessageChannel
func (ps *Parser) Run() {

        // Register metrics
        prometheus.MustRegister(Ptp4lRootMeanSquareValue)

        // Including these stats kills performance when Prometheus polls with multiple targets
        prometheus.Unregister(prometheus.NewProcessCollector(prometheus.ProcessCollectorOpts{}))
        prometheus.Unregister(prometheus.NewGoCollector())

	for {
		select {
		case msg := <-ps.messageChannel:
			switch msg.Type {
			case "ptp4l":
				glog.V(2).Infof("Parser message(%s) received: %v", msg.Type, msg.Content)
			case "phc2sys":
				glog.V(2).Infof("Parser message(%s) received: %v", msg.Type, msg.Content)
			}
		case <-ps.stopCh:
			glog.Infof("linuxPTP Parser stop signal received, existing..")
			return
		}
	}
	return
}

func StartHTTPMetricServer(address string) {
        mux := http.NewServeMux()
        mux.Handle(metricsPath, promhttp.Handler())

        // Add healthzPath
        mux.HandleFunc(healthzPath, func(w http.ResponseWriter, r *http.Request) {
                w.WriteHeader(http.StatusOK)
                w.Write([]byte(http.StatusText(http.StatusOK)))
        })
        // Add index
        mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
                w.Write([]byte(`<html>
                 <head><title>Linuxptp Daemon Metrics Server</title></head>
                 <body>
                 <h1>Kube Metrics</h1>
                 <ul>
                 <li><a href='` + metricsPath + `'>metrics</a></li>
                 <li><a href='` + healthzPath + `'>healthz</a></li>
                 </ul>
                 </body>
                 </html>`))
        })

        go utilwait.Until(func() {
                err := http.ListenAndServe(address, mux)
                if err != nil {
                        utilruntime.HandleError(fmt.Errorf("starting metrics server failed: %v", err))
                }
        }, 5*time.Second, utilwait.NeverStop)
}
