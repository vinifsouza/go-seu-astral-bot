ffmpeg -i ./5seconds.mp3 -f s16le -ar 48000 -ac 2 pipe:1 | ./dca > ../assets/5seconds.dca
