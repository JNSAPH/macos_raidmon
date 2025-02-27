#!/bin/bash

# Define paths
DEST_DIR="$HOME/Library/Application Support/macos_raidmon"
PLIST_PATH="$HOME/Library/LaunchAgents/"
PLIST_NAME="com.jnsaph.raidmon.plist"
EXECUTABLE_NAME="moss_raidmon"
EXECUTABLE_PATH="$DEST_DIR/$EXECUTABLE_NAME"

# Unload the Launch Agent if loaded
launchctl unload "$PLIST_PATH/$PLIST_NAME" 2>/dev/null

# Remove the plist file
rm -f "$PLIST_PATH/$PLIST_NAME"

# Remove the executable and the destination directory
rm -rf "$DEST_DIR"

# Provide confirmation message
echo "Uninstallation complete. Files removed from: $DEST_DIR and $PLIST_PATH"
