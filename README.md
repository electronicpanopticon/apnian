# Apnian

[![Build Status](https://api.travis-ci.com/electronicpanopticon/apnian.go.svg?branch=master)](https://travis-ci.com/electronicpanopticon/apnian.go)
[![Coverage Status](https://coveralls.io/repos/github/electronicpanopticon/apnian.go/badge.svg?branch=master)](https://coveralls.io/github/electronicpanopticon/apnian.go?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/electronicpanopticon/apnian.go)](https://goreportcard.com/report/github.com/electronicpanopticon/apnian.go)
[![License](https://img.shields.io/github/license/electronicpanopticon/apnian.go)](https://github.com/electronicpanopticon/apnian.go/blob/master/LICENSE)
[![Release](https://img.shields.io/github/v/release/electronicpanopticon/apnian.go.svg?style=flat-square)](https://github.com/electronicpanopticon/apnian.go/releases)

## About

Apnian is a wrapper around the [sideshow/apns2](https://github.com/sideshow/apns2) Apple Push Notification Service library. 
APNs allows you to send messages to your Apple devices such as an iPhone or Apple Watch. 

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

Place your Apple provided p8 file under $GOPATH/keys

### Step 3
Call Apnian from your code:

```
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

### Step 5

Profit!

## TODO

* More expressive APS messages 

## Libraries

* [APNS/2](https://github.com/sideshow/apns2)
* [go-homedir](https://github.com/mitchellh/go-homedir)
* [gobrick](https://github.com/electronicpanopticon/gobrick)
* [Goveralls](https://github.com/mattn/goveralls)
* [Viper](https://github.com/spf13/viper)

## Reference

* [Local and Remote Notification Programming Guide](https://developer.apple.com/library/archive/documentation/NetworkingInternet/Conceptual/RemoteNotificationsPG/APNSOverview.html)
* [Sending Notification Requests to APNs](https://developer.apple.com/documentation/usernotifications/setting_up_a_remote_notification_server/sending_notification_requests_to_apns)
* [Creating the Remote Notification Payload](https://developer.apple.com/library/archive/documentation/NetworkingInternet/Conceptual/RemoteNotificationsPG/CreatingtheNotificationPayload.html)

## TODO

* [How to display image in ios push notification?](https://stackoverflow.com/questions/37839171/how-to-display-image-in-ios-push-notification)


[![Via Appia](files/1083px-Appian_Way.jpg "Via Appia")](https://en.wikipedia.org/wiki/Appian_Way)
