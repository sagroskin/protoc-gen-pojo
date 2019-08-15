package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
	"github.com/sagroskin/protoc-gen-pojo/generator"
	"golang.org/x/crypto/ssh/terminal"
)

var (
	flagDeclareNamespace      = flag.Bool("declare_namespace", true, "if true, generate a namespace declaration")
	flagEnumsAsStrings        = flag.Bool("string_enums", false, "if true, generate string enums")
	flagOutputFilenamePattern = flag.String("outpattern", "{{.Dir}}/{{.BaseName}}.java", "output filename pattern")
)

func main() {
	g := generator.New()

	if terminal.IsTerminal(0) {
		flag.Usage()
		log.Fatalln("stdin appears to be a tty device. This tool is meant to be invoked via the protoc command via a --pojo_out directive.")
	}

	data, err := ioutil.ReadAll(os.Stdin)

	if err != nil {
		log.Fatalln(errors.Wrap(err, "Failed to read input."))
	}

	if err := proto.Unmarshal(data, g.Request); err != nil {
		log.Fatalln(errors.Wrap(err, "Failed to parse input."))
	}

	if len(g.Request.FileToGenerate) == 0 {
		log.Fatalln(errors.Wrap(err, "No files to generate."))
	}

	parseFlags(g.Request.Parameter)

	g.GenerateAllFiles(&generator.Parameters{
		DeclareNamespace:  *flagDeclareNamespace,
		EnumsAsStrings:    *flagEnumsAsStrings,
		OutputNamePattern: *flagOutputFilenamePattern,
	})

	data, err = proto.Marshal(g.Response)

	if err != nil {
		log.Fatalln(errors.Wrap(err, "Failed to marshal output proto."))
	}

	_, err = os.Stdout.Write(data)

	if err != nil {
		log.Fatalln(errors.Wrap(err, "Failed to write output proto."))
	}
}

func parseFlags(s *string) {
	if s == nil {
		return
	}

	for _, p := range strings.Split(*s, ",") {
		spec := strings.SplitN(p, "=", 2)

		if len(spec) == 1 {
			if err := flag.CommandLine.Set(spec[0], ""); err != nil {
				log.Fatalln("Cannot set flag", p, err)
			}
			continue
		}

		name, value := spec[0], spec[1]

		if err := flag.CommandLine.Set(name, value); err != nil {
			log.Fatalln("Cannot set flag", p)
		}
	}
}
