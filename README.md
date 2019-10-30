
# [Audio Recognition](https://www.acrcloud.com/music-recognition) -- File Scan Tool (Golang)



## Overview
  [ACRCloud](https://www.acrcloud.com/) provides [Automatic Content Recognition](https://www.acrcloud.com/docs/introduction/automatic-content-recognition/) services for [Audio Fingerprinting](https://www.acrcloud.com/docs/introduction/audio-fingerprinting/) based applications such as **[Audio Recognition](https://www.acrcloud.com/music-recognition)** (supports music, video, ads for both online and offline), **[Broadcast Monitoring](https://www.acrcloud.com/broadcast-monitoring)**, **[Second Screen](https://www.acrcloud.com/second-screen-synchronization)**, **[Copyright Protection](https://www.acrcloud.com/copyright-protection-de-duplication)** and etc.<br>
  
  This tool can scan audio/video files and detect audios you want to recognize such as music, ads.

  Supported Format:
  
```  
 Audio: mp3, wav, m4a, flac, aac, amr, ape, ogg ...<br>
 Video: mp4, mkv, wmv, flv, ts, avi ...
```


## Usage

Download the latest execution file from [release page](https://github.com/acrcloud/acrcloud_scan_files_go/releases)

Before you use this tool, you must have acrcloud host, access_key and access_secret.

If you haven't have these, you can [register one](https://console.acrcloud.com/signup).

**Rename config.yaml.example to config.yaml and fill in host, key and secret**
```diff
-  host:
-  access_key:
-  access_secret:
+  host: HOST ADDRESS in your console
+  access_key:  ACCESS KEY in your console
+  access_secret: ACCESS SECRET in your console 
```

**Notice**: the libacrcloud_extr_tool.dylib file must be in the same directory with main file

```
GLOBAL OPTIONS:
   --mode MODE, -m MODE        MODE: local, network (default: "local")
   --type TYPE, -t TYPE        TYPE: folder, file (default: "file")
   --filename PATH, -f PATH    PATH: the file need to scan
   --url URL, -u URL           URL: the network file you want to scan (when using network mode)
   --output OUTPUT, -o OUTPUT  OUTPUT: the directory to save the results
   --filter FILTER, -l FILTER  FILTER: combine, fuzzy
   --help, -h                  show help
   --version, -v               print the version
```


Scan a file and save the result:

```
./main scan -f ~/media_files/test.mp4 -o ~/media_files/report.csv
```

## Feature

-   [x] Scan file
-   [x] Scan network file
-   [x] Export the report
-   [ ] Scan folder
-   [ ] Custom report fields
-   [ ] Filter report result
