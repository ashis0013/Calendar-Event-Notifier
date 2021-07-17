# Calender Event Notifier

This is an event notifier for google calendar built in golang. If you are using calendar in browser you need keep the calendar tab open all day to get nofications. Also default notification is 10 minutes prior which is very early. This tool runs in background and notifies you 3 minutes prior.

This is built using <a href="https://github.com/deckarep/gosx-notifier">gosx-notifier</a> and googleapi bindings for go. To use this app you need to host a project in googleapi services and download a credential.json. The steps can be found here
- <a href="https://developers.google.com/workspace/guides/create-project" >Create a project and enable the API</a>
- <a href="https://developers.google.com/workspace/guides/create-credentials" >Create credentials.</a>