#!/bin/bash -e
killall present || true
present presentation.slide &
xdg-open "http://127.0.0.1:3999"
