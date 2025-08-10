package configs

import "log"

var AppName string

func Load() {
    AppName = "MyApp"
    log.Println("Configurações carregadas.")
}
