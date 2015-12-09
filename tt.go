// tt (template [cli] tool) is a Go text/html template command line tool.
//
// usage:
//
//  tt <(echo 'hello {{.data}}') <(echo '{"data": "world"}')
//  tt -t h <(echo 'data is {{.a}}') <(echo '{"a": "<tag>"}')
package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"

	htemplate "html/template"
	ttemplate "text/template"
)

func main() {
	typ := flag.String("t", "text", "specify template type: t/text, h/html")
	flag.Parse()
	if len(flag.Args()) != 2 {
		println("usage: tt template json")
		os.Exit(1)
	}

	tmpl, err := ioutil.ReadFile(flag.Args()[0])
	if err != nil {
		log.Fatal(err)
	}
	dataSrc, err := ioutil.ReadFile(flag.Args()[1])
	if err != nil {
		log.Fatal(err)
	}
	var data interface{}
	if err := json.Unmarshal(dataSrc, &data); err != nil {
		log.Fatal(err)
	}

	if *typ == "h" || *typ == "html" {
		if err := htemplate.Must(htemplate.New("main").Parse(string(tmpl))).Execute(os.Stdout, data); err != nil {
			log.Fatal(err)
		}
	} else {
		if err := ttemplate.Must(ttemplate.New("main").Parse(string(tmpl))).Execute(os.Stdout, data); err != nil {
			log.Fatal(err)
		}
	}
}
