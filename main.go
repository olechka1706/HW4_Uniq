package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

func main() {

	opt := Options{}
	flag.BoolVar(&opt.Count, "c", false, "Подсчитать количество встречаний строки во входных данных.")
	flag.BoolVar(&opt.Repeated, "d", false, "Вывести только те строки, которые повторились во входных данных.")
	flag.BoolVar(&opt.Unique, "u", false, "Вывести только те строки, которые не повторились во входных данных.")
	flag.IntVar(&opt.SkipFields, "f", 0, "Не учитывать первые num_fields полей в строке.")
	flag.IntVar(&opt.SkipChars, "s", 0, "Не учитывать первые num_chars символов в строке.")
	flag.BoolVar(&opt.IgnoreCase, "i", false, "Не учитывать регистр букв.")

	flag.Parse()

	if (opt.Count && opt.Repeated) || (opt.Count && opt.Unique) || (opt.Repeated && opt.Unique) {
		fmt.Println(errors.New("используйте только один из флагов -с -d -u"))
		return
	}

	if len(flag.Args()) == 1 {
		input_file, err := os.Open(flag.Arg(0))
		if err != nil {
			panic(err)
		}
		defer input_file.Close()
		Uniq(input_file, os.Stdout, opt)

	} else if len(flag.Args()) == 2 {
		input_file, err := os.Open(flag.Arg(0))
		if err != nil {
			panic(err)
		}
		output_file, err := os.OpenFile(flag.Arg(1), os.O_RDWR|os.O_CREATE, 0644)
		if err != nil {
			panic(err)
		}
		defer func() {
			input_file.Close()
			output_file.Close()
		}()
		Uniq(input_file, output_file, opt)
	} else {
		Uniq(os.Stdin, os.Stdout, opt)
	}

}


