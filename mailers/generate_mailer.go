package mailers

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/go-gomail/gomail"
	ottoutils "ottodigital.id/library/utils"
)

// UploadJobMailer ..
func UploadJobMailer(email,url,sender string) {
	bccEmails := strings.Split(email, ", ")
	to := sender

	subject := "OttoSFA - Upload Job Management Report "

	templateEmail := `
	<p>Berikut Link Download Report Upload Job Management :</p>
	
	<a href="` + url + `" download= "report.csv">Report Upload Job Management</a>
	
	<br/>
	<br/>
	<br/>
	
    Terima kasih,<br/>
    Team OttoPay
	`
	err := sendUploadReport(to, bccEmails, subject, templateEmail)
	if err != nil {
		logs.Error("failed send email ----", err.Error())
		return
	}

	logs.Info("Mail sent!")

	// return nil
}

// RegenerateCallPlanMailer ..
func RegenerateCallPlanMailer(email string, firstDate, lastDate time.Time, cycle, limit, salesCallPlan, targetedMerchants, successCall, failedCall int, msgErr string) {
	var templateEmail string
	bccEmails := strings.Split(email, ", ")
	to := "fina.rezalina@ottodigital.id"

	subject := "OttoSFA - ReGenerate Call Plan Report " + ottoutils.GetEnv("EMAIL_ENV", "")
	if msgErr == "" {
		templateEmail = `
		<p>Berikut Report Generate Call Plan yang dilaksanakan pada :</p>
		<ul>
		<li>Tanggal : ` + time.Now().Format("02/01/2006") + `</li>
		<li>Cycle : ` + strconv.Itoa(cycle) + ` Hari</li>
		<li>Periode Call Plan : ` + firstDate.Format("02/01/2006") + ` - ` + lastDate.Format("02/01/2006") + ` </li>
		<li>Jumlah Merchant dalam Cycle Call Plan : ` + strconv.Itoa(targetedMerchants) + ` </li>
		<li>Jumlah Salesman dalam Cycle Call Plan : ` + strconv.Itoa(salesCallPlan) + `</li>
		<li>Jumlah Sukses Generate Call Plan : ` + strconv.Itoa(successCall) + `</li>
		<li>Jumlah Gagal Generate Call Plan : ` + strconv.Itoa(failedCall) + `</li>
		<li>Total Generate Call Plan : ` + strconv.Itoa(successCall+failedCall) + `</li>
		</ul>
		
		Terima kasih,<br/>
		Team OttoPay
		`
	} else {
		templateEmail = `<p> Maaf, terjadi kesalahan ketika Regenerate Callplan : ` + msgErr + ` </p>
		<p>Silahkan coba beberapa saat lagi.</p>
		
		
		Terima kasih,<br/>
		Team OttoPay`
	}
	err := sendCallPlanReport(to, bccEmails, subject, templateEmail)
	if err != nil {
		logs.Error("failed send email ----", err.Error())
		// return err
	}

	logs.Info("Mail sent!")

	// return nil
}

// sendCallPlanReport ..
func sendCallPlanReport(to string, bcc []string, subject, message string) error {
	configSMTP := beego.AppConfig.DefaultString("email.smtp.address", "smtp.gmail.com")
	configSMTPPort := beego.AppConfig.DefaultInt("email.smtp.port", 587)
	configEmailSender := beego.AppConfig.DefaultString("email.sender", "ottopay@ottopay.id")
	configPassSender := beego.AppConfig.DefaultString("email.password", "zbnadctvllmukhbi")

	mailer := gomail.NewMessage()

	mailer.SetHeaders(map[string][]string{
		"From": {configEmailSender},
		// "To":      {to},
		"Bcc":     bcc,
		"Subject": {subject},
	})

	mailer.SetBody("text/html", message)

	dialer := gomail.NewDialer(configSMTP, configSMTPPort, configEmailSender, configPassSender)
	// fmt.Println("this is the dialer => ", dialer)
	fmt.Println("this is the mailer => ", mailer)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		fmt.Println("error while sending mail ==> ", string(err.Error()))
		log.Fatal(err.Error())
		return err
	}

	return nil
}

// sendUploadReport ..
func sendUploadReport(to string, bcc []string, subject, message string) error {
	configSMTP := beego.AppConfig.DefaultString("email.smtp.address", "smtp.gmail.com")
	configSMTPPort := beego.AppConfig.DefaultInt("email.smtp.port", 587)
	configEmailSender := beego.AppConfig.DefaultString("email.sender", "ottopay@ottopay.id")
	configPassSender := beego.AppConfig.DefaultString("email.password", "zbnadctvllmukhbi")

	mailer := gomail.NewMessage()
	fmt.Println("sender :" ,to)
	mailer.SetHeaders(map[string][]string{
		"From": {configEmailSender},
		"To":      {to},
		// "Bcc":     bcc,
		"Subject": {subject},
	})

	mailer.SetBody("text/html", message)

	dialer := gomail.NewDialer(configSMTP, configSMTPPort, configEmailSender, configPassSender)
	// fmt.Println("this is the dialer => ", dialer)
	fmt.Println("this is the mailer => ", mailer)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		fmt.Println("error while sending mail ==> ", string(err.Error()))
		log.Println(err.Error())
		return err
	}

	return nil
}
