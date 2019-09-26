#! /bin/bash
# Build the UI
cd /Users/jjmeg/go/src/videoserver/web
go install
cp /Users/jjmeg/go/bin/web /Users/jjmeg/go/bin/video_server_web_ui/web
cp -R /Users/jjmeg/go/src/videoserver/templates /Users/jjmeg/go/bin/video_server_web_ui/