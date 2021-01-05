package main

import (
	"io/ioutil"
	"os"

	"github.com/gogo/protobuf/protoc-gen-gogo/generator"
	"google.golang.org/protobuf/proto"
)

func main() {
	g := generator.New()
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		g.Error(err, "reading input")
	}
	if err = proto.Unmarshal(data, g.Request); err != nil {
		g.Error(err, "parsing input proto")
	}
	if len(g.Request.FileToGenerate) == 0 {
		g.Fatal("no files to generate")
	}
	g.CommandLineParameters(g.Request.GetParameters())

	// Create a wrapped version of the Descriptor and EnumDescriptors
	// that point to the file that defines them.
	g.WrapTypes()
	g.SetPackageNames()
	g.BuildTypeNameMap()
	g.GenerateAllFiles()

	// Send back the results.
	data, err = proto.Marshal(g.Response)
	if err != nil {
		g.Error(err, "failed to marshal output proto")
	}
	if _, err := os.Stdout.Write(data); err != nil {
		g.Error(err, "failed to write output proto")
	}
}
