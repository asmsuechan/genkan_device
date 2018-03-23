# Genkan Device
This is genkan on Raspberry Pi zero wh.

# Struture
WIP: Write the role of golang files `daemon/updateManager`, `deployment/pushNewVersion` and `main.go`

# Installation
1. Add `.genkan.env` to `/home/pi`
2. Deploy `genkan_device` and `updateManager` to `/home/pi`
3. Move `*.service` files to `/etc/systemd/system/`
4. Enable daemons by `sudo systemctl start genkan` and `sudo systemctl start genkan-updater`
