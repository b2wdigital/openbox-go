package main

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
	"log"
	"strings"
	"text/template"

	"github.com/b2wdigital/goignite/pkg/config"
	"github.com/b2wdigital/goignite/pkg/log/logrus"
	"github.com/b2wdigital/openbox-go/internal/app/http/router/model"
)

const (
	Output   = "openbox.output"
	Template = "openbox.template"
)

func init() {

	log.Println("getting default configurations for generator")

	config.Add(Output, "gen", "generator output path")
	config.Add(Template, "templates/http/router/echo.tpl", "generator template")
}

func main() {

	var err error

	err = config.Parse()
	if err != nil {
		log.Fatal(err)
	}

	options := new(model.Options)

	err = config.UnmarshalWithPath("openbox", &options)
	if err != nil {
		log.Fatal(err)
	}

	if len(options.RequestMaps) == 0 {
		log.Fatal("no request maps found")
	}

	logrus.Start()

	tmpl := template.Must(template.ParseFiles(options.Template))

	data := model.Data{
		Start: options.Start,
		Stop: options.Stop,
		RequestMaps: options.RequestMaps,
	}

	packages := getPackagesFromRequestMaps(data.RequestMaps)

	if data.Start != nil {
		data.Start.Alias = getAlias(data.Start.Package)
		packages = append(packages, data.Start.Package)
	}

	if data.Stop != nil {
		data.Stop.Alias = getAlias(data.Stop.Package)
		packages = append(packages, data.Stop.Package)
	}

	for _, r := range data.RequestMaps {
		r.Handler.Alias = getAlias(r.Handler.Package)

		if r.Method == "POST" || r.Method == "PUT" || r.Method == "PATCH" {
			r.Body.Alias = getAlias(r.Body.Package)
		}
	}

	data.Packages = getUniquePackages(packages)

	var buf bytes.Buffer

	err = tmpl.Execute(&buf, data)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(strings.Join([]string{options.Output, "main.go"}, "/"), buf.Bytes(), 0644)
	if err != nil {
		log.Fatal(err)
	}

}

func getPackagesFromRequestMaps(maps []*model.RequestMap) []string {
	var packages []string

	for _, v := range maps {
		packages = append(packages, v.Handler.Package)
		if v.Body != nil {
			packages = append(packages, v.Body.Package)
		}
	}
	return packages
}

func getUniquePackages(packages []string) []*model.Package {
	packages = unique(packages)

	var p []*model.Package

	for _, uri := range packages {
		alias := getAlias(uri)
		p = append(p, &model.Package{Alias: alias, URI: uri})
	}

	return p
}

func getAlias(uri string) string {
	hasher := md5.New()
	hasher.Write([]byte(uri))
	return strings.Join([]string{"p", hex.EncodeToString(hasher.Sum(nil))}, "")
}

func unique(stringSlice []string) []string {
	keys := make(map[string]bool)
	var list []string
	for _, entry := range stringSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
