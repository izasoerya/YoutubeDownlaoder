# https://www.youtube.com/watch?v=rCLekTzKUrw&list=RDMMrCLekTzKUrw

import sys
from pytube import YouTube

def Download(link, savePath):
    youtubeObject = YouTube(link)
    youtubeObject = youtubeObject.streams.get_highest_resolution()
    try:
        youtubeObject.download(savePath)
    except:
        print("An error has occurred")
    print("Download is completed successfully")

if __name__ == "__main__":
    savePath = "/home/iza/Documents/Video"
    url = sys.argv[1]
    Download(url, savePath)