#! /bin/bash
# Download the specific releace
cd /home/pi
curl -sJLO https://github.com/asmsuechan/genkan_device/releases/download/v_$1/genkan_device_$1_linux_arm.tar.gz
tar -zxvf genkan_device_$1_linux_arm.tar.gz
mv genkan_device_$1_linux_arm/genkan_device /home/pi

# Run the downloaded program and pi-blaster
sudo systemctl restart genkan
rm -rf genkan_device_$1_linux_arm/ genkan_device_$1_linux_arm.tar.gz
