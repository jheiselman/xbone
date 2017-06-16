# xbone
Library for controlling an Xbox One

This library (and associated command) will send a command to an Xbox One over the local network to wake it up. To make use of it:

* [The Xbox Power Mode setting must be set to 'Instant On'](https://support.xbox.com/en-US/xbox-one/console/change-power-settings)
* [The environment variable XBOX_LIVE_ID must be set or the Live Device ID should be supplied as a command line option](https://support.xbox.com/en-US/my-account/warranty-and-service/find-xbox-one-kinect-serial-number)

# TODO
This library can currently only power on an Xbox One. Further interactions require an authenticated session with the device which hasn't been worked out. It is unlikely that I'll ever get around to doing that since powering on the device was my only goal. Contributions accepted if you get it worked out.

# Inspiration
The method for packing the network request came from [Scheamper/xbox-remote-power](https://github.com/Schamper/xbox-remote-power)
