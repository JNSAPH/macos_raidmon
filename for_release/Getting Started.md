# Getting Started Guide for MOSS RaidMON

MOSS RaidMON is a tool to monitor the health of your RAID array and alert you in case of any issues. This guide will help you set up MOSS RaidMON on your system.

## Configure RaidMON
Open the `config/config.yml` file in a text editor and update the necessary fields. The Configuration file has documentation for each field to help you set up the tool.

## (Optional) Customize Notifications
You can customize the email template by editing the files in the `templates` folder. 

## Start Monitoring
Before adding RaidMON to the system startup, i suggest you run the tool once to ensure everything is working as expected. You can start RaidMON by running the following command in the terminal:

```bash
./moss_raidmon
```

## Enable Auto-Start
To enable RaidMON to start on boot, run the following command:

```bash
./install.sh
```

## Done!
MOSS RaidMON is now set up and ready to monitor your drives. You will receive email notifications in case of any issues with your RAID array. RAIDMon adds a Icon to your system tray to show the status of your RAID array:
- 💾: RAIDMon is started and there are no Problems.
- ⚠: RAIDMon is started and there are Problems with your RAID array. You should check your email for more information.

# To Remove RaidMON
To remove RaidMON from your system, run the following command:
```bash
./uninstall.sh
```