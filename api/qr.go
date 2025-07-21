package api

import(
	"specture/internal/config"
	"github.com/go-chi/chi"
	"net/http"
	"time"
	"specture/utils"
	"github.com/charmbracelet/log"
	qrcode "github.com/skip2/go-qrcode"
	"fmt"
	"strconv"
)

func QRRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/{secret}", issueQR)
	r.Get("/*", bye)
	return r
}

func issueQR(w http.ResponseWriter, r *http.Request) {
	if chi.URLParam(r, "secret") == config.GetPresharedKey() { 
		if config.GetBaseUrl() == "" {
			log.Error("Url not specified")
			return
		}

		ts_now := int(time.Now().Unix())
		qrtext := fmt.Sprintf("%s/%s?issuetime=%d", 
			fmt.Sprintf("%s/secret",config.GetBaseUrl()), 
			utils.SHA256STR(config.GetPresharedKey() + strconv.Itoa(ts_now)),
			ts_now,
		)

		var png []byte
		log.Infof("Generating New QR Code for: %s",qrtext)
		png, err := qrcode.Encode(qrtext, qrcode.Medium, 256) 
		if err != nil {
			log.Error(err)
		}

		w.Header().Set("Content-Type", "image/png")
		w.Write(png)

	} else {
		http.Redirect(w, r, config.GetDummyUrl(), http.StatusMovedPermanently)
	}
}
