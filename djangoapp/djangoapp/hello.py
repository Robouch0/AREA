from django.http import HttpRequest, HttpResponse
import time

from djangoapp.models import Choice, Question


def index(_: HttpRequest):
    t = time.localtime()
    current_time = time.strftime("%Y:%H:%M:%S", t)
    obj = Question.objects.create(question_text="hello")
    return HttpResponse(f"Current time is: {current_time} obj created: {obj}")
