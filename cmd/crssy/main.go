package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/YukiKonishi/crssy"
	flag "github.com/spf13/pflag"
)

const VERSION = "0.1.8"

func main() {
	//status := crssy.WeatherTemplate(os.Args)
	status := goMain(os.Args)
	os.Exit(status)
}

type options struct {
	version  bool
	help     bool
	location string
	week     string
}

type CrssyError struct {
	statusCode int
	message    string
}

func (e *CrssyError) Error() string {
	return e.message
}

// ヘルプメッセージ
func helpMessage(args string) string {
	return fmt.Sprintf(`%s [OPTIONS]
	OPTIgitONS
		<LOCATION>   	引数に場所を指定をし，出力で指定した場所の天気を返す．
		-v, --version 			バージョンを表示し，修了します.
		-h, --help 			このメッセージを表示し，修了します.
		-w, --week <WEEK>   	    	出力で週の天気を返す.`, args)
}

// バージョン
func versionString(args []string) string {
	prog := "crssy"
	if len(args) > 0 {
		prog = filepath.Base(args[0])
	}
	return fmt.Sprintf("%s version %s", prog, VERSION)
}

// 定義
func buildOptions(args []string) (*options, *flag.FlagSet) {
	opts := &options{}
	flags := flag.NewFlagSet(args[0], flag.ContinueOnError)
	flags.Usage = func() { fmt.Println(helpMessage(args[0])) }
	flags.BoolVarP(&opts.version, "version", "v", false, "バージョンを表示し，修了します")
	flags.BoolVarP(&opts.help, "help", "h", false, "このメッセージを表示し，修了します．")
	flags.StringVarP(&opts.location, "location", "l", "", "引数に場所を指定をし，出力で指定した場所の天気を返す．")
	flags.StringVarP(&opts.week, "week", "w", "", "出力で週の天気を返す(場所は京都)")
	return opts, flags
}

// 引数に何も与えられていない時
func perform(opts *options, args []string) *CrssyError {
	//fmt.Println("Hello World")
	if len(args) == 0 {
		fmt.Println(helpMessage("crssy"))
		return &CrssyError{statusCode: 0, message: ""}
	}
	city, err := crssy.FindCity(args[0])
	if err != nil {
		return &CrssyError{statusCode: 0, message: err.Error()}
	}
	// fmt.Println(city)
	weathercode, err := crssy.ExpectWeather(city)
	weather, err := crssy.Translateweather(weathercode.Weathercode[0])
	fmt.Println(weathercode.Time[0], weather)
	return nil
}

// 引数が定義にあるものが与えられている時
func parseOptions(args []string) (*options, []string, *CrssyError) {
	opts, flags := buildOptions(args)
	flags.Parse(args[1:])
	if opts.help {
		fmt.Println(helpMessage(args[0]))
		return nil, nil, &CrssyError{statusCode: 0, message: ""}
	}
	if opts.version {
		fmt.Println(versionString(args))
		return nil, nil, &CrssyError{statusCode: 0, message: ""}
	}

	return opts, flags.Args(), nil
}

// main関数
func goMain(args []string) int {
	opts, args, err := parseOptions(args)
	if err != nil {
		if err.statusCode != 0 {
			fmt.Println(err.Error())
		}
		return err.statusCode
	}
	if err := perform(opts, args); err != nil {
		fmt.Println(err.Error())
		return err.statusCode
	}
	return 0
}
