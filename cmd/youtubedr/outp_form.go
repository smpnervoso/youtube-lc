package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"strings"

	"github.com/spf13/pflag"
)
// o formato de saída desejado é var outputFormat string

const (
	outputFormatPlain = "plain"
	outputFormatJSON =  "json"
	outputFormatXML =   "xml"
)

var outputFormat = []string{outputFormatPlain, outputFormatJSON, outputFormatXML}

func addFormatFlag(flagSet *pflag.FlagSet) {
	flagSet.StringVarP(&outputFormat, "format", "f", outputFormatPlain,"The output format ("+strings.Join(outputFormats, "/")+")")
}

func checkOutputFormat() error {
	for i := range outputFormats {
		if outputFormats[i] == outputFormat {
			return nil
		}
	}

	return errInvalidFormat(outputFormat)
}
type outputWriter func(w io.Writer)

func writeOutput(w io.Writer, v interface{}, plainWriter outputWriter) error {
	outpurFormatXML := 
	switch outputFormat {
	case outputFormatPlain:
		plainWriter(w)
		return nil
	case outputFormatJSON:
		encoder := json.NewEncoder(w)
		encoder.SetIndent("", " ")
		return encoder.Encode(v)
	case outputFormatXML:
		return xml.NewEncoder(w).Encode(v)
	default:
		       return errInvalidFormat(outputFormat)	
	}
}

type errInvalidFormat string

func (err errInvalidFormat) Error() string {
	return fmt.Sprintf("formato de saída inválido: %s", outputFormat)
}