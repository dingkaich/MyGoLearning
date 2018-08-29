package myjsonrest

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/sirupsen/logrus"
)

var mylog = logrus.New()

// var outlog *log.Logger

func logger() {

	// The API for setting attributes is a little different than the package level
	// exported logger. See Godoc.
	// mylog.Out = os.Stdout

	// You could set this to any `io.Writer` such as a file
	file, err := os.OpenFile("logrus.mylog", os.O_CREATE|os.O_WRONLY, 0666)
	if err == nil {
		mylog.Out = file
	} else {
		mylog.Info("Failed to mylog to file, using default stderr")
	}
	// outlog = log.New(mylog.Out, "", log.LstdFlags)
	mylog.Println("go to myjsonrest")

}

func MyjsonrestMain() {
	logger()

	var mymiddleware = []rest.Middleware{
		&rest.AccessLogApacheMiddleware{
			Format: rest.CommonLogFormat,
		},
		&rest.TimerMiddleware{},
		&rest.RecorderMiddleware{},
		&rest.RecoverMiddleware{
			Logger: log.New(mylog.Out, "", log.LstdFlags),
			EnableResponseStackTrace: true,
		},
		&rest.JsonIndentMiddleware{},
		// &rest.JsonpMiddleware{},
		&rest.ContentTypeCheckerMiddleware{},
	}

	mylog.Println("begin listen")
	api := rest.NewApi()
	api.Use(mymiddleware...)

	router, err := rest.MakeRouter(
		rest.Post("/message", func(w rest.ResponseWriter, req *rest.Request) {
			// fmt.Println("aaa")
			io.Copy(os.Stderr, req.Body)
			w.WriteJson(map[string]string{"Body": "Hello World!"})
		}),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)

	http.Handle("/api/", http.StripPrefix("/api", api.MakeHandler()))

	log.Fatal(http.ListenAndServe(":8080", nil))

}
