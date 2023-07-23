

#include <stdio.h>
#include <stdlib.h>


typedef struct list_t {
    int val;
    struct list_t* next;
}list_t;


list_t* set_list(int*, int len );
list_t* reverse_list(list_t*);
void insert_head(list_t**, int);

list_t* set_list(int* arr, int len)
{
    list_t* start = malloc(sizeof(list_t));
    start->val = arr[0];

    list_t* cur = start;
    for(int i = 1; i < len; i++)
    {
        cur->next = malloc(sizeof(list_t));
        cur->next->val = arr[i];
        cur = cur->next;
    }

    return start;
}

list_t* reverse_list(list_t* list)
{
    list_t* cur = list;
    list_t* next = cur->next;
    cur->next = NULL;

    while(next)
    {
        list_t* new_next = next->next;
        next->next = cur;
        cur = next;
        next = new_next;
    }

    return cur;
}

void insert_head(list_t** head, int val)
{
    list_t* new_head = malloc(sizeof(list_t));
    new_head->val = val;
    new_head->next = *head;
    *head = new_head;
}



int main()
{

    int arr[] = {3,4,5,7,10,1};
    int len = 6;
    list_t* list = set_list(arr, len);
    // list = reverse_list(list);
    insert_head(&list, 2);

    list_t* elt = list;
    while(elt != NULL)
    {
        printf("%d ", elt->val);
        elt = elt->next;
    }
    printf("\n");


}
