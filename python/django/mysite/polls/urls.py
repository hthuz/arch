
from django.urls import path

from . import views

# Namespace
app_name = "polls"

urlpatterns = [
        # /polls/
        path("",views.index, name="index"),

        # /polls/5/
        # Using angle brackets to "capture" part of url
        # int: converter
        # 'name' is called by {% url %} tag in template html
        path("<int:question_id>/",views.detail,name="detail"),

        # /polls/5/results/
        path("<int:question_id>/results/",views.results, name="results"),

        # /polls/5/vote/
        path("<int:question_id>/vote/",views.vote, name="vote"),

]
