

from django.urls import path

from . import views

urlpatterns = [
    # name is the unique id for this URL mapping
    # link to home page is like
    # <a href="{% url 'index' %}">Home </a>
    path('', views.index, name = 'index')

]
