#!/bin/bash

PLIST=~/Library/LaunchAgents/weistroff.marc.galoxy.plist
go build
cp galoxy /usr/local/bin/galoxy

cat <<EOF > $PLIST
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
  <dict>
    <key>Label</key>
    <string>weistroff.marc.galoxy</string>
    <key>RunAtLoad</key>
    <true/>
    <key>KeepAlive</key>
    <true/>
    <key>UserName</key>
    <string>marc</string>
    <key>ProgramArguments</key>
    <array>
        <string>/usr/local/bin/galoxy</string>
    </array>
    <key>WorkingDirectory</key>
    <string>/usr/local</string>
  </dict>
</plist>
EOF

launchctl unload $PLIST
launchctl load $PLIST
