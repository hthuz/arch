
from django.urls import path

from . import views

# Namespace
app_name = "polls"

urlpatterns = [
        # /polls/
        path("",views.IndexView.as_view(), name="index"),

        # /polls/5/
        # Using angle brackets to "capture" part of url
        # int: converter
        # 'name' is called by {% url %} tag in template html
        path("<int:pk>/",views.DetailView.as_view(),name="detail"),

        # /polls/5/results/
        path("<int:pk>/results/",views.ResultsView.as_view(), name="results"),

        # /polls/5/vote/
        path("<int:question_id>/vote/",views.vote, name="vote"),

]
