package model

import (
	"fmt"
)

//Parameter для хранения входящих параметров из запроса
type Parameter struct {
	Name      string
	Value     string
	DBName    string
	Operation string
}

//Parameters управление параметрами
type Parameters struct {
	Params     []Parameter
	Config     []ConfigParameter
	Controller string
}

//ConfigParameter параметры для правильного формирования запроса
type ConfigParameter struct {
	Controllrt string `json:"controller"`
	Name       string `json:"name"`
	DBName     string `json:"dbname"`
	Operation  string `json:"operation"`
	Single     bool   `json:"single"`
	Ignore     bool   `json:"ignore"`
}

func (p *Parameters) Search(name, controller string) (cfg ConfigParameter, ok bool) {
	for _, cp := range p.Config {
		if cp.Name == name && (cp.Controllrt == controller || cp.Controllrt == "*") {
			fmt.Println("search param", name, controller, cp)
			if cp.Ignore {
				ok = false
				return
			}
			cfg = cp
			ok = true
			return
		}
	}
	ok = false
	return
}

//ToQuery получить строку запроса к БД
func (p *Parameters) ToQuery() (query string, single bool) {
	and := ""
	sep := ""
	query = ""
	single = false
	for _, param := range p.Params {
		cfg, ok := p.Search(param.Name, p.Controller)
		fmt.Println("ok", ok, "param", param.Name)
		if !ok {
			continue
		}
		if cfg.Ignore {
			continue
		}
		if cfg.Operation == "LIKE" {
			sep = "%"
		}
		query += and + cfg.DBName + " " + cfg.Operation + " '" + sep + param.Value + sep + "'"
		if and == "" {
			and = "AND "
		}
		if !single {
			single = cfg.Single
		}
	}
	fmt.Println("PARAMETERS:", p.Params)
	fmt.Println(query)
	return
}

//Add добавить параметр
func (p *Parameters) Add(parameter Parameter) {
	p.Params = append(p.Params, parameter)
}

//New создать новый объект
func (p *Parameters) New(parameters []Parameter) Parameters {
	return Parameters{Params: parameters}
}
