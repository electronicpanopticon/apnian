# Apnian

[![Build Status](https://api.travis-ci.com/electronicpanopticon/apnian.go.svg?branch=master)](https://travis-ci.com/electronicpanopticon/apnian.go)
[![Coverage Status](https://coveralls.io/repos/github/electronicpanopticon/apnian.go/badge.svg?branch=master)](https://coveralls.io/github/electronicpanopticon/apnian.go?branch=master)

## About

Apnian is a wrapper around the sideshow/apns2 library. 

## Usage

### Step 1
Create an Apnian config file and place in of of the following four places: `.`, `..`, `$HOME`, or `$GOPATH\config`.

Here's an example called `apnian.yaml`:

```yaml
P8KeyName: YourP8KeyFileName.p8
Topic: TheBundleIdentifierYourSendingNotificationsTo
APNSKeyID: YourAPNSKeyID
TeamID: YourTeamID
```

### Step 2
Call Apnian from your code:

```go
package main

import (
    "fmt"
    "github.com/electronicpanopticon/apnian.go"
    "log"
)

func main() {
    apn, err := apnian.New("apnian")
    deviceID := "123456"
    apsMessage := apnian.GenerateAPS("Bawk! Bawk! Bawk! 🐔", "default", "https://electronicpanopticon.com")
    response, err := apn.Push(deviceID, apsMessage)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("%v %v %v\n", response.StatusCode, response.ApnsID, response.Reason)
}
```

### Step 4

Profit!

## TODO

* More expressive APS messages 

## Libraries

* [APNS/2](https://github.com/sideshow/apns2)
* [go-homedir](https://github.com/mitchellh/go-homedir)
* [Goveralls](https://github.com/mattn/goveralls)

## Reference

* [Local and Remote Notification Programming Guide](https://developer.apple.com/library/archive/documentation/NetworkingInternet/Conceptual/RemoteNotificationsPG/APNSOverview.html)
* [Sending Notification Requests to APNs](https://developer.apple.com/documentation/usernotifications/setting_up_a_remote_notification_server/sending_notification_requests_to_apns)
* [Creating the Remote Notification Payload](https://developer.apple.com/library/archive/documentation/NetworkingInternet/Conceptual/RemoteNotificationsPG/CreatingtheNotificationPayload.html)

## TODO

* [How to display image in ios push notification?](https://stackoverflow.com/questions/37839171/how-to-display-image-in-ios-push-notification)


[![Via Appia](files/1083px-Appian_Way.jpg "Via Appia")](https://en.wikipedia.org/wiki/Appian_Way)
