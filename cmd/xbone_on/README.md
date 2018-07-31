# xbone
Command line interface the github.com/jheiselman/xbone library

This command can be used to power-on an Xbox One over the local network. To make use of it:

* [The Xbox Power Mode setting must be set to 'Instant On'](https://support.xbox.com/en-US/xbox-one/console/change-power-settings)
* [The environment variable XBOX_LIVE_ID must be set or the Live Device ID should be supplied as a command line option](https://support.xbox.com/en-US/my-account/warranty-and-service/find-xbox-one-kinect-serial-number)

# EXAMPLE
```
$ xbone_on -hostname 192.168.100.234 -liveid FA100CD2919
```

or

```
$ XBOX_LIVE_ID=FA100CD2919 xbone_on -hostname 192.168.100.234
```

or for Windows

```
C:\> xbone_on -hostname 192.168.100.234 -liveid FA100CD2919
```

or

```
C:\> set XBOX_LIVE_ID=FA100CD2919
C:\> xbone_on -hostname 192.168.100.234
```

