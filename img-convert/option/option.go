package option

import (
	"flag"
	"os"

	"github.com/kzkick2nd/golang-sandbox/img-convert/encoder"
)

// TODO 各プロパティも文字列ではなく構造体にしよう

type Args struct {
	Dir     string
	Decoder string
	Encoder encoder.Encoder
}

func Parse(s []string) (Args, error) {

	f := flag.NewFlagSet(s[0], flag.ExitOnError)
	from := f.String("from", "jpg", "convert from.")
	to := f.String("to", "png", "convert to.")
	f.Parse(s[1:])
	args := f.Args()

	return Args{
		Dir:     validDir(args[0]),
		Decoder: validDecoder(from),
		Encoder: validEncoder(to),
	}, nil

}

func validDir(s string) string {
	if _, err := os.Stat(s); os.IsNotExist(err) {
		return ""
	}
	return s
}

func validDecoder(s *string) string {
	switch *s {
	case "jpg", "jpeg":
		return ".jpg"
	case "png":
		return ".png"
	default:
		return ".jpg"
	}
}

func validEncoder(s *string) encoder.Encoder {
	switch *s {
	case "jpg", "jpeg":
		return &encoder.Jpg{}
	case "png":
		return &encoder.Png{}
	default:
		return &encoder.Png{}
	}
}
