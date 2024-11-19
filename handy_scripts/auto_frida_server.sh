#!/bin/bash

# The script first checks if any devices are connected using adb devices. If there are no devices, it exits and prints an error message.
# If a device is connected, the script selects the first available device from the list (if you need to connect to a specific device, you can modify the logic to match a particular device).
# The rest of the script proceeds as usual by downloading, extracting, and pushing frida-server to the device.

# choose a version from https://github.com/frida/frida/releases/
# use arm if you're debugging arm apps, via houdini or native bridge (ndk)
HOST_ARCH=x86_64
#HOST_ARCH=architecture
# HOST_ARCH=arm
GUEST_SYS=android
FRIDA_RELEASE=frida-server

# Get the list of devices connected via ADB
DEVICE_LIST=$(adb devices | grep -w "device" | awk '{print $1}')

# Check if there are any connected devices
if [ -z "$DEVICE_LIST" ]; then
    echo "No devices found. Please connect a device via ADB."
    exit 1
fi

# Select the first available device
DEVICE=$(echo "$DEVICE_LIST" | head -n 1)
echo "Using device: $DEVICE"

# Get the download link for the frida-server version matching the architecture and OS
FRIDA_RELEASES=($(curl -s https://github.com/frida/frida/releases | grep -Po "(?<=\<a\ href\=\")(\/frida\/frida\/releases\/download\/\d+\.\d.\d+\/${FRIDA_RELEASE}-\d+\.\d+.\d+-${GUEST_SYS}-${HOST_ARCH}.xz)(?=\"\ )"))

RELEASE_LINK="https://github.com${FRIDA_RELEASES[0]}"
# Uncomment the next line to manually set a specific version if needed
# RELEASE_LINK='https://github.com/frida/frida/releases/download/15.0.8/frida-server-15.0.8-android-x86_64.xz'

# Create the frida directory
mkdir -p ./frida
cd ./frida

# Download the frida-server release
wget "${RELEASE_LINK}"

# Extract the .xz file
unxz -d "$(basename "${RELEASE_LINK}")"

# Rename the extracted frida-server binary
find -name "${FRIDA_RELEASE}-*.*.*-${GUEST_SYS}-${HOST_ARCH}" | xargs -i mv "{}" frida-server

# Push the frida-server binary to the device
adb -s "$DEVICE" push frida-server /data/local/tmp/frida-server

# Change the permissions of the frida-server binary
adb -s "$DEVICE" shell "su -c chmod 755 /data/local/tmp/frida-server"

# Start the frida-server on the device
adb -s "$DEVICE" shell "/data/local/tmp/frida-server"
