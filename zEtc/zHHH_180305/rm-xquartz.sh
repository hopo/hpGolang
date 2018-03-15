launchctl unload /Library/LaunchAgents/org.macosforge.xquartz.startx.plist && \
sudo launchctl unload /Library/LaunchDaemons/org.macosforge.xquartz.privileged_startx.plist && \
sudo rm -rf /opt/X11* /Library/Launch*/org.macosforge.xquartz.* /Applications/Utilities/XQuartz.app /etc/*paths.d/*XQuartz  && \
sudo pkgutil --forget org.macosforge.xquartz.pkg  && \
rm -rf ~/.serverauth* && rm -rf ~/.Xauthorit* && rm -rf ~/.cache && rm -rf ~/.rnd && \
rm -rf ~/Library/Caches/org.macosforge.xquartz.X11 && rm -rf ~/Library/Logs/X11
