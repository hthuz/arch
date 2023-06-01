from django.shortcuts import render, get_object_or_404
from django.http import HttpResponseRedirect
from django.urls import reverse
from django.views import generic


from .models import Question,Choice

# Create your views here.

# Generic ListView: display a list of objects
class IndexView(generic.ListView):
    template_name = "polls/index.html"
    context_object_name = "latest_question_list"
    
    def get_queryset(self):
        return Question.objects.order_by("-pub_date")[:5]

# Generic DetailView: display a detail page or a particular type of object
# It expects primary key value captured from URL to be called "pk"
class DetailView(generic.DetailView):
    model = Question
    template_name = "polls/detail.html"

class ResultsView(generic.DetailView):
    model = Question
    template_name = "polls/results.html"



def index(request):
    latest_question_list = Question.objects.order_by("-pub_date")[:5]
    # Mapping template variable to python objects
    context = {
            "latest_question_list":latest_question_list,
    }

    # render(request object, tempalte name, dictionary )
    # returns an HttpResponse object of given template rendered with
    # the given context
    return render(request, "polls/index.html",context)


# When user request a page like localhost:8000/polls/5
# input request is <HttpRequest object>, question_id is 5
def detail(request, question_id):

    # return a HttpResponse object containing content for 
    # requested page or Http404
    # return HttpResponse("You're looking at question %s" % question_id)

    # Similar to question = Question.objects.get(pk=question_id)
    question = get_object_or_404(Question,pk = question_id)

    return render(request, "polls/detail.html",{"question":question})



def results(request, question_id):
    question = get_object_or_404(Question, pk=question_id)
    return render(request, "polls/results.html", {"question":question})

def vote(request, question_id):
    question = get_object_or_404(Question, pk=question_id)
    try:
        # Note on the input of function request
        # request.POST access submitted data (from user) by key name.
        selected_choice = question.choice_set.get(pk=request.POST["choice"])
    except (KeyError,Choice.DoesNotExist):
        return render(
                request,
                "polls/detail.html",
                {
                    "question": question,
                    "error_message": "You didn't select a choice",
                }
        )
    else:
        selected_choice.votes += 1
        selected_choice.save()
        
        # Return a HttpResponseRedirect after dealing with POST data to prevent
        # data being posted twice if a user hits back button
        # Takes URL to which user will be redirected as input

        # reverse return a string like /polls/3/results/
        return HttpResponseRedirect(reverse("polls:results", args=(question.id,)))



