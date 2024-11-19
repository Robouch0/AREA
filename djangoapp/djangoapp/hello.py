from django.http import HttpRequest, HttpResponse
import time

def index(_: HttpRequest):
    t = time.localtime()
    current_time = time.strftime("%H:%M:%S", t)
    return HttpResponse(f"Current time is: {current_time}")
