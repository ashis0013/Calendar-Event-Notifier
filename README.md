# Calender Event Notifier

This is an event notifier for google calendar built in golang. If you are using calendar in browser you need keep the calendar tab open all day to get nofications. Also default notification is 10 minutes prior which is very early. This tool runs in background and notifies you 3 minutes prior.

This is built using <a href="https://github.com/deckarep/gosx-notifier">gosx-notifier</a> and googleapi bindings for go. To use this app you need to host a project in googleapi services and download a credential.json. The steps can be found here
- <a href="https://developers.google.com/workspace/guides/create-project" >Create a project and enable the API</a>
- <a href="https://developers.google.com/workspace/guides/create-credentials" >Create credentials.</a>

## Usage
Change the file locations in the config file. Use the shell script provided to install the cli.

```shell
git clone https://github.com/ashis0013/Calendar-Event-Notifier.git
cd Calendar-Event-Notifier
./install.sh
```

You can build this project if you have golang installed and add the executable to the path. 

```shell
go build
go install
cp config.cfg $GOPATH/bin # (Optional) Do this if you get "Unable to open config"
```

After installation you can use the cli calendar-notifier. Use -h flag to know about flags

```shell
calendar-notifier -d 3 #Fetches events upto 3 days
calendar-notifier -m 30 #Fetches events for next 30 minutes
calendar-notifier #Fetehces events for today
```