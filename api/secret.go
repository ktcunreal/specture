package api

import (
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/go-chi/chi"
	"github.com/tidwall/sjson"
	"net/http"
	"os/exec"
	"strconv"
	"specture/utils"
	"specture/internal/config"
)


func SecretRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/{secret}", addWhitelist)
	r.Get("/*", bye)
	return r
}

func addWhitelist(w http.ResponseWriter, r *http.Request) {
	var issuetime int
	issuetime, err := strconv.Atoi(r.URL.Query().Get("issuetime"))
	if err != nil || !utils.ValidateTimestamp(issuetime) {
		log.Errorf("Invalid cred issue time")
		http.Redirect(w, r, config.GetDummyUrl(), http.StatusMovedPermanently)
		return
	}

	log.Infof("Cred issue time is %v", issuetime)

	if chi.URLParam(r, "secret") == utils.SHA256STR(config.GetPresharedKey() + strconv.Itoa(issuetime)) {
		var resJsonStr string
		cmd := exec.Command("/usr/bin/addwhitelist", r.RemoteAddr, config.GetWhitelistPath())
		stdout, err := cmd.Output()
		if err != nil {
			log.Printf("%v\n", err)
		}
		resJsonStr, _ = sjson.Set(resJsonStr, "sourceIp", fmt.Sprintf("%s", r.RemoteAddr))
		resJsonStr, _ = sjson.Set(resJsonStr, "result", fmt.Sprintf("%s", string(stdout)))
		w.Write([]byte(resJsonStr))
	} else {
		http.Redirect(w, r, config.GetDummyUrl(), http.StatusMovedPermanently)
	}
}
