#!/bin/sh

while true
do
  time=$(date "+%H:%M %x")
  acpi=$(acpi)
  battery=${acpi:24:3}

  xsetroot -name "$battery | $time"

  # Depending on the time when system boots. The error of time is 60 seconds
  sleep 60
done
