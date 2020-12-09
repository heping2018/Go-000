package Week03

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var Addr = "120.0.0.1:9090"

func main() {
	g, ctx := errgroup.WithContext(context.Background())
	g.Go(func() error {
		return HttpServer()
	})
	g.Go(func() error {
		return SingleListen(ctx)
	})

	if err := g.Wait(); err != nil {
		fmt.Printf("server exception exited")
	}

}

//启动并监听http服务
func HttpServer() error {
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "hello!\n")
	}
	httpHandler := http.NewServeMux()
	httpHandler.HandleFunc("/run", helloHandler)
	s := http.Server{
		Addr:    Addr,
		Handler: httpHandler,
	}
	defer func() {
		if err := recover(); err != nil {
			s.Close()
		}
	}()
	return s.ListenAndServe()
}

//监听系统信号
func SingleListen(ctx context.Context) error {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGINT)

	defer signal.Stop(ch)

	select {
	case <-ctx.Done():
		return fmt.Errorf("http quit")
	case e := <-ch:
		return fmt.Errorf("system signal :%v\n", e)
	}
}
