package core

import (
	"html/template"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/JNSAPH/macos_raidmon/mail"
	"github.com/JNSAPH/macos_raidmon/structs"
	"github.com/JNSAPH/macos_raidmon/utils"
	"github.com/dustin/go-humanize"
	"github.com/getlantern/systray"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
)

func onReady() {
	// Get Icons
	driveIco := utils.GetDriveIco()
	degradedDriveIco := utils.GetDegradedDriveIco()
	emailSender := mail.NewEmailSender(
		AppConfig.Mail.Smtp.Host,
		AppConfig.Mail.Smtp.Port,
		AppConfig.Mail.Smtp.Username,
		AppConfig.Mail.Smtp.Password,
		AppConfig.Mail.Sender,
	)

	// Set up systray
	systray.SetIcon(driveIco)
	systray.AddMenuItem("RaidMON is running", "Raid Monitor is running")
	systray.AddSeparator()
	mQuit := systray.AddMenuItem("Quit", "Exit the application")

	// Get working directory
	wd, err := os.Getwd()
	if err != nil {
		logrus.Fatalf("Error getting working directory. Couldn't send email: ", err)
	}

	var foundDegraded []structs.AppleRAIDSet
	var lastEmailSent time.Time

	// Check if the drive is degraded
	go func() {
		for {
			details := utils.GetRaidDetails()
			if len(details.AppleRAIDSets) > 0 {
				logrus.Infof("Found %d RAID sets", len(details.AppleRAIDSets))

				// For each RAID set
				for _, raidSet := range details.AppleRAIDSets {
					logrus.Debugf("Checking RAID set: %s with status: %s", raidSet.Name, raidSet.Status)
					if raidSet.Status != "Online" {
						logrus.Warnf("RAID set %s is degraded", raidSet.Name)
						foundDegraded = append(foundDegraded, raidSet)
					}
				}
			} else {
				logrus.Info("No RAID sets found")
			}

			// log length of foundDegraded
			logrus.Debugf("Found %d degraded RAID sets", len(foundDegraded))

			// One or more degraded RAID sets found
			if len(foundDegraded) > 0 {
				systray.SetIcon(degradedDriveIco)

				// Send Mail if enabled
				if AppConfig.Mail.SendMail {
					// Check if last email was sent over X ago
					if lastEmailSent.IsZero() || time.Since(lastEmailSent).Minutes() >= float64(AppConfig.Mail.MaxSendEvery) {
						logrus.Debug("Sending email: either no previous email was sent or the last one was sent over an hour ago.")

						// Get E-Mail Template
						templatePath := filepath.Join(wd, "templates", "degraded.html")

						tmpl, err := template.New("degraded.html").Funcs(template.FuncMap{
							"lower": strings.ToLower,
						}).ParseFiles(templatePath)
						if err != nil {
							logrus.Error("Error parsing template file. Couldn't send email: ", err)
							for _, recipient := range AppConfig.Mail.Recipients {
								emailSender.SendErrorMail(recipient, err)
							}
							break
						}

						// Execute Template
						var renderedTemplate string
						outputBuffer := &structs.TemplateBuffer{
							Buffer: &renderedTemplate,
						}
						err = tmpl.Execute(outputBuffer, map[string]interface{}{
							"Sets": foundDegraded,
						})
						if err != nil {
							logrus.Error("Error executing template. Couldn't send email: ", err)
							for _, recipient := range AppConfig.Mail.Recipients {
								emailSender.SendErrorMail(recipient, err)
							}
							break
						}

						for _, recipient := range AppConfig.Mail.Recipients {
							logrus.Info("Sending email to ", recipient)
							err := emailSender.SendEmail(recipient, "⚠ IMPORTANT: Your RAID Sets State Changed!", renderedTemplate)
							if err != nil {
								logrus.Error("Error sending email: ", err)
								for _, recipient := range AppConfig.Mail.Recipients {
									emailSender.SendErrorMail(recipient, err)
								}
								break
							}
						}

						lastEmailSent = time.Now()
					} else {
						logrus.Debugf("Last email was sent less than %d minutes ago. Skipping email.", AppConfig.Mail.MaxSendEvery)
					}
				} // end of send mail
			} else {
				// No degraded RAID sets found
				logrus.Debug("No degraded RAID sets found")
				systray.SetIcon(driveIco)
			}

			// Reset foundDegraded
			foundDegraded = nil

			time.Sleep(time.Duration(AppConfig.Config.CheckValue) * time.Second)
		}
	}()

	// Handle quit menu item click
	go func() {
		for {
			select {
			case <-mQuit.ClickedCh:
				systray.Quit()
			}
		}
	}()

	go func() {
		// Register Chrono for daily report
		if AppConfig.Mail.SendMail {
			logrus.Infof("Registering daily report cron: %s", AppConfig.Mail.DailyReportChron)

			c := cron.New()
			_, err := c.AddFunc(AppConfig.Mail.DailyReportChron, func() {
				// Vars
				var overallStatus string = "Healthy"
				var externalDrives []structs.Disk

				// Chrono Code here pls uwu
				logrus.Info("Running daily report")

				// Get Drive Details
				driveDetails := utils.GetDriveDetails()
				raidDetails := utils.GetRaidDetails()

				// Check Overall Status
				for _, raidSet := range raidDetails.AppleRAIDSets {
					if raidSet.Status != "Online" {
						if raidSet.Status == "Degraded" {
							overallStatus = "Dangerous"
						} else {
							overallStatus = "Problematic"
						}
						break
					}
				}

				// Get Drive Details
				for _, drive := range driveDetails.AllDisksAndPartitions {
					if utils.IsExternalDisk(drive) {
						externalDrives = append(externalDrives, drive)
					}
				}

				// Get E-Mail Template
				templatePath := filepath.Join(wd, "templates", "dailyreport.html")

				tmpl, err := template.New("dailyreport.html").Funcs(template.FuncMap{
					"lower":    strings.ToLower,
					"humanize": humanize.Bytes,
				}).ParseFiles(templatePath)
				if err != nil {
					logrus.Error("Error parsing template file. Couldn't send email: ", err)
					for _, recipient := range AppConfig.Mail.Recipients {
						emailSender.SendErrorMail(recipient, err)
					}
					return
				}

				// Execute Template
				var renderedTemplate string
				outputBuffer := &structs.TemplateBuffer{
					Buffer: &renderedTemplate,
				}

				err = tmpl.Execute(outputBuffer, map[string]interface{}{
					"OverallStatus": overallStatus,
					"Drives":        externalDrives,
					"Raids":         raidDetails.AppleRAIDSets,
				})
				if err != nil {
					logrus.Error("Error executing template. Couldn't send email: ", err)
					for _, recipient := range AppConfig.Mail.Recipients {
						emailSender.SendErrorMail(recipient, err)
					}
					return
				}

				for _, recipient := range AppConfig.Mail.Recipients {
					logrus.Info("Sending email to ", recipient)
					err := emailSender.SendEmail(recipient, "Daily RAID Report", renderedTemplate)
					if err != nil {
						logrus.Error("Error sending email: ", err)
						for _, recipient := range AppConfig.Mail.Recipients {
							emailSender.SendErrorMail(recipient, err)
						}
						break
					}
				}
			})

			// Check for errors
			if err != nil {
				logrus.Error("Error adding daily report cron: ", err)
			}
			c.Start()

		}
	}()
}
