# https://www.youtube.com/watch?v=rCLekTzKUrw&list=RDMMrCLekTzKUrw

import sys
import os
from pytube import YouTube

def Download(link):
    youtubeObject = YouTube(link)
    stream = youtubeObject.streams.filter(file_extension='mp4', progressive=True).first()
    download_path = 'downloads'
    os.makedirs(download_path, exist_ok=True)
    file_path = os.path.join(download_path, stream.title + ".mp4")
    try:
        stream.download(download_path)
        return file_path
    except:
        print("An error has occurred")
    print("Download is completed successfully")

if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Usage: python youtube_downloader.py <video_url>")
        sys.exit(1)
        
    url = sys.argv[1]
    result = Download(url)
    print(result)