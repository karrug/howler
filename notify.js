"use strict";

var oldMsg, newMsg;

function enableNotifications() {
  console.log(Notification.permission)
  if (Notification.permission === "granted") {
    pollMessages();
  } else if (Notification.permission !== "denied") {
    Notification.requestPermission().then(function (permission) {
      if (permission === "granted") {
        pollMessages();
      }
    });
  }
}

function getMessage() {
  var xhttp = new XMLHttpRequest();
  xhttp.onreadystatechange = function() {
    if (this.readyState == 4 && this.status == 200) {
      newMsg = xhttp.responseText;
      if (newMsg !== oldMsg) {
        var notification = new Notification(newMsg);
        oldMsg = newMsg;
      }
    }
  };
  xhttp.open("GET", "http://192.168.1.2:8050/g");
  xhttp.send()
}


function pollMessages() {
  setInterval(getMessage, 15000);
}

enableNotifications();
