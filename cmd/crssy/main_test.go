package main

import "testing"

func ignore_Example_Main() {
	goMain([]string{})
}

func Example_Completion() {
	goMain([]string{"./crssy", "--generate-completions"})
	// Output:
	//
	//crssy [OPTIONS]
	//	OPTIONS
	//		<LOCATION>   	引数に場所を指定をし，出力で指定した場所の天気を返す．
	//		-v, --version 			バージョンを表示し，修了します.
	//		-h, --help 			このメッセージを表示し，修了します.
	//		-w, --week <WEEK>   	    	出力で週の天気を返す.
}

func Test_Main(t *testing.T) {
	if status := goMain([]string{"./crssy", "Gifu"}); status != 0 {
		t.Error("Expected 0, got ", status)
	}
}
