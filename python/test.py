# https://www.youtube.com/watch?v=rCLekTzKUrw&list=RDMMrCLekTzKUrw

from pytube import YouTube

def Download(link):
    youtubeObject = YouTube(link)
    youtubeObject = youtubeObject.streams.get_highest_resolution()
    try:
        youtubeObject.download(saveDir)
    except:
        print("An error has occurred")
    print("Download is completed successfully")


saveDir = "/home/iza/Documents/Video"
link = input("Enter the YouTube video URL: ")
Download("https://www.youtube.com/watch?v=rCLekTzKUrw&list=RDMMrCLekTzKUrw")