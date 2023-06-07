from django.shortcuts import render

from .models import File

# Create your views here.

def index(request):
    num_files = File.objects.all().count()

    return render(request, 'index.html', context = {
        'num_files': num_files
    })
