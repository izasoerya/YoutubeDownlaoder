# https://www.youtube.com/watch?v=rCLekTzKUrw&list=RDMMrCLekTzKUrw

import sys
from pytube import YouTube

def Download(link):
    youtubeObject = YouTube(link)
    stream = youtubeObject.streams.filter(file_extension='mp4', progressive=True).first()
    try:
        stream.download('downloads')
    except:
        print("An error has occurred")
    print("Download is completed successfully")

if __name__ == "__main__":
    savePath = "/home/iza/Videos"
    url = sys.argv[1]
    Download(url)