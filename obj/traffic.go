package obj

//import (
//	"net"
//	"time"
//)
//
//type Event struct {
//	Date     time.Time
//	SourceIP net.IP
//	DstIP    net.IP
//	Category int
//	Checksum string
//	FileName string
//}
//
//type Traffic struct {
//	Date   time.Time
//	BpsIn  int64
//	BpsOut int64
//	PpsIn  int64
//	PpsOut int64
//}
//
//type Status struct {
//	Date        time.Time
//	CpuUsage    float64
//	MemoryUsage float64
//	DiskUsage   float64
//}
//
//
//func main() {
//	c := colly.NewCollector()
//
//	// Find and visit all links
//	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
//		e.Request.Visit(e.Attr("href"))
//	})
//
//	c.OnRequest(func(r *colly.Request) {
//		fmt.Println("Visiting", r.URL)
//	})
//
//	c.Visit("http://go-colly.org/")
//}
