package climatic_webapp

import (
	"net/http"

	"appengine"
	"appengine/mail"
	"appengine/user"
)

var ATD = Reverse("moc.liamg@vedehthsihsa")
var BCC_ADDR = Reverse("moc.liamg@vedehthsihsa")

func CLWAToAddr(r *http.Request) string {
	if IsLocalHostedOrOnDevBranch(r) {
		DEV_TOLIST := ATD
		return DEV_TOLIST
	} else {
		//LIVE_TOLIST := Reverse("something") + " ," + Reverse("something else")
		LIVE_TOLIST := ATD
		return LIVE_TOLIST
	}
}

func CLWASendMail(r *http.Request, subject string, finalHTML string) error {
	c := appengine.NewContext(r)
	u := user.Current(c)
	msg := &mail.Message{
		Sender:   u.String() + "<" + u.Email + ">",
		To:       []string{CLWAToAddr(r)},
		Bcc:      []string{BCC_ADDR},
		Subject:  subject + " [CLWACRM]",
		HTMLBody: finalHTML,
	}

	if err := mail.Send(c, msg); err != nil {
		if err_1 := CLWAReportErrorThroughMail(r, subject, finalHTML); err_1 != nil {
			c.Errorf("Couldn't send email: %v", err_1.Error())
			return err_1
		}
		return err
	}
	return nil
}

func CLWAReportErrorThroughMail(r *http.Request, subject string, finalHTML string) error {
	c := appengine.NewContext(r)
	msg := &mail.Message{
		Sender:   ATD,
		To:       []string{ATD},
		Subject:  subject + " [CLWACRM-ERR][CLWACRM]",
		HTMLBody: finalHTML,
	}

	return mail.Send(c, msg)
}
