#!/bin/bash

# Define paths
DEST_DIR="$HOME/Library/Application Support/macos_raidmon"
PLIST_PATH="$HOME/Library/LaunchAgents/"
PLIST_NAME="com.jnsaph.raidmon.plist"
EXECUTABLE_NAME="moss_raidmon"
EXECUTABLE_PATH="$DEST_DIR/$EXECUTABLE_NAME"

# Create destination directory if it doesn't exist
mkdir -p "$DEST_DIR"
mkdir -p "$PLIST_PATH"


# Copy all files from the release folder
cp -R "$(dirname "$0")"/* "$DEST_DIR"

# Ensure the executable has the right permissions
chmod +x "$EXECUTABLE_PATH"

# Create the plist file
cat <<EOF > "$PLIST_PATH/$PLIST_NAME"
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
    <key>Label</key>
    <string>com.jnsaph.raidmon</string>

    <key>ProgramArguments</key>
    <array>
        <string>$EXECUTABLE_PATH</string>
    </array>

    <key>WorkingDirectory</key>
    <string>$DEST_DIR</string>

    <key>RunAtLoad</key>
    <true/>

    <key>KeepAlive</key>
    <true/>
</dict>
</plist>
EOF

# Load the Launch Agent
launchctl unload "$PLIST_PATH" 2>/dev/null
launchctl load "$PLIST_PATH"

echo "Autostart setup complete. Files copied to: $DEST_DIR"
