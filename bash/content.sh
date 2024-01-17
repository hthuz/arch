#!/bin/bash

block=1123
if grep -q "block #[0-9]* not found" ./content.txt;
then
    block=$((block))
else
    block=$((block+1))
fi

echo ${block}

