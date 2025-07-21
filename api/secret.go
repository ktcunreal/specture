package api

import (
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/go-chi/chi"
	"github.com/tidwall/sjson"
	"net/http"
	"strconv"
	"os/exec"
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
		result := utils.AppendIfNotExist(config.GetWhitelistPath(), fmt.Sprintf("%s/32", r.RemoteAddr))
		
		exec.Command("systemctl", "reload", "haproxy")

		var resJsonStr string
		resJsonStr, _ = sjson.Set(resJsonStr, "sourceIp", fmt.Sprintf("%s", r.RemoteAddr))
		resJsonStr, _ = sjson.Set(resJsonStr, "result", fmt.Sprintf("%s%s", result))

		w.Write([]byte(resJsonStr))
		
	} else {
		http.Redirect(w, r, config.GetDummyUrl(), http.StatusMovedPermanently)
	}
}
