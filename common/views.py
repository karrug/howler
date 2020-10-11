from django.shortcuts import render
from django.http import HttpResponse


def get_room(request):
    return render(request, "room.html")


def get_message(request):
    message = []
    return message


def set_message(request):
    return HttpResponse(status=201)
