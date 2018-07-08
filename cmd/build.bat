call render.bat
robocopy "../client" "../bin/client/" /mir
go build ../
move computingfun.org.exe "../bin/"
