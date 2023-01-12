#!/bin/sh

while true
do
  time=$(date "+%H:%M %x")
  acpi=$(acpi)
  battery=${acpi:24:3}

  xsetroot -name "$battery | $time"
  sleep 60
done
