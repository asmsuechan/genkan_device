# Genkan Device
This is genkan on Raspberry Pi zero wh.

# Struture
WIP: Write the role of golang files `daemon/updateManager`, `deployment/pushNewVersion` and `main.go`

# Installation
1. Move `updater.sh` to `/home/pi`
2. Add `.genkan.env` to `/home/pi`
3. Put `genkan_device` and `updateManager` to `/home/pi`
4. Move `*.service` files to `/etc/systemd/system/`
5. Enable daemons by `sudo systemctl enable genkan # or genkan-updater` and start it by `sudo systemctl start genkan # or genkan-updater`

# Release
TravisCI releases every job automatically only if the commit has tag.

```
$ git add .
$ git commit -m "commit message"
$ git tag v_1.0.0
$ git push origin v_1.0.0
```

# systemd
`genkan.service` and `genkan-updater.service` is files for systemd. They register the 2 jobs as daemons.

They use `.genkan.env` as an environment file. Don't forget to put it properly because it's very hard to debug what happens if it's forgotten...

## Detailed
Cross-compile is realized by `goxc` and the compiled binary is uploaded to GitHub Releases.

After upload, a MQTT message will be published to `genkan/update` with version number as a payload by `pushNewVersion`.
