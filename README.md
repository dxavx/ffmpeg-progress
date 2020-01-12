# ffmpeg-progress

Reads data from a TCP connection about the progress of the ffmpeg encoding process

1. ```go run main.go``` - application will start listening to the 9090 port, without this step running ````ffmpeg```` will give an error.
2. ```ffmpeg -progress tcp://127.0.0.1:9090 -i input.mov -c:v libx264 -preset veryslow -crf 15 out.mov```
##### the duration of the material is used to calculate the completion percentage. In your development use the structure output by ffprobe to get the duration of the material
