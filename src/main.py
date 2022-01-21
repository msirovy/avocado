#!/usr/bin/env python3


import RPi.GPIO as GPIO
import time

def callback(channel):
    if GPIO.input(channel):
        print("LED off")
    else:
        print("LED on")

# Set our GPIO numbering to BCM
GPIO.setmode(GPIO.BCM)

# Define the GPIO pin that we have our digital output from our sensor connected to
channel = 17
# Set the GPIO pin to an input
GPIO.setup(channel, GPIO.IN)


def status(hodnota):
    if hodnota > 0:
        print("                           HI  %%%%%%%%%%%%%%%%%%%%%%%%%%%%")
    else:
        print("%%%%%%%%%%%%%%%%%%%%%%%%  LOW ")


# This is an infinte loop to keep our script running
while True:
    try:
        # This line simply tells our script to wait 0.1 of a second, this is so the script doesnt hog all of the CPU
        time.sleep(1)
        status(GPIO.input(channel))

    except KeyboardInterrupt:
        print("Stopped by user")

