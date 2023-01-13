#!/bin/sh

while true
do
  time=$(date "+%H:%M %x")
  acpi=$(acpi)
  battery=${acpi:24:3}
  volume=$(pamixer --get-volume-human)

  xsetroot -name "V $volume | B $battery | $time"

  # Depending on the time when system boots. The error of time is 60 seconds
  sleep 1
done
