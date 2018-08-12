package myreflect

import (
	"errors"
	"fmt"
	"reflect"
)

type Config struct {
	LOG      *LOG     `cfgpath:"/log"`
	Network  *Network `cfgpath:"/network"`
	Network1 Network  `cfgpath:"/network"`
	Name     string   `cfgpath:"name"`
}

type Network struct {
	Ipaddress string `cfgpath:"/network/ipaddress"`
	ManagerIp string `cfgpath:"/network/managerIp"`
}

type LOG struct {
	Loglevel       int    `cfgpath:"/log/loglevel"`
	Filepath       string `cfgpath:"/log/filepath" logpath:"/cfg/filepath"`
	OutputTerminal bool   `cfgpath:"/log/OutputTerminal"`
}

func testlog1() {
	var cfg = new(LOG)
	cv := reflect.ValueOf(cfg)
	if cv.CanSet() {
		fmt.Println("can set")
	}
	fmt.Println(cv.Kind())

	for i := 0; i < cv.Elem().NumField(); i++ {
		fmt.Println(cv.Elem().Field(i).CanSet())
		switch cv.Elem().Field(i).Kind() {
		case reflect.String:
			fmt.Println("set string", cv.Elem().Type().Field(i).Tag.Get("logpath"))
			cv.Elem().Field(i).SetString("sss")
		case reflect.Int:
			fmt.Println("set int")
			cv.Elem().Field(i).SetInt(33)
		default:
			fmt.Println(cv.Elem().Kind())
		}
	}

	fmt.Println(cfg)
}
func testlog2() {
	type mynet struct {
		name  string
		value string
	}

	type myconfig struct {
		log1 *LOG
		log2 mynet
	}

	var sss = myconfig{
		log1: new(LOG),
	}
	sss.log1.Filepath = "sadf"
	sss.log2.name = "asdf"
	fmt.Println(sss.log1)
	fmt.Println(sss.log2)
}

func MyReflectMain() {
	var cfg = &Config{
		LOG:     new(LOG),
		Network: new(Network),
	}
	// cfgValue := reflect.ValueOf(cfg).Elem()
	// cfgValue.Field(i)
	LoadConfig(cfg)

	fmt.Println(cfg.LOG)
	fmt.Println(cfg.Network)
	fmt.Println(cfg.Network1)
	fmt.Println(cfg.Name)

	// testlog2()
}

//根据传入参数
func LoadConfig(cfg interface{}) error {
	if reflect.TypeOf(cfg).Kind() != reflect.Ptr {
		return errors.New("please use ptr to be input paramter")
	}

	cfgValue := reflect.ValueOf(cfg).Elem()
	LoadValue(cfgValue)
	return nil
}

func LoadValue(cfgValue reflect.Value) error {
	if cfgValue.Kind() != reflect.Struct {
		if !cfgValue.CanSet() {
			fmt.Println("not struct ,  not canset ")
			return nil
		}
		return nil
	}
	for i := 0; i < cfgValue.NumField(); i++ {
		if !cfgValue.Field(i).CanSet() {
			fmt.Println("field[%s] can not set", cfgValue.Field(i).String())
			continue
			// LoadConfig(cfgValue.Field(i))
		}
		str := cfgValue.Type().Field(i).Tag.Get("cfgpath")
		switch cfgValue.Field(i).Kind() {

		case reflect.String:

			fmt.Println("set string", str)
			cfgValue.Field(i).SetString(str)

		case reflect.Int:
			fmt.Println("set int", str)
			cfgValue.Field(i).SetInt(111)
			// cv.Field(i).SetInt(33)
		case reflect.Bool:
			fmt.Println("set bool", str)

		case reflect.Ptr:
			fmt.Println("set ptr", str)
			LoadValue(cfgValue.Field(i).Elem())
		case reflect.Struct:
			fmt.Println("set strcut", str, cfgValue.Field(i).Type())
			LoadValue(cfgValue.Field(i))

		default:
			fmt.Println("set default", str)
			// LoadValue(cfgValue.Addr().Elem())
		}
	}
	return nil
}
